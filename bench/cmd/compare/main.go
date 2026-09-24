// Command compare runs the rtcompare comparisons for one scenario (key kind and
// corpus size) and appends one JSON line per comparison to -out.
//
// Each scenario runs in its own process so that the live heap of one scenario
// never slows down the garbage collections of another. Candidate A is chosen
// with -a (default: the arena prototype); every other structure is candidate B
// in turn, so each line answers "how does A compare to X".
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	art "github.com/plar/go-adaptive-radix-tree/v2"
	hot "github.com/plar/go-hot-trie"
	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/proto/arenaart"
	"github.com/TomTonic/multimap/bench/proto/arenaflat"
	"github.com/TomTonic/multimap/bench/proto/ptrart"
	"github.com/TomTonic/rtcompare"
)

var sink uint64

const scanLen = 100

type result struct {
	Keys       string   `json:"keys"`
	N          int      `json:"n"`
	Op         string   `json:"op"`
	A          string   `json:"a"`
	B          string   `json:"b"`
	NsA        float64  `json:"ns_a"`
	NsB        float64  `json:"ns_b"`
	Delta      float64  `json:"delta"`
	Low        float64  `json:"low"`
	High       float64  `json:"high"`
	Resolved   bool     `json:"resolved"`
	NoiseFloor float64  `json:"noise_floor"`
	AutoCorr   float64  `json:"autocorr"`
	InnerLoops uint64   `json:"inner_loops"`
	Warnings   []string `json:"warnings"`
}

func main() {
	kind := flag.String("keys", "u64", "key kind: u64 or str")
	n := flag.Int("n", 4096, "number of keys")
	ops := flag.String("ops", "get,miss,scan", "comparisons to run: get,miss,scan,build")
	aName := flag.String("a", "arena-art", "candidate A")
	only := flag.String("only", "", "comma-separated B candidates to run (default all)")
	out := flag.String("out", "results.jsonl", "JSON lines output file")
	flag.Parse()

	c := keys.Generate(keys.Kind(*kind), *n, 0x5EED)
	f := build(c)
	f.verify()

	w, err := os.OpenFile(*out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fail(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "closing output:", err)
			os.Exit(1)
		}
	}()

	want := func(b string) bool { return *only == "" || slices.Contains(strings.Split(*only, ","), b) }
	for op := range strings.SplitSeq(*ops, ",") {
		for _, pair := range f.pairs(op, *aName) {
			if !want(pair[1].Name) {
				continue
			}
			// HOWTO: a tie rate of ~18% at the default showed up in the smoke test,
			// so batches are ten times longer than the default.
			opt := rtcompare.CompareOptions{Collect: rtcompare.CollectOptions{MaxQuantizationError: 0.0001}}
			if op == "build" {
				// building allocates; collect garbage between batches so one
				// candidate's garbage is not charged to the other
				opt.Collect.GCBetween = true
			}
			fmt.Fprintf(os.Stderr, "== %s n=%d %s: %s vs %s\n", *kind, *n, op, pair[0].Name, pair[1].Name)
			rep, err := rtcompare.Compare(pair[0], pair[1], opt)
			if err != nil {
				fail(err)
			}
			fmt.Fprintf(os.Stderr, "%s\n\n", rep)
			r := result{
				Keys: *kind, N: *n, Op: op, A: pair[0].Name, B: pair[1].Name,
				NsA: rep.NsPerOpA, NsB: rep.NsPerOpB,
				Delta: rep.Estimate.Delta, Low: rep.Estimate.Low, High: rep.Estimate.High,
				Resolved: rep.Resolved, NoiseFloor: rep.NoiseFloor, AutoCorr: rep.Autocorrelation,
				InnerLoops: rep.ValidationA.InnerLoops, Warnings: rep.Warnings,
			}
			if err := json.NewEncoder(w).Encode(r); err != nil {
				fail(err)
			}
		}
	}
	_ = sink
}

type fixtures struct {
	c     keys.Corpus
	arena *arenaart.Tree
	flat  *arenaflat.Tree
	ptr   *ptrart.Tree
	bt    *btree.Map[string, uint32]
	plar  art.Tree
	hot   hot.Tree
	m     map[string]uint32
}

