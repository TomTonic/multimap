// Package art is the adaptive radix tree behind multimap.Ordered.
//
// Design (see bench/README.md for the measurements behind each choice):
//
//   - Byte nodes have 5, 12, 26, 58 or 257 child slots in 64, 128, 256, 512
//     and 2080 bytes. Up to 512 bytes these are Go size classes, so every node
//     is cache-line aligned; above 512 bytes Go prepends a malloc header,
//     which only the rare 256-way node pays.
//   - Children are kept in key-byte order. The 5- and 12-way nodes store their
//     bytes in a sorted array searched with SWAR, eight bytes per step; the
//     26- and 58-way nodes find a child by the rank of its byte in a 256-bit
//     bitmap (branch-free); the 256-way node indexes directly.
//   - The node header is 16 bytes: type, child count, prefix length and the
//     first 12 bytes of the common prefix. The rest of a longer common prefix
//     follows the node in the same object (prefix.go), so every key byte on
//     the way down is checked in the nodes (pessimistic common prefix).
//   - A leaf holds only the part of its key below the node it was created
//     under (lazy expansion: a subtree with one key is just its leaf), inline
//     right after a 6-byte head; only a remainder of more than 254 bytes
//     makes the leaf hold its whole key as a separate string. The shared
//     parts of the keys are stored once, in the nodes.
//   - The values of an entry follow its key remainder in the same object, a
//     single-key page of 32, 64, 128, 256, 384 or 512 bytes (internal/skpage),
//     which grows through these classes as values arrive: as bytes with a
//     length each for strings (singlekey.go), as an array of T for small
//     pointer-free values and for values that are one pointer (fixedkey.go). A
//     page without pointers is never scanned by the garbage collector; one with
//     pointers is allocated as a Go type that marks them. An entry whose values
//     do not fit, and every entry of the other types of value, is a value
//     overflow: the key remainder and a pointer to a Set3 of the values.
//   - A key that ends at a byte node (a prefix of other keys) is that
//     node's end page. It takes the node's last child slot, which the byte
//     children reach only when there is no end page: few keys are prefixes of
//     others, so no node pays a field for them.
//
// The tree code is not generic. It works on singleKeyHead, the key part every leaf
// starts with, so no generic dictionary calls sit on the traversal path; only
// Map[T], which creates leaves and reaches their values, is generic.
//
// A Tree is not safe for concurrent use when any goroutine writes; concurrent
// readers alone are safe, because reads never modify the tree.
package art

