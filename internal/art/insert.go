package art

import (
	"bytes"

	"github.com/TomTonic/multimap/internal/swar"
)

// upsert returns the leaf of key, creating it with nl when it is missing.
//
// It descends like find, checking compressed paths and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, nl newLeafFunc) *leafHead {
	loc, depth := &t.root, 0
	for {
		n := *loc
		if n == nil {
			l := nl(key)
			*loc = leafHdr(l)
			t.size++
			return l
		}
		if n.kind == kLeaf {
			return t.splitLeaf(loc, asLeaf(n), key, depth, nl)
		}
		if n.plen > 0 {
			// MatchPrefix settles paths of up to 8 bytes; a longer path, or
			// one that does not match, needs the position of the mismatch.
			if n.plen > 8 || !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				var buf [8]byte
				pk := fullPrefix(n, depth, &buf)
				if mis := swar.Lcp(pk, key[depth:]); mis < len(pk) {
					return t.splitPrefix(loc, n, pk, mis, key, depth, nl)
				}
			}
			depth += int(n.plen)
		}
		if depth == len(key) {
			if n.term == nil {
				n.term = nl(key)
				t.size++
			}
			return n.term
		}
		b := key[depth]
		c := findLoc(n, b)
		if c == nil {
			l := nl(key)
			*loc = addChild(n, b, leafHdr(l))
			t.size++
			return l
		}
		loc, depth = c, depth+1
	}
}

// splitLeaf handles an insert that reaches leaf l: either it is the key's
// leaf, or both keys go below a new node holding their common path.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, nl newLeafFunc) *leafHead {
	lk := l.key()
	if bytes.Equal(lk, key) {
		return l
	}
	p := swar.Lcp(lk[depth:], key[depth:])
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(key[depth:depth+p], p)
	nl2 := nl(key)
	h := attach(&nn.header, lk, depth+p, l)
	*loc = attach(h, key, depth+p, nl2)
	t.size++
	return nl2
}

// splitPrefix handles an insert whose key leaves n's compressed path pk after
// mis bytes: a new node takes the common part, with n and the new leaf below.
func (t *Tree) splitPrefix(loc **header, n *header, pk []byte, mis int, key []byte, depth int, nl newLeafFunc) *leafHead {
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(pk[:mis], mis)
	old := pk[mis]
	n.setPrefix(pk[mis+1:], len(pk)-mis-1) // pk is a copy or a leaf key: safe to read while n changes
	l := nl(key)
	h := addChild(&nn.header, old, n)
	*loc = attach(h, key, depth+mis, l)
	t.size++
	return l
}

// attach hangs leaf l, whose key is k, below node h at key depth d: as h's
// term if k ends there, as a child otherwise. It returns h or its grown
// replacement.
func attach(h *header, k []byte, d int, l *leafHead) *header {
	if d == len(k) {
		h.term = l
		return h
	}
	return addChild(h, k[d], leafHdr(l))
}

// addChild inserts child c under byte b, which must not be present yet,
// growing n into the next larger kind when it is full. It returns n or its
// replacement.
func addChild(n *header, b byte, c *header) *header {
	switch n.kind {
	case kN4:
		x := asN4(n)
		if x.count < 4 {
			insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
			x.count++
			return n
		}
		y := &node11{header: x.header}
		y.kind = kN11
		copy(y.keys[:], x.keys[:4])
		copy(y.child[:], x.child[:])
		return addChild(&y.header, b, c)
	case kN11:
		x := asN11(n)
		if x.count < 11 {
			insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
			x.count++
			return n
		}
		y := &node25{header: x.header}
		y.kind = kN25
		copy(y.child[:], x.child[:])
		for _, k := range x.keys[:11] {
			swar.Set(&y.bitmap, k)
		}
		return addChild(&y.header, b, c)
	case kN25:
		x := asN25(n)
		if x.count < 25 {
			insertBitmap(&x.header, &x.bitmap, x.child[:], b, c)
			return n
		}
		y := &node57{header: x.header, bitmap: x.bitmap}
		y.kind = kN57
		copy(y.child[:], x.child[:])
		return addChild(&y.header, b, c)
	case kN57:
		x := asN57(n)
		if x.count < 57 {
			insertBitmap(&x.header, &x.bitmap, x.child[:], b, c)
			return n
		}
		y := &node256{header: x.header}
		y.kind = kN256
		i := 0
		for k := range 256 {
			if swar.Has(&x.bitmap, byte(k)) {
				y.child[k] = x.child[i]
				i++
			}
		}
		return addChild(&y.header, b, c)
	default:
		x := asN256(n)
		x.child[b] = c
		x.count++
		return n
	}
}

// insertBitmap inserts child c under byte b into a bitmap node with room.
func insertBitmap(n *header, bm *[4]uint64, child []*header, b byte, c *header) {
	r := swar.Rank(bm, b)
	copy(child[r+1:int(n.count)+1], child[r:n.count])
	child[r] = c
	swar.Set(bm, b)
	n.count++
}

func insertSorted(keys []byte, child []*header, count int, b byte, c *header) {
	i := 0
	for i < count && keys[i] < b {
		i++
	}
	copy(keys[i+1:count+1], keys[i:count])
	copy(child[i+1:count+1], child[i:count])
	keys[i] = b
	child[i] = c
}