// build inserts the corpus into every structure. All of them own their keys
// (copies), as a multimap must.
func build(c keys.Corpus) *fixtures {
	f := &fixtures{c: c, arena: &arenaart.Tree{}, flat: &arenaflat.Tree{}, ptr: &ptrart.Tree{},
		bt: btree.NewMap[string, uint32](0), plar: art.New(), hot: hot.New(), m: map[string]uint32{}}
	for i, k := range c.Keys.B {
		v := uint32(i)
		f.arena.Put(k, v)
		f.flat.Put(k, v)
		f.ptr.Put(k, v)
		f.bt.Set(string(k), v)
		f.plar.Insert(art.Key(k), v)
		f.hot.Insert(hot.Key(bytes.Clone(k)), v) // HOT keeps the caller's slice
		f.m[string(k)] = v
	}
	return f
}

// verify makes sure every structure answers identically before anything is
// timed, so a fast wrong answer cannot win.
func (f *fixtures) verify() {
	check := func(ok bool, what string, k []byte) {
		if !ok {
			fail(fmt.Errorf("fixture mismatch (%s) for key %q", what, k))
		}
	}
	for i, k := range f.c.Hits.B {
		want := f.m[string(k)]
		a, ok1 := f.arena.Get(k)
		fl, ok0 := f.flat.Get(k)
		p, ok2 := f.ptr.Get(k)
		b, ok3 := f.bt.Get(string(k))
		x, ok4 := f.plar.Search(art.Key(k))
		h, ok5 := f.hot.Search(hot.Key(k))
		check(ok0 && fl == want && ok1 && ok2 && ok3 && ok4 && ok5 && a == want && p == want && b == want && x.(uint32) == want && h.(uint32) == want, "hit", k)
		if i < 2000 {
			var sa, sf, sp, sb []uint32
			f.arena.Scan(k, func(_ []byte, v uint32) bool { sa = append(sa, v); return len(sa) < scanLen })
			f.flat.Scan(k, func(_ []byte, v uint32) bool { sf = append(sf, v); return len(sf) < scanLen })
			f.ptr.Scan(k, func(_ []byte, v uint32) bool { sp = append(sp, v); return len(sp) < scanLen })
			f.bt.Ascend(string(k), func(_ string, v uint32) bool { sb = append(sb, v); return len(sb) < scanLen })
			check(slices.Equal(sa, sb) && slices.Equal(sf, sb) && slices.Equal(sp, sb), "scan", k)
		}
	}
	for _, k := range f.c.Misses.B {
		_, ok0 := f.flat.Get(k)
		_, ok1 := f.arena.Get(k)
		_, ok2 := f.ptr.Get(k)
		_, ok3 := f.bt.Get(string(k))
		_, ok4 := f.plar.Search(art.Key(k))
		_, ok5 := f.m[string(k)]
		_, ok6 := f.hot.Search(hot.Key(k))
		check(!ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6, "miss", k)
	}
}

func (f *fixtures) pairs(op, a string) [][2]rtcompare.Candidate {
	var cands []rtcompare.Candidate
	switch op {
	case "get":
		cands = f.gets(f.c.Hits)
	case "miss":
		cands = f.gets(f.c.Misses)
	case "scan":
		cands = f.scans()
	case "build":
		cands = f.builds()
	default:
		fail(fmt.Errorf("unknown op %q", op))
	}
	i := slices.IndexFunc(cands, func(c rtcompare.Candidate) bool { return c.Name == a })
	if i < 0 {
		fail(fmt.Errorf("unknown candidate A %q", a))
	}
	var p [][2]rtcompare.Candidate
	for j, b := range cands {
		if j != i {
			p = append(p, [2]rtcompare.Candidate{cands[i], b})
		}
	}
	return p
}

// Every batch function keeps its own cursor across batches, so successive
// batches walk the whole probe sequence instead of re-probing a warm prefix.

func (f *fixtures) gets(p keys.Set) []rtcompare.Candidate {
	arena, flat, ptr, bt, pl, ht, m := f.arena, f.flat, f.ptr, f.bt, f.plar, f.hot, f.m
	var ja, jf, jp, jb, jl, jh, jm int
	next := func(j *int) {
		*j++
		if *j == len(p.B) {
			*j = 0
		}
	}
	return []rtcompare.Candidate{
		{Name: "arena-art", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := arena.Get(p.B[ja])
				acc += uint64(v) + b2u(ok)
				next(&ja)
			}
			sink += acc
		}},
		{Name: "arena-flat", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := flat.Get(p.B[jf])
				acc += uint64(v) + b2u(ok)
				next(&jf)
			}
			sink += acc
		}},
		{Name: "ptr-art", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := ptr.Get(p.B[jp])
				acc += uint64(v) + b2u(ok)
				next(&jp)
			}
			sink += acc
		}},
		{Name: "tidwall-btree", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := bt.Get(p.S[jb])
				acc += uint64(v) + b2u(ok)
				next(&jb)
			}
			sink += acc
		}},
		{Name: "plar-art", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := pl.Search(art.Key(p.B[jl]))
				if ok {
					acc += uint64(v.(uint32)) + 1
				}
				next(&jl)
			}
			sink += acc
		}},
		{Name: "plar-hot", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := ht.Search(hot.Key(p.B[jh]))
				if ok {
					acc += uint64(v.(uint32)) + 1
				}
				next(&jh)
			}
			sink += acc
		}},
		{Name: "go-map", Batch: func(n uint64) {
			var acc uint64
			for range n {
				v, ok := m[p.S[jm]]
				acc += uint64(v) + b2u(ok)
				next(&jm)
			}
			sink += acc
		}},
	}
}

