package art

import "github.com/TomTonic/multimap/internal/skpage"

// This file measures the objects of a tree against the cache-line rules of
// docs/redesign/STRATEGY.md: how many objects there are, how big each is, and
// where the Go allocator puts it. The bench's cmd/objstat reports it for every
// benchmark case. A new object type needs a case in Objects, which the test
// that adds up the keys of all objects (TestObjects) fails without.

// Object describes one allocated object of a tree.
type Object struct {
	// Label names the type of the object: "N5", "R56+tail", "page", "flat
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
// a value overflow allocates for the Set3 of its values.
func (m *Map[T]) Objects(fn func(Object)) {
	if m.t.root != nil {
		m.objects(m.t.root, fn)
	}
}

var objTypeLabels = [64]string{
	kN5: "N5", kN12: "N12", kN26: "N26", kN58: "N58", kN256: "N256",
}

// leafObject describes leaf l of a map of T.
func (m *Map[T]) leafObject(l *singleKeyHead) Object {
	o := m.leafKind(l)
	o.Values, o.Remainder = m.leafValues(l), len(l.stored())
	return o
}

// leafValues returns the number of values of leaf l.
func (m *Map[T]) leafValues(l *singleKeyHead) int {
	if l.isValueOverflow() {
		return int((*overflowSetOf[T](l)).Size())
	}
	return int(l.n)
}

// leafKind describes the object of leaf l of a map of T, without its values
// and remainder.
func (m *Map[T]) leafKind(l *singleKeyHead) Object {
	if l.cls() > 0 {
		return Object{Label: "single-key page", Size: asSK(l).Size(), Pointers: m.flat == 1 && skpage.HoldsPointers[T](), Keys: 1}
	}
	return Object{Label: "value overflow", Size: int(valueOverflowSize(l.rem())), Pointers: true, Keys: 1}
}

// multiKeyObject describes multi-key page n: its keys and values, and Remainder is the common
// prefix of its keys.
func (m *Map[T]) multiKeyObject(n *header) Object {
	if m.flat == 3 {
		p := asMKStr(n)
		return Object{Label: "multi-key page", Size: p.Size(), Keys: p.Keys(), Values: p.Len(), Remainder: p.PrefixLen()}
	}
	p := asMKFix(n)
	return Object{Label: "multi-key page", Size: p.Size(), Keys: p.Keys(), Values: p.Len(), Remainder: p.PrefixLen()}
}

// object describes the object n, a leaf, page or node, without what is below
// it.
func (m *Map[T]) object(n *header) Object {
	switch {
	case isSingleKey(n.objType):
		return m.leafObject(asSingleKey(n))
	case isMultiKey(n.objType) && isPage(n.objType):
		return m.multiKeyObject(n)
	}
	size, label := int(fixedSize[n.objType&objTypeMask]), objTypeLabels[n.objType]
	if tc := tailClass(n.prefixLen()); tc != tailNone {
		size += [...]int{tail16: 16, tail48: 48, tail112: 112, tailStr: 16}[tc]
		label += "+tail"
	}
	return Object{Label: label, Size: size, Pointers: true}
}

// objects reports the subtree n.
func (m *Map[T]) objects(n *header, fn func(Object)) {
	fn(m.object(n))
	if isPage(n.objType) {
		return
	}
	if t := endPageOf(n); t != nil {
		fn(m.leafObject(t))
	}
	eachByteNode(n, func(_ byte, c *header) { m.objects(c, fn) })
}
