package art

import (
	"bytes"
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

// touchChildren reads one byte from the start of every child of n with byte in
// [loB, hiB], and from leaves also the byte at leafTail, the last one of the
// smallest leaf, which lies in or near the value set. The scan would otherwise take the cache misses for these objects
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
		acc += uint8(c.objType)
		if isPage(c.objType) {
			acc += *(*uint8)(unsafe.Add(unsafe.Pointer(c), leafTail))
		}
	}
	switch n.objType {
	case kN26, kN58:
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
