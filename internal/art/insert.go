package art

import (
	"bytes"
	"slices"

	"github.com/TomTonic/multimap/internal/swar"
)

// spot is where upsert left a key that already was in a page: at position i of
// the page that the slot at loc points to, whose path starts at depth. A page
// below a range node is child pi of the node at *par.
type spot struct {
	loc   **header
	i     int
	depth int
	par   **header
	pi    int
}

// upsert returns the slot that holds the key's leaf or page, creating the key
// when it is missing: in a page with the raw value v when the key fits one (see
// pageable) and does not go below an inner node (see settle.go), else as a leaf
// made by nl, which the caller fills. The caller may replace a leaf in its
// slot, as a flat leaf does when it grows. It tells a key it has created by
// the tree's size, and, for a key that was in a page, where in t.at.
//
// It descends like find, checking compressed paths and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, v uint64, nl newLeafFunc) **header {
	loc, depth := &t.root, 0
	if t.root == nil {
		t.chooseKeyLen(key)
	}
	var par **header // the range node *loc is a child of, or nil
	pi := 0
	inner := false // *loc is below an inner node, where no pages go
	for {
		n := *loc
		if n == nil {
			*loc = t.newChild(key, v, nl, depth)
			t.size++
			return loc
		}
		if n.kind <= kLastPage {
			if n.kind > kLastLeaf {
				return t.upsertPage(loc, par, pi, key, depth, v, nl)
			}
			return t.splitLeaf(loc, asLeaf(n), key, depth, v, nl, inner)
		}
		if n.plen > 0 {
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, depth) {
				pl = n.pathLen()
				if !pathMatches(n, pl, key, depth) {
					mis, _ := pathLcp(n, pl, key[depth:])
					return t.splitPrefix(loc, n, mis, key, depth, v, nl)
				}
			}
			depth += pl
		}
		if depth == len(key) {
			if termOf(n) == nil {
				*loc = setTerm(n, nl(key, depth))
				t.size++
			}
			return termSlot(*loc)
		}
		b := key[depth]
		if isRange(n.kind) {
			x := asR(n)
			i := x.index(b)
			c := &x.children()[i]
			if cb, ok := firstByte(*c, depth); ok && cb != b {
				// The child's keys all have byte cb: the key gets a range of
				// its own, next to it.
				nc := t.newChild(key, v, nl, depth)
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
		par, inner = nil, true
		c := findLoc(n, b)
		if c == nil {
			var slot **header
			*loc, slot = addChild(n, b, leafHdr(nl(key, depth+1)))
			t.size++
			return slot
		}
		loc, depth = c, depth+1
	}
}

// hasPages reports whether the tree may hold pages: its root is a page or a
// range node. Below a root that is an inner node there are none, and there
// never are (see settle.go).
func (t *Tree) hasPages() bool {
	return t.root != nil && (isPage(t.root.kind) || isRange(t.root.kind))
}

// chooseKeyLen decides, when key is the first key of an empty tree, how long
// the keys of its pages are: as long as key, if it is short enough for a page. A
// tree whose keys all have the length of its first one, such as a tree of
// integers, holds them in pages; keys of other lengths get leaves.
func (t *Tree) chooseKeyLen(key []byte) {
	t.pk = 0
	if t.small && len(key) <= maxPageKey {
		t.pk = uint8(len(key)) + 1
	}
}

// pageable reports whether key may go into a page.
func (t *Tree) pageable(key []byte) bool { return t.pk != 0 && len(key) == int(t.pk)-1 }

// firstByte returns byte depth of every key below c, a child of a range
// node, if they all share it: the byte of a leaf, or the first path byte of a
// node. A page's keys may differ there.
func firstByte(c *header, depth int) (byte, bool) {
	switch {
	case isLeaf(c.kind):
		return asLeaf(c).from(depth)[0], true
	case isPage(c.kind) || c.plen == 0:
		return 0, false
	}
	return c.prefix[0], true
}

// newChild returns a new page holding key with the raw value v when the key
// fits a page, else a new leaf made by nl, for a child at depth.
func (t *Tree) newChild(key []byte, v uint64, nl newLeafFunc, depth int) *header {
	if !t.pageable(key) {
		return leafHdr(nl(key, depth))
	}
	p := newPage(0)
	p.count, p.klen = 1, uint8(len(key))
	w := keyWord(key)
	p.heads()[0], p.vals()[0], p.bloom = w, v, bloomBit(w)
	return pageHdr(p)
}

