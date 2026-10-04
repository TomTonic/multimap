package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map of strings (Map.flat == 3) a key lives in a single-key page
// (internal/skpage, docs/redesign/step3-skmv-design.md): the remainder and all
// the values as bytes, in one object without pointers of 32 to 512 bytes. The
// page starts like a leaf (kind, remainder length, number of values, length of
// the whole key), so everything the tree does with the key of a leaf works on it.
// Its kinds are kSet+1 to kSet+6, the kinds that flat and typed leaves have in
// other maps; kSet is still the set leaf, which takes a key whose values do not
// fit a page (the value overflow) or whose remainder is longer than a page holds.

func init() { skpage.KindBase = uint8(kSet) + 1 }

// stringType reports whether T is string, the value type that takes single-key
// pages.
func stringType[T comparable]() bool {
	var z T
	_, ok := any(z).(string)
	return ok
}

func asSK(l *leafHead) *skpage.Page   { return (*skpage.Page)(unsafe.Pointer(l)) }
func skLeaf(p *skpage.Page) *leafHead { return (*leafHead)(unsafe.Pointer(p)) }

// view returns the bytes of s without copying them. The page copies what it
// stores.
func view(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }

// newSK allocates an empty single-key page that holds key from base on, or a
// set leaf if the remainder or the key is too long for a page. It is a
// newLeafFunc; addSK adds the first value.
func newSK(key []byte, base int) *leafHead {
	if len(key) > maxKeyLen || len(key)-base > skpage.MaxRemainder {
		return newSetLeaf3(key, base, set3.EmptyWithCapacity[string](8))
	}
	return skLeaf(skpage.Empty(key[base:], len(key)))
}

// addSK adds v to the values of the page l of key, which sits in slot loc. A
// value that the page cannot hold makes the key a set leaf (spillSK).
func addSK(loc **header, l *leafHead, key []byte, v string) {
	p := asSK(l)
	switch q, res := p.Add(view(v)); {
	case res == skpage.Full:
		*loc = leafHdr(spillSK(p, key, v))
	case q != p:
		*loc = leafHdr(skLeaf(q))
	}
}

// spillSK returns a set leaf with the values of page p of key and v, the value
// overflow of a key that outgrew its page.
func spillSK(p *skpage.Page, key []byte, v string) *leafHead {
	nl := newSetLeaf3(key, skLeaf(p).base(), set3.EmptyWithCapacity[string](uint32(p.Len()+p.Len()/2+1)))
	p.Strings(func(x string) bool { strSetAdd(nl, x); return true })
	strSetAdd(nl, v)
	return nl
}

// removeSK removes v from the page l of key, which sits in the tree; rk is the
// map's rekeyFunc.
func (t *Tree) removeSK(l *leafHead, key []byte, v string, rk rekeyFunc) {
	switch q, ok := asSK(l).Remove(view(v)); {
	case !ok:
	case q == nil:
		t.remove(key, rk)
	case q != asSK(l):
		*t.findSlot(key) = leafHdr(skLeaf(q))
	}
}

// unspillSK returns the page for the set leaf l, whose values have shrunk to
// fit a page of skpage.BackLimit bytes, or nil if they do not.
func unspillSK(l *leafHead) *skpage.Page {
	if l.klen == longKey {
		return nil
	}
	s := *strSetOf(l)
	if int(s.Size()) > skpage.BackLimit/2 { // every value takes two bytes at least
		return nil
	}
	need, ok := skpage.Header+int(l.klen), true
	vs := make([][]byte, 0, s.Size())
	strSetEach(l, func(x string) bool {
		need += 1 + len(x)
		if ok = len(x) <= skpage.MaxValue && need <= skpage.BackLimit; ok {
			vs = append(vs, view(x))
		}
		return ok
	})
	if !ok {
		return nil
	}
	return skpage.Build(l.stored(), int(l.kl), vs)
}

// rekeySK is the map's rekeyFunc for strings (see rekeyFunc): a page takes the
// bytes in front of its remainder, or becomes a set leaf if they do not fit; a
// set leaf takes them in its key area if they fit, else it is copied.
func rekeySK(l *leafHead, pre []byte, b, pathLen int) *leafHead {
	if l.kind == kSet {
		if setPrepend(l, pre, b, pathLen, set3KeyCap) {
			return l
		}
		return newSetLeaf3(wholeKey(l, pre, b), pathLen, *strSetOf(l))
	}
	front := make([]byte, l.base()-pathLen)
	fillHead(front, pre, b, pathLen)
	p := asSK(l)
	if q := p.Prepend(front); q != nil {
		return skLeaf(q)
	}
	nl := newSetLeaf3(wholeKey(l, pre, b), pathLen, set3.EmptyWithCapacity[string](uint32(p.Len()+p.Len()/2+1)))
	p.Strings(func(x string) bool { strSetAdd(nl, x); return true })
	return nl
}

