package art

import (
	"bytes"
	"encoding/binary"
	"fmt"
	set3 "github.com/TomTonic/Set3"
	"maps"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
	"testing"
	"unsafe"
)

// keySets returns key corpora that exercise every structural case: keys that
// are prefixes of other keys, the empty key, paths longer than the 12 inline
// bytes and longer than the 64K a header can count, keys of every length from
// 0 to 99 (every size class of the leaves and beyond), zero bytes, dense and
// sparse integers (which drive nodes through every type), nodes whose slots a
// end page fills up, and shared string prefixes.
func keySets() map[string][][]byte {
	// size is the number of keys of a big set: full, or a tenth of it under the
	// race detector, which would otherwise take the suite past go test's timeout.
	size := func(n int) int {
		if underRace {
			return n / 10
		}
		return n
	}
	r := rand.New(rand.NewPCG(1, 2))
	sets := map[string][][]byte{}
	var u64, dense [][]byte
	for range size(5000) {
		u64 = append(u64, binary.BigEndian.AppendUint64(nil, r.Uint64()))
	}
	for i := range uint64(size(20000)) {
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
	for range size(8000) {
		strs = append(strs, fmt.Appendf(nil, "%s/%s/%s/%d", words[r.IntN(len(words))],
			words[r.IntN(len(words))], words[r.IntN(len(words))], r.IntN(300)))
	}
	sets["strings"] = strs

	var wide [][]byte
	long := []byte("a-compressed-path-longer-than-sixteen-bytes/")
	for _, fan := range []int{3, 4, 5, 10, 11, 12, 24, 25, 26, 56, 57, 58, 256} {
		p := append(append([]byte(nil), long...), byte(fan))
		wide = append(wide, p) // an end page next to fan children: fills the node's last slot
		for b := range fan {
			wide = append(wide, append(append(p[:len(p):len(p)], byte(b)), "tail"...))
		}
	}
	// ... and an end page that arrives after the fan children filled the node up to
	// its capacity, which makes it grow.
	late := []byte("a-prefix-of-keys-that-end-late/")
	for _, fan := range []int{5, 12, 26, 58} {
		p := append(append([]byte(nil), late...), byte(fan))
		for b := range fan {
			wide = append(wide, append(append(p[:len(p):len(p)], byte(b)), "tail"...))
		}
		wide = append(wide, p)
	}
	sets["long-prefix-wide"] = append(wide, long[:20], append(long[:30:30], 'Z'))

	// Paths of 64K bytes and more, which the header marks as longPrefix: a
	// long path with an end page, a split inside it, a 256-way node below it, and
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
	for range size(6000) {
		k := make([]byte, r.IntN(12))
		for i := range k {
			k[i] = byte(r.IntN(3))
		}
		small = append(small, k)
	}
	sets["random-small-alphabet"] = small

	var lengths [][]byte
	for range size(4000) {
		k := make([]byte, r.IntN(100))
		for i := range k {
			k[i] = 'a' + byte(r.IntN(4))
		}
		lengths = append(lengths, k)
	}
	sets["every-length"] = lengths

	// Keys whose pages lie deep, or that no page can hold: 200 shared bytes, a
	// byte that splits the keys in two families, 69 more bytes that each family
	// shares, and two bytes of 8 symbols. A page that forks the two families
	// holds the 70 bytes below it, a page that is full rebuilds at a pathLen
	// beyond 255 (where pages end, see maxPagePathLen), and a key of 300 bytes,
	// with no part of it in a page, comes along.
	top, mid := bytes.Repeat([]byte("t"), 200), bytes.Repeat([]byte("m"), 69)
	deep := [][]byte{ // the first two keys fork at byte 200, so that pages hold the others
		slices.Concat(top, []byte{1}, mid, []byte("aa")), slices.Concat(top, []byte{2}, mid, []byte("aa")),
	}
	for range size(500) {
		fam := byte(1 + r.IntN(2))
		deep = append(deep, slices.Concat(top, []byte{fam}, mid, []byte{'a' + byte(r.IntN(8)), 'a' + byte(r.IntN(8))}))
	}
	sets["deep"] = append(deep, bytes.Repeat([]byte("l"), 300))
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
// invariants demand (see checkInvariants). It runs every corpus with single-key
// pages of fixed-size values, which small pointer-free values get, with pages of
// pointers, with pages of strings, and with value overflows, which all other values get.
func TestAgainstReference(t *testing.T) {
	str := func(v uint64) string { return fmt.Sprint("value ", v) }
	// values of 20 to 140 bytes, so that a key with a few of them outgrows a page
	longStr := func(v uint64) string { return fmt.Sprint("value ", v, strings.Repeat("x", 20+int(v%120))) }
	for _, mode := range []int8{0, 1, 2, -1, -2, 3, 4} {
		for name, keys := range keySets() {
			t.Run(fmt.Sprintf("%s/leaves=%d", name, mode), func(t *testing.T) {
				if underRace && name == "paths-over-64k" && mode != 0 && mode != -2 {
					t.Skip("the long paths are the same for every leaf type; the race detector makes them slow")
				}
				switch mode {
				case 0:
					againstReference(t, keys, &Map[uint64]{}, id) // with pages
					uniqueAgainstReference(t, keys, &Map[uint64]{}, id)
				case 1, -1:
					againstReference(t, keys, &Map[uint64]{flat: mode}, id)
				case 2:
					againstReference(t, keys, &Map[*rec]{}, ptrValue) // pages of pointers
				case 3:
					againstReference(t, keys, &Map[string]{}, str) // single-key pages
				case 4:
					againstReference(t, keys, &Map[string]{}, longStr) // single-key pages that overflow
				default:
					againstReference(t, keys, &Map[string]{flat: -1}, str)
				}
			})
		}
	}
}

func againstReference[T comparable](t *testing.T, keys [][]byte, m *Map[T], mk func(uint64) T) {
	r := rand.New(rand.NewPCG(7, 8))
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
				m.Remove(k, mk(v))
				ref.remove(k, v)
			case !removing || op < 8:
				for range 1 + r.IntN(4) {
					v := uint64(r.IntN(8))
					if r.IntN(30) == 0 {
						v = uint64(r.IntN(200)) // grow flat leaves, spill sets into array and hash
					}
					m.Add(k, mk(v))
					ref.add(k, v)
				}
			}
		}
		compare(t, m, ref, mk, r)
		checkInvariants(t, &m.t)
	}
}

