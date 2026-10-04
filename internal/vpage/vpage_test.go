package vpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
	"unsafe"
)

// shapes generate the remainders the pages are tested with: integers (uniform
// pages), short remainders of several lengths whose zero padding makes their heads
// equal, remainders that share their first 8 bytes and so need their tails
// compared, and strings of every length up to the longest a page holds.
var shapes = map[string]func(r *rand.Rand) []byte{
	"u64":         func(r *rand.Rand) []byte { return rbytes(r, 8, 256) },
	"u32":         func(r *rand.Rand) []byte { return rbytes(r, 4, 256) },
	"short-mixed": func(r *rand.Rand) []byte { return rbytes(r, r.IntN(9), 3) },
	"ties": func(r *rand.Rand) []byte {
		return append([]byte("tie-pref"), rbytes(r, r.IntN(24), 3)...)
	},
	"uuid": func(r *rand.Rand) []byte {
		k := rbytes(r, 36, 16)
		for i := range k {
			k[i] = "0123456789abcdef"[k[i]]
		}
		return k
	},
	"path": func(r *rand.Rand) []byte {
		dirs := []string{"usr", "lib", "share", "doc", "python3", "dist-packages", "x86_64-linux-gnu"}
		var b []byte
		for range 1 + r.IntN(7) {
			b = append(b, '/')
			b = append(b, dirs[r.IntN(len(dirs))]...)
		}
		return append(b, rbytes(r, r.IntN(12), 26)...)
	},
	"deep": func(r *rand.Rand) []byte {
		return append([]byte("/a/very/deep/directory/of/some/repository/"), rbytes(r, 1+r.IntN(6), 3)...)
	},
	"twelve": func(r *rand.Rand) []byte { // 12 bytes that share 6: a page of them is uniform, with a prefix
		return append([]byte("shared"), rbytes(r, 6, 256)...)
	},
	"long":  func(r *rand.Rand) []byte { return rbytes(r, 200+r.IntN(56), 2) },
	"zeros": func(r *rand.Rand) []byte { return make([]byte, r.IntN(12)) },
	"any":   func(r *rand.Rand) []byte { return rbytes(r, r.IntN(maxRemainder+1), 256) },
}

// rbytes returns n random bytes below alphabet.
func rbytes(r *rand.Rand, n, alphabet int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(r.IntN(alphabet))
	}
	return b
}

// check fails the test if page p breaks an invariant of its layout: the keys
// in order, the head words, tags and lengths matching the keys without the
// prefix, the tails in the heap and apart from each other, the unused heads
// padded.
func check(t *testing.T, p *Page) {
	t.Helper()
	n, c, end := int(p.count), int(p.cap), p.Size()
	if n == 0 || n > c || arraysEnd(c, p.ulen != 0, int(p.plen)) > int(p.top) || int(p.top) > end {
		t.Fatalf("count %d, cap %d, arrays end %d, heap %d of %d", n, c, arraysEnd(c, p.ulen != 0, int(p.plen)), p.top, end)
	}
	for i := n; i < c; i++ {
		if p.slots()[i].Head != pad {
			t.Fatalf("head %d past the count is not padded", i)
		}
	}
	var prev, buf [maxRemainder]byte
	var prevKey []byte
	type span struct{ from, to int }
	var spans []span
	for i := range n {
		full := p.Key(i, &buf)
		if !bytes.HasPrefix(full, p.prefix()) {
			t.Fatalf("key %d does not start with the prefix", i)
		}
		k := full[p.plen:]
		if p.ulen != 0 && len(k) != int(p.ulen) {
			t.Fatalf("key %d has %d bytes in a uniform page of %d", i, len(k), p.ulen)
		}
		got := p.tag(i)
		if p.slots()[i].Head != word(k) || got != tag(word(k)) {
			t.Fatalf("key %d: head %x, word %x, tag %x, want %x", i, p.slots()[i].Head, word(k), got, tag(word(k)))
		}
		if i > 0 && bytes.Compare(prevKey, full) >= 0 {
			t.Fatalf("keys %d and %d are out of order: %x %x", i-1, i, prevKey, full)
		}
		prevKey = append(prev[:0], full...)
		if p.ulen == 0 && len(k) > headLen {
			off := int(p.fat()[i] >> 16)
			if off < int(p.top) || off+len(k)-headLen > end {
				t.Fatalf("tail %d at %d..%d is outside the heap %d..%d", i, off, off+len(k)-headLen, p.top, end)
			}
			spans = append(spans, span{off, off + len(k) - headLen})
		}
	}
	slices.SortFunc(spans, func(a, b span) int { return a.from - b.from })
	for i := 1; i < len(spans); i++ {
		if spans[i].from < spans[i-1].to {
			t.Fatalf("tails overlap: %v %v", spans[i-1], spans[i])
		}
	}
}

