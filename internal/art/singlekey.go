package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map of strings (Map.flat == 3) a key lives in a single-key page
// (internal/skpage, docs/redesign/step3-skmv-design.md): the remainder and all
// the values as bytes, in one object without pointers of 32 to 512 bytes. The
// page starts like a leaf (type, remainder length, number of values, length of
// the whole key), so everything the tree does with the key of a leaf works on it.
// Its types are kValueOverflow+1 to kValueOverflow+6, the types that flat and typed leaves have in
// other maps; kValueOverflow is still the value overflow, which takes a key whose values do not
// fit a page (the value overflow) or whose remainder is longer than a page holds.

func init() { skpage.TypeBase = uint8(kValueOverflow) + 2 }

// stringType reports whether T is string, the value type that takes single-key
// pages.
func stringType[T comparable]() bool {
	var z T
	_, ok := any(z).(string)
	return ok
}

func asSK(l *singleKeyHead) *skpage.Page   { return (*skpage.Page)(unsafe.Pointer(l)) }
func skHead(p *skpage.Page) *singleKeyHead { return (*singleKeyHead)(unsafe.Pointer(p)) }

// view returns the bytes of s without copying them. The page copies what it
// stores.
func view(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }

// newSK allocates an empty single-key page that holds key from base on, or a
// value overflow if the remainder or the key is too long for a page. It is a
// newLeafFunc; addSK adds the first value.
func newSK(key []byte, base int) *singleKeyHead {
	if len(key) > maxKeyLen || len(key)-base > skpage.MaxRemainder {
		return newValueOverflow(key, base, set3.EmptyWithCapacity[string](8))
	}
	return skHead(skpage.Empty(key[base:], len(key)))
}

// addSK adds v to the values of the page l of key, which sits in slot loc. A
// value that the page cannot hold makes the key a value overflow (toValueOverflow).
func addSK(loc **header, l *singleKeyHead, key []byte, v string) {
	p := asSK(l)
	switch q, res := p.Add(view(v)); {
	case res == skpage.Full:
		*loc = singleKeyHdr(toValueOverflow(p, key, v))
	case q != p:
		*loc = singleKeyHdr(skHead(q))
	}
}

// toValueOverflow returns a value overflow with the values of page p of key and v, the value
// overflow of a key that outgrew its page.
func toValueOverflow(p *skpage.Page, key []byte, v string) *singleKeyHead {
	nl := newValueOverflow(key, skHead(p).base(), set3.EmptyWithCapacity[string](uint32(p.Len()+p.Len()/2+1)))
	p.Strings(func(x string) bool { overflowAdd(nl, x); return true })
	overflowAdd(nl, v)
	return nl
}

// removeSK removes v from the page l of key, which sits in the tree; rk is the
// map's rekeyFunc.
func (t *Tree) removeSK(l *singleKeyHead, key []byte, v string, rk rekeyFunc) {
	switch q, ok := asSK(l).Remove(view(v)); {
	case !ok:
	case q == nil:
		t.remove(key, rk)
	case q != asSK(l):
		*t.findSlot(key) = singleKeyHdr(skHead(q))
	}
}

