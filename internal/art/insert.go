package art

import (
	"bytes"
	"slices"

	"github.com/TomTonic/multimap/internal/swar"
	"github.com/TomTonic/multimap/internal/vpage"
)

// spot is where upsert left a key that already was in a page: at position i of
// the page that the slot at loc points to, whose path starts at pathLen. A page
// below a range node is child pi of the node at *par.
type spot struct {
	loc     **header
	i       int
	pathLen int
	par     **header
	pi      int
}

// upsert returns the slot that holds the key's leaf or page, creating the key
// when it is missing: in a page with the raw value v when the key fits one (see
// pageable) and does not go below a byte node (see rebuild.go), else as a leaf
// made by nl, which the caller fills. The caller may replace a leaf in its
// slot, as a flat leaf does when it grows. It tells a key it has created by
// the tree's size, and, for a key that was in a page, where in t.at.
//
// It descends like find, checking common prefixes and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, v uint64, nl newLeafFunc) **header {
	loc, pathLen := &t.root, 0
	var par **header // the range node *loc is a child of, or nil
	pi := 0
	belowByteNode := false // *loc is below a byte node, where no pages go
	for {
		n := *loc
		if n == nil {
			*loc = t.newChild(key, v, nl, pathLen)
			t.size++
			return loc
		}
		if n.objType <= maxMultiKeyByte {
			if n.objType > maxSingleKeyByte {
				return t.upsertPage(loc, par, pi, key, pathLen, v, nl)
			}
			return t.splitLeaf(loc, asSingleKey(n), key, pathLen, v, nl, belowByteNode)
		}
		if n.plen > 0 {
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, pathLen) {
				pl = n.prefixLen()
				if !prefixMatches(n, pl, key, pathLen) {
					mis, _ := prefixLcp(n, pl, key[pathLen:])
					return t.splitPrefix(loc, n, mis, key, pathLen, v, nl)
				}
			}
			pathLen += pl
		}
		if pathLen == len(key) {
			if endPageOf(n) == nil {
				*loc = setEndPage(n, nl(key, pathLen))
				t.size++
			}
			return endPageSlot(*loc)
		}
		b := key[pathLen]
		if isRange(n.objType) {
			x := asR(n)
			i := x.index(b)
			c := &x.children()[i]
			if cb, ok := firstByte(*c, pathLen); ok && cb != b {
				// The child's keys all have byte cb: the key gets a range of
				// its own, next to it.
				nc := t.newChild(key, v, nl, pathLen)
				rs := []rng{{0, *c}, {b, nc}}
				if b < cb {
					rs = []rng{{0, nc}, {cb, *c}}
				}
				*loc = rSplice(n, i, rs)
				y := asR(*loc)
				t.size++
				return &y.children()[y.index(b)]
			}
			par, pi, loc = loc, i, c
			continue
		}
		par, belowByteNode = nil, true
		c := findLoc(n, b)
		if c == nil {
			var slot **header
			*loc, slot = addChild(n, b, singleKeyHdr(nl(key, pathLen+1)))
			t.size++
			return slot
		}
		loc, pathLen = c, pathLen+1
	}
}

// hasPages reports whether the tree may hold pages: its root is a page or a
// range node. Below a root that is a byte node there are none, and there
// never are (see rebuild.go).
func (t *Tree) hasPages() bool {
	return t.root != nil && (isMultiKey(t.root.objType) || isRange(t.root.objType))
}

// firstByte returns byte pathLen of every key below c, a child of a range
// node, if they all share it: the byte of a leaf, or the first byte of the common prefix of a
// node. A page's keys may differ there.
func firstByte(c *header, pathLen int) (byte, bool) {
	switch {
	case isSingleKey(c.objType):
		return asSingleKey(c).from(pathLen)[0], true
	case isMultiKey(c.objType) || c.plen == 0:
		return 0, false
	}
	return c.prefix[0], true
}

// newChild returns a new page holding key with the raw value v when the key
// fits a page, else a new leaf made by nl, for a child at pathLen.
func (t *Tree) newChild(key []byte, v uint64, nl newLeafFunc, pathLen int) *header {
	if !t.pageable(key, pathLen) {
		return singleKeyHdr(nl(key, pathLen))
	}
	return multiKeyHdr(newPageFor(key, pathLen, v))
}

