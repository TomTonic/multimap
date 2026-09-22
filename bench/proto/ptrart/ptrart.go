// Package ptrart is the pointer-based ART prototype: nodes are ordinary Go heap
// objects linked by pointers, the node kind lives in the first byte of every
// node, and kinds are sized to Go allocator size classes (64/128/512 B).
//
// It implements exactly the same algorithm as package arenaart (sorted keys,
// SWAR search, bitmap+popcount for the large node, lazy expansion, 8-byte
// pessimistic + optimistic path compression, full key in the leaf). The two
// differ only in memory representation, which is what the benchmark compares.
package ptrart

import (
	"bytes"
	"unsafe"

	"github.com/TomTonic/multimap/bench/proto/swar"
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
	term   *leaf // value for the key that ends exactly at this node
}

// leaf holds the full key (inline up to 16 bytes) and the value (48 B).
type leaf struct {
	kind   kind
	_      [3]byte
	val    uint32
	klen   uint32
	_      uint32
	inline [inlineKey]byte
	ext    string
}

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

func asLeaf(h *header) *leaf    { return (*leaf)(unsafe.Pointer(h)) }
func asN4(h *header) *node4     { return (*node4)(unsafe.Pointer(h)) }
func asN11(h *header) *node11   { return (*node11)(unsafe.Pointer(h)) }
func asN57(h *header) *node57   { return (*node57)(unsafe.Pointer(h)) }
func asN256(h *header) *node256 { return (*node256)(unsafe.Pointer(h)) }
func leafHdr(l *leaf) *header   { return (*header)(unsafe.Pointer(l)) }
func (l *leaf) key() []byte {
	if l.klen <= inlineKey {
		return l.inline[:l.klen]
	}
	return unsafe.Slice(unsafe.StringData(l.ext), len(l.ext))
}

func newLeaf(key []byte, val uint32) *leaf {
	l := &leaf{kind: kLeaf, val: val, klen: uint32(len(key))}
	if len(key) <= inlineKey {
		copy(l.inline[:], key)
	} else {
		l.ext = string(key)
	}
	return l
}

// Tree is a map from byte-string keys to uint32 values.
type Tree struct {
	root *header
	size int
}

// Len returns the number of keys.
func (t *Tree) Len() int { return t.size }

// Get returns the value stored for key.
func (t *Tree) Get(key []byte) (uint32, bool) {
	n := t.root
	depth := 0
	for n != nil {
		if n.kind == kLeaf {
			l := asLeaf(n)
			if bytes.Equal(l.key(), key) {
				return l.val, true
			}
			return 0, false
		}
		if n.plen != 0 {
			if !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				return 0, false
			}
			depth += int(n.plen)
		}
		if depth == len(key) {
			if n.term == nil {
				return 0, false
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
				return 0, false
			}
			n = x.child[i&3]
		case kN11:
			x := asN11(n)
			i := swar.Index8(swar.Word(x.keys[0:8]), b)
			if i == 8 {
				i = 8 + swar.Index8(swar.Word(x.keys[8:16]), b)
			}
			if i >= int(x.count) || i >= 11 {
				return 0, false
			}
			n = x.child[i]
		case kN57:
			x := asN57(n)
			if !swar.Has(&x.bitmap, b) {
				return 0, false
			}
			n = x.child[swar.Rank(&x.bitmap, b)]
		default:
			n = asN256(n).child[b]
		}
	}
	return 0, false
}

// Put inserts or overwrites key.
func (t *Tree) Put(key []byte, val uint32) { t.insert(&t.root, key, 0, val) }

func (t *Tree) insert(loc **header, key []byte, depth int, val uint32) {
	n := *loc
	if n == nil {
		*loc = leafHdr(newLeaf(key, val))
		t.size++
		return
	}
	if n.kind == kLeaf {
		l := asLeaf(n)
		lk := l.key()
		if bytes.Equal(lk, key) {
			l.val = val
			return
		}
		p := swar.Lcp(lk[depth:], key[depth:])
		nn := &node4{}
		nn.kind = kN4
		nn.setPrefix(key[depth:depth+p], p)
		d := depth + p
		h := &nn.header
		h = attach(h, lk, d, l)
		h = attach(h, key, d, newLeaf(key, val))
		*loc = h
		t.size++
		return
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
			h := addChild(&nn.header, old, n)
			h = attach(h, key, depth+mis, newLeaf(key, val))
			*loc = h
			t.size++
			return
		}
		depth += int(n.plen)
	}
	if depth == len(key) {
		if n.term != nil {
			n.term.val = val
		} else {
			n.term = newLeaf(key, val)
			t.size++
		}
		return
	}
	b := key[depth]
	if c := findLoc(n, b); c != nil {
		t.insert(c, key, depth+1, val)
		return
	}
	*loc = addChild(n, b, leafHdr(newLeaf(key, val)))
	t.size++
}

// attach hangs leaf l (whose key is k) below h at depth d.
func attach(h *header, k []byte, d int, l *leaf) *header {
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

func minLeaf(n *header) *leaf {
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

// Scan calls fn for every key >= from in ascending order until fn returns false.
func (t *Tree) Scan(from []byte, fn func(key []byte, val uint32) bool) {
	t.scan(t.root, from, 0, true, fn)
}

func (t *Tree) scan(n *header, from []byte, depth int, bounded bool, fn func([]byte, uint32) bool) bool {
	if n == nil {
		return true
	}
	if n.kind == kLeaf {
		l := asLeaf(n)
		if bounded && bytes.Compare(l.key(), from) < 0 {
			return true
		}
		return fn(l.key(), l.val)
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
	} else if n.term != nil && !fn(n.term.key(), n.term.val) {
		return false
	}
	return t.scanChildren(n, b, bounded, from, depth, fn)
}

// scanChildren visits the children with byte >= b in order; the child for b
// itself stays bounded, all later ones are unbounded.
func (t *Tree) scanChildren(n *header, b byte, bounded bool, from []byte, depth int, fn func([]byte, uint32) bool) bool {
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

func (t *Tree) scanOne(c *header, exact bool, from []byte, depth int, fn func([]byte, uint32) bool) bool {
	if exact {
		return t.scan(c, from, depth+1, true, fn)
	}
	return t.scan(c, nil, 0, false, fn)
}