// reference is the sorted map that pages are compared with.
type reference map[string]uint64

func (r reference) sorted() []string {
	ks := make([]string, 0, len(r))
	for k := range r {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

// TestRunAgainstReference makes sure that the keys of a multimap's index come
// back as stored however they look. It belongs to the page prototype of the
// redesign (docs/redesign, step 1), which holds many keys of any length in
// one object: the test fills a run of pages with every shape of key, through
// growth, splits, removals that shrink and merge the pages, and compares
// lookups of present and absent keys and ordered iteration from random
// starts with a sorted map, checking the layout of every page.
func TestRunAgainstReference(t *testing.T) {
	for name, gen := range shapes {
		for _, byBytes := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/split by bytes=%v", name, byBytes), func(t *testing.T) {
				SplitByBytes = byBytes
				defer func() { SplitByBytes = false }()
				r := rand.New(rand.NewPCG(1, 2))
				var run Run
				ref := reference{}
				for phase := range 6 {
					removing := phase%2 == 1
					for range 1500 {
						k := gen(r)
						v := r.Uint64()
						if removing && r.IntN(3) > 0 {
							ks := ref.sorted()
							if len(ks) > 0 && r.IntN(2) == 0 {
								k = []byte(ks[r.IntN(len(ks))])
							}
							_, had := ref[string(k)]
							if run.Delete(k) != had {
								t.Fatalf("Delete(%x) = %v, want %v", k, !had, had)
							}
							delete(ref, string(k))
							continue
						}
						if err := run.Insert(k, v); err != nil {
							t.Fatal(err)
						}
						ref[string(k)] = v
					}
					verify(t, &run, ref, gen, r)
				}
			})
		}
	}
}

// verify fails the test unless the run holds exactly ref, in order.
func verify(t *testing.T, run *Run, ref reference, gen func(*rand.Rand) []byte, r *rand.Rand) {
	t.Helper()
	for _, p := range run.Pages() {
		check(t, p)
	}
	if run.Len() != len(ref) {
		t.Fatalf("%d keys, want %d", run.Len(), len(ref))
	}
	ks := ref.sorted()
	for _, k := range ks {
		if v, ok := run.Get([]byte(k)); !ok || v != ref[k] {
			t.Fatalf("Get(%x) = %d, %v; want %d", k, v, ok, ref[k])
		}
	}
	for range 300 {
		k := gen(r)
		if _, want := ref[string(k)]; want {
			continue
		}
		if _, ok := run.Get(k); ok {
			t.Fatalf("Get(%x) found an absent key", k)
		}
	}
	for range 20 {
		var from []byte
		start := 0
		if r.IntN(4) > 0 && len(ks) > 0 {
			from = gen(r)
			start = sort.SearchStrings(ks, string(from))
		}
		var got []string
		run.Each(from, func(s []byte, v uint64) bool {
			if v != ref[string(s)] {
				t.Fatalf("Each gave %x the value %d, want %d", s, v, ref[string(s)])
			}
			got = append(got, string(s))
			return true
		})
		if !slices.Equal(got, ks[start:]) {
			t.Fatalf("Each(%x) gave %d keys, want %d", from, len(got), len(ks)-start)
		}
	}
}

