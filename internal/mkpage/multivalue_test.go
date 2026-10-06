package mkpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

// mvEntry is a key of the reference model of the pages that hold several values a key: the
// key from the end of the path on, and its values in the order they came in.
type mvEntry struct {
	key  string
	vals []string
}

type mvModel []mvEntry

func (m mvModel) find(key string) (int, bool) {
	i := sort.Search(len(m), func(i int) bool { return m[i].key >= key })
	return i, i < len(m) && m[i].key == key
}

func (m mvModel) slots() (n int) {
	for _, e := range m {
		n += len(e.vals)
	}
	return n
}

// rests returns the model as a page is built from it: the key once for every value.
func (m mvModel) rests() (rests, vals [][]byte) {
	for _, e := range m {
		for _, v := range e.vals {
			rests, vals = append(rests, []byte(e.key)), append(vals, []byte(v))
		}
	}
	return rests, vals
}

// clone returns a copy that shares nothing with m.
func (m mvModel) clone() mvModel {
	out := make(mvModel, len(m))
	for i, e := range m {
		out[i] = mvEntry{e.key, slices.Clone(e.vals)}
	}
	return out
}

// add adds the value to the key's values, and reports whether it was new.
func (m mvModel) add(key, val string) (mvModel, bool) {
	i, ok := m.find(key)
	if !ok {
		return slices.Insert(m, i, mvEntry{key, []string{val}}), true
	}
	if slices.Contains(m[i].vals, val) {
		return m, false
	}
	m[i].vals = append(m[i].vals, val)
	return m, true
}

// remove takes the value out of the key's values, and the key out once it has none.
func (m mvModel) remove(key, val string) (mvModel, bool) {
	i, ok := m.find(key)
	if !ok {
		return m, false
	}
	j := slices.Index(m[i].vals, val)
	if j < 0 {
		return m, false
	}
	m[i].vals = slices.Delete(m[i].vals, j, j+1)
	if len(m[i].vals) == 0 {
		m = slices.Delete(m, i, i+1)
	}
	return m, true
}

// mvFlat is one value of a page as Each reports it.
type mvFlat struct {
	key, val string
	first    bool
}

// mvPage is a page of either flavor as the tests of the values of a key see it.
type mvPage interface {
	add(key, val string) Result
	insert(key, val string) Result
	remove(key, val string) bool
	widen(key, val string) bool
	gone() bool
	cp() string
	need(slots, remBytes, valBytes int) int
	tooLong(key, val string) bool
	flat() []mvFlat
	values(key string) ([]string, bool)
	get(key string) (string, bool)
	counts() (slots, keys int)
	invariants(t *testing.T)
	build(m mvModel) mvPage
	value(r *rand.Rand) string
}

type strPage struct{ p *Page }

