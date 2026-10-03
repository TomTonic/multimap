// Package art is the adaptive radix tree behind multimap.Ordered.
//
// Design (see bench/README.md for the measurements behind each choice):
//
//   - Inner nodes have 5, 12, 26, 58 or 257 child slots in 64, 128, 256, 512
//     and 2080 bytes. Up to 512 bytes these are Go size classes, so every node
//     is cache-line aligned; above 512 bytes Go prepends a malloc header,
//     which only the rare 256-way node pays.
//   - Children are kept in key-byte order. The 5- and 12-way nodes store their
//     bytes in a sorted array searched with SWAR, eight bytes per step; the
//     26- and 58-way nodes find a child by the rank of its byte in a 256-bit
//     bitmap (branch-free); the 256-way node indexes directly.
//   - The node header is 16 bytes: kind, child count, path length and the
//     first 12 bytes of the compressed path. The rest of a longer path
//     follows the node in the same object (path.go), so every key byte on
//     the way down is checked in the nodes (pessimistic path compression).
//   - A leaf holds only the part of its key below the node it was created
//     under (lazy expansion: a subtree with one key is just its leaf), inline
//     right after a 6-byte head; only a remainder of more than 254 bytes
//     makes the leaf hold its whole key as a separate string. The shared
//     parts of the keys are stored once, in the nodes.
//   - For small pointer-free values (flat.go) the values follow the key in
//     the same object, which grows through Go size classes from 32 to 512
//     bytes as values arrive. Such a leaf holds no pointer, so the garbage
//     collector never scans it. Small values with a pointer, such as strings,
//     live in typed leaves (typed.go) the same way: the object is allocated
//     with its real type, one of eight value capacities up to 16 values and
//     one of eight key areas up to 58 bytes. Beyond that, and for other
//     values, a leaf holds a vset.Set after its key (set leaves).
//   - A key that ends at an inner node (a prefix of other keys) is that
//     node's term leaf. It takes the node's last child slot, which the byte
//     children reach only when there is no term: few keys are prefixes of
//     others, so no node pays a field for them.
//   - In a map whose values are small and pointer-free (at most 8 bytes), keys
//     with exactly one value need no leaf at all: they live in pages
//     (page.go), sorted arrays of up to 31 keys with their values that hold
//     whole keys of one length of at most 8 bytes, such as integers. Range
//     nodes (rnode.go) give the pages below them ranges of key bytes instead
//     of one child per byte, which keeps them full however many keys there
//     are, and let a range scan walk contiguous memory. A key
//     that gets a second value leaves its page for a leaf, and where such keys
//     crowd a range node, its subtree is rebuilt from inner nodes and leaves
//     (rebuild.go). Only keys in such a tree's pages and range nodes pay for
//     that; every other map has none of them.
//
// The tree code is not generic. It works on leafHead, the key part every leaf
// starts with, so no generic dictionary calls sit on the traversal path; only
// Map[T], which creates leaves and reaches their values, is generic.
//
// A Tree is not safe for concurrent use when any goroutine writes; concurrent
// readers alone are safe, because reads never modify the tree.
package artstr

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/lpage"
	"github.com/TomTonic/multimap/internal/swar"
	"github.com/TomTonic/multimap/internal/vset"
)

// kind tells the node or leaf type, and a flat leaf's size class: an object
// is a leaf iff its kind is at most kLastLeaf (see isLeaf).
type kind uint8

const (
	kSet      kind = iota + 1 // set leaf; kSet+c is a flat leaf of class c (see flatSizes), or in a map of values with a pointer a typed leaf of class c (see typedCaps)
	kLastLeaf kind = kSet + 9 // flat leaf of the largest class
	kPage     kind = kLastLeaf + 1
	// kPage to kLastPage are the first bytes of a page (see page.go).
	kLastPage kind = kPage + lpage.Kinds - 1
	kN5       kind = kLastPage + 1 + iota - 4
	kN12
	kN26
	kN58
	kN256
	kR8 // range nodes, see rnode.go
	kR24
	kR56
	kR256
)

// A descent ends at an object of a kind up to kLastPage: leaves and pages come
// first, so that one comparison detects them.

// kindMask maps a kind to an index of the tables below, which hold every kind.
const kindMask = 255

// Shrink thresholds: a node turns into the next smaller kind once it holds
// this many byte children or fewer. They lie below the next smaller capacity
// less the term's slot, so a node that has just grown does not shrink back
// after one removal.
const (
	shrink12  = 3
	shrink26  = 9
	shrink58  = 22
	shrink256 = 48
)

// maxInline is the longest key remainder a leaf holds inline: in a set leaf
// in an array of 16 to 256 bytes, whichever is the smallest that fits, in a
// flat leaf right before its values. A longer one makes the leaf hold its
// whole key as a string, which costs a separate allocation and a pointer
// chase on every comparison.
const maxInline = longKey - 1

// maxKeyLen is the longest key a leaf holds a remainder of; leafHead.kl must
// hold its length. A longer key is held whole, as a string.
const maxKeyLen = 1<<16 - 1

