package main

import (
	"fmt"

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
	case "addRemove":
		b = f.addRemove(impl)
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
// map-sets has no order and therefore no such operation.
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
	}
	return nil
}

// addRemove adds a value a random existing key does not hold, then removes
// it again; the key keeps its other values throughout.
func (f *fixture) addRemove(impl string) func(uint64) {
	p, j := f.c.Hits, 0
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) {
			for range n {
				k, v := p.B[j], absent(j)
				m.AddValue(k, v)
				m.RemoveValue(k, v)
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink++
		}
	case hashed:
		m := f.hsh
		return func(n uint64) {
			for range n {
				k, v := p.B[j], absent(j)
				m.AddValue(k, v)
				m.RemoveValue(k, v)
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink++
		}
	case btreeSets:
		m := f.bt
		return func(n uint64) {
			for range n {
				k, v := p.S[j], absent(j)
				btreeAdd(m, k, v)
				btreeRemove(m, k, v)
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink++
		}
	case mapSets:
		m := f.gm
		return func(n uint64) {
			for range n {
				k, v := p.S[j], absent(j)
				mapAdd(m, k, v)
				mapRemove(m, k, v)
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink++
		}
	}
	return nil
}

// build constructs the whole multimap from the corpus as one operation.
func (f *fixture) build(impl string) func(uint64) {
	kb, ks, vals, offs := f.c.Keys.B, f.c.Keys.S, f.vals, f.offs
	switch impl {
	case ordered:
		return func(n uint64) {
			for range n {
				sink += buildOrdered(kb, vals, offs).NumberOfKeys()
			}
		}
	case hashed:
		return func(n uint64) {
			for range n {
				sink += buildHashed(kb, vals, offs).NumberOfKeys()
			}
		}
	case btreeSets:
		return func(n uint64) {
			for range n {
				sink += uint64(buildBtree(ks, vals, offs).Len())
			}
		}
	case mapSets:
		return func(n uint64) {
			for range n {
				sink += uint64(len(buildMap(ks, vals, offs)))
			}
		}
	}
	return nil
}
