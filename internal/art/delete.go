package art

import (
	"bytes"

	"github.com/TomTonic/multimap/internal/swar"
)

// remove deletes the leaf of key and returns it, or nil if key is absent.
func (t *Tree) remove(key []byte) *leafHead {
	l := del(&t.root, key, 0)
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
func del(loc **header, key []byte, depth int) *leafHead {
	n := *loc
	if n == nil {
		return nil
	}
	if isLeaf(n.kind) {
		l := asLeaf(n)
		if !bytes.Equal(l.key(), key) {
			return nil
		}
		*loc = nil
		return l
	}
	pl := n.pathLen(depth)
	if pl != 0 && !swar.Match8(&n.prefix, pl, key, depth) && !swar.MatchPrefix(&n.prefix, pl, key, depth) {
		return nil
	}
	d := depth + pl
	var l *leafHead
	if d == len(key) {
		if l = termOf(n); l == nil || !bytes.Equal(l.key(), key) {
			return nil
		}
		setTermSlot(n, nil)
	} else {
		b := key[d]
		c := findLoc(n, b)
		if c == nil {
			return nil
		}
		if l = del(c, key, d+1); l == nil {
			return nil
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n, pl, d+1)
	return l
}

// collapse replaces an inner node that no longer branches: without children
// it becomes its term leaf, and with a single child and no term it merges
// into that child, whose compressed path grows by n's path of pl bytes plus
// the child's byte. The child's path starts at key depth cd. It returns what
// should stand in n's place.
func collapse(n *header, pl, cd int) *header {
	switch {
	case n.count == 0:
		// The node held only its term leaf, which holds its full key (lazy
		// expansion). A node without a term never gets here: it collapsed
		// when it fell to one child.
		return leafHdr(termOf(n))
	case n.count > 1 || termOf(n) != nil:
		return n
	}
	b, c := onlyChild(n)
	if isLeaf(c.kind) {
		return c
	}
	cl := c.pathLen(cd)
	var buf [swar.PrefixLen]byte
	m := copy(buf[:], n.prefix[:min(pl, swar.PrefixLen)])
	if m < swar.PrefixLen {
		buf[m] = b
		m++
		copy(buf[m:], c.prefix[:min(cl, swar.PrefixLen)])
	}
	c.prefix = buf
	c.plen = uint16(min(pl+1+cl, longPath))
	return c
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
			y := &node26{header: x.header, bitmap: x.bitmap}
			y.kind = kN26
			copy(y.child[:], x.child[:x.count])
			setTermSlot(&y.header, termOf(n))
			return &y.header
		case n.kind == kN26 && n.count <= shrink26:
			x := asN26(n)
			y := &node12{header: x.header}
			y.kind = kN12
			copy(y.child[:], x.child[:x.count])
			i := 0
			for k := range 256 {
				if swar.Has(&x.bitmap, byte(k)) {
					y.keys[i] = byte(k)
					i++
				}
			}
			setTermSlot(&y.header, termOf(n))
			return &y.header
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
		y := &node5{header: *n}
		y.kind = kN5
		copy(y.keys[:], keys[:last])
		copy(y.child[:], child[:last])
		setTermSlot(&y.header, termOf(n))
		return &y.header
	}
	return n
}

// n256To58 copies a 256-way node that has fallen to shrink256 byte children
// into a 58-way node.
func n256To58(x *node256) *header {
	y := &node58{header: x.header}
	y.kind, y.count = kN58, uint8(x.total)
	i := 0
	for k, c := range x.child[:256] {
		if c != nil {
			swar.Set(&y.bitmap, byte(k))
			y.child[i] = c
			i++
		}
	}
	setTermSlot(&y.header, termOf(&x.header))
	return &y.header
}
