package art

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"

	"github.com/TomTonic/multimap/internal/vset"
)

type valset = vset.Set[uint64]

// keySets returns key corpora that exercise every structural case: keys that
// are prefixes of other keys, the empty key, keys longer than the 8 inline
// path bytes and the 16 inline leaf bytes, zero bytes, dense and sparse
// integers (which drive nodes through every kind), and shared string prefixes.
func keySets() map[string][][]byte {
	r := rand.New(rand.NewPCG(1, 2))
	sets := map[string][][]byte{}
	var u64, dense [][]byte
	for range 5000 {
		u64 = append(u64, binary.BigEndian.AppendUint64(nil, r.Uint64()))
	}
	for i := range uint64(20000) {
		dense = append(dense, binary.BigEndian.AppendUint64(nil, i))
	}
	sets["u64-random"], sets["u64-dense"] = u64, dense

	var prefixes [][]byte
	base := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	for i := 0; i <= len(base); i++ {
		prefixes = append(prefixes, base[:i:i],
			append(base[:i:i], 0), append(base[:i:i], 0xff, 'x'))
	}
	sets["prefix-chains"] = prefixes

	var strs [][]byte
	words := []string{"user", "users", "item", "items", "a", "", "product", "productcatalogue"}
	for range 8000 {
		strs = append(strs, fmt.Appendf(nil, "%s/%s/%s/%d", words[r.IntN(len(words))],
			words[r.IntN(len(words))], words[r.IntN(len(words))], r.IntN(300)))
	}
	sets["strings"] = strs

	var wide [][]byte
	long := []byte("a-compressed-path-longer-than-sixteen-bytes/")
	for _, fan := range []int{3, 10, 24, 56, 256} {
		p := append(append([]byte(nil), long...), byte(fan))
		for b := range fan {
			wide = append(wide, append(append(p[:len(p):len(p)], byte(b)), "tail"...))
		}
	}
	sets["long-prefix-wide"] = append(wide, long[:20], append(long[:30:30], 'Z'))

	var small [][]byte
	for range 6000 {
		k := make([]byte, r.IntN(12))
		for i := range k {
			k[i] = byte(r.IntN(3))
		}
		small = append(small, k)
	}
	sets["random-small-alphabet"] = small
	return sets
}

// reference is the trivially correct multimap the tree is compared with.
type reference map[string]map[uint64]bool

func (r reference) add(k []byte, v uint64) {
	if r[string(k)] == nil {
		r[string(k)] = map[uint64]bool{}
	}
	r[string(k)][v] = true
}

func (r reference) remove(k []byte, v uint64) {
	if s := r[string(k)]; s != nil {
		delete(s, v)
		if len(s) == 0 {
			delete(r, string(k))
		}
	}
}