func (s *strPage) build(m mvModel) mvPage {
	rests, vals := m.rests()
	return &strPage{BuildStrings(rests, vals)}
}
func (s *strPage) value(r *rand.Rand) string {
	switch r.IntN(40) {
	case 0:
		return strings.Repeat("L", MaxValue+1)
	case 1:
		return strings.Repeat("M", 60+r.IntN(100))
	}
	return randVal(r)
}
func (s *strPage) add(key, val string) Result {
	q, res := s.p.Add([]byte(key), []byte(val))
	s.p = q
	return res
}
func (s *strPage) insert(key, val string) Result {
	q, res := s.p.Insert([]byte(key), []byte(val))
	s.p = q
	return res
}
func (s *strPage) remove(key, val string) bool {
	q, ok := s.p.Remove([]byte(key), []byte(val))
	s.p = q
	return ok
}
func (s *strPage) widen(key, val string) bool {
	q := s.p.Widen([]byte(key), []byte(val))
	if q != nil {
		s.p = q
	}
	return q != nil
}
func (s *strPage) gone() bool { return s.p == nil }
func (s *strPage) cp() string { return string(s.p.CP()) }
func (s *strPage) need(slots, remBytes, valBytes int) int {
	return NeedStrings(slots, s.p.PrefixLen(), remBytes, valBytes)
}
func (s *strPage) tooLong(key, val string) bool { return len(val) > MaxValue }
func (s *strPage) flat() (out []mvFlat) {
	cp := string(s.p.CP())
	s.p.Each(func(rem, val []byte, first bool) bool {
		out = append(out, mvFlat{cp + string(rem), string(val), first})
		return true
	})
	return out
}
func (s *strPage) values(key string) (vs []string, ok bool) {
	ok = s.p.EachValue([]byte(key), func(v []byte) bool { vs = append(vs, string(v)); return true })
	return vs, ok
}
func (s *strPage) get(key string) (string, bool) {
	v, ok := s.p.Get([]byte(key))
	return string(v), ok
}
func (s *strPage) counts() (int, int) { return s.p.Len(), s.p.Keys() }
func (s *strPage) invariants(t *testing.T) {
	t.Helper()
	used := s.p.Used()
	if used > s.p.Size() {
		t.Fatalf("used %d exceeds the size %d", used, s.p.Size())
	}
	if bytes.Count(s.p.mem()[used:], []byte{0}) != s.p.Size()-used {
		t.Fatal("bytes behind the used part are not zero")
	}
}

type fixedPage struct{ p *Fixed }

func (f *fixedPage) build(m mvModel) mvPage {
	rests, vals := m.rests()
	vs := make([]uint64, len(vals))
	for i, v := range vals {
		vs[i] = fixedVal(string(v))
	}
	return &fixedPage{BuildFixed(rests, vs)}
}

func fixedVal(s string) uint64 { n, _ := strconv.ParseUint(s, 10, 64); return n }

func (f *fixedPage) value(r *rand.Rand) string { return strconv.Itoa(r.IntN(6)) }
func (f *fixedPage) add(key, val string) Result {
	q, res := f.p.Add([]byte(key), fixedVal(val))
	f.p = q
	return res
}
func (f *fixedPage) insert(key, val string) Result {
	q, res := f.p.Insert([]byte(key), fixedVal(val))
	f.p = q
	return res
}
func (f *fixedPage) remove(key, val string) bool {
	q, ok := f.p.Remove([]byte(key), fixedVal(val))
	f.p = q
	return ok
}
func (f *fixedPage) widen(key, val string) bool {
	q := f.p.Widen([]byte(key), fixedVal(val))
	if q != nil {
		f.p = q
	}
	return q != nil
}
func (f *fixedPage) gone() bool { return f.p == nil }
func (f *fixedPage) cp() string { return string(f.p.CP()) }
func (f *fixedPage) need(slots, remBytes, _ int) int {
	return NeedFixed[uint64](slots, f.p.PrefixLen(), remBytes)
}
func (f *fixedPage) tooLong(string, string) bool { return false }
func (f *fixedPage) flat() (out []mvFlat) {
	cp := string(f.p.CP())
	f.p.Each(func(rem []byte, v uint64, first bool) bool {
		out = append(out, mvFlat{cp + string(rem), strconv.FormatUint(v, 10), first})
		return true
	})
	return out
}
func (f *fixedPage) values(key string) (vs []string, ok bool) {
	ok = f.p.EachValue([]byte(key), func(v uint64) bool { vs = append(vs, strconv.FormatUint(v, 10)); return true })
	return vs, ok
}
func (f *fixedPage) get(key string) (string, bool) {
	v, ok := f.p.Get[uint64]([]byte(key))
	return strconv.FormatUint(v, 10), ok
}
func (f *fixedPage) counts() (int, int) { return f.p.Len(), f.p.Keys() }
func (f *fixedPage) invariants(t *testing.T) {
	t.Helper()
	used := f.p.Used()
	vs := valuesAt[uint64](used)
	end := vs + f.p.Len()*int(unsafe.Sizeof(uint64(0)))
	if end > f.p.Size() {
		t.Fatalf("values end at %d, the object has %d bytes", end, f.p.Size())
	}
	mem := f.p.mem()
	if bytes.Count(mem[end:], []byte{0}) != f.p.Size()-end || bytes.Count(mem[used:vs], []byte{0}) != vs-used {
		t.Fatal("bytes behind the used part are not zero")
	}
}

