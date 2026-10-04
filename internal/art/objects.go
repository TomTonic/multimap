package art

import "unsafe"

// This file measures the objects of a tree against the cache-line rules of
// docs/redesign/STRATEGY.md: how many objects there are, how big each is, and
// where the Go allocator puts it. The bench's cmd/objstat reports it for every
// benchmark case. A new object kind needs a case in Objects, which the test
// that adds up the keys of all objects (TestObjects) fails without.

// Object describes one allocated object of a tree.
type Object struct {
	// Label names the kind of the object: "N5", "R56+tail", "page", "flat
	// leaf" and so on. A "+tail" node holds a prefix tail beyond its header.
	Label string
	// Size is the size of the object in bytes, as Go allocates the type.
	Size int
	// Pointers says whether the object holds pointers. Go gives such an
	// object of more than 512 bytes a malloc header (see Block).
	Pointers bool
	// Keys is the number of keys the object holds: those of a page, one for a
	// leaf, none for a node.
	Keys int
	// Values is the number of values of a leaf, and Remainder the length of the
	// key remainder it stores (the key from its base on); both are 0 for pages
	// and nodes.
	Values, Remainder int
}

// goClasses are the size classes of Go's allocator up to 8 KiB
// (runtime/sizeclasses.go of Go 1.27). TestBlock checks them against the
// runtime of the Go that runs the tests.
var goClasses = [...]int{
	8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320,
	352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792,
	2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192,
}

// mallocHeader is what Go puts in front of an object with pointers of more
// than maxNoHeader bytes: its type, so that the garbage collector can find the
// pointers. It moves the object by that much inside its block.
const (
	mallocHeader = 8
	maxNoHeader  = 512
)

// Block returns the size of the block Go allocates for an object of size
// bytes, and the offset of the object inside the block: mallocHeader for an
// object with pointers of more than 512 bytes, else 0. Sizes above the largest
// class are rounded up to whole pages of 8 KiB.
func Block(size int, pointers bool) (block, offset int) {
	need := size
	if pointers && size > maxNoHeader {
		need, offset = size+mallocHeader, mallocHeader
	}
	for _, c := range goClasses {
		if c >= need {
			return c, offset
		}
	}
	const page = 8192
	return (need + page - 1) / page * page, offset
}

// Objects calls fn for every object of the tree, in key order of the paths to
// them, nodes before their children. It does not count what a value set of
// a set leaf allocates for itself (see vset).
func (m *Map[T]) Objects(fn func(Object)) {
	if m.t.root != nil {
		m.objects(m.t.root, fn)
	}
}

var kindLabels = [32]string{
	kN5: "N5", kN12: "N12", kN26: "N26", kN58: "N58", kN256: "N256",
	kR8: "R8", kR24: "R24", kR56: "R56", kR256: "R256",
}

// leafObject describes leaf l of a map of T.
func (m *Map[T]) leafObject(l *leafHead) Object {
	o := m.leafKind(l)
	o.Values, o.Remainder = m.leafValues(l), len(l.stored())
	return o
}

// leafValues returns the number of values of leaf l.
func (m *Map[T]) leafValues(l *leafHead) int {
	if l.kind == kSet {
		return vals[T](l).Len()
	}
	return int(l.n)
}

// leafKind describes the object of leaf l of a map of T, without its values
// and remainder.
func (m *Map[T]) leafKind(l *leafHead) Object {
	var z T
	switch {
	case l.cls() > 0 && m.flat == 1:
		return Object{Label: "flat leaf", Size: int(flatSizes[l.cls()]), Keys: 1}
	case l.cls() > 0 && m.flat == 2:
		return Object{Label: "typed leaf", Size: int(typedOff(int(l.klen))) + typedCaps[l.cls()]*int(unsafe.Sizeof(z)), Pointers: true, Keys: 1}
	}
	var size uintptr
	switch k := l.klen; {
	case k <= 16:
		size = unsafe.Sizeof(leaf[T, [16]byte]{})
	case k <= 32:
		size = unsafe.Sizeof(leaf[T, [32]byte]{})
	case k <= 48:
		size = unsafe.Sizeof(leaf[T, [48]byte]{})
	case k <= 64:
		size = unsafe.Sizeof(leaf[T, [64]byte]{})
	case k <= 96:
		size = unsafe.Sizeof(leaf[T, [96]byte]{})
	case k <= 128:
		size = unsafe.Sizeof(leaf[T, [128]byte]{})
	case k <= 192:
		size = unsafe.Sizeof(leaf[T, [192]byte]{})
	case k <= maxInline:
		size = unsafe.Sizeof(leaf[T, [256]byte]{})
	default:
		size = unsafe.Sizeof(leaf[T, string]{})
	}
	return Object{Label: "set leaf", Size: int(size), Pointers: true, Keys: 1}
}

// object describes the object n, a leaf, page or node, without what is below
// it.
func (m *Map[T]) object(n *header) Object {
	switch {
	case isLeaf(n.kind):
		return m.leafObject(asLeaf(n))
	case isPage(n.kind):
		p := asPage(n)
		return Object{Label: "page", Size: p.Size(), Keys: p.Len()}
	}
	size, label := int(fixedSize[n.kind&kindMask]), kindLabels[n.kind]
	if tc := tailClass(n.prefixLen()); tc != tailNone {
		size += [...]int{tail16: 16, tail48: 48, tail112: 112, tailStr: 16}[tc]
		label += "+tail"
	}
	return Object{Label: label, Size: size, Pointers: true}
}

// objects reports the subtree n.
func (m *Map[T]) objects(n *header, fn func(Object)) {
	fn(m.object(n))
	if isLeaf(n.kind) || isPage(n.kind) {
		return
	}
	if t := termOf(n); t != nil {
		fn(m.leafObject(t))
	}
	if isRange(n.kind) {
		for _, c := range asR(n).children()[:asR(n).n] {
			m.objects(c, fn)
		}
		return
	}
	eachInner(n, func(_ byte, c *header) { m.objects(c, fn) })
}