func (r reference) sortedKeys() []string {
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// TestAgainstReference checks that the ordered multimap behind
// multimap.Ordered stores, finds, removes and ranges over keys and values
// exactly like a trivially correct reference, through phases of growth and
// of heavy deletion, and that after every phase the tree has the shape its
// invariants demand (see checkInvariants).
func TestAgainstReference(t *testing.T) {
	for name, keys := range keySets() {
		t.Run(name, func(t *testing.T) {
			r := rand.New(rand.NewPCG(7, 8))
			var m Map[uint64]
			ref := reference{}
			for phase := range 6 {
				removing := phase%2 == 1
				for _, k := range keys {
					switch op := r.IntN(10); {
					case removing && op < 3:
						m.RemoveKey(k)
						delete(ref, string(k))
					case removing && op < 7:
						v := uint64(r.IntN(8))
						m.Remove(k, v)
						ref.remove(k, v)
					case !removing || op < 8:
						for range 1 + r.IntN(4) {
							v := uint64(r.IntN(8))
							if r.IntN(30) == 0 {
								v = uint64(r.IntN(200)) // spill into array and hash
							}
							m.Add(k, v)
							ref.add(k, v)
						}
					}
				}
				compare(t, &m, ref, r)
				checkInvariants(t, &m.t)
			}
		})
	}
}

func compare(t *testing.T, m *Map[uint64], ref reference, r *rand.Rand) {
	t.Helper()
	if m.Len() != len(ref) {
		t.Fatalf("Len = %d, want %d", m.Len(), len(ref))
	}
	for k, want := range ref {
		s := m.Values([]byte(k))
		if s == nil || s.Len() != len(want) {
			t.Fatalf("Values(%q) has wrong size", k)
		}
		s.Each(func(v uint64) bool {
			if !want[v] {
				t.Fatalf("Values(%q) holds %d unexpectedly", k, v)
			}
			return true
		})
	}
	sorted := ref.sortedKeys()
	if len(sorted) == 0 {
		return
	}
	for i := range 300 {
		b := randomBounds(r, sorted, i)
		checkRange(t, m, ref, sorted, b)
	}
	checkRange(t, m, ref, sorted, &Bounds{}) // everything
}

func randomBounds(r *rand.Rand, sorted []string, i int) *Bounds {
	pick := func() []byte {
		k := []byte(sorted[r.IntN(len(sorted))])
		switch r.IntN(4) {
		case 0:
			return append(k, 0) // just above a key
		case 1:
			return k[:r.IntN(len(k)+1)] // a prefix of a key
		}
		return k
	}
	b := &Bounds{From: pick(), To: pick(), HasFrom: r.IntN(4) > 0, HasTo: r.IntN(4) > 0,
		FromIncl: r.IntN(2) == 0, ToIncl: r.IntN(2) == 0}
	if i%5 == 0 {
		b.To = b.From // single key or empty
	}
	return b
}

func inRange(k string, b *Bounds) bool {
	if b.HasFrom {
		if c := bytes.Compare([]byte(k), b.From); c < 0 || (c == 0 && !b.FromIncl) {
			return false
		}
	}
	if b.HasTo {
		if c := bytes.Compare([]byte(k), b.To); c > 0 || (c == 0 && !b.ToIncl) {
			return false
		}
	}
	return true
}

func checkRange(t *testing.T, m *Map[uint64], ref reference, sorted []string, b *Bounds) {
	t.Helper()
	var want []string
	for _, k := range sorted {
		if inRange(k, b) {
			want = append(want, k)
		}
	}
	var got []string
	m.Range(b, func(k []byte, s *valset) bool {
		if s.Len() != len(ref[string(k)]) {
			t.Fatalf("Range: key %q has %d values, want %d", k, s.Len(), len(ref[string(k)]))
		}
		got = append(got, string(k))
		return true
	})
	if !slices.Equal(got, want) {
		t.Fatalf("Range(%+v) yielded %d keys, want %d\n got=%q\nwant=%q", *b, len(got), len(want), got, want)
	}
	// stopping early must stop
	n := 0
	m.Range(b, func([]byte, *valset) bool { n++; return n < 3 })
	if n != min(3, len(want)) {
		t.Fatalf("Range did not stop when asked: %d calls", n)
	}
}

// TestLeafTail makes sure that range queries touch ahead only memory that
// belongs to the leaf, for any value type. It covers the scan of the ART behind
// multimap.Ordered, which reads the last byte of every leaf it is about to
// visit: that byte must be the leaf's last one for uint64 values (80 B leaf)
// and for string values (104 B leaf), whose value sets differ in size.
func TestLeafTail(t *testing.T) {
	if got := leafTail[uint64](); got != 79 {
		t.Fatalf("leafTail[uint64] = %d, want 79", got)
	}
	if got := leafTail[string](); got != 103 {
		t.Fatalf("leafTail[string] = %d, want 103", got)
	}
}

// checkInvariants walks the whole tree and fails on any node that violates
// the structure insert and delete must maintain.
func checkInvariants(t *testing.T, tr *Tree) {
	t.Helper()
	if tr.root == nil {
		if tr.size != 0 {
			t.Fatalf("empty tree with size %d", tr.size)
		}
		return
	}
	if n := checkNode(t, tr.root, 0); n != tr.size {
		t.Fatalf("tree holds %d leaves, size says %d", n, tr.size)
	}
}

// checkNode checks the subtree n whose path starts at key depth depth and
// returns its number of leaves.
func checkNode(t *testing.T, n *header, depth int) int {
	t.Helper()
	if n.kind == kLeaf {
		if len(asLeaf(n).key()) < depth {
			t.Fatalf("leaf %q is shorter than its depth %d", asLeaf(n).key(), depth)
		}
		return 1
	}
	limits := map[kind][2]int{kN4: {1, 4}, kN11: {shrink11 + 1, 11}, kN25: {shrink25 + 1, 25},
		kN57: {shrink57 + 1, 57}, kN256: {shrink256 + 1, 256}}[n.kind]
	if lo, hi := limits[0], limits[1]; int(n.count) < lo || int(n.count) > hi {
		t.Fatalf("kind %d holds %d children, allowed %d..%d", n.kind, n.count, lo, hi)
	}
	if int(n.count)+b2i(n.term != nil) < 2 {
		t.Fatalf("node does not branch (count %d, term %v): it should have collapsed", n.count, n.term != nil)
	}
	end := depth + int(n.plen)
	checkKey := func(k []byte) {
		if len(k) < end || !bytes.Equal(k[depth:depth+min(int(n.plen), 8)], n.prefix[:min(n.plen, 8)]) {
			t.Fatalf("key %q does not match the node path at depth %d (plen %d)", k, depth, n.plen)
		}
	}
	leaves := 0
	if n.term != nil {
		checkKey(n.term.key())
		if len(n.term.key()) != end {
			t.Fatalf("term key %q does not end at depth %d", n.term.key(), end)
		}
		leaves++
	}
	children, bytesOf := 0, -1
	eachChild(n, func(b byte, c *header) {
		if int(b) <= bytesOf {
			t.Fatalf("child bytes not ascending")
		}
		bytesOf = int(b)
		children++
		walkLeaves(c, func(l *leafHead) {
			k := l.key()
			checkKey(k)
			if len(k) <= end || k[end] != b {
				t.Fatalf("key %q sits under child byte %d at depth %d", k, b, end)
			}
		})
		leaves += checkNode(t, c, end+1)
	})
	if children != int(n.count) {
		t.Fatalf("count %d but %d children", n.count, children)
	}
	return leaves
}

func eachChild(n *header, fn func(byte, *header)) {
	switch n.kind {
	case kN25, kN57:
		bm, child := bitmapOf(n)
		i := 0
		for k := range 256 {
			if (bm[k>>6]>>(k&63))&1 == 1 {
				fn(byte(k), child[i])
				i++
			}
		}
	case kN256:
		for k, c := range asN256(n).child {
			if c != nil {
				fn(byte(k), c)
			}
		}
	default:
		keys, child := sorted(n)
		for i, k := range keys {
			fn(k, child[i])
		}
	}
}

func walkLeaves(n *header, fn func(*leafHead)) {
	if n.kind == kLeaf {
		fn(asLeaf(n))
		return
	}
	if n.term != nil {
		fn(n.term)
	}
	eachChild(n, func(_ byte, c *header) { walkLeaves(c, fn) })
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// TestShrinkAndCollapse checks that deleting keys one by one takes every node
// kind back down through each smaller kind to nothing, collapsing and
// re-merging compressed paths (short and longer than 8 bytes) on the way, and
// that the tree satisfies its invariants after every single delete.
func TestShrinkAndCollapse(t *testing.T) {
	for _, prefix := range []string{"", "p", "a-path-longer-than-eight-bytes"} {
		for _, fan := range []int{2, 4, 11, 25, 57, 256} {
			t.Run(fmt.Sprintf("%q/%d", prefix, fan), func(t *testing.T) {
				r := rand.New(rand.NewPCG(uint64(fan), 9))
				var m Map[uint64]
				var keys [][]byte
				for b := range fan {
					for _, tail := range []string{"", "x", "xy-longer-tail-than-16"} {
						k := append([]byte(prefix), byte(b))
						keys = append(keys, append(k, tail...))
					}
				}
				for _, k := range keys {
					m.Add(k, 1)
				}
				checkInvariants(t, &m.t)
				r.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
				for i, k := range keys {
					m.RemoveKey(k)
					checkInvariants(t, &m.t)
					if m.Len() != len(keys)-i-1 {
						t.Fatalf("Len = %d after %d deletes", m.Len(), i+1)
					}
					for _, other := range keys[i+1:] {
						if m.Values(other) == nil {
							t.Fatalf("deleting %q lost %q", k, other)
						}
					}
				}
				if m.t.root != nil {
					t.Fatalf("tree not empty after deleting every key")
				}
				m.Add(keys[0], 1)
				m.Clear()
				if m.Len() != 0 || m.Values(keys[0]) != nil {
					t.Fatalf("Clear left keys behind")
				}
			})
		}
	}
}

// FuzzOperations drives the tree with arbitrary operation sequences over
// short keys from a tiny alphabet (which maximises shared paths, splits and
// merges), checking every result against the reference and the structural
// invariants at the end.
func FuzzOperations(f *testing.F) {
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	f.Add(bytes.Repeat([]byte{3, 0, 1, 2, 7, 1, 0, 0, 5}, 30))
	f.Fuzz(func(t *testing.T, ops []byte) {
		var m Map[uint64]
		ref := reference{}
		for len(ops) >= 2 {
			op, n := ops[0], int(ops[1]%6)
			ops = ops[2:]
			if len(ops) < n {
				break
			}
			k := make([]byte, n)
			for i := range k {
				k[i] = ops[i] % 4
			}
			ops = ops[n:]
			v := uint64(op >> 2)
			switch op % 4 {
			case 0, 1:
				m.Add(k, v)
				ref.add(k, v)
			case 2:
				m.Remove(k, v)
				ref.remove(k, v)
			default:
				m.RemoveKey(k)
				delete(ref, string(k))
			}
		}
		compare(t, &m, ref, rand.New(rand.NewPCG(1, 1)))
		checkInvariants(t, &m.t)
	})
}
