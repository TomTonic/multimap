package mkpage

import (
	"bytes"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
	"unsafe"
)

// fixedModel is the reference of the Fixed tests: the entries in key order.
type fixedModel[T comparable] []struct {
	key string
	val T
}

func (m fixedModel[T]) find(key string) (int, bool) {
	i, _ := slices.BinarySearchFunc(m, key, func(e struct {
		key string
		val T
	}, k string) int {
		return strings.Compare(e.key, k)
	})
	return i, i < len(m) && m[i].key == key
}

// fixedEnd returns the content of a page: its keys and its values (which sit at the end of the
// object, with zeros between).
func fixedEnd[T comparable](p *Fixed) int {
	return p.Used() + p.Len()*int(unsafe.Sizeof(*new(T)))
}

// checkFixed compares a page with the model and checks its invariants.
func checkFixed[T comparable](t *testing.T, p *Fixed, m fixedModel[T]) {
	t.Helper()
	if p == nil {
		if len(m) != 0 {
			t.Fatalf("page is nil, model has %d entries", len(m))
		}
		return
	}
	if p.Len() != len(m) {
		t.Fatalf("Len %d, model %d", p.Len(), len(m))
	}
	i := 0
	cp := string(p.CP())
	p.Each(func(rem []byte, v T, _ bool) bool {
		if got := cp + string(rem); got != m[i].key || v != m[i].val {
			t.Fatalf("entry %d: %q -> %v, model %q -> %v", i, got, v, m[i].key, m[i].val)
		}
		i++
		return true
	})
	end := fixedEnd[T](p)
	if end > p.Size() {
		t.Fatalf("content %d, the object has %d bytes", end, p.Size())
	}
	mem := p.mem()
	vs := p.vs(int(unsafe.Sizeof(*new(T))))
	if bytes.Count(mem[p.Used():vs], []byte{0}) != vs-p.Used() {
		t.Fatalf("bytes between the keys and the values are not zero")
	}
	for _, e := range m {
		if v, ok := p.Get[T]([]byte(e.key)); !ok || v != e.val {
			t.Fatalf("Get(%q) = %v, %v; want %v", e.key, v, ok, e.val)
		}
	}
}

func expectFixedInsert[T comparable](m fixedModel[T], cp string, key string, v T) Result {
	if !strings.HasPrefix(key, cp) {
		return Outside
	}
	if i, ok := m.find(key); ok && m[i].val == v {
		return Present
	}
	rems := len(key) - len(cp)
	for _, e := range m {
		rems += len(e.key) - len(cp)
	}
	if len(key)-len(cp) > MaxRemainder || NeedFixed[T](len(m)+1, len(cp), rems) > 512 {
		return Full
	}
	return Added
}

func runFixedModel[T comparable](t *testing.T, gen func(r *rand.Rand) T) {
	for seed := range uint64(40) {
		r := rand.New(rand.NewPCG(seed, 11))
		prefix := randKey(r, "")
		prefix = prefix[:min(len(prefix), r.IntN(3))]
		var m fixedModel[T]
		for range 1 + r.IntN(4) {
			k := randKey(r, prefix)
			if i, ok := m.find(k); !ok {
				m = slices.Insert(m, i, struct {
					key string
					val T
				}{k, gen(r)})
			}
		}
		rests := make([][]byte, len(m))
		vals := make([]T, len(m))
		for i, e := range m {
			rests[i], vals[i] = []byte(e.key), e.val
		}
		p := BuildFixed(rests, vals)
		cp := string(p.CP())
		checkFixed(t, p, m)
		for step := range 600 {
			key := randKey(r, cp)
			if r.IntN(20) == 0 {
				key = randKey(r, "")
			}
			v := gen(r)
			if i, ok := m.find(key); ok && r.IntN(2) == 0 {
				v = m[i].val
			}
			if r.IntN(3) > 0 {
				q, rm := p.Remove(key2b(key), v, HoldsPointers[T]())
				ok := rm != Absent
				i, found := m.find(key)
				want := found && strings.HasPrefix(key, cp) && m[i].val == v
				if ok != want {
					t.Fatalf("seed %d step %d: Remove(%q, %v) = %v, want %v", seed, step, key, v, ok, want)
				}
				if want {
					m = slices.Delete(m, i, i+1)
				}
				p = q
				if p == nil {
					break
				}
			} else {
				if i, ok := m.find(key); ok { // these models hold one value a key
					v = m[i].val
				}
				want := expectFixedInsert(m, cp, key, v)
				q, res := p.Add(key2b(key), v, HoldsPointers[T]())
				if res != want {
					t.Fatalf("seed %d step %d: Insert(%q, %v) = %v, want %v", seed, step, key, v, res, want)
				}
				if res == Added {
					i, _ := m.find(key)
					m = slices.Insert(m, i, struct {
						key string
						val T
					}{key, v})
				} else if q != p {
					t.Fatalf("seed %d step %d: another page for %v", seed, step, res)
				}
				p = q
			}
			checkFixed(t, p, m)
		}
	}
}

