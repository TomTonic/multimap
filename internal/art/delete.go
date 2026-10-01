package art

import (
	"github.com/TomTonic/multimap/internal/swar"
)

// remove deletes key and reports whether it was there. rk moves a leaf that
// takes the place of a node above it (see collapse).
func (t *Tree) remove(key []byte, rk rekeyFunc) bool {
	ok := del(&t.root, key, 0, rk)
	if ok {
		t.size--
	}
	return ok
}

// del deletes key from the subtree at *loc, whose compressed path starts at
// key depth depth, and reports whether it was there. On the way back up,
// every node on the path shrinks to the smallest kind that fits and collapses
// when it no longer branches, so the tree after a delete has the shape it
// would have had if the key had never been inserted. Below a range node, a
// range whose child is gone goes to its neighbour, and a page merges with a
// neighbouring page once both are thin (see rMerge).
func del(loc **header, key []byte, depth int, rk rekeyFunc) bool {
	n := *loc
	if n == nil {
		return false
	}
	if n.kind <= kLastPage {
		if isPage(n.kind) {
			return delFromPage(loc, key, depth)
		}
		if !asLeaf(n).matches(key) {
			return false
		}
		*loc = nil
		return true
	}
	pl := n.pathLen()
	if pl != 0 && !pathMatches(n, pl, key, depth) {
		return false
	}
	d := depth + pl
	switch {
	case d == len(key):
		// The term's key is the path to n, which key matched: it is key.
		if termOf(n) == nil {
			return false
		}
		setTermSlot(n, nil)
	case isRange(n.kind):
		r := asR(n)
		i := r.index(key[d])
		c := &r.children()[i]
		if !del(c, key, d, rk) {
			return false
		}
		switch {
		case *c == nil:
			n = rRemove(n, i)
		case isPage((*c).kind):
			n = rMerge(n, i)
		}
	default:
		b := key[d]
		c := findLoc(n, b)
		if c == nil || !del(c, key, d+1, rk) {
			return false
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n, key, depth, d, rk)
	return true
}

// delFromPage deletes key from the page at *loc, below which the nodes have
// checked the key up to depth, and reports whether it was there.
func delFromPage(loc **header, key []byte, depth int) bool {
	p := asPage(*loc)
	if p.kind == kPageK {
		i, ok, _ := p.kFind(key, depth)
		if ok {
			*loc = pageHdr(p.kRemoveAt(i))
		}
		return ok
	}
	if len(key) != int(p.klen) {
		return false
	}
	i, ok := p.search(keyWord(key))
	if ok {
		*loc = pageHdr(p.removeAt(i))
	}
	return ok
}

// collapse replaces an inner or range node n at depth, whose path ends at d,
// that no longer branches: without children it becomes its term leaf, and with
// a single child and no term it merges into that child, whose compressed path
// grows by n's path plus the child's byte (which a range node's child already
// starts with). key is the key just deleted below n, which agrees with every
// key below n up to d. It returns what should stand in n's place.
func collapse(n *header, key []byte, depth, d int, rk rekeyFunc) *header {
	switch c := childCount(n); {
	case c == 0:
		// The node held only its term leaf. A node without a term never gets
		// here: it collapsed when it fell to one child.
		return leafHdr(lift(termOf(n), key[:d], depth, rk))
	case c > 1 || termOf(n) != nil:
		return n
	}
	b, c := onlyChild(n)
	if c.kind <= kLastPage {
		if isPage(c.kind) {
			return c // pages hold whole keys: they just move up
		}
		l := asLeaf(c)
		if l.base() <= depth {
			return c
		}
		return leafHdr(rk(l, key[:d], b, depth))
	}
	var buf [pathBuf]byte // on the stack for the common short paths
	p := appendPath(buf[:0], n)
	if b >= 0 {
		p = append(p, byte(b))
	}
	return withPath(c, appendPath(p, c))
}

// childCount returns the number of byte children of an inner node, or of ranges
// of a range node.
func childCount(n *header) int {
	if isRange(n.kind) {
		return int(asR(n).n)
	}
	return int(n.count)
}

// lift returns leaf l, whose key is k, ready to stand at depth: l itself if it
// holds its key from there on, the leaf from rk otherwise.
func lift(l *leafHead, k []byte, depth int, rk rekeyFunc) *leafHead {
	if l.base() <= depth {
		return l
	}
	return rk(l, k, -1, depth)
}

// onlyChild returns the single child of n and the byte it sits under, or -1
// for a range node's child. Only a 5-way node or a range node of the smallest
// class can fall to one child: every larger kind shrinks into the next smaller
// one well before.
func onlyChild(n *header) (int, *header) {
	if isRange(n.kind) {
		return -1, asR(n).children()[0]
	}
	x := asN5(n)
	return int(x.keys[0]), x.child[0]
}

// removeChild deletes the child under byte b, which must be present, and
// shrinks n into the next smaller kind once it falls to that kind's
// threshold. It returns n or its replacement.
func removeChild(n *header, b byte) *header {
	switch n.kind {
	case kN26, kN58:
		bm, child := bitmapOf(n)
		r := swar.Rank(bm, b)
		copy(child[r:n.count-1], child[r+1:n.count])
		n.count--
		child[n.count] = nil
		bm[b>>6] &^= uint64(1) << (b & 63)
		switch {
		case n.kind == kN58 && n.count <= shrink58:
			x := asN58(n)
			y := newLike(n, kN26)
			z := asN26(y)
			z.bitmap = x.bitmap
			copy(z.child[:], x.child[:x.count])
			setTermSlot(y, termOf(n))
			return y
		case n.kind == kN26 && n.count <= shrink26:
			x := asN26(n)
			y := newLike(n, kN12)
			z := asN12(y)
			copy(z.child[:], x.child[:x.count])
			i := 0
			for k := range 256 {
				if swar.Has(&x.bitmap, byte(k)) {
					z.keys[i] = byte(k)
					i++
				}
			}
			setTermSlot(y, termOf(n))
			return y
		}
		return n
	case kN256:
		x := asN256(n)
		x.child[b] = nil
		x.total--
		if x.total <= shrink256 {
			return n256To58(x)
		}
		return n
	}
	keys, child := sorted(n)
	i := 0
	for keys[i] != b {
		i++
	}
	last := len(keys) - 1
	copy(keys[i:last], keys[i+1:])
	copy(child[i:last], child[i+1:])
	child[last] = nil
	n.count--
	if n.kind == kN12 && n.count <= shrink12 {
		y := newLike(n, kN5)
		z := asN5(y)
		copy(z.keys[:], keys[:last])
		copy(z.child[:], child[:last])
		setTermSlot(y, termOf(n))
		return y
	}
	return n
}

// n256To58 copies a 256-way node that has fallen to shrink256 byte children
// into a 58-way node.
func n256To58(x *node256) *header {
	y := newLike(&x.header, kN58)
	z := asN58(y)
	y.count = uint8(x.total)
	i := 0
	for k, c := range x.child[:256] {
		if c != nil {
			swar.Set(&z.bitmap, byte(k))
			z.child[i] = c
			i++
		}
	}
	setTermSlot(y, termOf(&x.header))
	return y
}
