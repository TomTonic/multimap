package skpage

import (
	"bytes"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"testing"
)

// model is what a page must hold: the remainder and the values in the order of
// their arrival.
type model struct {
	rest string
	vals []string
}

// check compares a page with its model and with its own layout: the header,
// the remainder, every value in order, the lookups, the copies of Strings, the
// used bytes and a zero tail.
func check(p *Page, m model) error {
	if p.Len() != len(m.vals) || string(p.Rest()) != m.rest || !p.Match([]byte(m.rest)) {
		return fmt.Errorf("page has %d values and remainder %q, want %d and %q", p.Len(), p.Rest(), len(m.vals), m.rest)
	}
	var got []string
	p.Each(func(v []byte) bool { got = append(got, string(v)); return true })
	var copied []string
	p.Strings(func(v string) bool { copied = append(copied, v); return true })
	if !slices.Equal(got, m.vals) || !slices.Equal(copied, m.vals) {
		return fmt.Errorf("values %q / %q, want %q", got, copied, m.vals)
	}
	for _, v := range m.vals {
		if !p.Has([]byte(v)) {
			return fmt.Errorf("Has(%q) is false", v)
		}
	}
	want := modelUsed(m)
	if p.Used() != want || want > p.Size() {
		return fmt.Errorf("Used = %d, want %d in an object of %d bytes", p.Used(), want, p.Size())
	}
	if tail := p.mem()[want:]; !bytes.Equal(tail, make([]byte, len(tail))) {
		return fmt.Errorf("the bytes behind the values are not zero: %x", tail)
	}
	return nil
}

// verify fails the test if check finds a difference.
func verify(t *testing.T, p *Page, m model) {
	t.Helper()
	if err := check(p, m); err != nil {
		t.Fatal(err)
	}
}

// TestNewAndBuild covers the entry of a key into the tree's single-key page: a
// remainder with its first value, or with a set of values that a value set gives
// back. The page takes everything within the limits of the layout and tells
// the tree with nil when a key does not fit, so that the tree keeps it in a
// value set. Each case names the limit it probes.
func TestNewAndBuild(t *testing.T) {
	long := func(n int) string { return strings.Repeat("x", n) }
	tests := []struct {
		name      string
		rest      string
		vals      []string
		new, bild int // the size of the page New and Build give, 0 for nil
	}{
		{"holds an empty remainder", "", []string{"a"}, 32, 32},
		{"holds an empty value", "ab", []string{""}, 32, 32},
		{"fits 32 bytes exactly", "ab", []string{long(26)}, 32, 32}, // 3 + 2 + 1 + 26
		{"needs 64 bytes for one byte more", "ab", []string{long(27)}, 64, 64},
		{"fits 64 bytes exactly", "ab", []string{long(58)}, 64, 64},
		{"needs 128 bytes for one byte more", "ab", []string{long(59)}, 128, 128},
		{"takes a value of 254 bytes", "", []string{long(254)}, 384, 384}, // 3 + 1 + 254 = 258
		{"takes the longest remainder", long(254), []string{"a"}, 384, 384},
		{"fits the largest class exactly", long(254), []string{long(254)}, 512, 512},
		{"holds many values", "k", []string{"a", "b", "c", "d"}, 32, 32},
		{"refuses a remainder of 255 bytes", long(255), []string{"a"}, 0, 0},
		{"refuses a value of 255 bytes", "a", []string{long(255)}, 0, 0},
		{"Build refuses content of 513 bytes", long(104), []string{long(200), long(204)}, 384, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New([]byte(tt.rest), []byte(tt.vals[0]))
			b := Build([]byte(tt.rest), bytesOf(tt.vals))
			if (p == nil) != (tt.new == 0) || (b == nil) != (tt.bild == 0) {
				t.Fatalf("New = %v, Build = %v", p, b)
			}
			if p != nil {
				if p.Size() != tt.new || sizes[p.Class()] != tt.new {
					t.Fatalf("New gives %d bytes, want %d", p.Size(), tt.new)
				}
				verify(t, p, model{tt.rest, tt.vals[:1]})
			}
			if b != nil {
				if b.Size() != tt.bild {
					t.Fatalf("Build gives %d bytes, want %d", b.Size(), tt.bild)
				}
				verify(t, b, model{tt.rest, tt.vals})
			}
		})
	}
	t.Run("Build refuses no values", func(t *testing.T) {
		if Build([]byte("a"), nil) != nil {
			t.Error("a page needs a value")
		}
	})
	t.Run("Build refuses many values that fill more than the largest class", func(t *testing.T) {
		var vals [][]byte
		for i := range 256 {
			vals = append(vals, []byte{byte(i)})
		}
		if Build(nil, vals) != nil {
			t.Error("256 values of one byte take 515 bytes")
		}
	})
}

