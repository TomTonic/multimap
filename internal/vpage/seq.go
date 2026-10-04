package vpage

import (
	"bytes"
	"sort"
)

// Run is a sorted sequence of pages: the leaf level of a B+ tree, without the
// levels above. It stands in for the range nodes of the tree in the
// experiments of PLAN step 1 and in the tests: it finds the page of a remainder
// by its separators, splits a page that is full and merges thin neighbours.
// The tree splits at a byte boundary, and keeps its own fill; the
// experiments approximate that.
type Run struct {
	pages []*Page
	seps  [][]byte // seps[i] is the first remainder of pages[i+1] when it was split off
}

// page returns the index of the page that remainder s belongs to.
func (r *Run) page(s []byte) int {
	return sort.Search(len(r.seps), func(i int) bool { return bytes.Compare(r.seps[i], s) > 0 })
}

// Len returns the number of keys.
func (r *Run) Len() int {
	n := 0
	for _, p := range r.pages {
		n += p.Len()
	}
	return n
}

// PageFor returns the page that remainder s belongs to, or nil if the run is empty.
// Benchmarks use it to take the choice of the page out of the timed part.
func (r *Run) PageFor(s []byte) *Page {
	if len(r.pages) == 0 {
		return nil
	}
	return r.pages[r.page(s)]
}

// Pages returns the pages in order.
func (r *Run) Pages() []*Page { return r.pages }

// Get returns the value of remainder s.
func (r *Run) Get(s []byte) (uint64, bool) {
	if len(r.pages) == 0 {
		return 0, false
	}
	return r.pages[r.page(s)].Get(s)
}

// Insert sets the value of remainder s.
func (r *Run) Insert(s []byte, v uint64) error {
	if len(s) > maxRemainder {
		return ErrTooLong
	}
	if len(r.pages) == 0 {
		p := New(0, 0)
		p, _, _ = p.Insert(s, v)
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
		// one long key fills the page: s goes into a page of its own, before or after it
		q, _, _ := New(0, 0).Insert(s, v)
		var buf [maxRemainder]byte
		left, right, placed = p, q, true
		if bytes.Compare(s, p.Key(0, &buf)) < 0 {
			left, right = q, p
		}
	} else {
		left, right, _ = p.Split() // at least two keys: it cannot fail
	}
	var sep [maxRemainder]byte
	sepKey := append([]byte(nil), right.Key(0, &sep)...)
	r.pages[i] = left
	r.pages = append(r.pages, nil)
	copy(r.pages[i+2:], r.pages[i+1:])
	r.pages[i+1] = right
	r.seps = append(r.seps, nil)
	copy(r.seps[i+1:], r.seps[i:])
	r.seps[i] = sepKey
	if placed {
		return nil
	}
	return r.Insert(s, v)
}

// Delete removes remainder s and reports whether it was there. A page that has
// become thin is merged into a neighbour if the two fit one page.
func (r *Run) Delete(s []byte) bool {
	if len(r.pages) == 0 {
		return false
	}
	i := r.page(s)
	p, ok := r.pages[i].Delete(s)
	if !ok {
		return false
	}
	if p == nil {
		r.removePage(i)
		return true
	}
	r.pages[i] = p
	r.mergeAround(i)
	return true
}

// removePage drops page i, which is empty, with the separator that led to it.
func (r *Run) removePage(i int) {
	r.pages = append(r.pages[:i], r.pages[i+1:]...)
	if len(r.seps) > 0 {
		j := min(i, len(r.seps)-1)
		r.seps = append(r.seps[:j], r.seps[j+1:]...)
	}
}

// mergeAround merges page i with its predecessor or successor, if they fit one page.
func (r *Run) mergeAround(i int) {
	for _, a := range []int{i - 1, i} {
		if a < 0 || a+1 >= len(r.pages) {
			continue
		}
		if m := Merge(r.pages[a], r.pages[a+1]); m != nil {
			r.pages[a] = m
			r.pages = append(r.pages[:a+1], r.pages[a+2:]...)
			r.seps = append(r.seps[:a], r.seps[a+1:]...)
			return
		}
	}
}

// Each calls fn for the remainders from from on, in order, until fn returns false.
func (r *Run) Each(from []byte, fn func(s []byte, v uint64) bool) {
	if len(r.pages) == 0 {
		return
	}
	start := 0
	if from != nil {
		start = r.page(from)
	}
	for i := start; i < len(r.pages); i++ {
		stop := false
		r.pages[i].Each(from, func(s []byte, v uint64) bool {
			if !fn(s, v) {
				stop = true
			}
			return !stop
		})
		if stop {
			return
		}
		from = nil
	}
}
