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

// del deletes key from the subtree at *loc, whose common prefix starts at
// key[pathLen:], and reports whether it was there. On the way back up,
// every node on the path shrinks to the smallest type that fits and collapses
// when it no longer branches, so the tree after a delete has the shape it
// would have had if the key had never been inserted.
func del(loc **header, key []byte, pathLen int, rk rekeyFunc) bool {
	n := *loc
	if n == nil {
		return false
	}
	if isPage(n.objType) {
		// A multi-key page holds no entry that this removes: the typed map removes
		// the entries of a page itself (mkkey.go).
		if !isSingleKey(n.objType) || !asSingleKey(n).matches(key) {
			return false
		}
		*loc = nil
		return true
	}
	pl := n.prefixLen()
	if pl != 0 && !prefixMatches(n, pl, key, pathLen) {
		return false
	}
	d := pathLen + pl
	if d == len(key) {
		// The end page's key is the path to n, which key matched: it is key.
		if endPageOf(n) == nil {
			return false
		}
		setEndPageSlot(n, nil)
	} else {
		b := key[d]
		c := findLoc(n, b)
		if c == nil {
			return false
		}
		if !del(c, key, d+1, rk) {
			return false
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n, key, pathLen, d, rk)
	return true
}

// collapse replaces a byte node n at pathLen, whose common prefix ends at d,
// that no longer branches: without children it becomes its end page, and with
// a single child and no end page it merges into that child, whose common prefix
// grows by n's common prefix plus the child's byte; a multi-key page that cannot take
// those bytes stays below n, which then has one child. key is the key just deleted
// below n, which agrees with every key below n up to d. It returns what should
// stand in n's place.
func collapse(n *header, key []byte, pathLen, d int, rk rekeyFunc) *header {
	switch c := int(n.count); {
	case c == 0:
		// The node held only its end page, or one child that was a multi-key page that could
		// not move up and has shrunk to one key (the tree makes it a single-key page): then the
		// child is the key just deleted, and nothing is left. A node with a single child and no
		// end page otherwise collapsed when it fell to one child.
		e := endPageOf(n)
		if e == nil {
			return nil
		}
		return singleKeyHdr(lift(e, key[:d], pathLen, rk))
	case c > 1 || endPageOf(n) != nil:
		return n
	}
	b, c := onlyChild(n)
	if isPage(c.objType) {
		l := asSingleKey(c)
		if !isSingleKey(c.objType) { // a multi-key page takes the bytes in front of its common prefix, if they fit
			if q := rk(l, key[:d], b, pathLen); q != nil {
				return singleKeyHdr(q)
			}
			return n
		}
		if l.base() <= pathLen {
			return c
		}
		return singleKeyHdr(rk(l, key[:d], b, pathLen))
	}
	var buf [prefixBuf]byte // on the stack for the common short prefixes
	p := append(appendPrefix(buf[:0], n), byte(b))
	return withPrefix(c, appendPrefix(p, c))
}

// lift returns leaf l, whose key is k, ready to stand at pathLen: l itself if it
// holds its key from there on, the leaf from rk otherwise.
func lift(l *singleKeyHead, k []byte, pathLen int, rk rekeyFunc) *singleKeyHead {
	if l.base() <= pathLen {
		return l
	}
	return rk(l, k, -1, pathLen)
}

// onlyChild returns the single child of n and the byte it sits under. Only a
// 5-way node can fall to one child: every larger type shrinks into the next
// smaller one well before.
func onlyChild(n *header) (int, *header) {
	x := asN5(n)
	return int(x.keys[0]), x.child[0]
}

// removeChild deletes the child under byte b, which must be present, and
// shrinks n into the next smaller type once it falls to that type's
// threshold. It returns n or its replacement.
func removeChild(n *header, b byte) *header {
	switch n.objType {
	case kN26, kN58:
		bm, child := bitmapOf(n)
		r := swar.Rank(bm, b)
		copy(child[r:n.count-1], child[r+1:n.count])
		n.count--
		child[n.count] = nil
		bm[b>>6] &^= uint64(1) << (b & 63)
		switch {
		case n.objType == kN58 && n.count <= shrink58:
			x := asN58(n)
			y := newLike(n, kN26)
			z := asN26(y)
			z.bitmap = x.bitmap
			copy(z.child[:], x.child[:x.count])
			setEndPageSlot(y, endPageOf(n))
			return y
		case n.objType == kN26 && n.count <= shrink26:
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
			setEndPageSlot(y, endPageOf(n))
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
	if n.objType == kN12 && n.count <= shrink12 {
		y := newLike(n, kN5)
		z := asN5(y)
		copy(z.keys[:], keys[:last])
		copy(z.child[:], child[:last])
		setEndPageSlot(y, endPageOf(n))
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
	setEndPageSlot(y, endPageOf(&x.header))
	return y
}