// checkMV compares a page with the model: every value of every key in order, the first
// flags, the counts, the lookups, and the layout invariants.
func checkMV(t *testing.T, pg mvPage, m mvModel) {
	t.Helper()
	if pg.gone() {
		if len(m) != 0 {
			t.Fatalf("page is gone, model has %d keys", len(m))
		}
		return
	}
	var want []mvFlat
	for _, e := range m {
		for i, v := range e.vals {
			want = append(want, mvFlat{e.key, v, i == 0})
		}
	}
	if got := pg.flat(); !slices.Equal(got, want) {
		t.Fatalf("entries %v, model %v", got, want)
	}
	if slots, keys := pg.counts(); slots != m.slots() || keys != len(m) {
		t.Fatalf("Len %d Keys %d, model %d and %d", slots, keys, m.slots(), len(m))
	}
	for _, e := range m {
		if vs, ok := pg.values(e.key); !ok || !slices.Equal(vs, e.vals) {
			t.Fatalf("values of %q: %v, %v; want %v", e.key, vs, ok, e.vals)
		}
		if v, ok := pg.get(e.key); !ok || v != e.vals[0] {
			t.Fatalf("Get(%q) = %q, %v; want %q", e.key, v, ok, e.vals[0])
		}
	}
	pg.invariants(t)
}

// expectAdd says what Add must answer for the model.
func expectAdd(pg mvPage, m mvModel, key, val string) Result {
	cp := pg.cp()
	if !strings.HasPrefix(key, cp) {
		return Outside
	}
	r := key[len(cp):]
	i, exists := m.find(key)
	if exists && slices.Contains(m[i].vals, val) {
		return Present
	}
	rems, vals := 0, len(val)
	for _, e := range m {
		rems += len(e.key) - len(cp)
		for _, v := range e.vals {
			vals += len(v)
		}
	}
	if !exists {
		rems += len(r)
	}
	if pg.tooLong(key, val) || (!exists && len(r) > MaxRemainder) || pg.need(m.slots()+1, rems, vals) > sizes[Classes-1] {
		return Full
	}
	if exists {
		return AddedValue
	}
	return Added
}

// TestMultiValuePageAgainstModel keeps a multi-key page that holds several values of a key
// and a plain model side by side through thousands of random adds, removes and widenings
// from many starting pages, for both flavors, and compares them after every step.
//
// For a user of the multimap this is the guarantee of step 5: a key of the index that gets
// a second, third, tenth value keeps every value, in the order they came in, in the page of
// its neighbours, and a value that goes away leaves the others where they are, whichever of
// the values it is.
//
// Expected: after every operation the result is the one the model predicts, the entries,
// the first flags, the counts and the lookups equal the model's, the layout invariants
// hold, and a refused Widen is one that a page built from all the entries would also refuse.
func TestMultiValuePageAgainstModel(t *testing.T) {
	for _, pg := range []mvPage{&strPage{}, &fixedPage{}} {
		t.Run(fmt.Sprintf("%T", pg), func(t *testing.T) { runMV(t, pg) })
	}
}

