package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/page"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// entry (a key with its values) lives in one single-key page of 32 to 512 bytes,
// the values right after the remainder of the key (see fixedkey.go, singlekey.go),
// or in a value overflow when they do not fit. Reading them costs no pointer chase
// beyond the page until an entry holds more values than a page of 512 bytes. The
// zero value is an empty map.
type Map[T comparable] struct {
	t    Tree
	flat int8 // 1: single-key pages of fixed-size values (page.Fixed), 3: single-key pages of strings (page.Str), -1: value overflows only, 0: not decided yet
	mk   bool // entries share multi-key pages (mkkey.go): strings and fixed-size values, with a pointer or none
	ptr  bool // the values are a word with a pointer: the multi-key pages of fixed-size values are typed objects (page.Fixed)
	cur  T    // the value of the Add in progress, for pair and reach of mkkey.go

	scrRests, scrVals [][]byte // scratch of pageOf, reused so that a burst allocates only its pages
	scrT              []T
}

// decide settles once per map how it holds the values of an entry: in
// single-key pages of fixed-size values if T takes them (page.Supported: small
// and without pointers, or one word that is a pointer), in single-key pages of
// strings if T is string, else in a value overflow for every entry.
func (m *Map[T]) decide() {
	switch {
	case page.Supported[T]():
		m.flat = 1
	case stringType[T]():
		m.flat = 3
	default:
		m.flat = -1
	}
	m.mk = m.flat == 3 || (m.flat == 1 && page.Supported[T]())
	m.ptr = m.flat == 1 && page.HoldsPointers[T]()
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
	m.cur = v
	m.upsert(key)
	var zero T
	m.cur = zero
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
	n, pathLen := m.t.find(key)
	if n == nil {
		return 0
	}
	if isMultiKey(n.objType) {
		return m.pageRemove(n, pathLen, key, v, false)
	}
	l := asSingleKey(n)
	if !l.isValueOverflow() {
		m.removeFromLeaf(l, key, pathLen, v)
		return 0
	}
	s := *overflowSetOf[T](l)
	switch {
	case !s.Remove(v):
	case s.Size() == 0:
		m.t.remove(key, m.rekey)
	case m.flat != -1: // the values may fit a page again
		if p := m.backToPage(l); p != nil {
			*m.t.findSlot(key) = singleKeyHdr(p)
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

// eachValue calls yield for every value of leaf l, a page or a value overflow, and reports whether it ran to
// completion.
func eachValue[T comparable](l *singleKeyHead, flat int8, yield func(T) bool) bool {
	switch {
	case l.isValueOverflow():
		return overflowEach(l, yield)
	case flat == 3: // T is string
		y := *(*func(string) bool)(unsafe.Pointer(&yield))
		return asSK(l).EachSingle(func(val []byte) bool { return y(string(val)) })
	}
	return asFixed(l).EachSingle(yield)
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
	var c Cursor[T]
	c.Init(m, b, true)
	for c.NextPage() {
		for _, k := range c.Keys {
			if !fn(k) {
				return
			}
		}
	}
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	var c Cursor[T]
	c.Init(m, b, false)
	for c.NextPage() {
		for _, v := range c.Vals {
			if !yield(v) {
				return
			}
		}
		if c.Set != nil { // a value overflow: its set, not copied
			for v := range c.Set.MutableRange() {
				if !yield(v) {
					return
				}
			}
		}
	}
}
