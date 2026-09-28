//go:build baseline

package main

import (
	multimap "github.com/TomTonic/multimap/bench/baseline"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/rtcompare/workload"
)

// The kit of the baseline candidate: the same code as the ordered candidate's,
// on the copy of Ordered in package baseline.
func init() {
	type om = multimap.Ordered[V]
	baseKit = &kit{
		ref: multimap.BaselineRef,
		build: func(k [][]byte, vals []V, offs []int) any {
			m := multimap.NewOrdered[V]()
			for i, key := range k {
				for _, v := range vals[offs[i]:offs[i+1]] {
					m.AddValue(key, v)
				}
			}
			return m
		},
		empty: func() any { return multimap.NewOrdered[V]() },
		apply: func(a any, k [][]byte, key []uint32, val []V, run []workload.Op) {
			m := a.(*om)
			for _, op := range run {
				if op.Kind == workload.Insert {
					m.AddValue(k[key[op.ID]], val[op.ID])
				} else {
					m.RemoveValue(k[key[op.ID]], val[op.ID])
				}
			}
		},
		valuesFor: func(a any, p keys.Set) func(uint64) {
			m, j := a.(*om), 0
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
		},
		between: func(a any, from, to keys.Set) func(uint64) {
			m, j := a.(*om), 0
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
		},
		sum:       func(a any, key []byte) uint64 { return sum(a.(*om).ValuesForSeq(key)) },
		rangeSum:  func(a any, from, to []byte) uint64 { return sum(a.(*om).ValuesBetweenInclusiveSeq(from, to)) },
		keys:      func(a any) int { return int(a.(*om).NumberOfKeys()) },
		removeKey: func(a any) func(key []byte) { m := a.(*om); return func(key []byte) { m.RemoveKey(key) } },
	}
}