func key2b(s string) []byte { return []byte(s) }

type pair struct {
	a uint16
	b uint32
}

// TestFixedAgainstModel keeps a multi-key page of fixed-size values and a plain sorted
// list side by side through thousands of random inserts and removes, for values of
// different sizes and alignments, and compares them after every step.
//
// For a user of the multimap this is the guarantee that the page for numbers (ids,
// counters, offsets) answers exactly like the simple structure it stands for. The
// page is the multi-key page of the redesign (step 4.1) for values without pointers.
//
// Expected: the result of every operation is the one the model predicts, the
// entries are equal, the values sit where their alignment wants them, nothing is
// left behind the used part, whether T is a byte, a pair, a word or two words.
func TestFixedAgainstModel(t *testing.T) {
	t.Run("uint8", func(t *testing.T) { runFixedModel(t, func(r *rand.Rand) uint8 { return uint8(r.IntN(4)) }) })
	t.Run("uint16", func(t *testing.T) { runFixedModel(t, func(r *rand.Rand) uint16 { return uint16(r.IntN(4)) }) })
	t.Run("uint64", func(t *testing.T) { runFixedModel(t, func(r *rand.Rand) uint64 { return uint64(r.IntN(4)) << 40 }) })
	t.Run("pair of 8 bytes with alignment 4", func(t *testing.T) {
		runFixedModel(t, func(r *rand.Rand) pair { return pair{uint16(r.IntN(3)), uint32(r.IntN(2))} })
	})
	t.Run("two words", func(t *testing.T) {
		runFixedModel(t, func(r *rand.Rand) [2]uint64 { return [2]uint64{uint64(r.IntN(3)), 7} })
	})
}

