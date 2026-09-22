// Package arenaflat is the arena ART prototype with one contiguous slice per
// node kind instead of a chunk table, so a level of the tree costs exactly one
// dependent load (the node), as with real pointers.
//
// Growing a slice moves it, which would invalidate pointers held during an
// insert. Put therefore reserves room for the worst case of one insert (one
// leaf, two nodes) before it starts, so nothing moves while it runs.
//
// Everything else is identical to package arenaart.
package arenaflat

import (
	"bytes"

	"github.com/TomTonic/multimap/bench/proto/swar"
)

type ref uint32

const (
	kindShift = 29
	idxMask   = 1<<kindShift - 1
)

const (
	kLeaf = iota + 1
	kN8
	kN22
	kN52
	kN256
)

func mk(kind, idx uint32) ref { return ref(kind<<kindShift | idx) }
func (r ref) kind() uint32    { return uint32(r) >> kindShift }
func (r ref) idx() uint32     { return uint32(r) & idxMask }

const inlineKey = 16

// header is the common part of all inner nodes (16 B).
type header struct {
	count  uint16
	plen   uint16
	term   ref // leaf for the key that ends exactly at this node
	prefix [8]byte
}

type leaf struct { // 32 B
	val    uint32
	klen   uint32
	off    uint32 // offset into Tree.keys when klen > inlineKey
	_      uint32
	inline [inlineKey]byte
}

type node8 struct { // 64 B
	keys [8]byte
	header
	child [8]ref
	_     [8]byte
}

type node22 struct { // 128 B
	keys [24]byte
	header
	child [22]ref
}

type node52 struct { // 256 B
	bitmap [4]uint64
	header
	child [52]ref
}

type node256 struct { // 1088 B = 17 x 64 B, keeps every slot line-aligned
	header
	child [256]ref
	_     [48]byte
}

// slab stores T in one contiguous slice. Index 0 is reserved as nil.
type slab[T any] struct {
	s    []T
	free []uint32
}

func (s *slab[T]) at(i uint32) *T { return &s.s[i] }

// reserve makes sure k more allocations will not move the slice.
func (s *slab[T]) reserve(k int) {
	if len(s.s) == 0 {
		s.s = make([]T, 1, 1024) // slot 0 = nil
	}
	if cap(s.s)-len(s.s) >= k {
		return
	}
	ns := make([]T, len(s.s), 2*cap(s.s))
	copy(ns, s.s)
	s.s = ns
}

func (s *slab[T]) alloc() (uint32, *T) {
	if n := len(s.free); n > 0 {
		i := s.free[n-1]
		s.free = s.free[:n-1]
		p := &s.s[i]
		var zero T
		*p = zero
		return i, p
	}
	i := len(s.s)
	s.s = s.s[:i+1] // never reallocates: reserve ran first
	return uint32(i), &s.s[i]
}

func (s *slab[T]) release(i uint32) { s.free = append(s.free, i) }

// Tree is a map from byte-string keys to uint32 values.
type Tree struct {
	root   ref
	size   int
	leaves slab[leaf]
	n8     slab[node8]
	n22    slab[node22]
	n52    slab[node52]
	n256   slab[node256]
	keys   []byte // key bytes of leaves with klen > inlineKey
}

// Len returns the number of keys.
func (t *Tree) Len() int { return t.size }

func (t *Tree) leafKey(l *leaf) []byte {
	if l.klen <= inlineKey {
		return l.inline[:l.klen]
	}
	return t.keys[l.off : l.off+l.klen]
}

func (t *Tree) newLeaf(key []byte, val uint32) ref {
	i, l := t.leaves.alloc()
	l.val, l.klen = val, uint32(len(key))
	if len(key) <= inlineKey {
		copy(l.inline[:], key)
	} else {
		l.off = uint32(len(t.keys))
		t.keys = append(t.keys, key...)
	}
	return mk(kLeaf, i)
}