// TestPageGrowth makes sure that a page keeps every key as it fills up. It
// belongs to the page prototype of the redesign (docs/redesign, step 1): the
// page of the smallest class grows into the next larger one when its slots or
// its heap run out, uniform pages hold 7, 14 and 29 integer keys, and a page
// of the largest class that is full says so and stays as it was, so that the
// tree can split it.
func TestPageGrowth(t *testing.T) {
	for _, tc := range []struct {
		name      string
		remainder func(i int) []byte
		wantFull  int // keys when the largest class is full, at least
		uniform   bool
		wantClass int
	}{
		{"uniform integers fill 29 slots of 512 bytes", func(i int) []byte { return []byte{0, 0, 0, 0, 0, 0, byte(i >> 8), byte(i)} }, 29, true, 2},
		{"short keys of several lengths", func(i int) []byte { return bytes.Repeat([]byte{byte(i)}, 1+i%7) }, 20, false, 2},
		{"remainders with tails use the heap", func(i int) []byte { return []byte(fmt.Sprintf("%03d-a-remainder-with-a-tail", i)) }, 12, false, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := New(0, 0)
			if tc.uniform {
				p = New(0, 8)
			}
			classes := map[int]bool{}
			i := 0
			for ; ; i++ {
				q, res, err := p.Insert(tc.remainder(i), uint64(i))
				if err != nil {
					t.Fatal(err)
				}
				if res == Full {
					if q != p {
						t.Fatal("a full page was replaced")
					}
					break
				}
				p = q
				classes[p.Class()] = true
				check(t, p)
			}
			if p.Len() != i || i < tc.wantFull || p.Class() != tc.wantClass || p.Uniform() != tc.uniform {
				t.Errorf("%d keys, class %d, uniform %v; want at least %d, %d, %v", i, p.Class(), p.Uniform(), tc.wantFull, tc.wantClass, tc.uniform)
			}
			if len(classes) < 2 {
				t.Errorf("the page grew through classes %v, want more than one", classes)
			}
			for j := range i {
				if v, ok := p.Get(tc.remainder(j)); !ok || v != uint64(j) {
					t.Fatalf("Get(%d) = %d, %v", j, v, ok)
				}
			}
		})
	}
}

// TestPageFlavors makes sure that a page of integers turns into a general page
// when a key of another length arrives, and back into a uniform page when it
// is rebuilt with keys of one length. It belongs to the page prototype of the
// redesign (docs/redesign, step 1), whose uniform pages cost 17 bytes a key,
// the general ones 20 and the bytes of their tails.
func TestPageFlavors(t *testing.T) {
	p := New(0, 4)
	for i := range 5 {
		p, _, _ = p.Insert([]byte{0, 0, 0, byte(i)}, uint64(i))
	}
	if !p.Uniform() {
		t.Fatal("a page of keys of 4 bytes is not uniform")
	}
	p, res, err := p.Insert([]byte{0, 0, 0, 2, 7}, 99) // after a key it shares its head with
	if err != nil || res != Inserted || p.Uniform() {
		t.Fatalf("Insert of a longer key: %v %v uniform %v", res, err, p.Uniform())
	}
	check(t, p)
	for k, want := range map[string]uint64{"\x00\x00\x00\x02\x07": 99, "\x00\x00\x00\x02": 2, "\x00\x00\x00\x03": 3} {
		if v, ok := p.Get([]byte(k)); !ok || v != want {
			t.Errorf("Get(%x) = %d, %v; want %d", k, v, ok, want)
		}
	}
	p, ok := p.Delete([]byte{0, 0, 0, 2, 7})
	if !ok {
		t.Fatal("Delete of the long key failed")
	}
	left, right, err := p.Split()
	if err != nil || !left.Uniform() || !right.Uniform() {
		t.Fatalf("a split of keys of one length gave uniform %v %v, err %v", left.Uniform(), right.Uniform(), err)
	}
	merged := Merge(left, right)
	if merged == nil || !merged.Uniform() || merged.Len() != p.Len() {
		t.Fatalf("Merge of two uniform pages: %v", merged)
	}
}