// TestFixedLimits shows which sets of entries make a page of fixed-size values.
//
// The tree asks BuildFixed whether the entries of a subtree fit one page before it
// decides between a page and a byte node, and a map of a type that has no page must
// be told so.
//
// Expected: nil for a type that is not supported (a string, three words),
// for no entry, a different number of values, 256 entries, a remainder of 256 bytes
// and content beyond 512 bytes; a page at the borders.
func TestFixedLimits(t *testing.T) {
	b := func(ss ...string) [][]byte {
		out := make([][]byte, len(ss))
		for i, s := range ss {
			out[i] = []byte(s)
		}
		return out
	}
	many := make([][]byte, 256)
	manyV := make([]uint8, 256)
	for i := range many {
		many[i] = []byte{byte(i)}
	}
	if BuildFixed(b("a", "b"), []string{"x", "y"}) != nil || BuildFixed(b("a"), [][3]uint64{{}}) != nil || BuildFixed(b("a"), []*int{nil}) == nil {
		t.Error("BuildFixed built a page for a type without one")
	}
	if Supported[string]() || !Supported[*int]() || HoldsPointers[uint64]() || !HoldsPointers[*int]() || !Supported[uint64]() || !Supported[[2]uint64]() || Supported[[3]uint64]() {
		t.Error("Supported")
	}
	if BuildFixed[uint64](nil, nil) != nil || BuildFixed(b("a", "b"), []uint64{1}) != nil {
		t.Error("BuildFixed built a page for no entry or the wrong number of values")
	}
	if BuildFixed(many, manyV) != nil || BuildFixed(many[:100], manyV[:100]) == nil {
		t.Error("256 entries or 100 entries of one byte")
	}
	if BuildFixed(b("a", "a"+strings.Repeat("x", 255)), []uint8{1, 2}) != nil || BuildFixed(b("a", "a"+strings.Repeat("x", 254)), []uint8{1, 2}) == nil {
		t.Error("remainder of 255 or 254 bytes")
	}
	if p := BuildFixed(b("a", "a"+strings.Repeat("x", 100)), []uint64{1, 2}); p == nil || p.Size() != 128 {
		t.Errorf("two entries of 101 bytes: %v", p)
	}
	// content: 3 + n + 0 + n + padding + 8n, n=30: 3+60 -> 64 + 240 = 304 -> 384; n=40: 3+80=83 -> 88+320 = 408; n=45: 3+90=93 -> 96+360 = 456; n=50: 3+100->104+400 = 504; n=51: 3+102->112+408=520
	for n, want := range map[int]bool{50: true, 51: false} {
		rs := make([][]byte, n)
		vs := make([]uint64, n)
		for i := range rs {
			rs[i] = []byte{byte(i)}
		}
		if got := BuildFixed(rs, vs) != nil; got != want {
			t.Errorf("%d entries of one byte: built %v, want %v (need %d)", n, got, want, NeedFixed[uint64](n, 0, n))
		}
	}
	if NeedFixed[uint64](2, 1, 4) != 3+2+1+4+16 {
		t.Errorf("NeedFixed = %d", NeedFixed[uint64](2, 1, 4))
	}
}

// TestFixedInsertResults shows each answer an insert can give on a page of
// fixed-size values.
//
// The tree acts on the answer: Present nothing, AddedValue the count of keys stays, Full a
// burst, Outside a byte node above the page.
//
// Expected: one case for every Result, with the page unchanged where nothing was
// added, and the removals that must refuse.
func TestFixedInsertResults(t *testing.T) {
	fresh := func() *Fixed {
		return BuildFixed([][]byte{[]byte("abc1"), []byte("abc2")}, []uint64{10, 20})
	}
	tests := []struct {
		name string
		key  string
		val  uint64
		want Result
	}{
		{"adds a key under the prefix", "abc3", 30, Added},
		{"adds the key that ends at the prefix", "abc", 30, Added},
		{"knows an entry that is there", "abc1", 10, Present},
		{"adds a second value to a key", "abc1", 11, AddedValue},
		{"refuses a key outside the prefix", "abd1", 10, Outside},
		{"refuses a key shorter than the prefix", "ab", 10, Outside},
		{"refuses a remainder of 256 bytes", "abc" + strings.Repeat("k", 256), 1, Full},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := fresh()
			q, res := p.Add([]byte(tt.key), tt.val, false)
			if res != tt.want {
				t.Fatalf("Insert = %v, want %v", res, tt.want)
			}
			if res != Added && res != AddedValue && (q != p || p.Len() != 2) {
				t.Errorf("the page changed without an addition")
			}
			if v, ok := q.Get[uint64]([]byte(tt.key)); res == Added && (!ok || v != tt.val) {
				t.Errorf("Get after Added = %v, %v", v, ok)
			}
		})
	}
	p := fresh()
	if _, ok := p.Get[uint64]([]byte("zzz")); ok {
		t.Error("Get found a key outside the prefix")
	}
	if _, ok := p.Get[uint64]([]byte("abc9")); ok {
		t.Error("Get found a key that is not there")
	}
	for _, c := range []struct {
		key string
		val uint64
	}{{"abd1", 10}, {"abc9", 10}, {"abc1", 20}} {
		if q, rm := p.Remove([]byte(c.key), c.val, false); rm != Absent || q != p {
			t.Errorf("Remove(%q, %d) took something it must not", c.key, c.val)
		}
	}
	if q, rm := BuildFixed([][]byte{[]byte("k")}, []uint64{1}).Remove([]byte("k"), 1, false); rm != Gone || q != nil {
		t.Error("Remove of the only entry")
	}
}

