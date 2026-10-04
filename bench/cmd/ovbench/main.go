// Command ovbench is a diagnosis for the design of the value overflow
// (docs/redesign/step3-overflow-design.md, section 7): on the keys of the real
// data sets whose values do not fit a page of 512 bytes, it builds the value
// set of each key in two ways, as internal/vset does today (inline, array, hash
// set) and as a plain Set3 (the idea of leaving vset out), and times and weighs
// them. The strings are allocated one by one, as an application's would be.
//
//	go run ./cmd/ovbench [-keys street,dirs] [-rounds 5] [-values string|words|pointers]
//
// -values words and pointers (8-byte values; step 3.5, docs/redesign/step3-fixed-design.md)
// build the same sets of uint64 or of pointers to records the caller owns; a key
// then overflows at more than (512 - 6 - remainder) / 8 values.
//
// It is a diagnosis of the containers in isolation, not of the tree.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"slices"
	"sort"
	"strings"
	"testing"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/vset"
)

var sink int

// recs are the records of the mode pointers, one per value number, allocated one by one.
var recs map[uint64]*rec

// valuesMode is the flag -values.
var valuesMode string

// rec is the caller's record of the mode pointers.
type rec struct{ id, aux uint64 }

// keyVals is the values of one key that overflows a page.
type keyVals struct {
	vals  []string
	words []uint64
	ptrs  []*rec
}

func main() {
	kinds := flag.String("keys", "street,dirs", "key kinds with natural values")
	rounds := flag.Int("rounds", 5, "rounds per measurement; the median is reported")
	values := flag.String("values", "string", "the values: string, words (uint64) or pointers (*rec)")
	flag.Parse()
	valuesMode = *values
	fmt.Println("| data | container | heap B/value (string bytes included) | contains, hit (ns) | contains, miss (ns) | add+remove (ns) | read all values of a key (ns/value) |\n|---|---|--:|--:|--:|--:|--:|")
	for _, k := range strings.Split(*kinds, ",") {
		if err := run(keys.Kind(k), *rounds); err != nil {
			fmt.Fprintln(os.Stderr, "ovbench:", err)
			os.Exit(1)
		}
	}
}

// load returns the keys whose content does not fit a page of 512 bytes, with
// their values, each string allocated on its own.
func load(kind keys.Kind) ([]keyVals, int) {
	recs = map[uint64]*rec{}
	c := keys.Generate(kind, keys.Capacity(kind), 0x5EED)
	all := make([]struct {
		key  []byte
		vals []uint64
	}, len(c.Keys.B))
	for i, k := range c.Keys.B {
		all[i].key, all[i].vals = k, c.Natural[i]
	}
	sort.Slice(all, func(i, j int) bool { return bytes.Compare(all[i].key, all[j].key) < 0 })
	var out []keyVals
	total := 0
	for _, e := range all {
		n, b := len(e.vals), 0
		for _, v := range e.vals {
			if valuesMode == "string" {
				b += 1 + len(name(c, v))
			} else {
				b += 8
			}
		}
		if 6+8+b <= 512 {
			continue
		}
		kv := keyVals{}
		for _, v := range e.vals {
			switch valuesMode {
			case "string":
				kv.vals = append(kv.vals, strings.Clone(name(c, v)))
			case "words":
				kv.words = append(kv.words, v)
			default:
				kv.ptrs = append(kv.ptrs, recOf(v))
			}
		}
		out = append(out, kv)
		total += n
	}
	rand.New(rand.NewSource(1)).Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
	return out, total
}

func name(c keys.Corpus, v uint64) string {
	if v >= 1 && v <= uint64(len(c.Names)) {
		return c.Names[v-1]
	}
	return fmt.Sprintf("%016x", v)
}

// recOf returns the record of value number v, made on first use.
func recOf(v uint64) *rec {
	r, ok := recs[v]
	if !ok {
		r = &rec{id: v, aux: v * 0x9e3779b97f4a7c15}
		recs[v] = r
	}
	return r
}

// container is what the two candidates have in common.
type container[V comparable] interface {
	Add(V) bool
	Remove(V) bool
	Contains(V) bool
	Each(func(V) bool)
}

type vsetC[V comparable] struct{ s vset.Set[V] }

