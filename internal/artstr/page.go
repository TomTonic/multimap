package artstr

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/lpage"
)

// A page holds the keys of one subtree in a single object of 128 to 512 bytes
// (package lpage): the part of each key below the page's depth, sorted, with the
// keys' values, which are strings, as bytes in the page. It holds no pointer, so
// that a lookup ends in one object, a range scan walks contiguous memory, and the
// garbage collector never scans it. It replaces a leaf per key and the pointer to
// each leaf, and a string header per value.
//
// Pages hold values of type string only (see Tree.small), and only keys with
// exactly one value keep it in the page. A key with more values has a leaf,
// which a range node holds like a page (see rnode.go).
//
// A page holds its keys from the depth it stands at, as a leaf does from its
// base: the bytes before are the path to it, stored in the nodes above, which a
// lookup has checked. A page that moves to another depth is rebuilt with the
// remainders to match (lpage.Page.Skip and Prepend).

func init() { lpage.KindBase = uint8(kPage) }

func asPage(h *header) *lpage.Page  { return (*lpage.Page)(unsafe.Pointer(h)) }
func pageHdr(p *lpage.Page) *header { return (*header)(unsafe.Pointer(p)) }

// bytesOf returns the bytes of s without copying them. The caller must not
// write to them.
func bytesOf(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }

// pageable reports whether key with value v may go into a page at depth: its
// remainder is not empty and entry and value fit a page.
func (t *Tree) pageable(key []byte, depth int, v string) bool {
	return t.small && lpage.FitsLen(len(key)-depth, len(v))
}

// newPageFor returns a page at depth that holds key and its value v.
func newPageFor(key []byte, depth int, v string) *lpage.Page {
	p, _ := lpage.Build([][]byte{key[depth:]}, [][]byte{bytesOf(v)}) // pageable: checked by the caller
	return p
}

// pageItems returns the keys of page p, which stands at depth, as items, in
// order, a key with several values as one item; pre holds the key bytes before
// depth.
func pageItems(p *lpage.Page, pre []byte, depth int) []item {
	n := p.Len()
	out := make([]item, 0, n)
	arena := make([]byte, 0, n*(depth+16))
	for i := 0; i < n; i++ {
		if p.IsCont(i) {
			it := &out[len(out)-1]
			it.more = append(it.more, string(p.ValueAt(i)))
			continue
		}
		start := len(arena)
		arena = p.AppendKey(append(arena, pre[:depth]...), i)
		out = append(out, item{key: arena[start:len(arena):len(arena)], val: string(p.ValueAt(i))})
	}
	return out
}

// pageFor returns the page at depth that holds items, which share their first
// depth bytes, or nil if they do not fit one: there are too many, one of them has
// a leaf, or its entry is too big or ends at depth.
func pageFor(items []item, depth int) *lpage.Page {
	pairs := 0
	for _, it := range items {
		pairs += 1 + len(it.more)
	}
	if pairs > lpage.MaxEntries() {
		return nil
	}
	keys, vals := make([][]byte, 0, pairs), make([][]byte, 0, pairs)
	for _, it := range items {
		if it.leaf != nil || !lpage.FitsLen(len(it.key)-depth, len(it.val)) {
			return nil
		}
		keys, vals = append(keys, it.key[depth:]), append(vals, bytesOf(it.val))
		for _, v := range it.more {
			if len(v) > 255 {
				return nil
			}
			keys, vals = append(keys, it.key[depth:]), append(vals, bytesOf(v))
		}
	}
	p, err := lpage.Build(keys, vals)
	if err != nil {
		return nil
	}
	return p
}

// b2i returns 1 for true and 0 for false.
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
