package lpage

import (
	"bytes"
	"math/rand/v2"
	"slices"
	"testing"
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
	q := cp.DeleteAt(k)
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
