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
		if setPrepend(l, pre, b, pathLen) {
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
// with a few values would use, are not needed; the leaf is 56 bytes smaller.
// Its key areas are those of the other set leaf, and the pointer lies where
// vset.Set lies (valsOff).
type setLeaf3[K keyArea] struct {
	leafHead
	k   K
	set *set3.Set3[string]
}

// strSetOf returns the slot that holds the value set of set leaf l of a string
// map.
func strSetOf(l *leafHead) **set3.Set3[string] {
	return (**set3.Set3[string])(unsafe.Add(unsafe.Pointer(l), valsOff(l)))
}

// newSetLeaf3 allocates a set leaf of a string map that holds key from base on,
// or the whole key as a string when the rest is too long to hold inline, and
// holds set as its value set.
func newSetLeaf3(key []byte, base int, set *set3.Set3[string]) *leafHead {
	var l *leafHead
	if len(key) > maxKeyLen || len(key)-base > maxInline {
		sl := &setLeaf3[string]{k: string(key)}
		sl.kind, sl.klen = kSet, longKey
		l = &sl.leafHead
	} else {
		s, kl := key[base:], len(key)
		switch k := len(s); {
		case k <= 16:
			l = newInline3[[16]byte](s, kl)
		case k <= 32:
			l = newInline3[[32]byte](s, kl)
		case k <= 48:
			l = newInline3[[48]byte](s, kl)
		case k <= 64:
			l = newInline3[[64]byte](s, kl)
		case k <= 96:
			l = newInline3[[96]byte](s, kl)
		case k <= 128:
			l = newInline3[[128]byte](s, kl)
		case k <= 192:
			l = newInline3[[192]byte](s, kl)
		default:
			l = newInline3[[256]byte](s, kl)
		}
	}
	*strSetOf(l) = set
	return l
}

// newInline3 allocates a set leaf of a string map that holds the remainder s
// of a key of kl bytes inline in an array of type K.
func newInline3[K [16]byte | [32]byte | [48]byte | [64]byte | [96]byte | [128]byte | [192]byte | [256]byte](s []byte, kl int) *leafHead {
	l := &setLeaf3[K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.kind, l.klen, l.kl = kSet, uint8(len(s)), uint16(kl)
	return &l.leafHead
}

// setLeaf3Size returns the size of the set leaf of a string map with a key
// remainder of klen bytes (longKey: a key held as a string).
func setLeaf3Size(klen uint8) uintptr {
	switch k := klen; {
	case k <= 16:
		return unsafe.Sizeof(setLeaf3[[16]byte]{})
	case k <= 32:
		return unsafe.Sizeof(setLeaf3[[32]byte]{})
	case k <= 48:
		return unsafe.Sizeof(setLeaf3[[48]byte]{})
	case k <= 64:
		return unsafe.Sizeof(setLeaf3[[64]byte]{})
	case k <= 96:
		return unsafe.Sizeof(setLeaf3[[96]byte]{})
	case k <= 128:
		return unsafe.Sizeof(setLeaf3[[128]byte]{})
	case k <= 192:
		return unsafe.Sizeof(setLeaf3[[192]byte]{})
	case k <= maxInline:
		return unsafe.Sizeof(setLeaf3[[256]byte]{})
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