func bytesOf(ss []string) [][]byte {
	out := make([][]byte, len(ss))
	for i, s := range ss {
		out[i] = []byte(s)
	}
	return out
}

// TestAddAndRemove covers what a user does to the values of one key: add one
// (a value that is there changes nothing), remove one. It belongs to the
// single-key page of the tree. The page grows into a larger class when the
// content outgrows it, shrinks when a removal leaves at most half of it
// in use, disappears with its last value, and tells the tree with Full when a
// value does not fit any more.
func TestAddAndRemove(t *testing.T) {
	t.Run("adds in place while the content fits", func(t *testing.T) {
		p := New([]byte("rest"), []byte("one"))
		q, res := p.Add([]byte("two"))
		if q != p || res != Added {
			t.Fatalf("got %p, %v; want the same page and Added", q, res)
		}
		verify(t, q, model{"rest", []string{"one", "two"}})
	})
	t.Run("answers Present for a value that is there", func(t *testing.T) {
		p := New([]byte("rest"), []byte("one"))
		p, _ = p.Add([]byte("two"))
		if q, res := p.Add([]byte("one")); q != p || res != Present {
			t.Fatalf("got %p, %v", q, res)
		}
		if q, res := p.Add([]byte("two")); q != p || res != Present {
			t.Fatalf("got %p, %v", q, res)
		}
		verify(t, p, model{"rest", []string{"one", "two"}})
		if p.Has([]byte("thr")) || p.Has([]byte("on")) || p.Has(nil) {
			t.Error("Has must be false for a value that is not there")
		}
	})
	t.Run("grows into the next class that holds the content", func(t *testing.T) {
		p := New([]byte("ab"), bytes.Repeat([]byte("x"), 26)) // 32 bytes
		q, res := p.Add([]byte("y"))
		if q == p || res != Added || q.Size() != 64 {
			t.Fatalf("got a page of %d bytes, %v", q.Size(), res)
		}
		verify(t, q, model{"ab", []string{strings.Repeat("x", 26), "y"}})
		verify(t, p, model{"ab", []string{strings.Repeat("x", 26)}}) // the old page is unchanged
	})
	t.Run("jumps over classes for a long value", func(t *testing.T) {
		p := New([]byte("ab"), []byte("v"))
		q, _ := p.Add(bytes.Repeat([]byte("z"), 200))
		if q.Size() != 256 {
			t.Fatalf("got %d bytes", q.Size())
		}
	})
	t.Run("answers Full for a value of 255 bytes", func(t *testing.T) {
		p := New([]byte("ab"), []byte("v"))
		if q, res := p.Add(make([]byte, 255)); q != p || res != Full {
			t.Fatalf("got %p, %v", q, res)
		}
		if p.Has(make([]byte, 255)) {
			t.Error("a value of 255 bytes cannot be there")
		}
	})
	t.Run("answers Full when the content would pass 512 bytes", func(t *testing.T) {
		p := New([]byte("ab"), bytes.Repeat([]byte("a"), 200))
		p, _ = p.Add(bytes.Repeat([]byte("b"), 200)) // 3 + 2 + 201 + 201 = 407
		before := p.Used()
		q, res := p.Add(bytes.Repeat([]byte("c"), 110)) // 407 + 111 = 518
		if q != p || res != Full || p.Used() != before {
			t.Fatalf("got %p, %v, used %d", q, res, p.Used())
		}
		if _, res := p.Add(bytes.Repeat([]byte("c"), 100)); res != Added { // 407 + 101 = 508
			t.Fatalf("got %v", res)
		}
	})
	t.Run("removes the first, a middle and the last value", func(t *testing.T) {
		for _, gone := range []string{"a", "bb", "ccc"} {
			p := Build([]byte("r"), bytesOf([]string{"a", "bb", "ccc"}))
			q, ok := p.Remove([]byte(gone))
			if !ok || q != p {
				t.Fatalf("remove %q: got %p, %v", gone, q, ok)
			}
			var want []string
			for _, v := range []string{"a", "bb", "ccc"} {
				if v != gone {
					want = append(want, v)
				}
			}
			verify(t, q, model{"r", want})
		}
	})
	t.Run("answers false for a value that is not there", func(t *testing.T) {
		p := New([]byte("r"), []byte("a"))
		if q, ok := p.Remove([]byte("b")); q != p || ok {
			t.Fatalf("got %p, %v", q, ok)
		}
		if q, ok := p.Remove(make([]byte, 255)); q != p || ok {
			t.Fatalf("got %p, %v", q, ok)
		}
	})
	t.Run("returns nil for the last value", func(t *testing.T) {
		p := New([]byte("r"), []byte("a"))
		if q, ok := p.Remove([]byte("a")); q != nil || !ok {
			t.Fatalf("got %p, %v", q, ok)
		}
	})
	t.Run("shrinks when the content fits a class of at most half the size", func(t *testing.T) {
		p := New([]byte("ab"), bytes.Repeat([]byte("x"), 26))
		p, _ = p.Add([]byte("y")) // 64 bytes, content 35
		q, ok := p.Remove([]byte("y"))
		if !ok || q == p || q.Size() != 32 {
			t.Fatalf("got a page of %d bytes", q.Size())
		}
		verify(t, q, model{"ab", []string{strings.Repeat("x", 26)}})
	})
	t.Run("keeps a page that would save less than half", func(t *testing.T) {
		p := New(nil, bytes.Repeat([]byte("a"), 200)) // 256
		p, _ = p.Add(bytes.Repeat([]byte("b"), 100))  // 3 + 201 + 101 = 305: 384
		if p.Size() != 384 {
			t.Fatalf("got %d bytes", p.Size())
		}
		q, _ := p.Remove(bytes.Repeat([]byte("b"), 100)) // content 204: class 256 saves a third
		if q != p || q.Size() != 384 {
			t.Fatalf("got a page of %d bytes", q.Size())
		}
	})
}

