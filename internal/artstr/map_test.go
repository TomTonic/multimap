package artstr

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
	"testing"
)

// model is the multimap the tree must agree with: sorted keys, each with a set
// of values.
type model map[string]map[string]bool

func (m model) add(k, v string) {
	if m[k] == nil {
		m[k] = map[string]bool{}
	}
	m[k][v] = true
}

func (m model) remove(k, v string) {
	delete(m[k], v)
	if len(m[k]) == 0 {
		delete(m, k)
	}
}

func (m model) sorted() []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

// keyShapes are generators of keys for the differential tests: short keys that
// collide a lot, text paths, keys with long common prefixes, long keys that no
// page holds, and the empty key.
var keyShapes = map[string]func(r *rand.Rand) string{
	"short": func(r *rand.Rand) string {
		b := make([]byte, 1+r.IntN(3))
		for i := range b {
			b[i] = "abc"[r.IntN(3)]
		}
		return string(b)
	},
	"names": func(r *rand.Rand) string {
		w := []string{"main", "street", "oak", "elm", "road", "lane", "north", "south"}
		return w[r.IntN(len(w))] + " " + w[r.IntN(len(w))] + fmt.Sprint(r.IntN(40))
	},
	"paths": func(r *rand.Rand) string {
		d := []string{"usr", "lib", "share", "doc", "python3", "dist-packages"}
		var sb strings.Builder
		for range 1 + r.IntN(6) {
			sb.WriteString("/" + d[r.IntN(len(d))])
		}
		return sb.String() + fmt.Sprint(r.IntN(30))
	},
	"binary": func(r *rand.Rand) string {
		b := make([]byte, r.IntN(12))
		for i := range b {
			b[i] = byte(r.IntN(256))
		}
		return string(b)
	},
	"long": func(r *rand.Rand) string {
		return strings.Repeat("p", 200+r.IntN(80)) + fmt.Sprint(r.IntN(100))
	},
	"mixed": func(r *rand.Rand) string {
		switch r.IntN(4) {
		case 0:
			return ""
		case 1:
			return strings.Repeat("q", r.IntN(300))
		case 2:
			return "k" + fmt.Sprint(r.IntN(50))
		}
		return "k" + fmt.Sprint(r.IntN(5)) + strings.Repeat("z", r.IntN(70))
	},
}

// valueShapes are generators of values: short names, equal lengths, empty
// strings and strings that no page holds.
var valueShapes = map[string]func(r *rand.Rand) string{
	"names": func(r *rand.Rand) string {
		return []string{"Aachen", "Bonn", "Köln", "Ulm", "x", "", "Frankfurt am Main", "Bad Homburg v. d. Höhe"}[r.IntN(8)]
	},
	"fixed": func(r *rand.Rand) string { return fmt.Sprintf("%016x", r.IntN(500)) },
	"mixed": func(r *rand.Rand) string {
		if r.IntN(30) == 0 {
			return strings.Repeat("v", 250+r.IntN(300))
		}
		return strings.Repeat("v", r.IntN(30)) + fmt.Sprint(r.IntN(9))
	},
}

// verify checks every observable thing of m against the model: the number of
// keys, each key's values, the ordered keys and values of Range and RangeValues
// over everything and over random bounds.
func verify(t *testing.T, m *Map[string], want model, r *rand.Rand, step int) {
	t.Helper()
	if m.Len() != len(want) {
		t.Fatalf("step %d: Len = %d, want %d", step, m.Len(), len(want))
	}
	for k, vs := range want {
		got := map[string]bool{}
		m.Each([]byte(k), func(v string) bool { got[v] = true; return true })
		if !maps(got, vs) {
			t.Fatalf("step %d: key %q has %v, want %v", step, k, got, vs)
		}
	}
	keys := want.sorted()
	rangeCheck(t, m, want, keys, &Bounds{}, step)
	for range 6 {
		var b Bounds
		if len(keys) > 0 {
			b.From, b.HasFrom, b.FromIncl = []byte(keys[r.IntN(len(keys))]), true, r.IntN(2) == 0
			b.To, b.HasTo, b.ToIncl = []byte(keys[r.IntN(len(keys))]), true, r.IntN(2) == 0
			if r.IntN(3) == 0 {
				b.HasFrom = false
			}
			if r.IntN(3) == 0 {
				b.HasTo = false
			}
			if r.IntN(4) == 0 {
				b.From = append(b.From, byte(r.IntN(256)))
			}
		}
		rangeCheck(t, m, want, keys, &b, step)
	}
}

