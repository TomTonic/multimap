package skpage

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

// rec is the record that the pointer values of the tests point to.
type rec struct{ id, aux uint64 }

// one is a struct with a single pointer: a word with a pointer, as a *rec is.
type one struct{ p *rec }

// padded is a value of 4 bytes with 2-byte alignment and padding inside.
type padded struct {
	a uint8
	b uint16
}

// pool holds the records the pointer values of a test point to, so that they stay alive.
var pool = func() []*rec {
	out := make([]*rec, 4096)
	for i := range out {
		out[i] = &rec{id: uint64(i)}
	}
	return out
}()

// valueCases are the types of value the Fixed page takes, each with a way to make
// the i-th different value of it.
type valueCase struct {
	name string
	run  func(t *testing.T, rests []int, steps int)
}

func valueCases() []valueCase {
	return []valueCase{
		{"uint64", modelTest(func(i int) uint64 { return uint64(i) * 0x9e3779b97f4a7c15 })},
		{"uint32", modelTest(func(i int) uint32 { return uint32(i) + 1 })},
		{"uint8", modelTest(func(i int) uint8 { return uint8(i) })},
		{"[3]byte", modelTest(func(i int) [3]byte { return [3]byte{byte(i), byte(i >> 8), 7} })},
		{"padded", modelTest(func(i int) padded { return padded{uint8(i), uint16(i >> 8)} })},
		{"[2]uint64", modelTest(func(i int) [2]uint64 { return [2]uint64{uint64(i), 5} })},
		{"*rec", modelTest(func(i int) *rec { return pool[i%len(pool)] })},
		{"one", modelTest(func(i int) one { return one{pool[i%len(pool)]} })},
		{"unsafe.Pointer", modelTest(func(i int) unsafe.Pointer { return unsafe.Pointer(pool[i%len(pool)]) })},
	}
}

// verifyFixed fails unless the page p holds exactly the key remainder rest of a
// key of kl bytes and the values want, in an object whose bytes outside head,
// remainder and values are zero.
func verifyFixed[T comparable](t *testing.T, p *Fixed, rest []byte, kl int, want map[T]bool) {
	t.Helper()
	vs := p.Values[T]()
	got := map[T]bool{}
	for _, v := range vs {
		got[v] = true
	}
	if p.Len() != len(want) || len(vs) != len(want) || len(got) != len(want) {
		t.Fatalf("page holds %d values (%d distinct), want %d", len(vs), len(got), len(want))
	}
	for v := range want {
		if !got[v] || !p.Has(v) {
			t.Fatalf("value %v is missing", v)
		}
	}
	n := 0
	if !p.Each(func(v T) bool { n++; return want[v] }) || n != len(want) {
		t.Fatalf("Each ran over %d of %d values", n, len(want))
	}
	if n = 0; p.Each(func(T) bool { n++; return false }) || n != 1 {
		t.Fatalf("Each did not stop at the first value: %d calls", n)
	}
	if p.KeyLen() != kl || !bytes.Equal(p.Rest(), rest) {
		t.Fatalf("remainder %q of a key of %d bytes, want %q of %d", p.Rest(), p.KeyLen(), rest, kl)
	}
	key := append(bytes.Repeat([]byte("p"), kl-len(rest)), rest...)
	if !p.Match(key) || p.Match(key[1:]) {
		t.Fatal("Match is wrong")
	}
	c := p.Class()
	if p.Size() != sizes[c] || capacityOf[T](c, len(rest)) < len(want) {
		t.Fatalf("class %d of %d bytes holds %d values, wants %d", c, p.Size(), capacityOf[T](c, len(rest)), len(want))
	}
	m, off := p.mem(), valuesAt[T](len(rest))
	var z T
	end := off + len(want)*int(unsafe.Sizeof(z))
	for i, b := range m {
		if inside := i < Header+len(rest) || (i >= off && i < end); !inside && b != 0 {
			t.Fatalf("byte %d of the object is %d, outside head, remainder and values", i, b)
		}
	}
}

