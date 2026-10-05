package art

import (
	"reflect"
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map whose values are small and without pointers (Map.flat == 1, a uint64
// for instance) a key lives in a skpage.Fixed (docs/redesign/step3-fixed-design.md):
// the remainder and all the values as an array of T, in one object without pointers
// of 32 to 512 bytes. Like the page of the strings it starts like a leaf, so
// everything the tree does with the key works on it, and a key whose values do
// not fit it becomes a value overflow with a Set3 of its values (Set3[T] is the
// same object for every T).

// fixedType reports whether T takes skpage.Fixed pages without pointers: small,
// not empty and free of pointers, so that its values may live in memory the
// garbage collector does not scan.
func fixedType[T comparable]() bool {
	var z T
	return unsafe.Sizeof(z) > 0 && unsafe.Sizeof(z) <= 16 && pointerFree(reflect.TypeFor[T]())
}

func asFixed(l *singleKeyHead) *skpage.Fixed   { return (*skpage.Fixed)(unsafe.Pointer(l)) }
func fixedHead(p *skpage.Fixed) *singleKeyHead { return (*singleKeyHead)(unsafe.Pointer(p)) }

// newFixedLeaf allocates an empty page of values of T that holds key from base on,
// or a value overflow if the remainder or the key is too long for a page. It
// is a newLeafFunc; the first value is added by addFixed.
func newFixedLeaf[T comparable](key []byte, base int) *singleKeyHead {
	if len(key) > maxKeyLen || len(key)-base > skpage.MaxRemainderFixed[T]() {
		return newValueOverflow(key, base, set3.EmptyWithCapacity[T](8))
	}
	return fixedHead(skpage.EmptyFixed[T](key[base:], len(key)))
}

// addFixed adds v to the values of the page l of key, which sits in slot loc. A
// value that the page cannot hold makes the key a value overflow.
func addFixed[T comparable](loc **header, l *singleKeyHead, key []byte, v T) {
	p := asFixed(l)
	switch q, res := p.Add(v); {
	case res == skpage.Full:
		nl := newValueOverflow(key, l.base(), set3.EmptyWithCapacity[T](uint32(p.Len()+p.Len()/2+1)))
		fixedValuesTo[T](p, nl)
		overflowAdd(nl, v)
		*loc = singleKeyHdr(nl)
	case q != p:
		*loc = singleKeyHdr(fixedHead(q))
	}
}

// fixedValuesTo adds the values of page p to the set of the value overflow nl.
func fixedValuesTo[T comparable](p *skpage.Fixed, nl *singleKeyHead) {
	for _, x := range p.Values[T]() {
		overflowAdd(nl, x)
	}
}

// removeFixed removes v from the page l of key, which sits in the tree; rk is the
// map's rekeyFunc.
func removeFixed[T comparable](t *Tree, l *singleKeyHead, key []byte, v T, rk rekeyFunc) {
	switch q, ok := asFixed(l).Remove(v); {
	case !ok:
	case q == nil:
		t.remove(key, rk)
	case q != asFixed(l):
		*t.findSlot(key) = singleKeyHdr(fixedHead(q))
	}
}

// removeFromFixedOverflow removes v from the value overflow l of key in a map of
// fixed-size values, and moves the key into a page once its values fit one.
func removeFromFixedOverflow[T comparable](t *Tree, l *singleKeyHead, key []byte, v T, rk rekeyFunc) {
	s := *overflowSetOf[T](l)
	switch {
	case !s.Remove(v):
	case s.Size() == 0:
		t.remove(key, rk)
	default:
		if p := fixedFromOverflow[T](l); p != nil {
			*t.findSlot(key) = singleKeyHdr(fixedHead(p))
		}
	}
}

// fixedFromOverflow returns the page for the value overflow l, whose values have
// shrunk to what skpage.BackFits allows, or nil if they have not.
func fixedFromOverflow[T comparable](l *singleKeyHead) *skpage.Fixed {
	if l.rem() == longKey {
		return nil
	}
	var z T
	s := *overflowSetOf[T](l)
	if !skpage.BackFits(l.rem(), int(s.Size())*int(unsafe.Sizeof(z))) {
		return nil
	}
	return skpage.BuildFixed(l.stored(), int(l.kl), s.ToArray())
}

// rekeyFixed is the map's rekeyFunc for fixed-size values (see rekeyFunc): a page
// takes the bytes in front of its remainder, or becomes a value overflow if
// they do not fit; a value overflow takes them in its key area if they fit, else it
// is copied.
func rekeyFixed[T comparable](l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	if l.isValueOverflow() {
		if setPrepend(l, pre, b, pathLen, overflowKeyCap) {
			return l
		}
		return newValueOverflow(wholeKey(l, pre, b), pathLen, *overflowSetOf[T](l))
	}
	front := make([]byte, l.base()-pathLen)
	fillHead(front, pre, b, pathLen)
	p := asFixed(l)
	if q := p.Prepend[T](front); q != nil {
		return fixedHead(q)
	}
	nl := newValueOverflow(wholeKey(l, pre, b), pathLen, set3.EmptyWithCapacity[T](uint32(p.Len()+p.Len()/2+1)))
	fixedValuesTo[T](p, nl)
	return nl
}
