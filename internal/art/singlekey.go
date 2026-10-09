package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/page"
	"github.com/TomTonic/multimap/internal/swar"
)

// An entry with one key lives in a leaf: a single-key page, which is the one-key form of the one page
// (internal/page, docs/redesign/step5-one-page.md): the remainder of the key (its key part) and all its values in one
// object of 32 to 512 bytes, for values of variable length (strings, Map.flat == 3) and of one size (a uint64 or
// a *T, Map.flat == 1) alike; or a value overflow, which takes an entry whose values do not fit a page or whose
// remainder is longer than a page holds, and which is the only leaf of the maps whose values have no page
// (Map.flat == -1). Both start with the head of the page (type, length of the key part, number of values,
// rawWords) and hold their key from the path length they stand at on: when the path changes the leaf changes with it,
// in place (Skip, Prepend). The types of the pages are kValueOverflow+2 to kValueOverflow+12 in steps of two; the
// type of the value overflow is kValueOverflow.

// stringType reports whether T is string, the value type that takes the pages of variable length.
func stringType[T comparable]() bool {
	var z T
	_, ok := any(z).(string)
	return ok
}

func asSK(l *singleKeyHead) *page.Str                { return (*page.Str)(unsafe.Pointer(l)) }
func skHead(p *page.Str) *singleKeyHead              { return (*singleKeyHead)(unsafe.Pointer(p)) }
func asFixed(l *singleKeyHead) *page.Fixed           { return (*page.Fixed)(unsafe.Pointer(l)) }
func fixedHead(p *page.Fixed) *singleKeyHead         { return (*singleKeyHead)(unsafe.Pointer(p)) }
func pageHdr[P *page.Str | *page.Fixed](p P) *header { return (*header)(unsafe.Pointer(p)) }

// view returns the bytes of s without copying them. The page copies what it stores.
func view(s string) []byte { return unsafe.Slice(unsafe.StringData(s), len(s)) }

// The value overflow of a key whose values do not fit its page: behind the head and the key area it holds
// one pointer to a Set3 (github.com/TomTonic/Set3), a hash set that takes any T. The page is the small stage,
// so the inline stage and the array stage of vset, which a key with a few values would use, are not needed. The type of
// the values is a parameter of the code and of the pointer only: the object, its size and the offset of the pointer
// are the same for every T.
//
// It is an object of the grid of the pages, 32, 64, 128, 256, 384 or 512 bytes: the head takes 4 bytes, the pointer the last
// 8, and the key area what is left, 20, 52, 116, 244, 372 or 500 bytes. A remainder of more than 500 bytes is held as
// a string.
type valueOverflow[K overflowKeyArea, T comparable] struct {
	singleKeyHead
	k   K
	set *set3.Set3[T]
}

// overflowKeyArea is the storage of the key of a value overflow: the key areas of the grid, or a string.
type overflowKeyArea interface {
	[20]byte | [52]byte | [116]byte | [244]byte | [372]byte | [500]byte | string
}

// maxInlineOverflow is the longest remainder a value overflow holds inline.
const maxInlineOverflow = 500

// overflowSetOf returns the slot that holds the value set of value overflow l, whose values are of type T. Its
// place is fixed by the object's size and recorded in rawWords (the head's last byte): the key area of the object
// does not change when the key part shrinks (Skip), the offset of the pointer must not.
func overflowSetOf[T comparable](l *singleKeyHead) **set3.Set3[T] {
	return (**set3.Set3[T])(unsafe.Add(unsafe.Pointer(l), uintptr(l.raw)*8))
}

// newValueOverflow allocates a value overflow that holds the key part rest, or the key part as a string when it is
// too long to hold inline, and holds set as its value set.
func newValueOverflow[T comparable](rest []byte, set *set3.Set3[T]) *singleKeyHead {
	var l *singleKeyHead
	if len(rest) > maxInlineOverflow {
		sl := &valueOverflow[string, T]{k: string(rest)}
		sl.objType = kValueOverflow
		sl.setRem(longKey)
		sl.raw = uint8((uintptr(unsafe.Pointer(&sl.set)) - uintptr(unsafe.Pointer(sl))) / 8)
		l = &sl.singleKeyHead
	} else {
		switch k := len(rest); {
		case k <= 20:
			l = newInlineOverflow[[20]byte, T](rest)
		case k <= 52:
			l = newInlineOverflow[[52]byte, T](rest)
		case k <= 116:
			l = newInlineOverflow[[116]byte, T](rest)
		case k <= 244:
			l = newInlineOverflow[[244]byte, T](rest)
		case k <= 372:
			l = newInlineOverflow[[372]byte, T](rest)
		default:
			l = newInlineOverflow[[500]byte, T](rest)
		}
	}
	*overflowSetOf[T](l) = set
	return l
}

// newInlineOverflow allocates a value overflow that holds the key part s inline in an array of type K.
func newInlineOverflow[K [20]byte | [52]byte | [116]byte | [244]byte | [372]byte | [500]byte, T comparable](s []byte) *singleKeyHead {
	l := &valueOverflow[K, T]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.objType = kValueOverflow
	l.setRem(len(s))
	l.raw = uint8((uintptr(unsafe.Pointer(&l.set)) - uintptr(unsafe.Pointer(l))) / 8)
	return &l.singleKeyHead
}