// modelTest returns the test that drives pages of T with random operations and
// compares them with a set.
func modelTest[T comparable](mk func(i int) T) func(t *testing.T, rests []int, steps int) {
	return func(t *testing.T, rests []int, steps int) {
		for _, rl := range rests {
			if rl > MaxRemainderFixed[T]() {
				continue
			}
			r := rand.New(rand.NewSource(int64(rl) + 1))
			rest := make([]byte, rl)
			for i := range rest {
				rest[i] = byte('a' + r.Intn(26))
			}
			const kl = 700 // the whole key: longer than any remainder
			want := map[T]bool{}
			maxN := capacityOf[T](Classes-1, rl)
			first := mk(0)
			p := NewFixed(rest, kl, first)
			if p == nil {
				t.Fatalf("remainder of %d bytes: no page", rl)
			}
			want[first] = true
			verifyFixed(t, p, rest, kl, want)
			for range steps {
				v := mk(r.Intn(2*maxN + 4))
				switch op := r.Intn(10); {
				case op < 5:
					q, res := p.Add(v)
					switch {
					case want[v]:
						if res != Present || q != p {
							t.Fatalf("Add of a present value: %v", res)
						}
					case len(want) == maxN:
						if res != Full || q != p {
							t.Fatalf("Add to a full page: %v", res)
						}
					default:
						if res != Added {
							t.Fatalf("Add of a new value: %v", res)
						}
						want[v] = true
						p = q
					}
				case op < 9:
					q, ok := p.Remove(v)
					if ok != want[v] {
						t.Fatalf("Remove of %v: %v, was there: %v", v, ok, want[v])
					}
					if ok {
						delete(want, v)
					}
					if q == nil {
						if len(want) != 0 {
							t.Fatal("Remove gave no page for a page that has values")
						}
						p = NewFixed(rest, kl, v)
						want[v] = true
					} else {
						p = q
					}
				default:
					pre := make([]byte, r.Intn(12))
					for i := range pre {
						pre[i] = byte('A' + r.Intn(26))
					}
					q := p.Prepend[T](pre)
					if fits := len(rest)+len(pre) <= MaxRemainderFixed[T]() && classHolding[T](len(want), len(rest)+len(pre)) >= 0; fits != (q != nil) {
						t.Fatalf("Prepend of %d bytes to a remainder of %d with %d values: %v", len(pre), len(rest), len(want), q != nil)
					}
					if q != nil {
						rest, p = append(slices.Clone(pre), rest...), q
						maxN = capacityOf[T](Classes-1, len(rest))
					}
				}
				verifyFixed(t, p, rest, kl, want)
			}
		}
	}
}

// TestFixedModel covers the values of one key as a user of the tree sees them:
// a set that grows and shrinks, which stays right whatever the size of the
// values and the length of the key's remainder. The page belongs to the single-key
// page of the tree for fixed-size values (step 3.5). Random additions, removals and
// moves up of a page are compared with a set, for every type of value the page
// takes and remainders that put the values at every alignment, and each step
// checks the page's invariants: values, remainder, the size class, and zeros
// outside head, remainder and values.
func TestFixedModel(t *testing.T) {
	rests := []int{0, 1, 2, 5, 6, 7, 8, 9, 17, 25, 58, 59, 100, 250, 498}
	for _, vc := range valueCases() {
		t.Run(vc.name, func(t *testing.T) { vc.run(t, rests, 3000) })
	}
}

// TestFixedSupported covers which values the Fixed page takes, as a user of the
// map sees it: small values without pointers, and one word that is a pointer;
// not an empty value, a string, an interface or a struct of two words.
func TestFixedSupported(t *testing.T) {
	type two [2]*rec
	for _, tc := range []struct {
		name string
		got  bool
		want bool
	}{
		{"uint64 is supported", Supported[uint64](), true},
		{"a struct of 16 bytes without pointers is supported", Supported[[2]uint64](), true},
		{"a value of 17 bytes is not", Supported[[17]byte](), false},
		{"an empty struct is not", Supported[struct{}](), false},
		{"a pointer is supported", Supported[*rec](), true},
		{"a struct of one pointer is supported", Supported[one](), true},
		{"unsafe.Pointer is supported", Supported[unsafe.Pointer](), true},
		{"a string is not", Supported[string](), false},
		{"a struct of two pointers is not", Supported[two](), false},
		{"an interface is not", Supported[any](), false},
	} {
		if tc.got != tc.want {
			t.Errorf("%s: got %v", tc.name, tc.got)
		}
	}
	if HoldsPointers[uint64]() || !HoldsPointers[*rec]() || HoldsPointers[string]() {
		t.Error("HoldsPointers is right for a pointer only")
	}
}