// fromValueOverflow returns the page for the value overflow l, whose values have shrunk to
// fit a page of skpage.BackLimit bytes, or nil if they do not.
func fromValueOverflow(l *singleKeyHead) *skpage.Page {
	if l.rem() == longKey {
		return nil
	}
	s := *overflowSetOf(l)
	if int(s.Size()) > skpage.BackLimit/2 { // every value takes two bytes at least
		return nil
	}
	need, ok := skpage.Header+l.rem(), true
	vs := make([][]byte, 0, s.Size())
	overflowEach(l, func(x string) bool {
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
// bytes in front of its remainder, or becomes a value overflow if they do not fit; a
// value overflow takes them in its key area if they fit, else it is copied.
func rekeySK(l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	if l.isValueOverflow() {
		if setPrepend(l, pre, b, pathLen, overflowKeyCap) {
			return l
		}
		return newValueOverflow(wholeKey(l, pre, b), pathLen, *overflowSetOf(l))
	}
	front := make([]byte, l.base()-pathLen)
	fillHead(front, pre, b, pathLen)
	p := asSK(l)
	if q := p.Prepend(front); q != nil {
		return skHead(q)
	}
	nl := newValueOverflow(wholeKey(l, pre, b), pathLen, set3.EmptyWithCapacity[string](uint32(p.Len()+p.Len()/2+1)))
	p.Strings(func(x string) bool { overflowAdd(nl, x); return true })
	return nl
}

// The value overflow of a map of strings (Map.flat == 3) is not the leaf of
// vset.Set: behind the head and the key area it holds one pointer to a Set3
// (github.com/TomTonic/Set3), a hash set that takes any string. The page is the
// small stage, so the inline stage and the array stage of vset, which a key
// with a few values would use, are not needed.
//
// It is an object of the grid of the pages, 32, 64, 128, 256, 384 or 512 bytes:
// the head takes 6 bytes, the pointer the last 8, and the key area what is left,
// 18, 50, 114, 242, 370 or 498 bytes, so the bytes the 64-byte vset.Set took go
// to the key and not to padding. A remainder of more than 498 bytes is held as
// a string.
type valueOverflow[K overflowKeyArea] struct {
	singleKeyHead
	k   K
	set *set3.Set3[string]
}

// overflowKeyArea is the storage of the key of the value overflow of a string map: the key
// areas of the grid, or a string.
type overflowKeyArea interface {
	[18]byte | [50]byte | [114]byte | [242]byte | [370]byte | [498]byte | string
}

// maxInlineOverflow is the longest remainder the value overflow of a string map holds inline.
const maxInlineOverflow = 498

// Offsets of the pointer in the value overflow of a string map: they follow from the
// key areas of the grid.
var (
	overflowSetOff18  = unsafe.Offsetof(valueOverflow[[18]byte]{}.set)
	overflowSetOff50  = unsafe.Offsetof(valueOverflow[[50]byte]{}.set)
	overflowSetOff114 = unsafe.Offsetof(valueOverflow[[114]byte]{}.set)
	overflowSetOff242 = unsafe.Offsetof(valueOverflow[[242]byte]{}.set)
	overflowSetOff370 = unsafe.Offsetof(valueOverflow[[370]byte]{}.set)
	overflowSetOff498 = unsafe.Offsetof(valueOverflow[[498]byte]{}.set)
	overflowSetOffStr = unsafe.Offsetof(valueOverflow[string]{}.set)
)

// overflowKeyCap returns the size of the key area of a value overflow of a string map
// whose key remainder is klen bytes long, at most maxInlineOverflow (see setPrepend).
func overflowKeyCap(klen int) int {
	for _, c := range [...]int{18, 50, 114, 242, 370} {
		if klen <= c {
			return c
		}
	}
	return 498
}

// overflowSetOf returns the slot that holds the value set of value overflow l of a string
// map.
func overflowSetOf(l *singleKeyHead) **set3.Set3[string] {
	off := overflowSetOffStr
	switch k := l.rem(); {
	case k <= 18:
		off = overflowSetOff18
	case k <= 50:
		off = overflowSetOff50
	case k <= 114:
		off = overflowSetOff114
	case k <= 242:
		off = overflowSetOff242
	case k <= 370:
		off = overflowSetOff370
	case k <= maxInlineOverflow:
		off = overflowSetOff498
	}
	return (**set3.Set3[string])(unsafe.Add(unsafe.Pointer(l), off))
}

// newValueOverflow allocates a value overflow of a string map that holds key from base on,
// or the whole key as a string when the rest is too long to hold inline, and
// holds set as its value set.
func newValueOverflow(key []byte, base int, set *set3.Set3[string]) *singleKeyHead {
	var l *singleKeyHead
	if len(key) > maxKeyLen || len(key)-base > maxInlineOverflow {
		sl := &valueOverflow[string]{k: string(key)}
		sl.objType = kValueOverflow
		sl.setRem(longKey)
		l = &sl.singleKeyHead
	} else {
		s, kl := key[base:], len(key)
		switch k := len(s); {
		case k <= 18:
			l = newInlineOverflow[[18]byte](s, kl)
		case k <= 50:
			l = newInlineOverflow[[50]byte](s, kl)
		case k <= 114:
			l = newInlineOverflow[[114]byte](s, kl)
		case k <= 242:
			l = newInlineOverflow[[242]byte](s, kl)
		case k <= 370:
			l = newInlineOverflow[[370]byte](s, kl)
		default:
			l = newInlineOverflow[[498]byte](s, kl)
		}
	}
	*overflowSetOf(l) = set
	return l
}

// newInlineOverflow allocates a value overflow of a string map that holds the remainder s
// of a key of kl bytes inline in an array of type K.
func newInlineOverflow[K [18]byte | [50]byte | [114]byte | [242]byte | [370]byte | [498]byte](s []byte, kl int) *singleKeyHead {
	l := &valueOverflow[K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.objType = kValueOverflow
	l.setRem(len(s))
	l.kl = uint16(kl)
	return &l.singleKeyHead
}

// valueOverflowSize returns the size of the value overflow of a string map with a key
// remainder of klen bytes (longKey: a key held as a string).
func valueOverflowSize(klen int) uintptr {
	switch k := klen; {
	case k <= 18:
		return unsafe.Sizeof(valueOverflow[[18]byte]{})
	case k <= 50:
		return unsafe.Sizeof(valueOverflow[[50]byte]{})
	case k <= 114:
		return unsafe.Sizeof(valueOverflow[[114]byte]{})
	case k <= 242:
		return unsafe.Sizeof(valueOverflow[[242]byte]{})
	case k <= 370:
		return unsafe.Sizeof(valueOverflow[[370]byte]{})
	case k <= maxInlineOverflow:
		return unsafe.Sizeof(valueOverflow[[498]byte]{})
	}
	return unsafe.Sizeof(valueOverflow[string]{})
}

// overflowAdd adds v to the set of leaf l.
func overflowAdd(l *singleKeyHead, v string) { (*overflowSetOf(l)).Add(v) }

// overflowEach calls yield for every value of the set of leaf l and reports whether
// it ran to completion.
func overflowEach(l *singleKeyHead, yield func(string) bool) bool {
	for v := range (*overflowSetOf(l)).MutableRange() {
		if !yield(v) {
			return false
		}
	}
	return true
}
