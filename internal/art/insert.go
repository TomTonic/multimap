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
			*loc = leafHdr(nl(key))
			t.size++
			return loc
		}
		if isLeaf(n.kind) {
			return t.splitLeaf(loc, asLeaf(n), key, depth, nl)
		}
		if n.plen > 0 {
			// MatchPrefix settles paths of up to 12 bytes; a longer path, or
			// one that does not match, needs the position of the mismatch.
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, depth) && (pl > swar.PrefixLen || !swar.MatchPrefix(&n.prefix, pl, key, depth)) {
				var buf [swar.PrefixLen]byte
				pk := fullPrefix(n, depth, &buf)
				if mis := swar.Lcp(pk, key[depth:]); mis < len(pk) {
					return t.splitPrefix(loc, n, pk, mis, key, depth, nl)
				}
				pl = len(pk)
			}
			depth += pl
		}
		if depth == len(key) {
			if termOf(n) == nil {
				*loc = setTerm(n, nl(key))
				t.size++
			}
			return termSlot(*loc)
		}
		b := key[depth]
		c := findLoc(n, b)
		if c == nil {
			var slot **header
			*loc, slot = addChild(n, b, leafHdr(nl(key)))
			t.size++
			return slot
		}
		loc, depth = c, depth+1
	}
}

// splitLeaf handles an insert that reaches leaf l: either it is the key's
// leaf, or both keys go below a new node holding their common path.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, nl newLeafFunc) **header {
	lk := l.key()
	if bytes.Equal(lk, key) {
		return loc
	}
	p := swar.Lcp(lk[depth:], key[depth:])
	nn := &node5{}
	nn.kind = kN5
	nn.setPrefix(key[depth:], p)
	nl2 := nl(key)
	h, _ := attach(&nn.header, lk, depth+p, l)
	h, slot := attach(h, key, depth+p, nl2)
	*loc = h
	t.size++
	return slot
}

// splitPrefix handles an insert whose key leaves n's compressed path pk after
// mis bytes: a new node takes the common part, with n and the new leaf below.
func (t *Tree) splitPrefix(loc **header, n *header, pk []byte, mis int, key []byte, depth int, nl newLeafFunc) **header {
	nn := &node5{}
	nn.kind = kN5
	nn.setPrefix(pk[:mis], mis)
	old := pk[mis]
	n.setPrefix(pk[mis+1:], len(pk)-mis-1) // pk is a copy or a leaf key: safe to read while n changes
	l := nl(key)
	h, _ := addChild(&nn.header, old, n)
	h, slot := attach(h, key, depth+mis, l)
	*loc = h
	t.size++
	return slot
}

// attach hangs leaf l, whose key is k, below node h at key depth d: as h's
// term if k ends there, as a child otherwise. It returns h or its grown
// replacement, and the slot that holds l.
func attach(h *header, k []byte, d int, l *leafHead) (*header, **header) {
	if d == len(k) {
		h = setTerm(h, l)
		return h, termSlot(h)
	}
	return addChild(h, k[d], leafHdr(l))
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

// grow copies the full node n, children and term, into the next larger kind.
func grow(n *header) *header {
	term := termOf(n)
	var y *header
	switch n.kind {
	case kN5:
		x := asN5(n)
		z := &node12{header: x.header}
		copy(z.keys[:], x.keys[:x.count])
		copy(z.child[:], x.child[:x.count])
		y = &z.header
		y.kind = kN12
	case kN12:
		x := asN12(n)
		z := &node26{header: x.header}
		copy(z.child[:], x.child[:x.count])
		for _, k := range x.keys[:x.count] {
			swar.Set(&z.bitmap, k)
		}
		y = &z.header
		y.kind = kN26
	case kN26:
		x := asN26(n)
		z := &node58{header: x.header, bitmap: x.bitmap}
		copy(z.child[:], x.child[:x.count])
		y = &z.header
		y.kind = kN58
	default:
		x := asN58(n)
		z := &node256{header: x.header, total: uint16(x.count)}
		i := 0
		for k := range 256 {
			if swar.Has(&x.bitmap, byte(k)) {
				z.child[k] = x.child[i]
				i++
			}
		}
		y = &z.header
		y.kind, y.count = kN256, 255
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