// TestFixedLimits covers what the Fixed page refuses, so that the tree knows to
// use a value overflow: a type without a page, a remainder that leaves no room for
// a value, more values than the largest class holds, and the longest
// remainder of each kind of value.
func TestFixedLimits(t *testing.T) {
	if NewFixed[string](nil, 1, "a") != nil || BuildFixed[string](nil, 1, []string{"a"}) != nil {
		t.Error("a page for a string")
	}
	for _, tc := range []struct {
		name string
		got  int
		want int
	}{
		{"uint64", MaxRemainderFixed[uint64](), 498},
		{"*rec", MaxRemainderFixed[*rec](), 498},
		{"uint8", MaxRemainderFixed[uint8](), 505},
		{"[2]uint64", MaxRemainderFixed[[2]uint64](), 490},
		{"[3]byte", MaxRemainderFixed[[3]byte](), 503},
	} {
		if tc.got != tc.want {
			t.Errorf("longest remainder for %s: %d, want %d", tc.name, tc.got, tc.want)
		}
	}
	long := bytes.Repeat([]byte("r"), 499)
	if NewFixed(long, 600, uint64(1)) != nil || BuildFixed(long, 600, []uint64{1}) != nil {
		t.Error("a remainder of 499 bytes leaves no room for a uint64")
	}
	if NewFixed(long[:498], 600, uint64(1)) == nil {
		t.Error("a remainder of 498 bytes fits")
	}
	if EmptyFixed[string](nil, 1) != nil || EmptyFixed[uint64](long, 600) != nil {
		t.Error("an empty page for a string or for a remainder of 499 bytes")
	}
	if e := EmptyFixed[uint64](long[:498], 600); e == nil || e.Len() != 0 || e.Size() != 512 {
		t.Errorf("an empty page for a remainder of 498 bytes: %v", e)
	} else if f, res := e.Add(uint64(5)); res != Added || f != e || !f.Has(uint64(5)) {
		t.Error("the first value of an empty page")
	}
	if BuildFixed(nil, 3, []uint64(nil)) != nil {
		t.Error("a page without values")
	}
	if BuildFixed(nil, 3, make([]uint64, 64)) != nil {
		t.Error("64 values of 8 bytes do not fit with the 6 bytes of head (8 with the padding): 63 do")
	}
	if p := BuildFixed(nil, 3, slices.Collect(func(yield func(uint64) bool) {
		for i := range 63 {
			if !yield(uint64(i)) {
				return
			}
		}
	})); p == nil || p.Size() != 512 || p.Len() != 63 {
		t.Errorf("63 values of 8 bytes fill a page of 512: %v", p)
	}
	if allocPtr[*rec](0, 4) != nil || allocPtr[*rec](6, 1) != nil || AllocPtr[*rec](0, 4) != nil || AllocPtr[*rec](2, 5) == nil {
		t.Error("allocPtr gives an object for a layout that no page has")
	}
	p := NewFixed(nil, 3, uint64(0))
	for i := 1; i < 63; i++ {
		q, res := p.Add(uint64(i))
		if res != Added {
			t.Fatalf("value %d: %v", i, res)
		}
		p = q
	}
	if q, res := p.Add(uint64(99)); res != Full || q != p || p.Len() != 63 {
		t.Errorf("the 64th value: %v", res)
	}
}

