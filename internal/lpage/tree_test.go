package lpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
	"unsafe"
)

// randomPage returns a page of random entries of one shape together with its
// sorted keys and values.
func randomPage(r *rand.Rand, shape string, n int) (*Page, [][]byte, [][]byte) {
	seen := map[string]bool{}
	var keys [][]byte
	for len(keys) < n {
		k := shapes[shape](r)
		if len(k) == 0 || seen[string(k)] {
			continue
		}
		seen[string(k)] = true
		keys = append(keys, k)
	}
	slices.SortFunc(keys, bytes.Compare)
	vals := make([][]byte, n)
	for i := range vals {
		vals[i] = rbytes(r, r.IntN(12), 256)
	}
	p, err := Build(keys, vals)
	if err != nil {
		return nil, nil, nil
	}
	return p, keys, vals
}

// TestTreeAccess checks the access by position that a tree uses.
//
// A tree that holds many keys in pages finds a key once and then reads its
// value, deletes it, cuts the page there or moves the page one level up or
// down. This belongs to the page layer of the tree and the test checks every
// one of these against the entries the page was built from: Find and Seek,
// ValueAt, ValueIs and EachValue, AppendKey and ByteAt, Shared, DeleteAt,
// SplitAt, Merge, Skip and Prepend.
func TestTreeAccess(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	for name := range shapes {
		for range 60 {
			p, keys, vals := randomPage(r, name, 1+r.IntN(9))
			if p == nil {
				continue
			}
			checkAccess(t, p, keys, vals, r)
		}
	}
}

func checkAccess(t *testing.T, p *Page, keys, vals [][]byte, r *rand.Rand) {
	t.Helper()
	n := len(keys)
	for i, k := range keys {
		if j, ok := p.Find(k); !ok || j != i {
			t.Fatalf("Find(%q) = %d, %v; want %d", k, j, ok, i)
		}
		if !bytes.Equal(p.ValueAt(i), vals[i]) || !p.ValueIs(i, string(vals[i])) {
			t.Fatalf("value %d is %q, want %q", i, p.ValueAt(i), vals[i])
		}
		if got := p.AppendKey([]byte("x"), i); !bytes.Equal(got, append([]byte("x"), k...)) {
			t.Fatalf("AppendKey(%d) = %q, want %q", i, got, k)
		}
		for off := range k {
			if p.ByteAt(i, off) != k[off] {
				t.Fatalf("ByteAt(%d, %d) wrong", i, off)
			}
		}
	}
	if p.ValueIs(0, string(vals[0])+"!") {
		t.Fatal("ValueIs accepts a longer value")
	}
	// Seek agrees with a binary search for random strings and for each key
	// with a byte appended or removed.
	probes := [][]byte{nil, {0}, {255, 255}}
	for _, k := range keys {
		probes = append(probes, append(bytes.Clone(k), 0), k[:len(k)-1], append(bytes.Clone(k[:len(k)-1]), 255))
		probes = append(probes, rbytes(r, 1+r.IntN(6), 4))
	}
	for _, s := range probes {
		want, found := slices.BinarySearchFunc(keys, s, bytes.Compare)
		if got, ok := p.Seek(s); got != want || ok != found {
			t.Fatalf("Seek(%q) = %d, %v; want %d, %v", s, got, ok, want, found)
		}
	}
	// EachValue over a run, and its early stop
	var got [][]byte
	i, j := n/3, n
	p.EachValue(i, j, func(v []byte) bool { got = append(got, bytes.Clone(v)); return true })
	if !slices.EqualFunc(got, vals[i:j], bytes.Equal) {
		t.Fatalf("EachValue(%d, %d) = %q, want %q", i, j, got, vals[i:j])
	}
	calls := 0
	if p.EachValue(0, n, func([]byte) bool { calls++; return false }) || calls != 1 {
		t.Fatal("EachValue does not stop")
	}
	want := lcp(keys[0], keys[n-1])
	if n == 1 {
		want = len(keys[0])
	}
	if p.Shared() != want {
		t.Fatalf("Shared = %d, want %d", p.Shared(), want)
	}
	if p.Thin() != (2*p.Used()*100 <= MergeFill*512) {
		t.Fatal("Thin disagrees with Used")
	}
	// SplitAt and Merge are inverse
	if n >= 2 {
		m := 1 + r.IntN(n-1)
		left, right := p.SplitAt(m)
		lk, lv := left.Entries()
		rk, rv := right.Entries()
		if !slices.EqualFunc(append(lk, rk...), keys, bytes.Equal) || !slices.EqualFunc(append(lv, rv...), vals, bytes.Equal) {
			t.Fatal("SplitAt loses entries")
		}
		if mg := Merge(left, right); mg != nil {
			mk, mv := mg.Entries()
			if !slices.EqualFunc(mk, keys, bytes.Equal) || !slices.EqualFunc(mv, vals, bytes.Equal) {
				t.Fatal("Merge loses entries")
			}
		}
	}
	// DeleteAt removes the entry and keeps the others
	k := r.IntN(n)
	cp, _ := Build(keys, vals) // DeleteAt works in place
	q := cp.DeleteAt(k, false)
	if n == 1 {
		if q != nil {
			t.Fatal("DeleteAt of the only entry leaves a page")
		}
		return
	}
	rest := slices.Delete(slices.Clone(keys), k, k+1)
	restV := slices.Delete(slices.Clone(vals), k, k+1)
	ks, vs := q.Entries()
	if !slices.EqualFunc(ks, rest, bytes.Equal) || !slices.EqualFunc(vs, restV, bytes.Equal) {
		t.Fatalf("DeleteAt(%d) left %q", k, ks)
	}
}

