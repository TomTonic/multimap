package lpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

// rbytes returns n random bytes below alphabet.
func rbytes(r *rand.Rand, n, alphabet int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(r.IntN(alphabet))
	}
	return b
}

// shapes are generators of suffixes that stress the page differently; u64, uuid
// and path are those of the benchmarks of internal/vpage.
var shapes = map[string]func(*rand.Rand) []byte{
	"u64": func(r *rand.Rand) []byte { return rbytes(r, 8, 256) },
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
	"tiny": func(r *rand.Rand) []byte {
		k := make([]byte, 1+r.IntN(4))
		for i := range k {
			k[i] = "ab"[r.IntN(2)]
		}
		return k
	},
	"long": func(r *rand.Rand) []byte {
		k := make([]byte, 100+r.IntN(156))
		for i := range k {
			k[i] = "xy"[r.IntN(2)]
		}
		return k
	},
	"mixed": func(r *rand.Rand) []byte {
		k := make([]byte, 1+r.IntN(255))
		for i := range k {
			k[i] = byte('a' + r.IntN(2))
		}
		return k
	},
}

// check verifies the invariants of a page against the entries expected in it.
func check(t *testing.T, p *Page, want map[string]uint64) {
	t.Helper()
	keys, vals := p.Entries()
	if len(keys) != len(want) || p.Len() != len(want) {
		t.Fatalf("%d entries, Len %d, want %d", len(keys), p.Len(), len(want))
	}
	if !slices.IsSortedFunc(keys, bytes.Compare) {
		t.Fatal("entries are not sorted")
	}
	for i, k := range keys {
		if v, ok := want[string(k)]; !ok || v != vals[i] {
			t.Fatalf("entry %q = %d, want %d (present %v)", k, vals[i], v, ok)
		}
		if got, ok := p.Get(k); !ok || got != vals[i] {
			t.Fatalf("Get(%q) = %d, %v, want %d", k, got, ok, vals[i])
		}
		if len(k) <= p.PrefixLen() {
			t.Fatalf("entry %q is not longer than the prefix %d", k, p.PrefixLen())
		}
		if !bytes.HasPrefix(k, keys[0][:p.PrefixLen()]) {
			t.Fatalf("entry %q does not start with the prefix", k)
		}
	}
	if h := p.HeaderLen(); h < headerFor(len(keys)) || h > MaxHeader {
		t.Fatalf("header %d for %d entries", h, len(keys))
	}
	if p.Used() > p.Size() {
		t.Fatalf("uses %d bytes of %d", p.Used(), p.Size())
	}
	m := p.mem()
	for _, b := range m[2+len(keys) : p.HeaderLen()] {
		if b != 0 {
			t.Fatal("a length after the last entry is not 0")
		}
	}
}

// TestPageAgainstModel makes sure that the page holds what was put into it
// whatever the keys look like and in whatever order they come and go. It
// belongs to the multi-key page candidate with a header of lengths
// (docs/redesign, PLAN step 3), which the tree would use for the keys of
// many entries with one value each. For every kind of key a run of pages gets
// random inserts, updates and deletes and must agree with a map after each
// round, with every page in order and within its size.
func TestPageAgainstModel(t *testing.T) {
	for name, gen := range shapes {
		for _, hdr := range []int{8, 16, 24, 32} {
			t.Run(fmt.Sprintf("%s/header %d", name, hdr), func(t *testing.T) {
				defer func(old int) { MaxHeader = old }(MaxHeader)
				MaxHeader = hdr
				r := rand.New(rand.NewPCG(1, uint64(hdr)))
				var run Run
				want := map[string]uint64{}
				for round := range 6 {
					for range 400 {
						k := gen(r)
						if r.IntN(3) == 0 && len(want) > 0 {
							for w := range want { // a random key that is there
								k = []byte(w)
								break
							}
						}
						if r.IntN(4) == 0 {
							if got, was := run.Delete(k), has(want, k); got != was {
								t.Fatalf("Delete(%q) = %v, want %v", k, got, was)
							}
							delete(want, string(k))
							continue
						}
						v := uint64(r.IntN(1 << 30))
						if err := run.Insert(k, v); err != nil {
							t.Fatalf("Insert(%q): %v", k, err)
						}
						want[string(k)] = v
					}
					if run.Len() != len(want) {
						t.Fatalf("round %d: Len %d, want %d", round, run.Len(), len(want))
					}
					for k, v := range want {
						if got, ok := run.Get([]byte(k)); !ok || got != v {
							t.Fatalf("Get(%q) = %d, %v, want %d", k, got, ok, v)
						}
					}
					var prev []byte
					for _, p := range run.Pages() {
						keys, _ := p.Entries()
						sub := map[string]uint64{}
						for _, k := range keys {
							sub[string(k)] = want[string(k)]
						}
						check(t, p, sub)
						if prev != nil && bytes.Compare(prev, keys[0]) >= 0 {
							t.Fatal("pages are not in order")
						}
						prev = keys[len(keys)-1]
					}
				}
			})
		}
	}
}

