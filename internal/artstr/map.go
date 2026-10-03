package artstr

import (
	"bytes"
	"slices"
	"unsafe"

	"github.com/TomTonic/multimap/internal/lpage"
	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// leaf holds its key's values: a flat leaf right after the key remainder
// (see flat.go), a set leaf in a vset.Set. Reading them costs no pointer chase beyond the
// leaf until a key holds more values than a flat leaf of 512 bytes. The zero
// value is an empty map.
type Map[T comparable] struct {
	// ZeroCopy, set before the first write, makes the pages immutable and
	// hands out the strings of single-valued keys as views of the page, with
	// no copy: a lookup costs no allocation. Every change then builds a new
	// page, and a string a caller keeps holds its page, up to 512 bytes, alive.
	ZeroCopy bool
	// Pairs, set before the first write, keeps a key with several values in its
	// page: the page stores the key once and a value after it. Without it a key
	// leaves its page for a leaf when it gets a second value.
	Pairs bool
	t     Tree
	flat  int8 // 1: T takes flat leaves, 2: typed leaves, -1: set leaves only, 0: not decided yet
}

// newSetLeaf allocates a set leaf that holds key from base on, in the
// smallest size class that fits it, or the whole key as a string when the
// rest is too long to hold inline. It captures nothing, so passing it as a
// newLeafFunc allocates no closure.
func newSetLeaf[T comparable](key []byte, base int) *leafHead {
	if len(key) > maxKeyLen || len(key)-base > maxInline {
		l := &leaf[T, string]{k: string(key)}
		l.kind, l.klen = kSet, longKey
		return &l.leafHead
	}
	return newSetLeafOf[T](key[base:], len(key))
}

// newSetLeafOf allocates a set leaf holding the key remainder s, at most
// maxInline bytes, of a key of kl bytes.
func newSetLeafOf[T comparable](s []byte, kl int) *leafHead {
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

// newInline allocates a set leaf that holds the key remainder s inline in an
// array of type K.
func newInline[T comparable, K [16]byte | [32]byte | [48]byte | [64]byte | [96]byte | [128]byte | [192]byte | [256]byte](s []byte, kl int) *leafHead {
	l := &leaf[T, K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.kind, l.klen, l.kl = kSet, uint8(len(s)), uint16(kl)
	return &l.leafHead
}

// Offsets of the value set in the set leaf of each key size class. They do
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

// vals returns the value set of a set leaf, created by newSetLeaf[T].
func vals[T comparable](l *leafHead) *vset.Set[T] {
	off := setOffStr
	switch k := l.klen; {
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
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), off))
}

// rekey returns a leaf with l's values that holds its key from depth on (see
// rekeyFunc). It captures nothing, so passing it allocates no closure. It
// keeps l where it is when the longer remainder fits its size class, and
// builds the whole key only when it has to copy l.
func rekey[T comparable](l *leafHead, pre []byte, b, depth int) *leafHead {
	if l.kind != kSet {
		if flatPrepend[T](l, pre, b, depth) {
			return l
		}
		return reflat[T](l, wholeKey(l, pre, b), depth)
	}
	if setPrepend(l, pre, b, depth) {
		return l
	}
	nl := newSetLeaf[T](wholeKey(l, pre, b), depth)
	*vals[T](nl) = *vals[T](l)
	return nl
}

// setKeyCap returns the size of the key area of a set leaf whose key
// remainder is klen bytes long, at most maxInline: the class that klen falls
// in (see vals), whose largest holds 256 bytes.
func setKeyCap(klen int) int {
	for _, c := range [...]int{16, 32, 48, 64, 96, 128, 192} {
		if klen <= c {
			return c
		}
	}
	return 256
}

// setPrepend makes set leaf l hold its key from depth on in place when the
// longer remainder still fits the key area of l's class, which keeps its value
// set where it is. It saves rekey the allocation of a new leaf, see
// flatPrepend, and reports whether it did.
func setPrepend(l *leafHead, pre []byte, b, depth int) bool {
	old, klen := int(l.klen), l.keyLen()-depth // a string leaf holds its whole key, and never gets here
	if klen > maxInline || klen > setKeyCap(old) {
		return false
	}
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), klen)
	copy(area[klen-old:], area[:old])
	fillHead(area[:klen-old], pre, b, depth)
	l.klen = uint8(klen)
	return true
}

