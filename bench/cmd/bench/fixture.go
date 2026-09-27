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

// vsOnly, if not empty, limits the candidates ordered is compared with to
// these (-vs).
var vsOnly []string

// implsFor returns the candidates compared under a value profile: ordered
// first, then those it is compared with, including the baseline in a bench
// built with it (see kit.go), all of them limited by vsOnly.
func implsFor(profile string) []string {
	others := []string{hashed, btreeSets, mapSets}
	if profile == unique {
		others = []string{btreeMapC}
	}
	if baseKit != nil {
		others = append(others, baseline)
	}
	out := []string{ordered}
	for _, b := range others {
		if len(vsOnly) == 0 || slices.Contains(vsOnly, b) {
			out = append(out, b)
		}
	}
	return out
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
	base    any      // the baseline, see kit.go
	others  []string // the candidates built besides ordered
	// ranges of rangeKeys consecutive keys, as []byte and string views, and
	// the ranges of the keys that start with the prefix of a random key (see
	// keys.Prefix; text keys only)
	from, to   keys.Set
	pfrom, pto keys.Set
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
	f.vals, f.offs = profileValues(f.c, profile, n)
	builds := map[string]func(){
		ordered:   func() { f.ord = buildOrdered(f.c.Keys.B, f.vals, f.offs) },
		hashed:    func() { f.hsh = buildHashed(f.c.Keys.B, f.vals, f.offs) },
		btreeSets: func() { f.bt = buildBtree(f.c.Keys.S, f.vals, f.offs) },
		mapSets:   func() { f.gm = buildMap(f.c.Keys.S, f.vals, f.offs) },
		btreeMapC: func() { f.bm = buildBtreeMap(f.c.Keys.S, f.vals, f.offs) },
		baseline:  func() { f.base = baseKit.build(f.c.Keys.B, f.vals, f.offs) },
	}
	f.others = slices.DeleteFunc(slices.Clone(impls), func(s string) bool { return s == ordered })
	order := append([]string(nil), impls...)
	layout.Shuffle(order)
	for _, name := range order {
		layout.Spacer()
		builds[name]()
	}
	f.from, f.to = ranges(f.c.Keys, n)
	if keys.Text(kind) {
		f.pfrom, f.pto = prefixes(kind, f.c.Keys, n)
	}
	f.ck = keys.Pack(append(slices.Clone(f.c.Keys.B), f.c.Misses.B[:extraKeys(profile, n)]...))
	f.ratio, f.cur = 2, map[string]*int{}
	for _, name := range impls {
		f.cur[name] = new(int)
	}
	return f
}

// profileValues returns the values of the n keys of c under a value profile:
// key i holds vals[offs[i]:offs[i+1]]. Under multi, keys with natural values
// (street names: their localities) hold those, others a skewed number of
// synthetic ones (see keys.Values).
func profileValues(c keys.Corpus, profile string, n int) (vals []uint64, offs []int) {
	vals, offs = keys.Values(n, 0xFA11)
	if c.Natural != nil {
		vals, offs = nil, make([]int, n+1)
		for i, vs := range c.Natural[:n] {
			vals = append(vals, vs...)
			offs[i+1] = len(vals)
		}
	}
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
		case baseline:
			baseKit.apply(f.base)(f.ck.B, rest)
		}
		*j = 0
	}
}

// ranges picks the probe ranges: rangeKeys consecutive keys in key order,
// starting at random keys.
// prefixes returns the bounds of prefix searches: the prefixes of random keys
// (see keys.Prefix), so that frequent prefixes are searched as often as users
// would search them, and the end of each prefix's range (see keys.PrefixEnd).
func prefixes(kind keys.Kind, all keys.Set, n int) (from, to keys.Set) {
	rng := rtcompare.NewDPRNG(0x9F1F)
	var f, t [][]byte
	for range min(n, 1<<16) {
		k := all.B[rng.Uint64()%uint64(n)]
		p := keys.Prefix(kind, k)
		f, t = append(f, p), append(t, keys.PrefixEnd(p))
	}
	return keys.Pack(f), keys.Pack(t)
}

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