// uniqueAgainstReference is againstReference for keys that hold one value, as
// they do in most maps, and a few that get a second one. It runs the pages of
// a map of T through growing, splitting, promoting a key to a leaf, shrinking
// and merging, and compares the map with the reference after each phase.
func uniqueAgainstReference[T comparable](t *testing.T, keys [][]byte, m *Map[T], mk func(uint64) T) {
	r := rand.New(rand.NewPCG(9, 10))
	ref := reference{}
	value := func(k []byte) uint64 { return uint64(len(k)*7+int(slices.Max(append([]byte{0}, k...)))) % 8 }
	for phase := range 6 {
		removing := phase%2 == 1
		for _, k := range keys {
			switch op := r.IntN(100); {
			case removing && op < 50:
				m.RemoveKey(k)
				delete(ref, string(k))
			case removing && op < 90:
				m.Remove(k, mk(value(k)))
				ref.remove(k, value(k))
			case !removing || op < 95:
				m.Add(k, mk(value(k)))
				ref.add(k, value(k))
				if r.IntN(40) == 0 {
					w := 100 + uint64(r.IntN(3)) // a second value gives the key a leaf
					m.Add(k, mk(w))
					ref.add(k, w)
				}
			}
		}
		compare(t, m, ref, mk, r)
		checkInvariants(t, &m.t)
	}
}