// header is the common start of all inner nodes (16 B).
type header struct {
	kind   kind
	count  uint8                // byte children; a 256-way node keeps its count in node256.total
	plen   uint16               // length of the compressed path, or longPath
	prefix [swar.PrefixLen]byte // first min(plen, 12) bytes of the compressed path; the rest is in the tail
}

// longPath in header.plen stands for a path of that many bytes or more, whose
// tail is a string: its length is then 12 plus the string's (see pathLen).
const longPath = 1<<16 - 1

// leafHead is the start of every leaf (6 B). The key remainder follows at
// keyOff; a whole key held as a string sits at strOff.
//
// A leaf holds its key from its base on, the depth it was created at: the
// bytes before are the path to it, stored in the nodes above. The leaf may
// move deeper later, when a node is split in above it, and still holds the
// bytes from its base, which are then also on its path; it moves up only
// with a new base (see rekeyFunc).
type leafHead struct {
	kind kind   // kSet, or kSet+c for a flat or typed leaf of class c
	klen uint8  // length of the inline key remainder, or longKey for a whole key held as a string
	n    uint16 // flat and typed leaves: number of values
	kl   uint16 // length of the whole key, if the remainder is inline
}

// isLeaf reports whether an object of kind k is a leaf.
func isLeaf(k kind) bool { return k <= kLastLeaf }

// isPage reports whether an object of kind k is a page.
func isPage(k kind) bool { return k-kPage <= kLastPage-kPage }

// isRange reports whether an object of kind k is a range node.
func isRange(k kind) bool { return k >= kR8 }

// cls returns the size class of a flat leaf (see flatSizes), or 0 for a set
// leaf.
func (l *leafHead) cls() uint8 { return uint8(l.kind - kSet) }

// longKey in leafHead.klen marks a whole key held as a string.
const longKey = 255

// keyOff is the offset of an inline key in every leaf: right after the
// leafHead. A string key sits at strOff, where a string is aligned.
const (
	keyOff = unsafe.Sizeof(leafHead{})
	strOff = 8
)

// keyArea is the storage of a set leaf's key: an inline array of one of the
// size classes for the remainder, or a string for the whole key.
type keyArea interface {
	[16]byte | [32]byte | [48]byte | [64]byte | [96]byte | [128]byte | [192]byte | [256]byte | string
}

// leaf is a set leaf: a leafHead followed by its key and the values of its
// key. For T = uint64 it is 64 B with a remainder of up to 16 bytes.
type leaf[T comparable, K keyArea] struct {
	leafHead
	k    K
	vals vset.Set[T]
}

// newLeafFunc creates a leaf for key with base base (see leafHead) and returns
// its head. Map[T] supplies it so that the tree code does not need to know T.
type newLeafFunc func(key []byte, base int) *leafHead

// rekeyFunc moves leaf l to depth, which is below its base: it returns a leaf
// with the same values that holds its key from depth on. The tree calls it when
// a node above l goes away and l takes its place. l's whole key is pre, then
// the byte b unless b is negative, then l's key from there on. Map[T] supplies
// it, and gets by without the whole key when the longer remainder still fits l.
type rekeyFunc func(l *leafHead, pre []byte, b int, depth int) *leafHead

// Every node kind but the 256-way one keeps its term leaf, if any, in its
// last child slot, which is free whenever there is a term (see termOf).

type node5 struct { // 64 B
	header
	keys  [8]byte
	child [5]*header
}

type node12 struct { // 128 B
	header
	keys  [16]byte
	child [12]*header
}

type node26 struct { // 256 B
	header
	bitmap [4]uint64
	child  [26]*header
}

type node58 struct { // 512 B
	header
	bitmap [4]uint64
	child  [58]*header
}

// node256 holds the child of byte b in child[b] and its term in child[256].
// header.count is fixed at 255, since the byte children can number 256.
type node256 struct { // 2080 B
	header
	total uint16 // byte children
	_     [6]byte
	child [257]*header
}

// The casts below are valid because every node and leaf type starts with a
// kind byte at offset 0, and a pointer is only ever cast to the type its kind
// names. The garbage collector traces objects by their allocated type, not by
// the static type of the pointer.

func asLeaf(h *header) *leafHead  { return (*leafHead)(unsafe.Pointer(h)) }
func asN5(h *header) *node5       { return (*node5)(unsafe.Pointer(h)) }
func asN12(h *header) *node12     { return (*node12)(unsafe.Pointer(h)) }
func asN26(h *header) *node26     { return (*node26)(unsafe.Pointer(h)) }
func asN58(h *header) *node58     { return (*node58)(unsafe.Pointer(h)) }
func asN256(h *header) *node256   { return (*node256)(unsafe.Pointer(h)) }
func leafHdr(l *leafHead) *header { return (*header)(unsafe.Pointer(l)) }

// stored returns the key bytes the leaf holds: its key from its base on, or
// its whole key. The slice aliases the leaf and must not be modified.
func (l *leafHead) stored() []byte {
	if l.klen != longKey {
		return unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), l.klen)
	}
	s := *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// keyLen returns the length of the leaf's whole key.