func has(m map[string]uint64, k []byte) bool { _, ok := m[string(k)]; return ok }

// TestPageEdges makes sure that the page says no where it has to and behaves at
// its limits. It belongs to the multi-key page candidate with a header of lengths
// (docs/redesign, PLAN step 3): an entry cannot be empty or longer than 255
// bytes, a page with the most entries or in the largest class has to be split,
// a miss is a miss whether the prefix or the remainder differs, and a page
// that thins out moves to a smaller class.
func TestPageEdges(t *testing.T) {
	t.Run("rejects an empty suffix and one beyond 255 bytes", func(t *testing.T) {
		if _, err := Build([][]byte{{}}, []uint64{1}); err != ErrEmpty {
			t.Errorf("Build empty: %v", err)
		}
		if _, err := Build([][]byte{make([]byte, 256)}, []uint64{1}); err != ErrTooLong {
			t.Errorf("Build long: %v", err)
		}
		p, _ := Build([][]byte{{1}}, []uint64{1})
		if _, _, err := p.Insert(nil, 1); err != ErrEmpty {
			t.Errorf("Insert empty: %v", err)
		}
		if _, _, err := p.Insert(make([]byte, 256), 1); err != ErrTooLong {
			t.Errorf("Insert long: %v", err)
		}
		var run Run
		if err := run.Insert(nil, 1); err != ErrEmpty {
			t.Errorf("Run.Insert empty: %v", err)
		}
		if err := run.Insert(make([]byte, 256), 1); err != ErrTooLong {
			t.Errorf("Run.Insert long: %v", err)
		}
	})
	t.Run("reports Full when the entries or the bytes run out", func(t *testing.T) {
		defer func(old int) { MaxHeader = old }(MaxHeader)
		MaxHeader = 8
		var keys [][]byte
		var vals []uint64
		for i := range 6 {
			keys, vals = append(keys, []byte{byte(i + 1)}), append(vals, uint64(i))
		}
		p, err := Build(keys, vals)
		if err != nil {
			t.Fatal(err)
		}
		if q, res, _ := p.Insert([]byte{9}, 9); res != Full || q != p {
			t.Errorf("a seventh entry: %v, want Full with the same page", res)
		}
		MaxHeader = 24
		big, _ := Build([][]byte{bytes.Repeat([]byte{'a'}, 240), bytes.Repeat([]byte{'b'}, 240)}, []uint64{1, 2})
		if big.Size() != 512 {
			t.Fatalf("two long keys: size %d", big.Size())
		}
		if _, res, _ := big.Insert(bytes.Repeat([]byte{'c'}, 240), 3); res != Full {
			t.Errorf("a third long key: %v, want Full", res)
		}
	})
	t.Run("misses on the prefix and on the remainder", func(t *testing.T) {
		p, _ := Build([][]byte{[]byte("abcX"), []byte("abcY")}, []uint64{1, 2})
		if p.PrefixLen() != 3 {
			t.Fatalf("prefix %d, want 3", p.PrefixLen())
		}
		for _, k := range []string{"ab", "abc", "abdX", "abcZ", "abcXX"} {
			if _, ok := p.Get([]byte(k)); ok {
				t.Errorf("Get(%q) found a key", k)
			}
			if q, ok := p.Delete([]byte(k)); ok || q != p {
				t.Errorf("Delete(%q) = %v", k, ok)
			}
		}
	})
	t.Run("a key outside the prefix shortens it", func(t *testing.T) {
		p, _ := Build([][]byte{[]byte("abcX"), []byte("abcY")}, []uint64{1, 2})
		q, res, err := p.Insert([]byte("abZ"), 3)
		if err != nil || res != Inserted || q.PrefixLen() != 2 {
			t.Fatalf("Insert: %v %v prefix %d", err, res, q.PrefixLen())
		}
		q, res, _ = q.Insert([]byte("abZ"), 4)
		if res != Updated {
			t.Errorf("a second insert of the key: %v", res)
		}
		if v, _ := q.Get([]byte("abZ")); v != 4 {
			t.Errorf("value %d, want 4", v)
		}
	})
	t.Run("thins out into smaller classes and then empties", func(t *testing.T) {
		var keys [][]byte
		var vals []uint64
		for i := range 20 {
			keys, vals = append(keys, []byte{'k', byte('a' + i), 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}), append(vals, uint64(i))
		}
		p, err := Build(keys, vals)
		if err != nil {
			t.Fatal(err)
		}
		start := p.Class()
		for _, k := range keys[1:] {
			var ok bool
			if p, ok = p.Delete(k); !ok {
				t.Fatalf("Delete(%q) failed", k)
			}
		}
		if p.Class() >= start {
			t.Errorf("class %d after deleting all but one of 20, started at %d", p.Class(), start)
		}
		if p, ok := p.Delete(keys[0]); !ok || p != nil {
			t.Errorf("deleting the last entry: %v %v", p, ok)
		}
	})
	t.Run("splits a page that one long suffix fills", func(t *testing.T) {
		var run Run
		a, b, c := bytes.Repeat([]byte{'a'}, 255), bytes.Repeat([]byte{'b'}, 255), bytes.Repeat([]byte{'c'}, 255)
		for _, k := range [][]byte{b, a, c} {
			if err := run.Insert(k, 1); err != nil {
				t.Fatal(err)
			}
		}
		if run.Len() != 3 || len(run.Pages()) < 2 {
			t.Fatalf("%d keys in %d pages", run.Len(), len(run.Pages()))
		}
		for _, k := range [][]byte{a, b, c} {
			if _, ok := run.Get(k); !ok {
				t.Errorf("lost %c", k[0])
			}
		}
		var empty Run
		if _, ok := empty.Get(a); ok || empty.Delete(a) {
			t.Error("an empty run found a key")
		}
	})
}