// TestFixedGrowsThroughTheClasses adds entries one after the other and shows that
// the page takes the smallest class that holds one more and keeps its entries.
//
// A tree builds its pages this way, one insert at a time, and an entry must never
// get lost when the object is copied into a larger one.
//
// Expected: the sizes go 32, 64, 128, 256, 384, 512 in this order and the page is
// refused only when 512 bytes cannot hold one more.
func TestFixedGrowsThroughTheClasses(t *testing.T) {
	p := BuildFixed([][]byte{{0}, {255}}, []uint64{0, 255})
	m := fixedModel[uint64]{{"\x00", 0}, {"\xff", 255}}
	sizes := []int{p.Size()}
	for i := 1; i < 255; i++ {
		q, res := p.Add([]byte{byte(i)}, uint64(i), false)
		if res == Full {
			break
		}
		p = q
		m = slices.Insert(m, i, struct {
			key string
			val uint64
		}{string([]byte{byte(i)}), uint64(i)})
		if p.Size() != sizes[len(sizes)-1] {
			sizes = append(sizes, p.Size())
		}
		checkFixed(t, p, m)
	}
	if !slices.Equal(sizes, []int{32, 64, 128, 256, 384, 512}) || p.Len() != 50 {
		t.Errorf("sizes %v, %d entries", sizes, p.Len())
	}
}

// TestFixedShrinksWithHysteresis removes entries from a full page and shows that it
// changes class only when the content fills half of the smaller one.
//
// Expected: the page shrinks as a page of strings does, and a page at the border does
// not change its object when an entry is added and removed again.
func TestFixedShrinksWithHysteresis(t *testing.T) {
	p := BuildFixed([][]byte{{0}, {255}}, []uint64{0, 255})
	last := 0
	for i := 1; ; i++ {
		q, res := p.Add([]byte{byte(i)}, uint64(i), false)
		if res == Full {
			break
		}
		p, last = q, i
	}
	classes := []int{p.Size()}
	for i := last; i >= 1; i-- {
		var rm Removal
		if p, rm = p.Remove([]byte{byte(i)}, uint64(i), false); rm == Absent {
			t.Fatalf("entry %d not removed", i)
		}
		if c := classFor(2 * fixedEnd[uint64](p)); c >= 0 && c < p.class() {
			t.Fatalf("content %d in %d bytes should have shrunk", fixedEnd[uint64](p), p.Size())
		}
		if p.Size() != classes[len(classes)-1] {
			classes = append(classes, p.Size())
		}
	}
	if !slices.Equal(classes, []int{512, 384, 256, 128, 64}) {
		t.Errorf("classes passed: %v", classes)
	}
	size := p.Size()
	for range 5 {
		q, res := p.Add([]byte{9}, 9, false)
		if res != Added {
			t.Fatalf("insert: %v", res)
		}
		r, rm := q.Remove([]byte{9}, 9, false)
		ok := rm != Absent
		if !ok || r.Size() != size {
			t.Fatalf("remove after insert: %v, size %d, was %d", ok, r.Size(), size)
		}
		p = r
	}
}