func compare[T comparable](t *testing.T, m *Map[T], ref reference, mk func(uint64) T, r *rand.Rand) {
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
			if !hasValue(want, v, mk) {
				t.Fatalf("key %q holds %v unexpectedly", k, v)
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
		checkRange(t, m, ref, mk, sorted, b)
	}
	checkRange(t, m, ref, mk, sorted, &Bounds{}) // everything
}

// id maps a reference value to itself, for maps of uint64.
func id(v uint64) uint64 { return v }

// hasValue reports whether the reference set want holds a value that mk maps to v.
func hasValue[T comparable](want map[uint64]bool, v T, mk func(uint64) T) bool {
	for w := range want {
		if mk(w) == v {
			return true
		}
	}
	return false
}

// valuesOf returns the values of key in m, in the map's order.
func valuesOf[T comparable](m *Map[T], key []byte) []T {
	var out []T
	m.Each(key, func(v T) bool { out = append(out, v); return true })
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
			if len(k) > 0 { // leaves the keys mid-path, often inside a common prefix
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

func checkRange[T comparable](t *testing.T, m *Map[T], ref reference, mk func(uint64) T, sorted []string, b *Bounds) {
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
	wantN, gotN := map[T]int{}, map[T]int{}
	for _, k := range want {
		for v := range ref[k] {
			wantN[mk(v)]++
		}
	}
	m.RangeValues(b, func(v T) bool { gotN[v]++; return true })
	if !maps.Equal(gotN, wantN) {
		t.Fatalf("RangeValues(%+v) yielded %v, want %v", *b, gotN, wantN)
	}
	n = 0
	m.RangeValues(b, func(T) bool { n++; return n < 3 })
	total := 0
	for _, c := range wantN {
		total += c
	}
	if n != min(3, total) {
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
// queries of any type call back not once.
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

// TestSmallestObject makes sure that the byte the range scan touches ahead lies
// inside the smallest object of every map: a single-key page or a value overflow
// is at least 32 bytes, whatever the type of the values, and a value overflow
// holds its key remainder where the untyped tree code reads it (keyOff) or, beyond
// the longest inline remainder, the whole key as a string (strOff).
func TestSmallestObject(t *testing.T) {
	if leafTail >= 32 || leafTail >= unsafe.Sizeof(valueOverflow[[20]byte, int]{}) {
		t.Errorf("leafTail = %d lies outside the smallest object (%d B)", leafTail, unsafe.Sizeof(valueOverflow[[20]byte, int]{}))
	}
	for _, n := range []int{0, 20, 52, 116, 244, 372, 500, 501} {
		l := newValueOverflow(bytes.Repeat([]byte("k"), n), set3.Empty[uint64]())
		if n > maxInlineOverflow != (l.rem() == longKey) || len(l.stored()) != n {
			t.Errorf("key of %d bytes: remainder %d", n, l.rem())
		}
	}
}

// TestNodeLayout makes sure that the byte nodes of multimap.Ordered stay the
// cache-line sized objects the tree is designed around: every node type of
// the adaptive radix tree fills a Go size class of 64, 128, 256 or 512 bytes
// (the 256-way node excepted), behind a 16-byte header shared by all types,
// the type byte sits where leaves keep theirs, and the tail of a long path
// starts right after the fixed part of every type, where the untyped tree
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
		{"type at the start of a node", unsafe.Offsetof(header{}.objType), 0},
		{"type at the start of a leaf", unsafe.Offsetof(singleKeyHead{}.objType), 0},
		{"leaf head", unsafe.Sizeof(singleKeyHead{}), 4},
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
		t.Fatalf("tree holds %d keys, size says %d", n, tr.size)
	}
}

// checkLeaf fails unless leaf l, reached through path, is a well-formed leaf: a page has at least one value
// and holds only its own key part (the path is in the nodes above), a value overflow's key part fits its kind.
func checkLeaf(t *testing.T, l *singleKeyHead, path []byte) {
	t.Helper()
	if !l.isValueOverflow() && l.n == 0 {
		t.Fatalf("a page without a value after the path %q", path)
	}
	if k := l.rem(); k != longKey && k != len(l.stored()) {
		t.Fatalf("leaf of %d key bytes announces %d", len(l.stored()), k)
	}
}

// checkNode checks the subtree n reached through path, the key bytes above
// it, and returns its number of leaves.
func checkNode(t *testing.T, n *header, path []byte) int {
	t.Helper()
	if isSingleKey(n.objType) {
		checkLeaf(t, asSingleKey(n), path)
		return 1
	}
	if isPage(n.objType) { // a multi-key page: two keys at least, with one value or more each
		p := asMKStr(n) // the head is that of both kinds of page
		if p.Keys() < 2 || p.Len() < p.Keys() {
			t.Fatalf("multi-key page with %d keys and %d values", p.Keys(), p.Len())
		}
		return p.Keys()
	}
	limits := map[objType][2]int{kN5: {1, 5}, kN12: {shrink12 + 1, 12}, kN26: {shrink26 + 1, 26},
		kN58: {shrink58 + 1, 58}, kN256: {shrink256 + 1, 256}}[n.objType]
	count, endPage := int(n.count), endPageOf(n)
	if n.objType == kN256 {
		if n.count != 255 {
			t.Fatalf("256-way node with header count %d, want 255", n.count)
		}
		count = int(asN256(n).total)
	}
	if lo, hi := limits[0], limits[1]; count < lo || count > hi {
		t.Fatalf("type %d holds %d children, allowed %d..%d", n.objType, count, lo, hi)
	}
	if count+b2i(endPage != nil) < 2 && (count != 1 || !isMultiKey(onlyChildOf(n).objType)) { // a multi-key page that could not move up stays below its node
		t.Fatalf("node does not branch (type %d, count %d, endPage %v): it should have collapsed", n.objType, count, endPage != nil)
	}
	if n.objType != kN256 {
		s := slots(n)
		if count+b2i(endPage != nil) > len(s) {
			t.Fatalf("type %d holds %d children and a endPage in %d slots", n.objType, count, len(s))
		}
		for i := count; i < len(s)-1; i++ {
			if s[i] != nil {
				t.Fatalf("type %d: unused slot %d is not empty", n.objType, i)
			}
		}
	}
	pl := n.prefixLen()
	if want := uint16(min(pl, longPrefix)); n.plen != want {
		t.Fatalf("plen %d for a path of %d bytes, want %d", n.plen, pl, want)
	}
	for i := min(pl, len(n.prefix)); i < len(n.prefix); i++ {
		if n.prefix[i] != 0 {
			t.Fatalf("path byte %d beyond a path of %d bytes is not zero", i, pl)
		}
	}
	end := appendPrefix(slices.Clip(path), n)
	leaves := 0
	if endPage != nil {
		checkLeaf(t, endPage, end)
		if len(endPage.stored()) != 0 {
			t.Fatalf("end page holds %d key bytes behind the path of %d bytes", len(endPage.stored()), len(end))
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

// onlyChildOf returns the first child of n.
func onlyChildOf(n *header) *header {
	var c *header
	eachChild(n, func(_ byte, x *header) {
		if c == nil {
			c = x
		}
	})
	return c
}

func eachChild(n *header, fn func(byte, *header)) {
	switch n.objType {
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

// TestShrinkAndCollapse checks that deleting keys one by one takes every node
// type back down through each smaller type to nothing, collapsing and
// re-merging common prefixes (short and longer than 12 bytes) on the way,
// moving the end page of the widest node along, and that the tree satisfies its
// invariants after every single delete.
func TestShrinkAndCollapse(t *testing.T) {
	for _, prefix := range []string{"", "p", "a-path-longer-than-twelve-bytes"} {
		for _, fan := range []int{2, 4, 5, 11, 12, 25, 26, 57, 58, 256} {
			t.Run(fmt.Sprintf("%q/%d", prefix, fan), func(t *testing.T) {
				r := rand.New(rand.NewPCG(uint64(fan), 9))
				m := Map[uint64]{flat: 1}        // tails of 260 bytes: two keys do not fit one page, so the keys make nodes
				keys := [][]byte{[]byte(prefix)} // the end page of the widest node
				for b := range fan {
					for _, tail := range []string{"", "x" + strings.Repeat("a", 260), "xy-longer-tail-than-16" + strings.Repeat("b", 260)} {
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
// invariants at the end. The same operations run on a map of uint64, whose
// first byte's lowest bit chooses flat or value overflows and third bit pages, and
// on a map of strings, whose second bit chooses typed or value overflows.
func FuzzOperations(f *testing.F) {
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	f.Add(bytes.Repeat([]byte{3, 0, 1, 2, 7, 1, 0, 0, 5}, 30))
	f.Fuzz(func(t *testing.T, ops []byte) {
		m := Map[uint64]{flat: 1}
		s := Map[*rec]{}  // the same operations on pages of pointers
		var p Map[string] // and on single-key pages, with values of up to 200 bytes
		lstr := func(v uint64) string { return fmt.Sprint("value ", v, strings.Repeat("x", int(v%7)*33)) }
		if len(ops) > 0 && ops[0]&1 == 1 {
			m.flat = -1
		}
		if len(ops) > 0 && ops[0]&4 == 4 {
			m = Map[uint64]{} // with pages
		}
		if len(ops) > 0 && ops[0]&2 == 2 {
			s.flat = -1
		}
		str := ptrValue
		ref := reference{}
		for len(ops) >= 2 {
			op, n := ops[0], int(ops[1]%24)
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
				s.Add(k, str(v))
				p.Add(k, lstr(v))
				ref.add(k, v)
			case 2:
				m.Remove(k, v)
				s.Remove(k, str(v))
				p.Remove(k, lstr(v))
				ref.remove(k, v)
			default:
				m.RemoveKey(k)
				s.RemoveKey(k)
				p.RemoveKey(k)
				delete(ref, string(k))
			}
		}
		compare(t, &m, ref, id, rand.New(rand.NewPCG(1, 1)))
		checkInvariants(t, &m.t)
		compare(t, &s, ref, str, rand.New(rand.NewPCG(1, 1)))
		checkInvariants(t, &s.t)
		compare(t, &p, ref, lstr, rand.New(rand.NewPCG(1, 1)))
		checkInvariants(t, &p.t)
	})
}

// TestLongPaths makes sure that multimap.Ordered keeps keys that share long
// common parts, such as URLs of one site or files of one directory, while
// other keys split those parts and merge them again and the number of keys
// below them grows and shrinks. It covers the prefix tails of the ART nodes
// (prefix.go): for paths at every tail class boundary and every node type, a
// key that leaves the path in its middle moves the rest of the path into a
// node of another tail class, removing that key merges the path back, and
// growing and shrinking the node carries its tail through every type. After
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
				m := Map[uint64]{flat: 1} // single-key pages only, so that the keys make nodes
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
				phase := func() { compare(t, &m, ref, id, rand.New(rand.NewPCG(uint64(pl), uint64(fan)))) }
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

// b2i returns 1 for true and 0 for false.
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// ptrPool holds the records that the pointer values of the tests point to, so
// that equal numbers are equal pointers.
var ptrPool = func() []*rec {
	out := make([]*rec, 1<<12)
	for i := range out {
		out[i] = &rec{id: uint64(i), aux: 1}
	}
	return out
}()

// ptrValue returns the pointer value for number v.
func ptrValue(v uint64) *rec { return ptrPool[v%uint64(len(ptrPool))] }