func (l *leafHead) keyLen() int {
	if l.klen != longKey {
		return int(l.kl)
	}
	return len(*(*string)(unsafe.Add(unsafe.Pointer(l), strOff)))
}

// base returns the depth the leaf holds its key from.
func (l *leafHead) base() int { return l.keyLen() - len(l.stored()) }

// from returns the leaf's key from depth on; depth must not be below its base.
func (l *leafHead) from(depth int) []byte {
	s := l.stored()
	return s[depth-(l.keyLen()-len(s)):]
}

// matches reports whether l is the leaf of key. l holds the key from its base
// on, which a key of the leaf's length has at the same distance from its end,
// so no depth is needed; the bytes from the base down to l are checked twice.
// Kept small enough to inline into find.
func (l *leafHead) matches(key []byte) bool {
	if l.klen == longKey {
		return string(key) == *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	}
	k := int(l.klen)
	return len(key) == int(l.kl) && string(key[len(key)-k:]) == unsafe.String((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), k)
}

// wholeKey returns l's whole key, from its position in a rekeyFunc call: pre,
// the byte b unless it is negative, then l's key from there on.
func wholeKey(l *leafHead, pre []byte, b int) []byte {
	k := append(make([]byte, 0, l.keyLen()), pre...)
	at := len(pre)
	if b >= 0 {
		k = append(k, byte(b))
		at++
	}
	return append(k, l.from(at)...)
}

// fillHead writes the bytes of a rekeyFunc call's whole key from depth on into
// dst, as many as dst holds: the rest of pre, then b.
func fillHead(dst, pre []byte, b, depth int) {
	if n := copy(dst, pre[depth:]); n < len(dst) {
		dst[n] = byte(b)
	}
}

// slots returns all child slots of n, including the one its term takes; n must
// not be a 256-way or a range node.
func slots(n *header) []*header {
	switch n.kind {
	case kN5:
		return asN5(n).child[:]
	case kN12:
		return asN12(n).child[:]
	case kN26:
		return asN26(n).child[:]
	}
	return asN58(n).child[:]
}

// slotCap and termOff give, by kind, the number of child slots and the offset
// of the last one, where the term sits. The 256-way node's header count
// (255) never equals its slotCap (0): its term slot is its own.
var (
	slotCap = [256]uint8{kN5: 5, kN12: 12, kN26: 26, kN58: 58}
	termOff = [256]uintptr{
		kN5:   unsafe.Offsetof(node5{}.child) + 4*ptrSize,
		kN12:  unsafe.Offsetof(node12{}.child) + 11*ptrSize,
		kN26:  unsafe.Offsetof(node26{}.child) + 25*ptrSize,
		kN58:  unsafe.Offsetof(node58{}.child) + 57*ptrSize,
		kN256: unsafe.Offsetof(node256{}.child) + 256*ptrSize,
		kR8:   rChildOff + 8*ptrSize,
		kR24:  rChildOff + 24*ptrSize,
		kR56:  rChildOff + 56*ptrSize,
		kR256: rChildOff + 256*ptrSize,
	}
)

const ptrSize = unsafe.Sizeof(uintptr(0))

// termOf returns the leaf of the key that ends exactly at n, or nil. It sits
// in n's last slot, unless the byte children fill every slot.
func termOf(n *header) *leafHead {
	k := n.kind & kindMask
	if n.count == slotCap[k] {
		return nil
	}
	return *(**leafHead)(unsafe.Add(unsafe.Pointer(n), termOff[k]))
}

// termSlot returns the address of n's term slot, which holds its term leaf
// whenever it has one.
func termSlot(n *header) **header {
	return (**header)(unsafe.Add(unsafe.Pointer(n), termOff[n.kind&kindMask]))
}

// setTermSlot stores l, or nil to remove the term, in n's term slot; n must
// have room for it.
func setTermSlot(n *header, l *leafHead) { *termSlot(n) = leafHdr(l) }

// full reports whether n has no room for another byte child or a term.
func full(n *header) bool {
	if n.kind == kN256 || isRange(n.kind) {
		return false
	}
	s := slots(n)
	return int(n.count) == len(s) || (int(n.count) == len(s)-1 && s[len(s)-1] != nil)
}

// sorted returns the child bytes and children of a node that keeps them in
// sorted arrays: n must be a 5- or 12-way node.
func sorted(n *header) ([]byte, []*header) {
	if n.kind == kN5 {
		x := asN5(n)
		return x.keys[:x.count], x.child[:x.count]
	}
	x := asN12(n)
	return x.keys[:x.count], x.child[:x.count]
}

// bitmapOf returns the bitmap and the full child array of a 26- or 58-way
// node; the first count children are in use, in key-byte order.
func bitmapOf(n *header) (*[4]uint64, []*header) {
	if n.kind == kN26 {
		x := asN26(n)
		return &x.bitmap, x.child[:]
	}
	x := asN58(n)
	return &x.bitmap, x.child[:]
}
