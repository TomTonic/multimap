package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// Tree is the untyped adaptive radix tree; Map[T] wraps it. The zero value is
// an empty tree.
type Tree struct {
	root *header
	size int
}

// Len returns the number of keys.
func (t *Tree) Len() int { return t.size }

// Clear removes all keys.
func (t *Tree) Clear() { t.root, t.size = nil, 0 }

// find returns the leaf of key, or nil.
//
// This is the hot path of every point operation. It is one loop over the
// levels with every node search written out, so that the search helpers are
// inlined and no call is made per level; that is why it is longer than the
// project's usual function size.
func (t *Tree) find(key []byte) *singleKeyHead {
	n := t.root
	pathLen := 0
	for n != nil {
		if isSingleKey(n.objType) {
			// The nodes have checked the key up to pathLen; the leaf holds the
			// rest.
			if l := asSingleKey(n); l.matches(key) {
				return l
			}
			return nil
		}
		if n.plen != 0 {
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, pathLen) {
				if pl = longMatch(n, key, pathLen); pl < 0 {
					return nil
				}
			}
			pathLen += pl
		}
		if pathLen == len(key) {
			t := endPageOf(n)
			if t == nil {
				return nil
			}
			n = singleKeyHdr(t)
			continue
		}
		b := key[pathLen]
		pathLen++
		switch n.objType {
		case kN5:
			x := asN5(n)
			i := swar.Index8(swar.Word(x.keys[:]), b)
			if i >= int(x.count) {
				return nil
			}
			n = childAt(&x.child[0], i) // i < count <= 5
		case kN12:
			x := asN12(n)
			i := swar.Index8(swar.Word(x.keys[0:8]), b)
			if i == 8 {
				i = 8 + swar.Index8(swar.Word(x.keys[8:16]), b)
			}
			if i >= int(x.count) {
				return nil
			}
			n = childAt(&x.child[0], i) // i < count <= 12
		case kN26:
			x := asN26(n)
			if !swar.Has(&x.bitmap, b) {
				return nil
			}
			n = x.child[swar.Rank(&x.bitmap, b)]
		case kN58:
			x := asN58(n)
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

// findSlot returns the slot that holds the leaf of key, which must be in the
// tree and have a leaf. Writers use it to replace a leaf they have found.
// Since the key is present, its common prefixes need no checking: the descent only
// follows them.
func (t *Tree) findSlot(key []byte) **header {
	loc, pathLen := &t.root, 0
	for !isSingleKey((*loc).objType) {
		n := *loc
		pathLen += n.prefixLen()
		if pathLen == len(key) {
			loc = endPageSlot(n)
		} else {
			loc = findLoc(n, key[pathLen])
			pathLen++
		}
	}
	return loc
}

// childAt returns the i-th child from the first child slot c on, without the
// bounds check the compiler cannot prove away: the callers have checked i
// against the node's count, which never exceeds its slots.
func childAt(c **header, i int) *header {
	return *(**header)(unsafe.Add(unsafe.Pointer(c), uintptr(i)*ptrSize))
}

// findLoc returns the slot holding the child for byte b, or nil.
func findLoc(n *header, b byte) **header {
	switch n.objType {
	case kN26, kN58:
		bm, child := bitmapOf(n)
		if swar.Has(bm, b) {
			return &child[swar.Rank(bm, b)]
		}
		return nil
	case kN256:
		x := asN256(n)
		if x.child[b] != nil {
			return &x.child[b]
		}
		return nil
	}
	x := asN5(n) // a 5- and a 12-way node start alike
	i := swar.Index8(swar.Word(x.keys[:]), b)
	if i == 8 && n.objType == kN12 {
		i = 8 + swar.Index8(swar.Word(asN12(n).keys[8:16]), b)
	}
	if i >= int(n.count) {
		return nil
	}
	if n.objType == kN5 {
		return &x.child[i]
	}
	return &asN12(n).child[i]
}

// longMatch returns the length of n's common prefix if key[pathLen:] starts with it, and
// -1 otherwise. It is kept out of find, whose loop stays small for the common
// prefixes of at most eight bytes.
//
//go:noinline
func longMatch(n *header, key []byte, pathLen int) int {
	pl := n.prefixLen()
	if !prefixMatches(n, pl, key, pathLen) {
		return -1
	}
	return pl
}