// FuzzOperations makes sure that no sequence of inserts and deletes breaks a
// run of pages. It belongs to the multi-key page candidate with a header of
// lengths (docs/redesign, PLAN step 3); the input is read as a list of
// operations on short keys, and the run must agree with a map after them.
func FuzzOperations(f *testing.F) {
	f.Add([]byte("abcabdabe"), uint8(1))
	f.Add(bytes.Repeat([]byte{7}, 300), uint8(2))
	f.Fuzz(func(t *testing.T, data []byte, hdr uint8) {
		defer func(old int) { MaxHeader = old }(MaxHeader)
		MaxHeader = 8 * (1 + int(hdr)%4)
		var run Run
		want := map[string]uint64{}
		for i := 0; i+2 <= len(data); i += 2 {
			k := bytes.Repeat(data[i:i+1], 1+int(data[i+1])%40)
			k = append(k, data[i+1]%3)
			if data[i+1]&0x80 != 0 {
				if run.Delete(k) != has(want, k) {
					t.Fatal("Delete disagrees with the model")
				}
				delete(want, string(k))
				continue
			}
			if err := run.Insert(k, uint64(i)); err != nil {
				t.Fatal(err)
			}
			want[string(k)] = uint64(i)
		}
		for k, v := range want {
			if got, ok := run.Get([]byte(k)); !ok || got != v {
				t.Fatalf("Get(%q) = %d, %v, want %d", k, got, ok, v)
			}
		}
		if run.Len() != len(want) {
			t.Fatalf("Len %d, want %d", run.Len(), len(want))
		}
	})
}
