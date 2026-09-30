package art

import (
	"github.com/TomTonic/multimap/internal/swar"
)

// remove deletes the leaf of key and returns it, or nil if key is absent. rk
// moves a leaf that takes the place of a node above it (see collapse).
func (t *Tree) remove(key []byte, rk rekeyFunc) *leafHead {
	l := del(&t.root, key, 0, rk)
	if l != nil {
		t.size--
	}
	return l
}

// del deletes the leaf of key from the subtree at *loc, whose compressed path
// starts at key depth depth. On the way back up, every node on the path
// shrinks to the smallest kind that fits and collapses when it no longer
// branches, so the tree after a delete has the shape it would have had if the
// key had never been inserted.
func del(loc **header, key []byte, depth int, rk rekeyFunc) *leafHead {
	n := *loc
	if n == nil {
		return nil
	}
	if isLeaf(n.kind) {
		l := asLeaf(n)
		if !l.matches(key) {
			return nil
		}
		*loc = nil
		return l
	}
	pl := n.pathLen()
	if pl != 0 && !pathMatches(n, pl, key, depth) {
		return nil
	}
	d := depth + pl
	var l *leafHead
	if d == len(key) {
		// The term's key is the path to n, which key matched: it is key.
		if l = termOf(n); l == nil {
			return nil
		}
		setTermSlot(n, nil)
	} else {
		b := key[d]
		c := findLoc(n, b)
		if c == nil {
			return nil
		}
		if l = del(c, key, d+1, rk); l == nil {
			return nil
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n, key, depth, d, rk)
	return l
}

// collapse replaces an inner node n at depth, whose path ends at d, that no
// longer branches: without children it becomes its term leaf, and with a
// single child and no term it merges into that child, whose compressed path
// grows by n's path plus the child's byte. key is the key just deleted below
// n, which agrees with every key below n up to d. It returns what should
// stand in n's place.
func collapse(n *header, key []byte, depth, d int, rk rekeyFunc) *header {
	switch {
	case n.count == 0:
		// The node held only its term leaf. A node without a term never gets
		// here: it collapsed when it fell to one child.
		return leafHdr(lift(termOf(n), key[:d], depth, rk))
	case n.count > 1 || termOf(n) != nil:
		return n
	}
	b, c := onlyChild(n)
	if isLeaf(c.kind) {
		l := asLeaf(c)
		if l.base() <= depth {
			return c
		}
		k := append(append(append(make([]byte, 0, l.keyLen()), key[:d]...), b), l.from(d+1)...)
		return leafHdr(rk(l, k, depth))
	}
	p := append(appendPath(nil, n), b)
	return withPath(c, appendPath(p, c))
}

// lift returns leaf l, whose key is k, ready to stand at depth: l itself if it
// holds its key from there on, a new leaf from rk otherwise.
func lift(l *leafHead, k []byte, depth int, rk rekeyFunc) *leafHead {
	if l.base() <= depth {
		return l
	}
	return rk(l, k, depth)
}

// onlyChild returns the single child of n. Only a 5-way node can fall to one
// child: every larger kind shrinks into the next smaller one well before.
func onlyChild(n *header) (byte, *header) {
	x := asN5(n)
	return x.keys[0], x.child[0]
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
