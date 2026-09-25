// Package art is the adaptive radix tree behind multimap.Ordered.
//
// Design (see bench/README.md for the measurements behind each choice):
//
//   - Inner nodes hold 4, 11, 25, 57 or 256 children in 64, 128, 256, 512 and
//     2072 bytes. Up to 512 bytes these are Go size classes, so every node is
//     cache-line aligned; above 512 bytes Go prepends a malloc header, which
//     only the rare 256-way node pays.
//   - Children are kept in key-byte order. The 4- and 11-way nodes store their
//     bytes in a sorted array searched with SWAR, eight bytes per step; the
//     25- and 57-way nodes find a child by the rank of its byte in a 256-bit
//     bitmap (branch-free); the 256-way node indexes directly.
//   - Every inner node stores up to 8 bytes of its compressed path; longer
//     paths are skipped optimistically and verified against the full key,
//     which every leaf holds (lazy expansion: a subtree with one key is just
//     its leaf).
//   - A leaf holds its key inline, in the smallest of four size classes (16,
//     32, 48 or 64 bytes) that fits, and its values right after it; only a
//     longer key is a separate string. Comparing a key at the leaf therefore
//     costs no second pointer chase, and a leaf with uint64 values fills a Go
//     size class (64 B for integer keys, one cache line).
//   - A key that ends at an inner node (a prefix of other keys) is stored as
//     that node's term leaf.
//
// The tree code is not generic. It works on leafHead, the key part every leaf
// starts with, so no generic dictionary calls sit on the traversal path; only
// Map[T], which creates leaves and reaches their values, is generic.
//
// A Tree is not safe for concurrent use when any goroutine writes; concurrent
// readers alone are safe, because reads never modify the tree.
package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

type kind uint8

const (
	kLeaf kind = iota + 1
	kN4
	kN11
	kN25
	kN57
	kN256
)

// Shrink thresholds: a node turns into the next smaller kind once it holds
// this many children or fewer. They lie below the next smaller capacity, so a
// node that has just grown does not shrink back after one removal.
const (
	shrink11  = 3
	shrink25  = 9
	shrink57  = 22
	shrink256 = 48
)

// maxInline is the longest key a leaf holds inline, in an array of 16, 32, 48
// or 64 bytes, whichever is the smallest that fits. A longer key is held as a
// string, which costs a separate allocation and a pointer chase on every
// comparison.
const maxInline = 64

// header is the common start of all inner nodes (24 B).
type header struct {
	kind   kind
	_      uint8
	count  uint16    // number of children
	plen   uint32    // length of the compressed path
	prefix [8]byte   // first min(plen, 8) bytes of the compressed path
	term   *leafHead // leaf of the key that ends exactly at this node
}

// leafHead is the start of every leaf (8 B). The key follows it at keyOff,
// inline or as a string (see maxInline); the key's values follow the key, at
// valsOff, which depends on the key's size class and on T.
type leafHead struct {
	kind    kind
	_       uint8
	valsOff uint16 // offset of the value set from the start of the leaf
	klen    uint32
}

// keyOff is the offset of the key in every leaf: right after the leafHead.
const keyOff = unsafe.Sizeof(leafHead{})

// keyArea is the storage of a leaf's key: an inline array of one of the
// size classes, or a string for keys longer than maxInline.
type keyArea interface {
	[16]byte | [32]byte | [48]byte | [64]byte | string
}

// leaf is a leafHead followed by its key and the values of its key. For
// T = uint64 it is 64, 80, 96 or 112 B with an inline key, all Go size
// classes, and 64 B plus the string with a longer key.
type leaf[T comparable, K keyArea] struct {
	leafHead
	k    K
	vals vset.Set[T]
}

// newLeafFunc creates a leaf for key and returns its head. Map[T] supplies it
// so that the tree code does not need to know T.
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

type node25 struct { // 256 B
	header
	bitmap [4]uint64
	child  [25]*header
}

type node57 struct { // 512 B
	header
	bitmap [4]uint64
	child  [57]*header
}

type node256 struct { // 2072 B
	header
	child [256]*header
}

// The casts below are valid because every node and leaf type starts with a
// kind byte at offset 0, and a pointer is only ever cast to the type its kind
// names. The garbage collector traces objects by their allocated type, not by
// the static type of the pointer.

func asLeaf(h *header) *leafHead  { return (*leafHead)(unsafe.Pointer(h)) }
func asN4(h *header) *node4       { return (*node4)(unsafe.Pointer(h)) }
func asN11(h *header) *node11     { return (*node11)(unsafe.Pointer(h)) }
func asN25(h *header) *node25     { return (*node25)(unsafe.Pointer(h)) }
func asN57(h *header) *node57     { return (*node57)(unsafe.Pointer(h)) }
func asN256(h *header) *node256   { return (*node256)(unsafe.Pointer(h)) }
func leafHdr(l *leafHead) *header { return (*header)(unsafe.Pointer(l)) }

// key returns the leaf's key. The slice aliases the leaf and must not be
// modified.
func (l *leafHead) key() []byte {
	p := unsafe.Add(unsafe.Pointer(l), keyOff)
	if l.klen <= maxInline {
		return unsafe.Slice((*byte)(p), l.klen)
	}
	s := *(*string)(p)
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// init marks a new leaf for a key of klen bytes whose values lie at valsOff.
func (l *leafHead) init(klen int, valsOff uintptr) {
	l.kind, l.klen, l.valsOff = kLeaf, uint32(klen), uint16(valsOff)
}

func (h *header) setPrefix(p []byte, plen int) {
	h.plen = uint32(plen)
	h.prefix = [8]byte{}
	copy(h.prefix[:], p[:min(plen, 8)])
}

// sorted returns the child bytes and children of a node that keeps them in
// sorted arrays: n must be a 4- or 11-way node.
func sorted(n *header) ([]byte, []*header) {
	if n.kind == kN4 {
		x := asN4(n)
		return x.keys[:x.count], x.child[:x.count]
	}
	x := asN11(n)
	return x.keys[:x.count], x.child[:x.count]
}

// bitmapOf returns the bitmap and the full child array of a 25- or 57-way
// node; the first count children are in use, in key-byte order.
func bitmapOf(n *header) (*[4]uint64, []*header) {
	if n.kind == kN25 {
		x := asN25(n)
		return &x.bitmap, x.child[:]
	}
	x := asN57(n)
	return &x.bitmap, x.child[:]
}