// hdr returns the header of inner node r.
func (t *Tree) hdr(r ref) *header {
	switch r.kind() {
	case kN8:
		return &t.n8.at(r.idx()).header
	case kN22:
		return &t.n22.at(r.idx()).header
	case kN52:
		return &t.n52.at(r.idx()).header
	default:
		return &t.n256.at(r.idx()).header
	}
}

// Get returns the value stored for key.
func (t *Tree) Get(key []byte) (uint32, bool) {
	lv, c8, c22, c52, c256 := t.leaves.s, t.n8.s, t.n22.s, t.n52.s, t.n256.s
	r := t.root
	depth := 0
	for {
		i := r.idx()
		switch r.kind() {
		case kLeaf:
			l := &lv[i]
			if bytes.Equal(t.leafKey(l), key) {
				return l.val, true
			}
			return 0, false
		case kN8:
			x := &c8[i]
			if x.plen != 0 {
				if !swar.MatchPrefix(&x.prefix, int(x.plen), key, depth) {
					return 0, false
				}
				depth += int(x.plen)
			}
			if depth == len(key) {
				r = x.term
				continue
			}
			j := swar.Index8(swar.Word(x.keys[:]), key[depth])
			if j >= int(x.count) {
				return 0, false
			}
			r = x.child[j&7]
		case kN22:
			x := &c22[i]
			if x.plen != 0 {
				if !swar.MatchPrefix(&x.prefix, int(x.plen), key, depth) {
					return 0, false
				}
				depth += int(x.plen)
			}
			if depth == len(key) {
				r = x.term
				continue
			}
			b := key[depth]
			j := swar.Index8(swar.Word(x.keys[0:8]), b)
			if j == 8 {
				j = 8 + swar.Index8(swar.Word(x.keys[8:16]), b)
				if j == 16 {
					j = 16 + swar.Index8(swar.Word(x.keys[16:24]), b)
				}
			}
			if j >= int(x.count) || j >= 22 {
				return 0, false
			}
			r = x.child[j]
		case kN52:
			x := &c52[i]
			if x.plen != 0 {
				if !swar.MatchPrefix(&x.prefix, int(x.plen), key, depth) {
					return 0, false
				}
				depth += int(x.plen)
			}
			if depth == len(key) {
				r = x.term
				continue
			}
			b := key[depth]
			if !swar.Has(&x.bitmap, b) {
				return 0, false
			}
			r = x.child[swar.Rank(&x.bitmap, b)]
		case kN256:
			x := &c256[i]
			if x.plen != 0 {
				if !swar.MatchPrefix(&x.prefix, int(x.plen), key, depth) {
					return 0, false
				}
				depth += int(x.plen)
			}
			if depth == len(key) {
				r = x.term
				continue
			}
			r = x.child[key[depth]]
		default: // nil ref
			return 0, false
		}
		depth++
	}
}

// Put inserts or overwrites key.
func (t *Tree) Put(key []byte, val uint32) {
	t.leaves.reserve(2)
	t.n8.reserve(2)
	t.n22.reserve(1)
	t.n52.reserve(1)
	t.n256.reserve(1)
	t.insert(&t.root, key, 0, val)
}