import (
	"math/bits"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// type tells the type of a node or page, and a page's size class: an object
// is a single-key page iff its type is at most maxSingleKeyByte (see isSingleKey). The types step by
// two: the lowest bit of the type byte of a leaf or page is bit 8 of the length
// of its key remainder (see singleKeyHead.rem), so the type of such an object is
// type&^1.
type objType uint8

const (
	kValueOverflow objType = (iota + 1) << 1 // value overflow; kValueOverflow+2c is a single-key page of class c (see skpage)
	_
	_
	_
	_
	_
	kLastSingleKey // single-key page of the largest class
	kMultiKey      // multi-key page of the smallest class; kMultiKey+2c is that of class c (see mkpage)
	_
	_
	_
	_
	kLastMultiKey // multi-key page of the largest class
	kN5
	kN12
	kN26
	kN58
	kN256
)

// A descent ends at an object of a type byte up to kLastMultiKey: the pages
// come first, so that one comparison detects them: single-key pages first (up to
// maxSingleKeyByte), then multi-key pages. The byte of a page may have its lowest bit set
// (a single-key page: a long remainder, a multi-key page: a common prefix of 256 bytes or more),
// so the largest byte of a class is its largest type plus one.
const (
	maxSingleKeyByte = kLastSingleKey | 1
	maxPageByte      = kLastMultiKey | 1
)

// objTypeMask maps a type to an index of the tables below, which hold every type.
const objTypeMask = 63

// Shrink thresholds: a node turns into the next smaller type once it holds
// this many byte children or fewer. They lie below the next smaller capacity
// less the end page's slot, so a node that has just grown does not shrink back
// after one removal.
const (
	shrink12  = 3
	shrink26  = 9
	shrink58  = 22
	shrink256 = 48
)

// header is the common start of all byte nodes (16 B).
type header struct {
	objType objType
	count   uint8                // byte children; a 256-way node keeps its count in node256.total
	plen    uint16               // length of the common prefix, or longPrefix
	prefix  [swar.PrefixLen]byte // first min(plen, 12) bytes of the common prefix; the rest is in the tail
}

// longPrefix in header.plen stands for a common prefix of that many bytes or more, whose
// tail is a string: its length is then 12 plus the string's (see prefixLen).
const longPrefix = 1<<16 - 1

// singleKeyHead is the start of every leaf, a single-key page or a value overflow (4 B): the head of the page
// (internal/page): type, the length of the key part, the number of values, rawWords. The key part, the remainder
// of the leaf's key, follows at keyOff; a whole key held as a string sits at strOff.
//
// A leaf holds its key from the path length it stands at: the bytes before are the path to it, stored in the
// nodes above. When a node is put above the leaf it loses the bytes of the path (Skip, in place); when a node above it goes
// away it gets them back (Prepend, see rekeyFunc).
type singleKeyHead struct {
	objType objType // kValueOverflow, or kValueOverflow+2c for a single-key page of class c; its lowest bit is bit 8 of the length of the key part
	klen    uint8   // bits 0 to 7 of the length of the inline key part, or of longKey for a key held as a string
	n       uint8   // single-key pages: number of values
	raw     uint8   // rawWords of a page of pointers (internal/page); in a value overflow the offset of its value set in words
}

// isSingleKey reports whether an object of type k is a single-key page (in any of its
// forms: the page, the value overflow).
func isSingleKey(k objType) bool { return k <= maxSingleKeyByte }

// isPage reports whether an object of type k is a page of either kind, which ends a
// descent; isMultiKey tells a multi-key page among them.
func isPage(k objType) bool { return k <= maxPageByte }

// isMultiKey reports whether an object of type k is a multi-key page, given that it is
// a page.
func isMultiKey(k objType) bool { return k > maxSingleKeyByte }

// cls returns the size class of a single-key page (1 for 32 bytes), or 0 for a
// value overflow.
func (l *singleKeyHead) cls() uint8 { return uint8(l.objType&^1-kValueOverflow) >> 1 }

// isValueOverflow reports whether l is a value overflow (class 0).
func (l *singleKeyHead) isValueOverflow() bool { return l.objType&^1 == kValueOverflow }

// rem returns the length of the inline key remainder, or longKey: nine bits,
// the lowest bit of the type byte on top of klen.
func (l *singleKeyHead) rem() int { return int(l.klen) | int(l.objType&1)<<8 }

// setRem sets the remainder length n, at most longKey, and keeps the type.
func (l *singleKeyHead) setRem(n int) {
	l.objType = l.objType&^1 | objType(n>>8)
	l.klen = uint8(n)
}

// longKey in the remainder length of a leaf marks a whole key held as a string.
const longKey = 1<<9 - 1

// keyOff is the offset of an inline key in every leaf: right after the
// singleKeyHead. A string key sits at strOff, where a string is aligned.
const (
	keyOff = unsafe.Sizeof(singleKeyHead{})
	strOff = 8
)

// rekeyFunc moves leaf l up to pathLen, which is below the path length it stands at: it returns a leaf with the same
// values whose key part has the bytes in front of it that the node above l no longer holds. The tree calls it
// when a node above l goes away and l takes its place. l's key from pathLen on is pre from there, the byte b
// unless b is negative, then l's key part. Map[T] supplies it; a multi-key page that cannot take the bytes
// answers nil.
type rekeyFunc func(l *singleKeyHead, pre []byte, b int, pathLen int) *singleKeyHead

// Every node type but the 256-way one keeps its end page, if any, in its
// last child slot, which is free whenever there is an end page (see endPageOf).

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

// node256 holds the child of byte b in child[b] and its end page in child[256].
// header.count is fixed at 255, since the byte children can number 256.
type node256 struct { // 2080 B
	header
	total uint16 // byte children
	_     [6]byte
	child [257]*header
}

// The casts below are valid because every node and leaf type starts with a
// type byte at offset 0, and a pointer is only ever cast to the type its type
// names. The garbage collector traces objects by their allocated type, not by
// the static type of the pointer.

func asSingleKey(h *header) *singleKeyHead  { return (*singleKeyHead)(unsafe.Pointer(h)) }
func asN5(h *header) *node5                 { return (*node5)(unsafe.Pointer(h)) }
func asN12(h *header) *node12               { return (*node12)(unsafe.Pointer(h)) }
func asN26(h *header) *node26               { return (*node26)(unsafe.Pointer(h)) }
func asN58(h *header) *node58               { return (*node58)(unsafe.Pointer(h)) }
func asN256(h *header) *node256             { return (*node256)(unsafe.Pointer(h)) }
func singleKeyHdr(l *singleKeyHead) *header { return (*header)(unsafe.Pointer(l)) }

// stored returns the key part of the leaf: its key from the path length on. The slice aliases the leaf and must
// not be modified.
func (l *singleKeyHead) stored() []byte {
	if n := l.rem(); n != longKey {
		return unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), n)
	}
	s := *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// matches reports whether the key part of l is rest, the key from the path length on. Kept small enough to inline