// upsertPage handles a key whose descent reaches the page at *loc, whose keys
// are the ones below pathLen: the key is there, or goes in, or the subtree is
// rebuilt with it (see build) because the page is full or the key does not fit
// it.
func (t *Tree) upsertPage(loc, par **header, pi int, key []byte, pathLen int, v uint64, nl newLeafFunc) **header {
	p := asMultiKey(*loc)
	base := p.Base()
	full := false // the key fits the page but for its room
	if len(key)-base <= maxPageRemainder {
		i, ok := p.LocateIn(key)
		if ok {
			t.at = spot{loc: loc, i: i, pathLen: pathLen, par: par, pi: pi}
			return loc
		}
		q, res := p.InsertIn(i, key, v) // the remainder is not too long: checked above
		if res != vpage.Full {
			*loc = multiKeyHdr(q)
			t.size++
			return loc
		}
		full = true
	}
	if full {
		if par != nil && splitFull(loc, par, pi, pathLen) {
			r := asR(*par)
			i := r.index(key[pathLen])
			return t.upsertPage(&r.children()[i], par, i, key, pathLen, v, nl)
		}
		return t.burst(loc, p, key, pathLen, v, nl)
	}
	it := item{key: key, val: v}
	if !t.pageable(key, pathLen) {
		it = item{key: key, leaf: nl(key, pathLen)}
	}
	items := pageItems(p, key)
	i, _ := slices.BinarySearchFunc(items, key, func(x item, k []byte) int { return bytes.Compare(x.key, k) })
	t.replace(loc, par, pi, slices.Insert(items, i, it), pathLen)
	t.size++
	return t.upsert(key, v, nl) // finds the key in the rebuilt subtree
}

// burst makes room for key where the full page p at *loc, whose keys start at
// pathLen and share their byte there (or the page is the root), cannot be split
// by a byte at that pathLen. The keys of p all start with the bytes up to some
// pathLen d. If key starts with them too, the page gets a range node of its own
// with those bytes as its common prefix, in which splitFull can cut it where its keys
// differ; else key leaves them where it differs, and a range node there holds p
// and a new page or leaf for key (see fork). Either way no key of the page is
// copied: it replaces a rebuild of the subtree from its keys.
func (t *Tree) burst(loc **header, p *vpage.Page, key []byte, pathLen int, v uint64, nl newLeafFunc) **header {
	var buf [maxPageRemainder]byte
	base, d := p.Base(), p.Base()+p.Shared()
	first := p.Key(0, &buf)[pathLen-base:] // the first key, from pathLen on: it holds the shared bytes
	shared := first[:d-pathLen]
	if m := swar.Lcp(shared, key[pathLen:]); pathLen+m < d {
		return t.fork(loc, multiKeyHdr(p), false, key, pathLen, pathLen+m, v, nl)
	}
	var endPage *singleKeyHead
	if len(first) == d-pathLen {
		// The first key is the shared bytes: it ends at the range node, so it is
		// the node's end page, and a leaf. The others are longer.
		whole := append(key[:pathLen:pathLen], first...)
		endPage = t.mk(whole, d, p.Val(0))
		p = p.DeleteAt(0) // the page was full: it has more keys
	}
	*loc = makeR(key[pathLen:d], endPage, []rng{{0, multiKeyHdr(p)}})
	return t.upsert(key, v, nl)
}

// splitFull splits the full page at *loc, child pi of the range node at
// *par, whose keys start at pathLen, at the byte boundary nearest its middle
// (see ranges): the page keeps the keys before it and a new page takes the
// others, with a range of its own. Unlike a rebuild, it copies half a page
// and allocates one. It reports false and changes nothing when all keys share
// their byte at pathLen.
func splitFull(loc, par **header, pi, pathLen int) bool {
	p := asMultiKey(*loc)
	n := p.Len()
	at := func(i int) byte { return p.ByteAt(i, pathLen-p.Base()) }
	if at(0) == at(n-1) {
		return false // all keys share byte pathLen
	}
	m := n / 2
	b, lo, hi := at(m), m, m
	for lo > 0 && at(lo-1) == b {
		lo--
	}
	for hi < n && at(hi) == b {
		hi++
	}
	s := lo
	if lo == 0 || (hi < n && hi-m < m-lo) {
		s = hi
	}
	rb := at(s)
	left, right := p.SplitOff(s)
	*par = rSplice(*par, pi, []rng{{0, multiKeyHdr(left)}, {rb, multiKeyHdr(right)}})
	return true
}

// replace puts items, the keys of the page at *at, in the page's place: as a
// subtree (see build), or, for a page below a range node, as ranges that
// take over the page's range (see ranges).
func (t *Tree) replace(at, par **header, pi int, items []item, pathLen int) {
	if par == nil {
		*at = t.build(items, pathLen)
		return
	}
	*par = rSplice(*par, pi, t.ranges(items, pathLen, nil))
}

