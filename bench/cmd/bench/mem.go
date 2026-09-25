package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"runtime/metrics"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/layout"
)

// memResult is what one candidate costs the rest of the program, measured in
// a process of its own. Heap figures are per key of the full corpus and
// exclude the corpus itself; the GC figure is raw and needs the "none"
// baseline subtracted.
type memResult struct {
	Impl           string  `json:"impl"`
	Keys           string  `json:"keys"`
	Values         string  `json:"values"`
	N              int     `json:"n"`
	LayoutSeed     uint64  `json:"layout_seed"`
	HeapPerKey     float64 `json:"heap_bytes_per_key"`
	ScanPerKey     float64 `json:"scannable_bytes_per_key"`
	GCCPUMs        float64 `json:"gc_cpu_ms_per_cycle"`
	HalfHeapPerKey float64 `json:"heap_bytes_per_key_after_removing_half"`
}

// runMem builds one candidate (or nothing, for impl "none"), measures its
// retained heap and the CPU time of a full GC cycle while it is alive, then
// removes every second key and measures the heap again. Memory is not an
// rtcompare comparison: GC cost depends on the whole live heap, so two
// candidates in one process would each pay for the other.
func runMem(kind keys.Kind, profile string, n int, impl string, cycles int, out io.Writer) error {
	c := keys.Generate(kind, n, 0x5EED)
	vals, offs := profileValues(profile, n)
	layout.Spacer()
	runtime.GC()
	before := heapStats()
	var keep any
	var remove func(i int)
	switch impl {
	case "none":
		remove = func(int) {}
	case ordered:
		m := buildOrdered(c.Keys.B, vals, offs)
		keep, remove = m, func(i int) { m.RemoveKey(c.Keys.B[i]) }
	case hashed:
		m := buildHashed(c.Keys.B, vals, offs)
		keep, remove = m, func(i int) { m.RemoveKey(c.Keys.B[i]) }
	case btreeSets:
		m := buildBtree(c.Keys.S, vals, offs)
		keep, remove = m, func(i int) { m.Delete(c.Keys.S[i]) }
	case mapSets:
		m := buildMap(c.Keys.S, vals, offs)
		keep, remove = m, func(i int) { delete(m, c.Keys.S[i]) }
	case btreeMapC:
		m := buildBtreeMap(c.Keys.S, vals, offs)
		keep, remove = m, func(i int) { m.Delete(c.Keys.S[i]) }
	default:
		return fmt.Errorf("unknown candidate %q", impl)
	}
	runtime.GC()
	full := heapStats()
	cpu := gcCPUPerCycle(cycles)
	for i := 0; i < n; i += 2 {
		remove(i)
	}
	runtime.GC()
	half := heapStats()
	runtime.KeepAlive(keep)
	runtime.KeepAlive(c)
	runtime.KeepAlive(vals)
	perKey := func(after, before uint64) float64 { return float64(int64(after-before)) / float64(n) }
	return json.NewEncoder(out).Encode(memResult{
		Impl: impl, Keys: string(kind), Values: profile, N: n, LayoutSeed: layout.Seed(),
		HeapPerKey: perKey(full.heap, before.heap), ScanPerKey: perKey(full.scan, before.scan),
		GCCPUMs: cpu, HalfHeapPerKey: perKey(half.heap, before.heap),
	})
}

type heap struct{ heap, scan uint64 }

func heapStats() heap {
	s := []metrics.Sample{{Name: "/gc/heap/live:bytes"}, {Name: "/gc/scan/heap:bytes"}}
	metrics.Read(s)
	return heap{heap: s[0].Value.Uint64(), scan: s[1].Value.Uint64()}
}

// gcCPUPerCycle forces cycles full collections and returns the GC CPU time
// per collection in milliseconds.
func gcCPUPerCycle(cycles int) float64 {
	read := func() float64 {
		s := []metrics.Sample{{Name: "/cpu/classes/gc/total:cpu-seconds"}}
		metrics.Read(s)
		return s[0].Value.Float64()
	}
	t0 := read()
	for range cycles {
		runtime.GC()
	}
	return (read() - t0) / float64(cycles) * 1000
}
