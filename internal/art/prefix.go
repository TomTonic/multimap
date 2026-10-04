package art

import (
	"bytes"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// Every byte of a common prefix is stored in its node (pessimistic prefix
// compression): the first PrefixLen bytes in the header, the rest in a tail
// right after the node's fixed part, in the same object. A lookup therefore
// checks every key byte on its way down, and a leaf needs to hold only the
// part of its key below its parent.
//
// Long prefixes are rare and short: with real URLs and file paths, 2-5% of the
// byte nodes have one, 18-20 bytes on average. The tail comes in a few
// classes, 16, 48 or 112 bytes inline or a string beyond, and the class
// follows from the prefix length, so the node needs no field for it. A node
// moves to a new object only when a prefix change crosses a class.

// Tail classes, see tailClass.
const (
	tailNone = iota
	tail16
	tail48
	tail112
	tailStr
)

// prefixBuf is the size of the stack buffers that hold a copy of a prefix while the
// tree code changes the node it came from: it takes the prefixes of the first two
// tail classes without an allocation, and longer ones grow it onto the heap.
const prefixBuf = swar.PrefixLen + 48

// tailClass returns the tail class of a node with a prefix of plen bytes.
func tailClass(plen int) int {
	switch t := plen - swar.PrefixLen; {
	case t <= 0:
		return tailNone
	case t <= 16:
		return tail16
	case t <= 48:
		return tail48
	case t <= 112:
		return tail112
	}
	return tailStr
}

// fixedSize is the size of each node type without its tail, which is where
// the tail starts.
var fixedSize = [64]uintptr{
	kN5:   unsafe.Sizeof(node5{}),
	kN12:  unsafe.Sizeof(node12{}),
	kN26:  unsafe.Sizeof(node26{}),
	kN58:  unsafe.Sizeof(node58{}),
	kN256: unsafe.Sizeof(node256{}),
}

type nodeTypes interface {
	node5 | node12 | node26 | node58 | node256
}

type tailTypes interface {
	[16]byte | [48]byte | [112]byte | string
}

// tailed is a node of type N with a prefix tail of type T. Every node type is a
// multiple of 8 bytes, so the tail starts right at fixedSize.
type tailed[N nodeTypes, T tailTypes] struct {
	n N
	t T
}

// allocOf allocates a zeroed node of type N with a tail of class tc. It is
// allocated with its real type, so the garbage collector sees its children.
func allocOf[N nodeTypes](tc int) *header {
	switch tc {
	case tailNone:
		return (*header)(unsafe.Pointer(new(N)))
	case tail16:
		return (*header)(unsafe.Pointer(new(tailed[N, [16]byte])))
	case tail48:
		return (*header)(unsafe.Pointer(new(tailed[N, [48]byte])))
	case tail112:
		return (*header)(unsafe.Pointer(new(tailed[N, [112]byte])))
	}
	return (*header)(unsafe.Pointer(new(tailed[N, string])))
}

// newNode allocates an empty node of type k with room for a prefix of plen
// bytes; the caller stores the prefix (storePrefix).
func newNode(k objType, plen int) *header {
	tc := tailClass(plen)
	var h *header
	switch k {
	case kN5:
		h = allocOf[node5](tc)
	case kN12:
		h = allocOf[node12](tc)
	case kN26:
		h = allocOf[node26](tc)
	case kN58:
		h = allocOf[node58](tc)
	default:
		h = allocOf[node256](tc)
		h.count = 255
	}
	h.objType = k
	return h
}

// newLike allocates a node of type k with n's header and prefix, for n to grow
// or shrink into. The caller copies the children.
func newLike(n *header, k objType) *header {
	pl := n.prefixLen()
	y := newNode(k, pl)
	*y = *n
	y.objType = k
	switch tailClass(pl) {
	case tailNone:
	case tailStr:
		*(*string)(tailPtr(y)) = *(*string)(tailPtr(n))
	default:
		copy(unsafe.Slice((*byte)(tailPtr(y)), pl-swar.PrefixLen), prefixTail(n))
	}
	return y
}

// tailPtr returns the address of n's tail.
func tailPtr(n *header) unsafe.Pointer {
	return unsafe.Add(unsafe.Pointer(n), fixedSize[n.objType&objTypeMask])
}

// prefixLen returns the length of n's common prefix.
func (n *header) prefixLen() int {
	if n.plen != longPrefix {
		return int(n.plen)
	}
	return swar.PrefixLen + len(*(*string)(tailPtr(n)))
}

// prefixTail returns the bytes of n's prefix beyond the first PrefixLen; n's prefix
// must be longer than that.
func prefixTail(n *header) []byte {
	pl := n.prefixLen()
	if tailClass(pl) == tailStr {
		s := *(*string)(tailPtr(n))
		return unsafe.Slice(unsafe.StringData(s), len(s))
	}
	return unsafe.Slice((*byte)(tailPtr(n)), pl-swar.PrefixLen)
}

// appendPrefix appends n's whole prefix to dst.
func appendPrefix(dst []byte, n *header) []byte {
	pl := n.prefixLen()
	dst = append(dst, n.prefix[:min(pl, swar.PrefixLen)]...)
	if pl > swar.PrefixLen {
		dst = append(dst, prefixTail(n)...)
	}
	return dst
}

// storePrefix stores prefix p in n, whose object must have p's tail class.
func storePrefix(n *header, p []byte) {
	n.plen = uint16(min(len(p), longPrefix))
	n.prefix = [swar.PrefixLen]byte{}
	copy(n.prefix[:], p)
	switch tailClass(len(p)) {
	case tailNone:
	case tailStr:
		*(*string)(tailPtr(n)) = string(p[swar.PrefixLen:])
	default:
		copy(unsafe.Slice((*byte)(tailPtr(n)), len(p)-swar.PrefixLen), p[swar.PrefixLen:])
	}
}

// withPrefix gives n the prefix p and returns n, or a copy of n in a new object
// when p needs another tail class. p must not alias n.
func withPrefix(n *header, p []byte) *header {
	if tailClass(len(p)) != tailClass(n.prefixLen()) {
		m := newNode(n.objType, len(p))
		copyFixed(m, n)
		n = m
	}
	storePrefix(n, p)
	return n
}

// copyFixed copies the fixed part of node src, children included, into dst of
// the same type. The copy is typed, so the garbage collector sees the pointers
// move.
func copyFixed(dst, src *header) {
	switch src.objType {
	case kN5:
		*asN5(dst) = *asN5(src)
	case kN12:
		*asN12(dst) = *asN12(src)
	case kN26:
		*asN26(dst) = *asN26(src)
	case kN58:
		*asN58(dst) = *asN58(src)
	default:
		*asN256(dst) = *asN256(src)
	}
}

// prefixMatches reports whether key[pathLen:] starts with n's whole prefix of pl
// bytes, pl > 0.
func prefixMatches(n *header, pl int, key []byte, pathLen int) bool {
	if !swar.MatchPrefix(&n.prefix, pl, key, pathLen) {
		return false
	}
	return pl <= swar.PrefixLen || bytes.Equal(prefixTail(n), key[pathLen+swar.PrefixLen:pathLen+pl])
}

// prefixLcp returns the length m of the common prefix of n's prefix of pl bytes
// and rest, and, if m < pl, n's prefix byte at m.
func prefixLcp(n *header, pl int, rest []byte) (m int, c byte) {
	h := min(pl, swar.PrefixLen)
	if m = swar.Lcp(n.prefix[:h], rest); m < h {
		return m, n.prefix[m]
	}
	if pl == h {
		return m, 0
	}
	t := prefixTail(n)
	if len(rest) > h {
		m += swar.Lcp(t, rest[h:])
	}
	if m < pl {
		c = t[m-h]
	}
	return m, c
}