// promote gives the key at sp, which has one value in its page, the leaf
// l, which holds its key from sp.pathLen on and its values, and lets the subtree
// around it fall back to byte nodes if keys with several values crowd it
// (see fallBack).
func (t *Tree) promote(sp spot, l *singleKeyHead, key []byte) {
	t.place(sp, l, key)
	t.fallBack(key)
}

// place puts the leaf l of key, the key at sp, in the key's place. A page
// cannot hold the leaf, so it gets a range of its own: the page is cut around
// the key when the key is alone with its byte below a range node, else the
// page's keys are rebuilt with the leaf in the key's place (see ranges), which
// gives it a range or a subtree of its own.
func (t *Tree) place(sp spot, l *singleKeyHead, key []byte) {
	p, i := asMultiKey(*sp.loc), sp.i
	n := p.Len()
	if n == 1 {
		*sp.loc = singleKeyHdr(l) // the page held only this key
		return
	}
	// Below a range node, a key alone with its byte gets a range of its own
	// by cutting the page around it.
	at := func(i int) byte { return p.ByteAt(i, sp.pathLen-p.Base()) }
	if sp.par != nil && (i == 0 || at(i-1) != at(i)) && (i == n-1 || at(i+1) != at(i)) {
		b := at(i)
		var after byte
		if i < n-1 {
			after = at(i + 1)
		}
		var rs []rng
		if i > 0 {
			var left *vpage.Page
			left, p = p.SplitAt(i)
			rs = append(rs, rng{0, multiKeyHdr(left)})
		}
		rs = append(rs, rng{b, singleKeyHdr(l)})
		if i < n-1 {
			_, right := p.SplitAt(1)
			rs = append(rs, rng{after, multiKeyHdr(right)})
		}
		*sp.par = rSplice(*sp.par, sp.pi, rs)
		return
	}
	items := pageItems(p, key)
	items[sp.i] = item{key: key, leaf: l}
	t.replace(sp.loc, sp.par, sp.pi, items, sp.pathLen)
}

// splitLeaf handles an insert that reaches leaf l at pathLen: either it is the
// key's leaf, or both keys go below a new node holding their common prefix, a
// range node if the new key goes into a page (unless l is below a byte node),
// else a byte node. l keeps its base and moves below the new node.
func (t *Tree) splitLeaf(loc **header, l *singleKeyHead, key []byte, pathLen int, v uint64, nl newLeafFunc, belowByteNode bool) **header {
	ls, rest := l.from(pathLen), key[pathLen:]
	if len(key) == l.keyLen() && bytes.Equal(ls, rest) {
		return loc
	}
	p := swar.Lcp(ls, rest)
	if d := pathLen + p; !belowByteNode && t.pageable(key, d) {
		return t.fork(loc, singleKeyHdr(l), d == l.keyLen(), key, pathLen, d, v, nl)
	}
	nn := newNode(kN5, p)
	storePrefix(nn, rest[:p])
	h, _ := attachAt(nn, ls[p:], l)
	h, slot := attachAt(h, rest[p:], nl(key, pathLen+p+min(1, len(rest)-p)))
	*loc = h
	t.size++
	return slot
}

// splitPrefix handles an insert whose key leaves n's common prefix after
// mis bytes: a new node takes the common part, with n and the new key below,
// a range node if n is one, else a byte node.
func (t *Tree) splitPrefix(loc **header, n *header, mis int, key []byte, pathLen int, v uint64, nl newLeafFunc) **header {
	var buf [prefixBuf]byte
	pk := appendPrefix(buf[:0], n) // a copy: n's common prefix changes below
	if isRange(n.objType) {
		// Below a range node, n keeps the byte it branches on in its common prefix.
		return t.fork(loc, withPrefix(n, pk[mis:]), false, key, pathLen, pathLen+mis, v, nl)
	}
	nn := newNode(kN5, mis)
	storePrefix(nn, pk[:mis])
	h, _ := addChild(nn, pk[mis], withPrefix(n, pk[mis+1:]))
	rest := key[pathLen+mis:]
	h, slot := attachAt(h, rest, nl(key, pathLen+mis+min(1, len(rest))))
	*loc = h
	t.size++
	return slot
}