// TestEachAndStrings covers reading all values of a key, as a lookup or a scan
// does. Each hands out views of the page and Strings copies of it, all in
// one allocation, and both stop when the caller says so.
func TestEachAndStrings(t *testing.T) {
	p := Build([]byte("rest"), bytesOf([]string{"alpha", "", "gamma"}))
	t.Run("Each stops when fn returns false", func(t *testing.T) {
		n := 0
		if p.Each(func([]byte) bool { n++; return n < 2 }) || n != 2 {
			t.Fatalf("called %d times", n)
		}
	})
	t.Run("Strings stops when fn returns false", func(t *testing.T) {
		n := 0
		if p.Strings(func(string) bool { n++; return n < 3 }) || n != 3 {
			t.Fatalf("called %d times", n)
		}
	})
	t.Run("Strings makes one allocation for all the values", func(t *testing.T) {
		allocs := testing.AllocsPerRun(100, func() { p.Strings(func(string) bool { return true }) })
		if allocs != 1 {
			t.Fatalf("%v allocations", allocs)
		}
	})
	t.Run("Strings hands out an empty value", func(t *testing.T) {
		e := New([]byte("r"), nil)
		var got []string
		e.Strings(func(v string) bool { got = append(got, v); return true })
		if !slices.Equal(got, []string{""}) {
			t.Fatalf("got %q", got)
		}
	})
	t.Run("the strings stay valid when the page changes", func(t *testing.T) {
		q := Build([]byte("r"), bytesOf([]string{"hello", "world"}))
		var s []string
		q.Strings(func(v string) bool { s = append(s, v); return true })
		q.Remove([]byte("hello")) // shifts the bytes of "world" down in the page
		if s[0] != "hello" || s[1] != "world" {
			t.Fatalf("got %q", s)
		}
	})
	t.Run("AppendKey appends the remainder", func(t *testing.T) {
		if got := p.AppendKey([]byte("path:")); string(got) != "path:rest" {
			t.Fatalf("got %q", got)
		}
	})
}

