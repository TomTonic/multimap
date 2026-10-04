package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vpage"
)

// A page holds the keys of one subtree in a single object of 128 to 1024 bytes
// (package vpage): the part of each key below the page's base, sorted, with
// the keys' values, and no pointer, so that a lookup ends in one object, a range
// scan walks contiguous memory, and the garbage collector never scans it. It
// replaces a leaf per key and the pointer to each leaf.
//
// Pages hold values of a small pointer-free type only (see Tree.small), and
// only keys with exactly one value keep it in the page. A key with more values
// has a leaf, which a range node holds like a page (see rnode.go).
//
// A page holds its keys from its base on, as a leaf does: the bytes before the
// base are the path to it, stored in the nodes above, which a lookup has
// checked. The base is the pathLen the page was created at. A page moves up only
// with a new base (see collapse); it may sit deeper than its base, when a node is
// split in above it.

func init() { vpage.TypeBase = uint8(kMultiKey) }

// maxPagePathLen is the deepest base a page has; a key whose path is longer gets
// a leaf. maxPageRemainder is the longest part of a key a page holds from its
// base.
const (
	maxPagePathLen   = 255
	maxPageRemainder = 255
)

func asMultiKey(h *header) *vpage.Page  { return (*vpage.Page)(unsafe.Pointer(h)) }
func multiKeyHdr(p *vpage.Page) *header { return (*header)(unsafe.Pointer(p)) }

// pageable reports whether the key of a page at pathLen may go into a page.
func (t *Tree) pageable(key []byte, pathLen int) bool {
	return t.small && pathLen <= maxPagePathLen && len(key)-pathLen <= maxPageRemainder
}

// pageBase returns the base of a page at pathLen whose shortest key has minLen
// bytes: the pathLen, unless that leaves less than a head word of 8 bytes of the
// shortest key in the page. A head word costs 8 bytes whatever it holds, so the
// page then starts earlier and holds a whole head word of each key, which the
// lookup takes from the key in one load, instead of a few bytes that it has to
// shift into place; the page's keys share the bytes before the pathLen, which the
// nodes above have checked and the page checks once more. Keys of at most 8
// bytes, such as integers, are held whole.
func pageBase(pathLen, minLen int) int {
	return min(pathLen, max(0, minLen-8))
}

// newPageFor returns a page at pathLen that holds key and its raw value v.
func newPageFor(key []byte, pathLen int, v uint64) *vpage.Page {
	base := pageBase(pathLen, len(key))
	return vpage.Build(base, 1, func(int) []byte { return key[base:] }, func(int) uint64 { return v })
}

// pageItems returns the keys of page p as items, in order; pre is the key bytes
// before the page's base.
func pageItems(p *vpage.Page, pre []byte) []item {
	n := p.Len()
	out := make([]item, n)
	arena := make([]byte, 0, n*(p.Base()+16))
	var buf [maxPageRemainder]byte
	for i := range out {
		start := len(arena)
		arena = append(append(arena, pre[:p.Base()]...), p.Key(i, &buf)...)
		out[i].key = arena[start:len(arena):len(arena)]
		out[i].val = p.Val(i)
	}
	return out
}

// pageFor returns the page that holds items, which share their first pathLen
// bytes, or nil if they do not fit one: there are too many, one of them has a
// leaf, or one has too long a path or remainder.
func pageFor(items []item, pathLen int) *vpage.Page {
	if pathLen > maxPagePathLen || len(items) > 255 {
		return nil
	}
	minLen := len(items[0].key)
	for _, it := range items {
		if it.leaf != nil || len(it.key)-pathLen > maxPageRemainder {
			return nil
		}
		minLen = min(minLen, len(it.key))
	}
	base := pageBase(pathLen, minLen)
	return vpage.Build(base, len(items), func(i int) []byte { return items[i].key[base:] }, func(i int) uint64 { return items[i].val })
}

// b2i returns 1 for true and 0 for false.
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
