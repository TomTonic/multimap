package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
	"github.com/TomTonic/multimap/internal/vset"
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
}

// newSetLeaf allocates a value overflow that holds key from base on, in the
// smallest size class that fits it, or the whole key as a string when the
// rest is too long to hold inline. It captures nothing, so passing it as a
// newLeafFunc allocates no closure.
func newSetLeaf[T comparable](key []byte, base int) *singleKeyHead {
	if len(key) > maxKeyLen || len(key)-base > maxInline {
		l := &leaf[T, string]{k: string(key)}
		l.objType = kValueOverflow
		l.setRem(longKey)
		return &l.singleKeyHead
	}
	return newSetLeafOf[T](key[base:], len(key))
}

// newSetLeafOf allocates a value overflow holding the key remainder s, at most
// maxInline bytes, of a key of kl bytes.
func newSetLeafOf[T comparable](s []byte, kl int) *singleKeyHead {
	switch n := len(s); {
	case n <= 16:
		return newInline[T, [16]byte](s, kl)
	case n <= 32:
		return newInline[T, [32]byte](s, kl)
	case n <= 48:
		return newInline[T, [48]byte](s, kl)
	case n <= 64:
		return newInline[T, [64]byte](s, kl)
	case n <= 96:
		return newInline[T, [96]byte](s, kl)
	case n <= 128:
		return newInline[T, [128]byte](s, kl)
	case n <= 192:
		return newInline[T, [192]byte](s, kl)
	}
	return newInline[T, [256]byte](s, kl)
}

// newInline allocates a value overflow that holds the key remainder s inline in an
// array of type K.
func newInline[T comparable, K [16]byte | [32]byte | [48]byte | [64]byte | [96]byte | [128]byte | [192]byte | [256]byte](s []byte, kl int) *singleKeyHead {
	l := &leaf[T, K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.objType = kValueOverflow
	l.setRem(len(s))
	l.kl = uint16(kl)
	return &l.singleKeyHead
}

// Offsets of the value set in the value overflow of each key size class. They do
// not depend on T: a vset.Set holds a pointer, so it is always 8-aligned.
var (
	setOff16  = unsafe.Offsetof(leaf[struct{}, [16]byte]{}.vals)
	setOff32  = unsafe.Offsetof(leaf[struct{}, [32]byte]{}.vals)
	setOff48  = unsafe.Offsetof(leaf[struct{}, [48]byte]{}.vals)
	setOff64  = unsafe.Offsetof(leaf[struct{}, [64]byte]{}.vals)
	setOff96  = unsafe.Offsetof(leaf[struct{}, [96]byte]{}.vals)
	setOff128 = unsafe.Offsetof(leaf[struct{}, [128]byte]{}.vals)
	setOff192 = unsafe.Offsetof(leaf[struct{}, [192]byte]{}.vals)
	setOff256 = unsafe.Offsetof(leaf[struct{}, [256]byte]{}.vals)
	setOffStr = unsafe.Offsetof(leaf[struct{}, string]{}.vals)
)

// valsOff returns the offset of the value set in value overflow l: it depends on the
// size class of the key area only.
func valsOff(l *singleKeyHead) uintptr {
	off := setOffStr
	switch k := l.rem(); {
	case k <= 16:
		off = setOff16
	case k <= 32:
		off = setOff32
	case k <= 48:
		off = setOff48
	case k <= 64:
		off = setOff64
	case k <= 96:
		off = setOff96
	case k <= 128:
		off = setOff128
	case k <= 192:
		off = setOff192
	case k <= maxInline:
		off = setOff256
	}
	return off
}

// vals returns the value set of a value overflow, created by newSetLeaf[T].
func vals[T comparable](l *singleKeyHead) *vset.Set[T] {
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), valsOff(l)))
}

// rekey returns a leaf with l's values that holds its key from pathLen on (see
// rekeyFunc). It captures nothing, so passing it allocates no closure. It
// keeps l where it is when the longer remainder fits its size class, and
// builds the whole key only when it has to copy l.
func rekey[T comparable](l *singleKeyHead, pre []byte, b, pathLen int) *singleKeyHead {
	if setPrepend(l, pre, b, pathLen, setKeyCap) {
		return l
	}
	nl := newSetLeaf[T](wholeKey(l, pre, b), pathLen)
	*vals[T](nl) = *vals[T](l)
	return nl
}

// setKeyCap returns the size of the key area of a value overflow whose key
// remainder is klen bytes long, at most maxInline: the class that klen falls
// in (see vals); the largest holds 256 bytes, of which a remainder uses
// maxInline.
func setKeyCap(klen int) int {
	for _, c := range [...]int{16, 32, 48, 64, 96, 128, 192} {
		if klen <= c {
			return c
		}
	}
	return maxInline
}

