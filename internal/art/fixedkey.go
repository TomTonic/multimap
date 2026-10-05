package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// In a map whose values are small and without pointers (a uint64) or one word with a
// pointer (a *T), Map.flat == 1, an entry lives in a skpage.Fixed (docs/redesign/step3-fixed-design.md):
// the remainder and all the values as an array of T, in one object without pointers
// of 32 to 512 bytes (with pointers: a Go type that marks them). Like the page of
// the strings it starts like a leaf, so everything the tree does with the key
// works on it, and an entry whose values do not fit it becomes a value overflow with a Set3 of its values (Set3[T] is the
// same object for every T).

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
// value that the page cannot hold makes the entry a value overflow.
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
		return rekeyOverflow[T](l, pre, b, pathLen)
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