// TestTreeMoves checks Skip and Prepend, which move a page to another level.
//
// When a subtree gains or loses a node above a page, the page's remainders lose
// or gain the bytes of that node's path. Skip must drop exactly the shared bytes,
// Prepend must put them back, or refuse when the entries no longer fit; Merge
// must refuse two pages that are too full together.
func TestTreeMoves(t *testing.T) {
	keys := [][]byte{[]byte("abcd1"), []byte("abcd2"), []byte("abcde")}
	vals := [][]byte{[]byte("1"), []byte("22"), []byte("")}
	p, err := Build(keys, vals)
	if err != nil {
		t.Fatal(err)
	}
	if s := p.Skip(0); s != p {
		t.Error("Skip(0) copies the page")
	}
	s := p.Skip(4)
	sk, sv := s.Entries()
	if !slices.EqualFunc(sk, [][]byte{[]byte("1"), []byte("2"), []byte("e")}, bytes.Equal) || !slices.EqualFunc(sv, vals, bytes.Equal) {
		t.Fatalf("Skip(4) = %q", sk)
	}
	b := s.Prepend([]byte("abcd"))
	bk, _ := b.Entries()
	if !slices.EqualFunc(bk, keys, bytes.Equal) {
		t.Fatalf("Prepend = %q", bk)
	}
	if s.Prepend(make([]byte, 600)) != nil {
		t.Error("Prepend accepts a prefix that is too long")
	}
	fat := [][]byte{bytes.Repeat([]byte{'x'}, 200), bytes.Repeat([]byte{'y'}, 200)}
	a, _ := Build(fat[:1], [][]byte{bytes.Repeat([]byte{'v'}, 200)})
	c, _ := Build(fat[1:], [][]byte{bytes.Repeat([]byte{'v'}, 200)})
	if Merge(a, c) != nil {
		t.Error("Merge builds a page of more than MergeFill percent")
	}
	if !Fits([]byte("a"), nil) || Fits(nil, nil) || Fits(make([]byte, 300), nil) || !FitsLen(1, 255) || FitsLen(0, 1) || FitsLen(300, 1) || FitsLen(255, 255) {
		t.Error("Fits and FitsLen decide wrongly")
	}
}

