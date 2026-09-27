package art

import (
	"bytes"
	"slices"

	"github.com/TomTonic/multimap/internal/swar"
)

// spot is where upsert left a key: in a leaf, or at position i of the page
// that the slot at points to, whose path starts at depth. A page below a
// range node is child pi of the node at *par.
type spot struct {
	leaf    *leafHead // the key's leaf, or nil when the key is in a page
	at      **header
	i       int
	depth   int
	created bool // the key is new; in a page it already holds the value
	par     **header
	pi      int
}

// upsert returns where key is, creating it when it is missing: in a page with
// the raw value v when the key fits one (see pageable) and does not go below
// an inner node (see settle.go), else as a leaf made by nl, which the caller
// fills.
//
// It descends like find, checking compressed paths and searching nodes the
// same fast way, and keeps the slot it came through. A missing key is then
// added right where the descent stopped, without a second traversal.
func (t *Tree) upsert(key []byte, v uint64, nl newLeafFunc) spot {
	loc, depth := &t.root, 0
	var par **header // the range node *loc is a child of, or nil
	pi := 0
	inner := false // *loc is below an inner node, where no pages go
	for {
		n := *loc
		if n == nil {
			*loc = t.newChild(key, v, nl)
			return t.created(*loc)
		}
		if isPage(n) {
			return t.upsertPage(loc, par, pi, key, depth, v, nl)
		}
		if n.kind == kLeaf {
			return t.splitLeaf(loc, asLeaf(n), key, depth, v, nl, inner)
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
		if n.kind == kR {
			x := asR(n)
			i := x.index(b)
			c := &x.children()[i]
			if cb, ok := firstByte(*c, depth); ok && cb != b {
				// The child's keys all have byte cb: the key gets a range of
				// its own, next to it.
				nc := t.newChild(key, v, nl)
				rs := []rng{{0, *c}, {b, nc}}
				if b < cb {
					rs = []rng{{0, nc}, {cb, *c}}
				}
				*loc = rSplice(n, i, rs)
				return t.created(nc)
			}
			par, pi, loc = loc, i, c
			continue
		}
		par, inner = nil, true
		c := findLoc(n, b)
		if c == nil {
			nc := leafHdr(nl(key))
			*loc = addChild(n, b, nc)
			return t.created(nc)
		}
		loc, depth = c, depth+1
	}
}

// pageable reports whether key may go into a page.
func (t *Tree) pageable(key []byte) bool { return t.small && len(key) <= maxPageKey }

// firstByte returns byte depth of every key below c, a child of a range
// node, if they all share it: the byte of a leaf, or the first path byte of a
// node. A page's keys may differ there.
func firstByte(c *header, depth int) (byte, bool) {
	switch {
	case c.kind == kLeaf:
		return asLeaf(c).key()[depth], true
	case isPage(c) || c.plen == 0:
		return 0, false
	}
	return c.prefix[0], true
}

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
	return pageHdr(kPack([]item{{key: key, val: v}}))
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
func (t *Tree) upsertPage(loc, par **header, pi int, key []byte, depth int, v uint64, nl newLeafFunc) spot {
	p := asPage(*loc)
	full := false // the key fits the page but for its room
	if p.kind == kPageK {
		i, found, match := p.kFind(key, depth)
		switch {
		case found:
			return spot{at: loc, i: i, depth: depth, par: par, pi: pi}
		case match && t.pageable(key) && (p.count < p.kcap || int(p.class) < len(kCaps)-1):
			if p.count == p.kcap {
				p = p.kResize(int(p.class) + 1)
				*loc = pageHdr(p)
			}
			p.kInsertAt(i, key, v)
			t.size++
			return spot{created: true}
		}
		full = match && t.pageable(key)
	} else if len(key) == int(p.klen) {
		w := keyWord(key)
		i, ok := p.search(w)
		switch {
		case ok:
			return spot{at: loc, i: i, depth: depth, par: par, pi: pi}
		case int(p.count) < pageCaps[len(pageCaps)-1]:
			*loc = pageHdr(p.insertAt(i, w, v))
			t.size++
			return spot{created: true}
		}
		full = true
	}
	if full && par != nil && splitFull(loc, par, pi, depth) {
		r := asR(*par)
		i := r.index(key[depth])
		return t.upsertPage(&r.children()[i], par, i, key, depth, v, nl)
	}
	it := item{key: key, val: v}
	if !t.pageable(key) {
		it = item{key: key, leaf: nl(key)}
	}
	items := pageItems(p)
	i, _ := slices.BinarySearchFunc(items, key, func(x item, k []byte) int { return bytes.Compare(x.key, k) })
	t.replace(loc, par, pi, slices.Insert(items, i, it), depth)
	t.size++
	sp := t.upsert(key, v, nl) // finds the key in the rebuilt subtree
	sp.created = true
	return sp
}

// splitFull splits the full page at *loc, child pi of the range node at
// *par, whose keys start at depth, at the byte boundary nearest its middle
// (see ranges): the page keeps the keys before it and a new page takes the
// others, with a range of its own. Unlike a rebuild, it copies half a page
// and allocates one, and the keys keep their encoding: a K page's half keeps
// the shared prefix, which its keys may share beyond. It reports false and
// changes nothing when all keys share their byte at depth.
func splitFull(loc, par **header, pi, depth int) bool {
	p := asPage(*loc)
	n := int(p.count)
	at := p.byteAt(depth)
	if at == nil || at(0) == at(n-1) {
		return false // all keys share byte depth
	}
	m := n / 2
	b, lo, hi := at(m), m, m
	for lo > 0 && at(lo-1) == b {
		lo--
	}
	for hi < n && at(hi) == b {
		hi++
	}
	s := lo
	if lo == 0 || (hi < n && hi-m < m-lo) {
		s = hi
	}
	rb := at(s)
	left, right := p.cut(s)
	*par = rSplice(*par, pi, []rng{{0, pageHdr(left)}, {rb, pageHdr(right)}})
	return true
}

// byteAt returns a function that returns byte depth of key i of p, below a
// range node at depth, or nil if all its keys share that byte: a K page
// whose shared prefix reaches beyond depth.
func (p *pageHead) byteAt(depth int) func(i int) byte {
	if p.kind == kPageK {
		if int(p.base) != depth {
			return nil
		}
		h := p.kHeads()
		return func(i int) byte { return byte(h[2*i] >> 56) }
	}
	h, sh := p.keys(), 56-8*depth
	return func(i int) byte { return byte(h[i] >> sh) }
}

// cut returns the page with the keys before s, or its smaller
// replacement, and a new page with the keys from s on (see split, kSplit).
func (p *pageHead) cut(s int) (*pageHead, *pageHead) {
	if p.kind == kPageK {
		return p.kSplit(s)
	}
	return p.split(s)
}

// replace puts items, the keys of the page at *at, in the page's place: as a
// subtree (see build), or, for a page below a range node, as ranges that
// take over the page's range (see ranges).
func (t *Tree) replace(at, par **header, pi int, items []item, depth int) {
	if par == nil {
		*at = t.build(items, depth)
		return
	}
	*par = rSplice(*par, pi, t.ranges(items, depth, nil))
}

// promote gives the key at sp, which has one value in its page, the leaf
// l, which holds its key and its values, and lets the subtree around it fall
// back to inner nodes if keys with several values crowd it (see settle).
func (t *Tree) promote(sp spot, l *leafHead) {
	t.place(sp, l)
	t.settle(l.key())
}

// place puts the leaf l of the key at sp in the key's place. A page cannot
// hold the leaf, so it gets a range of its own: the page is cut around the
// key when the key is alone with its byte below a range node, else the
// page's keys are rebuilt with the leaf in the key's place (see ranges),
// which gives it a range or a subtree of its own.
func (t *Tree) place(sp spot, l *leafHead) {
	p, i, n := asPage(*sp.at), sp.i, int(asPage(*sp.at).count)
	if n == 1 {
		*sp.at = leafHdr(l) // the page held only this key
		return
	}
	// Below a range node, a key alone with its byte gets a range of its own
	// by cutting the page around it.
	if at := p.byteAt(sp.depth); sp.par != nil && at != nil &&
		(i == 0 || at(i-1) != at(i)) && (i == n-1 || at(i+1) != at(i)) {
		b := at(i)
		var after byte
		if i < n-1 {
			after = at(i + 1)
		}
		var rs []rng
		if i > 0 {
			var left *pageHead
			left, p = p.cut(i)
			rs = append(rs, rng{0, pageHdr(left)})
		}
		rs = append(rs, rng{b, leafHdr(l)})
		if i < n-1 {
			_, right := p.cut(1)
			rs = append(rs, rng{after, pageHdr(right)})
		}
		*sp.par = rSplice(*sp.par, sp.pi, rs)
		return
	}
	items := pageItems(p)
	items[sp.i] = item{key: l.key(), leaf: l}
	t.replace(sp.at, sp.par, sp.pi, items, sp.depth)
}

// splitLeaf handles an insert that reaches leaf l: either it is the key's
// leaf, or both keys go below a new node holding their common path, a range
// node in a tree with pages unless l is below an inner node.
func (t *Tree) splitLeaf(loc **header, l *leafHead, key []byte, depth int, v uint64, nl newLeafFunc, inner bool) spot {
	lk := l.key()
	if bytes.Equal(lk, key) {
		return spot{leaf: l}
	}
	p := swar.Lcp(lk[depth:], key[depth:])
	if t.small && !inner {
		d := depth + p
		return t.fork(loc, leafHdr(l), d == len(lk), key, depth, d, v, nl)
	}
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(key[depth:depth+p], p)
	h := attach(&nn.header, lk, depth+p, l)
	return t.attachNew(loc, h, key, depth+p, nl)
}

// splitPrefix handles an insert whose key leaves n's compressed path pk after
// mis bytes: a new node takes the common part, with n and the new key below,
// a range node if n is one, else an inner node.
func (t *Tree) splitPrefix(loc **header, n *header, pk []byte, mis int, key []byte, depth int, v uint64, nl newLeafFunc) spot {
	if n.kind == kR {
		// Below a range node, n keeps the byte it branches on in its path.
		n.setPrefix(pk[mis:], len(pk)-mis)
		return t.fork(loc, n, false, key, depth, depth+mis, v, nl)
	}
	nn := &node4{}
	nn.kind = kN4
	nn.setPrefix(pk[:mis], mis)
	old := pk[mis]
	n.setPrefix(pk[mis+1:], len(pk)-mis-1) // pk is a copy or a leaf key: safe to read while n changes
	h := addChild(&nn.header, old, n)
	return t.attachNew(loc, h, key, depth+mis, nl)
}

// fork puts a range node with the path key[depth:d] at *loc, in a tree with
// pages, that holds old and the new key: old is a leaf whose key ends at d
// (oend), or its keys continue with one byte at d, as the new key does
// unless it ends there.
func (t *Tree) fork(loc **header, old *header, oend bool, key []byte, depth, d int, v uint64, nl newLeafFunc) spot {
	var h header
	h.setPrefix(key[depth:d], d-depth)
	if d == len(key) {
		h.term = nl(key)
		*loc = makeR(h, []rng{{0, old}})
		t.size++
		return spot{leaf: h.term, created: true}
	}
	nc := t.newChild(key, v, nl)
	rs := []rng{{0, nc}}
	if oend {
		h.term = asLeaf(old)
	} else if ob, _ := firstByte(old, d); key[d] < ob {
		rs = append(rs, rng{ob, old})
	} else {
		rs = []rng{{0, old}, {key[d], nc}}
	}
	*loc = makeR(h, rs)
	return t.created(nc)
}

// attachNew adds the new key below the inner node h at key depth d, as h's
// term leaf if it ends there, else as a new leaf, and stores h or its grown
// replacement in *loc.
func (t *Tree) attachNew(loc **header, h *header, key []byte, d int, nl newLeafFunc) spot {
	if d == len(key) {
		h.term = nl(key)
		*loc = h
		t.size++
		return spot{leaf: h.term, created: true}
	}
	c := leafHdr(nl(key))
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
