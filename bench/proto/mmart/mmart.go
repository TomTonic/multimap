// Package mmart is the multimap prototype on top of the pointer-based ART
// (package ptrart): every leaf carries the key's value set inline (vset.Set),
// so reading the values of a key costs no pointer chase beyond the leaf.
//
// The tree code is not generic: it works on leafHead, the key part every leaf
// starts with, so no generic dictionary calls sit on the traversal path. Only
// Map[T], which creates leaves and reads their values, is generic.
package mmart

import (
	"bytes"
	"iter"
	"math/bits"
	"unsafe"

	"github.com/TomTonic/multimap/bench/proto/swar"
	"github.com/TomTonic/multimap/bench/proto/vset"
)

type kind uint8

const (
	kLeaf kind = iota + 1
	kN4
	kN11
	kN57
	kN256
)

const inlineKey = 16

// header is the common prefix of all inner nodes (24 B).
type header struct {
	kind   kind
	_      uint8
	count  uint16
	plen   uint16
	_      uint16
	prefix [8]byte
	term   *leafHead // leaf for the key that ends exactly at this node
}

// leafHead is the key part every leaf starts with (40 B): the full key,
// inline up to 16 bytes.
type leafHead struct {
	kind   kind
	_      [3]byte
	klen   uint32
	inline [inlineKey]byte
	ext    string
}

// leaf is a leafHead followed by the key's values; 80 B for T = uint64.
type leaf[T comparable] struct {
	leafHead
	vals vset.Set[T]
}

// newLeafFunc creates a leaf for key and returns its head; Map[T] supplies it
// so the tree code does not need to know T.
type newLeafFunc func(key []byte) *leafHead

type node4 struct { // 64 B
	header
	keys  [8]byte
	child [4]*header
}

type node11 struct { // 128 B
	header
	keys  [16]byte
	child [11]*header
}

type node57 struct { // 512 B
	header
	bitmap [4]uint64
	child  [57]*header
}

type node256 struct { // 2072 B -> size class 2304 (malloc header, see README)
	header
	child [256]*header
}

func asLeaf(h *header) *leafHead  { return (*leafHead)(unsafe.Pointer(h)) }
func asN4(h *header) *node4       { return (*node4)(unsafe.Pointer(h)) }
func asN11(h *header) *node11     { return (*node11)(unsafe.Pointer(h)) }
func asN57(h *header) *node57     { return (*node57)(unsafe.Pointer(h)) }
func asN256(h *header) *node256   { return (*node256)(unsafe.Pointer(h)) }
func leafHdr(l *leafHead) *header { return (*header)(unsafe.Pointer(l)) }
func (l *leafHead) key() []byte {
	if l.klen <= inlineKey {
		return l.inline[:l.klen]
	}
	return unsafe.Slice(unsafe.StringData(l.ext), len(l.ext))
}

func (l *leafHead) init(key []byte) {
	l.kind, l.klen = kLeaf, uint32(len(key))
	if len(key) <= inlineKey {
		copy(l.inline[:], key)
	} else {
		l.ext = string(key) // the multimap owns (clones) its keys
	}
}

// Tree is the untyped ART; Map[T] wraps it.
type Tree struct {
	root *header
	size int
}

// Len returns the number of keys.
func (t *Tree) Len() int { return t.size }

// find returns the leaf of key, or nil.
func (t *Tree) find(key []byte) *leafHead {
	n := t.root
	depth := 0
	for n != nil {
		if n.kind == kLeaf {
			l := asLeaf(n)
			if bytes.Equal(l.key(), key) {
				return l
			}
			return nil
		}
		if n.plen != 0 {
			if !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				return nil
			}
			depth += int(n.plen)
		}
		if depth == len(key) {
			if n.term == nil {
				return nil
			}
			n = leafHdr(n.term)
			continue
		}
		b := key[depth]
		depth++
		switch n.kind {
		case kN4:
			x := asN4(n)
			i := swar.Index8(swar.Word(x.keys[:]), b)
			if i >= int(x.count) {
				return nil
			}
			n = x.child[i&3]
		case kN11:
			x := asN11(n)
			i := swar.Index8(swar.Word(x.keys[0:8]), b)
			if i == 8 {
				i = 8 + swar.Index8(swar.Word(x.keys[8:16]), b)
			}
			if i >= int(x.count) || i >= 11 {
				return nil
			}
			n = x.child[i]
		case kN57:
			x := asN57(n)
			if !swar.Has(&x.bitmap, b) {
				return nil
			}
			n = x.child[swar.Rank(&x.bitmap, b)]
		default:
			n = asN256(n).child[b]
		}
	}
	return nil
}