// valueOverflowSize returns the size of value overflow l: the pointer to the value set is its last word.
func valueOverflowSize(l *singleKeyHead) uintptr { return uintptr(l.raw)*8 + 8 }

// overflowAdd adds v to the set of value overflow l.
func overflowAdd[T comparable](l *singleKeyHead, v T) { (*overflowSetOf[T](l)).Add(v) }

// overflowEach calls yield for every value of the set of value overflow l and reports whether it ran to completion.
func overflowEach[T comparable](l *singleKeyHead, yield func(T) bool) bool {
	for v := range (*overflowSetOf[T](l)).MutableRange() {
		if !yield(v) {
			return false
		}
	}
	return true
}

// overflowSkip drops the first k bytes of the key part of value overflow l (k at most its length): a node has been put
// above it. A key held as a string stays a string, with a shorter one.
func overflowSkip(l *singleKeyHead, k int) {
	if l.rem() == longKey {
		s := (*string)(unsafe.Add(unsafe.Pointer(l), strOff))
		*s = (*s)[k:]
		return
	}
	n := l.rem()
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), n)
	copy(area, area[k:])
	clear(area[n-k:])
	l.setRem(n - k)
}

// overflowPrepend puts front before the key part of value overflow l in place if the longer key part still fits
// the key area of l's size, which keeps its value set where it is, and reports whether it did. A key held as a
// string always takes it.
func overflowPrepend(l *singleKeyHead, front []byte) bool {
	if l.rem() == longKey {
		s := (*string)(unsafe.Add(unsafe.Pointer(l), strOff))
		*s = string(front) + *s
		return true
	}
	n := l.rem()
	if n+len(front) > int(l.raw)*8-int(keyOff) {
		return false
	}
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), n+len(front))
	copy(area[len(front):], area[:n])
	copy(area, front)
	l.setRem(n + len(front))
	return true
}

// rekeyOverflow is the map's rekeyFunc for a value overflow (see rekeyFunc): it takes the bytes in front of its
// key part in its key area if they fit, else it is copied around the same set.
func rekeyOverflow[T comparable](l *singleKeyHead, front []byte) *singleKeyHead {
	if overflowPrepend(l, front) {
		return l
	}
	return newValueOverflow(append(front, l.stored()...), *overflowSetOf[T](l))
}

// newLeaf makes the leaf of the entry rest -> m.cur, the key part rest being the key from the path length of the
// leaf on: a page with its first value, or a value overflow if the page cannot hold them.
func (m *Map[T]) newLeaf(rest []byte) *singleKeyHead {
	switch m.flat {
	case 3:
		if p := page.NewStr(rest, view(strOf(m.cur))); p != nil {
			return skHead(p)
		}
	case 1:
		if p := page.NewFixed(rest, m.cur, m.ptr); p != nil {
			return fixedHead(p)
		}
	}
	l := newValueOverflow(rest, set3.EmptyWithCapacity[T](8))
	overflowAdd(l, m.cur)
	return l
}

// reachLeaf is the add for a descent that ended at leaf l (a single-key page or a value overflow) at pathLen:
// the value m.cur goes to the key if l is its leaf, a page makes a multi-key page with it if another key meets
// it and both fit one, else a node goes above l and the new key gets a leaf of its own below it.
func (m *Map[T]) reachLeaf(loc **header, l *singleKeyHead, key []byte, pathLen int) {
	rest := key[pathLen:]
	if l.isValueOverflow() {
		if l.matches(rest) {
			overflowAdd(l, m.cur)
			return
		}
		m.splitLeaf(loc, l, key, pathLen)
		return
	}
	if m.maxKeys == 1 && !l.matches(rest) { // every key its own page: another key gets a page below a node
		m.splitLeaf(loc, l, key, pathLen)
		return
	}
	var q *header
	var res page.Result
	if m.flat == 3 {
		r, x := asSK(l).Add(rest, view(strOf(m.cur)))
		q, res = pageHdr(r), x
	} else {
		r, x := asFixed(l).Add(rest, m.cur, m.ptr)
		q, res = pageHdr(r), x
	}
	switch res {
	case page.Added: // another key: the two make a multi-key page
		ev(evPair, 2)
		*loc = q
		m.t.size++
	case page.AddedValue, page.Present:
		*loc = q
	default: // page.Full
		if l.matches(rest) { // the values of the key outgrow the page
			*loc = singleKeyHdr(m.toValueOverflow(l, rest))
			return
		}
		ev(evPairNo, 1)
		m.splitLeaf(loc, l, key, pathLen)
	}
}

