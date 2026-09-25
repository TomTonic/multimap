package art

import (
	"bytes"

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
	skipped := false // whether path bytes beyond the eighth went unchecked
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
			if !swar.MatchPrefix(&n.prefix, int(n.plen), key, depth) {
				return nil
			}
			if n.plen > 8 {
				skipped = true
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
		case kN25:
			x := asN25(n)
			if !swar.Has(&x.bitmap, b) {
				return nil
			}
			n = x.child[swar.Rank(&x.bitmap, b)]
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

// findLoc returns the slot holding the child for byte b, or nil.
func findLoc(n *header, b byte) **header {
	switch n.kind {
	case kN25, kN57:
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
	x := asN4(n) // a 4- and an 11-way node start alike
	i := swar.Index8(swar.Word(x.keys[:]), b)
	if i == 8 && n.kind == kN11 {
		i = 8 + swar.Index8(swar.Word(asN11(n).keys[8:16]), b)
	}
	if i >= int(n.count) {
		return nil
	}
	if n.kind == kN4 {
		return &x.child[i&3]
	}
	return &asN11(n).child[i]
}

// minLeaf returns the leaf with the smallest key below n.
func minLeaf(n *header) *leafHead {
	for n.kind != kLeaf {
		if n.term != nil {
			return n.term // a prefix of every other key below n
		}
		switch n.kind {
		case kN25, kN57:
			_, child := bitmapOf(n)
			n = child[0]
		case kN256:
			for _, c := range asN256(n).child {
				if c != nil {
					n = c
					break
				}
			}
		default:
			_, child := sorted(n)
			n = child[0]
		}
	}
	return asLeaf(n)
}

// fullPrefix returns the complete compressed path of n, whose first byte is
// at key depth depth. Paths longer than the 8 bytes stored inline are read
// from a leaf below n; buf backs the result otherwise.
func fullPrefix(n *header, depth int, buf *[8]byte) []byte {
	if n.plen <= 8 {
		*buf = n.prefix
		return buf[:n.plen]
	}
	return minLeaf(n).key()[depth : depth+int(n.plen)]
}