func (t *Tree) insert(loc *ref, key []byte, depth int, val uint32) {
	r := *loc
	if r == 0 {
		*loc = t.newLeaf(key, val)
		t.size++
		return
	}
	if r.kind() == kLeaf {
		l := t.leaves.at(r.idx())
		lk := t.leafKey(l)
		if bytes.Equal(lk, key) {
			l.val = val
			return
		}
		p := swar.Lcp(lk[depth:], key[depth:])
		ni, nn := t.n8.alloc()
		nn.setPrefix(key[depth:depth+p], p)
		d := depth + p
		nr := mk(kN8, ni)
		nr = t.attach(nr, lk, d, r)
		nr = t.attach(nr, key, d, t.newLeaf(key, val))
		*loc = nr
		t.size++
		return
	}
	h := t.hdr(r)
	if h.plen > 0 {
		var buf [8]byte
		pk := t.fullPrefix(r, h, depth, &buf)
		mis := swar.Lcp(pk, key[depth:])
		if mis < int(h.plen) {
			ni, nn := t.n8.alloc()
			nn.setPrefix(pk[:mis], mis)
			old := pk[mis]
			h.setPrefix(pk[mis+1:], int(h.plen)-mis-1)
			nr := t.addChild(mk(kN8, ni), old, r)
			nr = t.attach(nr, key, depth+mis, t.newLeaf(key, val))
			*loc = nr
			t.size++
			return
		}
		depth += int(h.plen)
	}
	if depth == len(key) {
		if h.term != 0 {
			t.leaves.at(h.term.idx()).val = val
		} else {
			h.term = t.newLeaf(key, val)
			t.size++
		}
		return
	}
	b := key[depth]
	if c := t.findLoc(r, b); c != nil {
		t.insert(c, key, depth+1, val)
		return
	}
	*loc = t.addChild(r, b, t.newLeaf(key, val))
	t.size++
}

func (t *Tree) attach(n ref, k []byte, d int, l ref) ref {
	if d == len(k) {
		t.hdr(n).term = l
		return n
	}
	return t.addChild(n, k[d], l)
}

func (h *header) setPrefix(p []byte, plen int) {
	h.plen = uint16(plen)
	h.prefix = [8]byte{}
	copy(h.prefix[:], p[:min(plen, 8)])
}

func (t *Tree) fullPrefix(r ref, h *header, depth int, buf *[8]byte) []byte {
	if h.plen <= 8 {
		*buf = h.prefix
		return buf[:h.plen]
	}
	return t.leafKey(t.leaves.at(t.minLeaf(r).idx()))[depth : depth+int(h.plen)]
}

func (t *Tree) minLeaf(r ref) ref {
	for r.kind() != kLeaf {
		h := t.hdr(r)
		if h.term != 0 {
			return h.term
		}
		switch r.kind() {
		case kN8:
			r = t.n8.at(r.idx()).child[0]
		case kN22:
			r = t.n22.at(r.idx()).child[0]
		case kN52:
			r = t.n52.at(r.idx()).child[0]
		default:
			x := t.n256.at(r.idx())
			for _, c := range x.child {
				if c != 0 {
					r = c
					break
				}
			}
		}
	}
	return r
}

func (t *Tree) findLoc(r ref, b byte) *ref {
	switch r.kind() {
	case kN8:
		x := t.n8.at(r.idx())
		for i := 0; i < int(x.count); i++ {
			if x.keys[i] == b {
				return &x.child[i]
			}
		}
	case kN22:
		x := t.n22.at(r.idx())
		for i := 0; i < int(x.count); i++ {
			if x.keys[i] == b {
				return &x.child[i]
			}
		}
	case kN52:
		x := t.n52.at(r.idx())
		if swar.Has(&x.bitmap, b) {
			return &x.child[swar.Rank(&x.bitmap, b)]
		}
	default:
		x := t.n256.at(r.idx())
		if x.child[b] != 0 {
			return &x.child[b]
		}
	}
	return nil
}

