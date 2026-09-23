package multimap

// designed to be exactly 64 bytes in size so one node5 exactly
// fits into one cache line (most x86/x64 and arm64 CPUs)
type node5[T comparable] struct {
	node[T]
	numChildren       byte
	firstKeyByteChild [5]byte
	child             [5]*node[T]
}

func (n *node5[T]) currentChildCount() uint32 {
	return uint32(n.numChildren)
}

func (n *node5[T]) hasCapacityForChild() bool {
	return n.numChildren < 5
}

func (n *node5[T]) getChild(currentPrefix Key, searchKey Key) (child *node[T]) {
	lokalKey := n.appendLocalPrefix(currentPrefix)
	lcp := LongestCommonPrefix(lokalKey, searchKey)
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
		if n.firstKeyByteChild[i] == nextKeyByte && n.child[i] != nil {
			return n.child[i].getChild(lokalKey, searchKey)
		}
	}
	// no matching child found
	return nil
}

func (n *node5[T]) getOrCreateChild(currentPrefix Key, searchKey Key) (child *node[T], replacement *node[T]) {
	lcp := len(currentPrefix)
	lsk := len(searchKey)
	if lcp > lsk {
		panic("structural problem: lcp > lsk, prefix: " + string(currentPrefix) + ", searchKey: " + string(searchKey))
		// return nil
	}
	if lcp == lsk {
		if !currentPrefix.Equal(searchKey) {
			panic("structural problem: lcp == lsk but keys not equal, prefix: " + string(currentPrefix) + ", searchKey: " + string(searchKey))
		}
		return n.asNode(), nil
	}
	nextKeyByte := searchKey[lcp]
	for i := 0; i < int(n.numChildren); i++ {
		if n.firstKeyByteChild[i] == nextKeyByte && n.child[i] != nil {
			ch, rep := n.child[i].getOrCreateChild(n.appendLocalPrefix(currentPrefix), searchKey)
			if rep != nil {
				n.child[i] = rep
			}
			return ch, nil
		}
	}
	// child not found, create it
	newChild := &nodeLeaf[T]{}
	newChild.ntype = Leaf
	newChild.prefixLen = byte(len(searchKey) - (lcp + 1))
	copy(newChild.prefix[:], searchKey[lcp+1:])

	if n.hasCapacityForChild() {
		n.firstKeyByteChild[n.numChildren] = nextKeyByte
		n.child[n.numChildren] = newChild.asNode()
		n.numChildren++
		return newChild.asNode(), nil
	} else {
		// no capacity, need to grow to node51
		newMe := n.toNode51(nextKeyByte, newChild)
		return newChild.asNode(), newMe.asNode()
	}
}

func (n *node5[T]) toNode51(nextKeyByte byte, newChild *nodeLeaf[T]) *node51[T] {
	newMe := &node51[T]{}
	// copy common fields
	newMe.ntype = Node51
	newMe.prefixLen = n.prefixLen
	copy(newMe.prefix[:], n.prefix[:n.prefixLen])
	// copy existing children
	for i := 0; i < int(n.numChildren); i++ {
		kb := n.firstKeyByteChild[i]
		pos := newMe.getPositionForInsertingKeyByte(kb)
		newMe.insertChildAt(pos, kb, n.child[i])
	}
	// add the new child
	pos := newMe.getPositionForInsertingKeyByte(nextKeyByte)
	newMe.insertChildAt(pos, nextKeyByte, newChild.asNode())
	return newMe
}
