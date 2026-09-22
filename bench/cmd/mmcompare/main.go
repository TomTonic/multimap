// Command mmcompare runs the realistic multimap comparisons with rtcompare:
// keys carry sets of uint64 values (see keys.Values), results are consumed
// through iterators, and the ART multimap (A) is compared with the two ways of
// using tidwall/btree.Map (B), all sharing the same value container.
//
// Operations, each timed per call:
//
//	valuesFor      iterate all values of a random existing key
//	valuesBetween  iterate all values of 100 consecutive keys [from, to]
//	addRemove      add a new value to a random existing key, then remove it
//	build          build the whole multimap from scratch (small n only)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/proto/mmart"
	"github.com/TomTonic/multimap/bench/proto/mmbtree"
	"github.com/TomTonic/rtcompare"
)

var sink uint64

const rangeKeys = 100

type multimap interface {
	AddValue(key []byte, v uint64)
	RemoveValue(key []byte, v uint64)
	ValuesFor(key []byte) iter.Seq[uint64]
	ValuesBetween(from, to []byte) iter.Seq[uint64]
	Len() int
}

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
	Warnings   []string `json:"warnings"`
}

type data struct {
	c        keys.Corpus
	vals     []uint64
	offs     []int
	from, to keys.Set // range probes, 100 consecutive keys each
	art      *mmart.Map[uint64]
	btInline *mmbtree.Inline[uint64]
	btPtr    *mmbtree.Ptr[uint64]
}

func main() {
	kind := flag.String("keys", "u64", "u64 or str")
	n := flag.Int("n", 4096, "number of keys")
	ops := flag.String("ops", "valuesFor,valuesBetween,addRemove", "operations")
	out := flag.String("out", "results/mm.jsonl", "JSON lines output")
	only := flag.String("only", "", "comma-separated B candidates to run (default all)")
	flag.Parse()

	d := load(keys.Kind(*kind), *n)
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

	for _, op := range strings.Split(*ops, ",") {
		if op == "addRemove" {
			d.warmAddRemove() // settle every key's value set into its steady state first
		}
		cands := d.candidates(op)
		for _, b := range cands[1:] {
			if *only != "" && !slices.Contains(strings.Split(*only, ","), b.Name) {
				continue
			}
			opt := rtcompare.CompareOptions{Collect: rtcompare.CollectOptions{MaxQuantizationError: 0.0001}}
			if op == "build" {
				// building allocates a whole structure per operation; collect
				// between batches so one candidate's garbage is not charged to
				// the other. (Not for addRemove: after warmAddRemove it hardly
				// allocates, and a forced GC per batch over a 1M-key heap
				// would make the run take hours.)
				opt.Collect.GCBetween = true
			}
			fmt.Fprintf(os.Stderr, "== %s n=%d %s: %s vs %s\n", *kind, *n, op, cands[0].Name, b.Name)
			rep, err := rtcompare.Compare(cands[0], b, opt)
			if err != nil {
				fail(err)
			}
			fmt.Fprintf(os.Stderr, "%s\n\n", rep)
			if err := json.NewEncoder(w).Encode(result{
				Keys: *kind, N: *n, Op: op, A: cands[0].Name, B: b.Name,
				NsA: rep.NsPerOpA, NsB: rep.NsPerOpB, Delta: rep.Estimate.Delta,
				Low: rep.Estimate.Low, High: rep.Estimate.High, Resolved: rep.Resolved,
				NoiseFloor: rep.NoiseFloor, Warnings: rep.Warnings,
			}); err != nil {
				fail(err)
			}
		}
	}
	_ = sink
}