// upsertPage handles a key whose descent reaches the page at *loc: the key is
// there, or goes in, or the subtree is rebuilt with it (see build) because
// the page is full or the key does not fit it.
func (t *Tree) upsertPage(loc, par **header, pi int, key []byte, depth int, v uint64, nl newLeafFunc) **header {
	p := asPage(*loc)
	full := false // the key fits the page but for its room
	if t.pageable(key) {
		w := keyWord(key)
		i, ok := p.search(w)
		switch {
		case ok:
			t.at = spot{loc: loc, i: i, depth: depth, par: par, pi: pi}
			return loc
		case int(p.count) < pageCaps[len(pageCaps)-1]:
			*loc = pageHdr(p.insertAt(i, w, v))
			t.size++
			return loc
		}
		full = true
	}
	if full && par != nil && splitFull(loc, par, pi, depth) {
		r := asR(*par)
		i := r.index(key[depth])
		return t.upsertPage(&r.children()[i], par, i, key, depth, v, nl)
	}
	it := item{key: key, val: v}
	if !t.pageable(key) {
		it = item{key: key, leaf: nl(key, depth)}
	}
	items := pageItems(p)
	i, _ := slices.BinarySearchFunc(items, key, func(x item, k []byte) int { return bytes.Compare(x.key, k) })
	t.replace(loc, par, pi, slices.Insert(items, i, it), depth)
	t.size++
	return t.upsert(key, v, nl) // finds the key in the rebuilt subtree
}

// splitFull splits the full page at *loc, child pi of the range node at
// *par, whose keys start at depth, at the byte boundary nearest its middle
// (see ranges): the page keeps the keys before it and a new page takes the
// others, with a range of its own. Unlike a rebuild, it copies half a page
// and allocates one, and the keys keep their encoding: a K page's half keeps
// the shared prefix, which its keys may share beyond. It reports false and
// changes nothing when all keys share their byte at depth.
func splitFull(loc, par **header, pi, depth int) bool {
	p := asPage(*loc)
	n := int(p.count)
	at := p.byteAt(depth)
	if at(0) == at(n-1) {
		return false // all keys share byte depth
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
	left, right := p.split(s)
	*par = rSplice(*par, pi, []rng{{0, pageHdr(left)}, {rb, pageHdr(right)}})
	return true
}

// byteAt returns a function that returns byte depth of key i of p, below a
// range node at depth.
func (p *pageHead) byteAt(depth int) func(i int) byte {
	h, sh := p.keys(), 56-8*depth
	return func(i int) byte { return byte(h[i] >> sh) }
}

// replace puts items, the keys of the page at *at, in the page's place: as a
// subtree (see build), or, for a page below a range node, as ranges that
// take over the page's range (see ranges).
func (t *Tree) replace(at, par **header, pi int, items []item, depth int) {
	if par == nil {
		*at = t.build(items, depth)
		return
	}
	*par = rSplice(*par, pi, t.ranges(items, depth, nil))
}

// promote gives the key at sp, which has one value in its page, the leaf
// l, which holds its key from sp.depth on and its values, and lets the subtree
// around it fall back to inner nodes if keys with several values crowd it
// (see settle).
func (t *Tree) promote(sp spot, l *leafHead, key []byte) {
	t.place(sp, l, key)
	t.settle(key)
}

// place puts the leaf l of key, the key at sp, in the key's place. A page
// cannot hold the leaf, so it gets a range of its own: the page is cut around
// the key when the key is alone with its byte below a range node, else the
// page's keys are rebuilt with the leaf in the key's place (see ranges), which
// gives it a range or a subtree of its own.
func (t *Tree) place(sp spot, l *leafHead, key []byte) {
	p, i, n := asPage(*sp.loc), sp.i, int(asPage(*sp.loc).count)
	if n == 1 {
		*sp.loc = leafHdr(l) // the page held only this key
		return
	}
	// Below a range node, a key alone with its byte gets a range of its own
	// by cutting the page around it.
	if at := p.byteAt(sp.depth); sp.par != nil &&
		(i == 0 || at(i-1) != at(i)) && (i == n-1 || at(i+1) != at(i)) {
		b := at(i)
		var after byte
		if i < n-1 {
			after = at(i + 1)
		}
		var rs []rng
		if i > 0 {
			var left *pageHead
			left, p = p.split(i)
			rs = append(rs, rng{0, pageHdr(left)})
		}
		rs = append(rs, rng{b, leafHdr(l)})
		if i < n-1 {
			_, right := p.split(1)
			rs = append(rs, rng{after, pageHdr(right)})
		}
		*sp.par = rSplice(*sp.par, sp.pi, rs)
		return
	}
	items := pageItems(p)
	items[sp.i] = item{key: key, leaf: l}
	t.replace(sp.loc, sp.par, sp.pi, items, sp.depth)
}

// splitLeaf handles an insert that reaches leaf l at depth: either it is the
// key's leaf, or both keys go below a new node holding their common path, a
// range node if the new key goes into a page (unless l is below an inner node),
// else an inner node. l keeps its base and moves below the new node.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, v uint64, nl newLeafFunc, inner bool) **header {
	ls, rest := l.from(depth), key[depth:]
	if len(key) == l.keyLen() && bytes.Equal(ls, rest) {
		return loc
	}
	p := swar.Lcp(ls, rest)
	if !inner && t.pageable(key) {
		d := depth + p
		return t.fork(loc, leafHdr(l), d == l.keyLen(), key, depth, d, v, nl)
	}
	nn := newNode(kN5, p)
	storePath(nn, rest[:p])
	h, _ := attachAt(nn, ls[p:], l)
	h, slot := attachAt(h, rest[p:], nl(key, depth+p+min(1, len(rest)-p)))
	*loc = h
	t.size++
	return slot
}

