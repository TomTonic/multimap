// Command memgc measures what one structure costs the rest of the program:
// retained heap per key, how much of it the garbage collector has to scan, and
// the CPU time of a full GC cycle while the structure is alive.
//
// This is deliberately not an rtcompare comparison. GC cost is a property of
// the whole live heap, so two candidates in one process would each pay for the
// other. Each structure is measured in its own process instead (see run.sh),
// against a baseline process that holds only the key corpus.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/metrics"
	"slices"
	"time"

	art "github.com/plar/go-adaptive-radix-tree/v2"
	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/proto/arenaart"
	"github.com/TomTonic/multimap/bench/proto/arenaflat"
	"github.com/TomTonic/multimap/bench/proto/mmart"
	"github.com/TomTonic/multimap/bench/proto/mmart2"
	"github.com/TomTonic/multimap/bench/proto/mmbtree"
	"github.com/TomTonic/multimap/bench/proto/ptrart"
	"github.com/TomTonic/multimap/bench/proto/vset"
)

func main() {
	impl := flag.String("impl", "none", "none|arena-art|arena-flat|ptr-art|tidwall-btree|plar-art|go-map|mm-art|mm-art-v2|mm-btree-inline|mm-btree-ptr|vset-array|vset-hash")
	kind := flag.String("keys", "u64", "u64 or str")
	n := flag.Int("n", 1<<20, "number of keys")
	cycles := flag.Int("cycles", 15, "forced GC cycles to time")
	flag.Parse()

	c := keys.Generate(keys.Kind(*kind), *n, 0x5EED)
	vals, offs := keys.Values(*n, 0xFA11) // part of the baseline, like the corpus
	runtime.GC()
	before := heapStats()

	var keep any
	k := c.Keys.B
	switch *impl {
	case "none":
	case "arena-art":
		t := &arenaart.Tree{}
		for i, key := range k {
			t.Put(key, uint32(i))
		}
		keep = t
	case "arena-flat":
		t := &arenaflat.Tree{}
		for i, key := range k {
			t.Put(key, uint32(i))
		}
		keep = t
	case "ptr-art":
		t := &ptrart.Tree{}
		for i, key := range k {
			t.Put(key, uint32(i))
		}
		keep = t
	case "tidwall-btree":
		t := btree.NewMap[string, uint32](0)
		for i, key := range k {
			t.Set(string(key), uint32(i))
		}
		keep = t
	case "plar-art":
		t := art.New()
		for i, key := range k {
			t.Insert(art.Key(key), uint32(i))
		}
		keep = t
	case "go-map":
		t := map[string]uint32{}
		for i, key := range k {
			t[string(key)] = uint32(i)
		}
		keep = t
	case "mm-art":
		t := &mmart.Map[uint64]{}
		fillMM(t, k, vals, offs)
		keep = t
	case "mm-art-v2":
		t := &mmart2.Map[uint64]{}
		fillMM(t, k, vals, offs)
		keep = t
	case "mm-btree-inline":
		t := &mmbtree.Inline[uint64]{}
		fillMM(t, k, vals, offs)
		keep = t
	case "mm-btree-ptr":
		t := &mmbtree.Ptr[uint64]{}
		fillMM(t, k, vals, offs)
		keep = t
	case "vset-array": // the value containers alone, stored contiguously
		t := make([]vset.Set[uint64], *n)
		for i := range t {
			for _, v := range vals[offs[i]:offs[i+1]] {
				t[i].Add(v)
			}
		}
		keep = t
	case "vset-hash":
		t := make([]vset.HashSpill[uint64], *n)
		for i := range t {
			for _, v := range vals[offs[i]:offs[i+1]] {
				t[i].Add(v)
			}
		}
		keep = t
	default:
		fmt.Fprintln(os.Stderr, "unknown impl", *impl)
		os.Exit(2)
	}
	runtime.GC()
	after := heapStats()

	cpu0 := gcCPU()
	walls := make([]float64, *cycles)
	for i := range walls {
		t0 := time.Now()
		runtime.GC()
		walls[i] = float64(time.Since(t0).Microseconds()) / 1000
	}
	cpu := (gcCPU() - cpu0) / float64(*cycles) * 1000
	slices.Sort(walls)
	runtime.KeepAlive(keep)
	runtime.KeepAlive(c)
	runtime.KeepAlive(vals)

	fmt.Printf(`{"impl":%q,"keys":%q,"n":%d,"heap_bytes_per_key":%.1f,"scannable_bytes_per_key":%.1f,"gc_wall_ms_median":%.3f,"gc_cpu_ms_per_cycle":%.3f}`+"\n",
		*impl, *kind, *n,
		float64(int64(after.heap-before.heap))/float64(*n),
		float64(int64(after.scan-before.scan))/float64(*n),
		walls[len(walls)/2], cpu)
}

func fillMM(m interface{ AddValue([]byte, uint64) }, k [][]byte, vals []uint64, offs []int) {
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
}

type stats struct{ heap, scan uint64 }

func heapStats() stats {
	s := []metrics.Sample{{Name: "/gc/heap/live:bytes"}, {Name: "/gc/scan/heap:bytes"}}
	metrics.Read(s)
	return stats{heap: s[0].Value.Uint64(), scan: s[1].Value.Uint64()}
}

func gcCPU() float64 {
	s := []metrics.Sample{{Name: "/cpu/classes/gc/total:cpu-seconds"}}
	metrics.Read(s)
	return s[0].Value.Float64()
}
