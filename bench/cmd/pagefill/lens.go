package main

import (
	"bytes"
	"fmt"
	"slices"
	"sort"

	"github.com/TomTonic/multimap/bench/keys"
)

// This file models the second candidate for the multi-key page (docs/redesign,
// PLAN step 3): the page of the user's sketch in whataleafneedstostore.md, whose
// header lists the lengths of its entries instead of a directory of tags:
//
//	header | common prefix | remainder 1 | remainder 2 | ... | value 1 | value 2 | ...
//
// The header is a type byte, the length of the common prefix, one length byte
// per remainder and, if the values have different lengths, one per value,
// rounded up to a multiple of 8. It decides how many entries a page holds: 6 for
// a header of 8 bytes (3 with value lengths), 14 for 16 bytes, and so on.
//
// Only the sizes are modelled, not the bytes: a page is the sorted list of its
// stripped suffixes, its size is what the sketch needs for them, and a page that
// has too many entries or does not fit the largest class is split in the middle,
// as the page of step 2 does after random inserts. That is enough for the bytes
// a key costs, which is what decides whether the layout deserves a real
// implementation.

// lensLayout describes one variant of the sketched page.
type lensLayout struct {
	fixed     bool  // all values have the same length: no length byte per value
	value     int   // bytes of a value
	maxHeader int   // largest header in bytes, a multiple of 8
	classes   []int // object sizes
}

// capacity returns how many entries a header of h bytes has room for.
func (l lensLayout) capacity(h int) int {
	if l.fixed {
		return h - 2
	}
	return (h - 2) / 2
}

// header returns the smallest header that has room for n entries, or 0 if no
// allowed header has.
func (l lensLayout) header(n int) int {
	for h := 8; h <= l.maxHeader; h += 8 {
		if l.capacity(h) >= n {
			return h
		}
	}
	return 0
}

// lensSize describes a page of the sketch holding keys.
type lensSize struct {
	header, prefix, remainders, size int
}

// size returns what the sketch needs for the sorted suffixes keys, and whether
// it can hold them: a remainder takes one to 255 bytes after the common prefix,
// which takes up to 255, so the common prefix stops one byte before the
// shortest suffix ends.
func (l lensLayout) size(keys [][]byte) (lensSize, bool) {
	n := len(keys)
	h := l.header(n)
	if h == 0 {
		return lensSize{}, false
	}
	cp := 0
	if n > 1 {
		shortest := len(keys[0])
		for _, k := range keys {
			shortest = min(shortest, len(k))
		}
		cp = min(lcp(keys[0], keys[n-1]), shortest-1, 255)
	}
	rem := 0
	for _, k := range keys {
		if len(k)-cp > 255 || len(k)-cp < 1 {
			return lensSize{}, false
		}
		rem += len(k) - cp
	}
	return lensSize{h, cp, rem, h + cp + rem + n*l.value}, true
}

// class returns the smallest object size that holds size bytes, or 0.
func (l lensLayout) class(size int) int {
	for _, c := range l.classes {
		if size <= c {
			return c
		}
	}
	return 0
}

// lensPage is the sorted suffixes of one page.
type lensPage struct{ keys [][]byte }

// lensRun is a sorted sequence of pages, as vpage.Run is for the page of step 2.
type lensRun struct {
	l     lensLayout
	pages []*lensPage
	seps  [][]byte // seps[i] is the first suffix of pages[i+1]
}

func (r *lensRun) page(s []byte) int {
	return sort.Search(len(r.seps), func(i int) bool { return bytes.Compare(r.seps[i], s) > 0 })
}

// fits reports whether the keys fit a page of the largest class.
func (r *lensRun) fits(keys [][]byte) bool {
	sz, ok := r.l.size(keys)
	return ok && r.l.class(sz.size) != 0
}

