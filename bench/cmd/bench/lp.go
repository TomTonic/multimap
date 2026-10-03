package main

import (
	"iter"
	"os"
	"strconv"

	"github.com/TomTonic/multimap/internal/artstr"
	"github.com/TomTonic/multimap/internal/lpage"
)

// The environment variable LPAGE_MAXHEADER sets the largest header of a page in
// bytes (a multiple of 8, 8 to 64; the default is 48: 23 entries with values of
// different lengths), LPAGE_MINHEADER the smallest (default 8). Child processes inherit it, so a whole run uses one value.
func init() {
	if n, err := strconv.Atoi(os.Getenv("ARTSTR_CROWDED")); err == nil && n >= 0 {
		artstr.CrowdedRatio = n
	}
	if n, err := strconv.Atoi(os.Getenv("LPAGE_MAXHEADER")); err == nil && n >= 8 && n <= 64 && n%8 == 0 {
		lpage.MaxHeader = n
	}
	if n, err := strconv.Atoi(os.Getenv("LPAGE_MINHEADER")); err == nil && n >= 8 && n <= 64 && n%8 == 0 {
		lpage.MinHeader = n
	}
}

// lpMap is the candidate "ordered-lpage": the multimap of internal/artstr, the
// tree whose keys with one string value live in length-header pages
// (internal/lpage), behind the same methods and iterators as multimap.Ordered
// so that a batch of either costs the same to call. It is an experiment of
// step 3 of docs/redesign and exists only for strings: with uint64 values
// (hasPages is false) the bench never builds it.
type lpMap struct{ m artstr.Map[V] }

func newLP(zeroCopy, pairs bool) *lpMap {
	l := &lpMap{}
	l.m.ZeroCopy, l.m.Pairs = zeroCopy, pairs
	return l
}

// lpImpls are the candidates with pages for string values.
var lpImpls = []string{orderedLP, orderedLPZ, orderedLPM, orderedLPMZ}

// lpOptions returns how candidate impl, one of lpImpls, is set up: with
// immutable pages and views of them as strings, and with multi-value entries in
// the pages.
func lpOptions(impl string) (zeroCopy, pairs bool) {
	return impl == orderedLPZ || impl == orderedLPMZ, impl == orderedLPM || impl == orderedLPMZ
}

func (l *lpMap) AddValue(key []byte, v V)    { l.m.Add(key, v) }
func (l *lpMap) RemoveValue(key []byte, v V) { l.m.Remove(key, v) }
func (l *lpMap) RemoveKey(key []byte)        { l.m.RemoveKey(key) }
func (l *lpMap) NumberOfKeys() int           { return l.m.Len() }

func (l *lpMap) ValuesForSeq(key []byte) iter.Seq[V] {
	return func(yield func(V) bool) { l.m.Each(key, yield) }
}

// ValuesBetweenInclusiveSeq iterates over the values of all keys in [from, to].
func (l *lpMap) ValuesBetweenInclusiveSeq(from, to []byte) iter.Seq[V] {
	b := artstr.Bounds{From: from, To: to, HasFrom: true, HasTo: true, FromIncl: true, ToIncl: true}
	return func(yield func(V) bool) { l.m.RangeValues(&b, yield) }
}

func buildLP(k [][]byte, vals []V, offs []int, impl string) *lpMap {
	m := newLP(lpOptions(impl))
	for i, key := range k {
		for _, v := range vals[offs[i]:offs[i+1]] {
			m.AddValue(key, v)
		}
	}
	return m
}
