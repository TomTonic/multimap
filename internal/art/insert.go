package art

import (
	"bytes"
	"slices"
	"unsafe"

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
		if n.kind != kLeaf && n.kind <= kPageS {
			return t.upsertPage(loc, key, depth, v, nl)
		}
		if n.kind == kLeaf {
			return t.splitLeaf(loc, asLeaf(n), key, depth, v, nl)
		}
		if n.plen > 0 {
			// MatchPrefix settles paths of up to 8 bytes; a longer path, or
			// one that does not match, needs the position of the mismatch.
			if n.plen > 8 || !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				var buf keyBuf
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

// pageable reports whether key may go into a page.
func (t *Tree) pageable(key []byte) bool { return t.small && len(key) <= maxPageKey }

// newChild returns a new page holding key with the raw value v when the key
// fits a page, else a new leaf made by nl.
func (t *Tree) newChild(key []byte, v uint64, nl newLeafFunc) *header {
	switch {
	case !t.pageable(key):
		return leafHdr(nl(key))
	case len(key) <= 8:
		p := newPage(0)
		p.count, p.klen = 1, uint8(len(key))
		p.heads()[0], p.vals()[0] = keyWord(key), v
		return pageHdr(p)
	}
	return pageHdr(sPack([]item{{key: key, vals: []uint64{v}}}))
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
	if p.kind == kPageS {
		if p.sMatch(key) {
			s := key[p.base:]
			i, ok := p.sSearch(s)
			if ok {
				return spot{at: loc, i: i, depth: depth}
			}
			if p.sRoom(len(s)) {
				p.sInsertKey(i, s, v)
				t.size++
				return spot{created: true}
			}
		}
	} else if len(key) == int(p.klen) {
		w := keyWord(key)
		i, ok := p.search(w)
		switch {
		case ok:
			return spot{at: loc, i: i, depth: depth}
		case p.kind == kPage && int(p.count) < pageCaps[len(pageCaps)-1]:
			*loc = pageHdr(p.insertAt(i, w, v))
			t.size++
			return spot{created: true}
		case p.kind == kPageN:
			if c := nClassFor(int(p.count)+1, int(p.nv)+1, p.extUsed()); c >= 0 {
				if c > int(p.class) {
					p = p.nResize(c)
				}
				p.nInsertKey(i, w, v)
				*loc = pageHdr(p)
				t.size++
				return spot{created: true}
			}
		}
	}
	it := item{key: key, vals: []uint64{v}}
	if !t.pageable(key) {
		it = item{key: key, leaf: nl(key)}
	}
	items := pageItems(p)
	i, _ := slices.BinarySearchFunc(items, key, func(x item, k []byte) int { return bytes.Compare(x.key, k) })
	*loc = t.build(slices.Insert(items, i, it), depth)
	t.size++
	sp := t.upsert(key, v, nl) // finds the key in the rebuilt subtree
	sp.created = true
	return sp
}

// rebuild rebuilds the subtree of the page at sp after change modified the
// key at sp (see build).
func (t *Tree) rebuild(sp spot, change func(*item)) {
	items := pageItems(asPage(*sp.at))
	change(&items[sp.i])
	*sp.at = t.build(items, sp.depth)
}

// addToSingle gives the key at sp, in a U8-1 page, the second raw value v:
// the page becomes a U8-n page, or, with more keys than one holds, the
// subtree is rebuilt.
func (t *Tree) addToSingle(sp spot, v uint64) {
	p := asPage(*sp.at)
	if int(p.count) <= nLayouts[len(nLayouts)-1].keys {
		q := p.toN(1)
		q.nAddVal(sp.i, v)
		*sp.at = pageHdr(q)
		return
	}
	t.rebuild(sp, func(it *item) { it.vals = append(it.vals, v) })
}

// addInline adds the raw value v to the inline values of the key at sp, in a
// U8-n or S page, which holds fewer than inlineMax. A full U8-n page grows
// into a larger class; beyond the largest, and for a full S page, the subtree
// is rebuilt.
func (t *Tree) addInline(sp spot, v uint64) {
	p := asPage(*sp.at)
	if int(p.nv) == p.layout().vals {
		c := -1
		if p.kind == kPageN {
			c = nClassFor(int(p.count), int(p.nv)+1, p.extUsed())
		}
		if c < 0 {
			t.rebuild(sp, func(it *item) { it.vals = append(it.vals, v) })
			return
		}
		p = p.nResize(c)
		*sp.at = pageHdr(p)
	}
	p.nAddVal(sp.i, v)
}

// externalize moves the inline values of the key at sp, in a U8-n or S page,
// to the external set s, which the caller filled with them and the new value.
// A U8-n page without a free external slot grows into a class with one;
// beyond the largest, and for an S page, the subtree is rebuilt.
func (t *Tree) externalize(sp spot, s unsafe.Pointer) {
	p := asPage(*sp.at)
	if p.extUsed() == p.layout().ext {
		c := -1
		if p.kind == kPageN {
			_, n, _ := p.run(sp.i)
			c = nClassFor(int(p.count), int(p.nv)-n, p.extUsed()+1)
		}
		if c < 0 {
			t.rebuild(sp, func(it *item) { it.vals, it.set = nil, s })
			return
		}
		p = p.nResize(c)
		*sp.at = pageHdr(p)
	}
	p.nExternalize(sp.i, s)
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
