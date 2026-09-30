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
// are prefixes of other keys, the empty key, paths longer than the 12 inline
// bytes and longer than the 64K a header can count, keys of every length from
// 0 to 99 (every size class of the leaves and beyond), zero bytes, dense and
// sparse integers (which drive nodes through every kind), nodes whose slots a
// term fills up, and shared string prefixes.
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
	for _, fan := range []int{3, 4, 5, 10, 11, 12, 24, 25, 26, 56, 57, 58, 256} {
		p := append(append([]byte(nil), long...), byte(fan))
		wide = append(wide, p) // a term next to fan children: fills the node's last slot
		for b := range fan {
			wide = append(wide, append(append(p[:len(p):len(p)], byte(b)), "tail"...))
		}
	}
	sets["long-prefix-wide"] = append(wide, long[:20], append(long[:30:30], 'Z'))

	// Paths of 64K bytes and more, which the header marks as longPath: a
	// long path with a term, a split inside it, a 256-way node below it, and
	// two paths of 40,000 bytes that merge into one of 80,001 once the key
	// between them goes.
	x, y := bytes.Repeat([]byte("x"), 70000), bytes.Repeat([]byte("y"), 40000)
	long64k := [][]byte{
		x, append(x[:len(x):len(x)], 'a'), append(x[:len(x):len(x)], 'b'), append(x[:len(x):len(x)], "bc"...),
		append(x[:69000:69000], 'z'),
		slices.Concat(y, []byte("a"), y, []byte("1")), slices.Concat(y, []byte("a"), y, []byte("2")),
		slices.Concat(y, []byte("b")),
	}
	for b := range 64 {
		long64k = append(long64k, slices.Concat(x, []byte{'c', byte(b)}))
	}
	sets["paths-over-64k"] = long64k

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
// invariants demand (see checkInvariants). It runs every corpus with flat
// leaves, which small pointer-free values get, and with set leaves, which
// all other values get.
func TestAgainstReference(t *testing.T) {
	for _, mode := range []int8{1, -1} {
		for name, keys := range keySets() {
			t.Run(fmt.Sprintf("%s/flat=%d", name, mode), func(t *testing.T) {
				againstReference(t, keys, mode)
			})
		}
	}
}

func againstReference(t *testing.T, keys [][]byte, mode int8) {
	r := rand.New(rand.NewPCG(7, 8))
	m := Map[uint64]{flat: mode}
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
						v = uint64(r.IntN(200)) // grow flat leaves, spill sets into array and hash
					}
					m.Add(k, v)
					ref.add(k, v)
				}
			}
		}
		compare(t, &m, ref, r)
		checkInvariants(t, &m.t)
	}
}

