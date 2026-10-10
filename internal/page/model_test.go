package page

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

// mvEntry is a key of the reference model of the pages: the key from the end of the path on, and
// its values in the order they came in.
type mvEntry struct {
	key  string
	vals []string
}

type mvModel []mvEntry

func (m mvModel) find(key string) (int, bool) {
	i := sort.Search(len(m), func(i int) bool { return m[i].key >= key })
	return i, i < len(m) && m[i].key == key
}

func (m mvModel) nValues() (n int) {
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

// removalOf is what Remove of val from the key at m[i] must answer; inside says that the key is
// found and starts with the page's key part (many-key form) or is the key (one-key form).
func removalOf(m mvModel, i int, inside bool, val string) Removal {
	switch {
	case !inside || !slices.Contains(m[i].vals, val):
		return Absent
	case len(m[i].vals) == 1:
		return Gone
	}
	return Removed
}

// withPrefix returns the model with pre in front of every key.
func (m mvModel) withPrefix(pre string) mvModel {
	out := m.clone()
	for i := range out {
		out[i].key = pre + out[i].key
	}
	return out
}

// withoutPrefix returns the model with the first k bytes of every key cut off.
func (m mvModel) withoutPrefix(k int) mvModel {
	out := m.clone()
	for i := range out {
		out[i].key = out[i].key[k:]
	}
	return out
}

// mvFlat is one value of a page as Each reports it.
type mvFlat struct {
	key, val string
	first    bool
}

// mvPage is a page of either flavor as the model tests see it.
type mvPage interface {
	add(key, val string) Result
	remove(key, val string) Removal
	widen(key, val string) bool
	skip(k int)
	prepend(pre string) bool
	gone() bool
	cp() string
	oneKey() bool
	need(nValues, remBytes, valBytes int) int
	tooLong(key, val string) bool
	flat() []mvFlat
	values(key string) ([]string, bool)
	get(key string) (string, bool)
	counts() (nValues, keys int)
	invariants(t *testing.T)
	build(m mvModel) mvPage
	value(r *rand.Rand) string
}

func randKey(r *rand.Rand, prefix string) string {
	var b strings.Builder
	b.WriteString(prefix)
	for range r.IntN(7) {
		b.WriteByte("ab\x00\xff"[r.IntN(4)])
	}
	return b.String()
}

func randVal(r *rand.Rand) string {
	return strings.Repeat("v", r.IntN(12)) + string(rune('0'+r.IntN(3)))
}

func bs(ss ...string) [][]byte {
	out := make([][]byte, len(ss))
	for i, s := range ss {
		out[i] = []byte(s)
	}
	return out
}

type strPage struct{ p *Str }

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
func (s *strPage) remove(key, val string) Removal {
	q, rm := s.p.Remove([]byte(key), []byte(val))
	s.p = q
	return rm
}
func (s *strPage) widen(key, val string) bool {
	q := s.p.Widen([]byte(key), []byte(val))
	if q != nil {
		s.p = q
	}
	return q != nil
}
func (s *strPage) skip(k int) { s.p.Skip(k) }
func (s *strPage) prepend(pre string) bool {
	q := s.p.Prepend([]byte(pre))
	if q != nil {
		s.p = q
	}
	return q != nil
}
func (s *strPage) gone() bool   { return s.p == nil }
func (s *strPage) cp() string   { return string(s.p.CP()) }
func (s *strPage) oneKey() bool { return s.p.OneKey() }
func (s *strPage) need(nValues, remBytes, valBytes int) int {
	if s.p.OneKey() {
		return Header + s.p.PrefixLen() + nValues + valBytes
	}
	return Header + s.p.PrefixLen() + 2*nValues + remBytes + valBytes
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
	m := s.p.mem()
	la := s.p.lay(true)
	used := la.keyEnd(m)
	vb := sum(m[la.vl : la.vl+la.currentValues])
	if used+vb > s.p.Size() {
		t.Fatalf("used %d and values %d exceed the size %d", used, vb, s.p.Size())
	}
	if bytes.Count(m[used:len(m)-vb], []byte{0}) != len(m)-vb-used {
		t.Fatal("bytes between the keys and the values are not zero")
	}
}

// checkMV compares a page with the model: every value of every key in order, the first flags, the
// counts, the lookups, the form, and the layout invariants.
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
	if nValues, keys := pg.counts(); nValues != m.nValues() || keys != len(m) {
		t.Fatalf("Len %d Keys %d, model %d and %d", nValues, keys, m.nValues(), len(m))
	}
	if pg.oneKey() != (len(m) == 1) {
		t.Fatalf("a page of %d keys has OneKey %v", len(m), pg.oneKey())
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
	i, exists := m.find(key)
	if exists && slices.Contains(m[i].vals, val) {
		return Present
	}
	if pg.oneKey() && !exists { // another key: the two make a page of the many-key form, if they fit
		if all, _ := m.clone().add(key, val); pg.build(all).gone() {
			return Full
		}
		return Added
	}
	cp := pg.cp()
	if !pg.oneKey() && !strings.HasPrefix(key, cp) {
		return Outside
	}
	rems, vals := 0, len(val)
	for _, e := range m {
		if !pg.oneKey() {
			rems += len(e.key) - len(cp)
		}
		for _, v := range e.vals {
			vals += len(v)
		}
	}
	if !exists {
		rems += len(key) - len(cp)
	}
	if pg.tooLong(key, val) || (!exists && len(key)-len(cp) > MaxRemainder) || pg.need(m.nValues()+1, rems, vals) > sizes[Classes-1] {
		return Full
	}
	if exists {
		return AddedValue
	}
	return Added
}

// TestPageAgainstModel keeps a page and a plain model side by side through thousands of random
// adds, removes and widenings from many starting pages, for every flavor, and compares them after
// every step.
//
// For a user of the multimap this is the guarantee of the page: a key of the index that gets a second,
// third, tenth value keeps every value, in the order they came in, in the page of its neighbours, and a
// value that goes away leaves the others where they are, whichever of the values it is; a page of
// one key is the one-key form and a page of several keys the many-key form, and the keys of a page are
// found in either.
//
// Expected: after every operation the result is the one the model predicts, the entries, the first
// flags, the counts and the lookups equal the model's, the form is that of the number of keys, the
// layout invariants hold, and a refused Widen is one that a page built from all the entries would also refuse.
func TestPageAgainstModel(t *testing.T) {
	for _, pg := range []mvPage{&strPage{}, &fixedPage{}, &ptrPage{}} {
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
			opName := "add"
			switch op := r.IntN(22); {
			case op == 20 && len(cp) > 0 && !pg.oneKey(): // the path gains bytes: a byte node above the page
				opName = "skip"
				k := 1 + r.IntN(len(cp))
				pg.skip(k)
				m = m.withoutPrefix(k)
			case op == 20 && pg.oneKey() && len(m[0].key) > 0:
				opName = "skip"
				k := 1 + r.IntN(len(m[0].key))
				pg.skip(k)
				m = m.withoutPrefix(k)
			case op == 21: // the path loses bytes: a byte node above the page goes away
				opName = "prepend"
				pre := strings.Repeat("a", 1+r.IntN(40))
				all := m.withPrefix(pre)
				fits := !pg.build(all).gone()
				if done := pg.prepend(pre); done != fits {
					t.Fatalf("seed %d step %d: Prepend(%q) = %v, a page built from all the entries: %v", seed, step, pre, done, fits)
				} else if done {
					m = all
				}
			case op < 8: // remove
				opName = "remove"
				i, found := m.find(key)
				inside := found
				if !pg.oneKey() {
					inside = found && strings.HasPrefix(key, cp)
				}
				want := removalOf(m, i, inside, val)
				if got := pg.remove(key, val); got != want {
					t.Fatalf("seed %d step %d: Remove(%q, %q) = %v, want %v", seed, step, key, val, got, want)
				}
				if want != Absent {
					m, _ = m.remove(key, val)
				}
			default:
				want := expectAdd(pg, m, key, val)
				if want == Outside && r.IntN(2) == 0 { // the tree would widen the page, if the entries fit
					opName = "widen"
					all, _ := m.clone().add(key, val)
					fits := !pg.build(all).gone()
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
			t.Logf("seed %d step %d: %s %q %q", seed, step, opName, key, val)
			checkMV(t, pg, m)
		}
	}
}

type fixedPage struct{ p *Fixed }

func fixedVal(s string) uint64 { n, _ := strconv.ParseUint(s, 10, 64); return n }

func (f *fixedPage) build(m mvModel) mvPage {
	rests, vals := m.rests()
	vs := make([]uint64, len(vals))
	for i, v := range vals {
		vs[i] = fixedVal(string(v))
	}
	return &fixedPage{BuildFixed(rests, vs)}
}
func (f *fixedPage) value(r *rand.Rand) string { return strconv.Itoa(r.IntN(6)) }
func (f *fixedPage) add(key, val string) Result {
	q, res := f.p.Add([]byte(key), fixedVal(val), false)
	f.p = q
	return res
}
func (f *fixedPage) remove(key, val string) Removal {
	q, rm := f.p.Remove([]byte(key), fixedVal(val), false)
	f.p = q
	return rm
}
func (f *fixedPage) widen(key, val string) bool {
	q := f.p.Widen([]byte(key), fixedVal(val), false)
	if q != nil {
		f.p = q
	}
	return q != nil
}
func (f *fixedPage) skip(k int) { f.p.Skip(k) }
func (f *fixedPage) prepend(pre string) bool {
	q := f.p.Prepend[uint64]([]byte(pre), false)
	if q != nil {
		f.p = q
	}
	return q != nil
}
func (f *fixedPage) gone() bool   { return f.p == nil }
func (f *fixedPage) cp() string   { return string(f.p.CP()) }
func (f *fixedPage) oneKey() bool { return f.p.OneKey() }
func (f *fixedPage) need(nValues, remBytes, _ int) int {
	if f.p.OneKey() {
		return Header + f.p.PrefixLen() + 8*nValues
	}
	return Header + f.p.PrefixLen() + nValues + remBytes + 8*nValues
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
	m := f.p.mem()
	used := f.p.Used()
	vs := len(m) - f.p.Len()*int(unsafe.Sizeof(uint64(0)))
	if vs < used {
		t.Fatalf("values start at %d, the keys end at %d", vs, used)
	}
	if bytes.Count(m[used:vs], []byte{0}) != vs-used {
		t.Fatal("bytes between the keys and the values are not zero")
	}
}

// ptrPool holds the values of the pointer pages in the tests: the pointer of "n" is &ptrPool[n], and holds n.
var ptrPool = func() (p [64]uint64) {
	for i := range p {
		p[i] = uint64(i)
	}
	return p
}()

func ptrVal(s string) *uint64 { n, _ := strconv.ParseUint(s, 10, 64); return &ptrPool[n] }

type ptrPage struct{ p *Fixed }

func (f *ptrPage) build(m mvModel) mvPage {
	rests, vals := m.rests()
	vs := make([]*uint64, len(vals))
	for i, v := range vals {
		vs[i] = ptrVal(string(v))
	}
	return &ptrPage{BuildFixed(rests, vs)}
}
func (f *ptrPage) value(r *rand.Rand) string { return strconv.Itoa(r.IntN(6)) }
func (f *ptrPage) add(key, val string) Result {
	q, res := f.p.Add([]byte(key), ptrVal(val), true)
	f.p = q
	return res
}
func (f *ptrPage) remove(key, val string) Removal {
	q, rm := f.p.Remove([]byte(key), ptrVal(val), true)
	f.p = q
	return rm
}
func (f *ptrPage) widen(key, val string) bool {
	q := f.p.Widen([]byte(key), ptrVal(val), true)
	if q != nil {
		f.p = q
	}
	return q != nil
}
func (f *ptrPage) skip(k int) { f.p.Skip(k) }
func (f *ptrPage) prepend(pre string) bool {
	q := f.p.Prepend[*uint64]([]byte(pre), true)
	if q != nil {
		f.p = q
	}
	return q != nil
}
func (f *ptrPage) gone() bool   { return f.p == nil }
func (f *ptrPage) cp() string   { return string(f.p.CP()) }
func (f *ptrPage) oneKey() bool { return f.p.OneKey() }
func (f *ptrPage) need(nValues, remBytes, _ int) int {
	if f.p.OneKey() {
		return Header + f.p.PrefixLen() + 8*nValues
	}
	return Header + f.p.PrefixLen() + nValues + remBytes + 8*nValues
}
func (f *ptrPage) tooLong(string, string) bool { return false }
func (f *ptrPage) flat() (out []mvFlat) {
	cp := string(f.p.CP())
	f.p.Each(func(rem []byte, v *uint64, first bool) bool {
		out = append(out, mvFlat{cp + string(rem), strconv.FormatUint(*v, 10), first})
		return true
	})
	return out
}
func (f *ptrPage) values(key string) (vs []string, ok bool) {
	ok = f.p.EachValue([]byte(key), func(v *uint64) bool { vs = append(vs, strconv.FormatUint(*v, 10)); return true })
	return vs, ok
}
func (f *ptrPage) get(key string) (string, bool) {
	v, ok := f.p.Get[*uint64]([]byte(key))
	return strconv.FormatUint(*v, 10), ok
}
func (f *ptrPage) counts() (int, int) { return f.p.Len(), f.p.Keys() }
func (f *ptrPage) invariants(t *testing.T) {
	t.Helper()
	m := f.p.mem()
	used := f.p.Used()
	vs := len(m) - f.p.Len()*8
	if vs < used {
		t.Fatalf("values start at %d, the keys end at %d", vs, used)
	}
	if used > f.p.RawWords()*8 {
		t.Fatalf("the keys end at %d, the byte area has %d bytes", used, f.p.RawWords()*8)
	}
	if bytes.Count(m[used:vs], []byte{0}) != vs-used {
		t.Fatal("bytes between the keys and the values are not zero")
	}
}

// FuzzPage lets the fuzzer drive the pages of every flavor against the model with keys and values of any
// bytes. It covers what random keys of a small alphabet may not: any byte value, any lengths, in any
// order, with the values of a key piling up and leaving, and the page changing between its two forms.
//
// Expected: the page and the model agree after every operation.
func FuzzPage(f *testing.F) {
	f.Add([]byte("abc\x00x1abc\x00y2\x01abc\x00x1abd\x00z2"))
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, fl := range []mvPage{&strPage{}, &fixedPage{}, &ptrPage{}} {
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
				if _, isStr := fl.(*strPage); !isStr {
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
					inside := found
					if !pg.oneKey() {
						inside = found && strings.HasPrefix(key, pg.cp())
					}
					want := removalOf(m, i, inside, val)
					if got := pg.remove(key, val); got != want {
						t.Fatalf("Remove %q: %v, want %v", key, got, want)
					} else if want != Absent {
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
