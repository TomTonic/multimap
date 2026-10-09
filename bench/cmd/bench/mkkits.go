package main

import (
	"fmt"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/rtcompare/workload"
)

// mkKits are the kits of the candidates ordered-mk1, ordered-mk2 and ordered-mk4 (experiment exp-maxkeys, the curve
// of the autotune, docs/redesign/autotune-design.md): Ordered with at most 1, 2 or 4 keys a page. They are
// candidates only when -vs names them.
var mkKits = map[string]*kit{}

// mkNames are their names, in order.
var mkNames []string

func init() {
	for _, k := range []int{1, 2, 4} {
		name := fmt.Sprintf("ordered-mk%d", k)
		mkNames = append(mkNames, name)
		mkKits[name] = orderedKit(k)
	}
}

// kitOf returns the kit of candidate impl, or nil if it has none.
func kitOf(impl string) *kit {
	if impl == baseline {
		return baseKit
	}
	return mkKits[impl]
}

// orderedKit returns the kit of Ordered with at most maxKeys keys a page; its loops are those of the ordered
// candidate.
func orderedKit(maxKeys int) *kit {
	type om = multimap.Ordered[V]
	newMap := func() *om {
		m := multimap.NewOrdered[V]()
		m.SetMaxKeys(maxKeys)
		return m
	}
	return &kit{
		ref: fmt.Sprintf("this commit, at most %d keys a page", maxKeys),
		build: func(k [][]byte, vals []V, offs []int) any {
			m := newMap()
			for i, key := range k {
				for _, v := range vals[offs[i]:offs[i+1]] {
					m.AddValue(key, v)
				}
			}
			return m
		},
		empty: func() any { return newMap() },
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
