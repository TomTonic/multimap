package art

import "unsafe"

// U8-n pages hold keys like U8-1 pages (one length of at most 8 bytes, full
// keys as big-endian words at the same place), but each key with one or more
// values. A key's values lie inline, packed after those of the keys before
// it, as long as they are at most inlineMax; beyond that they move to a value
// set of their own that the page points to. The page's only pointers are
// those few external sets.
//
// Per key, one count byte says how many values it holds inline, or, with
// extBit set, which external slot holds its set.

// The U8-n classes: 128, 256 and 512 bytes, for 4, 8 and 16 keys with 9, 20
// and 41 inline values and 1, 2 and 4 external sets.
type (
	pageN4 struct {
		pageHead
		heads [4]uint64
		cnt   [8]uint8
		ext   [1]unsafe.Pointer
		vals  [9]uint64
	}
	pageN8 struct {
		pageHead
		heads [8]uint64
		cnt   [8]uint8
		ext   [2]unsafe.Pointer
		vals  [20]uint64
	}
	pageN16 struct {
		pageHead
		heads [16]uint64
		cnt   [16]uint8
		ext   [4]unsafe.Pointer
		vals  [41]uint64
	}
)

// nLayout describes where a page with values held like this keeps its keys,
// counts and values: a U8-n class, or an S page (see sLayout).
type nLayout struct {
	keys, vals, ext                   int
	headsOff, cntOff, extOff, valsOff uintptr
}

var nLayouts = [3]nLayout{
	{4, 9, 1, headsOff, unsafe.Offsetof(pageN4{}.cnt), unsafe.Offsetof(pageN4{}.ext), unsafe.Offsetof(pageN4{}.vals)},
	{8, 20, 2, headsOff, unsafe.Offsetof(pageN8{}.cnt), unsafe.Offsetof(pageN8{}.ext), unsafe.Offsetof(pageN8{}.vals)},
	{16, 41, 4, headsOff, unsafe.Offsetof(pageN16{}.cnt), unsafe.Offsetof(pageN16{}.ext), unsafe.Offsetof(pageN16{}.vals)},
}

const (
	// inlineMax is the most values a key keeps inline; with more, they move
	// to an external set, where adding and finding a value stays cheap.
	inlineMax = 8
	extBit    = 0x80
)

func newPageN(class int) *pageHead {
	var p *pageHead
	switch class {
	case 0:
		p = &(&pageN4{}).pageHead
	case 1:
		p = &(&pageN8{}).pageHead
	default:
		p = &(&pageN16{}).pageHead
	}
	p.kind, p.class = kPageN, uint8(class)
	return p
}

// nClassFor returns the smallest U8-n class for k keys, v inline values and
// e external sets, or -1 if none holds them.
func nClassFor(k, v, e int) int {
	for c, l := range nLayouts {
		if k <= l.keys && v <= l.vals && e <= l.ext {
			return c
		}
	}
	return -1
}

// layout returns where a U8-n or S page keeps its arrays.
func (p *pageHead) layout() nLayout {
	if p.kind == kPageS {
		return p.sLayout().nLayout
	}
	return nLayouts[p.class]
}

// nHeads, cnts, exts and nvals return the arrays of a U8-n or S page, over
// their full capacity.
func (p *pageHead) nHeads() []uint64 {
	l := p.layout()
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), l.headsOff)), l.keys)
}

func (p *pageHead) cnts() []uint8 {
	l := p.layout()
	return unsafe.Slice((*uint8)(unsafe.Add(unsafe.Pointer(p), l.cntOff)), l.keys)
}

func (p *pageHead) exts() []unsafe.Pointer {
	l := p.layout()
	return unsafe.Slice((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(p), l.extOff)), l.ext)
}

func (p *pageHead) nvals() []uint64 {
	l := p.layout()
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), l.valsOff)), l.vals)
}

// run returns where the values of key i lie: at nvals()[off:off+n], or, with
// e >= 0, in the external set exts()[e].
func (p *pageHead) run(i int) (off, n, e int) {
	c := p.cnts()
	for _, x := range c[:i] {
		if x < extBit {
			off += int(x)
		}
	}
	if c[i] >= extBit {
		return off, 0, int(c[i] &^ extBit)
	}
	return off, int(c[i]), -1
}

// extUsed returns the number of external sets in use.
func (p *pageHead) extUsed() int {
	n := 0
	for _, x := range p.exts() {
		if x != nil {
			n++
		}
	}
	return n
}

