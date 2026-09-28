package art

import (
	"bytes"
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
func (t *Tree) find(key []byte) *leafHead {
	n := t.root
	depth := 0
	skipped := false // whether path bytes beyond the twelfth went unchecked
	for n != nil {
		if n.kind == kLeaf {
			// Every key below a path starts with it, so once the whole path
			// to the leaf is checked, only the rest of the key needs comparing.
			from := depth
			if skipped {
				from = 0
			}
			if l := asLeaf(n); bytes.Equal(l.key()[from:], key[from:]) {
				return l
			}
			return nil
		}
		if n.plen != 0 {
			pl := int(n.plen)
			if !swar.Match8(&n.prefix, pl, key, depth) {
				pl = n.pathLen(depth)
				if !swar.MatchPrefix(&n.prefix, pl, key, depth) {
					return nil
				}
				if pl > swar.PrefixLen {
					skipped = true
				}
			}
			depth += pl
		}
		if depth == len(key) {
			t := termOf(n)
			if t == nil {
				return nil
			}
			n = leafHdr(t)
			continue
		}
		b := key[depth]
		depth++
		switch n.kind {
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
// tree. Writers use it to replace a leaf they have found. Since the key is
// present, its paths need no checking: the descent only follows them.
func (t *Tree) findSlot(key []byte) **header {
	loc, depth := &t.root, 0
	for (*loc).kind != kLeaf {
		n := *loc
		depth += n.pathLen(depth)
		if depth == len(key) {
			loc = termSlot(n)
			continue
		}
		loc = findLoc(n, key[depth])
		depth++
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
	switch n.kind {
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
	if i == 8 && n.kind == kN12 {
		i = 8 + swar.Index8(swar.Word(asN12(n).keys[8:16]), b)
	}
	if i >= int(n.count) {
		return nil
	}
	if n.kind == kN5 {
		return &x.child[i]
	}
	return &asN12(n).child[i]
}

// minLeaf returns the leaf with the smallest key below n.
func minLeaf(n *header) *leafHead {
	for n.kind != kLeaf {
		if t := termOf(n); t != nil {
			return t // a prefix of every other key below n
		}
		if n.kind == kN256 {
			n = asN256(n).child[firstChild(asN256(n), 0)]
			continue
		}
		n = slots(n)[0]
	}
	return asLeaf(n)
}

// maxLeaf returns the leaf with the largest key below n.
func maxLeaf(n *header) *leafHead {
	for n.kind != kLeaf {
		if n.kind == kN256 {
			x := asN256(n)
			k := 255
			for x.child[k] == nil {
				k--
			}
			n = x.child[k]
			continue
		}
		n = slots(n)[n.count-1]
	}
	return asLeaf(n)
}

// firstChild returns the smallest byte from b on that has a child in x; x
// must have one.
func firstChild(x *node256, b int) int {
	for x.child[b] == nil {
		b++
	}
	return b
}

// fullPrefix returns the complete compressed path of n, whose first byte is
// at key depth depth. Paths longer than the 12 bytes stored inline are read
// from a leaf below n; buf backs the result otherwise.
func fullPrefix(n *header, depth int, buf *[swar.PrefixLen]byte) []byte {
	if n.plen <= swar.PrefixLen {
		*buf = n.prefix
		return buf[:n.plen]
	}
	return minLeaf(n).key()[depth : depth+n.pathLen(depth)]
}