// TestFixedPrependCases covers the page's move up in the tree, when the node
// above it goes away: in place if the values keep their place, a bigger page
// if they do not fit, and for a page with pointers a page of another
// layout when the words in front of the values grow. The values and the key
// stay.
func TestFixedPrependCases(t *testing.T) {
	check := func(t *testing.T, p *Fixed, rest string, n int) {
		t.Helper()
		if string(p.Rest()) != rest || p.Len() != n {
			t.Fatalf("remainder %q, %d values; want %q, %d", p.Rest(), p.Len(), rest, n)
		}
	}
	t.Run("no pointer, values keep their word: in place", func(t *testing.T) {
		p := NewFixed([]byte("abc"), 9, uint64(7)) // values at 16
		q := p.Prepend[uint64]([]byte("zz"))       // 5 bytes: still 16
		if q != p {
			t.Fatal("moved")
		}
		check(t, q, "zzabc", 1)
		if !q.Has(uint64(7)) {
			t.Fatal("value lost")
		}
	})
	t.Run("no pointer, the values move up a word, in place", func(t *testing.T) {
		p := NewFixed([]byte("abc"), 20, uint64(7))
		for i := uint64(8); i < 10; i++ {
			p, _ = p.Add(i)
		} // 3 values in a page of 64: room for 6 at 16
		q := p.Prepend[uint64]([]byte("0123456789"))
		if q != p {
			t.Fatal("moved")
		}
		check(t, q, "0123456789abc", 3)
		for _, v := range []uint64{7, 8, 9} {
			if !q.Has(v) {
				t.Fatalf("value %d lost", v)
			}
		}
		verifyFixed(t, q, []byte("0123456789abc"), 20, map[uint64]bool{7: true, 8: true, 9: true})
	})
	t.Run("no pointer, the values do not fit any more: a bigger page", func(t *testing.T) {
		p := NewFixed([]byte("a"), 90, uint64(7))
		q := p.Prepend[uint64](bytes.Repeat([]byte("x"), 20))
		if q == p || q.Size() != 64 {
			t.Fatalf("got a page of %d bytes", q.Size())
		}
		check(t, q, strings20+"a", 1)
	})
	t.Run("pointers, same words in front: in place", func(t *testing.T) {
		p := NewFixed([]byte("abc"), 9, pool[1]) // 9 bytes: 2 words
		q := p.Prepend[*rec]([]byte("zz"))       // 11 bytes: 2 words
		if q != p {
			t.Fatal("moved")
		}
		check(t, q, "zzabc", 1)
	})
	t.Run("pointers, one more word in front: another layout", func(t *testing.T) {
		p := NewFixed([]byte("abc"), 9, pool[1])
		q := p.Prepend[*rec](bytes.Repeat([]byte("x"), 8)) // 17 bytes: 3 words
		if q == p {
			t.Fatal("not moved")
		}
		check(t, q, "xxxxxxxxabc", 1)
		if !q.Has(pool[1]) {
			t.Fatal("value lost")
		}
	})
	t.Run("too long a remainder: nil", func(t *testing.T) {
		p := NewFixed(bytes.Repeat([]byte("a"), 490), 600, uint64(1))
		if p.Prepend[uint64](bytes.Repeat([]byte("b"), 9)) != nil {
			t.Fatal("remainder of 499 bytes")
		}
	})
}

const strings20 = "xxxxxxxxxxxxxxxxxxxx"

// TestFixedKeepsPointersAlive covers the promise that matters for a map of
// pointers: a record that a page points to is not freed while the page holds
// it, and is freed once the page has let go of it. The remainder, which is only
// bytes, is not mistaken for a pointer. Every class and every number of words in
// front of the values (166 layouts) is built, filled, and checked after a
// garbage collection; an object whose type had its pointers marked one word
// off would keep a record that the page does not hold, or lose one that it
// does.
func TestFixedKeepsPointersAlive(t *testing.T) {
	for c, size := range sizes {
		w := size / 8
		for j := 1; j < w; j++ {
			t.Run(fmt.Sprintf("class %d, %d words in front", c, j), func(t *testing.T) {
				p, held, decoy := fillWithRecords(c, j)
				for range 2 {
					runtime.GC()
				}
				for i, wp := range held {
					if wp.Value() == nil {
						t.Fatalf("record %d of %d was freed while the page holds it", i, len(held))
					}
				}
				if decoy != nil && decoy.Value() != nil {
					t.Fatal("a record that only the bytes of the remainder point to was kept: the remainder is scanned")
				}
				// let go of the first record
				q, ok := removeFirst(p, held)
				if !ok {
					t.Fatal("Remove did not find the value")
				}
				for range 2 {
					runtime.GC()
				}
				if held[0].Value() != nil {
					t.Fatal("the removed record is still held")
				}
				if q != nil && !q.Has(held[1].Value()) {
					t.Fatal("another value was lost")
				}
				runtime.KeepAlive(q)
			})
		}
	}
}

// fillWithRecords returns a page of class c with j words in front of the
// values, full of values that are pointers to records which only the page
// holds, and weak pointers to them and, for j of at least 2, to a decoy record whose
// address is the first bytes of the remainder. It does not inline, so that no
// stack slot of the caller holds a record.
//
//go:noinline
func fillWithRecords(c, j int) (*Fixed, []weak.Pointer[rec], *weak.Pointer[rec]) {
	rest := make([]byte, 8*j-Header)
	var decoy *weak.Pointer[rec]
	if len(rest) >= 8 {
		d := &rec{id: 1 << 40}
		w := weak.Make(d)
		decoy = &w
		*(*uintptr)(unsafe.Pointer(&rest[0])) = uintptr(unsafe.Pointer(d))
	}
	n := capacityOf[*rec](c, len(rest))
	vals := make([]*rec, n)
	held := make([]weak.Pointer[rec], n)
	for i := range vals {
		vals[i] = &rec{id: uint64(i)}
		held[i] = weak.Make(vals[i])
	}
	p := BuildFixed(rest, 8*j, vals)
	if p == nil || p.Class() != c {
		panic("test setup")
	}
	return p, held, decoy
}

