package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/mkpage"
	"github.com/TomTonic/multimap/internal/skpage"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// entry (a key with its values) lives in one single-key page of 32 to 512 bytes,
// the values right after the remainder of the key (see fixedkey.go, singlekey.go),
// or in a value overflow when they do not fit. Reading them costs no pointer chase
// beyond the page until an entry holds more values than a page of 512 bytes. The
// zero value is an empty map.
type Map[T comparable] struct {
	t    Tree
	flat int8 // 1: single-key pages of fixed-size values (skpage.Fixed), 3: single-key pages of strings (skpage.Page), -1: value overflows only, 0: not decided yet
	mk   bool // entries share multi-key pages (mkkey.go): strings and fixed-size values, with a pointer or none
	ptr  bool // the values are a word with a pointer: the multi-key pages of fixed-size values are typed objects (mkpage.Fixed)
	cur  T    // the value of the Add in progress, for the pager methods of mkkey.go

	scrRests, scrVals [][]byte // scratch of pageOf, reused so that a burst allocates only its pages
	scrT              []T
}

// setPrepend makes value overflow l hold its key from pathLen on in place when the
// longer remainder still fits the key area of l's class, which keeps its value
// set where it is. It saves rekey the allocation of a new object, and reports
// whether it did.
func setPrepend(l *singleKeyHead, pre []byte, b, pathLen int) bool {
	old, klen := l.rem(), l.keyLen()-pathLen // a value overflow that holds its whole key as a string never gets here
	if klen > overflowKeyCap(old) {
		return false
	}
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), klen)
	copy(area[klen-old:], area[:old])
	fillHead(area[:klen-old], pre, b, pathLen)
	l.setRem(klen)
	return true
}

// decide settles once per map how it holds the values of an entry: in
// single-key pages of fixed-size values if T takes them (skpage.Supported: small
// and without pointers, or one word that is a pointer), in single-key pages of
// strings if T is string, else in a value overflow for every entry.
func (m *Map[T]) decide() {
	switch {
	case skpage.Supported[T]():
		m.flat = 1
	case stringType[T]():
		m.flat = 3
	default:
		m.flat = -1
	}
	m.mk = m.flat == 3 || (m.flat == 1 && mkpage.Supported[T]())
	m.ptr = m.flat == 1 && mkpage.HoldsPointers[T]()
}

// Len returns the number of keys.
func (m *Map[T]) Len() int { return m.t.Len() }

// Clear removes all keys and values.
func (m *Map[T]) Clear() { m.t.Clear() }

// Add adds v to the values of key, creating the key if needed. The map keeps
// its own copy of key.
func (m *Map[T]) Add(key []byte, v T) {
	if m.flat == 0 {
		m.decide()
	}
	// Chosen here rather than returned from a helper: a function value that
	// does not escape stays on the stack, one that is returned is allocated.
	var nl newLeafFunc = newOverflowLeaf[T]
	switch m.flat {
	case 1:
		nl = newFixedLeaf[T]
	case 3:
		nl = newSK
	}
	if m.mk {
		m.cur = v
	}
	loc := m.t.upsert(key, nl, m)
	if m.mk {
		var zero T
		m.cur = zero
	}
	if loc != nil {
		m.addToLeaf(loc, asSingleKey(*loc), key, v)
	}
}