func load(kind keys.Kind, n int) *data {
	d := &data{c: keys.Generate(kind, n, 0x5EED)}
	d.vals, d.offs = keys.Values(n, 0xFA11)
	d.art, d.btInline, d.btPtr = &mmart.Map[uint64]{}, &mmbtree.Inline[uint64]{}, &mmbtree.Ptr[uint64]{}
	for _, m := range []multimap{d.art, d.btInline, d.btPtr} {
		d.fill(m)
	}

	sorted := keys.Sorted(d.c.Keys)
	rng := rtcompare.NewDPRNG(0xB0B)
	var from, to [][]byte
	for range min(n, 1<<16) {
		i := int(rng.Uint64() % uint64(n-rangeKeys+1))
		from, to = append(from, sorted.B[i]), append(to, sorted.B[i+rangeKeys-1])
	}
	d.from, d.to = keys.Pack(from), keys.Pack(to)
	d.verify()
	return d
}

func (d *data) fill(m multimap) {
	for i, k := range d.c.Keys.B {
		for _, v := range d.vals[d.offs[i]:d.offs[i+1]] {
			m.AddValue(k, v)
		}
	}
}

// verify makes sure all three answer identically before anything is timed.
func (d *data) verify() {
	sum := func(s iter.Seq[uint64]) (a, c uint64) {
		for v := range s {
			a += v
			c++
		}
		return
	}
	for i := range min(len(d.c.Hits.B), 20000) {
		k := d.c.Hits.B[i]
		a1, c1 := sum(d.art.ValuesFor(k))
		a2, c2 := sum(d.btInline.ValuesFor(k))
		a3, c3 := sum(d.btPtr.ValuesFor(k))
		if a1 != a2 || a1 != a3 || c1 != c2 || c1 != c3 || c1 == 0 {
			fail(fmt.Errorf("ValuesFor mismatch for %q", k))
		}
	}
	for i := range min(len(d.from.B), 2000) {
		f, t := d.from.B[i], d.to.B[i]
		a1, c1 := sum(d.art.ValuesBetween(f, t))
		if a0, c0 := sum(d.art.ValuesBetweenLinear(f, t)); a0 != a1 || c0 != c1 {
			fail(fmt.Errorf("ValuesBetweenLinear mismatch for %q..%q", f, t))
		}
		a2, c2 := sum(d.btInline.ValuesBetween(f, t))
		a3, c3 := sum(d.btPtr.ValuesBetween(f, t))
		if a1 != a2 || a1 != a3 || c1 != c2 || c1 != c3 || c1 < rangeKeys {
			fail(fmt.Errorf("ValuesBetween mismatch for %q..%q", f, t))
		}
	}
}

// warmAddRemove performs the addRemove operation once for every probe key on
// every structure, so the measured runs start from the state they keep.
func (d *data) warmAddRemove() {
	for i, k := range d.c.Hits.B {
		v := absent(i)
		for _, m := range []multimap{d.art, d.btInline, d.btPtr} {
			m.AddValue(k, v)
			m.RemoveValue(k, v)
		}
	}
}

func absent(i int) uint64 { return 1<<63 | uint64(i) }

func (d *data) candidates(op string) []rtcompare.Candidate {
	switch op {
	case "valuesFor":
		return []rtcompare.Candidate{
			valuesForArt("art", d.art, d.c.Hits), valuesForBtInline("btree-inline", d.btInline, d.c.Hits),
			valuesForBtPtr("btree-ptr", d.btPtr, d.c.Hits)}
	case "valuesBetween":
		return []rtcompare.Candidate{
			valuesBetweenArt("art", d.art, d.from, d.to), valuesBetweenArtLinear("art-linear", d.art, d.from, d.to),
			valuesBetweenBtInline("btree-inline", d.btInline, d.from, d.to), valuesBetweenBtPtr("btree-ptr", d.btPtr, d.from, d.to)}
	case "addRemove":
		return []rtcompare.Candidate{
			addRemoveArt("art", d.art, d.c.Hits), addRemoveBtInline("btree-inline", d.btInline, d.c.Hits),
			addRemoveBtPtr("btree-ptr", d.btPtr, d.c.Hits)}
	case "build":
		return []rtcompare.Candidate{d.buildArt("art"), d.buildBtInline("btree-inline"), d.buildBtPtr("btree-ptr")}
	}
	fail(fmt.Errorf("unknown op %q", op))
	return nil
}