// upsert returns the leaf of key, creating it with nl when it is missing.
func (t *Tree) upsert(key []byte, nl newLeafFunc) *leafHead {
	return t.insert(&t.root, key, 0, nl)
}

func (t *Tree) insert(loc **header, key []byte, depth int, nl newLeafFunc) *leafHead {
	n := *loc
	if n == nil {
		l := nl(key)
		*loc = leafHdr(l)
		t.size++
		return l
	}
	if n.kind == kLeaf {
		l := asLeaf(n)
		lk := l.key()
		if bytes.Equal(lk, key) {
			return l
		}
		p := swar.Lcp(lk[depth:], key[depth:])
		nn := &node4{}
		nn.kind = kN4
		nn.setPrefix(key[depth:depth+p], p)
		d := depth + p
		nl2 := nl(key)
		h := &nn.header
		h = attach(h, lk, d, l)
		h = attach(h, key, d, nl2)
		*loc = h
		t.size++
		return nl2
	}
	if n.plen > 0 {
		var buf [8]byte
		pk := t.fullPrefix(n, depth, &buf)
		mis := swar.Lcp(pk, key[depth:])
		if mis < int(n.plen) {
			nn := &node4{}
			nn.kind = kN4
			nn.setPrefix(pk[:mis], mis)
			old := pk[mis]
			n.setPrefix(pk[mis+1:], int(n.plen)-mis-1)
			nl2 := nl(key)
			h := addChild(&nn.header, old, n)
			h = attach(h, key, depth+mis, nl2)
			*loc = h
			t.size++
			return nl2
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
	if c := findLoc(n, b); c != nil {
		return t.insert(c, key, depth+1, nl)
	}
	l := nl(key)
	*loc = addChild(n, b, leafHdr(l))
	t.size++
	return l
}

// attach hangs leaf l (whose key is k) below h at depth d.
func attach(h *header, k []byte, d int, l *leafHead) *header {
	if d == len(k) {
		h.term = l
		return h
	}
	return addChild(h, k[d], leafHdr(l))
}

func (h *header) setPrefix(p []byte, plen int) {
	h.plen = uint16(plen)
	h.prefix = [8]byte{}
	copy(h.prefix[:], p[:min(plen, 8)])
}

// fullPrefix returns the complete compressed path of n, reading it from the
// minimum leaf when it is longer than the 8 bytes stored inline.
func (t *Tree) fullPrefix(n *header, depth int, buf *[8]byte) []byte {
	if n.plen <= 8 {
		*buf = n.prefix
		return buf[:n.plen]
	}
	return minLeaf(n).key()[depth : depth+int(n.plen)]
}

func minLeaf(n *header) *leafHead {
	for {
		if n.kind == kLeaf {
			return asLeaf(n)
		}
		if n.term != nil {
			return n.term
		}
		switch n.kind {
		case kN4:
			n = asN4(n).child[0]
		case kN11:
			n = asN11(n).child[0]
		case kN57:
			n = asN57(n).child[0]
		default:
			x := asN256(n)
			for _, c := range x.child {
				if c != nil {
					n = c
					break
				}
			}
		}
	}
}

func findLoc(n *header, b byte) **header {
	switch n.kind {
	case kN4:
		x := asN4(n)
		for i := 0; i < int(x.count); i++ {
			if x.keys[i] == b {
				return &x.child[i]
			}
		}
	case kN11:
		x := asN11(n)
		for i := 0; i < int(x.count); i++ {
			if x.keys[i] == b {
				return &x.child[i]
			}
		}
	case kN57:
		x := asN57(n)
		if swar.Has(&x.bitmap, b) {
			return &x.child[swar.Rank(&x.bitmap, b)]
		}
	default:
		x := asN256(n)
		if x.child[b] != nil {
			return &x.child[b]
		}
	}
	return nil
}

// addChild inserts child c under byte b, growing the node when it is full.
// It returns the (possibly new) node.
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
		y := &node57{header: x.header}
		y.kind = kN57
		copy(y.child[:], x.child[:])
		for i := 0; i < 11; i++ {
			swar.Set(&y.bitmap, x.keys[i])
		}
		return addChild(&y.header, b, c)
	case kN57:
		x := asN57(n)
		if x.count < 57 {
			r := swar.Rank(&x.bitmap, b)
			copy(x.child[r+1:int(x.count)+1], x.child[r:x.count])
			x.child[r] = c
			swar.Set(&x.bitmap, b)
			x.count++
			return n
		}
		y := &node256{header: x.header}
		y.kind = kN256
		i := 0
		for k := 0; k < 256; k++ {
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

// scanFrom calls fn for every leaf with key >= from in ascending order until
// fn returns false.
func (t *Tree) scanFrom(from []byte, fn func(*leafHead) bool) {
	t.scan(t.root, from, 0, true, fn)
}

func (t *Tree) scan(n *header, from []byte, depth int, bounded bool, fn func(*leafHead) bool) bool {
	if n == nil {
		return true
	}
	if n.kind == kLeaf {
		l := asLeaf(n)
		if bounded && bytes.Compare(l.key(), from) < 0 {
			return true
		}
		return fn(l)
	}
	if bounded {
		if n.plen > 0 {
			var buf [8]byte
			pk := t.fullPrefix(n, depth, &buf)
			rest := from[depth:]
			m := swar.Lcp(pk, rest)
			if m < len(pk) {
				if m < len(rest) && pk[m] < rest[m] {
					return true // whole subtree < from
				}
				bounded = false // whole subtree > from
			}
			depth += int(n.plen)
		}
		if bounded && depth == len(from) {
			bounded = false
		}
	}
	var b byte
	if bounded {
		b = from[depth] // term key is a proper prefix of from -> smaller, skip it
	} else if n.term != nil && !fn(n.term) {
		return false
	}
	return t.scanChildren(n, b, bounded, from, depth, fn)
}

// scanChildren visits the children with byte >= b in order; the child for b
// itself stays bounded, all later ones are unbounded.
func (t *Tree) scanChildren(n *header, b byte, bounded bool, from []byte, depth int, fn func(*leafHead) bool) bool {
	var keys []byte
	var child []*header
	switch n.kind {
	case kN4:
		x := asN4(n)
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN11:
		x := asN11(n)
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN57:
		x := asN57(n)
		start := swar.Rank(&x.bitmap, b)
		for i := start; i < int(x.count); i++ {
			exact := bounded && i == start && swar.Has(&x.bitmap, b)
			if !t.scanOne(x.child[i], exact, from, depth, fn) {
				return false
			}
		}
		return true
	default:
		x := asN256(n)
		for k := int(b); k < 256; k++ {
			if x.child[k] != nil && !t.scanOne(x.child[k], bounded && k == int(b), from, depth, fn) {
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

func (t *Tree) scanOne(c *header, exact bool, from []byte, depth int, fn func(*leafHead) bool) bool {
	if exact {
		return t.scan(c, from, depth+1, true, fn)
	}
	return t.scan(c, nil, 0, false, fn)
}

// scanRange calls fn for every leaf with from <= key <= to in ascending order
// until fn returns false, and returns false once it has stopped.
//
// lo and hi say whether n lies on the path of from and of to. Both bounds are
// checked structurally: a node's prefix and child bytes are compared with the
// bound only while on its path. A subtree strictly inside the range is visited
// without any key comparison, so its leaves' keys (which live outside the leaf
// for keys over 16 bytes) are never read.
func (t *Tree) scanRange(n *header, from, to []byte, depth int, lo, hi bool, fn func(*leafHead) bool) bool {
	if n == nil {
		return true
	}
	if n.kind == kLeaf {
		l := asLeaf(n)
		if lo && bytes.Compare(l.key(), from) < 0 {
			return true
		}
		if hi && bytes.Compare(l.key(), to) > 0 {
			return false
		}
		return fn(l)
	}
	if (lo || hi) && n.plen > 0 {
		var buf [8]byte
		pk := t.fullPrefix(n, depth, &buf)
		if lo {
			rest := from[depth:]
			if m := swar.Lcp(pk, rest); m < len(pk) {
				if m < len(rest) && pk[m] < rest[m] {
					return true // whole subtree < from
				}
				lo = false // whole subtree > from
			}
		}
		if hi {
			rest := to[depth:]
			if m := swar.Lcp(pk, rest); m < len(pk) {
				if m == len(rest) || pk[m] > rest[m] {
					return false // whole subtree > to: done
				}
				hi = false // whole subtree < to
			}
		}
	}
	depth += int(n.plen)
	if lo && depth == len(from) {
		lo = false
	}
	// The term leaf's key is the path itself: below from while still on its
	// path (skipped), equal to to when to ends here (last key in range).
	if !lo && n.term != nil && !fn(n.term) {
		return false
	}
	if hi && depth == len(to) {
		return false
	}
	var loB, hiB byte = 0, 255
	if lo {
		loB = from[depth]
	}
	if hi {
		hiB = to[depth]
	}
	if !t.rangeChildren(n, from, to, depth, lo, hi, loB, hiB, fn) {
		return false
	}
	return !hi // on to's path, everything after the hiB child is > to
}

// rangeChildren visits the children with byte in [loB, hiB] in order. Only
// the child for loB stays on from's path, only the one for hiB on to's.
func (t *Tree) rangeChildren(n *header, from, to []byte, depth int, lo, hi bool, loB, hiB byte, fn func(*leafHead) bool) bool {
	visit := func(c *header, kb byte) bool {
		return t.scanRange(c, from, to, depth+1, lo && kb == loB, hi && kb == hiB, fn)
	}
	var keys []byte
	var child []*header
	switch n.kind {
	case kN4:
		x := asN4(n)
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN11:
		x := asN11(n)
		keys, child = x.keys[:x.count], x.child[:x.count]
	case kN57:
		x := asN57(n)
		i := swar.Rank(&x.bitmap, loB)
		for w := int(loB >> 6); w < 4; w++ {
			set := x.bitmap[w]
			if w == int(loB>>6) {
				set &= ^uint64(0) << (loB & 63)
			}
			for ; set != 0; set &= set - 1 {
				kb := byte(w<<6 + bits.TrailingZeros64(set))
				if kb > hiB {
					return true
				}
				if !visit(x.child[i], kb) {
					return false
				}
				i++
			}
		}
		return true
	default:
		x := asN256(n)
		for k := int(loB); k <= int(hiB); k++ {
			if x.child[k] != nil && !visit(x.child[k], byte(k)) {
				return false
			}
		}
		return true
	}
	for i, kb := range keys {
		if kb < loB {
			continue
		}
		if kb > hiB {
			break
		}
		if !visit(child[i], kb) {
			return false
		}
	}
	return true
}

// Map is a multimap from byte-string keys to sets of T.
type Map[T comparable] struct {
	t Tree
}

// newLeaf allocates a leaf[T] for key. It captures nothing, so passing it as
// a newLeafFunc allocates no closure.
func newLeaf[T comparable](key []byte) *leafHead {
	l := &leaf[T]{}
	l.init(key)
	return &l.leafHead
}

func vals[T comparable](l *leafHead) *vset.Set[T] {
	return &(*leaf[T])(unsafe.Pointer(l)).vals
}

// Len returns the number of keys.
func (m *Map[T]) Len() int { return m.t.Len() }

// AddValue adds v to the values of key, creating the key if needed.
func (m *Map[T]) AddValue(key []byte, v T) {
	vals[T](m.t.upsert(key, newLeaf[T])).Add(v)
}

// RemoveValue removes v from the values of key. The key stays, possibly with
// no values (deleting keys is not part of this prototype).
func (m *Map[T]) RemoveValue(key []byte, v T) {
	if l := m.t.find(key); l != nil {
		vals[T](l).Remove(v)
	}
}

// ValuesFor iterates over the values of key.
func (m *Map[T]) ValuesFor(key []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		if l := m.t.find(key); l != nil {
			vals[T](l).Each(yield)
		}
	}
}

// ValuesBetween iterates over the values of all keys in [from, to], in key
// order.
func (m *Map[T]) ValuesBetween(from, to []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		m.t.scanRange(m.t.root, from, to, 0, true, true, func(l *leafHead) bool {
			return vals[T](l).Each(yield)
		})
	}
}

// ValuesBetweenLinear is ValuesBetween with the upper bound checked by
// comparing every key, as the first prototype did. It stays only so the
// benchmark can measure what the structural check is worth.
func (m *Map[T]) ValuesBetweenLinear(from, to []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		m.t.scanFrom(from, func(l *leafHead) bool {
			if bytes.Compare(l.key(), to) > 0 {
				return false
			}
			return vals[T](l).Each(yield)
		})
	}
}

// Keys iterates over all keys in ascending order.
func (m *Map[T]) Keys() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		m.t.scanFrom(nil, func(l *leafHead) bool { return yield(l.key()) })
	}
}