// TestFixedSkipAndPrepend moves a page down and up in the tree: a byte node arrives
// above it and takes bytes of its common prefix, or goes away and gives them back.
//
// The values move with the end of the keys, as far as their alignment wants. The page
// keeps its entries through both, may need a larger object for Prepend, and refuses
// at 255 bytes of prefix or 512 bytes of content.
//
// Expected: the entries read as the keys without their first bytes after Skip, and as
// before after Prepend; Prepend grows the class when needed and refuses beyond the limits.
func TestFixedSkipAndPrepend(t *testing.T) {
	p := BuildFixed([][]byte{[]byte("abcdef1"), []byte("abcdef2")}, []uint64{1, 2})
	p.Skip[uint64](2)
	checkFixed(t, p, fixedModel[uint64]{{"cdef1", 1}, {"cdef2", 2}})
	p.Skip[uint64](4)
	checkFixed(t, p, fixedModel[uint64]{{"1", 1}, {"2", 2}})
	q := p.Prepend[uint64]([]byte("abcd"), false)
	checkFixed(t, q, fixedModel[uint64]{{"abcd1", 1}, {"abcd2", 2}})
	if q != p {
		t.Error("Prepend allocated although the class holds it")
	}
	big := q.Prepend[uint64]([]byte(strings.Repeat("P", 40)), false)
	if big == nil || big == q || big.Size() <= q.Size() {
		t.Fatalf("Prepend of 40 bytes: %v", big)
	}
	checkFixed(t, big, fixedModel[uint64]{{strings.Repeat("P", 40) + "abcd1", 1}, {strings.Repeat("P", 40) + "abcd2", 2}})
	if big.Prepend[uint64]([]byte(strings.Repeat("P", MaxPrefix)), false) != nil {
		t.Error("Prepend beyond MaxPrefix")
	}
	if big.Prepend[uint64]([]byte(strings.Repeat("P", 470)), false) != nil {
		t.Error("Prepend beyond the largest class")
	}
}

// TestFixedEach covers the early stop of a scan.
//
// Expected: Each stops at the first false.
func TestFixedEach(t *testing.T) {
	p := BuildFixed([][]byte{[]byte("a"), []byte("b")}, []uint64{1, 2})
	n := 0
	if p.Each(func(rem []byte, v uint64, _ bool) bool { n++; return false }) || n != 1 {
		t.Errorf("Each did not stop: %d", n)
	}
}

// FuzzFixed runs the page against the model on operations chosen by the fuzzer.
//
// Expected: the page and the model agree after every operation.
func FuzzFixed(f *testing.F) {
	f.Add([]byte("abc\x00x1abd\x00y2\x01abc"))
	f.Fuzz(func(t *testing.T, data []byte) {
		var p *Fixed
		var m fixedModel[uint32]
		for len(data) >= 3 {
			op, kl, v := data[0]%3, int(data[1])%9, uint32(data[2]%3)
			data = data[3:]
			if len(data) < kl {
				return
			}
			key := string(data[:kl])
			data = data[kl:]
			i, found := m.find(key)
			switch {
			case p == nil && op != 1:
				p = BuildFixed([][]byte{[]byte(key)}, []uint32{v})
				m = fixedModel[uint32]{{key, v}}
			case p == nil:
			case op == 1:
				q, rm := p.Remove([]byte(key), v, false)
				ok := rm != Absent
				if want := found && strings.HasPrefix(key, string(p.CP())) && m[i].val == v; ok != want {
					t.Fatalf("Remove %q: %v, want %v", key, ok, want)
				} else if want {
					m = slices.Delete(m, i, i+1)
				}
				p = q
			default:
				if found { // this model holds one value a key
					v = m[i].val
				}
				want := expectFixedInsert(m, string(p.CP()), key, v)
				q, res := p.Add([]byte(key), v, false)
				if res != want {
					t.Fatalf("Insert %q: %v, want %v", key, res, want)
				}
				if res == Added {
					m = slices.Insert(m, i, struct {
						key string
						val uint32
					}{key, v})
				}
				p = q
			}
			if p == nil {
				m = nil
			}
			checkFixed(t, p, m)
		}
	})
}