// TestPageEdges makes sure that a page copes with what a tree will hand it by
// mistake or by design: a remainder too long for any page, an empty remainder, a
// delete of what is not there, updating a value in place, and the last key
// going. It belongs to the page prototype of the redesign (docs/redesign,
// step 1).
func TestPageEdges(t *testing.T) {
	p := New(0, 0)
	long := make([]byte, maxRemainder+1)
	if _, _, err := p.Insert(long, 1); err != ErrTooLong {
		t.Errorf("Insert of 256 bytes: %v, want ErrTooLong", err)
	}
	if _, ok := p.Get(long); ok {
		t.Error("Get found a remainder that is too long")
	}
	if _, ok := p.Delete(long); ok {
		t.Error("Delete found a remainder that is too long")
	}
	p, _, _ = p.Insert(nil, 7) // the empty remainder: the key ends where the page starts
	p, res, _ := p.Insert([]byte{}, 8)
	if res != Updated {
		t.Errorf("the second insert of the empty remainder: %v, want Updated", res)
	}
	if v, ok := p.Get(nil); !ok || v != 8 {
		t.Errorf("Get(empty) = %d, %v; want 8", v, ok)
	}
	if _, ok := p.Delete([]byte("absent")); ok {
		t.Error("Delete of an absent key reported success")
	}
	p, ok := p.Delete(nil)
	if !ok || p != nil {
		t.Errorf("Delete of the last key: %v %v, want the page gone", p, ok)
	}
}

// TestSplitAndMerge makes sure that a page that is full divides into two pages
// that hold all its keys and leave room to grow, and that thin neighbours join
// again, and not when they are too full. It belongs to the page prototype of the
// redesign (docs/redesign, step 1): whether to split by count or by bytes is one
// of the questions of that step.
func TestSplitAndMerge(t *testing.T) {
	if _, _, err := New(0, 0).Split(); err == nil {
		t.Error("Split of an empty page succeeded")
	}
	for _, byBytes := range []bool{false, true} {
		SplitByBytes = byBytes
		p := New(0, 0)
		var keys []string
		for i := 0; ; i++ {
			k := fmt.Sprintf("%s%02d", bytes.Repeat([]byte("x"), 1+i*3%11), i)
			q, res, _ := p.Insert([]byte(k), uint64(i))
			if res == Full {
				break
			}
			p = q
			keys = append(keys, k)
		}
		left, right, err := p.Split()
		SplitByBytes = false
		if err != nil || left.Len()+right.Len() != p.Len() || left.Len() == 0 || right.Len() == 0 {
			t.Fatalf("split by bytes=%v: %v, %d + %d of %d keys", byBytes, err, left.Len(), right.Len(), p.Len())
		}
		check(t, left)
		check(t, right)
		var lb, rb [maxRemainder]byte
		if bytes.Compare(left.Key(left.Len()-1, &lb), right.Key(0, &rb)) >= 0 {
			t.Errorf("the pages of the split are not in order")
		}
		for _, f := range []*Page{left, right} {
			if f.Size()*85 < need(f.planFor(0, f.Len(), nil, false))*100 {
				t.Errorf("a half of the split fills more than %d%%", SplitFill)
			}
		}
		if m := Merge(left, right); m != nil && m.Len() != len(keys) {
			t.Errorf("Merge lost keys: %d of %d", m.Len(), len(keys))
		}
	}
	a, b := New(0, 0), New(0, 0)
	a, _, _ = a.Insert([]byte("a"), 1)
	b, _, _ = b.Insert([]byte("b"), 2)
	m := Merge(a, b)
	if m == nil || m.Len() != 2 || m.Class() != 0 {
		t.Fatalf("Merge of two thin pages: %v", m)
	}
	check(t, m)
	big := New(2, 0)
	for i := range 20 {
		big, _, _ = big.Insert([]byte(fmt.Sprintf("big-%02d", i)), 1)
	}
	if Merge(big, big) != nil {
		t.Error("Merge of two pages that fill a class succeeded")
	}
}

