// Command ovbench is a diagnosis for the design of the value overflow
// (docs/redesign/step3-overflow-design.md, section 7): on the keys of the real
// data sets whose values do not fit a page of 512 bytes, it builds the value
// set of each key in two ways, as internal/vset does today (inline, array, hash
// set) and as a plain Set3 (the idea of leaving vset out), and times and weighs
// them. The strings are allocated one by one, as an application's would be.
//
//	go run ./cmd/ovbench [-keys street,dirs] [-rounds 5]
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

// keyVals is the values of one key that overflows a page.
type keyVals struct{ vals []string }

func main() {
	kinds := flag.String("keys", "street,dirs", "key kinds with natural values")
	rounds := flag.Int("rounds", 5, "rounds per measurement; the median is reported")
	flag.Parse()
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
			b += 1 + len(name(c, v))
		}
		if 6+8+b <= 512 {
			continue
		}
		kv := keyVals{}
		for _, v := range e.vals {
			kv.vals = append(kv.vals, strings.Clone(name(c, v)))
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

// container is what the two candidates have in common.
type container interface {
	Add(string) bool
	Remove(string) bool
	Contains(string) bool
	Each(func(string) bool)
}

type vsetC struct{ s vset.Set[string] }

func (c *vsetC) Add(v string) bool      { return c.s.Add(v) }
func (c *vsetC) Remove(v string) bool   { return c.s.Remove(v) }
func (c *vsetC) Contains(v string) bool { return c.s.Contains(v) }
func (c *vsetC) Each(f func(string) bool) {
	c.s.Each(f)
}

type set3C struct{ s *set3.Set3[string] }

func (c *set3C) Add(v string) bool      { n := c.s.Size(); c.s.Add(v); return c.s.Size() != n }
func (c *set3C) Remove(v string) bool   { return c.s.Remove(v) }
func (c *set3C) Contains(v string) bool { return c.s.Contains(v) }
func (c *set3C) Each(f func(string) bool) {
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
	for _, cand := range []struct {
		name string
		mk   func(n int) container
	}{
		{"vset (today)", func(int) container { return &vsetC{} }},
		{"Set3 on its own", func(int) container { return &set3C{set3.Empty[string]()} }},
	} {
		var before runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&before)
		sets := make([]container, len(data))
		for i, kv := range data {
			sets[i] = cand.mk(len(kv.vals))
			for _, v := range kv.vals {
				sets[i].Add(v)
			}
		}
		runtime.GC()
		var after runtime.MemStats
		runtime.ReadMemStats(&after)
		// the strings are shared by both and counted in neither run's difference: add them
		strBytes := 0
		for _, kv := range data {
			for _, v := range kv.vals {
				strBytes += max(16, (len(v)+7)/8*8) // a rough Go size class for a string of its own
			}
		}
		perValue := float64(int64(after.HeapAlloc)-int64(before.HeapAlloc))/float64(total) + float64(strBytes)/float64(total)
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
				kv := &data[j]
				if !sets[j].Contains(kv.vals[rng.Intn(len(kv.vals))]) {
					b.Fatal("missing")
				}
			}
		})
		miss := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(3))
			probe := "no such value"
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
				s.Add("a new value")
				s.Remove("a new value")
			}
		})
		each := med(func(b *testing.B) {
			rng := rand.New(rand.NewSource(5))
			n := 0
			for i := 0; i < b.N; {
				s := sets[rng.Intn(len(sets))]
				s.Each(func(v string) bool { sink += int(v[0]); n++; return true })
				i = n
			}
		})
		fmt.Printf("| %s | %s | %.1f | %.0f | %.0f | %.0f | %.1f |\n", kind, cand.name, perValue, hit, miss, churn, each)
	}
	return nil
}