func runMV(t *testing.T, flavor mvPage) {
	for seed := range uint64(60) {
		r := rand.New(rand.NewPCG(seed, 13))
		prefix := randKey(r, "")
		prefix = prefix[:min(len(prefix), r.IntN(3))]
		var m mvModel
		for range 1 + r.IntN(4) {
			m, _ = m.add(randKey(r, prefix), flavor.value(r))
		}
		for range r.IntN(4) { // some keys with several values
			m, _ = m.add(m[r.IntN(len(m))].key, flavor.value(r))
		}
		pg := flavor.build(m)
		if pg.gone() {
			continue // the values did not fit one page
		}
		checkMV(t, pg, m)
		for step := range 800 {
			cp := pg.cp()
			key := randKey(r, cp)
			if r.IntN(25) == 0 {
				key = randKey(r, "")
			}
			if i := r.IntN(len(m) + 1); i < len(m) && r.IntN(2) == 0 {
				key = m[i].key // often a key that is there, to give it more values
			}
			val := flavor.value(r)
			if i, ok := m.find(key); ok && r.IntN(3) == 0 {
				val = m[i].vals[r.IntN(len(m[i].vals))]
			}
			switch r.IntN(10) {
			case 0, 1, 2, 3: // remove
				i, found := m.find(key)
				want := found && strings.HasPrefix(key, cp) && slices.Contains(m[i].vals, val)
				ok := pg.remove(key, val)
				if ok != want {
					t.Fatalf("seed %d step %d: Remove(%q, %q) = %v, want %v", seed, step, key, val, ok, want)
				}
				if want {
					m, _ = m.remove(key, val)
				}
				if pg.gone() {
					if len(m) != 0 {
						t.Fatalf("seed %d step %d: page gone, model has %d keys", seed, step, len(m))
					}
				}
			default:
				want := expectAdd(pg, m, key, val)
				if want == Outside && r.IntN(2) == 0 { // the tree would widen the page, if the entries fit
					all, _ := m.clone().add(key, val)
					rests, vals := all.rests()
					vs := make([]uint64, len(vals))
					for i, v := range vals {
						vs[i] = fixedVal(string(v))
					}
					var fits bool
					if _, isStr := flavor.(*strPage); isStr {
						fits = BuildStrings(rests, vals) != nil
					} else {
						fits = BuildFixed(rests, vs) != nil
					}
					done := pg.widen(key, val)
					if done != fits {
						t.Fatalf("seed %d step %d: Widen(%q, %q) = %v, a page built from all the entries: %v", seed, step, key, val, done, fits)
					}
					if done {
						m, _ = m.add(key, val)
					}
					break
				}
				res := pg.add(key, val)
				if res != want {
					t.Fatalf("seed %d step %d: Add(%q, %q) = %v, want %v", seed, step, key, val, res, want)
				}
				if res == Added || res == AddedValue {
					m, _ = m.add(key, val)
				}
			}
			if pg.gone() {
				break
			}
			checkMV(t, pg, m)
		}
	}
}

func bs(ss ...string) [][]byte {
	out := make([][]byte, len(ss))
	for i, s := range ss {
		out[i] = []byte(s)
	}
	return out
}

// exampleKeys are the keys of the layout example of step5-mkmv-design.md: "Bahnhof" is the
// common prefix, "strasse" has three values.
var exampleKeys = bs("Bahnhof", "Bahnhofsallee", "Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofstrasse", "Bahnhofweg")