// TestTryInsert checks the insertion of a tree that keeps one value per key.
//
// A tree puts a key with its first value into a page and moves it elsewhere when
// it gets a second one, so the page must not replace a value that is there: it
// reports the key as present, with its position, and leaves the page alone, for
// every form of page (values of one width or of different lengths, a new value
// of another length), and it grows the page in place, with a longer header or in a
// bigger class, or says Full.
func TestTryInsert(t *testing.T) {
	var model []struct{ k, v []byte }
	var p *Page
	r := rand.New(rand.NewPCG(9, 9))
	full := false
	for round := 0; round < 400 && !full; round++ {
		k := []byte(fmt.Sprintf("key-%03d", r.IntN(60)))
		v := rbytes(r, r.IntN(5), 256)
		if round%7 == 0 {
			v = []byte("12345678") // a run of equal widths now and then
		}
		if p == nil {
			var err error
			if p, err = Build([][]byte{k}, [][]byte{v}); err != nil {
				t.Fatal(err)
			}
			model = append(model, struct{ k, v []byte }{k, v})
			continue
		}
		before, _ := p.Entries()
		q, res, at, err := p.TryInsert(k, v, false)
		if err != nil {
			t.Fatal(err)
		}
		switch res {
		case Present:
			if q != p || !bytes.Equal(before[at], k) {
				t.Fatalf("Present at %d for %q", at, k)
			}
		case Inserted:
			p = q
			model = append(model, struct{ k, v []byte }{k, v})
		case Full:
			full = true
			if q != p {
				t.Fatal("Full returns another page")
			}
		default:
			t.Fatalf("TryInsert result %v", res)
		}
	}
	slices.SortFunc(model, func(a, b struct{ k, v []byte }) int { return bytes.Compare(a.k, b.k) })
	keys, vals := p.Entries()
	for i, e := range model {
		if !bytes.Equal(keys[i], e.k) || !bytes.Equal(vals[i], e.v) {
			t.Fatalf("entry %d is %q=%q, want %q=%q", i, keys[i], vals[i], e.k, e.v)
		}
	}
	if !full {
		t.Fatal("the page never got full")
	}
	if _, _, _, err := p.TryInsert(nil, nil, false); err != ErrEmpty {
		t.Errorf("empty suffix: %v", err)
	}
	if _, _, _, err := p.TryInsert(make([]byte, 300), nil, false); err != ErrTooLong {
		t.Errorf("long suffix: %v", err)
	}
	// a page of values of one width, and a key that is there with a value of
	// another length: present, not replaced
	w, _ := Build([][]byte{[]byte("a"), []byte("b")}, [][]byte{[]byte("1234"), []byte("5678")})
	if q, res, at, _ := w.TryInsert([]byte("b"), []byte("x"), false); res != Present || at != 1 || q != w || !w.ValueIs(1, "5678") {
		t.Errorf("another length replaces a value: %v at %d", res, at)
	}
	if q, res, _, _ := w.TryInsert([]byte("zzzz"), []byte("x"), false); res != Inserted || q.Len() != 3 {
		t.Errorf("new key with another length: %v", res)
	}
	if MaxEntries() != min(MaxHeader-2, maxEnts-1) {
		t.Errorf("MaxEntries = %d with header %d", MaxEntries(), MaxHeader)
	}
}

// TestEachString checks the strings a scan hands out.
//
// A scan over many entries of a page copies their values in one piece; each
// string must still be its own value, and stopping early must work.
func TestEachString(t *testing.T) {
	keys := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}
	vals := [][]byte{[]byte("one"), nil, []byte("three"), []byte("4")}
	p, err := Build(keys, vals)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	if !p.EachString(1, 4, false, func(s string) bool { got = append(got, s); return true }) || !slices.Equal(got, []string{"", "three", "4"}) {
		t.Fatalf("EachString(1, 4) = %q", got)
	}
	n := 0
	if p.EachString(0, 4, false, func(string) bool { n++; return n < 2 }) {
		t.Error("EachString does not stop")
	}
	if !p.EachString(2, 2, false, func(string) bool { t.Fatal("called for an empty run"); return true }) {
		t.Error("an empty run is not complete")
	}
}

// TestMergeRefuses checks the quick refusals of Merge.
//
// A tree tries to merge a thin page with its neighbours after almost every delete,
// so most tries must fail without laying anything out: too many bytes, too many
// entries for any header.
func TestMergeRefuses(t *testing.T) {
	defer func(old int) { MaxHeader = old }(MaxHeader)
	MaxHeader = 24
	var ka, kb [][]byte
	var va, vb [][]byte
	for i := range 7 {
		ka = append(ka, []byte{'a', byte('a' + i)})
		kb = append(kb, []byte{'b', byte('a' + i)})
		va = append(va, rbytes(rand.New(rand.NewPCG(1, 1)), i, 256))
		vb = append(vb, nil)
	}
	a, err := Build(ka, va)
	if err != nil {
		t.Fatal(err)
	}
	b, err := Build(kb, vb)
	if err != nil {
		t.Fatal(err)
	}
	if Merge(a, b) != nil {
		t.Error("14 entries merge into a page of variable values")
	}
	one, _ := Build([][]byte{[]byte("c")}, [][]byte{[]byte("1")})
	if m := Merge(a, one); m == nil || m.Len() != 8 {
		t.Error("a page and one more entry do not merge")
	}
}

