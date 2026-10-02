package vpage

import (
	"bytes"
	"math/rand/v2"
	"os"
	"strconv"
	"testing"
)

// The benchmarks of the page prototype (docs/redesign, PLAN step 1) are for
// diagnosis: they say what a page costs to read and to change, and where its
// lookup waits for memory. They are not claims of speed; the tree's own
// benchmarks (cmd/bench, interleaved against a baseline) are.

func init() { // VPAGE_MINPREFIX=256 runs the benchmarks without the page prefix, for comparison
	if v := os.Getenv("VPAGE_MINPREFIX"); v != "" {
		MinPrefix, _ = strconv.Atoi(v)
	}
}

// workload is keys in pages, and the lookups to run against them.
type workload struct {
	pages []*Page
	keys  [][]byte // every key, in random order
	page  []int32  // keys[i] is in pages[page[i]]
	run   *Run
}

// build returns the pages of n keys of gen, inserted at random in one run.
func build(n int, gen func(*rand.Rand) []byte) *workload {
	r := rand.New(rand.NewPCG(5, 6))
	var run Run
	w := &workload{run: &run}
	seen := map[string]bool{}
	for len(w.keys) < n {
		k := gen(r)
		if seen[string(k)] {
			continue
		}
		seen[string(k)] = true
		_ = run.Insert(k, uint64(len(w.keys)))
		w.keys = append(w.keys, k)
	}
	w.pages = run.Pages()
	w.page = make([]int32, n)
	for i, k := range w.keys {
		w.page[i] = int32(run.page(k))
	}
	return w
}

var gens = []struct {
	name string
	gen  func(*rand.Rand) []byte
	n    int
}{
	{"u64", shapes["u64"], 3_000_000},
	{"uuid", shapes["uuid"], 1_000_000},
	{"path", shapes["path"], 1_000_000},
}

// leaf is what a tree of leaves holds per key: a 64-byte object with the key
// remainder and the value, found by a pointer from its node.
type leaf struct {
	key  [47]byte
	klen uint8
	val  uint64
	_    [8]byte
}

var sink uint64

// BenchmarkGet reads one value per iteration from a random page of a set far
// bigger than the CPU caches (cold), or from one page (hot). The leaf rows
// read a random 64-byte leaf of as many keys for comparison: one cache miss.
func BenchmarkGet(b *testing.B) {
	for _, g := range gens {
		w := build(g.n, g.gen)
		order := rand.New(rand.NewPCG(7, 8)).Perm(len(w.keys))
		b.Run(g.name+"/page cold", func(b *testing.B) {
			var acc uint64
			for i := 0; b.Loop(); i++ {
				j := order[i%len(order)]
				v, _ := w.pages[w.page[j]].Get(w.keys[j])
				acc += v
			}
			sink += acc
		})
		b.Run(g.name+"/page hot", func(b *testing.B) {
			p := w.pages[len(w.pages)/2]
			var buf [maxSuffix]byte
			ks := make([][]byte, p.Len())
			for i := range ks {
				ks[i] = bytes.Clone(p.Key(i, &buf))
			}
			var acc uint64
			for i := 0; b.Loop(); i++ {
				v, _ := p.Get(ks[i%len(ks)])
				acc += v
			}
			sink += acc
		})
		// keys that are not there and share no head word with a key that is
		var absent [][]byte
		var absentPage []int32
		for r := rand.New(rand.NewPCG(11, 12)); len(absent) < 200_000; {
			k := g.gen(r)
			if _, ok := w.run.Get(k); !ok {
				absent = append(absent, k)
				absentPage = append(absentPage, int32(w.run.page(k)))
			}
		}
		b.Run(g.name+"/page cold absent", func(b *testing.B) {
			var acc uint64
			for i := 0; b.Loop(); i++ {
				j := order[i%len(absent)] % len(absent)
				v, _ := w.pages[absentPage[j]].Get(absent[j])
				acc += v
			}
			sink += acc
		})
		b.Run(g.name+"/page cold miss", func(b *testing.B) {
			var acc uint64
			for i := 0; b.Loop(); i++ {
				j := order[i%len(order)]
				k := w.keys[j]
				miss := append(k[:len(k):len(k)], 0x7f) // another key, with the same head word
				v, _ := w.pages[w.page[j]].Get(miss)
				acc += v
			}
			sink += acc
		})
		leaves := make([]leaf, len(w.keys))
		for i, k := range w.keys {
			leaves[i].klen = uint8(copy(leaves[i].key[:], k))
			leaves[i].val = uint64(i)
		}
		b.Run(g.name+"/leaf cold", func(b *testing.B) {
			var acc uint64
			for i := 0; b.Loop(); i++ {
				j := order[i%len(order)]
				l := &leaves[j]
				if bytes.Equal(l.key[:l.klen], w.keys[j]) {
					acc += l.val
				}
			}
			sink += acc
		})
	}
}