// TestMultiValuePageLayout pins the bytes of a page that holds a key with several values.
//
// The format is what the tree and every later change of the page rely on: the head of three
// bytes (type, the length of the common prefix, the number of values), the common prefix right
// behind it, a length byte for every value with 255 for a further value of the key before it,
// and the remainders, which only the first value of a key has.
//
// Expected: the page of `Fixed` values is the 128-byte page with the bytes of the design note, the page
// of strings the 64-byte page with its values behind their keys, and a page without a key that
// has several values has no 255 in it.
func TestMultiValuePageLayout(t *testing.T) {
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	want := []byte{4, 7, 6} // type of class 128 (2·2), cpl, n
	want = append(want, "Bahnhof"...)
	want = append(want, 0, 6, 7, Further, Further, 3)
	want = append(want, "sallee"+"strasse"+"weg"...)
	for _, v := range []uint64{7, 1, 2, 5, 9, 4} {
		want = binaryLE(want, v)
	}
	if got := f.mem()[:f.Used()+6*8]; !bytes.Equal(got, want) || f.Size() != 128 || f.Used() != 32 {
		t.Fatalf("Fixed page:\n got %v\nwant %v (size %d, used %d)", got, want, f.Size(), f.Used())
	}
	if f.Len() != 6 || f.Keys() != 4 || f.PrefixLen() != 7 {
		t.Errorf("Len %d Keys %d PrefixLen %d", f.Len(), f.Keys(), f.PrefixLen())
	}

	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	want = []byte{2, 7, 6} // class 64
	want = append(want, "Bahnhof"...)
	want = append(want, 0, 6, 7, Further, Further, 3)
	want = append(want, 5) // "Bahnhof" itself: the empty remainder, then the length of its value
	want = append(want, "Mitte"...)
	want = append(want, "sallee"...)
	want = append(want, 4)
	want = append(want, "Nord"...)
	want = append(want, "strasse"...)
	want = append(want, 3)
	want = append(want, "Ost"...)
	want = append(want, 4)
	want = append(want, "Sued"...)
	want = append(want, 4)
	want = append(want, "West"...)
	want = append(want, "weg"...)
	want = append(want, 4)
	want = append(want, "Ring"...)
	if got := p.mem()[:p.Used()]; !bytes.Equal(got, want) || p.Size() != 64 || p.Used() != 62 {
		t.Fatalf("string page:\n got %v\nwant %v (size %d, used %d)", got, want, p.Size(), p.Used())
	}

	single := BuildFixed(bs("ab", "ac", "b"), []uint64{1, 2, 3})
	for _, rl := range single.mem()[Header+single.PrefixLen() : Header+single.PrefixLen()+3] {
		if rl == Further {
			t.Fatal("a page whose keys have one value has a Further")
		}
	}
}

func binaryLE(b []byte, v uint64) []byte {
	for i := range 8 {
		b = append(b, byte(v>>(8*i)))
	}
	return b
}

// TestMultiValuePageRemovals shows each way a value leaves a page, deterministically.
//
// A user who removes the values of a key one by one, in any order, must find the others where
// they were: the key's first value is the one with the remainder, so removing it hands the key
// to the next value.
//
// Expected: from a key with the values 2, 5, 9 the middle, the first and the last value go,
// the others stay in order each time, the key stays until the last value is gone, and the
// only value of a page makes the page gone.
func TestMultiValuePageRemovals(t *testing.T) {
	for _, pg := range []mvPage{&strPage{}, &fixedPage{}} {
		t.Run(fmt.Sprintf("%T", pg), func(t *testing.T) {
			m := mvModel{{"Bahnhof", []string{"7"}}, {"Bahnhofsallee", []string{"1"}}, {"Bahnhofstrasse", []string{"2", "5", "9"}}, {"Bahnhofweg", []string{"4"}}}
			for _, order := range [][]string{{"5", "2", "9"}, {"2", "5", "9"}, {"9", "5", "2"}} {
				cur := m.clone()
				p := pg.build(cur)
				for _, v := range order {
					if !p.remove("Bahnhofstrasse", v) {
						t.Fatalf("Remove(%s) failed", v)
					}
					cur, _ = cur.remove("Bahnhofstrasse", v)
					checkMV(t, p, cur)
				}
				if p.remove("Bahnhofstrasse", "9") || p.remove("Bahnhofstrasse", "5") {
					t.Error("removed a value that is gone")
				}
			}
			p := pg.build(mvModel{{"a", []string{"1", "2"}}})
			if !p.remove("a", "1") || !p.remove("a", "2") || !p.gone() {
				t.Error("the page is not gone with its last value")
			}
		})
	}
}

