package vpage

import (
	"errors"
	"math/bits"
)

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
	// MinPrefix (at least 1) is the shortest prefix a page stores once instead of in each
	// key; above 255 it turns the prefix off. A rebuild takes what all keys share.
	MinPrefix = 1
	// MinGain is the bytes a prefix must save a page, after the bytes it takes
	// (its length rounded up to a multiple of 8), for the page to have one: a
	// prefix shortens the tails, which only keys longer than a head have, and may
	// make the page uniform. A prefix that does not pay costs the lookup its
	// shifts and a rebuild that changes it cuts all keys anew.
	MinGain = 16
	// PrefixSlack makes a rebuild keep this many bytes less than the keys
	// share, so that keys that differ a little later do not shorten it again.
	PrefixSlack = 0
)

// Counts is what the pages did since the program started, for the experiments
// of PLAN step 1 (not safe for concurrent use).
var Counts struct {
	Rebuilds      int // pages built anew: growth, shrink, split, merge, a shorter prefix
	PrefixChanges int // of them, those that had to cut every key anew because the prefix changed
	Merges        int
	Shrinks       int // of them, by a delete that let the page move to a smaller class
}

// Result says what Insert did.
type Result int

const (
	Inserted Result = iota
	Updated
	Full // the page is of the largest class and has no room: split it
)

// plan describes a page to be built: its entries and the bytes of their
// tails, the length of the prefix they share (the prefix itself is cut from the
// first key when the page is built), and the length all suffixes have after
// it, if they have one of 1 to 8.
type plan struct {
	n, tails, ulen, plen int
}

// need returns the bytes of a page of plan pl, without spare slots.
func need(pl plan) int {
	return arraysEnd(pl.n, pl.ulen != 0, pl.plen) + pl.tails
}

// shape returns the bytes of the tails and the common length of the suffixes
// of the entries from to to of p once the first q bytes of each are cut off
// (-2 if they differ, -1 if there are none).
func (p *Page) shape(from, to, q int) (tails, uni int) {
	uni = -1
	for i := from; i < to; i++ {
		tails, uni = see(tails, uni, int(p.plen)+p.length(i)-q)
	}
	return tails, uni
}

// see adds a suffix of length l to the tails and the common length.
func see(tails, uni, l int) (int, int) {
	if l > headLen {
		tails += l - headLen
	}
	switch {
	case uni == -1:
		uni = l
	case uni != l:
		uni = -2
	}
	return tails, uni
}

// lcpEntries returns the length of the longest common prefix of the suffixes
// at positions i and j, not counting the page's prefix; it compares head words,
// and the tails only if the heads are equal.
func (p *Page) lcpEntries(i, j int) int {
	m := min(p.length(i), p.length(j))
	if x := p.heads()[i] ^ p.heads()[j]; x != 0 {
		return min(bits.LeadingZeros64(x)/8, m)
	}
	if m <= headLen {
		return m
	}
	return headLen + lcp(p.tail(i)[:m-headLen], p.tail(j)[:m-headLen])
}

// lcpWith returns the length of the longest common prefix of the key at
// position i, with the page's prefix, and s.
func (p *Page) lcpWith(i int, s []byte) int {
	pl := int(p.plen)
	if l := lcp(s, p.prefix()); l < pl {
		return l
	}
	rest := s[pl:]
	m := min(len(rest), p.length(i))
	if x := p.heads()[i] ^ word(rest); x != 0 {
		return pl + min(bits.LeadingZeros64(x)/8, m)
	}
	if m <= headLen {
		return pl + m
	}
	return pl + headLen + lcp(p.tail(i)[:m-headLen], rest[headLen:m])
}

// sharedBy returns how many bytes the keys from to to of p, and if extra one
// more suffix s, share and might be stored once (0 if none).
func (p *Page) sharedBy(from, to int, s []byte, extra bool) (q int) {
	m := to - from
	if m == 0 || m == 1 && !extra || MinPrefix > maxSuffix {
		return 0
	}
	q = int(p.plen) + p.length(from) // the whole first key
	if m >= 2 {
		q = int(p.plen) + p.lcpEntries(from, to-1)
	}
	if extra {
		q = min(q, p.lcpWith(from, s))
	}
	if q = max(q-PrefixSlack, 0); q < MinPrefix {
		return 0
	}
	return q
}

