package vpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

// sortedKeys returns up to n distinct keys of gen, sorted.
func sortedKeys(r *rand.Rand, gen func(*rand.Rand) []byte, n int) [][]byte {
	seen := map[string]bool{}
	var ks [][]byte
	for tries := 0; len(ks) < n && tries < 50*n; tries++ { // a shape may have fewer distinct keys than n
		if k := gen(r); !seen[string(k)] {
			seen[string(k)] = true
			ks = append(ks, k)
		}
	}
	slices.SortFunc(ks, bytes.Compare)
	return ks
}

// fitting returns the most keys of gen, up to n, sorted, that fit one page.
func fitting(r *rand.Rand, gen func(*rand.Rand) []byte, n int) [][]byte {
	for ; n > 1; n-- {
		ks := sortedKeys(r, gen, n)
		if Build(0, len(ks), func(i int) []byte { return ks[i] }, func(int) uint64 { return 0 }) != nil {
			return ks
		}
	}
	return sortedKeys(r, gen, 1)
}

// pageOf builds a page of keys with the values 0, 1, 2 ... or fails the test.
func pageOf(t *testing.T, base int, ks [][]byte) *Page {
	t.Helper()
	p := Build(base, len(ks), func(i int) []byte { return ks[i] }, func(i int) uint64 { return uint64(i) })
	if p == nil {
		t.Fatalf("%d keys did not fit a page", len(ks))
	}
	return p
}

// TestBuild makes sure that a page can be built from sorted keys in one step,
// which is how the tree rebuilds a subtree that has outgrown its pages. It
// belongs to the page of the redesign (docs/redesign, step 2), whose tree turns
// the keys of a bursting page into pages again: the keys of every shape come
// back in order with their values and the base they were given, keys that
// do not fit are refused, and so are too many, too long and none.
func TestBuild(t *testing.T) {
	for name, gen := range shapes {
		t.Run(name, func(t *testing.T) {
			r := rand.New(rand.NewPCG(11, 12))
			for n := 1; n <= 40; n++ {
				ks := sortedKeys(r, gen, min(n, 3))
				if name != "long" && name != "any" {
					ks = sortedKeys(r, gen, n)
				}
				p := Build(7, len(ks), func(i int) []byte { return ks[i] }, func(i int) uint64 { return uint64(100 + i) })
				if p == nil {
					continue // too many keys, or too long, for the largest class
				}
				check(t, p)
				if p.Base() != 7 || p.Len() != len(ks) {
					t.Fatalf("base %d, %d keys, want 7 and %d", p.Base(), p.Len(), len(ks))
				}
				var buf [maxSuffix]byte
				for i, k := range ks {
					if !bytes.Equal(p.Key(i, &buf), k) || p.Val(i) != uint64(100+i) {
						t.Fatalf("entry %d is %x = %d, want %x = %d", i, p.Key(i, &buf), p.Val(i), k, 100+i)
					}
					if v, ok := p.Get(k); !ok || v != uint64(100+i) {
						t.Fatalf("Get(%x) = %d, %v", k, v, ok)
					}
				}
			}
		})
	}
	ks := sortedKeys(rand.New(rand.NewPCG(1, 2)), shapes["u64"], 300)
	if Build(0, 0, nil, nil) != nil {
		t.Error("no keys made a page")
	}
	if Build(0, len(ks), func(i int) []byte { return ks[i] }, func(int) uint64 { return 0 }) != nil {
		t.Error("300 keys made a page")
	}
	long := [][]byte{make([]byte, maxSuffix+1)}
	if Build(0, 1, func(i int) []byte { return long[i] }, func(int) uint64 { return 0 }) != nil {
		t.Error("a suffix of 256 bytes made a page")
	}
	if p := pageOf(t, 0, [][]byte{bytes.Repeat([]byte("x"), maxSuffix)}); p.Len() != 1 {
		t.Error("a suffix of 255 bytes did not make a page")
	}
}

