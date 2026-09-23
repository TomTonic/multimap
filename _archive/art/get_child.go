package art

import (
	"unsafe"

	mm "github.com/TomTonic/multimap"
)

func (n *Node[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	switch n.GetNodeType() {
	case NodeTypeLeaf:
		return (*LeafNode[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case NodeType64:
		return (*Node64[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case NodeType128:
		return (*Node128[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case NodeType256:
		return (*Node256[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case NodeType512:
		return (*Node512[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case NodeType1024:
		return (*Node1024[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	case FullNodeType:
		return (*FullNode[T])(unsafe.Pointer(n)).getChild(currentPrefix, searchKey)
	}
	return nil
}

func (n *LeafNode[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	if n.child != nil {
		// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
		return n.child.getChild(lokalKey, searchKey)
	}

	// child did not match
	return nil
}

// implementations for nodes with multiple children
// but without bitmap and linear/unrolled search
// these are Node64, Node128, Node256
// code is identical except for the node type

func (n *Node64[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	for i := 0; i < int(n.numChildren); i++ {
		if n.firstKeyByte[i] == nextKeyByte && n.child[i] != nil {
			return n.child[i].getChild(lokalKey, searchKey)
		}
	}
	// no matching child found
	return nil
}

func (n *Node128[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	for i := 0; i < int(n.numChildren); i++ {
		if n.firstKeyByte[i] == nextKeyByte && n.child[i] != nil {
			return n.child[i].getChild(lokalKey, searchKey)
		}
	}
	// no matching child found
	return nil
}

func (n *Node256[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	for i := 0; i < int(n.numChildren); i++ {
		if n.firstKeyByte[i] == nextKeyByte && n.child[i] != nil {
			return n.child[i].getChild(lokalKey, searchKey)
		}
	}
	// no matching child found
	return nil
}

// implementations for nodes with multiple children
// with bitmap and binary search
// these are Node512, Node1024
// code is identical except for the node type

func (n *Node512[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	if !n.bitmap.Get(nextKeyByte) {
		// no child for the next key byte
		return nil
	}
	position := n.binarySearchKeyByte(nextKeyByte)
	if position >= 0 {
		return n.child[position].getChild(lokalKey, searchKey)
	} else {
		panic("structural problem: presence bit set but key byte not found, prefix: " + string(currentPrefix) + ", searchKey: " + string(searchKey))
	}
}

// This function tries to find an existing child for the given searchKey.
// Returns the index of the child in the firstKeyByteChild and child arrays
// if found, otherwise returns -1.
func (n *Node512[T]) binarySearchKeyByte(target byte) (position int) {
	lo, hi := 0, int(n.numChildren)
	for lo < hi {
		mid := (lo + hi) / 2
		if n.firstKeyByte[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(n.firstKeyByte) && n.firstKeyByte[lo] == target {
		return lo
	}
	return -1
}

func (n *Node1024[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	if !n.bitmap.Get(nextKeyByte) {
		// no child for the next key byte
		return nil
	}
	position := n.binarySearchKeyByte(nextKeyByte)
	if position >= 0 {
		return n.child[position].getChild(lokalKey, searchKey)
	} else {
		panic("structural problem: presence bit set but key byte not found, prefix: " + string(currentPrefix) + ", searchKey: " + string(searchKey))
	}
}

// This function tries to find an existing child for the given searchKey.
// Returns the index of the child in the firstKeyByteChild and child arrays
// if found, otherwise returns -1.
func (n *Node1024[T]) binarySearchKeyByte(target byte) (position int) {
	lo, hi := 0, int(n.numChildren)
	for lo < hi {
		mid := (lo + hi) / 2
		if n.firstKeyByte[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(n.firstKeyByte) && n.firstKeyByte[lo] == target {
		return lo
	}
	return -1
}

// implementations for FullNode

func (n *FullNode[T]) getChild(currentPrefix mm.Key, searchKey mm.Key) (child *Node[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)
	lcp := mm.LongestCommonPrefix(lokalKey, searchKey)
	llk := uint(len(lokalKey))
	lsk := uint(len(searchKey))

	// always: lcp <= llk and lcp <= lsk

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return nil
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the node
		return n.asNode()
	}

	// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
	nextKeyByte := searchKey[lcp]
	// first check bitmap for presence - this data should already reside in the same cache line
	if !n.bitmap.Get(nextKeyByte) {
		// no child for the next key byte
		return nil
	}
	// presence bit is set - directly access child array using nextKeyByte as index
	if n.child[nextKeyByte] != nil {
		return n.child[nextKeyByte].getChild(lokalKey, searchKey)
	}
	// no matching child found
	return nil
}