func (f *fixtures) scans() []rtcompare.Candidate {
	p := f.c.Hits
	arena, flat, ptr, bt := f.arena, f.flat, f.ptr, f.bt
	var ja, jf, jp, jb int
	next := func(j *int) {
		*j++
		if *j == len(p.B) {
			*j = 0
		}
	}
	return []rtcompare.Candidate{
		{Name: "arena-art", Batch: func(n uint64) {
			var acc uint64
			for range n {
				cnt := 0
				arena.Scan(p.B[ja], func(_ []byte, v uint32) bool { acc += uint64(v); cnt++; return cnt < scanLen })
				next(&ja)
			}
			sink += acc
		}},
		{Name: "arena-flat", Batch: func(n uint64) {
			var acc uint64
			for range n {
				cnt := 0
				flat.Scan(p.B[jf], func(_ []byte, v uint32) bool { acc += uint64(v); cnt++; return cnt < scanLen })
				next(&jf)
			}
			sink += acc
		}},
		{Name: "ptr-art", Batch: func(n uint64) {
			var acc uint64
			for range n {
				cnt := 0
				ptr.Scan(p.B[jp], func(_ []byte, v uint32) bool { acc += uint64(v); cnt++; return cnt < scanLen })
				next(&jp)
			}
			sink += acc
		}},
		{Name: "tidwall-btree", Batch: func(n uint64) {
			var acc uint64
			for range n {
				cnt := 0
				bt.Ascend(p.S[jb], func(_ string, v uint32) bool { acc += uint64(v); cnt++; return cnt < scanLen })
				next(&jb)
			}
			sink += acc
		}},
	}
}

// builds measures constructing a whole structure from the corpus as one
// operation, in the corpus' random insertion order.
func (f *fixtures) builds() []rtcompare.Candidate {
	k := f.c.Keys.B
	return []rtcompare.Candidate{
		{Name: "arena-art", Batch: func(n uint64) {
			for range n {
				t := &arenaart.Tree{}
				for i, key := range k {
					t.Put(key, uint32(i))
				}
				sink += uint64(t.Len())
			}
		}},
		{Name: "arena-flat", Batch: func(n uint64) {
			for range n {
				t := &arenaflat.Tree{}
				for i, key := range k {
					t.Put(key, uint32(i))
				}
				sink += uint64(t.Len())
			}
		}},
		{Name: "ptr-art", Batch: func(n uint64) {
			for range n {
				t := &ptrart.Tree{}
				for i, key := range k {
					t.Put(key, uint32(i))
				}
				sink += uint64(t.Len())
			}
		}},
		{Name: "tidwall-btree", Batch: func(n uint64) {
			for range n {
				t := btree.NewMap[string, uint32](0)
				for i, key := range k {
					t.Set(string(key), uint32(i))
				}
				sink += uint64(t.Len())
			}
		}},
		{Name: "plar-art", Batch: func(n uint64) {
			for range n {
				t := art.New()
				for i, key := range k {
					t.Insert(art.Key(key), uint32(i))
				}
				sink += uint64(t.Size())
			}
		}},
		{Name: "plar-hot", Batch: func(n uint64) {
			for range n {
				t := hot.New()
				for i, key := range k {
					t.Insert(hot.Key(bytes.Clone(key)), uint32(i))
				}
				sink += uint64(t.Size())
			}
		}},
		{Name: "go-map", Batch: func(n uint64) {
			for range n {
				t := map[string]uint32{}
				for i, key := range k {
					t[string(key)] = uint32(i)
				}
				sink += uint64(len(t))
			}
		}},
	}
}

func b2u(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "compare:", err)
	os.Exit(1)
}