// insert adds suffix s (a copy is not made); it reports false for a suffix the
// sketch cannot hold alone.
func (r *lensRun) insert(s []byte) bool {
	if !r.fits([][]byte{s}) {
		return false
	}
	if len(r.pages) == 0 {
		r.pages = []*lensPage{{keys: [][]byte{s}}}
		return true
	}
	i := r.page(s)
	p := r.pages[i]
	at, found := slices.BinarySearchFunc(p.keys, s, bytes.Compare)
	if found {
		return true
	}
	keys := slices.Insert(slices.Clone(p.keys), at, s)
	if r.fits(keys) {
		p.keys = keys
		return true
	}
	// Split in the middle by count, until both halves fit; a single suffix always fits.
	m := len(keys) / 2
	left, right := &lensPage{keys: keys[:m]}, &lensPage{keys: keys[m:]}
	r.pages = slices.Insert(r.pages, i+1, right)
	r.pages[i] = left
	r.seps = slices.Insert(r.seps, i, right.keys[0])
	for _, q := range []*lensPage{left, right} {
		for !r.fits(q.keys) && len(q.keys) > 1 { // a half too big: only for suffixes that fill a page
			j := slices.Index(r.pages, q)
			mm := len(q.keys) / 2
			nq := &lensPage{keys: q.keys[mm:]}
			q.keys = q.keys[:mm]
			r.pages = slices.Insert(r.pages, j+1, nq)
			r.seps = slices.Insert(r.seps, j, nq.keys[0])
		}
	}
	return true
}

// lensModel is the tree around the pages of the sketch.
type lensModel struct {
	*model
	l    lensLayout
	runs []lensRun
}

func newLensModel(kind keys.Kind, n, chunk int, l lensLayout) *lensModel {
	m := &lensModel{model: newModel(kind, n, chunk), l: l}
	m.runs = make([]lensRun, len(m.bounds))
	for i := range m.runs {
		m.runs[i].l = l
	}
	return m
}

// lensRow measures a layout for the keys of a kind and formats the row of the
// table: as the row of the page of step 2 (measure), with the header and the
// common prefix in place of the uniform pages and the tails, and the share of
// the pages whose header, common prefix and remainders end within the first 128
// bytes, so that a lookup finds the entry in the first round (R5).
func lensRow(kind keys.Kind, n, chunk int, l lensLayout) string {
	m := newLensModel(kind, n, chunk, l)
	tooLong, suffixBytes := 0, 0
	for _, key := range m.keys {
		ci := m.chunk(key)
		if !m.runs[ci].insert(key[m.bases[ci]:]) {
			tooLong++
			continue
		}
		suffixBytes += len(key) - m.bases[ci]
	}
	var pages, keysIn, size, used, router, hdr, cp, scan, early int
	for i := range m.runs {
		router += routerBytes(len(m.runs[i].pages))
		for _, p := range m.runs[i].pages {
			sz, _ := l.size(p.keys)
			pages++
			keysIn += len(p.keys)
			size += l.class(sz.size)
			used += sz.size
			hdr += sz.header
			cp += sz.prefix
			scan += sz.header + sz.prefix + sz.remainders
			if sz.header+sz.prefix+sz.remainders <= 128 {
				early++
			}
		}
	}
	f := func(x int) float64 { return float64(x) / float64(keysIn) }
	return fmt.Sprintf("| %s | %d | %.1f | %d | %.1f | %.0f %% | %.1f | %.1f | %.1f | %.1f | %.1f | %.0f | %.0f %% | %d |",
		kind, keysIn, f(suffixBytes), pages, float64(keysIn)/float64(pages), 100*float64(used)/float64(size),
		f(size), f(router), f(size+router), float64(hdr)/float64(pages), float64(cp)/float64(pages),
		float64(scan)/float64(pages), 100*float64(early)/float64(pages), tooLong)
}

const lensTableHeader = "| kind | keys | suffix B | pages | keys/page | fill | page B/key | router B/key | total B/key | header B/page | prefix B/page | keys end at B | keys end within 128 B | too long |\n|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|"
