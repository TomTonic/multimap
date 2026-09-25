package main

import (
	"fmt"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/rtcompare"
)

// candidate returns the timed batch of one operation on one candidate. Each
// batch is written out per concrete type, as a user would call it, so that
// no interface or generic dictionary call distorts the comparison. Every
// batch keeps its own cursor, so successive batches walk the whole probe
// sequence instead of re-probing a warm prefix.
func (f *fixture) candidate(op, impl string) rtcompare.Candidate {
	var b func(uint64)
	switch op {
	case "valuesFor":
		b = f.valuesFor(impl)
	case "valuesBetween":
		b = f.valuesBetween(impl)
	case "churn":
		b = f.churn(impl)
	case "build":
		b = f.build(impl)
	}
	if b == nil {
		panic(fmt.Sprintf("no %s batch for %s", op, impl))
	}
	return rtcompare.Candidate{Name: impl, Batch: b}
}

// valuesFor iterates over all values of a random existing key.
func (f *fixture) valuesFor(impl string) func(uint64) {
	p, j := f.c.Hits, 0
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesForSeq(p.B[j]) {
					acc += v
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	case hashed:
		m := f.hsh
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesForSeq(p.B[j]) {
					acc += v
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	case btreeSets:
		m := f.bt
		return func(n uint64) {
			var acc uint64
			for range n {
				if s, ok := m.Get(p.S[j]); ok {
					for v := range s {
						acc += v
					}
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	case mapSets:
		m := f.gm
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m[p.S[j]] {
					acc += v
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	}
	return nil
}

// valuesBetween iterates over all values of rangeKeys consecutive keys.
// hashed and map-sets have no order and scan every key.
func (f *fixture) valuesBetween(impl string) func(uint64) {
	from, to, j := f.from, f.to, 0
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
					acc += v
				}
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	case hashed:
		m := f.hsh
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
					acc += v
				}
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	case btreeSets:
		m := f.bt
		return func(n uint64) {
			var acc uint64
			for range n {
				acc += btreeRangeSum(m, from.S[j], to.S[j])
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	case mapSets:
		m := f.gm
		return func(n uint64) {
			var acc uint64
			for range n {
				acc += mapRangeSum(m, from.S[j], to.S[j])
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	}
	return nil
}

// churn applies the next mutations of the churn cycle, wrapping around at
// its end. The candidate's position persists across comparisons, because the
// multimap is only in the matching state there.
func (f *fixture) churn(impl string) func(uint64) {
	ms, j := f.churnWorkload(), f.cur[impl]
	if j == nil {
		return nil
	}
	// step replays n mutations from *j on, in slices that end at the cycle's end.
	step := func(n uint64, apply func([]mutation)) {
		for n > 0 {
			c := min(n, uint64(len(ms)-*j))
			apply(ms[*j : *j+int(c)])
			if *j += int(c); *j == len(ms) {
				*j = 0
			}
			n -= c
		}
		sink++
	}
	kb, ks := f.ck.B, f.ck.S
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) { step(n, func(s []mutation) { applyOrdered(m, kb, s) }) }
	case hashed:
		m := f.hsh
		return func(n uint64) { step(n, func(s []mutation) { applyHashed(m, kb, s) }) }
	case btreeSets:
		m := f.bt
		return func(n uint64) { step(n, func(s []mutation) { applyBtree(m, ks, s) }) }
	case mapSets:
		m := f.gm
		return func(n uint64) { step(n, func(s []mutation) { applyMap(m, ks, s) }) }
	}
	return nil
}

// build replays the build workload on an empty multimap as one operation:
// the corpus with as many transient values inserted and deleted in between.
func (f *fixture) build(impl string) func(uint64) {
	ms, kb, ks := f.buildWorkload(), f.ck.B, f.ck.S
	switch impl {
	case ordered:
		return func(n uint64) {
			for range n {
				m := multimap.NewOrdered[uint64]()
				applyOrdered(m, kb, ms)
				sink += m.NumberOfKeys()
			}
		}
	case hashed:
		return func(n uint64) {
			for range n {
				m := multimap.NewHashed[uint64]()
				applyHashed(m, kb, ms)
				sink += m.NumberOfKeys()
			}
		}
	case btreeSets:
		return func(n uint64) {
			for range n {
				m := &btreeMM{}
				applyBtree(m, ks, ms)
				sink += uint64(m.Len())
			}
		}
	case mapSets:
		return func(n uint64) {
			for range n {
				m := mapMM{}
				applyMap(m, ks, ms)
				sink += uint64(len(m))
			}
		}
	}
	return nil
}