// TestMultiValuePageEachValue covers the lookups of the values of a key that find nothing and
// that the caller stops.
//
// The tree reads the values of a key (Each of the multimap) and may stop after some.
//
// Expected: a key outside the common prefix and a key the page does not have report false and call
// nothing; a stop after the first of three values reports true and sees one value.
func TestMultiValuePageEachValue(t *testing.T) {
	f := BuildFixed(exampleKeys, []uint64{7, 1, 2, 5, 9, 4})
	p := BuildStrings(exampleKeys, bs("Mitte", "Nord", "Ost", "Sued", "West", "Ring"))
	calls := 0
	if f.EachValue([]byte("Bahn"), func(uint64) bool { calls++; return true }) || f.EachValue([]byte("Bahnhofx"), func(uint64) bool { calls++; return true }) ||
		p.EachValue([]byte("Bahn"), func([]byte) bool { calls++; return true }) || p.EachValue([]byte("Bahnhofx"), func([]byte) bool { calls++; return true }) || calls != 0 {
		t.Fatalf("a key that is not there: %d calls", calls)
	}
	if !f.EachValue([]byte("Bahnhofstrasse"), func(uint64) bool { calls++; return false }) || !p.EachValue([]byte("Bahnhofstrasse"), func([]byte) bool { calls++; return false }) || calls != 2 {
		t.Fatalf("a stop after the first value: %d calls", calls)
	}
}

// TestMultiValuePageInsertKeepsOneValue shows that Insert, which the tree still uses until it
// holds several values in a page, leaves a key that has a value alone.
//
// Expected: Differs for another value of a key that is there, Present for the same, Added for a
// new key; the page is the same page after Differs.
func TestMultiValuePageInsertKeepsOneValue(t *testing.T) {
	p := BuildStrings(bs("a", "b"), bs("1", "2"))
	if q, res := p.Insert([]byte("a"), []byte("9")); res != Differs || q != p {
		t.Errorf("Differs: %v", res)
	}
	if _, res := p.Insert([]byte("a"), []byte("1")); res != Present {
		t.Errorf("Present: %v", res)
	}
	f := BuildFixed(bs("a", "b"), []uint64{1, 2})
	if q, res := f.Insert([]byte("a"), uint64(9)); res != Differs || q != f {
		t.Errorf("Fixed Differs: %v", res)
	}
	if _, res := f.Insert([]byte("c"), uint64(9)); res != Added {
		t.Errorf("Fixed Added: %v", res)
	}
}

// TestMultiValuePageBuildLimits shows what a page built from the values of keys refuses.
//
// The tree asks whether the values of a subtree fit one page before it makes one; a key
// with many values must count by its values.
//
// Expected: 256 values of one key and a value too long in a run give no page; 100 values of one
// key do (as strings and as one-byte values), and they are one key.
func TestMultiValuePageBuildLimits(t *testing.T) {
	many := make([][]byte, 256)
	one := make([]uint8, 256)
	vals := make([][]byte, 256)
	for i := range many {
		many[i], vals[i] = []byte("k"), []byte{byte(i)}
	}
	if BuildFixed(many, one) != nil || BuildStrings(many, vals) != nil {
		t.Error("256 values")
	}
	if p := BuildFixed(many[:100], one[:100]); p == nil || p.Keys() != 1 || p.Len() != 100 {
		t.Errorf("100 one-byte values of a key: %v", p)
	}
	if p := BuildStrings(many[:100], vals[:100]); p == nil || p.Keys() != 1 {
		t.Errorf("100 values of a key: %v", p)
	}
	if BuildStrings(bs("k", "k"), [][]byte{[]byte("a"), []byte(strings.Repeat("v", MaxValue+1))}) != nil {
		t.Error("a value that is too long in a run")
	}
	if BuildFixed(bs("a", "a"+strings.Repeat("x", 300), "a"+strings.Repeat("x", 300)), []uint8{1, 2, 3}) != nil {
		t.Error("a remainder that is too long before a run")
	}
}

