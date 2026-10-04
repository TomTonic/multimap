package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// leaf holds its key's values: a flat leaf right after the key remainder
// (see flat.go), a value overflow in a vset.Set. Reading them costs no pointer chase beyond the
// leaf until a key holds more values than a flat leaf of 512 bytes. The zero
// value is an empty map.
type Map[T comparable] struct {
	t    Tree
	flat int8 // 1: T takes flat leaves, 2: typed leaves, 3: single-key pages (string), -1: value overflows only, 0: not decided yet
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
	if !l.isValueOverflow() {
		if flatPrepend[T](l, pre, b, pathLen) {
			return l
		}
		return reflat[T](l, wholeKey(l, pre, b), pathLen)
	}
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

// decide settles once per map whether T takes flat leaves (see flatType) or
// typed leaves (see typedType).
func (m *Map[T]) decide() {
	switch {
	case flatType[T]():
		m.flat = 1
	case stringType[T]():
		m.flat = 3
	case typedType[T]():
		m.flat = 2
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
		nl = newFlatLeaf[T]
	case 2:
		nl = newTypedLeaf[T]
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
	case l.isValueOverflow():
		vals[T](l).Add(v)
	case m.flat == 3:
		addSK(loc, l, key, *(*string)(unsafe.Pointer(&v)))
	case m.flat == 2:
		if nl := typedAdd(l, v); nl != nil {
			*loc = singleKeyHdr(nl)
		}
	default:
		if nl := flatAdd(l, v); nl != nil {
			*loc = singleKeyHdr(nl)
		}
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
		if m.flat == 3 {
			m.t.removeSK(l, key, *(*string)(unsafe.Pointer(&v)), rk)
			return
		}
		if m.flat == 2 {
			m.removeTyped(l, key, v, rk)
			return
		}
		switch c, empty := flatRemove(l, v); {
		case empty:
			m.t.remove(key, rk)
		case c != 0:
			*m.t.findSlot(key) = singleKeyHdr(resize[T](l, c))
		}
		return
	}
	if m.flat == 3 {
		m.removeFromOverflowSet(l, key, *(*string)(unsafe.Pointer(&v)), rk)
		return
	}
	s := vals[T](l)
	switch {
	case !s.Remove(v):
	case s.Len() == 0:
		m.t.remove(key, rk)
	case m.flat == 1:
		if c := unspillClass[T](l); c != 0 {
			*m.t.findSlot(key) = singleKeyHdr(unspill[T](l, c))
		}
	case m.flat == 2:
		if c := unspillTypedClass[T](l); c != 0 {
			*m.t.findSlot(key) = singleKeyHdr(unspillTyped[T](l, c))
		}
	}
}

// removeFromOverflowSet removes v from the value overflow l of key in a map of strings, and
// moves the key into a page once its values fit one.
func (m *Map[T]) removeFromOverflowSet(l *singleKeyHead, key []byte, v string, rk rekeyFunc) {
	s := *overflowSetOf(l)
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

// removeTyped removes v from the typed leaf l of key; rk is the map's rekeyFunc.
func (m *Map[T]) removeTyped(l *singleKeyHead, key []byte, v T, rk rekeyFunc) {
	switch c, empty := typedRemove(l, v); {
	case empty:
		m.t.remove(key, rk)
	case c != 0:
		*m.t.findSlot(key) = singleKeyHdr(resizeTyped[T](l, c))
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) {
	m.t.remove(key, m.rekeyFunc())
}

// rekeyFunc returns the map's rekeyFunc: the one for its type of leaf.
func (m *Map[T]) rekeyFunc() rekeyFunc {
	switch m.flat {
	case 2:
		return rekeyTyped[T]
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
	if flat == 3 { // T is string
		y := *(*func(string) bool)(unsafe.Pointer(&yield))
		if l.isValueOverflow() {
			return overflowEach(l, y)
		}
		return asSK(l).Strings(y)
	}
	if l.isValueOverflow() {
		return vals[T](l).Each(yield)
	}
	for _, v := range flatVals[T](l) {
		if !yield(v) {
			return false
		}
	}
	return true
}

// leafTail is the offset of the last byte of the smallest leaf this map
// creates, which the scan touches ahead (see touchChildren). Every leaf is
// at least that large, and a constant offset keeps the load independent of
// the leaf's head.
func (m *Map[T]) leafTail() uintptr {
	var z T
	switch m.flat {
	case 1:
		return flatSizes[1] - 1
	case 2:
		return typedOff(0) + unsafe.Sizeof(z) - 1
	case 3:
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