// TestBackFits covers when a key may go back into a page from its value overflow:
// when its values take at most half the room that the largest page has for them,
// whatever the remainder.
func TestBackFits(t *testing.T) {
	for _, tc := range []struct {
		name     string
		rem, val int
		want     bool
	}{
		{"short remainder, half the room", 0, 253, true},
		{"short remainder, one byte more", 0, 254, false},
		{"long remainder: the room is smaller", 400, 53, true},
		{"long remainder, one byte more", 400, 54, false},
		{"no remainder, no values", 0, 0, true},
	} {
		if got := BackFits(tc.rem, tc.val); got != tc.want {
			t.Errorf("%s: %v", tc.name, got)
		}
	}
}

// FuzzFixed drives pages of uint64 with the operations a byte string spells,
// and checks them against a set.
func FuzzFixed(f *testing.F) {
	f.Add([]byte{0, 1, 2, 3, 4, 5, 6, 7}, uint8(3))
	f.Add(bytes.Repeat([]byte{0, 9, 1, 8}, 40), uint8(200))
	f.Fuzz(func(t *testing.T, ops []byte, rl uint8) {
		rest := bytes.Repeat([]byte("k"), int(rl)%MaxRemainderFixed[uint64]())
		want := map[uint64]bool{0: true}
		p := NewFixed(rest, len(rest)+1, uint64(0))
		for i := 0; i+1 < len(ops); i += 2 {
			v := uint64(ops[i+1]) % 80
			if ops[i]%3 == 0 {
				q, ok := p.Remove(v)
				if ok != want[v] {
					t.Fatalf("Remove: %v", ok)
				}
				delete(want, v)
				if q == nil {
					p = NewFixed(rest, len(rest)+1, v)
					want[v] = true
				} else {
					p = q
				}
			} else if q, res := p.Add(v); res == Added {
				want[v], p = true, q
			} else if res == Present != want[v] {
				t.Fatalf("Add: %v", res)
			}
			verifyFixed(t, p, rest, len(rest)+1, want)
		}
	})
}

// removeFirst removes the first record of held from p. It does not inline, so
// that no stack slot of the caller holds the record afterwards.
//
//go:noinline
func removeFirst(p *Fixed, held []weak.Pointer[rec]) (*Fixed, bool) {
	return p.Remove(held[0].Value())
}

// TestFixedHysteresis covers a key whose number of values hovers at the border
// of a size class, which must not get a new object with every value that comes
// and goes: the page grows by one class when it is full, shrinks only when its
// values fill at most half of the smaller class, and so is not copied by a value
// that is added and removed again. Values of 8 bytes without a remainder: 3 fit 32
// bytes, 7 fit 64.
func TestFixedHysteresis(t *testing.T) {
	p := NewFixed(nil, 1, uint64(0))
	for i := uint64(1); i < 3; i++ {
		p, _ = p.Add(i)
	}
	if p.Size() != 32 {
		t.Fatalf("three values take %d bytes", p.Size())
	}
	q, _ := p.Add(3)
	if q == p || q.Size() != 64 {
		t.Fatalf("the fourth value gives a page of %d bytes, the same: %v", q.Size(), q == p)
	}
	for range 5 { // added and removed again: the page stays
		r, _ := q.Remove(3)
		if r != q || r.Size() != 64 {
			t.Fatal("a page of 64 bytes with three values shrank")
		}
		r, _ = r.Add(3)
		if r != q {
			t.Fatal("a page with room was copied for an added value")
		}
	}
	for v := uint64(3); v > 1; v-- { // down to two values: 4 would fill the smaller class
		q, _ = q.Remove(v)
		if q.Size() != 64 {
			t.Fatalf("%d values left: the page shrank to %d bytes", v-1, q.Size())
		}
	}
	if q, _ = q.Remove(1); q.Size() != 32 || q.Len() != 1 {
		t.Fatalf("one value left: a page of %d bytes with %d values", q.Size(), q.Len())
	}
}
