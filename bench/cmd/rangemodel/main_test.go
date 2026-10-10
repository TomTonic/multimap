package main

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

// holds counts the keys a subtree holds and checks that a node child of a range covers exactly its byte.
func holds(t *testing.T, m *model, o obj) int {
	t.Helper()
	switch o.k {
	case kPage:
		return len(o.p.items)
	case kBig:
		return 1
	case kNode:
		n := o.n
		k := holds(t, m, n.end)
		for i, c := range n.kids {
			if m.ranges && c.k == kNode {
				hi := 256
				if i+1 < len(n.starts) {
					hi = n.starts[i+1]
				}
				if hi != n.starts[i]+1 {
					t.Fatalf("a node child covers the bytes %d to %d", n.starts[i], hi-1)
				}
			}
			k += holds(t, m, c)
		}
		if !slices.IsSorted(n.starts) {
			t.Fatalf("the starts of a node are not sorted: %v", n.starts)
		}
		return k
	}
	return 0
}

// TestModelFindsEveryKey makes sure that the model of the routing layer counts real trees.
//
// The model of step 6 (docs/redesign/step6-design.md) decides whether byte nodes with range children are built: it
// builds today's tree and the proposal's from a corpus and counts what a lookup of every key touches. A tree that
// loses a key, or a range node whose node child covers more than its byte, would count something else.
//
// Expected: for keys of shared prefixes, of one byte and of many values (some too many for a page, so that they
// get an object of their own), inserted in random order, both trees under every limit hold every key once, find
// every key, and keep the starts of a node sorted and a node child's range at one byte.
func TestModelFindsEveryKey(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	var items []*item
	seen := map[string]bool{}
	for len(items) < 3000 {
		var b strings.Builder
		b.WriteString([]string{"/usr/share/doc/", "/usr/lib/", "/etc/", "a", ""}[r.IntN(5)])
		for range r.IntN(12) {
			b.WriteByte("abcde/"[r.IntN(6)])
		}
		k := b.String()
		if seen[k] {
			continue
		}
		seen[k] = true
		v := 1 + r.IntN(3)
		if r.IntN(100) == 0 {
			v = 200 // a key whose values need an object of its own
		}
		items = append(items, &item{key: []byte(k), vals: v, vb: 8 * v})
	}
	for _, m := range []*model{{limit: 512}, {ranges: true, limit: 512}, {ranges: true, limit: 256}, {ranges: true, limit: 128}} {
		t.Run(fmt.Sprintf("ranges %v limit %d", m.ranges, m.limit), func(t *testing.T) {
			for _, it := range items {
				m.insert(it)
			}
			if got := holds(t, m, m.root); got != len(items) {
				t.Fatalf("the tree holds %d keys, want %d", got, len(items))
			}
			var c counts
			for _, it := range items {
				m.lookup(it, &c)
			}
			if c.lookups != len(items) || c.cmpLin < len(items) || c.linesLin < len(items) {
				t.Fatalf("counts %+v", c)
			}
		})
	}
}

// TestRunPrintsTheCases makes sure the model's table has a row for every tree of every case asked for.
//
// The rows are what the design note of step 6 quotes; a case that is missing would leave a gap in the decision.
//
// Expected: for one key kind, one profile and one size with two limits, three rows (today and two limits).
func TestRunPrintsTheCases(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "street", "-values", "single-value", "-sizes", "4096", "-limits", "512,256"}); err != nil {
		t.Fatal(err)
	}
	if n := strings.Count(out.String(), "| street single-value 4096 |"); n != 3 {
		t.Fatalf("%d rows, want 3:\n%s", n, out.String())
	}
}