// TestImmutablePages checks the mode of a tree whose pages are never changed.
//
// A tree that hands out strings which alias its pages must never modify a page
// after building it, or a string a user holds would change. In this mode an
// insertion or deletion leaves the old page exactly as it was, produces the
// same entries as the in-place one, and the aliased strings are the page's own
// bytes.
func TestImmutablePages(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	for range 300 {
		p, keys, vals := randomPage(r, "path", 1+r.IntN(8))
		if p == nil {
			continue
		}
		snapshot := bytes.Clone(p.mem())
		k := []byte(fmt.Sprintf("name %d", r.IntN(1000)))
		v := rbytes(r, r.IntN(9), 256)
		q, res, _, err := p.TryInsert(k, v, true)
		if err != nil || !bytes.Equal(snapshot, p.mem()) {
			t.Fatalf("TryInsert with cow changes the page or fails: %v", err)
		}
		if res == Inserted {
			if q == p {
				t.Fatal("TryInsert with cow returns the same page")
			}
			ks, vs := q.Entries()
			if len(ks) != len(keys)+1 || !has2(ks, k) || len(vs) != len(vals)+1 {
				t.Fatalf("TryInsert with cow lost entries: %q", ks)
			}
		}
		if p.Len() > 1 {
			i := r.IntN(p.Len())
			d := p.DeleteAt(i, true)
			if !bytes.Equal(snapshot, p.mem()) || d == p || d.Len() != p.Len()-1 {
				t.Fatal("DeleteAt with cow changes the page")
			}
			rest := slices.Delete(slices.Clone(keys), i, i+1)
			if dk, _ := d.Entries(); !slices.EqualFunc(dk, rest, bytes.Equal) {
				t.Fatalf("DeleteAt(%d) with cow left %q", i, dk)
			}
		}
		if len(vals[0]) > 0 {
			s := p.StringAt(0, true)
			if unsafe.StringData(s) != &p.ValueAt(0)[0] {
				t.Fatal("StringAt with alias copies")
			}
			p.EachString(0, p.Len(), true, func(s string) bool { return true })
		}
		if s := p.StringAt(0, false); s != string(vals[0]) {
			t.Fatalf("StringAt = %q, want %q", s, vals[0])
		}
	}
}

func has2(keys [][]byte, k []byte) bool {
	return slices.ContainsFunc(keys, func(x []byte) bool { return bytes.Equal(x, k) })
}

// pairModel is the entries of a page as the page must hold them: keys in order,
// each with its values in the order they were added.
type pairModel struct {
	keys []string
	vals map[string][]string
}

func (m *pairModel) entries() (keys, vals [][]byte) {
	for _, k := range m.keys {
		for _, v := range m.vals[k] {
			keys, vals = append(keys, []byte(k)), append(vals, []byte(v))
		}
	}
	return keys, vals
}