// addToLeaf adds v to the values of leaf l of key, which sits in slot loc.
func (m *Map[T]) addToLeaf(loc **header, l *singleKeyHead, key []byte, v T) {
	switch {
	case l.isValueOverflow():
		overflowAdd(l, v)
	case m.flat == 3:
		addSK(loc, l, key, *(*string)(unsafe.Pointer(&v)))
	default:
		addFixed(loc, l, key, v)
	}
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
//
// It finds the leaf as a lookup does and looks for the leaf's slot only when
// the leaf has to move into a smaller one.
func (m *Map[T]) Remove(key []byte, v T) {
	before := m.t.size
	left := m.removeValue(key, v)
	if m.mk && m.t.size != before && left <= mergeBelow {
		m.mergeUp(&m.t.root, key, 0)
	}
}

// removeValue is Remove without the merge of the nodes above a removed key. It returns the
// number of entries left in the multi-key page that held the entry, and 0 if it did not.
func (m *Map[T]) removeValue(key []byte, v T) int {
	// Chosen here, as in Add: a function value that does not escape stays on
	// the stack.
	rk := m.rekey
	n, pathLen := m.t.find(key)
	if n == nil {
		return 0
	}
	if isMultiKey(n.objType) {
		return m.pageRemove(n, pathLen, key, v, false)
	}
	l := asSingleKey(n)
	if !l.isValueOverflow() {
		if m.flat == 1 {
			removeFixed(&m.t, l, key, v, rk)
		} else {
			m.t.removeSK(l, key, *(*string)(unsafe.Pointer(&v)), rk)
		}
		return 0
	}
	s := *overflowSetOf[T](l)
	switch {
	case !s.Remove(v):
	case s.Size() == 0:
		m.t.remove(key, rk)
	case m.flat == 1: // the values may fit a page again
		if p := fixedFromOverflow[T](l); p != nil {
			*m.t.findSlot(key) = singleKeyHdr(fixedHead(p))
		}
	case m.flat == 3:
		if p := fromValueOverflow(l); p != nil {
			*m.t.findSlot(key) = singleKeyHdr(skHead(p))
		}
	}
	return 0
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) {
	n, pathLen := m.t.find(key)
	left := 0
	switch {
	case n == nil:
		return
	case isMultiKey(n.objType):
		var zero T
		left = m.pageRemove(n, pathLen, key, zero, true)
	default:
		m.t.remove(key, m.rekey)
	}
	if m.mk && left <= mergeBelow {
		m.mergeUp(&m.t.root, key, 0)
	}
}

// rekey is the map's rekeyFunc: the one for its type of leaf, which also moves a
// multi-key page up.
func (m *Map[T]) rekey(l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	switch {
	case isMultiKey(l.objType):
		return m.rekeyPage(l, pre, b, pathLen)
	case m.flat == 1:
		return rekeyFixed[T](l, pre, b, pathLen)
	case m.flat == 3:
		return rekeySK(l, pre, b, pathLen)
	}
	return rekeyOverflow[T](l, pre, b, pathLen)
}

// Has reports whether key holds any values.
func (m *Map[T]) Has(key []byte) bool {
	n, pathLen := m.t.find(key)
	if n != nil && isMultiKey(n.objType) {
		return m.pageHas(n, pathLen, key)
	}
	return n != nil
}

// Each calls yield for every value of key, in unspecified order, until it
// returns false. yield must not modify the map.
func (m *Map[T]) Each(key []byte, yield func(T) bool) {
	n, pathLen := m.t.find(key)
	switch {
	case n == nil:
	case isMultiKey(n.objType):
		m.pageEach(n, pathLen, key, yield)
	default:
		eachValue(asSingleKey(n), m.flat, yield)
	}
}

// eachValue calls yield for every value of leaf l and reports whether it ran
// to completion.
func eachValue[T comparable](l *singleKeyHead, flat int8, yield func(T) bool) bool {
	switch {
	case l.isValueOverflow():
		return overflowEach(l, yield)
	case flat == 3: // T is string
		return asSK(l).Strings(*(*func(string) bool)(unsafe.Pointer(&yield)))
	}
	return asFixed(l).Each(yield)
}

// leafTail is the offset of the last byte of the smallest object of any map, a
// page or a value overflow of 32 bytes, which the scan touches ahead (see
// touchChildren). Every object is at least that large, and a constant offset
// keeps the load independent of the object's head.
const leafTail = 32 - 1

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key is assembled for fn, from the path to its leaf and
// the rest the leaf holds: fn must not modify or retain it, and must not
// modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte) bool) {
	var kb keyBuf
	m.t.scan(b, leafTail, &kb, func(*singleKeyHead) bool { return fn(kb.key) },
		func(p *header, pathLen int, lo, hi bool) bool { return m.scanPage(p, pathLen, b, lo, hi, &kb, fn, nil) })
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, leafTail, nil, func(l *singleKeyHead) bool { return eachValue(l, m.flat, yield) },
		func(p *header, pathLen int, lo, hi bool) bool {
			return m.scanPage(p, pathLen, b, lo, hi, nil, nil, yield)
		})
}