func maps(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func rangeCheck(t *testing.T, m *Map[string], want model, keys []string, b *Bounds, step int) {
	t.Helper()
	var inRange []string
	for _, k := range keys {
		if b.Contains([]byte(k)) {
			inRange = append(inRange, k)
		}
	}
	var gotKeys []string
	m.Range(b, func(k []byte) bool { gotKeys = append(gotKeys, string(k)); return true })
	if !slices.Equal(gotKeys, inRange) {
		t.Fatalf("step %d: Range %+v = %q, want %q", step, b, gotKeys, inRange)
	}
	var gotVals, wantVals []string
	m.RangeValues(b, func(v string) bool { gotVals = append(gotVals, v); return true })
	for _, k := range inRange {
		for v := range want[k] {
			wantVals = append(wantVals, v)
		}
	}
	slices.Sort(gotVals)
	slices.Sort(wantVals)
	if !slices.Equal(gotVals, wantVals) {
		t.Fatalf("step %d: RangeValues %+v has %d values, want %d", step, b, len(gotVals), len(wantVals))
	}
}

// TestMapAgainstModel checks the multimap of strings, whose single-valued keys
// live in length-header pages, against a plain model.
//
// A user adds and removes values of many shapes of keys and values: short keys
// that crowd the pages, names, paths, binary keys, keys that no page can hold,
// the empty key, and values that are empty, of one length or too long for a
// page. After every few operations the tree must hold exactly the keys and
// values of the model, find them all and list them in order within any bounds.
func TestMapAgainstModel(t *testing.T) {
	for _, mode := range []struct{ zeroCopy, pairs bool }{{false, false}, {true, false}, {false, true}, {true, true}} {
		for kn, kf := range keyShapes {
			for vn, vf := range valueShapes {
				t.Run(fmt.Sprintf("zerocopy %v pairs %v/%s/%s", mode.zeroCopy, mode.pairs, kn, vn), func(t *testing.T) {
					testModel(t, mode.zeroCopy, mode.pairs, kf, vf, uint64(len(kn))+seedBase, uint64(len(vn))+7)
				})
			}
		}
	}
}

// testModel runs one workload of TestMapAgainstModel.
func testModel(t *testing.T, zeroCopy, pairs bool, kf, vf func(*rand.Rand) string, seed1, seed2 uint64) {
	r := rand.New(rand.NewPCG(seed1, seed2))
	m := Map[string]{ZeroCopy: zeroCopy, Pairs: pairs}
	want := model{}
	var held []string // strings a user holds on to while the map changes
	var heldBytes []string
	var added [][2]string
	var log []string // the last operations, for a failure
	for step := 1; step <= mapSteps; step++ {
		switch op := r.IntN(10); {
		case op < 6:
			k, v := kf(r), vf(r)
			log = append(log, fmt.Sprintf("Add(%q, %q)", k, v))
			m.Add([]byte(k), v)
			want.add(k, v)
			added = append(added, [2]string{k, v})
		case op < 9 && len(added) > 0:
			kv := added[r.IntN(len(added))]
			log = append(log, fmt.Sprintf("Remove(%q, %q)", kv[0], kv[1]))
			m.Remove([]byte(kv[0]), kv[1])
			want.remove(kv[0], kv[1])
		case len(added) > 0:
			kv := added[r.IntN(len(added))]
			log = append(log, fmt.Sprintf("RemoveKey(%q)", kv[0]))
			m.RemoveKey([]byte(kv[0]))
			delete(want, kv[0])
		}
		if m.Len() != len(want) {
			t.Fatalf("step %d: Len = %d, want %d after %q", step, m.Len(), len(want), log[max(0, len(log)-5):])
		}
		if step%250 == 0 {
			verify(t, &m, want, r, step)
		}
		if step%50 == 0 && len(added) > 0 {
			kv := added[r.IntN(len(added))]
			m.Each([]byte(kv[0]), func(v string) bool {
				held, heldBytes = append(held, v), append(heldBytes, strings.Clone(v))
				return false
			})
		}
	}
	verify(t, &m, want, r, -1)
	for _, k := range want.sorted() {
		if !m.Has([]byte(k)) {
			t.Fatalf("Has(%q) is false", k)
		}
	}
	for i, s := range held {
		if s != heldBytes[i] {
			t.Fatalf("a string handed out changed from %q to %q", heldBytes[i], s)
		}
	}
	m.Clear()
	if m.Len() != 0 || m.Has([]byte("a")) {
		t.Fatal("Clear leaves keys")
	}
}

// TestMapPagesAreUsed checks that single-valued string keys do live in pages,
// and move to leaves with a second value.
//
// A user who loads names with one value each should get the compact pages, and
// keep getting correct answers once some names gain values.
func TestMapPagesAreUsed(t *testing.T) {
	var m Map[string]
	for i := range 2000 {
		m.Add([]byte(fmt.Sprintf("key-%05d", i)), fmt.Sprintf("value %d", i))
	}
	count := func() (pages, leaves int) {
		m.Objects(func(o Object) {
			switch o.Label {
			case "page":
				pages++
			case "typed leaf", "set leaf":
				leaves++
			}
		})
		return
	}
	if p, l := count(); p < 100 || l != 0 {
		t.Fatalf("2000 single-valued keys: %d pages, %d leaves", p, l)
	}
	m.Add([]byte("key-00007"), "second")
	var got []string
	m.Each([]byte("key-00007"), func(v string) bool { got = append(got, v); return true })
	slices.Sort(got)
	if !slices.Equal(got, []string{"second", "value 7"}) {
		t.Fatalf("key with two values has %q", got)
	}
	if _, l := count(); l == 0 {
		t.Fatal("a key with two values has no leaf")
	}
	var buf bytes.Buffer
	m.Range(&Bounds{From: []byte("key-00005"), HasFrom: true, FromIncl: true, To: []byte("key-00009"), HasTo: true}, func(k []byte) bool {
		buf.Write(k)
		buf.WriteByte(' ')
		return true
	})
	if buf.String() != "key-00005 key-00006 key-00007 key-00008 " {
		t.Fatalf("range = %q", buf.String())
	}
}

// seedBase and mapSteps let a longer run of TestMapAgainstModel vary the
// workloads (go test -run Model -args is not needed: edit and rerun).
var (
	seedBase uint64
	mapSteps = 4000
)

// TestRemoveLastKeyBelowRangeNode checks that a range node whose only page is
// emptied disappears.
//
// When a key is removed and a range node is left with one page below it and no
// term, the tree moves the page up if its keys still fit a page with the node's
// path in front of them, and leaves the node standing if they do not (keys of 300
// bytes and more). Removing the last key of that page then leaves a node with
// nothing in it, which must vanish and not be treated as a node with a term.
func TestRemoveLastKeyBelowRangeNode(t *testing.T) {
	var m Map[string]
	m.Add([]byte("seed"), "v") // decides the map's kind
	m.RemoveKey([]byte("seed"))
	path := bytes.Repeat([]byte{'p'}, 200)
	key := append(bytes.Clone(path), bytes.Repeat([]byte{'k'}, 100)...)
	m.t.root = makeR(path, nil, []rng{{0, pageHdr(newPageFor(key, len(path), "value"))}})
	m.t.size = 1
	if got := collect(&m, key); !slices.Equal(got, []string{"value"}) {
		t.Fatalf("the key has %q", got)
	}
	m.RemoveKey(key)
	if m.Len() != 0 || m.Has(key) || m.t.root != nil {
		t.Fatalf("Len %d, root %v after removing the only key", m.Len(), m.t.root)
	}
}

func collect(m *Map[string], key []byte) []string {
	var out []string
	m.Each(key, func(v string) bool { out = append(out, v); return true })
	return out
}

// TestRemoveLongKeysOneByOne is the same check as in internal/art, through the
// operations a user has: two keys of 300 bytes with a common beginning, removed one
// after the other.
func TestRemoveLongKeysOneByOne(t *testing.T) {
	for _, pairs := range []bool{false, true} {
		m := Map[string]{Pairs: pairs}
		prefix := bytes.Repeat([]byte{'p'}, 100)
		k1 := append(append(bytes.Clone(prefix), '1'), bytes.Repeat([]byte{'x'}, 199)...)
		k2 := append(append(bytes.Clone(prefix), '2'), bytes.Repeat([]byte{'x'}, 199)...)
		m.Add(k1, "1")
		m.Add(k2, "2")
		m.RemoveKey(k1)
		m.RemoveKey(k2)
		if m.Len() != 0 || m.Has(k1) || m.Has(k2) {
			t.Fatalf("pairs %v: Len %d after removing both keys", pairs, m.Len())
		}
	}
}
