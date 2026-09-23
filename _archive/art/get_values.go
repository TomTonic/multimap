package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	mm "github.com/TomTonic/multimap"
)

func (n *Node[T]) getValuesFromTo(currentPrefix mm.Key, fromKey, toKey mm.Key, includingFrom, includingTo bool, result *set3.Set3[T]) {
	switch n.GetNodeType() {
	case NodeTypeLeaf:
		(*LeafNode[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case NodeType64:
		(*Node64[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case NodeType128:
		(*Node128[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case NodeType256:
		(*Node256[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case NodeType512:
		(*Node512[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case NodeType1024:
		(*Node1024[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	case FullNodeType:
		(*FullNode[T])(unsafe.Pointer(n)).getValuesFromTo(currentPrefix, fromKey, toKey, includingFrom, includingTo, result)
	}
}

func (n *LeafNode[T]) getValuesFromTo(currentPrefix mm.Key, fromKey, toKey mm.Key, includingFrom, includingTo bool, result *set3.Set3[T]) {
	lokalKey := n.appendLocalPrefixTo(currentPrefix)

	// first the simple cases where we can decide based on simple comparisons
	// without looking at the special cases where localKey might be a prefix of fromKey
	if fromKey.LessThan(lokalKey) {
		// fromKey is either strictly less than or a prefix of lokalKey
		if lokalKey.LessThan(toKey) {
			// lokalKey is either strictly less than or a prefix of toKey
			// lokalKey is strictly between fromKey and toKey -> match
			result.AddAll(n.value)
			if n.child != nil {
				n.child.getValuesFromTo(lokalKey, fromKey, toKey, includingFrom, includingTo, result)
			}
			return
		} else {
			// fromKey is either strictly less than or a prefix of lokalKey
			// and lokalKey is not(either strictly less than or a prefix of toKey)
			// i.e.
			// fromKey is either strictly less than or a prefix of lokalKey
			// and lokalKey is not strictly less than and not a prefix of toKey
			// -> lokalKey is greater than or equal to toKey
			if includingTo && lokalKey.Equal(toKey) {
				// lokalKey equals toKey and we include toKey -> match
				result.AddAll(n.value)
				// no need to go to child, lokalKey == toKey excludes any longer keys
				return
			}
			// else: no match
			return
		}
	}

	// now we know that fromKey is not less than lokalKey
	// i.e. fromKey is either equal to or greater than lokalKey
	if !includingFrom {
		// localKey is not allowed to equal fromKey
		// -> if localKey equals fromKey -> no match
		// -> if localKey is less than fromKey -> no match
		// => no match
		return
	}

	// now we know that fromKey is not less than lokalKey and
	// that localKey is allowed to equal fromKey
	// i.e.
	// if lokalKey is less than fromKey -> no match
	// if lokalKey is a prefix of fromKey -> potential match
	if lokalKey.LessThan(fromKey) {
		// lokalKey is strictly less than fromKey -> no match
		return
	}

	lcpFrom := mm.LongestCommonPrefixAsKey(lokalKey, fromKey)

	lcpF := mm.LongestCommonPrefix(lokalKey, fromKey)
	lcpT := mm.LongestCommonPrefix(lokalKey, toKey)
	llk := uint(len(lokalKey))
	lskF := uint(len(fromKey))
	lskT := uint(len(toKey))

	// always: lcpF <= llk and lcpF <= lskF
	// always: lcpT <= llk and lcpT <= lskT

	contained := getContained(lokalKey, fromKey, toKey, includingFrom, includingTo)

	if lcp == llk && lcp == lsk {
		// llk == lcp == lsk -> localKey and searchKey are equal
		if cmpOp&CmpEqual != 0 {
			// we look for less than equal, greater than equal or equal
			target.AddAll(n.value)
			return
		} else {
			// but we either look for strictly less than or strictly greater than -> no match
			return
		}
	}

	cmpVal := compareNextPos(lokalKey, searchKey, int(lcp))

	if cmpVal&cmpOp == 0 {
		// keys differ at position lcp
		// and the comparison result does not match the requested comparison operation(s) -> no match
		return
	}

	if lcp < llk {
		// case a) lcp < lsk -> keys differ -> no match
		// case b) lcp = lsk -> searchKey is true prefix of lokalKey -> no match
		return
	}

	if lcp == lsk {
		// lcp = llk and lcp = lsk -> localKey and searchKey are equal - we found the
		target.AddAll(n.value)
	}

	if n.child != nil {
		// lcp = llk and lcp < lsk -> localKey is a true prefix of searchKey
		n.child.getValuesFromTo(lokalKey, searchKey, target)
	}

	// child did not match
	return
}

// implementations for nodes with multiple children
// but without bitmap and linear/unrolled search
// these are Node64, Node128, Node256
// code is identical except for the node type

func (n *Node64[T]) getValuesFromTo(currentPrefix mm.Key, fromKey, toKey mm.Key, includingFrom, includingTo bool, result *set3.Set3[T]) (child *Node[T]) {
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