// TestShrink makes sure that a page that has lost most of its keys moves into
// a smaller class and keeps the keys it has left, and that the tags of the
// removed keys no longer answer lookups. It belongs to the
// page prototype of the redesign (docs/redesign, step 1), whose pages should
// not stay as big as their fullest moment.
func TestShrink(t *testing.T) {
	p := New(0, 0)
	for i := range 20 {
		p, _, _ = p.Insert([]byte(fmt.Sprintf("key-%02d", i)), uint64(i))
	}
	big := p.Class()
	for i := range 17 {
		p, _ = p.Delete([]byte(fmt.Sprintf("key-%02d", i)))
	}
	if p.Class() >= big || p.Len() != 3 {
		t.Errorf("after the removals: class %d of %d, %d keys", p.Class(), big, p.Len())
	}
	check(t, p)
	q := New(2, 8)
	for i := range 20 {
		q, _, _ = q.Insert([]byte{0, 0, 0, 0, 0, 0, 0, byte(i)}, 1)
	}
	for i := range 9 {
		q, _ = q.Delete([]byte{0, 0, 0, 0, 0, 0, 0, byte(i)})
	}
	check(t, q)
	for i := 9; i < 20; i++ {
		if _, ok := q.Get([]byte{0, 0, 0, 0, 0, 0, 0, byte(i)}); !ok {
			t.Fatalf("key %d lost after the removals", i)
		}
	}
	if _, ok := q.Get([]byte{0, 0, 0, 0, 0, 0, 0, 3}); ok {
		t.Error("a removed key is still found")
	}
}