// toValueOverflow returns a value overflow with the values of page l and m.cur, for the key part rest: the value
// overflow of a key that outgrew its page.
func (m *Map[T]) toValueOverflow(l *singleKeyHead, rest []byte) *singleKeyHead {
	n := int(l.n)
	nl := newValueOverflow(rest, set3.EmptyWithCapacity[T](uint32(n+n/2+1)))
	eachValue(l, m.flat, func(x T) bool { overflowAdd(nl, x); return true })
	overflowAdd(nl, m.cur)
	return nl
}

// splitLeaf handles an add whose key meets leaf l at pathLen and is not its key: both keys go below a new node
// holding their common prefix. l loses the bytes the node takes (Skip), the new key gets a leaf of its own.
func (m *Map[T]) splitLeaf(loc **header, l *singleKeyHead, key []byte, pathLen int) {
	ls, rest := l.stored(), key[pathLen:]
	p := swar.Lcp(ls, rest)
	nn := newNode(kN5, p)
	storePrefix(nn, rest[:p])
	var first [1]byte // the byte l hangs under, taken before the page loses its bytes
	n := min(1, len(ls)-p)
	if n > 0 {
		first[0] = ls[p]
	}
	m.skipLeaf(l, p+n)
	ev(evSplitLeaf, 1)
	h, _ := attachAt(nn, first[:n], l)
	h, _ = attachAt(h, rest[p:], m.newLeaf(rest[p+min(1, len(rest)-p):]))
	*loc = h
	m.t.size++
}

// skipLeaf drops the first k bytes of the key part of leaf l in place: a node has been put above it.
func (m *Map[T]) skipLeaf(l *singleKeyHead, k int) {
	switch {
	case l.isValueOverflow():
		overflowSkip(l, k)
	case m.flat == 3:
		asSK(l).Skip(k)
	default:
		asFixed(l).Skip(k)
	}
}

// rekey is the map's rekeyFunc (see rekeyFunc): a leaf takes the bytes in front of its key part, a page that no longer
// holds them becomes a value overflow; a value overflow takes them in its key area if they fit, else it is copied.
// A multi-key page that cannot take them answers nil.
func (m *Map[T]) rekey(l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	ev(evRekeyLeaf, 1)
	front := frontOf(pre, b, pathLen)
	switch {
	case isMultiKey(l.objType):
		return m.rekeyPage(l, front)
	case l.isValueOverflow():
		return rekeyOverflow[T](l, front)
	case m.flat == 3:
		if q := asSK(l).Prepend(front); q != nil {
			return skHead(q)
		}
	default:
		if q := asFixed(l).Prepend[T](front, m.ptr); q != nil {
			return fixedHead(q)
		}
	}
	n := int(l.n)
	nl := newValueOverflow(append(front, l.stored()...), set3.EmptyWithCapacity[T](uint32(n+n/2+1)))
	eachValue(l, m.flat, func(x T) bool { overflowAdd(nl, x); return true })
	return nl
}

// removeFromLeaf removes v from the values of page l of key, which sits in the tree, and the key with it when it
// was the last.
func (m *Map[T]) removeFromLeaf(l *singleKeyHead, key []byte, pathLen int, v T) {
	rest := key[pathLen:]
	var q *header
	var rm page.Removal
	if m.flat == 3 {
		r, x := asSK(l).Remove(rest, view(strOf(v)))
		q, rm = pageHdr(r), x
	} else {
		r, x := asFixed(l).Remove(rest, v, m.ptr)
		q, rm = pageHdr(r), x
	}
	switch {
	case rm == page.Gone:
		m.t.remove(key, m.rekey)
	case rm == page.Removed && q != singleKeyHdr(l):
		*m.t.findSlot(key) = q
	}
}

// backToPage returns the page for value overflow l, whose values have shrunk to what page.BackFits allows, or nil
// if they have not.
func (m *Map[T]) backToPage(l *singleKeyHead) *singleKeyHead {
	if l.rem() == longKey {
		return nil
	}
	s := *overflowSetOf[T](l)
	n := int(s.Size())
	// the test of the size comes first and allocates nothing: it runs at every removal from a value overflow
	if m.flat == 1 {
		var z T
		if !page.BackFits(l.rem(), n*int(unsafe.Sizeof(z))) {
			return nil
		}
		return fixedHead(page.BuildFixedOf(restsOf(l, n), s.ToArray(), m.ptr)) // a nil page is a nil head
	}
	if 2*n > page.Room(l.rem()) { // every value takes a byte at least, and they may take half the room
		return nil
	}
	valueBytes, ok := 0, true
	overflowEach(l, func(x string) bool {
		valueBytes += 1 + len(x)
		ok = len(x) <= page.MaxValue && page.BackFits(l.rem(), valueBytes)
		return ok
	})
	if !ok {
		return nil
	}
	vs := make([][]byte, 0, n)
	overflowEach(l, func(x string) bool { vs = append(vs, view(x)); return true })
	return skHead(page.BuildStrings(restsOf(l, n), vs))
}

// restsOf returns n times the key part of value overflow l: the keys of its n values, as the page builders take them.
func restsOf(l *singleKeyHead, n int) [][]byte {
	rests := make([][]byte, n)
	for i := range rests {
		rests[i] = l.stored()
	}
	return rests
}