// decide settles once per map whether T takes flat leaves (see flatType) or
// typed leaves (see typedType), and tells the tree whether it may hold pages:
// only a map of strings does.
func (m *Map[T]) decide() {
	switch {
	case flatType[T]():
		m.flat = 1
	case typedType[T]():
		m.flat = 2
		if stringType[T]() {
			m.t.small, m.t.cow, m.t.pairs = true, m.ZeroCopy, m.Pairs
			m.t.mk, m.t.mkAll = leafWith[T], leafWithAll[T]
		}
	default:
		m.flat = -1
	}
}

// stringType reports whether T is string, the one type whose values go into
// pages, as bytes.
func stringType[T comparable]() bool {
	var z T
	_, ok := any(z).(string)
	return ok
}

// asString returns v, a T of a map that has pages, which is a string.
func asString[T comparable](v T) string { return *(*string)(unsafe.Pointer(&v)) }

// fromString returns s as a T of a map that has pages, which is a string.
func fromString[T comparable](s string) T { return *(*T)(unsafe.Pointer(&s)) }

// leafWith allocates a typed leaf that holds key from base on and the value val
// of a page entry as its only value. It is the Tree's mk.
func leafWith[T comparable](key []byte, base int, val string) *leafHead {
	l := newTypedLeaf[T](key, base) // a set leaf if the key is too long for a typed one
	if v := fromString[T](val); l.kind == kSet {
		vals[T](l).Add(v)
	} else {
		appendTyped(l, v)
	}
	return l
}

// leafWithAll is leafWith for a key with several values. It is the Tree's mkAll.
func leafWithAll[T comparable](key []byte, base int, svals []string) *leafHead {
	l := newTypedLeaf[T](key, base)
	for _, s := range svals {
		v := fromString[T](s)
		if l.kind == kSet {
			vals[T](l).Add(v)
		} else if nl := typedAdd(l, v); nl != nil {
			l = nl
		}
	}
	return l
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
	}
	var sv string
	if m.t.small {
		sv = asString(v)
	}
	size := m.t.size
	loc := m.t.upsert(key, sv, nl)
	n := *loc
	switch {
	case isLeaf(n.kind):
		m.addToLeaf(loc, asLeaf(n), v)
		if m.t.size != size && m.t.small && isRange(m.t.root.kind) {
			// A new leaf among pages: keys that need leaves may crowd them.
			m.t.settle(key)
		}
	case m.t.size != size:
		// A page took the key and its value.
	default:
		// The key has one value in a page; a second one gives it a leaf, or in
		// pairs mode stays in its page.
		sp := m.t.at
		p := asPage(n)
		if m.t.pairs {
			m.addToPage(sp, key, sv)
			return
		}
		if p.ValueIs(sp.i, sv) {
			return
		}
		l := m.t.mk(key, sp.depth, string(p.ValueAt(sp.i)))
		slot := leafHdr(l)
		m.addToLeaf(&slot, l, v)
		m.t.promote(sp, asLeaf(slot), key)
	}
}

// addToPage adds the value v to the key at sp, which is in a page and may have
// values already (pairs mode). A full page splits where its keys differ, and when
// it cannot, or the value fits no page, the page's keys are laid out anew with the
// key's values in a leaf if no page holds them.
func (m *Map[T]) addToPage(sp spot, key []byte, v string) {
	p := asPage(*sp.loc)
	if lpage.FitsLen(len(key)-sp.depth, len(v)) {
		q, res := p.AddValueAt(sp.i, bytesOf(v), m.t.cow)
		switch res {
		case lpage.Present:
			return
		case lpage.Added:
			*sp.loc = pageHdr(q)
			return
		}
		if sp.par != nil && splitFull(sp.loc, sp.par, sp.pi, sp.depth) {
			m.Add(key, fromString[T](v)) // the key's page is smaller now
			return
		}
	}
	items := pageItems(p, key, sp.depth)
	for k := range items {
		if it := &items[k]; bytes.Equal(it.key, key) {
			if it.val == v || slices.Contains(it.more, v) {
				return
			}
			it.more = append(it.more, v)
		}
	}
	m.t.replace(sp.loc, sp.par, sp.pi, items, sp.depth)
}