// splitPrefix handles an insert whose key leaves n's compressed path after
// mis bytes: a new node takes the common part, with n and the new key below,
// a range node if n is one, else an inner node.
func (t *Tree) splitPrefix(loc **header, n *header, mis int, key []byte, depth int, v uint64, nl newLeafFunc) **header {
	var buf [pathBuf]byte
	pk := appendPath(buf[:0], n) // a copy: n's path changes below
	if isRange(n.kind) {
		// Below a range node, n keeps the byte it branches on in its path.
		return t.fork(loc, withPath(n, pk[mis:]), false, key, depth, depth+mis, v, nl)
	}
	nn := newNode(kN5, mis)
	storePath(nn, pk[:mis])
	h, _ := addChild(nn, pk[mis], withPath(n, pk[mis+1:]))
	rest := key[depth+mis:]
	h, slot := attachAt(h, rest, nl(key, depth+mis+min(1, len(rest))))
	*loc = h
	t.size++
	return slot
}

// fork puts a range node with the path key[depth:d] at *loc, in a tree with
// pages, that holds old and the new key: old is a leaf whose key ends at d
// (oend), or its keys continue with one byte at d, as the new key does unless
// it ends there.
func (t *Tree) fork(loc **header, old *header, oend bool, key []byte, depth, d int, v uint64, nl newLeafFunc) **header {
	path := key[depth:d]
	if d == len(key) {
		*loc = makeR(path, nl(key, d), []rng{{0, old}})
		t.size++
		return termSlot(*loc)
	}
	nc := t.newChild(key, v, nl, d)
	var term *leafHead
	rs := []rng{{0, nc}}
	if oend {
		term = asLeaf(old)
	} else if ob, _ := firstByte(old, d); key[d] < ob {
		rs = append(rs, rng{ob, old})
	} else {
		rs = []rng{{0, old}, {key[d], nc}}
	}
	*loc = makeR(path, term, rs)
	r := asR(*loc)
	t.size++
	return &r.children()[r.index(key[d])]
}

// attachAt hangs leaf l below node h, where rest is l's key from h's child
// byte on: as h's term if rest is empty, as the child under rest[0]
// otherwise. It returns h or its grown replacement, and the slot that holds
// l.
func attachAt(h *header, rest []byte, l *leafHead) (*header, **header) {
	if len(rest) == 0 {
		h = setTerm(h, l)
		return h, termSlot(h)
	}
	return addChild(h, rest[0], leafHdr(l))
}

// setTerm makes l the term leaf of n, which has none yet, growing n into the
// next larger kind when its slots are full. It returns n or its replacement.
func setTerm(n *header, l *leafHead) *header {
	if full(n) {
		n = grow(n)
	}
	setTermSlot(n, l)
	return n
}

// addChild inserts child c under byte b, which must not be present yet,
// growing n into the next larger kind when it is full. It returns n or its
// replacement, and the slot that holds c.
func addChild(n *header, b byte, c *header) (*header, **header) {
	if full(n) {
		n = grow(n)
	}
	var slot **header
	switch n.kind {
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

// grow copies the full node n, path, children and term, into the next larger
// kind.
func grow(n *header) *header {
	term := termOf(n)
	var y *header
	switch n.kind {
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
	setTermSlot(y, term)
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
