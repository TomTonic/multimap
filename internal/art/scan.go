package art

import (
	"bytes"
	"math/bits"
	"runtime"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// Bounds selects a key range. A bound that is not set leaves that side open.
type Bounds struct {
	From, To         []byte
	HasFrom, HasTo   bool
	FromIncl, ToIncl bool
}

// scan calls fn for every leaf within b, in ascending key order, until fn
// returns false. leafTail is the offset of the last byte of a leaf, which
// the scan touches ahead (see touchChildren).
func (t *Tree) scan(b *Bounds, leafTail uintptr, fn func(*leafHead) bool) {
	scanRange(t.root, b, 0, b.HasFrom, b.HasTo, leafTail, fn)
}

// scanRange visits the leaves of the subtree n within b in order and returns
// false once the scan is over (fn stopped it, or the upper bound was passed).
//
// lo and hi say whether n lies on the path of b.From and of b.To. The bounds
// are checked structurally: a node's path and child bytes are compared with a
// bound only while on its path, so a subtree strictly inside the range is
// visited without reading a single key.
func scanRange(n *header, b *Bounds, depth int, lo, hi bool, leafTail uintptr, fn func(*leafHead) bool) bool {
	if n == nil {
		return true
	}
	if n.kind == kLeaf {
		return scanLeaf(asLeaf(n), b, lo, hi, fn)
	}
	if (lo || hi) && n.plen > 0 {
		var buf [8]byte
		pk := fullPrefix(n, depth, &buf)
		if lo {
			rest := b.From[depth:]
			if m := swar.Lcp(pk, rest); m < len(pk) {
				if m < len(rest) && pk[m] < rest[m] {
					return true // the whole subtree lies below From
				}
				lo = false // the whole subtree lies above From
			}
		}
		if hi {
			rest := b.To[depth:]
			if m := swar.Lcp(pk, rest); m < len(pk) {
				if m == len(rest) || pk[m] > rest[m] {
					return false // the whole subtree lies above To: done
				}
				hi = false // the whole subtree lies below To
			}
		}
	}
	depth += int(n.plen)
	// The term leaf's key is the path to n. On From's path it is below From
	// unless the path is From itself; on To's path it is below To unless the
	// path is To itself, in which case it is the last key in range.
	termIsFrom := lo && depth == len(b.From)
	termIsTo := hi && depth == len(b.To)
	if termIsFrom {
		lo = false
	}
	if n.term != nil && !lo && (!termIsFrom || b.FromIncl) && (!termIsTo || b.ToIncl) {
		if !fn(n.term) {
			return false
		}
	}
	if termIsTo {
		return false
	}
	var loB, hiB byte = 0, 255
	if lo {
		loB = b.From[depth]
	}
	if hi {
		hiB = b.To[depth]
	}
	touchChildren(n, loB, hiB, leafTail)
	if !scanChildren(n, b, depth, lo, hi, loB, hiB, leafTail, fn) {
		return false
	}
	return !hi // on To's path, everything after the child for hiB is above To
}

func scanLeaf(l *leafHead, b *Bounds, lo, hi bool, fn func(*leafHead) bool) bool {
	if lo {
		if c := bytes.Compare(l.key(), b.From); c < 0 || (c == 0 && !b.FromIncl) {
			return true
		}
	}
	if hi {
		if c := bytes.Compare(l.key(), b.To); c > 0 || (c == 0 && !b.ToIncl) {
			return false
		}
	}
	return fn(l)
}

// scanChildren visits the children with byte in [loB, hiB] in order. Only the
// child for loB stays on From's path, only the one for hiB on To's.
func scanChildren(n *header, b *Bounds, depth int, lo, hi bool, loB, hiB byte, leafTail uintptr, fn func(*leafHead) bool) bool {
	visit := func(c *header, k byte) bool {
		return scanRange(c, b, depth+1, lo && k == loB, hi && k == hiB, leafTail, fn)
	}
	switch n.kind {
	case kN25, kN57:
		bm, child := bitmapOf(n)
		i := swar.Rank(bm, loB)
		for w := int(loB >> 6); w < 4; w++ {
			set := bm[w]
			if w == int(loB>>6) {
				set &= ^uint64(0) << (loB & 63)
			}
			for ; set != 0; set &= set - 1 {
				k := byte(w<<6 + bits.TrailingZeros64(set))
				if k > hiB {
					return true
				}
				if !visit(child[i], k) {
					return false
				}
				i++
			}
		}
		return true
	case kN256:
		x := asN256(n)
		for k := int(loB); k <= int(hiB); k++ {
			if x.child[k] != nil && !visit(x.child[k], byte(k)) {
				return false
			}
		}
		return true
	}
	keys, child := sorted(n)
	for i, k := range keys {
		if k < loB {
			continue
		}
		if k > hiB {
			break
		}
		if !visit(child[i], k) {
			return false
		}
	}
	return true
}

// touchChildren reads one byte from the start of every child of n with byte in
// [loB, hiB], and from leaves also the byte at leafTail, which lies in the
// value set. The scan would otherwise take the cache misses for these objects
// one after another, each only once it has finished the previous child's
// subtree. These loads do not depend on each other, so the CPU keeps all their
// misses in flight at once. A node with fewer than two children in range has
// nothing to overlap and is skipped.
//
// Touching is unconditional: once a tree no longer fits in the cache it makes
// range scans up to 20% faster, while a tree that stays in the cache loses at
// most about 4% (bench/README.md). Where that line lies depends on the cache
// size and on what else the program keeps in it, so no fixed tree size could
// draw it.
func touchChildren(n *header, loB, hiB byte, leafTail uintptr) {
	if loB >= hiB { // at most one child in range
		return
	}
	var acc uint8
	touch := func(c *header) {
		acc += uint8(c.kind)
		if c.kind == kLeaf {
			acc += *(*uint8)(unsafe.Add(unsafe.Pointer(c), leafTail))
		}
	}
	switch n.kind {
	case kN25, kN57:
		bm, child := bitmapOf(n)
		end := swar.Rank(bm, hiB)
		if swar.Has(bm, hiB) {
			end++
		}
		start := swar.Rank(bm, loB)
		if end-start < 2 {
			return
		}
		for _, c := range child[start:end] {
			touch(c)
		}
	case kN256:
		x := asN256(n)
		for k := int(loB); k <= int(hiB); k++ {
			if c := x.child[k]; c != nil {
				touch(c)
			}
		}
	default:
		keys, child := sorted(n)
		start, end := 0, len(keys)
		for start < end && keys[start] < loB {
			start++
		}
		for end > start && keys[end-1] > hiB {
			end--
		}
		if end-start < 2 {
			return
		}
		for _, c := range child[start:end] {
			touch(c)
		}
	}
	runtime.KeepAlive(acc) // keeps the loads from being optimized away
}

// Contains reports whether key lies within b. Unordered structures use it to
// filter their keys with exactly the semantics of an ordered scan.
func (b *Bounds) Contains(key []byte) bool {
	if b.HasFrom {
		if c := bytes.Compare(key, b.From); c < 0 || (c == 0 && !b.FromIncl) {
			return false
		}
	}
	if b.HasTo {
		if c := bytes.Compare(key, b.To); c > 0 || (c == 0 && !b.ToIncl) {
			return false
		}
	}
	return true
}