// TestMultiValuePageSkipAndPrepend moves pages with several values up and down the tree.
//
// A byte node that comes in above a page, or goes away, takes bytes of the common prefix or gives
// them back; the values of a key must not notice.
//
// Expected: after Skip the keys are the old ones without their first bytes, with all their values
// in order; Prepend restores the old keys.
func TestMultiValuePageSkipAndPrepend(t *testing.T) {
	check := func(name string, pg mvPage, m mvModel, skip int) {
		t.Helper()
		var stripped mvModel
		for _, e := range m {
			stripped = append(stripped, mvEntry{e.key[skip:], e.vals})
		}
		switch q := pg.(type) {
		case *strPage:
			q.p.Skip(skip)
			checkMV(t, pg, stripped)
			q.p = q.p.Prepend([]byte(m[0].key[:skip]))
		case *fixedPage:
			q.p.Skip[uint64](skip)
			checkMV(t, pg, stripped)
			q.p = q.p.Prepend[uint64]([]byte(m[0].key[:skip]))
		}
		checkMV(t, pg, m)
	}
	m := mvModel{{"Bahnhof", []string{"7"}}, {"Bahnhofsallee", []string{"1"}}, {"Bahnhofstrasse", []string{"2", "5", "9"}}, {"Bahnhofweg", []string{"4"}}}
	for _, pg := range []mvPage{&strPage{}, &fixedPage{}} {
		check("short", pg.build(m.clone()), m, 3)
		check("all of it", pg.build(m.clone()), m, 7)
	}
}

// TestMultiValuePagePrefixOfNineBits makes sure that a common prefix of 256 bytes or more, which
// needs the ninth bit that the type byte lends, is kept through every change of the page.
//
// The head of the page starts like the head of a single-key page: the type, whose lowest bit is bit 8
// of the next byte, and the length of the key part stored. A tree whose types start somewhere
// else (TypeBase) must see the same.
//
// Expected: for a prefix of 300 bytes the type byte is odd and the next byte 44; the page keeps
// both when it grows into a larger class, when Skip takes bytes (the bit goes at 255 and below)
// and Prepend gives them back; Get finds the keys all the time.
func TestMultiValuePagePrefixOfNineBits(t *testing.T) {
	defer func(b uint8) { TypeBase = b }(TypeBase)
	for _, base := range []uint8{0, 20} {
		TypeBase = base
		long := strings.Repeat("p", 300)
		keys := bs(long+"a", long+"b", long+"b")
		f := BuildFixed(keys, []uint64{1, 2, 3})
		p := BuildStrings(keys, bs("1", "2", "3"))
		for _, h := range []*head{&f.head, &p.head} {
			if h.PrefixLen() != 300 || h.objType&1 != 1 || h.klen != 44 || h.Class() != 4 {
				t.Fatalf("base %d: prefix %d type %d klen %d class %d", base, h.PrefixLen(), h.objType, h.klen, h.Class())
			}
		}
		// grows into the next class with the bit
		q := p
		for i := range 20 {
			var res Result
			q, res = q.Add([]byte(long+"c"+string(rune('a'+i))), []byte(strings.Repeat("v", 20)))
			if res != Added && res != Full {
				t.Fatalf("Add: %v", res)
			}
		}
		if q.Class() != 5 || q.PrefixLen() != 300 {
			t.Errorf("after growing: class %d prefix %d", q.Class(), q.PrefixLen())
		}
		g := f
		for i := range 40 {
			g, _ = g.Add([]byte(long+"c"+string(rune('a'+i))), uint64(i))
		}
		if g.Class() != 5 || g.PrefixLen() != 300 {
			t.Errorf("fixed after growing: class %d prefix %d", g.Class(), g.PrefixLen())
		}
		if v, ok := g.Get[uint64]([]byte(long + "b")); !ok || v != 2 {
			t.Errorf("Get after growing: %v %v", v, ok)
		}
		// Skip below 256 clears the bit, Prepend sets it again
		g.Skip[uint64](100)
		q.Skip(100)
		if g.PrefixLen() != 200 || g.objType&1 != 0 || q.PrefixLen() != 200 || q.objType&1 != 0 {
			t.Errorf("after Skip: %d %d", g.PrefixLen(), q.PrefixLen())
		}
		g = g.Prepend[uint64]([]byte(strings.Repeat("p", 100)))
		q = q.Prepend([]byte(strings.Repeat("p", 100)))
		if g.PrefixLen() != 300 || g.objType&1 != 1 || q.PrefixLen() != 300 || q.objType&1 != 1 {
			t.Errorf("after Prepend: %d %d", g.PrefixLen(), q.PrefixLen())
		}
		if v, ok := q.Get([]byte(long + "b")); !ok || string(v) != "2" {
			t.Errorf("Get after Prepend: %q %v", v, ok)
		}
		// the largest prefix and the one beyond
		if BuildFixed(bs(strings.Repeat("p", MaxPrefix)+"a", strings.Repeat("p", MaxPrefix)+"b"), []uint8{1, 2}) != nil {
			t.Error("a prefix of 511 bytes and entries do not fit 512")
		}
	}
}

