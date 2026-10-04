package main

import (
	"fmt"

	"github.com/TomTonic/multimap/bench/keys"
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
		b = f.valuesBetween(impl, f.from, f.to)
	case "prefix":
		b = f.valuesBetween(impl, f.pfrom, f.pto)
	}
	if b == nil {
		panic(fmt.Sprintf("no %s batch for %s", op, impl))
	}
	return rtcompare.Candidate{Name: impl, Batch: b}
}

// valuesFor iterates over all values of a random existing key, in the order
// of the corpus's Probes.
func (f *fixture) valuesFor(impl string) func(uint64) {
	p, j := f.c.Probes, 0
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesForSeq(p.B[j]) {
					acc += weigh(v)
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	case baseline:
		return baseKit.valuesFor(f.base, p)
	case hashed:
		m := f.hsh
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesForSeq(p.B[j]) {
					acc += weigh(v)
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
						acc += weigh(v)
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
					acc += weigh(v)
				}
				if j++; j == len(p.B) {
					j = 0
				}
			}
			sink += acc
		}
	case btreeMapC:
		m := f.bm
		return func(n uint64) {
			var acc uint64
			for range n {
				if v, ok := m.Get(p.S[j]); ok {
					acc += weigh(v)
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

// valuesBetween iterates over all values of the keys in [from[j], to[j]] for
// successive j: rangeKeys consecutive keys (valuesBetween), or all keys with
// a prefix (prefix). hashed and map-sets have no order and scan every key.
func (f *fixture) valuesBetween(impl string, from, to keys.Set) func(uint64) {
	j := 0
	switch impl {
	case ordered:
		m := f.ord
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
					acc += weigh(v)
				}
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	case baseline:
		return baseKit.between(f.base, from, to)
	case hashed:
		m := f.hsh
		return func(n uint64) {
			var acc uint64
			for range n {
				for v := range m.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
					acc += weigh(v)
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
	case btreeMapC:
		m := f.bm
		return func(n uint64) {
			var acc uint64
			for range n {
				acc += btreeMapRangeSum(m, from.S[j], to.S[j])
				if j++; j == len(from.B) {
					j = 0
				}
			}
			sink += acc
		}
	}
	return nil
}
