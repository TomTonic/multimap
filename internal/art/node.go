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
	_
	_
	_
	kLastSingleKey // single-key page of the largest class
	kN5
	kN12
	kN26
	kN58
	kN256
)

// A descent ends at an object of a type byte up to maxSingleKeyByte: single-key
// pages come first, so that one comparison detects them. The byte of a page
// may have its lowest bit set (a long remainder), so the largest byte of a
// class is its largest type plus one.
const maxSingleKeyByte = kLastSingleKey | 1

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

// maxKeyLen is the longest key a leaf holds a remainder of; singleKeyHead.kl must
// hold its length. A longer key is held whole, as a string.
const maxKeyLen = 1<<16 - 1

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

// singleKeyHead is the start of every single-key page (6 B). The key remainder follows at
// keyOff; a whole key held as a string sits at strOff.
//
// A leaf holds its key from its base on, the pathLen it was created at: the
// bytes before are the path to it, stored in the nodes above. The leaf may
// move deeper later, when a node is split in above it, and still holds the
// bytes from its base, which are then also on its path; it moves up only
// with a new base (see rekeyFunc).
type singleKeyHead struct {
	objType objType // kValueOverflow, or kValueOverflow+2c for a single-key page of class c; its lowest bit is bit 8 of the remainder length
	klen    uint8   // bits 0 to 7 of the length of the inline key remainder, or of longKey for a whole key held as a string
	n       uint16  // single-key pages: number of values
	kl      uint16  // length of the whole key, if the remainder is inline
}

// isSingleKey reports whether an object of type k is a single-key page (in any of its
// forms: the page, the value overflow).
func isSingleKey(k objType) bool { return k <= maxSingleKeyByte }

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

// newLeafFunc creates a leaf for key with base base (see singleKeyHead) and returns
// its head. Map[T] supplies it so that the tree code does not need to know T.
type newLeafFunc func(key []byte, base int) *singleKeyHead

// rekeyFunc moves leaf l to pathLen, which is below its base: it returns a leaf
// with the same values that holds its key from pathLen on. The tree calls it when
// a node above l goes away and l takes its place. l's whole key is pre, then
// the byte b unless b is negative, then l's key from there on. Map[T] supplies
// it, and gets by without the whole key when the longer remainder still fits l.
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

// stored returns the key bytes the leaf holds: its key from its base on, or
// its whole key. The slice aliases the leaf and must not be modified.
func (l *singleKeyHead) stored() []byte {
	if n := l.rem(); n != longKey {
		return unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), n)
	}
	s := *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// keyLen returns the length of the leaf's whole key.
func (l *singleKeyHead) keyLen() int {
	if l.rem() != longKey {
		return int(l.kl)
	}
	return len(*(*string)(unsafe.Add(unsafe.Pointer(l), strOff)))
}

// base returns the pathLen the leaf holds its key from.
func (l *singleKeyHead) base() int { return l.keyLen() - len(l.stored()) }

// from returns the leaf's key from pathLen on; pathLen must not be below its base.
func (l *singleKeyHead) from(pathLen int) []byte {
	s := l.stored()
	return s[pathLen-(l.keyLen()-len(s)):]
}

// matches reports whether l is the leaf of key. l holds the key from its base
// on, which a key of the leaf's length has at the same distance from its end,
// so no pathLen is needed; the bytes from the base down to l are checked twice.
// Kept small enough to inline into find.
func (l *singleKeyHead) matches(key []byte) bool {
	k := l.rem()
	if k == longKey {
		return string(key) == *(*string)(unsafe.Add(unsafe.Pointer(l), strOff))
	}
	return len(key) == int(l.kl) && string(key[len(key)-k:]) == unsafe.String((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), k)
}

// wholeKey returns l's whole key, from its position in a rekeyFunc call: pre,
// the byte b unless it is negative, then l's key from there on.
func wholeKey(l *singleKeyHead, pre []byte, b int) []byte {
	k := append(make([]byte, 0, l.keyLen()), pre...)
	at := len(pre)
	if b >= 0 {
		k = append(k, byte(b))
		at++
	}
	return append(k, l.from(at)...)
}

// fillHead writes the bytes of a rekeyFunc call's whole key from pathLen on into
// dst, as many as dst holds: the rest of pre, then b.
func fillHead(dst, pre []byte, b, pathLen int) {
	if n := copy(dst, pre[pathLen:]); n < len(dst) {
		dst[n] = byte(b)
	}
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
