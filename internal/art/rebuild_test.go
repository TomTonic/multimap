package art

import (
	"encoding/binary"
	"fmt"
	"testing"
)

// shape counts the objects of a tree by what they are.
type shape struct{ pages, ranges, inner, leaves int }

func shapeOf(n *header) shape {
	var s shape
	var walk func(n *header)
	walk = func(n *header) {
		switch {
		case n == nil:
			return
		case isLeaf(n.kind):
			s.leaves++
			return
		case isPage(n.kind):
			s.pages++
			return
		case isRange(n.kind):
			s.ranges++
			for _, c := range asR(n).children()[:asR(n).n] {
				walk(c)
			}
		default:
			s.inner++
			eachInner(n, func(_ byte, c *header) { walk(c) })
		}
		if termOf(n) != nil {
			s.leaves++
		}
	}
	walk(n)
	return s
}

// TestFallback makes sure a multimap whose keys hold several values keeps
// the structure that is fast for them, without being told. It covers the
// key index of Ordered: a tree starts with pages for keys with one value
// and rebuilds a subtree from inner nodes and leaves once keys with several
// values crowd it (rebuild.go), and it keeps its pages where they do not.
// Every case checks the tree's invariants and its contents.
func TestFallback(t *testing.T) {
	u64 := func(i int) []byte { return binary.BigEndian.AppendUint64(nil, uint64(i)) }
	tests := []struct {
		name  string
		keys  func(i int) []byte
		n     int
		multi func(i int) bool // the keys that get a second value
		check func(t *testing.T, s shape)
	}{
		{
			name: "falls back entirely when every key holds two values", keys: u64, n: 5000,
			multi: func(int) bool { return true },
			check: func(t *testing.T, s shape) {
				if s.pages != 0 || s.ranges != 0 {
					t.Errorf("%+v: pages or range nodes left", s)
				}
			},
		},
		{
			name: "keeps its pages when one key in a hundred holds two values", keys: u64, n: 5000,
			multi: func(i int) bool { return i%100 == 0 },
			check: func(t *testing.T, s shape) {
				if s.pages < 5000/32 || s.leaves > 5000/100 {
					t.Errorf("%+v: pages lost", s)
				}
			},
		},
		{
			name: "falls back only where keys hold several values",
			keys: func(i int) []byte { return fmt.Appendf(nil, "%c%05d", 'a'+i%2, i) }, n: 4000,
			multi: func(i int) bool { return i%2 == 1 },
			check: func(t *testing.T, s shape) {
				if s.pages < 2000/32 || s.inner == 0 || s.leaves < 2000 {
					t.Errorf("%+v: want the pages of the a-keys and the inner nodes of the b-keys", s)
				}
			},
		},
		{
			name: "rebuilds nodes of every width, and keys that end inside others",
			keys: func(i int) []byte {
				// groups of 3, 10, 20, 50 and 200 keys below one byte each,
				// and every group's byte as a key of its own
				for g, w := range []int{4, 11, 21, 51, 201} {
					if i < w {
						return []byte{byte(g), byte(i)}[:min(i, 1)+1]
					}
					i -= w
				}
				panic("out of keys")
			},
			n:     4 + 11 + 21 + 51 + 201,
			multi: func(int) bool { return true },
			check: func(t *testing.T, s shape) {
				if s.pages != 0 || s.ranges != 0 || s.inner != 6 {
					t.Errorf("%+v: want one inner node per group and the root", s)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, ref := &Map[uint64]{}, reference{}
			for i := range tt.n {
				m.Add(tt.keys(i), uint64(i))
				ref.add(tt.keys(i), uint64(i))
			}
			for i := range tt.n {
				if tt.multi(i) {
					m.Add(tt.keys(i), uint64(i)+1<<32)
					ref.add(tt.keys(i), uint64(i)+1<<32)
				}
			}
			checkInvariants(t, &m.t)
			if m.Len() != len(ref) {
				t.Fatalf("Len %d, want %d", m.Len(), len(ref))
			}
			for k, vs := range ref {
				if got := valuesOf(m, []byte(k)); len(got) != len(vs) {
					t.Fatalf("key %q: %d values, want %d", k, len(got), len(vs))
				}
			}
			tt.check(t, shapeOf(m.t.root))
		})
	}
}

// TestFallbackKeepsTerm makes sure a key that is a prefix of other keys
// keeps all its values when it gets a second one. It covers the fallback of
// the key index (rebuild.go) where the key's leaf ends at a range node
// instead of below it: the check starts at that range node.
func TestFallbackKeepsTerm(t *testing.T) {
	m := &Map[uint64]{}
	for _, k := range []string{"abc", "abd", "ab"} { // the first key's length is the pages' key length
		m.Add([]byte(k), 1)
	}
	m.Add([]byte("ab"), 2)
	checkInvariants(t, &m.t)
	if s := shapeOf(m.t.root); !isRange(m.t.root.kind) || termOf(m.t.root) == nil || s.pages != 1 {
		t.Fatalf("%+v: want a range node with the term and one page", s)
	}
	m.Add([]byte("abc"), 2)
	checkInvariants(t, &m.t)
	if s := shapeOf(m.t.root); s.pages != 0 || s.ranges != 0 || len(valuesOf(m, []byte("ab"))) != 2 {
		t.Fatalf("%+v: want the term kept in an inner node", s)
	}
}
