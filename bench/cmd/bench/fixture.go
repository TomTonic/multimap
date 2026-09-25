package main

import (
	"slices"
	"strings"

	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/layout"
	"github.com/TomTonic/rtcompare"
)

// The candidates a user chooses between. "btree-sets" and "map-sets" are what
// one would write by hand without this library: a tidwall/btree.Map or a Go
// map from string keys to Go-map sets, with empty keys removed as the library
// does. "btree-map" is a plain tidwall/btree.Map with one value per key, for
// users whose keys (almost) never hold more than one value.
const (
	ordered   = "ordered"
	hashed    = "hashed"
	btreeSets = "btree-sets"
	mapSets   = "map-sets"
	btreeMapC = "btree-map"
)

// The value profiles: multi gives keys a skewed number of values (see
// keys.Values), unique exactly one value per key.
const (
	multi  = "multi"
	unique = "unique"
)

// implsFor returns the candidates compared under a value profile.
func implsFor(profile string) []string {
	if profile == unique {
		return []string{ordered, btreeMapC}
	}
	return []string{ordered, hashed, btreeSets, mapSets}
}

type (
	btreeMM  = btree.Map[string, map[uint64]struct{}]
	mapMM    = map[string]map[uint64]struct{}
	btreeMap = btree.Map[string, uint64]
)

// fixture holds one scenario's corpus and every candidate built from it.
type fixture struct {
	c       keys.Corpus
	profile string
	vals    []uint64
	offs    []int
	ord     *multimap.Ordered[uint64]
	hsh     *multimap.Hashed[uint64]
	bt      *btreeMM
	gm      mapMM
	bm      *btreeMap
	// ranges of rangeKeys consecutive keys, as []byte and string views
	from, to keys.Set
	// index workloads: churn keys (corpus keys, then extra keys), the ratio of
	// insertions to final values, the workloads (made on first use) and each
	// candidate's position in the churn cycle
	ck             keys.Set
	ratio          float64
	buildW, churnW []mutation
	cur            map[string]*int
}

const rangeKeys = 100

// newFixture generates the corpus and its values under the value profile and
// builds the candidates named in impls, one after another. With -layoutseed
// the build order is shuffled and a random spacer precedes each build, so
// that each process samples its own layout.
func newFixture(kind keys.Kind, n int, profile string, impls []string) *fixture {
	f := &fixture{c: keys.Generate(kind, n, 0x5EED), profile: profile}
	f.vals, f.offs = profileValues(profile, n)
	builds := map[string]func(){
		ordered:   func() { f.ord = buildOrdered(f.c.Keys.B, f.vals, f.offs) },
		hashed:    func() { f.hsh = buildHashed(f.c.Keys.B, f.vals, f.offs) },
		btreeSets: func() { f.bt = buildBtree(f.c.Keys.S, f.vals, f.offs) },
		mapSets:   func() { f.gm = buildMap(f.c.Keys.S, f.vals, f.offs) },
		btreeMapC: func() { f.bm = buildBtreeMap(f.c.Keys.S, f.vals, f.offs) },
	}
	order := append([]string(nil), impls...)
	layout.Shuffle(order)
	for _, name := range order {
		layout.Spacer()
		builds[name]()
	}
	f.from, f.to = ranges(f.c.Keys, n)
	f.ck = keys.Pack(append(slices.Clone(f.c.Keys.B), f.c.Misses.B[:extraKeys(profile, n)]...))
	f.ratio, f.cur = 2, map[string]*int{}
	for _, name := range impls {
		f.cur[name] = new(int)
	}
	return f
}

// profileValues returns the values of n keys under a value profile: key i
// holds vals[offs[i]:offs[i+1]].
func profileValues(profile string, n int) (vals []uint64, offs []int) {
	vals, offs = keys.Values(n, 0xFA11)
	if profile != unique {
		return vals, offs
	}
	// one value per key: the first of each key's multi values
	one, idx := make([]uint64, n), make([]int, n+1)
	for i := range n {
		one[i], idx[i+1] = vals[offs[i]], i+1
	}
	return one, idx
}

// extraKeys is the number of keys that only the index workloads use: n/2 for
// multi, where transient values also go to corpus keys; n for unique, where
// each transient value needs a key of its own (see keyPool).
func extraKeys(profile string, n int) int {
	if profile == unique {
		return n
	}
	return max(1, n/2)
}

// buildWorkload returns the build workload, computing it on first use.
func (f *fixture) buildWorkload() []mutation {
	if f.buildW == nil {
		f.buildW = buildWorkload(len(f.c.Keys.B), f.vals, f.offs, f.ratio, 0xB11D, f.profile == unique)
	}
	return f.buildW
}

// churnWorkload returns the churn cycle, computing it on first use.
func (f *fixture) churnWorkload() []mutation {
	if f.churnW == nil {
		f.churnW = churnWorkload(len(f.c.Keys.B), f.vals, f.ratio, 0xC4A2, f.profile == unique)
	}
	return f.churnW
}

// settle plays every candidate's churn cycle to its end, untimed, so that
// each holds exactly the corpus again before other operations are timed.
func (f *fixture) settle() {
	ms := f.churnWorkload()
	for name, j := range f.cur {
		rest := ms[*j:]
		switch name {
		case ordered:
			applyOrdered(f.ord, f.ck.B, rest)
		case hashed:
			applyHashed(f.hsh, f.ck.B, rest)
		case btreeSets:
			applyBtree(f.bt, f.ck.S, rest)
		case mapSets:
			applyMap(f.gm, f.ck.S, rest)
		case btreeMapC:
			applyBtreeMap(f.bm, f.ck.S, rest)
		}
		*j = 0
	}
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

func buildBtreeMap(k []string, vals []uint64, offs []int) *btreeMap {
	m := &btreeMap{}
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.Set(strings.Clone(key), v)
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
