package lpage

import (
	"bytes"
	"encoding/binary"
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

// values are generators of values: of different lengths (from none to 40
// bytes, like a name), of one width (8 bytes, like a number), and mostly of one
// width with others now and then, which moves a page from one form to the other.
var values = map[string]func(*rand.Rand) []byte{
	"names":  func(r *rand.Rand) []byte { return rbytes(r, r.IntN(41), 256) },
	"eights": func(r *rand.Rand) []byte { return binary.LittleEndian.AppendUint64(nil, r.Uint64()) },
	"mostly eights": func(r *rand.Rand) []byte {
		if r.IntN(5) == 0 {
			return rbytes(r, r.IntN(20), 256)
		}
		return rbytes(r, 8, 256)
	},
}

// check verifies the invariants of a page against the entries expected in it.
func check(t *testing.T, p *Page, want map[string][]byte) {
	t.Helper()
	keys, vals := p.Entries()
	if len(keys) != len(want) || p.Len() != len(want) {
		t.Fatalf("%d entries, Len %d, want %d", len(keys), p.Len(), len(want))
	}
	if !slices.IsSortedFunc(keys, bytes.Compare) {
		t.Fatal("entries are not sorted")
	}
	for i, k := range keys {
		if v, ok := want[string(k)]; !ok || !bytes.Equal(v, vals[i]) {
			t.Fatalf("entry %q = %x, want %x (present %v)", k, vals[i], v, ok)
		}
		if got, ok := p.Get(k); !ok || !bytes.Equal(got, vals[i]) {
			t.Fatalf("Get(%q) = %x, %v, want %x", k, got, ok, vals[i])
		}
		if len(k) <= p.PrefixLen() || !bytes.HasPrefix(k, keys[0][:p.PrefixLen()]) {
			t.Fatalf("entry %q does not continue the prefix of %d bytes", k, p.PrefixLen())
		}
		if w := p.Width(); w != 0 && len(vals[i]) != w {
			t.Fatalf("a page of width %d holds a value of %d bytes", w, len(vals[i]))
		}
	}
	if h := p.HeaderLen(); h < headerFor(len(keys), p.Width()) || h > MaxHeader {
		t.Fatalf("header %d for %d entries of width %d", h, len(keys), p.Width())
	}
	if p.Used() > p.Size() {
		t.Fatalf("uses %d bytes of %d", p.Used(), p.Size())
	}
	r, v := p.lens()
	for i := len(keys); i < len(r); i++ {
		if r[i] != 0 || v != nil && v[i] != 0 {
			t.Fatal("a length after the last entry is not 0")
		}
	}
}

func has(m map[string][]byte, k []byte) bool { _, ok := m[string(k)]; return ok }

// TestPageAgainstModel makes sure that the page holds what was put into it
// whatever the keys and the values look like and in whatever order they come and
// go. It belongs to the multi-key page candidate with a header of lengths
// (docs/redesign, PLAN step 3), which the tree would use for the keys of many
// entries with one value each, a number or a name. For every kind of key, of
// value and of header a run of pages gets random inserts, updates and deletes
// and must agree with a map after each round, with every page in order and
// within its size.
func TestPageAgainstModel(t *testing.T) {
	for name, gen := range shapes {
		for vname, vgen := range values {
			for _, hdr := range []int{8, 16, 24, 32, 48, 64} {
				t.Run(fmt.Sprintf("%s/%s/header %d", name, vname, hdr), func(t *testing.T) {
					defer func(old int) { MaxHeader = old }(MaxHeader)
					MaxHeader = hdr
					r := rand.New(rand.NewPCG(1, uint64(hdr)))
					var run Run
					want := map[string][]byte{}
					for round := range 5 {
						for range 300 {
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
							v := vgen(r)
							if err := run.Insert(k, v); err != nil {
								if tooBig(k, v) {
									continue
								}
								t.Fatalf("Insert(%q): %v", k, err)
							}
							want[string(k)] = v
						}
						if run.Len() != len(want) {
							t.Fatalf("round %d: Len %d, want %d", round, run.Len(), len(want))
						}
						for k, v := range want {
							if got, ok := run.Get([]byte(k)); !ok || !bytes.Equal(got, v) {
								t.Fatalf("Get(%q) = %x, %v, want %x", k, got, ok, v)
							}
						}
						var prev []byte
						for _, p := range run.Pages() {
							keys, _ := p.Entries()
							sub := map[string][]byte{}
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
}

// TestPageEdges makes sure that the page says no where it has to and behaves at
// its limits. It belongs to the multi-key page candidate with a header of lengths
// (docs/redesign, PLAN step 3): an entry cannot have an empty suffix, a suffix or
// value beyond 255 bytes, or more than a page holds; a page with the most
// entries or in the largest class has to be split; a miss is a miss whether the
// prefix or the remainder differs; a value of another length changes the form of
// the page; and a page that thins out moves to a smaller class.
func TestPageEdges(t *testing.T) {
	one := []byte{1}
	t.Run("rejects an empty suffix and what does not fit", func(t *testing.T) {
		if _, err := Build([][]byte{{}}, [][]byte{one}); err != ErrEmpty {
			t.Errorf("Build empty: %v", err)
		}
		if _, err := Build([][]byte{make([]byte, 256)}, [][]byte{one}); err != ErrTooLong {
			t.Errorf("Build long: %v", err)
		}
		p, _ := Build([][]byte{{1}}, [][]byte{one})
		for name, s := range map[string][2][]byte{"empty": {nil, one}, "long suffix": {make([]byte, 256), one}, "long value": {{1}, make([]byte, 256)}, "both together": {make([]byte, 250), make([]byte, 255)}} {
			if _, _, err := p.Insert(s[0], s[1]); err == nil {
				t.Errorf("Insert with %s was accepted", name)
			}
		}
		var run Run
		if err := run.Insert(nil, one); err != ErrEmpty {
			t.Errorf("Run.Insert empty: %v", err)
		}
		if err := run.Insert(make([]byte, 256), one); err != ErrTooLong {
			t.Errorf("Run.Insert long: %v", err)
		}
	})
	t.Run("reports Full when the entries or the bytes run out", func(t *testing.T) {
		defer func(old int) { MaxHeader = old }(MaxHeader)
		MaxHeader = 8
		keys, vals := [][]byte{{1}, {2}, {3}}, [][]byte{{1}, {2, 2}, {3}}
		p, err := Build(keys, vals)
		if err != nil {
			t.Fatal(err)
		}
		if q, res, _ := p.Insert([]byte{9}, one); res != Full || q != p {
			t.Errorf("a fourth entry with values of different lengths: %v, want Full with the same page", res)
		}
		MaxHeader = 24
		big, _ := Build([][]byte{bytes.Repeat([]byte{'a'}, 200), bytes.Repeat([]byte{'b'}, 200)}, [][]byte{one, one})
		if big.Size() != 512 {
			t.Fatalf("two long keys: size %d", big.Size())
		}
		if _, res, _ := big.Insert(bytes.Repeat([]byte{'c'}, 200), one); res != Full {
			t.Errorf("a third long key: %v, want Full", res)
		}
	})
	t.Run("misses on the prefix and on the remainder", func(t *testing.T) {
		p, _ := Build([][]byte{[]byte("abcX"), []byte("abcY")}, [][]byte{one, {2}})
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
	t.Run("a key outside the prefix shortens it, a value of another length changes the form", func(t *testing.T) {
		eight := []byte("12345678")
		p, _ := Build([][]byte{[]byte("abcX"), []byte("abcY")}, [][]byte{eight, eight})
		if p.Width() != 8 {
			t.Fatalf("two values of 8 bytes: width %d", p.Width())
		}
		q, res, err := p.Insert([]byte("abZ"), eight)
		if err != nil || res != Inserted || q.PrefixLen() != 2 || q.Width() != 8 {
			t.Fatalf("Insert: %v %v prefix %d width %d", err, res, q.PrefixLen(), q.Width())
		}
		q, res, _ = q.Insert([]byte("abZ"), []byte("87654321"))
		if v, _ := q.Get([]byte("abZ")); res != Updated || string(v) != "87654321" {
			t.Errorf("a second insert of the key: %v, value %q", res, v)
		}
		q, res, _ = q.Insert([]byte("abZ"), []byte("short"))
		if v, _ := q.Get([]byte("abZ")); res != Updated || string(v) != "short" || q.Width() != 0 {
			t.Errorf("a value of another length: %v, value %q, width %d", res, v, q.Width())
		}
		q, res, _ = q.Insert([]byte("abQ"), []byte("a new one"))
		if v, _ := q.Get([]byte("abQ")); res != Inserted || string(v) != "a new one" {
			t.Errorf("a new entry in the other form: %v, value %q", res, v)
		}
		r, res, _ := p.Insert([]byte("abcZ"), []byte("odd"))
		if v, _ := r.Get([]byte("abcZ")); res != Inserted || string(v) != "odd" || r.Width() != 0 {
			t.Errorf("a new value of another length: %v %q width %d", res, v, r.Width())
		}
	})
	t.Run("thins out into smaller classes and then empties", func(t *testing.T) {
		var keys, vals [][]byte
		for i := range 20 {
			keys, vals = append(keys, []byte{'k', byte('a' + i), 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}), append(vals, []byte{byte(i), 1, 2, 3, 4, 5, 6, 7})
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
	t.Run("splits a page that one long entry fills", func(t *testing.T) {
		var run Run
		a, b, c := bytes.Repeat([]byte{'a'}, 250), bytes.Repeat([]byte{'b'}, 250), bytes.Repeat([]byte{'c'}, 250)
		for _, k := range [][]byte{b, a, c} {
			if err := run.Insert(k, bytes.Repeat([]byte{'v'}, 250)); err != nil {
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
// operations on short keys with values of the length the input says, and the
// run must agree with a map after them.
func FuzzOperations(f *testing.F) {
	f.Add([]byte("abcabdabe"), uint8(1))
	f.Add(bytes.Repeat([]byte{7}, 300), uint8(2))
	f.Fuzz(func(t *testing.T, data []byte, hdr uint8) {
		defer func(old int) { MaxHeader = old }(MaxHeader)
		MaxHeader = 8 * (1 + int(hdr)%8)
		var run Run
		want := map[string][]byte{}
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
			v := bytes.Repeat(data[i+1:i+2], int(data[i])%24)
			if err := run.Insert(k, v); err != nil {
				t.Fatal(err)
			}
			want[string(k)] = v
		}
		for k, v := range want {
			if got, ok := run.Get([]byte(k)); !ok || !bytes.Equal(got, v) {
				t.Fatalf("Get(%q) = %x, %v, want %x", k, got, ok, v)
			}
		}
		if run.Len() != len(want) {
			t.Fatalf("Len %d, want %d", run.Len(), len(want))
		}
	})
}

// TestRunPageFor makes sure that a benchmark can learn which page a suffix
// belongs to. It belongs to the multi-key page candidate with a header of lengths
// (docs/redesign, PLAN step 3) and its run of pages: the page of a suffix is the
// one that holds it, and an empty run has none.
func TestRunPageFor(t *testing.T) {
	var r Run
	if r.PageFor([]byte("a")) != nil {
		t.Error("an empty run has a page")
	}
	for i := range 200 {
		if err := r.Insert([]byte(fmt.Sprintf("key%04d", i)), []byte{byte(i)}); err != nil {
			t.Fatal(err)
		}
	}
	for i := range 200 {
		k := []byte(fmt.Sprintf("key%04d", i))
		if v, ok := r.PageFor(k).Get(k); !ok || v[0] != byte(i) {
			t.Fatalf("the page for %q holds %v, %v", k, v, ok)
		}
	}
}