// BenchmarkMutate times what changes a page: an insert into free room, an
// insert that rebuilds the page, a delete, and the split of a full page.
func BenchmarkMutate(b *testing.B) {
	r := rand.New(rand.NewPCG(9, 10))
	for _, g := range gens {
		keys := make([][]byte, 4096)
		for i := range keys {
			keys[i] = g.gen(r)
		}
		b.Run(g.name+"/insert and delete in a page with room", func(b *testing.B) {
			p := New(2, 0)
			for i := 0; i < 8; i++ { // a page that neither grows nor shrinks by one key
				p, _, _ = p.Insert(keys[i], 1)
			}
			for i := 8; b.Loop(); i++ {
				k := keys[8+i%(len(keys)-8)]
				q, res, _ := p.Insert(k, 1)
				if res == Inserted {
					p, _ = q.Delete(k)
				}
			}
		})
		b.Run(g.name+"/fill, grow and split", func(b *testing.B) {
			for i := 0; b.Loop(); i++ {
				p := New(0, 0)
				for j := 0; ; j++ {
					q, res, _ := p.Insert(keys[(i*31+j)%len(keys)], 1)
					if res == Full {
						_, _, _ = p.Split()
						break
					}
					p = q
				}
			}
		})
		b.Run(g.name+"/delete and insert back", func(b *testing.B) {
			p := New(2, 0)
			var in [][]byte
			for _, k := range keys {
				q, res, _ := p.Insert(k, 1)
				if res == Full {
					break
				}
				p = q
				in = append(in, k)
			}
			for i := 0; b.Loop(); i++ {
				k := in[i%len(in)]
				q, _ := p.Delete(k)
				p, _, _ = q.Insert(k, 1)
			}
		})
	}
}

// getTrace walks the lookup of s in page p as Get does and returns the 64-byte
// lines of the page it reads, and the number of rounds of loads where one has to
// wait for the one before: the directory, then the head, the value and the tail
// of the key the tags point at (their addresses need only the directory entry,
// and the CPU speculates on the match). It mirrors Get (a test checks that the
// results agree); it must be kept in step with it.
func getTrace(p *Page, s []byte) (lines map[int]bool, rounds int, v uint64, ok bool) {
	lines = map[int]bool{}
	read := func(off, n int) {
		for l := off / 64; l <= (off+n-1)/64; l++ {
			lines[l] = true
		}
	}
	n, c, u := int(p.count), int(p.cap), p.ulen != 0
	stride := 4
	if u {
		stride = 1
	}
	pl := int(p.plen)
	read(0, dirAt(pl)+8*((stride*n+7)/8)) // the header, the prefix and the directory
	rounds = 1
	if len(s) > maxSuffix {
		return lines, rounds, 0, false
	}
	s, rel := stripPrefix(p, s)
	if rel != 0 || (u && len(s) != int(p.ulen)) {
		return lines, rounds, 0, false
	}
	w, t := word(s), tag(word(s))
	for i := range n {
		if p.tag(i) != t {
			continue
		}
		rounds = 2
		read(base(c, u, pl)+16*i, 16) // the head and the value, side by side
		if !u && p.length(i) > headLen {
			read(int(p.fat()[i]>>16), p.length(i)-headLen)
		}
		if p.slots()[i].Head != w || (!u && p.length(i) != len(s)) {
			continue
		}
		if !u && p.length(i) > headLen && !bytes.Equal(p.tail(i), s[headLen:]) {
			continue
		}
		return lines, rounds, p.slots()[i].Val, true
	}
	return lines, rounds, 0, false
}

// TestLinesPerGet makes sure that the trace of a lookup agrees with the lookup
// and reports, for each shape of key, how many cache lines and how many rounds
// of dependent loads a lookup of a present and an absent key takes in pages of
// each class. It belongs to the page prototype of the redesign (docs/redesign,
// step 1), whose rule R5 asks for a decision in the first line or two of a page
// and the payload in one line more. Run with -v to see the figures.
func TestLinesPerGet(t *testing.T) {
	for _, g := range []string{"u64", "uuid", "path", "short-mixed"} {
		n := 20000
		if g == "short-mixed" {
			n = 3000 // there are only 9841 distinct keys of up to 8 bytes over 3 symbols
		}
		w := build(n, shapes[g])
		type sum struct{ n, lines, rounds int }
		hit, miss := map[int]*sum{}, map[int]*sum{}
		for i, k := range w.keys {
			p := w.pages[w.page[i]]
			lines, rounds, v, ok := getTrace(p, k)
			want, wantOK := p.Get(k)
			if ok != wantOK || v != want || !ok {
				t.Fatalf("%s: trace of %x says %v %d, Get %v %d", g, k, ok, v, wantOK, want)
			}
			m := hit[p.Class()]
			if m == nil {
				m = &sum{}
				hit[p.Class()] = m
			}
			m.n, m.lines, m.rounds = m.n+1, m.lines+len(lines), m.rounds+rounds
			lines, rounds, _, ok = getTrace(p, append(bytes.Clone(k), 0x7f))
			if ok {
				t.Fatalf("%s: an absent key was found", g)
			}
			m = miss[p.Class()]
			if m == nil {
				m = &sum{}
				miss[p.Class()] = m
			}
			m.n, m.lines, m.rounds = m.n+1, m.lines+len(lines), m.rounds+rounds
		}
		for c := range 4 {
			if h := hit[c]; h != nil {
				ms := miss[c]
				t.Logf("%-11s class %d (%4d B): lookup of a key: %.1f lines, %.1f rounds; of an absent key: %.1f lines, %.1f rounds (%d lookups)",
					g, c, sizes[c], float64(h.lines)/float64(h.n), float64(h.rounds)/float64(h.n), float64(ms.lines)/float64(ms.n), float64(ms.rounds)/float64(ms.n), h.n)
				if float64(h.rounds)/float64(h.n) > 5 {
					t.Errorf("%s: more than 5 rounds per lookup", g)
				}
			}
		}
	}
}

// stripPrefix returns suffix s without the page's prefix and 0, or nil and -1 if
// s sorts below all keys of the page, +1 if above (it does not start with the
// prefix).
func stripPrefix(p *Page, s []byte) ([]byte, int) {
	n := int(p.plen)
	if n == 0 {
		return s, 0
	}
	pre := p.prefix()
	if len(s) >= n && string(s[:n]) == string(pre) {
		return s[n:], 0
	}
	if compare(s[:min(len(s), n)], pre) < 0 {
		return nil, -1
	}
	return nil, 1
}
