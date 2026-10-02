package vpage

import "errors"

// The policies of the page, in percent of a class's size. They are variables
// so that the experiments of PLAN step 1 can try other values.
var (
	// SplitFill is how full a half of a split page may be: the class that
	// holds it with this much to spare is chosen.
	SplitFill = 85
	// ShrinkFill is how full a page may be in the next smaller class for it to
	// move there after a removal.
	ShrinkFill = 70
	// MergeFill is how full the two pages merged into one may be.
	MergeFill = 60
	// SplitByBytes splits a page where the bytes of its entries are halved,
	// not where its count is.
	SplitByBytes = false
)

// Result says what Insert did.
type Result int

const (
	Inserted Result = iota
	Updated
	Full // the page is of the largest class and has no room: split it
)

// plan describes a page to be built: its entries and the bytes of their
// tails, and the length all suffixes have, if they have one of 1 to 8.
type plan struct {
	n, tails, ulen int
}

// need returns the bytes of a page of plan pl, without spare slots.
func need(pl plan) int {
	return arraysEnd(pl.n, pl.ulen != 0) + pl.tails
}

// planFor returns the plan of the entries from to to of p and, if extra, of one
// more suffix s.
func (p *Page) planFor(from, to int, s []byte, extra bool) plan {
	pl := plan{n: to - from}
	uni := -1 // the length all entries have; -2: they differ
	see := func(l int) {
		if l > headLen {
			pl.tails += l - headLen
		}
		switch {
		case uni == -1:
			uni = l
		case uni != l:
			uni = -2
		}
	}
	for i := from; i < to; i++ {
		see(p.length(i))
	}
	if extra {
		pl.n++
		see(len(s))
	}
	if uni >= 1 && uni <= headLen {
		pl.ulen = uni
	}
	return pl
}

// fitClass returns the smallest class of at least min that holds pl with
// fill percent of its size, or -1.
func fitClass(pl plan, min, fill int) int {
	for c := min; c <= MaxClass; c++ {
		if need(pl)*100 <= fill*sizes[c] {
			return c
		}
	}
	return -1
}

// newPage returns an empty page of class c laid out for plan pl: its capacity
// is the entries of pl and as many more as fit with tails of the average length.
func newPage(c int, pl plan) *Page {
	q := alloc(c)
	q.class, q.ulen = uint8(c), uint8(pl.ulen)
	size := sizes[c]
	q.top = uint16(size)
	if pl.ulen != 0 {
		k := (size - hdr - 7) / 17
		for ; arraysEnd(k+1, true) <= size; k++ {
		}
		q.cap = uint8(k)
	} else {
		avg := 12 // an empty page: assume tails of this length
		if pl.n > 0 {
			avg = (pl.tails + pl.n - 1) / pl.n
		}
		k := pl.n + (size-need(pl))/(20+avg)
		for k = min(k, 255); arraysEnd(k, false)+pl.tails > size; k-- {
		}
		q.cap = uint8(max(k, 1))
	}
	h := q.heads()
	for i := range h {
		h[i] = pad
	}
	return q
}

// New returns an empty page of class c for suffixes of length ulen, 1 to 8,
// or of any length if ulen is 0.
func New(c, ulen int) *Page { return newPage(c, plan{ulen: ulen}) }

// appendEntry adds an entry after the last; the page must have room.
func (q *Page) appendEntry(w, v uint64, l int, tail []byte) {
	n := int(q.count)
	q.heads()[n], q.vals()[n] = w, v
	if q.ulen != 0 {
		q.tags()[n] = tag(w)
	} else {
		off := 0
		if len(tail) > 0 {
			q.top -= uint16(len(tail))
			copy(q.mem()[q.top:], tail)
			off = int(q.top)
		}
		q.fat()[n] = entry(tag(w), l, off)
	}
	q.count++
}

// insertion is a suffix that rebuild adds in front of the entry at position pos.
type insertion struct {
	pos int
	s   []byte
	v   uint64
}

// rebuild returns a page of class c and plan pl that holds the entries from to
// to of p and, if in is not nil, its suffix, in order. They must fit.
func (p *Page) rebuild(c int, pl plan, from, to int, in *insertion) *Page {
	q := newPage(c, pl)
	add := func() {
		var tail []byte
		if len(in.s) > headLen {
			tail = in.s[headLen:]
		}
		q.appendEntry(word(in.s), in.v, len(in.s), tail)
	}
	for i := from; i < to; i++ {
		if in != nil && in.pos == i {
			add()
		}
		q.appendEntry(p.heads()[i], p.vals()[i], p.length(i), p.tail(i))
	}
	if in != nil && in.pos == to {
		add()
	}
	return q
}

// Insert sets the value of suffix s, adding it if it is not there. It returns
// the page, or the larger page that replaces it; the old page must not be used
// any more then. Full means that nothing changed because the page holds as much
// as its largest class can.
func (p *Page) Insert(s []byte, v uint64) (*Page, Result, error) {
	if len(s) > maxSuffix {
		return p, Inserted, ErrTooLong
	}
	w := word(s)
	i, found := p.find(s, w)
	if found {
		p.vals()[i] = v
		return p, Updated, nil
	}
	tail := max(len(s)-headLen, 0)
	fits := p.ulen == 0 || len(s) == int(p.ulen)
	if fits && int(p.count) < int(p.cap) && (p.ulen != 0 || tail <= p.heapFree()) {
		p.insertAt(i, w, v, s)
		return p, Inserted, nil
	}
	pl := p.planFor(0, int(p.count), s, true)
	c := fitClass(pl, int(p.class), 100)
	if c < 0 {
		return p, Full, nil
	}
	return p.rebuild(c, pl, 0, int(p.count), &insertion{i, s, v}), Inserted, nil
}