// addToLeaf adds v to the values of leaf l, which sits in slot loc.
func (m *Map[T]) addToLeaf(loc **header, l *leafHead, v T) {
	switch {
	case l.kind == kSet:
		vals[T](l).Add(v)
	case m.flat == 2:
		if nl := typedAdd(l, v); nl != nil {
			*loc = leafHdr(nl)
		}
	default:
		if nl := flatAdd(l, v); nl != nil {
			*loc = leafHdr(nl)
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
	var rk rekeyFunc = rekey[T]
	if m.flat == 2 {
		rk = rekeyTyped[T]
	}
	if m.t.hasPages() {
		if m.t.removeValue(key, asString(v), rk) != keptLeaf {
			return
		}
	}
	n, _ := m.t.find(key)
	if n == nil {
		return
	}
	l := asLeaf(n)
	if l.kind != kSet {
		if m.flat == 2 {
			m.removeTyped(l, key, v, rk)
			return
		}
		switch c, empty := flatRemove(l, v); {
		case empty:
			m.t.remove(key, rk)
		case c != 0:
			*m.t.findSlot(key) = leafHdr(resize[T](l, c))
		}
		return
	}
	s := vals[T](l)
	switch {
	case !s.Remove(v):
	case s.Len() == 0:
		m.t.remove(key, rk)
	case m.flat == 1:
		if c := unspillClass[T](l); c != 0 {
			*m.t.findSlot(key) = leafHdr(unspill[T](l, c))
		}
	case m.flat == 2:
		if c := unspillTypedClass[T](l); c != 0 {
			*m.t.findSlot(key) = leafHdr(unspillTyped[T](l, c))
		}
	}
}

// removeTyped removes v from the typed leaf l of key; rk is the map's rekeyFunc.
func (m *Map[T]) removeTyped(l *leafHead, key []byte, v T, rk rekeyFunc) {
	switch c, empty := typedRemove(l, v); {
	case empty:
		m.t.remove(key, rk)
	case c != 0:
		*m.t.findSlot(key) = leafHdr(resizeTyped[T](l, c))
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) {
	var rk rekeyFunc = rekey[T]
	if m.flat == 2 {
		rk = rekeyTyped[T]
	}
	m.t.remove(key, rk)
}

// Has reports whether key holds any values.
func (m *Map[T]) Has(key []byte) bool {
	n, _ := m.t.find(key)
	return n != nil
}

// Each calls yield for every value of key, in unspecified order, until it
// returns false. yield must not modify the map.
func (m *Map[T]) Each(key []byte, yield func(T) bool) {
	n, i := m.t.find(key)
	switch {
	case n == nil:
	case isPage(n.kind):
		p := asPage(n)
		p.EachString(i, p.RunEnd(i), m.ZeroCopy, func(s string) bool { return yield(fromString[T](s)) })
	default:
		eachValue(asLeaf(n), yield)
	}
}

// eachValue calls yield for every value of leaf l and reports whether it ran
// to completion.
func eachValue[T comparable](l *leafHead, yield func(T) bool) bool {
	if l.kind == kSet {
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
	}
	return unsafe.Sizeof(leaf[T, [16]byte]{}) - 1
}

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key is assembled for fn, from the path to its leaf and
// the rest the leaf holds, or from its page: fn must not modify or retain it,
// and must not modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte) bool) {
	var kb keyBuf
	m.t.scan(b, m.leafTail(), &kb, func(n *header, i, j int) bool {
		if isLeaf(n.kind) {
			return fn(kb.key)
		}
		p := asPage(n)
		for k := i; k < j; k++ {
			if !p.IsCont(k) && !fn(kb.pageKey(p, k)) {
				return false
			}
		}
		return true
	})
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, m.leafTail(), nil, func(n *header, i, j int) bool {
		if isLeaf(n.kind) {
			return eachValue(asLeaf(n), yield)
		}
		return asPage(n).EachString(i, j, m.ZeroCopy, func(s string) bool { return yield(fromString[T](s)) })
	})
}