// The set leaf of a map of strings (Map.flat == 3) is not the leaf of
// vset.Set: behind the head and the key area it holds one pointer to a Set3
// (github.com/TomTonic/Set3), a hash set that takes any string. The page is the
// small stage, so the inline stage and the array stage of vset, which a key
// with a few values would use, are not needed.
//
// It is an object of the grid of the pages, 32, 64, 128 or 256 bytes: the head
// takes 6 bytes, the pointer the last 8, and the key area what is left, 18, 50,
// 114 or 242 bytes, so the bytes the 64-byte vset.Set took go to the key and not
// to padding. A remainder of more than 242 bytes is held as a string.
type setLeaf3[K keyArea3] struct {
	leafHead
	k   K
	set *set3.Set3[string]
}

// keyArea3 is the storage of the key of the set leaf of a string map: the key
// areas of the grid, or a string.
type keyArea3 interface {
	[18]byte | [50]byte | [114]byte | [242]byte | string
}

// maxInline3 is the longest remainder the set leaf of a string map holds inline.
const maxInline3 = 242

// Offsets of the pointer in the set leaf of a string map: they follow from the
// key areas of the grid.
var (
	setOff3a   = unsafe.Offsetof(setLeaf3[[18]byte]{}.set)
	setOff3b   = unsafe.Offsetof(setLeaf3[[50]byte]{}.set)
	setOff3c   = unsafe.Offsetof(setLeaf3[[114]byte]{}.set)
	setOff3d   = unsafe.Offsetof(setLeaf3[[242]byte]{}.set)
	setOff3Str = unsafe.Offsetof(setLeaf3[string]{}.set)
)

// set3KeyCap returns the size of the key area of a set leaf of a string map
// whose key remainder is klen bytes long, at most maxInline3 (see setPrepend).
func set3KeyCap(klen int) int {
	for _, c := range [...]int{18, 50, 114} {
		if klen <= c {
			return c
		}
	}
	return maxInline3
}

// strSetOf returns the slot that holds the value set of set leaf l of a string
// map.
func strSetOf(l *leafHead) **set3.Set3[string] {
	off := setOff3Str
	switch k := l.klen; {
	case k <= 18:
		off = setOff3a
	case k <= 50:
		off = setOff3b
	case k <= 114:
		off = setOff3c
	case k <= maxInline3:
		off = setOff3d
	}
	return (**set3.Set3[string])(unsafe.Add(unsafe.Pointer(l), off))
}

// newSetLeaf3 allocates a set leaf of a string map that holds key from base on,
// or the whole key as a string when the rest is too long to hold inline, and
// holds set as its value set.
func newSetLeaf3(key []byte, base int, set *set3.Set3[string]) *leafHead {
	var l *leafHead
	if len(key) > maxKeyLen || len(key)-base > maxInline3 {
		sl := &setLeaf3[string]{k: string(key)}
		sl.kind, sl.klen = kSet, longKey
		l = &sl.leafHead
	} else {
		s, kl := key[base:], len(key)
		switch k := len(s); {
		case k <= 18:
			l = newInline3[[18]byte](s, kl)
		case k <= 50:
			l = newInline3[[50]byte](s, kl)
		case k <= 114:
			l = newInline3[[114]byte](s, kl)
		default:
			l = newInline3[[242]byte](s, kl)
		}
	}
	*strSetOf(l) = set
	return l
}

// newInline3 allocates a set leaf of a string map that holds the remainder s
// of a key of kl bytes inline in an array of type K.
func newInline3[K [18]byte | [50]byte | [114]byte | [242]byte](s []byte, kl int) *leafHead {
	l := &setLeaf3[K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.kind, l.klen, l.kl = kSet, uint8(len(s)), uint16(kl)
	return &l.leafHead
}

// setLeaf3Size returns the size of the set leaf of a string map with a key
// remainder of klen bytes (longKey: a key held as a string).
func setLeaf3Size(klen uint8) uintptr {
	switch k := klen; {
	case k <= 18:
		return unsafe.Sizeof(setLeaf3[[18]byte]{})
	case k <= 50:
		return unsafe.Sizeof(setLeaf3[[50]byte]{})
	case k <= 114:
		return unsafe.Sizeof(setLeaf3[[114]byte]{})
	case k <= maxInline3:
		return unsafe.Sizeof(setLeaf3[[242]byte]{})
	}
	return unsafe.Sizeof(setLeaf3[string]{})
}

// strSetAdd adds v to the set of leaf l.
func strSetAdd(l *leafHead, v string) { (*strSetOf(l)).Add(v) }

// strSetEach calls yield for every value of the set of leaf l and reports whether
// it ran to completion.
func strSetEach(l *leafHead, yield func(string) bool) bool {
	for v := range (*strSetOf(l)).MutableRange() {
		if !yield(v) {
			return false
		}
	}
	return true
}