// TestEach makes sure that a range scan over a page starts at the right key
// and stops when asked. It belongs to the page prototype of the redesign
// (docs/redesign, step 1): ranges are the reason for pages, and a scan walks
// the keys of a page in one run of memory.
func TestEach(t *testing.T) {
	p := New(0, 0)
	for _, k := range []string{"b", "d", "f", "h"} {
		p, _, _ = p.Insert([]byte(k), 1)
	}
	for _, tc := range []struct {
		name string
		from []byte
		stop int
		want string
	}{
		{"from the start", nil, 0, "bdfh"},
		{"from a key that is there", []byte("d"), 0, "dfh"},
		{"from a key between two", []byte("e"), 0, "fh"},
		{"from beyond the last", []byte("z"), 0, ""},
		{"stopped after two", nil, 2, "bd"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []byte
			p.Each(tc.from, func(s []byte, v uint64) bool {
				got = append(got, s...)
				return tc.stop == 0 || len(got) < tc.stop
			})
			if string(got) != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
	if p.Val(0) != 1 {
		t.Errorf("Val(0) = %d, want 1", p.Val(0))
	}
	var r Run
	r.Each(nil, func([]byte, uint64) bool { t.Error("an empty run called fn"); return true })
	if err := r.Insert(make([]byte, maxRemainder+1), 1); err != ErrTooLong {
		t.Errorf("Run.Insert of a remainder that is too long: %v", err)
	}
	for i := range 200 {
		_ = r.Insert([]byte(fmt.Sprintf("run-key-%03d", i)), uint64(i))
	}
	if len(r.Pages()) < 3 {
		t.Fatalf("%d pages for 200 keys", len(r.Pages()))
	}
	n := 0
	r.Each([]byte("run-key-050"), func(s []byte, v uint64) bool { n++; return n < 120 }) // across pages, stopped inside one
	if n != 120 {
		t.Errorf("Run.Each visited %d keys, want 120", n)
	}
}

// TestLayout makes sure that a page is exactly as big as its class and that its
// arrays and its heap fit in it for every capacity a rebuild can choose. It
// belongs to the page prototype of the redesign (docs/redesign, step 1), whose
// rules are that an object is 128, 256, 512 or 1024 bytes, that its header is
// 16 bytes, and that nothing but the heap grows into the free space.
func TestLayout(t *testing.T) {
	if unsafe.Sizeof(Page{}) != hdr {
		t.Fatalf("the header is %d bytes, want %d", unsafe.Sizeof(Page{}), hdr)
	}
	for c, size := range sizes {
		if got := len(New(c, 0).mem()); got != size {
			t.Errorf("class %d is %d bytes, want %d", c, got, size)
		}
		for _, ulen := range []int{0, 1, 8} {
			p := New(c, ulen)
			if end := arraysEnd(int(p.cap), ulen != 0, 0); end > int(p.top) || int(p.top) != size {
				t.Errorf("class %d, ulen %d: arrays end at %d, heap at %d of %d", c, ulen, end, p.top, size)
			}
		}
		for tails := 0; tails <= size; tails += 7 {
			for n := 1; n <= 40; n++ {
				pl := plan{n: n, tails: tails}
				if need(pl) > size {
					continue
				}
				q := newPage(c, pl, nil)
				if int(q.cap) < n || arraysEnd(int(q.cap), false, 0)+tails > size {
					t.Fatalf("class %d, %d keys, %d bytes of tails: cap %d, arrays end %d", c, n, tails, q.cap, arraysEnd(int(q.cap), false, 0))
				}
			}
		}
	}
	for c := 1; c < 40; c++ {
		if base(c, true, 0)%16 != 0 || base(c, false, 0)%16 != 0 || base(c, true, 0) < hdr+c || base(c, false, 0) < hdr+4*c {
			t.Errorf("the arrays of capacity %d start at %d (uniform) and %d: not a multiple of 16, or inside the directory", c, base(c, true, 0), base(c, false, 0))
		}
	}
}

// FuzzRun makes sure that no sequence of inserts and deletes of arbitrary keys
// breaks a page's layout or loses a key. It belongs to the page prototype of
// the redesign (docs/redesign, step 1), and checks the invariants of every page
// of a run after the sequence, against a sorted map.
func FuzzRun(f *testing.F) {
	f.Add([]byte("abc\x00de\x00\x00abcdefghij\x01"))
	f.Add(bytes.Repeat([]byte{0xff, 1, 0}, 40))
	f.Fuzz(func(t *testing.T, data []byte) {
		var run Run
		ref := reference{}
		for len(data) > 1 {
			n := min(int(data[0])%40, len(data)-1)
			k, op := data[1:1+n], data[0]>>6
			data = data[1+n:]
			if op == 0 {
				if run.Delete(k) != (ref[string(k)] != 0 || contains(ref, k)) {
					t.Fatalf("Delete(%x) disagrees with the reference", k)
				}
				delete(ref, string(k))
				continue
			}
			if err := run.Insert(k, uint64(len(k))+1); err != nil {
				t.Fatal(err)
			}
			ref[string(k)] = uint64(len(k)) + 1
		}
		verify(t, &run, ref, shapes["any"], rand.New(rand.NewPCG(3, 4)))
	})
}

func contains(r reference, k []byte) bool { _, ok := r[string(k)]; return ok }

// TestStats makes sure that the numbers the fill experiments rely on are what
// the layout says. It belongs to the page prototype of the redesign
// (docs/redesign, step 1), whose bytes per key come from them: a page of
// integers needs 17 bytes a key after its header, a page of strings 20 and the
// bytes of their tails.
func TestStats(t *testing.T) {
	u := New(1, 8)
	for i := range 10 {
		u, _, _ = u.Insert([]byte{0, 0, 0, 0, 0, 0, 0, byte(i)}, 1)
	}
	// the keys share 7 bytes, but short remainders have no tails to shorten: no prefix
	if got, want := u.Stats(), (Stats{Keys: 10, Size: 256, Used: arraysEnd(10, true, 0), Uniform: true}); got != want {
		t.Errorf("uniform page: %+v, want %+v", got, want)
	}
	// long keys with a long common prefix: a rebuild (as every growth of the page is) finds it
	p := New(0, 0)
	for i := range 10 {
		p, _, _ = p.Insert([]byte(fmt.Sprintf("a/long/prefix/of/the/keys/%02d", i)), 1)
	}
	if p.PrefixLen() != 27 || p.Stats().Prefix != 27 {
		t.Errorf("the page has a prefix of %d bytes, a rebuild would find %d, want 27", p.PrefixLen(), p.Stats().Prefix)
	}
	if l, r, _ := p.Split(); l.PrefixLen() != 27 || r.PrefixLen() != 27 || Merge(l, r).PrefixLen() != 27 {
		t.Errorf("the halves and their merge have prefixes of %d, %d and %d bytes, want 27", l.PrefixLen(), r.PrefixLen(), Merge(l, r).PrefixLen())
	}
	g := New(1, 0)
	for _, k := range []string{"short", "a-key-of-twenty-chars", "another-key-of-more-than-twenty"} {
		g, _, _ = g.Insert([]byte(k), 1)
	}
	tails := (21 - headLen) + (31 - headLen)
	if got, want := g.Stats(), (Stats{Keys: 3, Size: 256, Used: arraysEnd(3, false, 0) + tails, Tails: tails}); got != want {
		t.Errorf("general page: %+v, want %+v", got, want)
	}
}

// TestSharedBy makes sure that a page finds the bytes its keys share, which it
// can then store once. It belongs to the page prototype of the redesign
// (docs/redesign, step 1), whose pages shorten the tails of long keys by a
// prefix: a prefix that is too long would corrupt keys, one that is too short
// wastes memory. The test builds pages of every shape of key and compares what
// the page computes from its head words and tails, for the keys of a range of
// positions with and without one more remainder, with the common prefix of the
// keys written out in full.
func TestSharedBy(t *testing.T) {
	defer func(a, b int) { MinPrefix, MinGain = a, b }(MinPrefix, MinGain)
	MinPrefix, MinGain = 1, 0
	for name, gen := range shapes {
		t.Run(name, func(t *testing.T) {
			r := rand.New(rand.NewPCG(3, 4))
			if name == "long" { // keys of 200 bytes and more leave room for one in a page
				gen = func(r *rand.Rand) []byte { return rbytes(r, 100+r.IntN(20), 2) }
			}
			p := New(2, 0)
			for range 50 { // fill a page, until it is full: a page of long keys holds two
				q, res, _ := p.Insert(gen(r), 1)
				if res == Full {
					break
				}
				p = q
			}
			if p.Len() < 2 {
				t.Fatalf("the page holds %d keys, want at least 2", p.Len())
			}
			var buf [maxRemainder]byte
			full := func(i int) []byte { return bytes.Clone(p.Key(i, &buf)) }
			for from := range p.Len() {
				for to := from; to <= p.Len(); to++ {
					for _, extra := range []bool{false, true} {
						s := gen(r)
						if extra && to > from && r.IntN(2) == 0 { // often share a long prefix with a key
							k := full(from)
							n := r.IntN(len(k) + 1)
							s = append(k[:n:n], s...)
							s = s[:min(len(s), maxRemainder)]
						}
						group := [][]byte{}
						for i := from; i < to; i++ {
							group = append(group, full(i))
						}
						if extra {
							group = append(group, s)
						}
						want := 0
						if len(group) >= 2 {
							want = len(group[0])
							for _, k := range group[1:] {
								want = min(want, lcp(group[0], k))
							}
						}
						if got := p.sharedBy(from, to, s, extra); got != want {
							t.Fatalf("keys %d..%d, extra %v: shared %d bytes, want %d", from, to, extra, got, want)
						}
					}
				}
			}
		})
	}
}