// The batch functions below are written out once per concrete multimap type
// from the same template. Neither the interface nor a generic function would
// do: with pointer type arguments Go shares one instantiation and calls
// methods through a dictionary, which forbids inlining. Each batch keeps its
// own cursor across batches.

func valuesForArt(name string, m *mmart.Map[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesFor(p.B[j]) {
				acc += v
			}
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func valuesBetweenArt(name string, m *mmart.Map[uint64], from, to keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesBetween(from.B[j], to.B[j]) {
				acc += v
			}
			if j++; j == len(from.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

// valuesBetweenArtLinear measures the ART's former range scan, which compared
// every key with the upper bound. It runs on the same tree as "art", so its
// cursor starts half-way through the probes: otherwise, with ABBA ordering,
// each batch would often re-scan the ranges the other candidate had just
// pulled into the cache.
func valuesBetweenArtLinear(name string, m *mmart.Map[uint64], from, to keys.Set) rtcompare.Candidate {
	j := len(from.B) / 2
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesBetweenLinear(from.B[j], to.B[j]) {
				acc += v
			}
			if j++; j == len(from.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func addRemoveArt(name string, m *mmart.Map[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			k, v := p.B[j], absent(j)
			m.AddValue(k, v)
			m.RemoveValue(k, v)
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink++
	}}
}

func (d *data) buildArt(name string) rtcompare.Candidate {
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			m := &mmart.Map[uint64]{}
			for i, k := range d.c.Keys.B {
				for _, v := range d.vals[d.offs[i]:d.offs[i+1]] {
					m.AddValue(k, v)
				}
			}
			sink += uint64(m.Len())
		}
	}}
}

func valuesForBtInline(name string, m *mmbtree.Inline[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesFor(p.B[j]) {
				acc += v
			}
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func valuesBetweenBtInline(name string, m *mmbtree.Inline[uint64], from, to keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesBetween(from.B[j], to.B[j]) {
				acc += v
			}
			if j++; j == len(from.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func addRemoveBtInline(name string, m *mmbtree.Inline[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			k, v := p.B[j], absent(j)
			m.AddValue(k, v)
			m.RemoveValue(k, v)
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink++
	}}
}

func (d *data) buildBtInline(name string) rtcompare.Candidate {
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			m := &mmbtree.Inline[uint64]{}
			for i, k := range d.c.Keys.B {
				for _, v := range d.vals[d.offs[i]:d.offs[i+1]] {
					m.AddValue(k, v)
				}
			}
			sink += uint64(m.Len())
		}
	}}
}

func valuesForBtPtr(name string, m *mmbtree.Ptr[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesFor(p.B[j]) {
				acc += v
			}
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func valuesBetweenBtPtr(name string, m *mmbtree.Ptr[uint64], from, to keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		var acc uint64
		for range n {
			for v := range m.ValuesBetween(from.B[j], to.B[j]) {
				acc += v
			}
			if j++; j == len(from.B) {
				j = 0
			}
		}
		sink += acc
	}}
}

func addRemoveBtPtr(name string, m *mmbtree.Ptr[uint64], p keys.Set) rtcompare.Candidate {
	j := 0
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			k, v := p.B[j], absent(j)
			m.AddValue(k, v)
			m.RemoveValue(k, v)
			if j++; j == len(p.B) {
				j = 0
			}
		}
		sink++
	}}
}

func (d *data) buildBtPtr(name string) rtcompare.Candidate {
	return rtcompare.Candidate{Name: name, Batch: func(n uint64) {
		for range n {
			m := &mmbtree.Ptr[uint64]{}
			for i, k := range d.c.Keys.B {
				for _, v := range d.vals[d.offs[i]:d.offs[i+1]] {
					m.AddValue(k, v)
				}
			}
			sink += uint64(m.Len())
		}
	}}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "mmcompare:", err)
	os.Exit(1)
}
