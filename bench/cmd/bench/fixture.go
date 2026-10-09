package main

import (
	"slices"
	"strings"

	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/rtcompare"
	"github.com/TomTonic/rtcompare/multiproc"
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

// The value profiles: natural gives keys a skewed number of values (see
// keys.Values), single-value exactly one value per key.
const (
	natural     = "natural"
	singleValue = "single-value"
)

// vsOnly, if not empty, limits the candidates ordered is compared with to
// these (-vs).
var vsOnly []string

// implsFor returns the candidates compared under a value profile: ordered
// first, then those it is compared with, including the baseline in a bench
// built with it (see kit.go), all of them limited by vsOnly.
func implsFor(profile string) []string {
	others := []string{hashed, btreeSets, mapSets}
	if profile == singleValue {
		others = []string{btreeMapC}
	}
	if baseKit != nil {
		others = append(others, baseline)
	}
	others = append(others, mkNames...) // only if -vs names them
	out := []string{ordered}
	for _, b := range others {
		if (len(vsOnly) == 0 && mkKits[b] == nil) || slices.Contains(vsOnly, b) {
			out = append(out, b)
		}
	}
	return out
}

type (
	btreeMM  = btree.Map[string, map[V]struct{}]
	mapMM    = map[string]map[V]struct{}
	btreeMap = btree.Map[string, V]
)

// fixture holds one scenario's corpus and every candidate built from it.
type fixture struct {
	c       keys.Corpus
	profile string
	vals    []V
	offs    []int
	ord     *multimap.Ordered[V]
	hsh     *multimap.Hashed[V]
	bt      *btreeMM
	gm      mapMM
	bm      *btreeMap
	kitMaps map[string]any // the candidates with a kit (the baseline, ordered-mkN), see kit.go
	others  []string       // the candidates built besides ordered
	// ranges of rangeKeys consecutive keys, as []byte and string views, and
	// the ranges of the keys that start with the prefix of a random key (see
	// keys.Prefix; text keys only)
	from, to   keys.Set
	pfrom, pto keys.Set
	// index workloads (see workload.go): churn keys (corpus keys, then extra
	// keys that only the workloads use), the shape of their streams, and the
	// key-value pair of every workload element
	ck     keys.Set
	stream stream
	pairs  pairs
}

const rangeKeys = 100

// newFixture generates the corpus and its values under the value profile and
// builds the candidates named in impls, one after another, in the order
// arrange leaves them in. Whichever is built last can be consistently a few
// percent faster, so the speed processes vary the order (see buildOrder).
func newFixture(kind keys.Kind, n int, profile string, impls []string, st stream, arrange func([]string)) *fixture {
	f := &fixture{c: keys.Generate(kind, n, 0x5EED), profile: profile}
	nums, offs := profileValues(f.c, profile, n)
	f.vals, f.offs = toVs(nums, f.c.Names), offs
	builds := map[string]func(){
		ordered:   func() { f.ord = buildOrdered(f.c.Keys.B, f.vals, f.offs) },
		hashed:    func() { f.hsh = buildHashed(f.c.Keys.B, f.vals, f.offs) },
		btreeSets: func() { f.bt = buildBtree(f.c.Keys.S, f.vals, f.offs) },
		mapSets:   func() { f.gm = buildMap(f.c.Keys.S, f.vals, f.offs) },
		btreeMapC: func() { f.bm = buildBtreeMap(f.c.Keys.S, f.vals, f.offs) },
	}
	f.kitMaps = map[string]any{}
	for _, name := range append([]string{baseline}, mkNames...) {
		if k := kitOf(name); k != nil {
			builds[name] = func() { f.kitMaps[name] = k.build(f.c.Keys.B, f.vals, f.offs) }
		}
	}
	f.others = slices.DeleteFunc(slices.Clone(impls), func(s string) bool { return s == ordered })
	order := append([]string(nil), impls...)
	arrange(order)
	for _, name := range order {
		builds[name]()
	}
	f.from, f.to = ranges(f.c.Keys, n)
	if keys.Text(kind) {
		f.pfrom, f.pto = prefixes(kind, f.c.Keys, n)
	}
	f.ck = keys.Pack(append(slices.Clone(f.c.Keys.B), f.c.Misses.B[:extraKeys(profile, n)]...))
	f.stream = st
	f.pairs = newPairs(n, f.vals, f.offs, st.ratio, profile == singleValue)
	return f
}

// buildOrder returns how process p of a speed scenario orders the builds: two
// candidates alternate by process, which balances exactly over a pair of
// processes, and more are shuffled from the process's seed.
func buildOrder(p *multiproc.Process) func([]string) {
	return func(order []string) {
		switch {
		case len(order) == 2 && p.Index%2 == 1:
			order[0], order[1] = order[1], order[0]
		case len(order) > 2:
			p.Rand().Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
		}
	}
}

// profileValues returns the value numbers of the n keys of c under a value
// profile: key i holds vals[offs[i]:offs[i+1]], which toVs turns into values. Under natural, keys with natural values
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
	if profile != singleValue {
		return vals, offs
	}
	// one value per key: the first of each key's natural values
	one, idx := make([]uint64, n), make([]int, n+1)
	for i := range n {
		one[i], idx[i+1] = vals[offs[i]], i+1
	}
	return one, idx
}

// extraKeys is the number of keys that only the index workloads use: n/2 for
// natural, where transient values also go to corpus keys; n for single-value, where
// each transient value needs a key of its own (see newPairs).
func extraKeys(profile string, n int) int {
	if profile == singleValue {
		return n
	}
	return max(1, n/2)
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

func buildOrdered(k [][]byte, vals []V, offs []int) *multimap.Ordered[V] {
	m := multimap.NewOrdered[V]()
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
	return m
}

func buildHashed(k [][]byte, vals []V, offs []int) *multimap.Hashed[V] {
	m := multimap.NewHashed[V]()
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
	return m
}

func buildBtree(k []string, vals []V, offs []int) *btreeMM {
	m := &btreeMM{}
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			btreeAdd(m, key, v)
		}
	}
	return m
}

func buildMap(k []string, vals []V, offs []int) mapMM {
	m := mapMM{}
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			mapAdd(m, key, v)
		}
	}
	return m
}

func buildBtreeMap(k []string, vals []V, offs []int) *btreeMap {
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

func btreeAdd(m *btreeMM, k string, v V) {
	s, ok := m.Get(k)
	if !ok {
		s = map[V]struct{}{}
		m.Set(strings.Clone(k), s)
	}
	s[v] = struct{}{}
}

func btreeRemove(m *btreeMM, k string, v V) {
	if s, ok := m.Get(k); ok {
		delete(s, v)
		if len(s) == 0 {
			m.Delete(k)
		}
	}
}

func mapAdd(m mapMM, k string, v V) {
	s := m[k]
	if s == nil {
		s = map[V]struct{}{}
		m[strings.Clone(k)] = s
	}
	s[v] = struct{}{}
}

func mapRemove(m mapMM, k string, v V) {
	if s := m[k]; s != nil {
		delete(s, v)
		if len(s) == 0 {
			delete(m, k)
		}
	}
}