func compare(t *testing.T, m *Map[uint64], ref reference, r *rand.Rand) {
	t.Helper()
	if m.Len() != len(ref) {
		t.Fatalf("Len = %d, want %d", m.Len(), len(ref))
	}
	for k, want := range ref {
		got := valuesOf(m, []byte(k))
		if len(got) != len(want) {
			t.Fatalf("key %q holds %d values, want %d", k, len(got), len(want))
		}
		for _, v := range got {
			if !want[v] {
				t.Fatalf("key %q holds %d unexpectedly", k, v)
			}
		}
	}
	for k := range ref {
		for _, probe := range nearMisses([]byte(k)) {
			_, want := ref[string(probe)]
			if got := m.Has(probe); got != want {
				t.Fatalf("Has(%q) = %v, want %v", probe, got, want)
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

// valuesOf returns the values of key in m, in the map's order.
func valuesOf(m *Map[uint64], key []byte) []uint64 {
	var out []uint64
	m.Each(key, func(v uint64) bool { out = append(out, v); return true })
	return out
}

// nearMisses returns keys that differ from k only slightly: shorter, longer,
// or with one byte changed at the end, in the middle, or at positions 9, 12
// and 13, which lie around the 12 path bytes a node checks itself. The last byte
// is also changed in its top bit, which a dense node rarely holds.
func nearMisses(k []byte) [][]byte {
	out := [][]byte{append(bytes.Clone(k), 0)}
	if len(k) > 0 {
		x := bytes.Clone(k)
		x[len(x)-1] ^= 0x80
		out = append(out, k[:len(k)-1], x)
	}
	for _, p := range []int{len(k) - 1, len(k) / 2, 9, 12, 13} {
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
	m.Range(b, func(k []byte) bool {
		got = append(got, string(k))
		return true
	})
	if !slices.Equal(got, want) {
		t.Fatalf("Range(%+v) yielded %d keys, want %d\n got=%q\nwant=%q", *b, len(got), len(want), got, want)
	}
	// stopping early must stop
	n := 0
	m.Range(b, func([]byte) bool { n++; return n < 3 })
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
	if m.Len() != 0 || m.Has([]byte("k")) || m.Has(nil) || valuesOf(&m, []byte("k")) != nil {
		t.Fatalf("empty map: Len %d, or a key was found", m.Len())
	}
	for _, b := range []*Bounds{{}, {From: []byte("a"), To: []byte("z"), HasFrom: true, HasTo: true}} {
		m.Range(b, func([]byte) bool { t.Fatalf("Range called back on an empty map"); return false })
		m.RangeValues(b, func(uint64) bool { t.Fatalf("RangeValues called back on an empty map"); return false })
	}
}

// leafLayout returns the size of a leaf[T, K] and the offsets of its key and
// its values, as the compiler lays them out.
func leafLayout[T comparable, K keyArea]() (size, kOff, vOff uintptr) {
	var l leaf[T, K]
	return unsafe.Sizeof(l), unsafe.Offsetof(l.k), unsafe.Offsetof(l.vals)
}

// TestLeafLayout makes sure that every key of multimap.Ordered keeps its bytes
// and its values, whatever its length and whatever the value type. It covers
// the set leaves of the ART behind Ordered, which hold a key's remainder
// inline in the smallest of eight size classes, or the whole key as a string
// beyond 254 bytes, and which the untyped tree code reads through fixed
// offsets. For each class boundary it checks that the leaf holds an
// independent copy of its key and the key's length, that its values lie
// where the compiler put them, and that the byte the range scan touches
// ahead lies inside even the smallest leaf.
func TestLeafLayout(t *testing.T) {
	type class struct{ size, kOff, vOff uintptr }
	layouts := func(get ...func() (uintptr, uintptr, uintptr)) (out []class) {
		for _, f := range get {
			var c class
			c.size, c.kOff, c.vOff = f()
			out = append(out, c)
		}
		return out
	}
	u64 := layouts(leafLayout[uint64, [16]byte], leafLayout[uint64, [32]byte], leafLayout[uint64, [48]byte],
		leafLayout[uint64, [64]byte], leafLayout[uint64, [96]byte], leafLayout[uint64, [128]byte],
		leafLayout[uint64, [192]byte], leafLayout[uint64, [256]byte], leafLayout[uint64, string])
	str := layouts(leafLayout[string, [16]byte], leafLayout[string, [32]byte], leafLayout[string, [48]byte],
		leafLayout[string, [64]byte], leafLayout[string, [96]byte], leafLayout[string, [128]byte],
		leafLayout[string, [192]byte], leafLayout[string, [256]byte], leafLayout[string, string])

	if u64[0].size != 64 || u64[len(u64)-1].size != 64 {
		t.Errorf("smallest leaves for uint64 values are %d and %d bytes, want one cache line", u64[0].size, u64[len(u64)-1].size)
	}
	for _, cs := range [][]class{u64, str} {
		for i, c := range cs {
			want := keyOff
			if i == len(cs)-1 {
				want = strOff
			}
			if c.kOff != want {
				t.Errorf("class %d: key at offset %d, want %d", i, c.kOff, want)
			}
		}
	}
	if tail := (&Map[uint64]{flat: -1}).leafTail(); tail >= u64[0].size {
		t.Errorf("leafTail = %d lies outside the smallest set leaf for uint64 (%d B)", tail, u64[0].size)
	}
	if tail := (&Map[string]{flat: -1}).leafTail(); tail >= str[0].size {
		t.Errorf("leafTail = %d lies outside the smallest set leaf for string (%d B)", tail, str[0].size)
	}
	if tail := (&Map[uint64]{flat: 1}).leafTail(); tail >= flatSizes[1] {
		t.Errorf("leafTail = %d lies outside the smallest flat leaf (%d B)", tail, flatSizes[1])
	}

	for _, tc := range []struct {
		name  string
		n     int
		class int
	}{
		{"empty key is inline", 0, 0},
		{"16 bytes are inline in 16", 16, 0},
		{"17 bytes are inline in 32", 17, 1},
		{"33 bytes are inline in 48", 33, 2},
		{"49 bytes are inline in 64", 49, 3},
		{"65 bytes are inline in 96", 65, 4},
		{"97 bytes are inline in 128", 97, 5},
		{"129 bytes are inline in 192", 129, 6},
		{"193 bytes are inline in 256", 193, 7},
		{"254 bytes are inline in 256", maxInline, 7},
		{"255 bytes are a string", maxInline + 1, 8},
		{"300 bytes are a string", 300, 8},
	} {
		t.Run(tc.name, func(t *testing.T) {
			key := make([]byte, tc.n)
			for i := range key {
				key[i] = byte(i*7 + 1)
			}
			want := bytes.Clone(key)

			l := newSetLeaf[uint64](key, 0)
			ls := newSetLeaf[string](key, 0)
			clear(key) // the leaves must hold copies
			wantLen := uint8(tc.n)
			if tc.n > maxInline {
				wantLen = longKey
			}
			for _, x := range []*leafHead{l, ls} {
				if x.kind != kSet || x.klen != wantLen || x.keyLen() != tc.n || x.base() != 0 || !bytes.Equal(x.stored(), want) {
					t.Fatalf("leaf holds kind %d, klen %d, key %v; want a set leaf of %d bytes %v", x.kind, x.klen, x.stored(), tc.n, want)
				}
			}
			gotU := uintptr(unsafe.Pointer(vals[uint64](l))) - uintptr(unsafe.Pointer(l))
			gotS := uintptr(unsafe.Pointer(vals[string](ls))) - uintptr(unsafe.Pointer(ls))
			if gotU != u64[tc.class].vOff || gotS != str[tc.class].vOff {
				t.Fatalf("values at offsets %d and %d, want %d and %d (size class %d)",
					gotU, gotS, u64[tc.class].vOff, str[tc.class].vOff, tc.class)
			}
			vals[uint64](l).Add(42)
			vals[string](ls).Add("v")
			if !vals[uint64](l).Contains(42) || !vals[string](ls).Contains("v") || !bytes.Equal(l.stored(), want) {
				t.Fatalf("adding values changed the key or lost the values")
			}
		})
	}
}

// TestNodeLayout makes sure that the inner nodes of multimap.Ordered stay the
// cache-line sized objects the tree is designed around: every node kind of
// the adaptive radix tree fills a Go size class of 64, 128, 256 or 512 bytes
// (the 256-way node excepted), behind a 16-byte header shared by all kinds,
// the kind byte sits where leaves keep theirs, and the tail of a long path
// starts right after the fixed part of every kind, where the untyped tree
// code reads it.
func TestNodeLayout(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  uintptr
		want uintptr
	}{
		{"header", unsafe.Sizeof(header{}), 16},
		{"5-way node", unsafe.Sizeof(node5{}), 64},
		{"12-way node", unsafe.Sizeof(node12{}), 128},
		{"26-way node", unsafe.Sizeof(node26{}), 256},
		{"58-way node", unsafe.Sizeof(node58{}), 512},
		{"256-way node", unsafe.Sizeof(node256{}), 2080},
		{"kind at the start of a node", unsafe.Offsetof(header{}.kind), 0},
		{"kind at the start of a leaf", unsafe.Offsetof(leafHead{}.kind), 0},
		{"leaf head", unsafe.Sizeof(leafHead{}), 6},
		{"tail of a 5-way node", unsafe.Offsetof(tailed[node5, [16]byte]{}.t), fixedSize[kN5]},
		{"string tail of a 12-way node", unsafe.Offsetof(tailed[node12, string]{}.t), fixedSize[kN12]},
		{"tail of a 26-way node", unsafe.Offsetof(tailed[node26, [48]byte]{}.t), fixedSize[kN26]},
		{"tail of a 58-way node", unsafe.Offsetof(tailed[node58, [112]byte]{}.t), fixedSize[kN58]},
		{"string tail of a 256-way node", unsafe.Offsetof(tailed[node256, string]{}.t), fixedSize[kN256]},
	} {
		if tc.got != tc.want {
			t.Errorf("%s: %d bytes, want %d", tc.name, tc.got, tc.want)
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
	if n := checkNode(t, tr.root, nil); n != tr.size {
		t.Fatalf("tree holds %d leaves, size says %d", n, tr.size)
	}
}

// checkLeaf fails unless leaf l, reached through path, holds its key from a
// base within path on, and the bytes it holds of path agree with it.
func checkLeaf(t *testing.T, l *leafHead, path []byte) {
	t.Helper()
	b := l.base()
	if b < 0 || b > len(path) || l.keyLen() < len(path) || !bytes.Equal(l.stored()[:len(path)-b], path[b:]) {
		t.Fatalf("leaf holding %q from %d does not continue its path %q", l.stored(), b, path)
	}
}

// checkNode checks the subtree n reached through path, the key bytes above
// it, and returns its number of leaves.
func checkNode(t *testing.T, n *header, path []byte) int {
	t.Helper()
	if isLeaf(n.kind) {
		checkLeaf(t, asLeaf(n), path)
		return 1
	}
	limits := map[kind][2]int{kN5: {1, 5}, kN12: {shrink12 + 1, 12}, kN26: {shrink26 + 1, 26},
		kN58: {shrink58 + 1, 58}, kN256: {shrink256 + 1, 256}}[n.kind]
	count, term := int(n.count), termOf(n)
	if n.kind == kN256 {
		if n.count != 255 {
			t.Fatalf("256-way node with header count %d, want 255", n.count)
		}
		count = int(asN256(n).total)
	}
	if lo, hi := limits[0], limits[1]; count < lo || count > hi {
		t.Fatalf("kind %d holds %d children, allowed %d..%d", n.kind, count, lo, hi)
	}
	if count+b2i(term != nil) < 2 {
		t.Fatalf("node does not branch (count %d, term %v): it should have collapsed", count, term != nil)
	}
	if s := slots(n); n.kind != kN256 {
		if count+b2i(term != nil) > len(s) {
			t.Fatalf("kind %d holds %d children and a term in %d slots", n.kind, count, len(s))
		}
		for i := count; i < len(s)-1; i++ {
			if s[i] != nil {
				t.Fatalf("kind %d: unused slot %d is not empty", n.kind, i)
			}
		}
	}
	pl := n.pathLen()
	if want := uint16(min(pl, longPath)); n.plen != want {
		t.Fatalf("plen %d for a path of %d bytes, want %d", n.plen, pl, want)
	}
	for i := min(pl, len(n.prefix)); i < len(n.prefix); i++ {
		if n.prefix[i] != 0 {
			t.Fatalf("path byte %d beyond a path of %d bytes is not zero", i, pl)
		}
	}
	end := appendPath(slices.Clip(path), n)
	leaves := 0
	if term != nil {
		checkLeaf(t, term, end)
		if term.keyLen() != len(end) {
			t.Fatalf("term key of %d bytes does not end at depth %d", term.keyLen(), len(end))
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
		leaves += checkNode(t, c, append(slices.Clip(end), b))
	})
	if children != count {
		t.Fatalf("count %d but %d children", count, children)
	}
	return leaves
}

func eachChild(n *header, fn func(byte, *header)) {
	switch n.kind {
	case kN26, kN58:
		bm, child := bitmapOf(n)
		i := 0
		for k := range 256 {
			if (bm[k>>6]>>(k&63))&1 == 1 {
				fn(byte(k), child[i])
				i++
			}
		}
	case kN256:
		for k, c := range asN256(n).child[:256] {
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

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// TestShrinkAndCollapse checks that deleting keys one by one takes every node
// kind back down through each smaller kind to nothing, collapsing and
// re-merging compressed paths (short and longer than 12 bytes) on the way,
// moving the term of the widest node along, and that the tree satisfies its
// invariants after every single delete.
func TestShrinkAndCollapse(t *testing.T) {
	for _, prefix := range []string{"", "p", "a-path-longer-than-twelve-bytes"} {
		for _, fan := range []int{2, 4, 5, 11, 12, 25, 26, 57, 58, 256} {
			t.Run(fmt.Sprintf("%q/%d", prefix, fan), func(t *testing.T) {
				r := rand.New(rand.NewPCG(uint64(fan), 9))
				var m Map[uint64]
				keys := [][]byte{[]byte(prefix)} // the term of the widest node
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
						if !m.Has(other) {
							t.Fatalf("deleting %q lost %q", k, other)
						}
					}
				}
				if m.t.root != nil {
					t.Fatalf("tree not empty after deleting every key")
				}
				m.Add(keys[0], 1)
				m.Clear()
				if m.Len() != 0 || m.Has(keys[0]) {
					t.Fatalf("Clear left keys behind")
				}
			})
		}
	}
}

// FuzzOperations drives the tree with arbitrary operation sequences over
// short keys from a tiny alphabet (which maximises shared paths, splits and
// merges), checking every result against the reference and the structural
// invariants at the end. The first byte chooses flat or set leaves.
func FuzzOperations(f *testing.F) {
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	f.Add(bytes.Repeat([]byte{3, 0, 1, 2, 7, 1, 0, 0, 5}, 30))
	f.Fuzz(func(t *testing.T, ops []byte) {
		m := Map[uint64]{flat: 1}
		if len(ops) > 0 && ops[0]&1 == 1 {
			m.flat = -1
		}
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

// TestLongPaths makes sure that multimap.Ordered keeps keys that share long
// common parts, such as URLs of one site or files of one directory, while
// other keys split those parts and merge them again and the number of keys
// below them grows and shrinks. It covers the path tails of the ART nodes
// (path.go): for paths at every tail class boundary and every node kind, a
// key that leaves the path in its middle moves the rest of the path into a
// node of another tail class, removing that key merges the path back, and
// growing and shrinking the node carries its tail through every kind. After
// every step the map must agree with a reference and satisfy its invariants.
func TestLongPaths(t *testing.T) {
	for _, pl := range []int{12, 13, 28, 29, 60, 61, 124, 125, 300} {
		for _, fan := range []int{2, 6, 13, 27, 59} {
			t.Run(fmt.Sprintf("path of %d bytes, %d children", pl, fan), func(t *testing.T) {
				common := make([]byte, pl)
				for i := range common {
					common[i] = byte(i%251 + 1)
				}
				key := func(b int) []byte { return append(append(slices.Clip(common), byte(b)), "tail"...) }
				var m Map[uint64]
				ref := reference{}
				step := func(add bool, k []byte) {
					if add {
						m.Add(k, 1)
						ref.add(k, 1)
					} else {
						m.RemoveKey(k)
						delete(ref, string(k))
					}
					checkInvariants(t, &m.t)
					for k := range ref {
						if !m.Has([]byte(k)) {
							t.Fatalf("lost key %q", k)
						}
					}
				}
				phase := func() { compare(t, &m, ref, rand.New(rand.NewPCG(uint64(pl), uint64(fan)))) }
				for b := range fan {
					step(true, key(b))
				}
				phase()
				split := append(slices.Clip(common[:pl/2]), 0xff)
				step(true, split)
				phase()
				step(false, split)
				phase()
				for b := fan; b < 60; b++ {
					step(true, key(b))
				}
				phase()
				for b := 59; b >= 0; b-- {
					step(false, key(b))
				}
				phase()
			})
		}
	}
}