// TestMultiValuePageWidenRefusesLongRemainder shows that a page does not widen into a remainder
// that no length byte can hold.
//
// A key that leaves the common prefix early adds the rest of the old prefix to every remainder; a
// page of one key with a prefix of 300 bytes would get a remainder of 300 bytes.
//
// Expected: Widen answers nil for both flavors, and the page is unchanged.
func TestMultiValuePageWidenRefusesLongRemainder(t *testing.T) {
	long := strings.Repeat("p", 300)
	f := BuildFixed(bs(long), []uint64{1})
	p := BuildStrings(bs(long), bs("1"))
	if f.Widen([]byte("zzz"), uint64(2)) != nil || p.Widen([]byte("zzz"), []byte("2")) != nil {
		t.Fatal("widened into a remainder of 300 bytes")
	}
	if f.PrefixLen() != 300 || p.PrefixLen() != 300 {
		t.Error("the page changed")
	}
	if g := f.Widen([]byte(strings.Repeat("p", 100)+"zzz"), uint64(2)); g == nil || g.PrefixLen() != 100 || g.Keys() != 2 {
		t.Errorf("widened at 100: %v", g)
	}
	if g := f.Widen([]byte(strings.Repeat("z", 300)), uint64(2)); g != nil {
		t.Error("a remainder of 300 bytes for the new key")
	}
}

// FuzzMultiValuePage runs both flavors of the page against the model of several values a key,
// on operations the fuzzer chooses.
//
// It covers what random keys of a small alphabet may not: any byte value, any lengths, in any
// order, with the values of a key piling up and leaving.
//
// Expected: the page and the model agree after every operation.
func FuzzMultiValuePage(f *testing.F) {
	f.Add([]byte("abc\x00x1abc\x00y2\x01abc\x00x1abd\x00z2"))
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, fl := range []mvPage{&strPage{}, &fixedPage{}} {
			var pg mvPage
			var m mvModel
			d := data
			for len(d) >= 3 {
				op, kl, vl := d[0]%3, int(d[1])%9, int(d[2])%9
				d = d[3:]
				if len(d) < kl+vl {
					break
				}
				key, val := string(d[:kl]), string(d[kl:kl+vl])
				d = d[kl+vl:]
				if _, isFixed := fl.(*fixedPage); isFixed {
					val = strconv.Itoa(len(val) % 6)
				}
				switch {
				case pg == nil && op != 1:
					m = mvModel{{key, []string{val}}}
					if pg = fl.build(m); pg.gone() {
						pg = nil
						m = nil
					}
				case pg == nil:
				case op == 1:
					i, found := m.find(key)
					want := found && strings.HasPrefix(key, pg.cp()) && slices.Contains(m[i].vals, val)
					if ok := pg.remove(key, val); ok != want {
						t.Fatalf("Remove %q: %v, want %v", key, ok, want)
					} else if want {
						m, _ = m.remove(key, val)
					}
				default:
					want := expectAdd(pg, m, key, val)
					if res := pg.add(key, val); res != want {
						t.Fatalf("Add %q: %v, want %v", key, res, want)
					} else if res == Added || res == AddedValue {
						m, _ = m.add(key, val)
					}
				}
				if pg != nil && pg.gone() {
					pg, m = nil, nil
				}
				if pg != nil {
					checkMV(t, pg, m)
				}
			}
		}
	})
}
