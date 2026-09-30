package main

import (
	"strings"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/rtcompare"
	"github.com/TomTonic/rtcompare/workload"
)

// Churn and build are the insertions and deletions of a database index, as
// streams from rtcompare's workload package: bursts of insertions alternate
// with bursts of deletions of values inserted earlier (see workload.Cycle and
// workload.Build). workload.Compare answers two questions per pair of
// candidates: churn, the cost of one insertion or deletion in a multimap in
// use, and build, the cost of building one from empty.
//
// A workload element is one key-value pair. Elements 0 to target-1 are the
// corpus's values, key by key; the elements from target on are transient
// values, which the streams insert and delete again (see pairs).

// transientBit marks the values of transient elements. Corpus values have it
// clear, so the two never collide and every operation changes the multimap.
const transientBit = 1 << 63

// pairs maps workload elements to key-value pairs: element id is the value
// val[id] of the churn key key[id] (see fixture.ck).
type pairs struct {
	key []uint32
	val []uint64
}

// newPairs lays out the elements of a corpus of n keys whose key i holds
// vals[offs[i]:offs[i+1]], and of the transient values a stream with ratio r
// inserts in all: (r-1) times as many as the corpus holds, each with its own
// value. Under unique, transient value t takes extra key t, so no key ever
// holds two values and every transient insertion creates a key; that needs
// one extra key per transient value, which extraKeys provides for r <= 2.
// Otherwise half of them go to random corpus keys, whose value sets then grow
// and shrink, and half to n/2 extra keys, which appear and disappear as whole
// keys.
func newPairs(n int, vals []uint64, offs []int, r float64, unique bool) pairs {
	target := len(vals)
	t := int((r-1)*float64(target) + 0.5)
	p := pairs{key: make([]uint32, target+t), val: make([]uint64, target+t)}
	for i := range n {
		for j := offs[i]; j < offs[i+1]; j++ {
			p.key[j], p.val[j] = uint32(i), vals[j]
		}
	}
	rng := rtcompare.NewDPRNG(0x7A15)
	extra := uint64(max(1, n/2))
	for i := range t {
		k := uint64(n + i)
		if !unique {
			k = rng.Uint64() % uint64(n)
			if rng.Uint64()&1 == 1 {
				k = uint64(n) + rng.Uint64()%extra
			}
		}
		p.key[target+i], p.val[target+i] = uint32(k), transientBit|uint64(i)
	}
	return p
}

// stream is the shape of the churn and build streams of a scenario: the
// insertions per value the multimap holds in the end, and the share of the
// deletions that take out a long-lived value (a corpus value) and put it back
// later, so that the merges and shrinks of real deletions reach the whole
// multimap and not only its newest part.
type stream struct{ ratio, permChurn float64 }

// workloadConfig is the configuration of both streams of a scenario. The
// streams carry no lookups (workload.Config.Lookups), so an Apply function
// only sees insertions and deletions. Keys stay at rtcompare's default, a
// seeded permutation, though the driver maps the elements' IDs to corpus keys
// itself (see newPairs): the IDs alone are sequential, and every transient one
// lies above the permanent ones.
func workloadConfig(s stream) workload.Config {
	return workload.Config{Seed: 0xC4A2, Ratio: s.ratio, PermanentChurn: s.permChurn}
}

// structure describes candidate impl to workload.Compare. Each Apply asserts
// the concrete type once per run of operations and loops over the run
// itself, so that the candidate's methods are called directly.
func (f *fixture) structure(impl string) workload.Structure[any] {
	kb, ks, p := f.ck.B, f.ck.S, f.pairs
	s := workload.Structure[any]{Name: impl}
	switch impl {
	case ordered:
		s.New = func() any { return multimap.NewOrdered[uint64]() }
		s.Apply = func(a any, run []workload.Op) {
			m := a.(*multimap.Ordered[uint64])
			for _, op := range run {
				if op.Kind == workload.Insert {
					m.AddValue(kb[p.key[op.ID]], p.val[op.ID])
				} else {
					m.RemoveValue(kb[p.key[op.ID]], p.val[op.ID])
				}
			}
		}
	case hashed:
		s.New = func() any { return multimap.NewHashed[uint64]() }
		s.Apply = func(a any, run []workload.Op) {
			m := a.(*multimap.Hashed[uint64])
			for _, op := range run {
				if op.Kind == workload.Insert {
					m.AddValue(kb[p.key[op.ID]], p.val[op.ID])
				} else {
					m.RemoveValue(kb[p.key[op.ID]], p.val[op.ID])
				}
			}
		}
	case btreeSets:
		s.New = func() any { return &btreeMM{} }
		s.Apply = func(a any, run []workload.Op) {
			m := a.(*btreeMM)
			for _, op := range run {
				if op.Kind == workload.Insert {
					btreeAdd(m, ks[p.key[op.ID]], p.val[op.ID])
				} else {
					btreeRemove(m, ks[p.key[op.ID]], p.val[op.ID])
				}
			}
		}
	case mapSets:
		s.New = func() any { return mapMM{} }
		s.Apply = func(a any, run []workload.Op) {
			m := a.(mapMM)
			for _, op := range run {
				if op.Kind == workload.Insert {
					mapAdd(m, ks[p.key[op.ID]], p.val[op.ID])
				} else {
					mapRemove(m, ks[p.key[op.ID]], p.val[op.ID])
				}
			}
		}
	case btreeMapC:
		s.New = func() any { return &btreeMap{} }
		s.Apply = func(a any, run []workload.Op) {
			m := a.(*btreeMap)
			for _, op := range run {
				if op.Kind == workload.Insert {
					m.Set(strings.Clone(ks[p.key[op.ID]]), p.val[op.ID])
				} else {
					m.Delete(ks[p.key[op.ID]])
				}
			}
		}
	case baseline:
		s.New = baseKit.empty
		s.Apply = func(a any, run []workload.Op) { baseKit.apply(a, kb, p.key, p.val, run) }
	}
	return s
}
