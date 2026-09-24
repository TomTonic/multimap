// Command vsetcompare compares the three-representation value container
// (vset.Set: inline, array, hash) with the first design (vset.HashSpill:
// inline, then straight to a set3 hash set), using rtcompare.
//
// n containers are filled with the skewed value-set sizes of keys.Values and
// stored contiguously, as the ART leaves and B-tree items hold them inline.
// Probes pick a random container (and a random present value in it).
//
//	contains   Contains of a present value
//	each       iterate all values of the container
//	addRemove  Add then Remove of an absent value
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/proto/vset"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/rtcompare"
)

var sink uint64

const numProbes = 1 << 20

func main() {
	n := flag.Int("n", 1<<20, "number of containers (keys)")
	ops := flag.String("ops", "contains,each,addRemove", "operations")
	out := flag.String("out", "results/vset.jsonl", "JSON lines output")
	flag.Parse()

	vals, offs := keys.Values(*n, 0xFA11)
	news := make([]vset.Set[uint64], *n)
	olds := make([]vset.HashSpill[uint64], *n)
	for i := range *n {
		for _, v := range vals[offs[i]:offs[i+1]] {
			news[i].Add(v)
			olds[i].Add(v)
		}
	}
	rng := rtcompare.NewDPRNG(0xC0FFEE)
	idx := make([]uint32, numProbes)
	val := make([]uint64, numProbes)
	for j := range idx {
		i := int(rng.Uint64() % uint64(*n))
		idx[j] = uint32(i)
		val[j] = vals[offs[i]+int(rng.Uint64()%uint64(offs[i+1]-offs[i]))]
	}
	for j := range idx { // the addRemove state both start from: settle hysteresis
		i := idx[j]
		news[i].Add(1<<63 | uint64(j))
		news[i].Remove(1<<63 | uint64(j))
		olds[i].Add(1<<63 | uint64(j))
		olds[i].Remove(1<<63 | uint64(j))
	}

	w, err := os.OpenFile(*out, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "closing output:", err)
			os.Exit(1)
		}
	}()
	for op := range strings.SplitSeq(*ops, ",") {
		var a, b rtcompare.Candidate
		switch op {
		case "contains":
			a, b = containsNew(news, idx, val), containsOld(olds, idx, val)
		case "each":
			a, b = eachNew(news, idx), eachOld(olds, idx)
		case "addRemove":
			a, b = addRemoveNew(news, idx), addRemoveOld(olds, idx)
		default:
			panic("unknown op " + op)
		}
		fmt.Fprintf(os.Stderr, "== n=%d %s: %s vs %s\n", *n, op, a.Name, b.Name)
		rep, err := rtcompare.Compare(a, b, rtopt.Options(a, b, false))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "%s\n\n", rep)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"n": *n, "op": op, "a": a.Name, "b": b.Name, "ns_a": rep.NsPerOpA, "ns_b": rep.NsPerOpB,
			"delta": rep.Estimate.Delta, "low": rep.Estimate.Low, "high": rep.Estimate.High,
			"resolved": rep.Resolved, "noise_floor": rep.NoiseFloor, "warnings": rep.Warnings,
		})
	}
	_ = sink
}

// One batch function per container type and operation, so each loop calls
// the concrete methods directly. The old container's cursor starts half-way
// so the two never re-probe the containers the other just pulled into cache.

func containsNew(s []vset.Set[uint64], idx []uint32, val []uint64) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: "array-spill", Batch: func(n uint64) {
		var acc uint64
		for range n {
			if s[idx[j]].Contains(val[j]) {
				acc++
			}
			j = (j + 1) & (numProbes - 1)
		}
		sink += acc
	}}
}

func containsOld(s []vset.HashSpill[uint64], idx []uint32, val []uint64) rtcompare.Candidate {
	j := numProbes / 2
	return rtcompare.Candidate{Name: "hash-spill", Batch: func(n uint64) {
		var acc uint64
		for range n {
			if s[idx[j]].Contains(val[j]) {
				acc++
			}
			j = (j + 1) & (numProbes - 1)
		}
		sink += acc
	}}
}

func eachNew(s []vset.Set[uint64], idx []uint32) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: "array-spill", Batch: func(n uint64) {
		var acc uint64
		for range n {
			s[idx[j]].Each(func(v uint64) bool { acc += v; return true })
			j = (j + 1) & (numProbes - 1)
		}
		sink += acc
	}}
}

func eachOld(s []vset.HashSpill[uint64], idx []uint32) rtcompare.Candidate {
	j := numProbes / 2
	return rtcompare.Candidate{Name: "hash-spill", Batch: func(n uint64) {
		var acc uint64
		for range n {
			s[idx[j]].Each(func(v uint64) bool { acc += v; return true })
			j = (j + 1) & (numProbes - 1)
		}
		sink += acc
	}}
}

func addRemoveNew(s []vset.Set[uint64], idx []uint32) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: "array-spill", Batch: func(n uint64) {
		for range n {
			v := 1<<63 | uint64(j)
			s[idx[j]].Add(v)
			s[idx[j]].Remove(v)
			j = (j + 1) & (numProbes - 1)
		}
		sink++
	}}
}

func addRemoveOld(s []vset.HashSpill[uint64], idx []uint32) rtcompare.Candidate {
	j := numProbes / 2
	return rtcompare.Candidate{Name: "hash-spill", Batch: func(n uint64) {
		for range n {
			v := 1<<63 | uint64(j)
			s[idx[j]].Add(v)
			s[idx[j]].Remove(v)
			j = (j + 1) & (numProbes - 1)
		}
		sink++
	}}
}
