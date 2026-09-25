package art

import (
	"bytes"

	"github.com/TomTonic/multimap/internal/swar"
)

// remove deletes key and reports whether it was there.
func (t *Tree) remove(key []byte) bool {
	ok := del(&t.root, key, 0)
	if ok {
		t.size--
	}
	return ok
}

// del deletes key from the subtree at *loc, whose compressed path starts at
// key depth depth, and reports whether it was there. On the way back up,
// every node on the path shrinks to the smallest kind that fits and collapses
// when it no longer branches, so the tree after a delete has the shape it
// would have had if the key had never been inserted.
func del(loc **header, key []byte, depth int) bool {
	n := *loc
	if n == nil {
		return false
	}
	switch n.kind {
	case kPage, kPageN:
		p := asPage(n)
		if len(key) != int(p.klen) {
			return false
		}
		i, ok := p.search(keyWord(key))
		switch {
		case !ok:
		case n.kind == kPage:
			*loc = pageHdr(p.removeAt(i))
		default:
			*loc = pageHdr(p.nRemoveKey(i))
		}
		return ok
	case kLeaf:
		if !bytes.Equal(asLeaf(n).key(), key) {
			return false
		}
		*loc = nil
		return true
	}
	if n.plen != 0 && !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
		return false
	}
	d := depth + int(n.plen)
	if d == len(key) {
		if n.term == nil || !bytes.Equal(n.term.key(), key) {
			return false
		}
		n.term = nil
	} else {
		b := key[d]
		c := findLoc(n, b)
		if c == nil || !del(c, key, d+1) {
			return false
		}
		if *c == nil {
			n = removeChild(n, b)
		}
	}
	*loc = collapse(n)
	return true
}

// collapse replaces an inner node that no longer branches: without children
// it becomes its term leaf, and with a single child and no
// term it merges into that child, whose compressed path grows by n's path
// plus the child's byte. It returns what should stand in n's place.
func collapse(n *header) *header {
	switch {
	case n.count == 0:
		// The node held only its term leaf, which holds its full key (lazy
		// expansion). A node without a term never gets here: it collapsed
		// when it fell to one child.
		return leafHdr(n.term)
	case n.count > 1 || n.term != nil:
		return n
	}
	b, c := onlyChild(n)
	if c.kind <= kPageN { // leaves and pages hold full keys: they just move up
		return c
	}
	var buf [8]byte
	m := copy(buf[:], n.prefix[:min(n.plen, 8)])
	if m < 8 {
		buf[m] = b
		m++
		copy(buf[m:], c.prefix[:min(c.plen, 8)])
	}
	c.prefix = buf
	c.plen += n.plen + 1
	return c
}

// onlyChild returns the single child of n. Only a 4-way node can fall to one
// child: every larger kind shrinks into the next smaller one well before.
func onlyChild(n *header) (byte, *header) {
	x := asN4(n)
	return x.keys[0], x.child[0]
}

// removeChild deletes the child under byte b, which must be present, and
// shrinks n into the next smaller kind once it falls to that kind's
// threshold. It returns n or its replacement.
func removeChild(n *header, b byte) *header {
	switch n.kind {
	case kN25, kN57:
		bm, child := bitmapOf(n)
		r := swar.Rank(bm, b)
		copy(child[r:n.count-1], child[r+1:n.count])
		n.count--
		child[n.count] = nil
		bm[b>>6] &^= uint64(1) << (b & 63)
		switch {
		case n.kind == kN57 && n.count <= shrink57:
			x := asN57(n)
			y := &node25{header: x.header, bitmap: x.bitmap}
			y.kind = kN25
			copy(y.child[:], x.child[:x.count])
			return &y.header
		case n.kind == kN25 && n.count <= shrink25:
			x := asN25(n)
			y := &node11{header: x.header}
			y.kind = kN11
			copy(y.child[:], x.child[:x.count])
			i := 0
			for k := range 256 {
				if swar.Has(&x.bitmap, byte(k)) {
					y.keys[i] = byte(k)
					i++
				}
			}
			return &y.header
		}
		return n
	case kN256:
		x := asN256(n)
		x.child[b] = nil
		x.count--
		if x.count <= shrink256 {
			return n256To57(x)
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
	if n.kind == kN11 && n.count <= shrink11 {
		y := &node4{header: *n}
		y.kind = kN4
		copy(y.keys[:], keys[:last])
		copy(y.child[:], child[:last])
		return &y.header
	}
	return n
}

func n256To57(x *node256) *header {
	y := &node57{header: x.header}
	y.kind = kN57
	i := 0
	for k, c := range x.child {
		if c != nil {
			swar.Set(&y.bitmap, byte(k))
			y.child[i] = c
			i++
		}
	}
	return &y.header
}