// TestSkipAndPrepend covers the moves of a page in the tree: deeper when a
// node is split in above it (Skip, in place) and up when the node above it goes
// away (Prepend, into a larger class if need be). The remainder changes,
// the values stay, and a remainder or content beyond the limits is refused.
func TestSkipAndPrepend(t *testing.T) {
	t.Run("Skip(0) changes nothing", func(t *testing.T) {
		p := Build([]byte("abcd"), bytesOf([]string{"v", "w"}))
		p.Skip(0)
		verify(t, p, model{"abcd", []string{"v", "w"}})
	})
	t.Run("Skip removes the first bytes of the remainder and keeps the values", func(t *testing.T) {
		for k := 0; k <= 4; k++ {
			p := Build([]byte("abcd"), bytesOf([]string{"v", "w"}))
			p.Skip(k)
			verify(t, p, model{"abcd"[k:], []string{"v", "w"}})
		}
	})
	t.Run("Prepend puts bytes in front of the remainder in place", func(t *testing.T) {
		p := Build([]byte("cd"), bytesOf([]string{"v", "w"}))
		q := p.Prepend([]byte("ab"))
		if q != p {
			t.Fatal("the content fits, so the page stays")
		}
		verify(t, q, model{"abcd", []string{"v", "w"}})
	})
	t.Run("Prepend moves to a larger class when the content outgrows it", func(t *testing.T) {
		p := New([]byte("cd"), bytes.Repeat([]byte("x"), 24)) // 3 + 2 + 25 = 30
		q := p.Prepend([]byte("abcd"))
		if q == p || q.Size() != 64 {
			t.Fatalf("got a page of %d bytes", q.Size())
		}
		verify(t, q, model{"abcdcd", []string{strings.Repeat("x", 24)}})
		verify(t, p, model{"cd", []string{strings.Repeat("x", 24)}})
	})
	t.Run("Prepend refuses a remainder beyond 254 bytes", func(t *testing.T) {
		p := New(bytes.Repeat([]byte("a"), 200), []byte("v"))
		if p.Prepend(bytes.Repeat([]byte("b"), 55)) != nil {
			t.Fatal("255 bytes of remainder")
		}
		if p.Prepend(bytes.Repeat([]byte("b"), 54)) == nil {
			t.Fatal("254 bytes of remainder fit")
		}
	})
	t.Run("Prepend refuses content beyond 512 bytes", func(t *testing.T) {
		p := New([]byte("a"), bytes.Repeat([]byte("v"), 200))
		p, _ = p.Add(bytes.Repeat([]byte("w"), 200)) // 3 + 1 + 201 + 201 = 406
		if p.Prepend(bytes.Repeat([]byte("b"), 107)) != nil {
			t.Fatal("513 bytes")
		}
		if p.Prepend(bytes.Repeat([]byte("b"), 106)) == nil {
			t.Fatal("512 bytes fit")
		}
	})
}