// TestLocateAndFind makes sure that the two ways to look for a key agree, and
// that the ordered one places keys that are not there, also longer ones than a
// page holds. It belongs to the page of the redesign (docs/redesign, step 2):
// the tree finds a key with Find on every lookup and with Locate where the
// position among the others matters, as at the bounds of a range scan.
func TestLocateAndFind(t *testing.T) {
	for name, gen := range shapes {
		t.Run(name, func(t *testing.T) {
			r := rand.New(rand.NewPCG(13, 14))
			ks := fitting(r, gen, 12)
			p := pageOf(t, 0, ks)
			for i, k := range ks {
				if j, ok := p.Find(k); !ok || j != i {
					t.Fatalf("Find(%x) = %d, %v, want %d", k, j, ok, i)
				}
				if j, ok := p.Locate(k); !ok || j != i {
					t.Fatalf("Locate(%x) = %d, %v, want %d", k, j, ok, i)
				}
			}
			probes := append(sortedKeys(r, gen, 50), bytes.Repeat([]byte{0xff}, maxSuffix+20), nil,
				ks[0][:len(ks[0])/2], append(bytes.Clone(ks[0]), 7)) // a prefix of a key, and a key and a byte
			for _, k := range probes {
				want, found := sort2(ks, k)
				if i, ok := p.Locate(k); i != want || ok != found {
					t.Fatalf("Locate(%x) = %d, %v, want %d, %v", k, i, ok, want, found)
				}
				if _, ok := p.Find(k); ok != found {
					t.Fatalf("Find(%x) = %v, want %v", k, ok, found)
				}
			}
		})
	}
}

// sort2 returns where k would go among the sorted keys ks, and whether it is one.
func sort2(ks [][]byte, k []byte) (int, bool) {
	i, found := slices.BinarySearchFunc(ks, k, bytes.Compare)
	return i, found
}

// TestPageAccessors makes sure that the pieces the tree reads a page through
// give back what the page holds: the byte of a key at an offset, all keys with
// their prefix, the values, the address of a value that the tree may change,
// and whether a page is thin enough to look for a neighbour. It belongs to the
// page of the redesign (docs/redesign, step 2), whose tree reads keys, splits
// pages at a byte boundary and merges thin pages.
func TestPageAccessors(t *testing.T) {
	pre := "common/prefix/of/some/long/directory/"
	ks := [][]byte{[]byte(pre + "a"), []byte(pre + "abcdefghijkl"), []byte(pre + "b")}
	p := pageOf(t, 3, ks)
	if p.PrefixLen() == 0 {
		t.Fatal("the keys of the page share 37 bytes, but it has no prefix")
	}
	for i, k := range ks {
		for off := range len(k) + 2 {
			want := byte(0)
			if off < len(k) {
				want = k[off]
			}
			if got := p.ByteAt(i, off); got != want {
				t.Fatalf("ByteAt(%d, %d) = %q, want %q", i, off, got, want)
			}
		}
		if got := p.AppendKey([]byte("x"), i); !bytes.Equal(got, append([]byte("x"), k...)) {
			t.Fatalf("AppendKey(%d) = %q", i, got)
		}
	}
	if !slices.Equal(p.Vals(), []uint64{0, 1, 2}) {
		t.Fatalf("Vals = %v", p.Vals())
	}
	*p.ValPtr(1) = 42
	if v, _ := p.Get(ks[1]); v != 42 {
		t.Fatalf("Get after writing through ValPtr = %d, want 42", v)
	}
	if !p.Thin() {
		t.Error("a page of three keys is not thin")
	}
	var many [][]byte
	for i := range 28 {
		many = append(many, []byte{0, 0, 0, 0, 0, 0, 0, byte(i)})
	}
	if pageOf(t, 0, many).Thin() {
		t.Error("a page of 28 integers is thin")
	}
}

// TestRebase makes sure that a page can move up to a shallower base, which the
// tree does when the node above it goes away: every key gets the bytes between
// the two bases in front, and a page that has no room for them says so. It
// belongs to the page of the redesign (docs/redesign, step 2).
func TestRebase(t *testing.T) {
	ks := [][]byte{[]byte("aa"), []byte("ab"), []byte("b")}
	p := pageOf(t, 10, ks)
	q := p.Rebase(4, []byte("pqrstu"))
	if q == nil || q.Base() != 4 || q.Len() != 3 {
		t.Fatalf("rebased page %v", q)
	}
	check(t, q)
	var buf [maxSuffix]byte
	for i, k := range ks {
		if got := q.Key(i, &buf); !bytes.Equal(got, append([]byte("pqrstu"), k...)) {
			t.Fatalf("key %d is %q after the rebase", i, got)
		}
	}
	var many [][]byte
	for i := range 29 {
		many = append(many, []byte{0, 0, 0, 0, 0, 0, 0, byte(i)})
	}
	if pageOf(t, 40, many).Rebase(0, bytes.Repeat([]byte("p"), 40)) != nil {
		t.Error("a full page took 40 more bytes in each key")
	}
}

