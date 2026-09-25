package art

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
	"unsafe"
)

// keySets returns key corpora that exercise every structural case: keys that
// are prefixes of other keys, the empty key, keys longer than the 8 inline
// path bytes, keys of every length from 0 to 99 (every size class of the
// leaves and beyond), zero bytes, dense and sparse integers (which drive nodes
// through every kind), and shared string prefixes.
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

	var lengths [][]byte
	for range 4000 {
		k := make([]byte, r.IntN(100))
		for i := range k {
			k[i] = 'a' + byte(r.IntN(4))
		}
		lengths = append(lengths, k)
	}
	sets["every-length"] = lengths
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
		if !s.Found() || s.Len() != len(want) {
			t.Fatalf("Values(%q) has wrong size", k)
		}
		s.Each(func(v uint64) bool {
			if !want[v] {
				t.Fatalf("Values(%q) holds %d unexpectedly", k, v)
			}
			return true
		})
	}
	for k := range ref {
		for _, probe := range nearMisses([]byte(k)) {
			_, want := ref[string(probe)]
			if got := m.Values(probe).Found(); got != want {
				t.Fatalf("Values(%q) found = %v, want %v", probe, got, want)
			}
		}
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

// nearMisses returns keys that differ from k only slightly: shorter, longer,
// or with one byte changed at the end, in the middle, or at positions 9 and
// 12, which lie beyond the 8 path bytes a node checks itself. The last byte
// is also changed in its top bit, which a dense node rarely holds.
func nearMisses(k []byte) [][]byte {
	out := [][]byte{append(bytes.Clone(k), 0)}
	if len(k) > 0 {
		x := bytes.Clone(k)
		x[len(x)-1] ^= 0x80
		out = append(out, k[:len(k)-1], x)
	}
	for _, p := range []int{len(k) - 1, len(k) / 2, 9, 12} {
		if p >= 0 && p < len(k) {
			x := bytes.Clone(k)
			x[p] ^= 1
			out = append(out, x)
		}
	}
	return out
}

func randomBounds(r *rand.Rand, sorted []string, i int) *Bounds {
	pick := func() []byte {
		k := []byte(sorted[r.IntN(len(sorted))])
		switch r.IntN(5) {
		case 0:
			return append(k, 0) // just above a key
		case 1:
			return k[:r.IntN(len(k)+1)] // a prefix of a key
		case 2:
			if len(k) > 0 { // leaves the keys mid-path, often inside a compressed path
				k[r.IntN(len(k))] += byte(1 - 2*r.IntN(2))
			}
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
	m.Range(b, func(k []byte, s View[uint64]) bool {
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
	m.Range(b, func([]byte, View[uint64]) bool { n++; return n < 3 })
	if n != min(3, len(want)) {
		t.Fatalf("Range did not stop when asked: %d calls", n)
	}
	// RangeValues yields the values of exactly these keys (Range above
	// checked their order and counts)
	var wantN, wantSum, gotN, gotSum uint64
	for _, k := range want {
		for v := range ref[k] {
			wantN, wantSum = wantN+1, wantSum+v
		}
	}
	m.RangeValues(b, func(v uint64) bool { gotN, gotSum = gotN+1, gotSum+v; return true })
	if gotN != wantN || gotSum != wantSum {
		t.Fatalf("RangeValues(%+v) yielded %d values summing to %d, want %d summing to %d", *b, gotN, gotSum, wantN, wantSum)
	}
	n = 0
	m.RangeValues(b, func(uint64) bool { n++; return n < 3 })
	if n != min(3, int(wantN)) {
		t.Fatalf("RangeValues did not stop when asked: %d calls", n)
	}
}

// TestBoundsContains makes sure that a range query on multimap.Hashed, which
// filters every key, selects exactly the keys an ordered range query on
// multimap.Ordered visits. It covers Bounds.Contains in the ART package,
// which both multimaps share, and checks each combination of open, inclusive
// and exclusive bounds at, below, above and between the bounds.
func TestBoundsContains(t *testing.T) {
	b := func(from, to string, hasFrom, hasTo, fromIncl, toIncl bool) *Bounds {
		return &Bounds{From: []byte(from), To: []byte(to), HasFrom: hasFrom, HasTo: hasTo, FromIncl: fromIncl, ToIncl: toIncl}
	}
	for _, tc := range []struct {
		name string
		b    *Bounds
		key  string
		want bool
	}{
		{"open bounds contain every key", b("", "", false, false, false, false), "anything", true},
		{"open bounds contain the empty key", b("", "", false, false, false, false), "", true},
		{"inclusive From contains From", b("b", "", true, false, true, false), "b", true},
		{"exclusive From excludes From", b("b", "", true, false, false, false), "b", false},
		{"From excludes a key below", b("b", "", true, false, true, false), "a", false},
		{"From excludes a prefix of From", b("bb", "", true, false, true, false), "b", false},
		{"From contains an extension of From", b("b", "", true, false, false, false), "b\x00", true},
		{"inclusive To contains To", b("", "d", false, true, false, true), "d", true},
		{"exclusive To excludes To", b("", "d", false, true, false, false), "d", false},
		{"To excludes a key above", b("", "d", false, true, false, true), "e", false},
		{"To excludes an extension of To", b("", "d", false, true, false, true), "da", false},
		{"To contains a prefix of To", b("", "dd", false, true, false, false), "d", true},
		{"both bounds contain a key between", b("b", "d", true, true, false, false), "c", true},
		{"both bounds exclude a key above", b("b", "d", true, true, true, true), "z", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.b.Contains([]byte(tc.key)); got != tc.want {
				t.Fatalf("Contains(%q) = %v, want %v", tc.key, got, tc.want)
			}
		})
	}
}

// TestEmptyMap makes sure that a new, empty multimap.Ordered answers every
// query with nothing. It covers the ART package's Map with no keys, where the
// tree has no root: lookups find nothing, removals are ignored and range
// queries of any kind call back not once.
func TestEmptyMap(t *testing.T) {
	var m Map[uint64]
	m.Remove([]byte("k"), 1)
	m.RemoveKey([]byte("k"))
	if m.Len() != 0 || m.Values([]byte("k")).Found() || m.Values(nil).Found() {
		t.Fatalf("empty map: Len %d, or a key was found", m.Len())
	}
	if v := m.Values([]byte("k")); v.Len() != 0 || !v.Each(func(uint64) bool { t.Fatalf("an absent key yielded a value"); return false }) {
		t.Fatalf("the values of an absent key: Len %d", v.Len())
	}
	for _, b := range []*Bounds{{}, {From: []byte("a"), To: []byte("z"), HasFrom: true, HasTo: true}} {
		m.Range(b, func([]byte, View[uint64]) bool { t.Fatalf("Range called back on an empty map"); return false })
		m.RangeValues(b, func(uint64) bool { t.Fatalf("RangeValues called back on an empty map"); return false })
	}
}

// leafLayout returns the size of a leaf[T, K] and the offsets of its key and
// its values, as the compiler lays them out.
func leafLayout[T comparable, K keyArea]() (size, kOff, vOff uintptr) {
	var l leaf[T, K]
	return unsafe.Sizeof(l), unsafe.Offsetof(l.k), unsafe.Offsetof(l.vals)
}

// TestLeafLayout makes sure that every key of multimap.Ordered that ends in a
// generic leaf keeps its bytes and its values, whatever its length and
// whatever the value type. It covers the generic leaves of the ART behind
// Ordered, which hold a key inline up to 16 bytes or as a string beyond, and
// which the untyped tree code reads through fixed offsets. It checks that a
// leaf returns an independent copy of its key, that its values lie where the
// compiler put them, that the byte the range scan touches ahead lies inside
// the leaf, and that leaves with uint64 values fill one 64-byte cache line.
func TestLeafLayout(t *testing.T) {
	type class struct{ size, kOff, vOff uintptr }
	u64 := [2]class{}
	u64[0].size, u64[0].kOff, u64[0].vOff = leafLayout[uint64, [16]byte]()
	u64[1].size, u64[1].kOff, u64[1].vOff = leafLayout[uint64, string]()
	str := [2]class{}
	str[0].size, str[0].kOff, str[0].vOff = leafLayout[string, [16]byte]()
	str[1].size, str[1].kOff, str[1].vOff = leafLayout[string, string]()

	if u64[0].size != 64 || u64[1].size != 64 {
		t.Errorf("leaf sizes for uint64 values = %d and %d, want 64 and 64", u64[0].size, u64[1].size)
	}
	for _, cs := range [][2]class{u64, str} {
		for i, c := range cs {
			if c.kOff != keyOff {
				t.Errorf("class %d: key at offset %d, want keyOff = %d", i, c.kOff, keyOff)
			}
		}
	}
	if tail := leafTail[uint64](); tail >= u64[0].size {
		t.Errorf("leafTail[uint64] = %d lies outside the leaf (%d B)", tail, u64[0].size)
	}
	if tail := leafTail[string](); tail >= str[0].size {
		t.Errorf("leafTail[string] = %d lies outside the leaf (%d B)", tail, str[0].size)
	}

	for _, tc := range []struct {
		name  string
		n     int
		class int
	}{
		{"empty key is inline", 0, 0},
		{"1 byte is inline", 1, 0},
		{"16 bytes are inline", 16, 0},
		{"17 bytes are a string", 17, 1},
		{"300 bytes are a string", 300, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			key := make([]byte, tc.n)
			for i := range key {
				key[i] = byte(i*7 + 1)
			}
			want := bytes.Clone(key)

			l := newLeaf[uint64](key)
			ls := newLeaf[string](key)
			clear(key) // the leaves must hold copies
			for _, x := range []*leafHead{l, ls} {
				if x.kind != kLeaf || int(x.klen) != tc.n || !bytes.Equal(x.key(), want) {
					t.Fatalf("leaf holds kind %d, klen %d, key %v; want a leaf of %d bytes %v", x.kind, x.klen, x.key(), tc.n, want)
				}
			}
			if uintptr(l.valsOff) != u64[tc.class].vOff || uintptr(ls.valsOff) != str[tc.class].vOff {
				t.Fatalf("values at offsets %d and %d, want %d and %d (class %d)",
					l.valsOff, ls.valsOff, u64[tc.class].vOff, str[tc.class].vOff, tc.class)
			}
			vals[uint64](l).Add(42)
			vals[string](ls).Add("v")
			if !vals[uint64](l).Contains(42) || !vals[string](ls).Contains("v") || !bytes.Equal(l.key(), want) {
				t.Fatalf("adding values changed the key or lost the values")
			}
		})
	}
}

// TestPageLayout makes sure that pages, which hold the short keys of
// multimap.Ordered together with their single values, fill whole cache-line
// sized Go size classes and keep keys and values where the untyped code
// expects them. It covers the four page classes of the ART behind Ordered
// and checks each class's size, capacity and the offsets of its arrays.
func TestPageLayout(t *testing.T) {
	for _, tc := range []struct {
		name              string
		size, want        uintptr
		heads, vals, capa uintptr
	}{
		{"3 keys fit 64 bytes", unsafe.Sizeof(page3{}), 56, unsafe.Offsetof(page3{}.heads), unsafe.Offsetof(page3{}.vals), 3},
		{"7 keys fit 128 bytes", unsafe.Sizeof(page7{}), 120, unsafe.Offsetof(page7{}.heads), unsafe.Offsetof(page7{}.vals), 7},
		{"15 keys fit 256 bytes", unsafe.Sizeof(page15{}), 248, unsafe.Offsetof(page15{}.heads), unsafe.Offsetof(page15{}.vals), 15},
		{"31 keys fit 512 bytes", unsafe.Sizeof(page31{}), 504, unsafe.Offsetof(page31{}.heads), unsafe.Offsetof(page31{}.vals), 31},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.size != tc.want || tc.heads != headsOff || tc.vals != headsOff+8*tc.capa {
				t.Fatalf("size %d, heads at %d, vals at %d; want %d, %d and %d",
					tc.size, tc.heads, tc.vals, tc.want, headsOff, headsOff+8*tc.capa)
			}
		})
	}
	for c := range pageCaps {
		p := newPage(c)
		if len(p.heads()) != pageCaps[c] || len(p.vals()) != pageCaps[c] ||
			uintptr(unsafe.Pointer(&p.vals()[0]))-uintptr(unsafe.Pointer(p)) != headsOff+8*uintptr(pageCaps[c]) {
			t.Fatalf("class %d: arrays do not match capacity %d", c, pageCaps[c])
		}
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
	if n.kind == kPage {
		return checkPage(t, asPage(n), depth)
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
		walkKeys(c, func(k []byte) {
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

// walkKeys calls fn with every key below n, in order.
func walkKeys(n *header, fn func([]byte)) {
	switch n.kind {
	case kLeaf:
		fn(asLeaf(n).key())
		return
	case kPage:
		p := asPage(n)
		var buf [8]byte
		for _, w := range p.heads()[:p.count] {
			fn(wordKey(w, int(p.klen), &buf))
		}
		return
	}
	if n.term != nil {
		fn(n.term.key())
	}
	eachChild(n, func(_ byte, c *header) { walkKeys(c, fn) })
}

// checkPage checks the invariants of a page whose path starts at depth and
// returns its number of keys: keys of one length of at most 8 bytes, strictly
// ascending, no longer than its depth, and a class that fits the count
// without being one a removal should have shrunk.
func checkPage(t *testing.T, p *pageHead, depth int) int {
	t.Helper()
	n, c := int(p.count), int(p.class)
	if n < 1 || n > pageCaps[c] || (c > 0 && n <= pageShrink[c]) {
		t.Fatalf("page of class %d (capacity %d) holds %d keys", c, pageCaps[c], n)
	}
	if p.klen > 8 || int(p.klen) < depth {
		t.Fatalf("page keys of length %d at depth %d", p.klen, depth)
	}
	h := p.heads()[:n]
	for i := 1; i < n; i++ {
		if h[i-1] >= h[i] {
			t.Fatalf("page keys not strictly ascending: %x", h)
		}
	}
	return n
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
						if !m.Values(other).Found() {
							t.Fatalf("deleting %q lost %q", k, other)
						}
					}
				}
				if m.t.root != nil {
					t.Fatalf("tree not empty after deleting every key")
				}
				m.Add(keys[0], 1)
				m.Clear()
				if m.Len() != 0 || m.Values(keys[0]).Found() {
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
