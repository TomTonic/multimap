package art

import (
	"github.com/TomTonic/multimap/internal/swar"
)

// The outcomes of del.
const (
	absent    = iota // the key is not there
	deleted          // the key is gone
	keptLeaf         // the key has a leaf, which del left in place
	keptEntry        // the key's entry in a page holds another value, and is still there
)

// remove deletes key and reports whether it was there. rk moves a leaf that
// takes the place of a node above it (see collapse).
func (t *Tree) remove(key []byte, rk rekeyFunc) bool {
	ok := del(&t.root, key, 0, rk, nil) == deleted
	if ok {
		t.size--
	}
	return ok
}

// removeRaw deletes key if it has one value in a page, which is the raw word
// want: in one descent, where a lookup and a delete would make two. It
// reports absent, deleted, keptEntry for another value, or keptLeaf if key has
// a leaf, which the caller removes values from. rk is as in remove.
func (t *Tree) removeRaw(key []byte, want uint64, rk rekeyFunc) int8 {
	r := del(&t.root, key, 0, rk, &want)
	if r == deleted {
		t.size--
	}
	return r
}

// del deletes key from the subtree at *loc, whose common prefix starts at
// key[pathLen:], and reports whether it was there. On the way back up,
// every node on the path shrinks to the smallest kind that fits and collapses
// when it no longer branches, so the tree after a delete has the shape it
// would have had if the key had never been inserted. Below a range node, a
// range whose child is gone goes to its neighbour, and a page merges with a
// neighbouring page once both are thin (see rMerge).
//
// If want is not nil, del deletes only a key that is in a page with the raw
// value *want (see removeRaw).
func del(loc **header, key []byte, pathLen int, rk rekeyFunc, want *uint64) int8 {
	n := *loc
	if n == nil {
		return absent
	}
	if n.kind <= maxPageByte {
		if isPage(n.kind) {
			return delFromPage(loc, key, want)
		}
		switch {
		case !asLeaf(n).matches(key):
			return absent
		case want != nil:
			return keptLeaf
		}
		*loc = nil
		return deleted
	}
	pl := n.prefixLen()
	if pl != 0 && !prefixMatches(n, pl, key, pathLen) {
		return absent
	}
	d := pathLen + pl
	switch {
	case d == len(key):
		// The end page's key is the path to n, which key matched: it is key.
		if endPageOf(n) == nil {
			return absent
		}
		if want != nil {
			return keptLeaf
		}
		setEndPageSlot(n, nil)
	case isRange(n.kind):
		r := asR(n)
		i := r.index(key[d])
		c := &r.children()[i]
		if r := del(c, key, d, rk, want); r != deleted {
			return r
		}
		switch {
		case *c == nil:
			n = rRemove(n, i)
		case isPage((*c).kind) && asPage(*c).Thin():
			// Two pages merge only if they fit one page together, so one that is
			// not thin does not need its neighbours looked at.
			n = rMerge(n, i)
		}
	default:
		b := key[d]
		c := findLoc(n, b)
		if c == nil {
			return absent
		}
		if r := del(c, key, d+1, rk, want); r != deleted {
			return r
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n, key, pathLen, d, rk)
	return deleted
}

// delFromPage deletes key from the page at *loc, unless want is not nil and
// the key's value is not *want, and reports what it found (see del).
func delFromPage(loc **header, key []byte, want *uint64) int8 {
	p := asPage(*loc)
	i, ok := p.FindIn(key)
	switch {
	case !ok:
		return absent
	case want != nil && p.Val(i) != *want:
		return keptEntry
	}
	*loc = pageHdr(p.DeleteAt(i))
	return deleted
}

// collapse replaces a byte or range node n at pathLen, whose common prefix ends at d,
// that no longer branches: without children it becomes its end page, and with
// a single child and no end page it merges into that child, whose common prefix
// grows by n's common prefix plus the child's byte (which a range node's child already
// starts with). key is the key just deleted below n, which agrees with every
// key below n up to d. It returns what should stand in n's place.
func collapse(n *header, key []byte, pathLen, d int, rk rekeyFunc) *header {
	switch c := childCount(n); {
	case c == 0:
		// The node held only its end page, or nothing: a range node whose only
		// page could not move up (see pageUp) stays, and its last key may go.
		if endPageOf(n) == nil {
			return nil
		}
		return leafHdr(lift(endPageOf(n), key[:d], pathLen, rk))
	case c > 1 || endPageOf(n) != nil:
		return n
	}
	b, c := onlyChild(n)
	if c.kind <= maxPageByte {
		if isPage(c.kind) {
			return pageUp(n, c, key, pathLen)
		}
		l := asLeaf(c)
		if l.base() <= pathLen {
			return c
		}
		return leafHdr(rk(l, key[:d], b, pathLen))
	}
	var buf [prefixBuf]byte // on the stack for the common short prefixes
	p := appendPrefix(buf[:0], n)
	if b >= 0 {
		p = append(p, byte(b))
	}
	return withPrefix(c, appendPrefix(p, c))
}

// pageUp returns what should stand in the place of the range node n at pathLen,
// which has no end page and one child, page c: the page, which then starts at pathLen,
// or n itself if the page cannot take the bytes of n's common prefix in front of its
// keys.
func pageUp(n, c *header, key []byte, pathLen int) *header {
	p := asPage(c)
	if p.Base() <= pathLen {
		return c
	}
	if q := p.Rebase(pathLen, key[pathLen:p.Base()]); q != nil {
		return pageHdr(q)
	}
	return n
}

// childCount returns the number of byte children of a byte node, or of ranges
// of a range node.
func childCount(n *header) int {
	if isRange(n.kind) {
		return int(asR(n).n)
	}
	return int(n.count)
}

// lift returns leaf l, whose key is k, ready to stand at pathLen: l itself if it
// holds its key from there on, the leaf from rk otherwise.
func lift(l *leafHead, k []byte, pathLen int, rk rekeyFunc) *leafHead {
	if l.base() <= pathLen {
		return l
	}
	return rk(l, k, -1, pathLen)
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
			setEndPageSlot(y, endPageOf(n))
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
	if n.kind == kN12 && n.count <= shrink12 {
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