// TestMultiValueEntries checks pages that hold keys with several values.
//
// A tree for strings keeps every key in a page, also one that has several values;
// the page stores its key once and a value for each. Adding values to a key,
// removing them one by one (also the first one, which carries the key's bytes)
// and reading the entries back must agree with a plain model, in place and
// with immutable pages, for values of any length and of one width.
func TestMultiValueEntries(t *testing.T) {
	for _, cow := range []bool{false, true} {
		for _, vgen := range []string{"short", "width"} {
			t.Run(fmt.Sprintf("cow %v/%s", cow, vgen), func(t *testing.T) {
				r := rand.New(rand.NewPCG(7, 8))
				for round := range 200 {
					m := &pairModel{vals: map[string][]string{}}
					var p *Page
					for step := 0; step < 60; step++ {
						k := fmt.Sprintf("key%02d", r.IntN(9))
						v := strings.Repeat("v", r.IntN(4)) + fmt.Sprint(r.IntN(6))
						if vgen == "width" {
							v = fmt.Sprintf("%04d", r.IntN(6))
							if r.IntN(12) == 0 {
								v = "x" // another length in a page of one width
							}
						}
						if r.IntN(4) == 0 && p != nil {
							// remove a random value of a random key
							var pairs [][2]string
							for key, vs := range m.vals {
								for _, x := range vs {
									pairs = append(pairs, [2]string{key, x})
								}
							}
							if len(pairs) == 0 {
								continue
							}
							slices.SortFunc(pairs, func(a, b [2]string) int { return strings.Compare(a[0]+a[1], b[0]+b[1]) })
							pr := pairs[r.IntN(len(pairs))]
							keys, vals := p.Entries()
							at := -1
							for i := range keys {
								if string(keys[i]) == pr[0] && string(vals[i]) == pr[1] {
									at = i
								}
							}
							if at < 0 {
								t.Fatalf("round %d: pair %v is not in the page", round, pr)
							}
							p = p.DeleteAt(at, cow)
							m.vals[pr[0]] = slices.DeleteFunc(m.vals[pr[0]], func(x string) bool { return x == pr[1] })
							if len(m.vals[pr[0]]) == 0 {
								delete(m.vals, pr[0])
								m.keys = slices.DeleteFunc(m.keys, func(x string) bool { return x == pr[0] })
							}
						} else if p == nil {
							var err error
							if p, err = Build([][]byte{[]byte(k)}, [][]byte{[]byte(v)}); err != nil {
								t.Fatal(err)
							}
							m.keys, m.vals[k] = []string{k}, []string{v}
						} else if i, ok := p.Find([]byte(k)); ok {
							q, res := p.AddValueAt(i, []byte(v), cow)
							switch res {
							case Present:
								if !slices.Contains(m.vals[k], v) {
									t.Fatalf("Present for the new value %q of %q", v, k)
								}
							case Added:
								p = q
								m.vals[k] = append(m.vals[k], v)
							case Full:
							default:
								t.Fatalf("AddValueAt result %v", res)
							}
						} else {
							q, res, _, err := p.TryInsert([]byte(k), []byte(v), cow)
							if err != nil {
								t.Fatal(err)
							}
							if res == Inserted {
								p = q
								m.keys = append(m.keys, k)
								slices.Sort(m.keys)
								m.vals[k] = []string{v}
							}
						}
						if p == nil {
							m = &pairModel{vals: map[string][]string{}}
							continue
						}
						checkPairs(t, p, m)
					}
				}
			})
		}
	}
}

func checkPairs(t *testing.T, p *Page, m *pairModel) {
	t.Helper()
	wk, wv := m.entries()
	gk, gv := p.Entries()
	if !slices.EqualFunc(gk, wk, bytes.Equal) || !slices.EqualFunc(gv, wv, bytes.Equal) {
		t.Fatalf("page holds %q=%q, model %q=%q", gk, gv, wk, wv)
	}
	if p.Keys() != len(m.keys) || p.Len() != len(wk) {
		t.Fatalf("Keys %d, Len %d; want %d, %d", p.Keys(), p.Len(), len(m.keys), len(wk))
	}
	i := 0
	for _, k := range m.keys {
		if got, ok := p.Find([]byte(k)); !ok || got != i || p.IsCont(i) {
			t.Fatalf("Find(%q) = %d, %v; want %d", k, got, ok, i)
		}
		if p.RunEnd(i) != i+len(m.vals[k]) {
			t.Fatalf("RunEnd(%d) = %d, want %d", i, p.RunEnd(i), i+len(m.vals[k]))
		}
		for j, v := range m.vals[k] {
			if string(p.ValueAt(i+j)) != v || string(p.AppendKey(nil, i+j)) != k || (j > 0) != p.IsCont(i+j) {
				t.Fatalf("entry %d of %q is wrong", i+j, k)
			}
			if p.ByteAt(i+j, 2) != k[2] {
				t.Fatalf("ByteAt(%d, 2) of %q", i+j, k)
			}
		}
		if got, ok := p.Seek([]byte(k)); !ok || got != i {
			t.Fatalf("Seek(%q) = %d, %v", k, got, ok)
		}
		i += len(m.vals[k])
	}
	if p.Shared() > len(m.keys[0]) {
		t.Fatalf("Shared = %d", p.Shared())
	}
	if v, ok := p.Get([]byte(m.keys[0])); !ok || string(v) != m.vals[m.keys[0]][0] {
		t.Fatalf("Get(first key) = %q, %v", v, ok)
	}
}