// setPrepend makes value overflow l hold its key from pathLen on in place when the
// longer remainder still fits the key area of l's class, which keeps its value
// set where it is. It saves rekey the allocation of a new leaf, see
// flatPrepend, and reports whether it did. keyCap is setKeyCap, or overflowKeyCap
// for the value overflow of a string map.
func setPrepend(l *singleKeyHead, pre []byte, b, pathLen int, keyCap func(int) int) bool {
	old, klen := l.rem(), l.keyLen()-pathLen // a string leaf holds its whole key, and never gets here
	if klen > keyCap(old) {
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
	var nl newLeafFunc = newSetLeaf[T]
	switch m.flat {
	case 1:
		nl = newFixedLeaf[T]
	case 3:
		nl = newSK
	}
	loc := m.t.upsert(key, nl)
	m.addToLeaf(loc, asSingleKey(*loc), key, v)
}

// addToLeaf adds v to the values of leaf l of key, which sits in slot loc.
func (m *Map[T]) addToLeaf(loc **header, l *singleKeyHead, key []byte, v T) {
	switch {
	case l.isValueOverflow() && m.flat == 3:
		overflowAdd(l, *(*string)(unsafe.Pointer(&v)))
	case l.isValueOverflow() && m.flat == 1:
		overflowAdd(l, v)
	case l.isValueOverflow():
		vals[T](l).Add(v)
	case m.flat == 3:
		addSK(loc, l, key, *(*string)(unsafe.Pointer(&v)))
	case m.flat == 1:
		addFixed(loc, l, key, v)
	}
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
//
// It finds the leaf as a lookup does and looks for the leaf's slot only when
// the leaf has to move into a smaller one.
func (m *Map[T]) Remove(key []byte, v T) {
	// Chosen here, as in Add: a function value that does not escape stays on
	// the stack.
	rk := m.rekeyFunc()
	l := m.t.find(key)
	if l == nil {
		return
	}
	if !l.isValueOverflow() {
		if m.flat == 1 {
			removeFixed(&m.t, l, key, v, rk)
		} else {
			m.t.removeSK(l, key, *(*string)(unsafe.Pointer(&v)), rk)
		}
		return
	}
	switch m.flat {
	case 1:
		removeFromFixedOverflow(&m.t, l, key, v, rk)
		return
	case 3:
		m.removeFromOverflowSet(l, key, *(*string)(unsafe.Pointer(&v)), rk)
		return
	}
	s := vals[T](l)
	switch {
	case !s.Remove(v):
	case s.Len() == 0:
		m.t.remove(key, rk)
	}
}

// removeFromOverflowSet removes v from the value overflow l of key in a map of strings, and
// moves the key into a page once its values fit one.
func (m *Map[T]) removeFromOverflowSet(l *singleKeyHead, key []byte, v string, rk rekeyFunc) {
	s := *overflowSetOf[string](l)
	switch {
	case !s.Remove(v):
	case s.Size() == 0:
		m.t.remove(key, rk)
	default:
		if p := fromValueOverflow(l); p != nil {
			*m.t.findSlot(key) = singleKeyHdr(skHead(p))
		}
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) {
	m.t.remove(key, m.rekeyFunc())
}

// rekeyFunc returns the map's rekeyFunc: the one for its type of leaf.
func (m *Map[T]) rekeyFunc() rekeyFunc {
	switch m.flat {
	case 1:
		return rekeyFixed[T]
	case 3:
		return rekeySK
	}
	return rekey[T]
}

// Has reports whether key holds any values.
func (m *Map[T]) Has(key []byte) bool {
	return m.t.find(key) != nil
}

// Each calls yield for every value of key, in unspecified order, until it
// returns false. yield must not modify the map.
func (m *Map[T]) Each(key []byte, yield func(T) bool) {
	if l := m.t.find(key); l != nil {
		eachValue(l, m.flat, yield)
	}
}

// eachValue calls yield for every value of leaf l and reports whether it ran
// to completion.
func eachValue[T comparable](l *singleKeyHead, flat int8, yield func(T) bool) bool {
	switch flat {
	case 3: // T is string
		y := *(*func(string) bool)(unsafe.Pointer(&yield))
		if l.isValueOverflow() {
			return overflowEach(l, y)
		}
		return asSK(l).Strings(y)
	case 1:
		if l.isValueOverflow() {
			return overflowEach(l, yield)
		}
		return asFixed(l).Each(yield)
	}
	return vals[T](l).Each(yield)
}

// leafTail is the offset of the last byte of the smallest leaf this map
// creates, which the scan touches ahead (see touchChildren). Every leaf is
// at least that large, and a constant offset keeps the load independent of
// the leaf's head.
func (m *Map[T]) leafTail() uintptr {
	if m.flat != -1 {
		return 32 - 1 // the smallest page
	}
	return unsafe.Sizeof(leaf[T, [16]byte]{}) - 1
}

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key is assembled for fn, from the path to its leaf and
// the rest the leaf holds: fn must not modify or retain it, and must not
// modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte) bool) {
	var kb keyBuf
	m.t.scan(b, m.leafTail(), &kb, func(*singleKeyHead) bool { return fn(kb.key) })
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, m.leafTail(), nil, func(l *singleKeyHead) bool { return eachValue(l, m.flat, yield) })
}
