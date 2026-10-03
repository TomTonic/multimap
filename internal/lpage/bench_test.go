package lpage

import (
	"math/rand/v2"
	"os"
	"strconv"
	"testing"
)

// The benchmarks of the length-header page are for diagnosis, as those of
// internal/vpage are, and are written to run against them with the same keys:
// they say what a page costs to read and to change, not how fast the tree is
// (cmd/bench, interleaved against a baseline, does that).

func init() { // LPAGE_MAXHEADER=16 runs the benchmarks with pages of at most 14 entries
	if v := os.Getenv("LPAGE_MAXHEADER"); v != "" {
		MaxHeader, _ = strconv.Atoi(v)
	}
}

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
		if len(k) == 0 || seen[string(k)] {
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

var sink uint64

// BenchmarkGet reads one value per iteration from a random page of a set far
// bigger than the CPU caches (cold), or from one page (hot); the absent rows
// look for keys that are not there.
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
			ks, _ := p.Entries()
			var acc uint64
			for i := 0; b.Loop(); i++ {
				v, _ := p.Get(ks[i%len(ks)])
				acc += v
			}
			sink += acc
		})
		var absent [][]byte
		var absentPage []int32
		for r := rand.New(rand.NewPCG(11, 12)); len(absent) < 200_000; {
			k := g.gen(r)
			if _, ok := w.run.Get(k); !ok && len(k) > 0 {
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
	}
}

// BenchmarkMutate times what changes a page: an insert into free room followed
// by the delete of the same key, filling a page until it must split, and
// deleting a key and putting it back.
func BenchmarkMutate(b *testing.B) {
	r := rand.New(rand.NewPCG(9, 10))
	for _, g := range gens {
		keys := make([][]byte, 4096)
		for i := range keys {
			keys[i] = g.gen(r)
			for len(keys[i]) == 0 {
				keys[i] = g.gen(r)
			}
		}
		b.Run(g.name+"/insert and delete in a page with room", func(b *testing.B) {
			p, _ := Build(keys[:1], []uint64{1})
			for i := 1; i < 8; i++ { // a page that neither grows nor shrinks by one key
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
				p, _ := Build(keys[(i*31)%len(keys):(i*31)%len(keys)+1], []uint64{1})
				for j := 1; ; j++ {
					q, res, _ := p.Insert(keys[(i*31+j)%len(keys)], 1)
					if res == Full {
						_, _ = p.Split()
						break
					}
					p = q
				}
			}
		})
		b.Run(g.name+"/delete and insert back", func(b *testing.B) {
			p, _ := Build(keys[:1], []uint64{1})
			in := [][]byte{keys[0]}
			for _, k := range keys[1:] {
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