// lcp returns the length of the longest common prefix of a and b.
func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	for i := range n {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

// planWith returns the plan of the entries from to to of p and, if extra, of one
// more suffix s, as if they had a prefix of q bytes (setPrefix fills it in).
func (p *Page) planWith(from, to int, s []byte, extra bool, q int) plan {
	pl := plan{n: to - from, plen: q}
	tails, uni := p.shape(from, to, q)
	if extra {
		pl.n++
		tails, uni = see(tails, uni, len(s)-q)
	}
	pl.tails = tails
	if uni >= 1 && uni <= headLen {
		pl.ulen = uni
	}
	return pl
}

// planFor returns the plan of the entries from to to of p and, if extra, of one
// more suffix s: with the prefix they share if that saves MinGain bytes.
func (p *Page) planFor(from, to int, s []byte, extra bool) plan {
	pl := p.planWith(from, to, s, extra, 0)
	if pl.tails < MinGain { // a prefix only shortens tails
		return pl
	}
	if q := p.sharedBy(from, to, s, extra); q > 0 {
		if with := p.planWith(from, to, s, extra, q); need(with)+MinGain <= need(pl) {
			return with
		}
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
func newPage(c int, pl plan, pfx []byte) *Page {
	q := alloc(c)
	q.class, q.ulen, q.plen = uint8(c), uint8(pl.ulen), uint8(pl.plen)
	size := sizes[c]
	q.top = uint16(size)
	copy(q.mem()[hdr:], pfx)
	if pl.ulen != 0 {
		k := (size - dirAt(pl.plen) - 7) / 17
		for ; arraysEnd(k+1, true, pl.plen) <= size; k++ {
		}
		q.cap = uint8(k)
	} else {
		avg := 12 // an empty page: assume tails of this length
		if pl.n > 0 {
			avg = (pl.tails + pl.n - 1) / pl.n
		}
		k := pl.n + (size-need(pl))/(20+avg)
		for k = min(k, 255); arraysEnd(k, false, pl.plen)+pl.tails > size; k-- {
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
func New(c, ulen int) *Page { return newPage(c, plan{ulen: ulen}, nil) }

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

// appendFull adds a key, given with the prefix, after the last.
func (q *Page) appendFull(key []byte, v uint64) {
	rest := key[q.plen:]
	var tail []byte
	if len(rest) > headLen {
		tail = rest[headLen:]
	}
	q.appendEntry(word(rest), v, len(rest), tail)
}

// insertion is a suffix that rebuild adds in front of the entry at position pos.
type insertion struct {
	pos int
	s   []byte
	v   uint64
}

// rebuild returns a page of class c and plan pl that holds the entries from to
// to of p and, if in is not nil, its suffix, in order. They must fit. If the
// new prefix is as long as the old one the entries are copied; else every key is
// cut anew, since its head word changes.
func (p *Page) rebuild(c int, pl plan, from, to int, in *insertion) *Page {
	var buf [maxSuffix]byte
	var pfx []byte
	if pl.plen > 0 { // the keys share it: any of them tells it
		pfx = p.Key(from, &buf)[:pl.plen]
	}
	q := newPage(c, pl, pfx)
	same := pl.plen == int(p.plen)
	Counts.Rebuilds++
	if !same {
		Counts.PrefixChanges++
	}
	for i := from; i < to; i++ {
		if in != nil && in.pos == i {
			q.appendFull(in.s, in.v)
		}
		if same {
			q.appendEntry(p.heads()[i], p.vals()[i], p.length(i), p.tail(i))
		} else {
			q.appendFull(p.Key(i, &buf), p.vals()[i])
		}
	}
	if in != nil && in.pos == to {
		q.appendFull(in.s, in.v)
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
	i, found := p.find(s)
	if found {
		p.vals()[i] = v
		return p, Updated, nil
	}
	rest, rel := p.strip(s)
	tail := max(len(rest)-headLen, 0)
	fits := rel == 0 && (p.ulen == 0 || len(rest) == int(p.ulen))
	if fits && int(p.count) < int(p.cap) && (p.ulen != 0 || tail <= p.heapFree()) {
		p.insertAt(i, word(rest), v, rest)
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
	i, found := p.find(s)
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
			Counts.Shrinks++
			return p.rebuild(int(p.class)-1, pl, 0, n-1, nil), true
		}
	}
	return p, true
}

// cost is what the entry at position i takes (not counting the prefix).
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
	with := func(q int) plan {
		pl := plan{n: int(a.count + b.count), plen: q}
		ta, ua := a.shape(0, int(a.count), q)
		tb, ub := b.shape(0, int(b.count), q)
		pl.tails = ta + tb
		if ua == ub && ua >= 1 && ua <= headLen {
			pl.ulen = ua
		}
		return pl
	}
	pl := with(0)
	if arraysEnd(pl.n, true, 0) > MergeFill*sizes[MaxClass]/100 { // too much even for the leanest page: no prefix helps
		return nil
	}
	var fb, lb [maxSuffix]byte
	first, last := a.Key(0, &fb), b.Key(int(b.count)-1, &lb)
	if q := lcp(first, last) - PrefixSlack; pl.tails >= MinGain && q >= max(MinPrefix, 1) {
		if cand := with(q); need(cand)+MinGain <= need(pl) {
			pl = cand
		}
	}
	c := fitClass(pl, 0, MergeFill)
	if c < 0 {
		return nil
	}
	m := newPage(c, pl, first[:pl.plen])
	Counts.Merges++
	var buf [maxSuffix]byte
	for _, src := range []*Page{a, b} {
		for i := range int(src.count) {
			m.appendFull(src.Key(i, &buf), src.vals()[i])
		}
	}
	return m
}

// Each calls fn for the suffixes from from on, in order, with their values,
// until fn returns false. A nil from starts at the first.
func (p *Page) Each(from []byte, fn func(s []byte, v uint64) bool) {
	i := 0
	if from != nil {
		i, _ = p.find(from)
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
	Prefix  int  // bytes of the prefix the keys share, stored once
	Uniform bool // the uniform flavor
}

// Stats returns the statistics of the page.
func (p *Page) Stats() Stats {
	pl := p.planFor(0, int(p.count), nil, false)
	return Stats{Keys: pl.n, Size: p.Size(), Used: need(pl), Tails: pl.tails, Prefix: pl.plen, Uniform: p.ulen != 0}
}
