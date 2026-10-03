package lpage

import "bytes"

// This file holds what a tree needs from a page beyond Run's Insert, Get and
// Delete: positions instead of keys, so that a tree can find a key once and
// then read, delete or split at the position; the pieces of an entry by
// position; and the operations that move a page to another depth of the tree.
// The keys are the remainders the tree hands in, the key from the page's depth
// on; the page does not know its depth.

// MergeFill is how full, in percent of the largest class, the page of a Merge
// may be.
var MergeFill = 75

// Find returns the position of remainder s and whether it is there. Unlike
// Get it tells where the entry is.
func (p *Page) Find(s []byte) (int, bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	if len(s) <= cp || !bytes.Equal(m[h:h+cp], s[:cp]) {
		return 0, false
	}
	rem := s[cp:]
	first, off := rem[0], h+cp
	for i, x := range m[2 : 2+p.capacity()] {
		if x == 0 {
			break
		}
		if int(x) == len(rem) && m[off] == first && bytes.Equal(m[off:off+int(x)], rem) {
			return i, true
		}
		off += int(x)
	}
	return 0, false
}

// Seek returns how many entries of the page are below remainder s, which need
// not start with the common prefix, and whether the entry at that position
// equals s.
func (p *Page) Seek(s []byte) (int, bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	k := min(len(s), cp)
	if c := bytes.Compare(s[:k], m[h:h+k]); c != 0 {
		if c < 0 {
			return 0, false
		}
		return p.Len(), false
	}
	if len(s) <= cp {
		return 0, false // a prefix of the common prefix sorts below every entry
	}
	i, _, found := p.locate(s)
	return i, found
}

// offsets returns where the remainder and the value of entry i start.
func (p *Page) offsets(i int) (koff, voff int) {
	r, v := p.lens()
	koff = p.hdr() + int(p.cp)
	kend := koff + sumBytes(r)
	return koff + sumBytes(r[:i]), p.voffset(v, kend, i)
}

// ValueAt returns the value of entry i. The slice aliases the page and is valid
// until the page changes.
func (p *Page) ValueAt(i int) []byte {
	_, v := p.lens()
	_, voff := p.offsets(i)
	return p.mem()[voff : voff+p.vlen(v, i)]
}

// ValueIs reports whether the value of entry i is val.
func (p *Page) ValueIs(i int, val string) bool {
	return string(p.ValueAt(i)) == val
}

// EachValue calls fn with the value of every entry from i up to j, in order,
// until fn returns false, and reports whether it ran to completion. The slices
// alias the page.
func (p *Page) EachValue(i, j int, fn func(val []byte) bool) bool {
	m := p.mem()
	_, v := p.lens()
	_, off := p.offsets(i)
	for k := i; k < j; k++ {
		x := p.vlen(v, k)
		if !fn(m[off : off+x]) {
			return false
		}
		off += x
	}
	return true
}

// EachString calls fn with the value of every entry from i up to j, in order,
// until fn returns false, and reports whether it ran to completion. The strings
// are copies of the page's bytes, made together: one allocation for the run
// instead of one for each value, so a string that a caller keeps holds the
// bytes of its neighbours alive until it is dropped.
func (p *Page) EachString(i, j int, fn func(val string) bool) bool {
	if i >= j {
		return true
	}
	_, v := p.lens()
	_, off := p.offsets(i)
	end := off
	for k := i; k < j; k++ {
		end += p.vlen(v, k)
	}
	run := string(p.mem()[off:end])
	off = 0
	for k := i; k < j; k++ {
		x := p.vlen(v, k)
		if !fn(run[off : off+x]) {
			return false
		}
		off += x
	}
	return true
}

// AppendKey appends the remainder of entry i, with the common prefix, to dst.
func (p *Page) AppendKey(dst []byte, i int) []byte {
	m := p.mem()
	r, _ := p.lens()
	koff, _ := p.offsets(i)
	h, cp := p.hdr(), int(p.cp)
	return append(append(dst, m[h:h+cp]...), m[koff:koff+int(r[i])]...)
}