// TestDeleteAtAndSplitAt makes sure that the tree can remove an entry it has
// found by its position and cut a page where it likes, with every key still
// where it was. It belongs to the page of the redesign (docs/redesign, step 2).
func TestDeleteAtAndSplitAt(t *testing.T) {
	r := rand.New(rand.NewPCG(15, 16))
	ks := fitting(r, shapes["uuid"], 12)
	p := pageOf(t, 0, ks)
	if len(ks) < 8 {
		t.Fatalf("only %d uuids fit a page", len(ks))
	}
	l, rt := p.SplitAt(5)
	check(t, l)
	check(t, rt)
	if l.Len() != 5 || rt.Len() != len(ks)-5 || l.Base() != 0 || rt.Base() != 0 {
		t.Fatalf("halves of %d and %d keys", l.Len(), rt.Len())
	}
	var buf [maxSuffix]byte
	for i := range 5 {
		if !bytes.Equal(l.Key(i, &buf), ks[i]) {
			t.Fatalf("left key %d is wrong", i)
		}
	}
	for i := range rt.Len() {
		if !bytes.Equal(rt.Key(i, &buf), ks[5+i]) {
			t.Fatalf("right key %d is wrong", i)
		}
	}
	before := rt.Len()
	q := rt.DeleteAt(2)
	check(t, q)
	if q.Len() != before-1 {
		t.Fatalf("%d keys after a delete", q.Len())
	}
	if _, ok := q.Get(ks[7]); ok {
		t.Fatal("the deleted key is still there")
	}
	one := pageOf(t, 0, ks[:1])
	if one.DeleteAt(0) != nil {
		t.Fatal("deleting the last entry did not empty the page")
	}
}

// TestKeyBasedCalls makes sure that the calls that take a whole key and cut it
// at the page's base give the answers of those that take the suffix. It belongs
// to the page of the redesign (docs/redesign, step 2), whose tree holds the keys
// of a page from a base on and passes the whole key down: the head word is
// taken from the key's last 8 bytes when the suffix is shorter, so the keys here
// are short, 8 bytes and long, and their bases are zero and below their length.
func TestKeyBasedCalls(t *testing.T) {
	r := rand.New(rand.NewPCG(21, 22))
	for off := 0; off < 6; off++ { // headWord against word
		for n := off; n < 20; n++ {
			k := rbytes(r, n, 256)
			if got, want := headWord(k, off), word(k[off:]); got != want {
				t.Fatalf("headWord(%x, %d) = %x, want %x", k, off, got, want)
			}
		}
	}
	for _, base := range []int{0, 2, 5} {
		for name, gen := range map[string]func(*rand.Rand) []byte{
			"short": func(r *rand.Rand) []byte { return rbytes(r, 3+r.IntN(4), 6) },
			"eight": func(r *rand.Rand) []byte { return rbytes(r, 8, 256) },
			"long":  func(r *rand.Rand) []byte { return rbytes(r, 20+r.IntN(20), 4) },
			"deep":  func(r *rand.Rand) []byte { return append(bytes.Repeat([]byte("d"), 30), rbytes(r, 3, 4)...) },
		} {
			t.Run(fmt.Sprintf("%s from base %d", name, base), func(t *testing.T) {
				whole := func(s []byte) []byte { return append(bytes.Repeat([]byte{'p'}, base), s...) }
				ks := fitting(r, gen, 12)
				p := pageOf(t, base, ks)
				for i, k := range ks {
					w := whole(k)
					if j, ok := p.FindIn(w); !ok || j != i {
						t.Fatalf("FindIn(%x) = %d, %v, want %d", w, j, ok, i)
					}
					if j, ok := p.LocateIn(w); !ok || j != i {
						t.Fatalf("LocateIn(%x) = %d, %v, want %d", w, j, ok, i)
					}
				}
				for _, s := range sortedKeys(r, gen, 30) {
					want, found := sort2(ks, s)
					if i, ok := p.LocateIn(whole(s)); i != want || ok != found {
						t.Fatalf("LocateIn(%x) = %d, %v, want %d, %v", s, i, ok, want, found)
					}
					if _, ok := p.FindIn(whole(s)); ok != found {
						t.Fatalf("FindIn(%x) = %v, want %v", s, ok, found)
					}
					if found || p.Len() >= 40 {
						continue
					}
					q, res := p.InsertIn(want, whole(s), 99)
					if res == Full {
						break
					}
					p = q
					if v, ok := p.Get(s); !ok || v != 99 {
						t.Fatalf("after InsertIn(%x): Get = %d, %v", s, v, ok)
					}
					check(t, p)
					ks = slices.Insert(ks, want, s)
				}
			})
		}
	}
}