// addChild inserts child c under byte b, growing the node when it is full,
// and returns the (possibly new) ref of the node.
func (t *Tree) addChild(r ref, b byte, c ref) ref {
	switch r.kind() {
	case kN8:
		x := t.n8.at(r.idx())
		if x.count < 8 {
			insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
			x.count++
			return r
		}
		yi, y := t.n22.alloc()
		y.header = x.header
		copy(y.keys[:], x.keys[:])
		copy(y.child[:], x.child[:])
		t.n8.release(r.idx())
		return t.addChild(mk(kN22, yi), b, c)
	case kN22:
		x := t.n22.at(r.idx())
		if x.count < 22 {
			insertSorted(x.keys[:], x.child[:], int(x.count), b, c)
			x.count++
			return r
		}
		yi, y := t.n52.alloc()
		y.header = x.header
		copy(y.child[:], x.child[:])
		for i := 0; i < 22; i++ {
			swar.Set(&y.bitmap, x.keys[i])
		}
		t.n22.release(r.idx())
		return t.addChild(mk(kN52, yi), b, c)
	case kN52:
		x := t.n52.at(r.idx())
		if x.count < 52 {
			k := swar.Rank(&x.bitmap, b)
			copy(x.child[k+1:int(x.count)+1], x.child[k:x.count])
			x.child[k] = c
			swar.Set(&x.bitmap, b)
			x.count++
			return r
		}
		yi, y := t.n256.alloc()
		y.header = x.header
		i := 0
		for k := 0; k < 256; k++ {
			if swar.Has(&x.bitmap, byte(k)) {
				y.child[k] = x.child[i]
				i++
			}
		}
		t.n52.release(r.idx())
		return t.addChild(mk(kN256, yi), b, c)
	default:
		x := t.n256.at(r.idx())
		x.child[b] = c
		x.count++
		return r
	}
}

func insertSorted(keys []byte, child []ref, count int, b byte, c ref) {
	i := 0
	for i < count && keys[i] < b {
		i++
	}
	copy(keys[i+1:count+1], keys[i:count])
	copy(child[i+1:count+1], child[i:count])
	keys[i] = b
	child[i] = c
}

// Scan calls fn for every key >= from in ascending order until fn returns false.
func (t *Tree) Scan(from []byte, fn func(key []byte, val uint32) bool) {
	t.scan(t.root, from, 0, true, fn)
}

func (t *Tree) scan(r ref, from []byte, depth int, bounded bool, fn func([]byte, uint32) bool) bool {
	if r == 0 {
		return true
	}
	if r.kind() == kLeaf {
		l := t.leaves.at(r.idx())
		k := t.leafKey(l)
		if bounded && bytes.Compare(k, from) < 0 {
			return true
		}
		return fn(k, l.val)
	}
	h := t.hdr(r)
	if bounded {
		if h.plen > 0 {
			var buf [8]byte
			pk := t.fullPrefix(r, h, depth, &buf)
			rest := from[depth:]
			m := swar.Lcp(pk, rest)
			if m < len(pk) {
				if m < len(rest) && pk[m] < rest[m] {
					return true
				}
				bounded = false
			}
			depth += int(h.plen)
		}
		if bounded && depth == len(from) {
			bounded = false
		}
	}
	var b byte
	if bounded {
		b = from[depth]
	} else if h.term != 0 {
		l := t.leaves.at(h.term.idx())
		if !fn(t.leafKey(l), l.val) {
			return false
		}
	}
	return t.scanChildren(r, b, bounded, from, depth, fn)
}

func (t *Tree) scanChildren(r ref, b byte, bounded bool, from []byte, depth int, fn func([]byte, uint32) bool) bool {
	var keys []byte
	var child []ref
	switch r.kind() {
	case kN8:
		x := t.n8.at(r.idx())
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN22:
		x := t.n22.at(r.idx())
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN52:
		x := t.n52.at(r.idx())
		start := swar.Rank(&x.bitmap, b)
		for i := start; i < int(x.count); i++ {
			exact := bounded && i == start && swar.Has(&x.bitmap, b)
			if !t.scanOne(x.child[i], exact, from, depth, fn) {
				return false
			}
		}
		return true
	default:
		x := t.n256.at(r.idx())
		for k := int(b); k < 256; k++ {
			if x.child[k] != 0 && !t.scanOne(x.child[k], bounded && k == int(b), from, depth, fn) {
				return false
			}
		}
		return true
	}
	for i, k := range keys {
		if k < b {
			continue
		}
		if !t.scanOne(child[i], bounded && k == b, from, depth, fn) {
			return false
		}
	}
	return true
}

func (t *Tree) scanOne(c ref, exact bool, from []byte, depth int, fn func([]byte, uint32) bool) bool {
	if exact {
		return t.scan(c, from, depth+1, true, fn)
	}
	return t.scan(c, nil, 0, false, fn)
}
