package lpage

import "bytes"

// An ent is an entry of a page under construction: a key in up to three
// pieces, laid end to end, and a value. The pieces and the value alias the
// page the entry came from or the caller's bytes, so that a page is rebuilt
// from another without a copy of its entries in between: a rebuild that
// inserts, deletes, splits, merges or moves a page one level lays the entries
// out anew, and has them all on the stack.
type ent struct {
	k0, k1, k2, v []byte
	// cont marks an entry of the same key as the entry before it: the page
	// stores the key once, for the first entry of the run (a multi-value entry).
	cont bool
}

// maxEnts is the most entries any page holds, two pages of them for a merge, and
// one more for an insertion.
const maxEnts = 32

func (e *ent) klen() int { return len(e.k0) + len(e.k1) + len(e.k2) }

// copyKey appends the bytes from to to of the key of e to dst.
func (e *ent) copyKey(dst []byte, from, to int) []byte {
	if to <= from {
		return dst
	}
	l0, l1 := len(e.k0), len(e.k1)
	if from < l0 {
		dst = append(dst, e.k0[from:min(to, l0)]...)
	}
	if to > l0 && from < l0+l1 {
		dst = append(dst, e.k1[max(from-l0, 0):min(to-l0, l1)]...)
	}
	if to > l0+l1 {
		dst = append(dst, e.k2[max(from-l0-l1, 0):to-l0-l1]...)
	}
	return dst
}

// load fills es with the entries of p, which alias p, and returns their number.
func (p *Page) load(es []ent) int {
	m := p.mem()
	r, v := p.lens()
	n := p.Len()
	h, cp := p.hdr(), int(p.cp)
	pre := m[h : h+cp]
	koff := h + cp
	voff := koff + remSum(r[:n])
	var key []byte
	for i := range n {
		x, vl := int(r[i]), p.vlen(v, i)
		if x != cont {
			key = m[koff : koff+x]
			koff += x
		}
		es[i] = ent{k0: pre, k1: key, v: m[voff : voff+vl], cont: x == cont}
		voff += vl
	}
	return n
}

// planEnts chooses the prefix and the sizes for the sorted entries es: the
// code byte of the page, the prefix length and the bytes needed. ok is false if
// there are too many entries, a remainder or a value is too long, or the
// largest class is too small. Values of one length that is a width of a page (1,
// 2, 4, 8, 16, 32 or 64) go into a page of that width.
func planEnts(es []ent) (code uint8, cp, need int, ok bool) {
	n := len(es)
	wcode := 0
	for i := 1; i < len(widths); i++ {
		if widths[i] == len(es[0].v) {
			wcode = i
		}
	}
	for i := range es {
		if len(es[i].v) != len(es[0].v) {
			wcode = 0
		}
	}
	h := headerFor(n, widths[wcode])
	if h == 0 {
		return 0, 0, 0, false
	}
	shortest := es[0].klen()
	for i := range es {
		shortest = min(shortest, es[i].klen())
	}
	if n > 1 {
		var fb, lb [2*maxField + 2]byte
		a := es[0].copyKey(fb[:0], 0, es[0].klen())
		b := es[n-1].copyKey(lb[:0], 0, es[n-1].klen())
		cp = min(lcp(a, b), shortest-1, maxField)
	}
	need = h + cp
	for i := range es {
		if r := es[i].klen() - cp; len(es[i].v) > maxField || !es[i].cont && (r < 1 || r > maxRem) {
			return 0, 0, 0, false
		}
		if !es[i].cont {
			need += es[i].klen() - cp
		}
		need += len(es[i].v)
	}
	for c, s := range sizes {
		if need <= s {
			return uint8(c) | uint8(h/8-1)<<2 | uint8(wcode)<<5, cp, need, true
		}
	}
	return 0, 0, 0, false
}

// writeEnts lays out the page chosen by planEnts.
func writeEnts(code uint8, cp int, es []ent) *Page {
	p := alloc(int(code & 3))
	p.code, p.cp = code+KindBase, uint8(cp)
	m := p.mem()
	r, v := p.lens()
	off := p.hdr()
	off += len(es[0].copyKey(m[off:off], 0, cp))
	for i := range es {
		if es[i].cont {
			r[i] = cont
			continue
		}
		k := es[i].klen()
		r[i] = byte(k - cp)
		off += len(es[i].copyKey(m[off:off], cp, k))
	}
	for i := range es {
		if v != nil {
			v[i] = byte(len(es[i].v))
		}
		off += copy(m[off:], es[i].v)
	}
	return p
}

// buildEnts returns the page of es, or nil if they do not fit one.
func buildEnts(es []ent) *Page {
	code, cp, _, ok := planEnts(es)
	if !ok {
		return nil
	}
	return writeEnts(code, cp, es)
}

// fitEnts returns the page of es if they fit one in at most maxPct percent of
// the class below class c, else nil.
func fitEnts(es []ent, c, maxPct int) *Page {
	if code, cp, need, ok := planEnts(es); ok && need*100 <= maxPct*sizes[c-1] {
		return writeEnts(code, cp, es)
	}
	return nil
}

// position returns where s goes among es (sorted, whole keys; the first entry of
// a run of one key counts), and whether it is there.
func position(es []ent, s []byte) (int, bool) {
	var buf [2*maxField + 2]byte
	for i := range es {
		if es[i].cont {
			continue
		}
		switch c := bytes.Compare(es[i].copyKey(buf[:0], 0, es[i].klen()), s); {
		case c == 0:
			return i, true
		case c > 0:
			return i, false
		}
	}
	return len(es), false
}