// ByteAt returns byte off of the remainder of entry i, with the common prefix.
// off must be within the remainder.
func (p *Page) ByteAt(i, off int) byte {
	m := p.mem()
	cp := int(p.cp)
	if off < cp {
		return m[p.hdr()+off]
	}
	koff, _ := p.offsets(i)
	return m[koff+off-cp]
}

// Shared returns how many bytes all remainders of the page start with: for a
// page of one entry, its whole remainder.
func (p *Page) Shared() int {
	n := p.Len()
	if n == 1 {
		r, _ := p.lens()
		return int(p.cp) + int(r[0])
	}
	var a, b [maxField]byte
	return lcp(p.AppendKey(a[:0], 0), p.AppendKey(b[:0], n-1))
}

// Thin reports whether the page holds so little that it may fit a page with a
// neighbour (see Merge): at most half of what a merge allows.
func (p *Page) Thin() bool {
	return 2*p.Used()*100 <= MergeFill*sizes[len(sizes)-1]
}

// DeleteAt removes entry i and returns the page that holds the rest, nil if it
// was the only entry.
func (p *Page) DeleteAt(i int) *Page {
	q, _ := p.deleteAt(i)
	return q
}

// SplitAt divides a page into the page of its first m entries and the page of
// the others (0 < m < Len), each in the smallest class that holds it.
func (p *Page) SplitAt(m int) (left, right *Page) {
	var es [maxEnts]ent
	n := p.load(es[:])
	return buildEnts(es[:m]), buildEnts(es[m:n])
}

// Merge returns the page that holds the entries of a and then those of b, which
// must all be above a's, if it holds them with at most MergeFill percent of the
// largest class, else nil.
func Merge(a, b *Page) *Page {
	// The merged page needs at least the bytes of both without their headers and
	// prefixes, and the smallest header: a test that most merges fail at.
	room := capacityOf(MaxHeader, 0)
	if w := a.width(); w != 0 && w == b.width() {
		room = capacityOf(MaxHeader, w)
	}
	if a.Len()+b.Len() > room || 100*(a.Used()-a.hdr()-int(a.cp)+b.Used()-b.hdr()-int(b.cp)+8) > MergeFill*sizes[len(sizes)-1] {
		return nil
	}
	var es [maxEnts]ent
	if a.Len()+b.Len() > len(es) {
		return nil
	}
	n := a.load(es[:])
	n += b.load(es[n:])
	code, cp, need, ok := planEnts(es[:n])
	if !ok || need*100 > MergeFill*sizes[len(sizes)-1] {
		return nil
	}
	return writeEnts(code, cp, es[:n])
}

// Skip returns the page without the first k bytes of every remainder, which
// all remainders must share and which must leave each at least one byte: the
// page of the same entries one level deeper in the tree.
func (p *Page) Skip(k int) *Page {
	if k == 0 {
		return p
	}
	var es [maxEnts]ent
	n := p.load(es[:])
	for i := range es[:n] {
		e := &es[i]
		d := k
		for _, piece := range []*[]byte{&e.k0, &e.k1, &e.k2} {
			c := min(d, len(*piece))
			*piece = (*piece)[c:]
			d -= c
		}
	}
	return buildEnts(es[:n])
}

// Prepend returns the page with pre in front of every remainder, the page of the
// same entries one level higher in the tree, or nil if they do not fit.
func (p *Page) Prepend(pre []byte) *Page {
	var es [maxEnts]ent
	n := p.load(es[:])
	for i := range es[:n] {
		// the page's own prefix and remainder follow pre
		es[i] = ent{k0: pre, k1: es[i].k0, k2: es[i].k1, v: es[i].v}
	}
	return buildEnts(es[:n])
}

// Fits reports whether an entry with remainder s and value val can go into a
// page at all.
func Fits(s, val []byte) bool { return len(s) > 0 && !tooBig(s, val) }

// FitsLen is Fits for lengths.
func FitsLen(slen, vlen int) bool {
	return slen > 0 && slen <= maxField && vlen <= maxField && slen+vlen <= sizes[len(sizes)-1]-8
}

// MaxEntries returns the most entries any page can hold: the capacity of the
// largest header for values of one width.
func MaxEntries() int { return capacityOf(MaxHeader, 1) }
