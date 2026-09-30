package art

import (
	"bytes"

	"github.com/TomTonic/multimap/internal/swar"
)

// upsert returns the slot that holds the leaf of key, creating the leaf with
// nl when it is missing. The caller may replace the leaf in that slot, as a
// flat leaf does when it grows.
//
// It descends like find, checking compressed paths and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, nl newLeafFunc) **header {
	loc, depth := &t.root, 0
	for {
		n := *loc
		if n == nil {
			*loc = leafHdr(nl(key, depth))
			t.size++
			return loc
		}
		if isLeaf(n.kind) {
			return t.splitLeaf(loc, asLeaf(n), key, depth, nl)
		}
		if n.plen > 0 {
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, depth) {
				pl = n.pathLen()
				if !pathMatches(n, pl, key, depth) {
					mis, _ := pathLcp(n, pl, key[depth:])
					return t.splitPrefix(loc, n, mis, key, depth, nl)
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

// splitLeaf handles an insert that reaches leaf l at depth: either it is the
// key's leaf, or both keys go below a new node holding their common path. l
// keeps its base and moves below the new node.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, nl newLeafFunc) **header {
	ls, rest := l.from(depth), key[depth:]
	if len(key) == l.keyLen() && bytes.Equal(ls, rest) {
		return loc
	}
	p := swar.Lcp(ls, rest)
	nn := newNode(kN5, p)
	storePath(nn, rest[:p])
	h, _ := attachAt(nn, ls[p:], l)
	h, slot := attachAt(h, rest[p:], nl(key, depth+p+min(1, len(rest)-p)))
	*loc = h
	t.size++
	return slot
}

// splitPrefix handles an insert whose key leaves n's compressed path after
// mis bytes: a new node takes the common part, with n and the new leaf below.
func (t *Tree) splitPrefix(loc **header, n *header, mis int, key []byte, depth int, nl newLeafFunc) **header {
	var buf [pathBuf]byte
	pk := appendPath(buf[:0], n) // a copy: n's path changes below
	nn := newNode(kN5, mis)
	storePath(nn, pk[:mis])
	h, _ := addChild(nn, pk[mis], withPath(n, pk[mis+1:]))
	rest := key[depth+mis:]
	h, slot := attachAt(h, rest, nl(key, depth+mis+min(1, len(rest))))
	*loc = h
	t.size++
	return slot
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