// fork puts a range node with the common prefix key[pathLen:d] at *loc, in a tree with
// pages, that holds old and the new key: old is a leaf whose key ends at d
// (oend), or its keys continue with one byte at d, as the new key does unless
// it ends there.
func (t *Tree) fork(loc **header, old *header, oend bool, key []byte, pathLen, d int, v uint64, nl newLeafFunc) **header {
	prefix := key[pathLen:d]
	if d == len(key) {
		*loc = makeR(prefix, nl(key, d), []rng{{0, old}})
		t.size++
		return endPageSlot(*loc)
	}
	nc := t.newChild(key, v, nl, d)
	var endPage *singleKeyHead
	rs := []rng{{0, nc}}
	if oend {
		endPage = asSingleKey(old)
	} else if ob := forkByte(old, d); key[d] < ob {
		rs = append(rs, rng{ob, old})
	} else {
		rs = []rng{{0, old}, {key[d], nc}}
	}
	*loc = makeR(prefix, endPage, rs)
	r := asR(*loc)
	t.size++
	return &r.children()[r.index(key[d])]
}

// forkByte returns byte d of every key below old, a leaf, a node whose common prefix
// starts there or a page whose keys all share that byte, which fork puts next to
// a new key.
func forkByte(old *header, d int) byte {
	if isMultiKey(old.objType) {
		return asMultiKey(old).ByteAt(0, d-asMultiKey(old).Base())
	}
	b, _ := firstByte(old, d)
	return b
}

// attachAt hangs leaf l below node h, where rest is l's key from h's child
// byte on: as h's end page if rest is empty, as the child under rest[0]
// otherwise. It returns h or its grown replacement, and the slot that holds
// l.
func attachAt(h *header, rest []byte, l *singleKeyHead) (*header, **header) {
	if len(rest) == 0 {
		h = setEndPage(h, l)
		return h, endPageSlot(h)
	}
	return addChild(h, rest[0], singleKeyHdr(l))
}

// setEndPage makes l the end page of n, which has none yet, growing n into the
// next larger type when its slots are full. It returns n or its replacement.
func setEndPage(n *header, l *singleKeyHead) *header {
	if full(n) {
		n = grow(n)
	}
	setEndPageSlot(n, l)
	return n
}

// addChild inserts child c under byte b, which must not be present yet,
// growing n into the next larger type when it is full. It returns n or its
// replacement, and the slot that holds c.
func addChild(n *header, b byte, c *header) (*header, **header) {
	if full(n) {
		n = grow(n)
	}
	var slot **header
	switch n.objType {
	case kN5:
		x := asN5(n)
		slot = insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
		x.count++
	case kN12:
		x := asN12(n)
		slot = insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
		x.count++
	case kN26, kN58:
		bm, child := bitmapOf(n)
		slot = insertBitmap(n, bm, child, b, c)
	default:
		x := asN256(n)
		x.child[b] = c
		x.total++
		slot = &x.child[b]
	}
	return n, slot
}

// grow copies the full node n, common prefix, children and end page, into the next larger
// type.
func grow(n *header) *header {
	endPage := endPageOf(n)
	var y *header
	switch n.objType {
	case kN5:
		x := asN5(n)
		y = newLike(n, kN12)
		z := asN12(y)
		copy(z.keys[:], x.keys[:x.count])
		copy(z.child[:], x.child[:x.count])
	case kN12:
		x := asN12(n)
		y = newLike(n, kN26)
		z := asN26(y)
		copy(z.child[:], x.child[:x.count])
		for _, k := range x.keys[:x.count] {
			swar.Set(&z.bitmap, k)
		}
	case kN26:
		x := asN26(n)
		y = newLike(n, kN58)
		z := asN58(y)
		z.bitmap = x.bitmap
		copy(z.child[:], x.child[:x.count])
	default:
		x := asN58(n)
		y = newLike(n, kN256)
		z := asN256(y)
		z.total = uint16(x.count)
		i := 0
		for k := range 256 {
			if swar.Has(&x.bitmap, byte(k)) {
				z.child[k] = x.child[i]
				i++
			}
		}
		y.count = 255
	}
	setEndPageSlot(y, endPage)
	return y
}

// insertBitmap inserts child c under byte b into a bitmap node with room and
// returns its slot.
func insertBitmap(n *header, bm *[4]uint64, child []*header, b byte, c *header) **header {
	r := swar.Rank(bm, b)
	copy(child[r+1:int(n.count)+1], child[r:n.count])
	child[r] = c
	swar.Set(bm, b)
	n.count++
	return &child[r]
}

// insertSorted inserts child c under byte b into a sorted node with room and
// returns its slot.
func insertSorted(keys []byte, child []*header, count int, b byte, c *header) **header {
	i := 0
	for i < count && keys[i] < b {
		i++
	}
	copy(keys[i+1:count+1], keys[i:count])
	copy(child[i+1:count+1], child[i:count])
	keys[i] = b
	child[i] = c
	return &child[i]
}