// nResize copies a U8-n page into a new page of the given class, packing
// the external sets into the first slots.
func (p *pageHead) nResize(class int) *pageHead {
	q := newPageN(class)
	q.count, q.klen, q.nv = p.count, p.klen, p.nv
	n := int(p.count)
	copy(q.nHeads(), p.nHeads()[:n])
	copy(q.nvals(), p.nvals()[:p.nv])
	pc, qc, pe, qe := p.cnts(), q.cnts(), p.exts(), q.exts()
	slot := 0
	for i, x := range pc[:n] {
		qc[i] = x
		if x >= extBit {
			qe[slot] = pe[x&^extBit]
			qc[i] = extBit | uint8(slot)
			slot++
		}
	}
	return q
}

// toN turns a U8-1 page of at most 16 keys into a U8-n page with room for
// extra more inline values.
func (p *pageHead) toN(extra int) *pageHead {
	n := int(p.count)
	q := newPageN(nClassFor(n, n+extra, 0))
	q.count, q.klen, q.nv = p.count, p.klen, p.count
	copy(q.nHeads(), p.heads()[:n])
	copy(q.nvals(), p.vals()[:n])
	c := q.cnts()
	for i := range n {
		c[i] = 1
	}
	return q
}

// nInsertKey inserts key w with the one value v at key position i. The page
// must have room for one more key and value.
func (p *pageHead) nInsertKey(i int, w, v uint64) {
	n := int(p.count)
	h, c := p.nHeads(), p.cnts()
	copy(h[i+1:n+1], h[i:n])
	copy(c[i+1:n+1], c[i:n])
	h[i], c[i] = w, 0
	off, _, _ := p.run(i)
	vs := p.nvals()
	copy(vs[off+1:int(p.nv)+1], vs[off:p.nv])
	vs[off], c[i] = v, 1
	p.count++
	p.nv++
}

// nAddVal appends v to the inline values of key i. The page must have room
// for one more value, and the key must hold fewer than inlineMax.
func (p *pageHead) nAddVal(i int, v uint64) {
	off, n, _ := p.run(i)
	vs := p.nvals()
	copy(vs[off+n+1:int(p.nv)+1], vs[off+n:p.nv])
	vs[off+n] = v
	p.cnts()[i]++
	p.nv++
}

// nRemoveVal removes the k-th inline value of key i, which holds at least two.
func (p *pageHead) nRemoveVal(i, k int) {
	off, _, _ := p.run(i)
	vs := p.nvals()
	copy(vs[off+k:p.nv-1], vs[off+k+1:p.nv])
	p.cnts()[i]--
	p.nv--
}

// nExternalize moves the inline values of key i out: the external set s,
// which the caller filled with them, takes their place. The page must have a
// free external slot.
func (p *pageHead) nExternalize(i int, s unsafe.Pointer) {
	e := 0
	for p.exts()[e] != nil {
		e++
	}
	off, n, _ := p.run(i)
	vs := p.nvals()
	copy(vs[off:int(p.nv)-n], vs[off+n:p.nv])
	p.nv -= uint8(n)
	p.exts()[e] = s
	p.cnts()[i] = extBit | uint8(e)
}

// nRemoveKey removes key i with all its values and returns the page, its
// smaller replacement once it holds few enough keys and values, or nil once
// it is empty.
func (p *pageHead) nRemoveKey(i int) *pageHead {
	if p.count == 1 {
		return nil
	}
	p.nDropKey(i)
	if p.class > 0 {
		l := nLayouts[p.class-1]
		if int(p.count) <= l.keys/2 && int(p.nv) <= l.vals/2 && p.extUsed() <= l.ext {
			return p.nResize(int(p.class) - 1)
		}
	}
	return p
}

// nDropKey removes key i with its values, inline or external, from its
// head, count and value arrays.
func (p *pageHead) nDropKey(i int) {
	cnt := int(p.count)
	off, n, e := p.run(i)
	if e >= 0 {
		p.exts()[e] = nil
	} else {
		vs := p.nvals()
		copy(vs[off:int(p.nv)-n], vs[off+n:p.nv])
		p.nv -= uint8(n)
	}
	h, c := p.nHeads(), p.cnts()
	copy(h[i:cnt-1], h[i+1:cnt])
	copy(c[i:cnt-1], c[i+1:cnt])
	p.count--
}

// fillVals stores the values of items, inline or as external sets, in the new
// U8-n or S page p, key by key.
func (p *pageHead) fillVals(items []item) {
	cs, ex, vs := p.cnts(), p.exts(), p.nvals()
	off, slot := 0, 0
	for i, it := range items {
		if it.set != nil {
			ex[slot], cs[i] = it.set, extBit|uint8(slot)
			slot++
			continue
		}
		off += copy(vs[off:], it.vals)
		cs[i] = uint8(len(it.vals))
	}
}