func (c *vsetC[V]) Add(v V) bool      { return c.s.Add(v) }
func (c *vsetC[V]) Remove(v V) bool   { return c.s.Remove(v) }
func (c *vsetC[V]) Contains(v V) bool { return c.s.Contains(v) }
func (c *vsetC[V]) Each(f func(V) bool) {
	c.s.Each(f)
}

type set3C[V comparable] struct{ s *set3.Set3[V] }

func (c *set3C[V]) Add(v V) bool      { n := c.s.Size(); c.s.Add(v); return c.s.Size() != n }
func (c *set3C[V]) Remove(v V) bool   { return c.s.Remove(v) }
func (c *set3C[V]) Contains(v V) bool { return c.s.Contains(v) }
func (c *set3C[V]) Each(f func(V) bool) {
	for v := range c.s.MutableRange() {
		if !f(v) {
			return
		}
	}
}

func run(kind keys.Kind, rounds int) error {
	data, total := load(kind)
	if len(data) == 0 {
		return fmt.Errorf("%s: no overflow keys", kind)
	}
	fmt.Fprintf(os.Stderr, "ovbench: %s: %d keys over a page, %d values\n", kind, len(data), total)
	switch valuesMode {
	case "string":
		measure(kind, rounds, data, total, func(kv keyVals) []string { return kv.vals }, "no such value", "a new value", func(v string) int { return int(v[0]) }, true)
	case "words":
		measure(kind, rounds, data, total, func(kv keyVals) []uint64 { return kv.words }, 1<<62, 1<<62+1, func(v uint64) int { return int(v) }, false)
	default:
		measure(kind, rounds, data, total, func(kv keyVals) []*rec { return kv.ptrs }, &rec{id: 1 << 62}, &rec{id: 1<<62 + 1}, func(v *rec) int { return int(v.id) }, false)
	}
	return nil
}

// measure weighs and times the two containers on the sets of data. strBytes
// tells that the values are strings of their own, whose bytes the heap figure
// then includes.
func measure[V comparable](kind keys.Kind, rounds int, data []keyVals, total int, vals func(keyVals) []V, probe, fresh V, weigh func(V) int, strBytes bool) {
	for _, cand := range []struct {
		name string
		mk   func() container[V]
	}{
		{"vset (today)", func() container[V] { return &vsetC[V]{} }},
		{"Set3 on its own", func() container[V] { return &set3C[V]{set3.Empty[V]()} }},
	} {
		var before runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		sets := make([]container[V], len(data))
		for i, kv := range data {
			sets[i] = cand.mk()
			for _, v := range vals(kv) {
				sets[i].Add(v)
			}
		}
		runtime.GC()
		var after runtime.MemStats
		runtime.ReadMemStats(&after)
		perValue := float64(int64(after.HeapAlloc)-int64(before.HeapAlloc)) / float64(total)
		if strBytes { // the strings are shared by both and counted in neither run's difference: add them
			b := 0
			for _, kv := range data {
				for _, v := range kv.vals {
					b += max(16, (len(v)+7)/8*8) // a rough Go size class for a string of its own
				}
			}
			perValue += float64(b) / float64(total)
		}
		med := func(f func(b *testing.B)) float64 {
			ns := make([]float64, rounds)
			for r := range rounds {
				res := testing.Benchmark(f)
				ns[r] = float64(res.T.Nanoseconds()) / float64(res.N)
			}
			slices.Sort(ns)
			return ns[len(ns)/2]
		}
		hit := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(2))
			for i := 0; i < b.N; i++ {
				j := rng.Intn(len(data))
				vs := vals(data[j])
				if !sets[j].Contains(vs[rng.Intn(len(vs))]) {
					b.Fatal("missing")
				}
			}
		})
		miss := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(3))
			for i := 0; i < b.N; i++ {
				if sets[rng.Intn(len(sets))].Contains(probe) {
					b.Fatal("found")
				}
			}
		})
		churn := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(4))
			for i := 0; i < b.N; i++ {
				s := sets[rng.Intn(len(sets))]
				s.Add(fresh)
				s.Remove(fresh)
			}
		})
		each := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(5))
			n := 0
			for i := 0; i < b.N; {
				s := sets[rng.Intn(len(sets))]
				s.Each(func(v V) bool { sink += weigh(v); n++; return true })
				i = n
			}
		})
		fmt.Printf("| %s | %s | %.1f | %.0f | %.0f | %.0f | %.1f |\n", kind, cand.name, perValue, hit, miss, churn, each)
	}
}
