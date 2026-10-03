package lpage

import (
	"bytes"
	"sort"
)

// Run is a sorted sequence of pages: the leaf level of a B+ tree without the
// levels above, standing in for the range nodes of the tree in the experiments
// and the tests, as vpage.Run does for the other candidate. It finds the page
// of a suffix by its separators and splits a page that is full.
type Run struct {
	pages []*Page
	seps  [][]byte // seps[i] is the first suffix of pages[i+1] when it was split off
}

func (r *Run) page(s []byte) int {
	return sort.Search(len(r.seps), func(i int) bool { return bytes.Compare(r.seps[i], s) > 0 })
}

// Pages returns the pages in order.
func (r *Run) Pages() []*Page { return r.pages }

// Len returns the number of keys.
func (r *Run) Len() int {
	n := 0
	for _, p := range r.pages {
		n += p.Len()
	}
	return n
}

// Get returns the value of suffix s.
func (r *Run) Get(s []byte) (uint64, bool) {
	if len(r.pages) == 0 {
		return 0, false
	}
	return r.pages[r.page(s)].Get(s)
}

// Insert sets the value of suffix s.
func (r *Run) Insert(s []byte, v uint64) error {
	switch {
	case len(s) == 0:
		return ErrEmpty
	case len(s) > maxSuffix:
		return ErrTooLong
	}
	if len(r.pages) == 0 {
		p, _ := Build([][]byte{s}, []uint64{v}) // one suffix of at most 255 bytes fits
		r.pages = []*Page{p}
		return nil
	}
	i := r.page(s)
	p, res, _ := r.pages[i].Insert(s, v) // s is not too long: checked above
	if res != Full {
		r.pages[i] = p
		return nil
	}
	var left, right *Page
	placed := false // s is in one of the new pages already
	if p.Len() < 2 {
		// one long suffix fills the page: s goes into a page of its own
		q, _ := Build([][]byte{s}, []uint64{v})
		first, _ := p.Entries()
		left, right, placed = p, q, true
		if bytes.Compare(s, first[0]) < 0 {
			left, right = q, p
		}
	} else {
		left, right = p.Split()
	}
	keys, _ := right.Entries()
	r.pages[i] = left
	r.pages = append(r.pages, nil)
	copy(r.pages[i+2:], r.pages[i+1:])
	r.pages[i+1] = right
	r.seps = append(r.seps, nil)
	copy(r.seps[i+1:], r.seps[i:])
	r.seps[i] = keys[0]
	if placed {
		return nil
	}
	return r.Insert(s, v)
}

// Delete removes suffix s and reports whether it was there. Pages are not merged.
func (r *Run) Delete(s []byte) bool {
	if len(r.pages) == 0 {
		return false
	}
	i := r.page(s)
	p, ok := r.pages[i].Delete(s)
	switch {
	case !ok:
		return false
	case p != nil:
		r.pages[i] = p
		return true
	}
	r.pages = append(r.pages[:i], r.pages[i+1:]...)
	if len(r.seps) > 0 {
		j := min(i, len(r.seps)-1)
		r.seps = append(r.seps[:j], r.seps[j+1:]...)
	}
	return true
}
