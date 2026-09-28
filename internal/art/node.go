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
//     first 12 bytes of the compressed path. Longer paths are skipped
//     optimistically and verified against the full key, which every leaf
//     holds (lazy expansion: a subtree with one key is just its leaf).
//   - A leaf holds its key inline, in the smallest of four size classes (16,
//     32, 48 or 64 bytes) that fits, and its values right after it; only a
//     longer key is a separate string. Comparing a key at the leaf therefore
//     costs no second pointer chase, and a leaf with uint64 values fills a Go
//     size class (64 B for integer keys, one cache line).
//   - A key that ends at an inner node (a prefix of other keys) is that
//     node's term leaf. It takes the node's last child slot, which the byte
//     children reach only when there is no term: few keys are prefixes of
//     others, so no node pays a field for them.
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

	"github.com/TomTonic/multimap/internal/swar"
	"github.com/TomTonic/multimap/internal/vset"
)

type kind uint8

const (
	kLeaf kind = iota + 1
	kN5
	kN12
	kN26
	kN58
	kN256
)

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

// maxInline is the longest key a leaf holds inline, in an array of 16, 32, 48
// or 64 bytes, whichever is the smallest that fits. A longer key is held as a
// string, which costs a separate allocation and a pointer chase on every
// comparison.
const maxInline = 64

// header is the common start of all inner nodes (16 B).
type header struct {
	kind   kind
	count  uint8                // byte children; a 256-way node keeps its count in node256.total
	plen   uint16               // length of the compressed path, or longPath
	prefix [swar.PrefixLen]byte // first min(plen, 12) bytes of the compressed path
}

// longPath in header.plen stands for a path of that many bytes or more; its
// length is then read from the leaves (see pathLen).
const longPath = 1<<16 - 1

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

// setPrefix stores a compressed path of plen bytes, of which p holds at least
// the first min(plen, 12).
func (h *header) setPrefix(p []byte, plen int) {
	h.plen = uint16(min(plen, longPath))
	h.prefix = [swar.PrefixLen]byte{}
	copy(h.prefix[:], p[:min(plen, swar.PrefixLen)])
}

// pathLen returns the length of n's compressed path, which starts at key depth
// depth. A path of longPath bytes or more is as long as the keys below n
// agree from depth on: n branches, so its smallest and largest key differ
// right after the path.
func (n *header) pathLen(depth int) int {
	if n.plen != longPath {
		return int(n.plen)
	}
	return longPathLen(n, depth)
}

func longPathLen(n *header, depth int) int {
	return swar.Lcp(minLeaf(n).key()[depth:], maxLeaf(n).key()[depth:])
}

// slots returns all child slots of n, including the one its term takes.
func slots(n *header) []*header {
	switch n.kind {
	case kN5:
		return asN5(n).child[:]
	case kN12:
		return asN12(n).child[:]
	case kN26:
		return asN26(n).child[:]
	case kN58:
		return asN58(n).child[:]
	}
	return asN256(n).child[:]
}

// slotCap and termOff give, by kind, the number of child slots and the offset
// of the last one, where the term sits. The 256-way node's header count
// (255) never equals its slotCap (0): its term slot is its own.
var (
	slotCap = [16]uint8{kN5: 5, kN12: 12, kN26: 26, kN58: 58}
	termOff = [16]uintptr{
		kN5:   unsafe.Offsetof(node5{}.child) + 4*ptrSize,
		kN12:  unsafe.Offsetof(node12{}.child) + 11*ptrSize,
		kN26:  unsafe.Offsetof(node26{}.child) + 25*ptrSize,
		kN58:  unsafe.Offsetof(node58{}.child) + 57*ptrSize,
		kN256: unsafe.Offsetof(node256{}.child) + 256*ptrSize,
	}
)

const ptrSize = unsafe.Sizeof(uintptr(0))

// termOf returns the leaf of the key that ends exactly at n, or nil. It sits
// in n's last slot, unless the byte children fill every slot.
func termOf(n *header) *leafHead {
	k := n.kind & 15
	if n.count == slotCap[k] {
		return nil
	}
	return *(**leafHead)(unsafe.Add(unsafe.Pointer(n), termOff[k]))
}

// setTermSlot stores l, or nil to remove the term, in n's term slot; n must
// have room for it.
func setTermSlot(n *header, l *leafHead) {
	s := slots(n)
	s[len(s)-1] = leafHdr(l)
}

// full reports whether n has no room for another byte child or a term.
func full(n *header) bool {
	if n.kind == kN256 {
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
