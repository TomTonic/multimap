// Command skbench is the diagnosis of the single-key page (internal/skpage,
// docs/redesign/step3-skmv-design.md) on the real data sets: one page per key,
// with the remainder a tree of byte nodes would cut and all the key's real
// values, and the page's own operations timed on them, with the choice of the
// page taken out of the timed part. It is a diagnosis of the page in
// isolation and not a claim of speed; the tree's own benchmarks (cmd/bench)
// are.
//
//	go run ./cmd/skbench [-keys street,dirs] [-rounds 5]
//
// Operations (ns per operation, the median of the rounds; allocations per
// operation):
//
//   - lookup, hot: Match and Strings on pages in random order out of the first
//     4,096, which stay in the cache;
//   - lookup, all: the same on all pages, which do not;
//   - add+remove: add a value of 10 bytes to a page and remove it again, on all
//     pages in random order (the churn of values; where the content crosses a
//     size class the page is copied twice);
//   - add+remove at a class border: the same on a page whose content fills its class exactly
//     (32 bytes), so that every add changes the class and
//     every remove changes it back: the worst case of the page;
//   - build: New and then Add of the other values, for every key.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/skpage"
)

var sink int

// entry is one key as the tree holds it: the remainder below the byte node and its values.
type entry struct {
	key  []byte // the whole key
	rest []byte
	vals [][]byte
}

func main() {
	kinds := flag.String("keys", "street,dirs", "key kinds with natural values")
	rounds := flag.Int("rounds", 5, "rounds per measurement; the median is reported")
	flag.Parse()
	fmt.Println("| data | operation | ns per operation | allocations per operation |\n|---|---|--:|--:|")
	for _, k := range strings.Split(*kinds, ",") {
		if err := run(keys.Kind(k), *rounds); err != nil {
			fmt.Fprintln(os.Stderr, "skbench:", err)
			os.Exit(1)
		}
	}
}

// load returns the entries of the largest corpus of a data set, in a random
// order, whose page fits: the others are the value overflow's.
func load(kind keys.Kind) ([]entry, error) {
	c := keys.Generate(kind, keys.Capacity(kind), 0x5EED)
	if c.Natural == nil {
		return nil, fmt.Errorf("%s has no natural values", kind)
	}
	type kv struct {
		key  []byte
		vals []uint64
	}
	all := make([]kv, len(c.Keys.B))
	for i, k := range c.Keys.B {
		all[i] = kv{k, c.Natural[i]}
	}
	sort.Slice(all, func(i, j int) bool { return bytes.Compare(all[i].key, all[j].key) < 0 })
	var es []entry
	for i, e := range all {
		l := 0
		if i > 0 {
			l = max(l, lcp(e.key, all[i-1].key))
		}
		if i+1 < len(all) {
			l = max(l, lcp(e.key, all[i+1].key))
		}
		en := entry{key: e.key, rest: e.key[min(l+1, len(e.key)):]}
		for _, v := range e.vals {
			en.vals = append(en.vals, []byte(name(c, v)))
		}
		if skpage.Build(en.rest, len(e.key), en.vals) != nil {
			es = append(es, en)
		}
	}
	rand.New(rand.NewSource(1)).Shuffle(len(es), func(i, j int) { es[i], es[j] = es[j], es[i] })
	return es, nil
}

func name(c keys.Corpus, v uint64) string {
	if v >= 1 && v <= uint64(len(c.Names)) {
		return c.Names[v-1]
	}
	return fmt.Sprintf("%016x", v)
}

func lcp(a, b []byte) int {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i
}

func run(kind keys.Kind, rounds int) error {
	es, err := load(kind)
	if err != nil {
		return err
	}
	pages := make([]*skpage.Page, len(es))
	for i, e := range es {
		pages[i] = skpage.Build(e.rest, len(e.key), e.vals)
	}
	hot := min(4096, len(pages))
	row := func(op string, f func(b *testing.B)) {
		ns := make([]float64, rounds)
		var allocs float64
		for r := range rounds {
			res := testing.Benchmark(f)
			ns[r] = float64(res.T.Nanoseconds()) / float64(res.N)
			allocs = float64(res.AllocsPerOp())
		}
		slices.Sort(ns)
		fmt.Printf("| %s | %s | %.1f | %.2f |\n", kind, op, ns[len(ns)/2], allocs)
	}
	lookup := func(n int) func(b *testing.B) {
		return func(b *testing.B) {
			b.ReportAllocs()
			rng := rand.New(rand.NewSource(2))
			perm := rng.Perm(n)
			for i := 0; i < b.N; i++ {
				j := perm[i%n]
				if !pages[j].Match(es[j].key) {
					b.Fatal("no match")
				}
				pages[j].Strings(func(v string) bool { sink += len(v); return true })
			}
		}
	}
	row("lookup, hot (4,096 pages)", lookup(hot))
	row("lookup, all pages", lookup(len(pages)))
	val := []byte("0123456789")
	row("add+remove a value, all pages", func(b *testing.B) {
		b.ReportAllocs()
		own := slices.Clone(pages)
		perm := rand.New(rand.NewSource(3)).Perm(len(own))
		for i := 0; i < b.N; i++ {
			j := perm[i%len(own)]
			q, res := own[j].Add(val)
			if res == skpage.Added {
				q, _ = q.Remove(val)
			}
			own[j] = q
		}
	})
	row("add+remove at a class border", func(b *testing.B) {
		b.ReportAllocs()
		p := skpage.New([]byte("ab"), 10, bytes.Repeat([]byte("x"), 23)) // content 6 + 2 + 1 + 23 = 32: the 32-byte class is full
		for i := 0; i < b.N; i++ {
			q, _ := p.Add(val)
			p, _ = q.Remove(val)
		}
	})
	row("build, per key", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			e := &es[i%len(es)]
			p := skpage.New(e.rest, len(e.key), e.vals[0])
			for _, v := range e.vals[1:] {
				var res skpage.Result
				if p, res = p.Add(v); res == skpage.Full {
					break
				}
			}
			sink += p.Len()
		}
	})
	return nil
}
