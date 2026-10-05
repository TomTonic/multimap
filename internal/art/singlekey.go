package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map of strings (Map.flat == 3) an entry lives in a single-key page
// (internal/skpage, docs/redesign/step3-skmv-design.md): the remainder and all
// the values as bytes, in one object without pointers of 32 to 512 bytes. The
// page starts like a leaf (type, remainder length, number of values, length of
// the whole key), so everything the tree does with the key of a leaf works on it.
// Its types are kValueOverflow+2 to kValueOverflow+12 in steps of two, those of the
// pages of fixed-size values (fixedkey.go) too; kValueOverflow is the value
// overflow, which takes an entry whose values do not fit a page or whose
// remainder is longer than a page holds.

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
// what skpage.BackFits allows, or nil if they have not.
func fromValueOverflow(l *singleKeyHead) *skpage.Page {
	if l.rem() == longKey {
		return nil
	}
	s := *overflowSetOf[string](l)
	if 2*int(s.Size()) > skpage.Room(l.rem()) { // every value takes a byte at least, and they may take half the room
		return nil
	}
	valueBytes, ok := 0, true
	vs := make([][]byte, 0, s.Size())
	overflowEach(l, func(x string) bool {
		valueBytes += 1 + len(x)
		if ok = len(x) <= skpage.MaxValue && skpage.BackFits(l.rem(), valueBytes); ok {
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
		return rekeyOverflow[string](l, pre, b, pathLen)
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

// The value overflow of a key whose values do not fit its page, in the maps
// that have single-key pages: behind the head and the key area it holds one
// pointer to a Set3 (github.com/TomTonic/Set3), a hash set that takes any T. The
// page is the small stage, so the inline stage and the array stage of vset, which
// a key with a few values would use, are not needed. The type of the values
// is a parameter of the code and of the pointer only: the object, its size and the
// offset of the pointer are the same for every T.
//
// It is an object of the grid of the pages, 32, 64, 128, 256, 384 or 512 bytes:
// the head takes 6 bytes, the pointer the last 8, and the key area what is left,
// 18, 50, 114, 242, 370 or 498 bytes, so the bytes the 64-byte vset.Set took go
// to the key and not to padding. A remainder of more than 498 bytes is held as
// a string.
type valueOverflow[K overflowKeyArea, T comparable] struct {
	singleKeyHead
	k   K
	set *set3.Set3[T]
}

// overflowKeyArea is the storage of the key of a value overflow: the key
// areas of the grid, or a string.
type overflowKeyArea interface {
	[18]byte | [50]byte | [114]byte | [242]byte | [370]byte | [498]byte | string
}

// maxInlineOverflow is the longest remainder a value overflow holds inline.
const maxInlineOverflow = 498

// Offsets of the pointer in the value overflow: they follow from the key areas
// of the grid and do not depend on the type of the values (T = int here).
var (
	overflowSetOff18  = unsafe.Offsetof(valueOverflow[[18]byte, int]{}.set)
	overflowSetOff50  = unsafe.Offsetof(valueOverflow[[50]byte, int]{}.set)
	overflowSetOff114 = unsafe.Offsetof(valueOverflow[[114]byte, int]{}.set)
	overflowSetOff242 = unsafe.Offsetof(valueOverflow[[242]byte, int]{}.set)
	overflowSetOff370 = unsafe.Offsetof(valueOverflow[[370]byte, int]{}.set)
	overflowSetOff498 = unsafe.Offsetof(valueOverflow[[498]byte, int]{}.set)
	overflowSetOffStr = unsafe.Offsetof(valueOverflow[string, int]{}.set)
)

// overflowKeyCap returns the size of the key area of a value overflow
// whose key remainder is klen bytes long, at most maxInlineOverflow (see setPrepend).
func overflowKeyCap(klen int) int {
	for _, c := range [...]int{18, 50, 114, 242, 370} {
		if klen <= c {
			return c
		}
	}
	return 498
}

// overflowSetOf returns the slot that holds the value set of value overflow l,
// whose values are of type T.
func overflowSetOf[T comparable](l *singleKeyHead) **set3.Set3[T] {
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
	return (**set3.Set3[T])(unsafe.Add(unsafe.Pointer(l), off))
}

// newValueOverflow allocates a value overflow that holds key from base on,
// or the whole key as a string when the rest is too long to hold inline, and
// holds set as its value set.
func newValueOverflow[T comparable](key []byte, base int, set *set3.Set3[T]) *singleKeyHead {
	var l *singleKeyHead
	if len(key) > maxKeyLen || len(key)-base > maxInlineOverflow {
		sl := &valueOverflow[string, T]{k: string(key)}
		sl.objType = kValueOverflow
		sl.setRem(longKey)
		l = &sl.singleKeyHead
	} else {
		s, kl := key[base:], len(key)
		switch k := len(s); {
		case k <= 18:
			l = newInlineOverflow[[18]byte, T](s, kl)
		case k <= 50:
			l = newInlineOverflow[[50]byte, T](s, kl)
		case k <= 114:
			l = newInlineOverflow[[114]byte, T](s, kl)
		case k <= 242:
			l = newInlineOverflow[[242]byte, T](s, kl)
		case k <= 370:
			l = newInlineOverflow[[370]byte, T](s, kl)
		default:
			l = newInlineOverflow[[498]byte, T](s, kl)
		}
	}
	*overflowSetOf[T](l) = set
	return l
}

// newInlineOverflow allocates a value overflow that holds the remainder s
// of a key of kl bytes inline in an array of type K.
func newInlineOverflow[K [18]byte | [50]byte | [114]byte | [242]byte | [370]byte | [498]byte, T comparable](s []byte, kl int) *singleKeyHead {
	l := &valueOverflow[K, T]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.objType = kValueOverflow
	l.setRem(len(s))
	l.kl = uint16(kl)
	return &l.singleKeyHead
}

// valueOverflowSize returns the size of the value overflow with a key
// remainder of klen bytes (longKey: a key held as a string); the same for every
// type of value.
func valueOverflowSize(klen int) uintptr {
	switch k := klen; {
	case k <= 18:
		return unsafe.Sizeof(valueOverflow[[18]byte, int]{})
	case k <= 50:
		return unsafe.Sizeof(valueOverflow[[50]byte, int]{})
	case k <= 114:
		return unsafe.Sizeof(valueOverflow[[114]byte, int]{})
	case k <= 242:
		return unsafe.Sizeof(valueOverflow[[242]byte, int]{})
	case k <= 370:
		return unsafe.Sizeof(valueOverflow[[370]byte, int]{})
	case k <= maxInlineOverflow:
		return unsafe.Sizeof(valueOverflow[[498]byte, int]{})
	}
	return unsafe.Sizeof(valueOverflow[string, int]{})
}

// overflowAdd adds v to the set of value overflow l.
func overflowAdd[T comparable](l *singleKeyHead, v T) { (*overflowSetOf[T](l)).Add(v) }

// overflowEach calls yield for every value of the set of value overflow l and
// reports whether it ran to completion.
func overflowEach[T comparable](l *singleKeyHead, yield func(T) bool) bool {
	for v := range (*overflowSetOf[T](l)).MutableRange() {
		if !yield(v) {
			return false
		}
	}
	return true
}

// newOverflowLeaf allocates an empty value overflow that holds key from base on:
// the leaf of an entry in a map whose type of value has no page, and for every
// type the home of an entry whose remainder is too long for one. It is a
// newLeafFunc.
func newOverflowLeaf[T comparable](key []byte, base int) *singleKeyHead {
	return newValueOverflow(key, base, set3.Empty[T]())
}

// rekeyOverflow is the map's rekeyFunc for a value overflow (see rekeyFunc): it
// takes the bytes in front of its remainder in its key area if they fit, else
// it is copied around the same set.
func rekeyOverflow[T comparable](l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	if setPrepend(l, pre, b, pathLen) {
		return l
	}
	return newValueOverflow(wholeKey(l, pre, b), pathLen, *overflowSetOf[T](l))
}