// TestEqual covers the comparison tests use: two pages are equal when
// remainder and values in order agree, whatever their classes.
func TestEqual(t *testing.T) {
	a := Build([]byte("r"), bytesOf([]string{"x", "y"}))
	b := Build([]byte("r"), bytesOf([]string{"x", "y"}))
	b2, _ := b.Add(bytes.Repeat([]byte("z"), 40))
	b2, _ = b2.Remove(bytes.Repeat([]byte("z"), 40))
	cases := []struct {
		name string
		x, y *Page
		want bool
	}{
		{"same content", a, b, true},
		{"same content after growing and shrinking", a, b2, true},
		{"other order", a, Build([]byte("r"), bytesOf([]string{"y", "x"})), false},
		{"other remainder", a, Build([]byte("s"), bytesOf([]string{"x", "y"})), false},
		{"other number of values", a, New([]byte("r"), []byte("x")), false},
		{"remainder and value swap bytes", New([]byte("ab"), []byte("c")), New([]byte("a"), []byte("bc")), false},
	}
	for _, tt := range cases {
		if got := Equal(tt.x, tt.y); got != tt.want {
			t.Errorf("%s: Equal = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestKindBase covers a tree that numbers its object kinds: pages take the
// kinds from KindBase on, and a page still finds its class.
func TestKindBase(t *testing.T) {
	KindBase = 40
	defer func() { KindBase = 0 }()
	p := New([]byte("ab"), []byte("v"))
	q := p.Prepend(bytes.Repeat([]byte("x"), 100))
	if p.kind != 40 || q.kind != 40+2 || q.Size() != 128 {
		t.Fatalf("kinds %d and %d", p.kind, q.kind)
	}
	verify(t, q, model{strings.Repeat("x", 100) + "ab", []string{"v"}})
}

// TestAgainstModel makes sure that a page behaves like a list of different
// strings with a remainder, under many random operations, as a user of the
// tree would drive it: adding values of all lengths up to the limits,
// removing, moving deeper and up, with the page replaced whenever an
// operation returns another.
func TestAgainstModel(t *testing.T) {
	for seed := int64(1); seed <= 40; seed++ {
		rng := rand.New(rand.NewSource(seed))
		drive(t, rng, func(n int) []byte { return randBytes(rng, n) }, 600)
	}
}

// randBytes returns n random bytes from a small alphabet, so that equal values
// are likely.
func randBytes(rng *rand.Rand, n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = "abc"[rng.Intn(3)]
	}
	return b
}

// drive applies steps random operations to a page and its model.
func drive(t *testing.T, rng *rand.Rand, gen func(n int) []byte, steps int) {
	t.Helper()
	m := model{rest: string(gen(rng.Intn(6))), vals: []string{""}}
	p := New([]byte(m.rest), nil)
	var hist []string
	for range steps {
		var op string
		switch rng.Intn(10) {
		case 0, 1, 2, 3:
			v := gen(valueLen(rng))
			q, res := p.Add(v)
			op = fmt.Sprintf("add %q", v)
			switch res {
			case Added:
				m.vals = append(m.vals, string(v))
			case Present:
				if !slices.Contains(m.vals, string(v)) {
					t.Fatalf("%s: Present for a new value", op)
				}
			default:
				if len(v) <= MaxValue && modelUsed(m)+1+len(v) <= 512 {
					t.Fatalf("%s: Full although it fits", op)
				}
			}
			p = q
		case 4, 5, 6:
			v := m.vals[rng.Intn(len(m.vals))]
			if rng.Intn(5) == 0 {
				v = string(gen(valueLen(rng)))
			}
			op = fmt.Sprintf("remove %q", v)
			q, ok := p.Remove([]byte(v))
			if ok != slices.Contains(m.vals, v) {
				t.Fatalf("%s: ok = %v", op, ok)
			}
			if ok {
				i := slices.Index(m.vals, v)
				m.vals = slices.Delete(m.vals, i, i+1)
			}
			if q == nil { // the last value went: a key begins again
				m = model{rest: string(gen(rng.Intn(6))), vals: []string{""}}
				q = New([]byte(m.rest), nil)
			}
			p = q
		case 7:
			k := rng.Intn(len(m.rest) + 1)
			op = fmt.Sprintf("skip %d", k)
			p.Skip(k)
			m.rest = m.rest[k:]
		default:
			pre := gen(rng.Intn(20))
			op = fmt.Sprintf("prepend %q", pre)
			q := p.Prepend(pre)
			if q == nil {
				if len(pre)+len(m.rest) <= MaxRemainder && modelUsed(m)+len(pre) <= 512 {
					t.Fatalf("%s: nil although it fits", op)
				}
				break
			}
			p = q
			m.rest = string(pre) + m.rest
		}
		hist = append(hist, op)
		if err := check(p, m); err != nil {
			t.Fatalf("%v\nafter %d operations, the last ones: %q", err, len(hist), hist[max(0, len(hist)-6):])
		}
	}
}

// valueLen returns a random value length: mostly short, sometimes long.
func valueLen(rng *rand.Rand) int {
	switch rng.Intn(10) {
	case 0:
		return rng.Intn(260)
	case 1, 2:
		return rng.Intn(60)
	}
	return rng.Intn(8)
}

// modelUsed returns the content of the model in bytes.
func modelUsed(m model) int {
	n := Header + len(m.rest)
	for _, v := range m.vals {
		n += 1 + len(v)
	}
	return n
}

// FuzzPage lets the fuzzer drive a page with operations given as bytes and checks it
// against the model after each one.
func FuzzPage(f *testing.F) {
	f.Add([]byte{0, 3, 'a', 'b', 'c', 1, 2, 'a', 'b'})
	f.Fuzz(func(t *testing.T, data []byte) {
		rng := rand.New(rand.NewSource(int64(len(data))*7919 + int64(sum(data))))
		pos := 0
		gen := func(n int) []byte {
			b := make([]byte, n)
			for i := range b {
				if pos < len(data) {
					b[i] = data[pos]
					pos++
				} else {
					b[i] = byte(rng.Intn(256))
				}
			}
			return b
		}
		drive(t, rng, gen, 100)
	})
}

func sum(b []byte) int {
	s := 0
	for _, x := range b {
		s += int(x)
	}
	return s
}
