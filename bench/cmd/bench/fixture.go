package main

import (
	"strings"

	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/layout"
	"github.com/TomTonic/rtcompare"
)

// The four candidates a user chooses between. "btree-sets" and "map-sets" are
// what one would write by hand without this library: a tidwall/btree.Map or a
// Go map from string keys to Go-map sets, with empty keys removed as the
// library does.
const (
	ordered   = "ordered"
	hashed    = "hashed"
	btreeSets = "btree-sets"
	mapSets   = "map-sets"
)

var allImpls = []string{ordered, hashed, btreeSets, mapSets}

type (
	btreeMM = btree.Map[string, map[uint64]struct{}]
	mapMM   = map[string]map[uint64]struct{}
)

// fixture holds one scenario's corpus and every candidate built from it.
type fixture struct {
	c    keys.Corpus
	vals []uint64
	offs []int
	ord  *multimap.Ordered[uint64]
	hsh  *multimap.Hashed[uint64]
	bt   *btreeMM
	gm   mapMM
	// ranges of rangeKeys consecutive keys, as []byte and string views
	from, to keys.Set
}

const rangeKeys = 100

// newFixture generates the corpus and builds the candidates named in impls,
// one after another. With -layoutseed the build order is shuffled and a random
// spacer precedes each build, so that each process samples its own layout.
func newFixture(kind keys.Kind, n int, impls []string) *fixture {
	f := &fixture{c: keys.Generate(kind, n, 0x5EED)}
	f.vals, f.offs = keys.Values(n, 0xFA11)
	builds := map[string]func(){
		ordered:   func() { f.ord = buildOrdered(f.c.Keys.B, f.vals, f.offs) },
		hashed:    func() { f.hsh = buildHashed(f.c.Keys.B, f.vals, f.offs) },
		btreeSets: func() { f.bt = buildBtree(f.c.Keys.S, f.vals, f.offs) },
		mapSets:   func() { f.gm = buildMap(f.c.Keys.S, f.vals, f.offs) },
	}
	order := append([]string(nil), impls...)
	layout.Shuffle(order)
	for _, name := range order {
		layout.Spacer()
		builds[name]()
	}
	f.from, f.to = ranges(f.c.Keys, n)
	return f
}

// ranges picks the probe ranges: rangeKeys consecutive keys in key order,
// starting at random keys.
func ranges(all keys.Set, n int) (from, to keys.Set) {
	sorted := keys.Sorted(all)
	rng := rtcompare.NewDPRNG(0xB0B)
	var f, t [][]byte
	for range min(n, 1<<16) {
		i := int(rng.Uint64() % uint64(n-rangeKeys+1))
		f, t = append(f, sorted.B[i]), append(t, sorted.B[i+rangeKeys-1])
	}
	return keys.Pack(f), keys.Pack(t)
}

func buildOrdered(k [][]byte, vals []uint64, offs []int) *multimap.Ordered[uint64] {
	m := multimap.NewOrdered[uint64]()
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
	return m
}

func buildHashed(k [][]byte, vals []uint64, offs []int) *multimap.Hashed[uint64] {
	m := multimap.NewHashed[uint64]()
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
	return m
}

func buildBtree(k []string, vals []uint64, offs []int) *btreeMM {
	m := &btreeMM{}
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			btreeAdd(m, key, v)
		}
	}
	return m
}

func buildMap(k []string, vals []uint64, offs []int) mapMM {
	m := mapMM{}
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			mapAdd(m, key, v)
		}
	}
	return m
}

// The hand-written candidates own their keys, as the library does: the corpus
// strings are views into one shared buffer.

func btreeAdd(m *btreeMM, k string, v uint64) {
	s, ok := m.Get(k)
	if !ok {
		s = map[uint64]struct{}{}
		m.Set(strings.Clone(k), s)
	}
	s[v] = struct{}{}
}

func btreeRemove(m *btreeMM, k string, v uint64) {
	if s, ok := m.Get(k); ok {
		delete(s, v)
		if len(s) == 0 {
			m.Delete(k)
		}
	}
}

func mapAdd(m mapMM, k string, v uint64) {
	s := m[k]
	if s == nil {
		s = map[uint64]struct{}{}
		m[strings.Clone(k)] = s
	}
	s[v] = struct{}{}
}

func mapRemove(m mapMM, k string, v uint64) {
	if s := m[k]; s != nil {
		delete(s, v)
		if len(s) == 0 {
			delete(m, k)
		}
	}
}