// into find.
func (l *singleKeyHead) matches(rest []byte) bool {
	k := l.rem()
	if k == longKey {
		return string(rest) == *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	}
	return len(rest) == k && string(rest) == unsafe.String((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), k)
}

// frontOf returns the bytes that a leaf that moves up to pathLen gets in front of its key part (see
// rekeyFunc): the rest of pre from there on, then the byte b unless it is negative.
func frontOf(pre []byte, b, pathLen int) []byte {
	front := append(make([]byte, 0, len(pre)-pathLen+1), pre[pathLen:]...)
	if b >= 0 {
		front = append(front, byte(b))
	}
	return front
}

// slots returns all child slots of n, including the one its end page takes; n must
// not be a 256-way node.
func slots(n *header) []*header {
	switch n.objType {
	case kN5:
		return asN5(n).child[:]
	case kN12:
		return asN12(n).child[:]
	case kN26:
		return asN26(n).child[:]
	}
	return asN58(n).child[:]
}

// slotCap and endPageOff give, by type, the number of child slots and the offset
// of the last one, where the end page sits. The 256-way node's header count
// (255) never equals its slotCap (0): its end page slot is its own.
var (
	slotCap    = [64]uint8{kN5: 5, kN12: 12, kN26: 26, kN58: 58}
	endPageOff = [64]uintptr{
		kN5:   unsafe.Offsetof(node5{}.child) + 4*ptrSize,
		kN12:  unsafe.Offsetof(node12{}.child) + 11*ptrSize,
		kN26:  unsafe.Offsetof(node26{}.child) + 25*ptrSize,
		kN58:  unsafe.Offsetof(node58{}.child) + 57*ptrSize,
		kN256: unsafe.Offsetof(node256{}.child) + 256*ptrSize,
	}
)

const ptrSize = unsafe.Sizeof(uintptr(0))

// endPageOf returns the leaf of the key that ends exactly at n, or nil. It sits
// in n's last slot, unless the byte children fill every slot.
func endPageOf(n *header) *singleKeyHead {
	k := n.objType & objTypeMask
	if n.count == slotCap[k] {
		return nil
	}
	return *(**singleKeyHead)(unsafe.Add(unsafe.Pointer(n), endPageOff[k]))
}

// endPageSlot returns the address of n's end page slot, which holds its end page
// whenever it has one.
func endPageSlot(n *header) **header {
	return (**header)(unsafe.Add(unsafe.Pointer(n), endPageOff[n.objType&objTypeMask]))
}

// setEndPageSlot stores l, or nil to remove the end page, in n's end page slot; n must
// have room for it.
func setEndPageSlot(n *header, l *singleKeyHead) { *endPageSlot(n) = singleKeyHdr(l) }

// full reports whether n has no room for another byte child or an end page.
func full(n *header) bool {
	if n.objType == kN256 {
		return false
	}
	s := slots(n)
	return int(n.count) == len(s) || (int(n.count) == len(s)-1 && s[len(s)-1] != nil)
}

// sorted returns the child bytes and children of a node that keeps them in
// sorted arrays: n must be a 5- or 12-way node.
func sorted(n *header) ([]byte, []*header) {
	if n.objType == kN5 {
		x := asN5(n)
		return x.keys[:x.count], x.child[:x.count]
	}
	x := asN12(n)
	return x.keys[:x.count], x.child[:x.count]
}

// bitmapOf returns the bitmap and the full child array of a 26- or 58-way
// node; the first count children are in use, in key-byte order.
func bitmapOf(n *header) (*[4]uint64, []*header) {
	if n.objType == kN26 {
		x := asN26(n)
		return &x.bitmap, x.child[:]
	}
	x := asN58(n)
	return &x.bitmap, x.child[:]
}

// eachByteNode calls fn for every byte child of the byte node n in byte order.
func eachByteNode(n *header, fn func(b byte, c *header)) {
	switch n.objType {
	case kN5, kN12:
		keys, child := sorted(n)
		for i, c := range child {
			fn(keys[i], c)
		}
	case kN26, kN58:
		bm, child := bitmapOf(n)
		i := 0
		for w, set := range bm {
			for ; set != 0; set &= set - 1 {
				fn(byte(w<<6+bits.TrailingZeros64(set)), child[i])
				i++
			}
		}
	default:
		for b, c := range asN256(n).child[:256] {
			if c != nil {
				fn(byte(b), c)
			}
		}
	}
}