// insertAt puts suffix s at position i of a page with room for it.
func (p *Page) insertAt(i int, w, v uint64, s []byte) {
	n := int(p.count)
	h, vs := p.heads(), p.vals()
	copy(h[i+1:n+1], h[i:n])
	copy(vs[i+1:n+1], vs[i:n])
	h[i], vs[i] = w, v
	if p.ulen != 0 {
		tg := p.tags()
		copy(tg[i+1:n+1], tg[i:n])
		tg[i] = tag(w)
	} else {
		f := p.fat()
		copy(f[i+1:n+1], f[i:n])
		off := 0
		if len(s) > headLen {
			p.top -= uint16(len(s) - headLen)
			copy(p.mem()[p.top:], s[headLen:])
			off = int(p.top)
		}
		f[i] = entry(tag(w), len(s), off)
	}
	p.count++
}

// Delete removes suffix s. It returns the page, or the smaller page that
// replaces it, or nil once the page is empty, and whether s was there.
func (p *Page) Delete(s []byte) (*Page, bool) {
	if len(s) > maxSuffix {
		return p, false
	}
	i, found := p.find(s, word(s))
	if !found {
		return p, false
	}
	n := int(p.count)
	if n == 1 {
		return nil, true
	}
	h, vs := p.heads(), p.vals()
	copy(h[i:n-1], h[i+1:n])
	copy(vs[i:n-1], vs[i+1:n])
	h[n-1] = pad
	if p.ulen != 0 {
		tg := p.tags()
		copy(tg[i:n-1], tg[i+1:n])
	} else {
		f := p.fat()
		copy(f[i:n-1], f[i+1:n])
	}
	p.count--
	if p.class > 0 {
		pl := p.planFor(0, n-1, nil, false)
		if need(pl)*100 <= ShrinkFill*sizes[p.class-1] {
			return p.rebuild(int(p.class)-1, pl, 0, n-1, nil), true
		}
	}
	return p, true
}

// cost is what the entry at position i takes.
func (p *Page) cost(i int) int {
	if p.ulen != 0 {
		return 17
	}
	return 20 + max(p.length(i)-headLen, 0)
}

// Split divides a page of at least two keys into the page of its first keys and
// the page of the others, at the middle of the count or, with SplitByBytes, of
// the bytes. The first suffix of the second page is the separator.
func (p *Page) Split() (left, right *Page, err error) {
	n := int(p.count)
	if n < 2 {
		return nil, nil, errors.New("vpage: cannot split a page of fewer than two keys")
	}
	m := n / 2
	if SplitByBytes {
		total := 0
		for i := range n {
			total += p.cost(i)
		}
		acc := 0
		for m = 1; m < n-1; m++ {
			if acc += p.cost(m - 1); 2*acc >= total {
				break
			}
		}
	}
	half := func(from, to int) *Page {
		pl := p.planFor(from, to, nil, false)
		c := fitClass(pl, 0, SplitFill)
		if c < 0 { // two long suffixes: no room to spare, but a half never needs more than the page had
			c = fitClass(pl, 0, 100)
		}
		return p.rebuild(c, pl, from, to, nil)
	}
	return half(0, m), half(m, n), nil
}

// Merge returns the page that holds the keys of a and then those of b, which
// must all be above a's, if it holds them with at most MergeFill percent of
// its class, else nil.
func Merge(a, b *Page) *Page {
	pa, pb := a.planFor(0, int(a.count), nil, false), b.planFor(0, int(b.count), nil, false)
	pl := plan{n: pa.n + pb.n, tails: pa.tails + pb.tails}
	if pa.ulen == pb.ulen {
		pl.ulen = pa.ulen
	}
	c := fitClass(pl, 0, MergeFill)
	if c < 0 {
		return nil
	}
	q := newPage(c, pl)
	for _, src := range []*Page{a, b} {
		for i := range int(src.count) {
			q.appendEntry(src.heads()[i], src.vals()[i], src.length(i), src.tail(i))
		}
	}
	return q
}

// Each calls fn for the suffixes from from on, in order, with their values,
// until fn returns false. A nil from starts at the first.
func (p *Page) Each(from []byte, fn func(s []byte, v uint64) bool) {
	i := 0
	if from != nil {
		i, _ = p.find(from, word(from))
	}
	var buf [maxSuffix]byte
	for ; i < int(p.count); i++ {
		if !fn(p.Key(i, &buf), p.vals()[i]) {
			return
		}
	}
}

// Stats describes a page for the experiments of PLAN step 1.
type Stats struct {
	Keys    int  // keys in the page
	Size    int  // bytes of the object
	Used    int  // bytes the keys need: what a page of just this size would be
	Tails   int  // bytes of tails among them
	Uniform bool // the uniform flavor
}

// Stats returns the statistics of the page.
func (p *Page) Stats() Stats {
	pl := p.planFor(0, int(p.count), nil, false)
	return Stats{Keys: pl.n, Size: p.Size(), Used: need(pl), Tails: pl.tails, Uniform: p.ulen != 0}
}
