// Command nodesearch compares the in-node child search strategies in isolation,
// with rtcompare: for a given node capacity, which way of finding the child for
// the next key byte is fastest once the node is in cache?
//
// Nodes are filled to capacity with random sorted key bytes; probes are random
// (node, present byte) pairs, so every lookup is a hit and the branch
// predictor cannot learn the position.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/TomTonic/multimap/bench/proto/swar"
	"github.com/TomTonic/rtcompare"
)

var sink uint64

const (
	numNodes  = 256 // 256 nodes of up to 256 B stay within L1/L2
	numProbes = 1 << 16
)

type node struct {
	keys   [64]byte // sorted, zero padded
	bitmap [4]uint64
	count  int
}

type probe struct {
	n *node
	b byte
}

func main() {
	out := flag.String("out", "results/nodesearch.jsonl", "JSON lines output")
	flag.Parse()
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

	linear := strat{"linear", findLinear}
	binary := strat{"bitmap+binary", findBinary}
	swar1 := strat{"swar", findSwar1}
	swar3 := strat{"swar", findSwar3}
	popcnt := strat{"bitmap+popcount", findPopcnt}

	cases := []struct {
		capacity int
		a, b     strat
	}{
		{4, swar1, linear},
		{8, swar1, linear},
		{22, swar3, linear},
		{22, swar3, binary},
		{22, popcnt, swar3},
		{52, popcnt, binary},
		{52, popcnt, linear},
	}
	for _, c := range cases {
		probes := makeProbes(c.capacity)
		check(probes, c.a.find, c.b.find)
		fmt.Fprintf(os.Stderr, "== capacity %d: %s vs %s\n", c.capacity, c.a.name, c.b.name)
		rep, err := rtcompare.Compare(batchFor(c.a.name, probes), batchFor(c.b.name, probes),
			rtcompare.CompareOptions{Collect: rtcompare.CollectOptions{MaxQuantizationError: 0.0001}})
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "%s\n\n", rep)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"capacity": c.capacity, "a": c.a.name, "b": c.b.name,
			"ns_a": rep.NsPerOpA, "ns_b": rep.NsPerOpB, "delta": rep.Estimate.Delta,
			"low": rep.Estimate.Low, "high": rep.Estimate.High, "resolved": rep.Resolved,
			"noise_floor": rep.NoiseFloor, "warnings": rep.Warnings,
		})
	}
	_ = sink
}

// strat names a search strategy; find is only used to cross-check results.
type strat struct {
	name string
	find func(*node, byte) int
}

// batchFor returns the batch loop for a strategy. Each loop calls its find
// function directly (not through a func value), so the compiler can inline it
// exactly as it would inside the tree.
func batchFor(name string, p []probe) rtcompare.Candidate {
	{
		j := 0
		var body func(n uint64)
		switch name {
		case "linear":
			body = func(n uint64) {
				var acc uint64
				for range n {
					acc += uint64(findLinear(p[j].n, p[j].b))
					j = (j + 1) & (numProbes - 1)
				}
				sink += acc
			}
		case "bitmap+binary":
			body = func(n uint64) {
				var acc uint64
				for range n {
					acc += uint64(findBinary(p[j].n, p[j].b))
					j = (j + 1) & (numProbes - 1)
				}
				sink += acc
			}
		case "bitmap+popcount":
			body = func(n uint64) {
				var acc uint64
				for range n {
					// findPopcnt, written out: it is inlined inside the tree's Get
					// but just over the inliner budget as a standalone function.
					x, b := p[j].n, p[j].b
					r := -1
					if swar.Has(&x.bitmap, b) {
						r = swar.Rank(&x.bitmap, b)
					}
					acc += uint64(r)
					j = (j + 1) & (numProbes - 1)
				}
				sink += acc
			}
		default: // swar: one word for small nodes, three for the 22-way node
			if p[0].n.count <= 8 {
				body = func(n uint64) {
					var acc uint64
					for range n {
						acc += uint64(findSwar1(p[j].n, p[j].b))
						j = (j + 1) & (numProbes - 1)
					}
					sink += acc
				}
			} else {
				body = func(n uint64) {
					var acc uint64
					for range n {
						// findSwar3, written out for the same reason.
						x, b := p[j].n, p[j].b
						i := swar.Index8(swar.Word(x.keys[0:8]), b)
						if i == 8 {
							i = 8 + swar.Index8(swar.Word(x.keys[8:16]), b)
							if i == 16 {
								i = 16 + swar.Index8(swar.Word(x.keys[16:24]), b)
							}
						}
						if i >= x.count {
							i = -1
						}
						acc += uint64(i)
						j = (j + 1) & (numProbes - 1)
					}
					sink += acc
				}
			}
		}
		return rtcompare.Candidate{Name: name, Batch: body}
	}
}

func findLinear(n *node, b byte) int {
	for i := 0; i < n.count; i++ {
		if n.keys[i&63] == b {
			return i
		}
	}
	return -1
}

func findBinary(n *node, b byte) int {
	if !swar.Has(&n.bitmap, b) {
		return -1
	}
	lo, hi := 0, n.count
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if n.keys[mid&63] < b {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func findSwar1(n *node, b byte) int {
	i := swar.Index8(swar.Word(n.keys[0:8]), b)
	if i >= n.count {
		return -1
	}
	return i
}

func findSwar3(n *node, b byte) int {
	i := swar.Index8(swar.Word(n.keys[0:8]), b)
	if i == 8 {
		i = 8 + swar.Index8(swar.Word(n.keys[8:16]), b)
		if i == 16 {
			i = 16 + swar.Index8(swar.Word(n.keys[16:24]), b)
		}
	}
	if i >= n.count {
		return -1
	}
	return i
}

func findPopcnt(n *node, b byte) int {
	if !swar.Has(&n.bitmap, b) {
		return -1
	}
	return swar.Rank(&n.bitmap, b)
}

func makeProbes(capacity int) []probe {
	rng := rtcompare.NewDPRNG(uint64(capacity) * 7919)
	nodes := make([]*node, numNodes)
	for i := range nodes {
		n := &node{count: capacity}
		seen := map[byte]bool{}
		var ks []byte
		for len(ks) < capacity {
			b := byte(rng.Uint64())
			if !seen[b] {
				seen[b] = true
				ks = append(ks, b)
			}
		}
		slices.Sort(ks)
		copy(n.keys[:], ks)
		for _, b := range ks {
			swar.Set(&n.bitmap, b)
		}
		nodes[i] = n
	}
	p := make([]probe, numProbes)
	for i := range p {
		n := nodes[rng.Uint64()%numNodes]
		p[i] = probe{n, n.keys[rng.Uint64()%uint64(capacity)]}
	}
	return p
}

func check(p []probe, a, b func(*node, byte) int) {
	for _, x := range p {
		if a(x.n, x.b) != b(x.n, x.b) || x.n.keys[a(x.n, x.b)] != x.b {
			panic(fmt.Sprintf("strategies disagree for byte %d", x.b))
		}
	}
}
