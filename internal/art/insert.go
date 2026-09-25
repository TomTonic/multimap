package art

import (
	"bytes"
	"slices"

	"github.com/TomTonic/multimap/internal/swar"
)

// spot is where upsert left a key: in a leaf, or at position i of the page
// that the slot at points to, whose path starts at depth.
type spot struct {
	leaf    *leafHead // the key's leaf, or nil when the key is in a page
	at      **header
	i       int
	depth   int
	created bool // the key is new; in a page it already holds the value
}

// upsert returns where key is, creating it when it is missing: in a page with
// the raw value v when the key fits one (see pageable), else as a leaf made
// by nl, which the caller fills.
//
// It descends like find, checking compressed paths and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, v uint64, nl newLeafFunc) spot {
	loc, depth := &t.root, 0
	for {
		n := *loc
		if n == nil {
			*loc = t.newChild(key, v, nl)
			return t.created(*loc)
		}
		if n.kind == kPage {
			return t.upsertPage(loc, key, depth, v, nl)
		}
		if n.kind == kLeaf {
			return t.splitLeaf(loc, asLeaf(n), key, depth, v, nl)
		}
		if n.plen > 0 {
			// MatchPrefix settles paths of up to 8 bytes; a longer path, or
			// one that does not match, needs the position of the mismatch.
			if n.plen > 8 || !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				var buf [8]byte
				pk := fullPrefix(n, depth, &buf)
				if mis := swar.Lcp(pk, key[depth:]); mis < len(pk) {
					return t.splitPrefix(loc, n, pk, mis, key, depth, v, nl)
				}
			}
			depth += int(n.plen)
		}
		if depth == len(key) {
			if n.term == nil {
				n.term = nl(key)
				t.size++
				return spot{leaf: n.term, created: true}
			}
			return spot{leaf: n.term}
		}
		b := key[depth]
		c := findLoc(n, b)
		if c == nil {
			nc := t.newChild(key, v, nl)
			*loc = addChild(n, b, nc)
			return t.created(nc)
		}
		loc, depth = c, depth+1
	}
}

// newChild returns a new page holding key with the raw value v when the key
// fits a page, else a new leaf made by nl.
func (t *Tree) newChild(key []byte, v uint64, nl newLeafFunc) *header {
	if t.small && len(key) <= 8 {
		p := newPage(0)
		p.count, p.klen = 1, uint8(len(key))
		p.heads()[0], p.vals()[0] = keyWord(key), v
		return pageHdr(p)
	}
	return leafHdr(nl(key))
}

// created counts the new key held by c, a child made by newChild.
func (t *Tree) created(c *header) spot {
	t.size++
	if c.kind == kLeaf {
		return spot{leaf: asLeaf(c), created: true}
	}
	return spot{created: true}
}

// upsertPage handles a key whose descent reaches the page at *loc: the key is
// there, or goes in, or the subtree is rebuilt with it (see build) because
// the page is full or the key does not fit it.
func (t *Tree) upsertPage(loc **header, key []byte, depth int, v uint64, nl newLeafFunc) spot {
	p := asPage(*loc)
	if len(key) == int(p.klen) {
		w := keyWord(key)
		i, ok := p.search(w)
		if ok {
			return spot{at: loc, i: i, depth: depth}
		}
		if int(p.count) < pageCaps[len(pageCaps)-1] {
			*loc = pageHdr(p.insertAt(i, w, v))
			t.size++
			return spot{created: true}
		}
	}
	it := item{key: key, val: v}
	if !t.small || len(key) > 8 {
		it.leaf = nl(key)
	}
	items := pageItems(p)
	i, _ := slices.BinarySearchFunc(items, key, func(x item, k []byte) int { return bytes.Compare(x.key, k) })
	*loc = t.build(slices.Insert(items, i, it), depth)
	t.size++
	sp := t.upsert(key, v, nl) // finds the key in the rebuilt subtree
	sp.created = true
	return sp
}

// promote moves the key at sp out of its page into a leaf that holds its
// value, because the key is about to get a second one, and returns the leaf.
// The page's subtree is rebuilt around it (see build).
func (t *Tree) promote(sp spot) *leafHead {
	items := pageItems(asPage(*sp.at))
	it := &items[sp.i]
	it.leaf = t.mk(it.key, it.val)
	*sp.at = t.build(items, sp.depth)
	return it.leaf
}

// splitLeaf handles an insert that reaches leaf l: either it is the key's
// leaf, or both keys go below a new node holding their common path.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, v uint64, nl newLeafFunc) spot {
	lk := l.key()
	if bytes.Equal(lk, key) {
		return spot{leaf: l}
	}
	p := swar.Lcp(lk[depth:], key[depth:])
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(key[depth:depth+p], p)
	h := attach(&nn.header, lk, depth+p, l)
	return t.attachNew(loc, h, key, depth+p, v, nl)
}

// splitPrefix handles an insert whose key leaves n's compressed path pk after
// mis bytes: a new node takes the common part, with n and the new key below.
func (t *Tree) splitPrefix(loc **header, n *header, pk []byte, mis int, key []byte, depth int, v uint64, nl newLeafFunc) spot {
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(pk[:mis], mis)
	old := pk[mis]
	n.setPrefix(pk[mis+1:], len(pk)-mis-1) // pk is a copy or a leaf key: safe to read while n changes
	h := addChild(&nn.header, old, n)
	return t.attachNew(loc, h, key, depth+mis, v, nl)
}

// attachNew adds the new key below node h at key depth d, as h's term leaf if
// it ends there, else as a new child, and stores h or its grown replacement
// in *loc.
func (t *Tree) attachNew(loc **header, h *header, key []byte, d int, v uint64, nl newLeafFunc) spot {
	if d == len(key) {
		h.term = nl(key)
		*loc = h
		t.size++
		return spot{leaf: h.term, created: true}
	}
	c := t.newChild(key, v, nl)
	*loc = addChild(h, key[d], c)
	return t.created(c)
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