// TestSplitOffAndShared makes sure that a page can be cut in place: the first
// entries stay in the page, the others go to a new one, and the first ones move
// to a smaller class only if they fit it with room. It also checks the bytes
// the keys of a page share. It belongs to the page of the redesign
// (docs/redesign, step 2), whose tree splits full pages this way and, to tell
// where the keys of a page that cannot be split by a byte differ, asks how
// much of them is common.
func TestSplitOffAndShared(t *testing.T) {
	var ints [][]byte
	for i := range 29 {
		ints = append(ints, []byte{0, 0, 0, 0, 0, 0, 1, byte(i)})
	}
	p := pageOf(t, 0, ints)
	if p.Class() != 2 || p.Shared() != 7 {
		t.Fatalf("29 integers: class %d sharing %d bytes, want class 2 and 7", p.Class(), p.Shared())
	}
	l, r := p.SplitOff(14)
	check(t, l)
	check(t, r)
	if l != p || l.Len() != 14 || r.Len() != 15 {
		t.Fatalf("a split at 14: halves of %d and %d keys, first one in place: %v", l.Len(), r.Len(), l == p)
	}
	p = pageOf(t, 0, ints)
	l, r = p.SplitOff(3)
	check(t, l)
	if l == p || l.Class() >= 2 || l.Len() != 3 || r.Len() != 26 {
		t.Fatalf("a split at 3: halves of %d and %d keys, class %d, first one in place: %v; want a smaller class", l.Len(), r.Len(), l.Class(), l == p)
	}
	for i, k := range ints[:3] {
		if v, ok := l.Get(k); !ok || v != uint64(i) {
			t.Fatalf("Get(%x) = %d, %v after the split", k, v, ok)
		}
	}
	one := pageOf(t, 0, [][]byte{[]byte("alone")})
	if one.Shared() != 5 {
		t.Fatalf("a page of one key of 5 bytes shares %d", one.Shared())
	}
}

// TestCompaction makes sure that a page whose keys come and go does not have to
// be rebuilt in a new object each time its heap has run out of room: the tails
// of the keys that left are given back by moving the others together. It
// belongs to the page of the redesign (docs/redesign, step 2), whose tree runs
// pages through churn, where a key is removed and put back again and again.
func TestCompaction(t *testing.T) {
	r := rand.New(rand.NewPCG(23, 24))
	ks := fitting(r, shapes["uuid"], 9)
	p := pageOf(t, 0, ks)
	rebuilds := Counts.Rebuilds
	for round := range 300 {
		k := ks[round%len(ks)]
		q := p.DeleteAt(mustFind(t, p, k))
		if q == nil {
			t.Fatal("the page emptied")
		}
		i, _ := q.Locate(k)
		q, res := q.InsertAt(i, k, uint64(round))
		if res == Full {
			t.Fatal("a key that had just left did not fit its page")
		}
		p = q
		check(t, p)
	}
	if got := Counts.Rebuilds - rebuilds; got > 3 {
		t.Errorf("%d rebuilds in 300 deletes and inserts of the same key, want almost none", got)
	}
	for _, k := range ks {
		if _, ok := p.Get(k); !ok {
			t.Fatalf("key %x lost", k)
		}
	}
}

// mustFind returns the position of k in p or fails the test.
func mustFind(t *testing.T, p *Page, k []byte) int {
	t.Helper()
	i, ok := p.Find(k)
	if !ok {
		t.Fatalf("key %x is not in the page", k)
	}
	return i
}
