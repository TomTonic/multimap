package swar

import (
	"bytes"
	"math/bits"
	"math/rand/v2"
	"testing"
)

// TestIndex8 checks that finding a key byte in an ART node returns the first
// position holding that byte, or 8 when none does, for every byte value and
// for words containing repeated and zero bytes (where the underlying
// zero-byte trick can flag false positives above a genuine match).
func TestIndex8(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	for range 20000 {
		var w [8]byte
		for i := range w {
			w[i] = byte(r.IntN(4)) // small alphabet: many repeats and zeros
		}
		b := byte(r.IntN(5))
		want := bytes.IndexByte(w[:], b)
		if want < 0 {
			want = 8
		}
		if got := Index8(Word(w[:]), b); got != want {
			t.Fatalf("Index8(%v, %d) = %d, want %d", w, b, got, want)
		}
	}
}

// TestRankAndHas checks the bitmap operations of the 57-way node: a child's
// index is the number of present bytes below it.
func TestRankAndHas(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	for range 2000 {
		var bm [4]uint64
		present := map[byte]bool{}
		for range r.IntN(100) {
			b := byte(r.IntN(256))
			Set(&bm, b)
			present[b] = true
		}
		for b := range 256 {
			want := 0
			for x := range b {
				if present[byte(x)] {
					want++
				}
			}
			if got := Rank(&bm, byte(b)); got != want {
				t.Fatalf("Rank(%d) = %d, want %d", b, got, want)
			}
			if Has(&bm, byte(b)) != present[byte(b)] {
				t.Fatalf("Has(%d) wrong", b)
			}
		}
		n := 0
		for _, w := range bm {
			n += bits.OnesCount64(w)
		}
		if n != len(present) {
			t.Fatalf("bitmap holds %d bits, want %d", n, len(present))
		}
	}
}

// TestMatchPrefix checks the compressed-path test of an inner node: a key can
// pass the node only if it is long enough and its next bytes equal the stored
// prefix, of which at most 8 bytes are compared.
func TestMatchPrefix(t *testing.T) {
	cases := []struct {
		name   string
		prefix string
		plen   int
		key    string
		depth  int
		want   bool
	}{
		{"short prefix matches", "ab", 2, "xabc", 1, true},
		{"short prefix differs", "ab", 2, "xacc", 1, false},
		{"key too short", "ab", 2, "xa", 1, false},
		{"exactly consumed", "ab", 2, "xab", 1, true},
		{"8 byte prefix via word compare", "abcdefgh", 8, "abcdefghij", 0, true},
		{"8 byte prefix differs in last byte", "abcdefgh", 8, "abcdefgXij", 0, false},
		{"long prefix checks only first 8", "abcdefgh", 12, "abcdefghXXXX", 0, true},
		{"long prefix needs full length", "abcdefgh", 12, "abcdefghXXX", 0, false},
		{"word path with short prefix", "ab", 2, "abXXXXXXXX", 0, true},
		{"word path with short prefix differing", "ab", 2, "aXXXXXXXXX", 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var p [8]byte
			copy(p[:], c.prefix)
			if got := MatchPrefix(&p, c.plen, []byte(c.key), c.depth); got != c.want {
				t.Fatalf("MatchPrefix = %v, want %v", got, c.want)
			}
		})
	}
}

// TestLcp checks the longest-common-prefix helper used when splitting paths.
func TestLcp(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	for range 20000 {
		a := make([]byte, r.IntN(30))
		for i := range a {
			a[i] = byte(r.IntN(3))
		}
		b := append([]byte(nil), a[:r.IntN(len(a)+1)]...)
		for range r.IntN(20) {
			b = append(b, byte(r.IntN(3)))
		}
		want := 0
		for want < len(a) && want < len(b) && a[want] == b[want] {
			want++
		}
		if got := Lcp(a, b); got != want {
			t.Fatalf("Lcp(%v, %v) = %d, want %d", a, b, got, want)
		}
	}
}
