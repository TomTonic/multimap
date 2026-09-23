package multimap

// Designed to be exactly 512 bytes in size so one node51 exactly
// fits into 8 cache lines of most x86/x64 and arm64 CPUs. Due to
// prefetching and burst reads (DDR memory), a node51 should be loaded
// in only a couple of nanoseconds.
type node51[T comparable] struct {
	node[T]
	numChildren       byte
	firstKeyByteChild [51]byte
	presence          bitfield256
	child             [51]*node[T]
}

func (n *node51[T]) currentChildCount() uint32 {
	return uint32(n.numChildren)
}

func (n *node51[T]) hasCapacityForChild() bool {
	return n.numChildren < 51
}

func (n *node51[T]) getChild(currentPrefix Key, searchKey Key) (child *node[T]) {
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
		return n.asNode()
	}
	nextKeyByte := searchKey[lcp]
	if !n.presence.get(nextKeyByte) {
		return nil
	}
	position := n.binarySearchKeyByte(nextKeyByte)
	if position >= 0 {
		return n.child[position].getChild(n.appendLocalPrefix(currentPrefix), searchKey)
	} else {
		panic("structural problem: presence bit set but key byte not found, prefix: " + string(currentPrefix) + ", searchKey: " + string(searchKey))
		//return nil
	}
}

// This function tries to find an existing child for the given searchKey.
// Returns the index of the child in the firstKeyByteChild and child arrays
// if found, otherwise returns -1.
func (n *node51[T]) binarySearchKeyByte(target byte) (position int) {
	a := &(n.firstKeyByteChild)
	lo, hi := 0, int(n.numChildren)
	for lo < hi {
		mid := (lo + hi) / 2
		if (*a)[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(a) && (*a)[lo] == target {
		return lo
	}
	return -1
}

// This is an internal function that assumes there is capacity for a new child
// and that the position is valid (0 <= position <= numChildren)
// and that keyByte is not already present
// and that the child is not nil and that numChildren < 51
// and that the position is the correct position to keep firstKeyByteChild sorted
// and that keyByte is the first byte of the child's key at the current prefix level
func (n *node51[T]) insertChildAt(position int, keyByte byte, child *node[T]) {
	for i := int(n.numChildren); i > position; i-- {
		n.firstKeyByteChild[i] = n.firstKeyByteChild[i-1]
		n.child[i] = n.child[i-1]
	}
	n.firstKeyByteChild[position] = keyByte
	n.child[position] = child
	n.presence.set(keyByte)
	n.numChildren++
}

// this is an internal function that returns the position of a keyByte that
// should be inserted with insertChildAt() to keep firstKeyByteChild sorted
// it assumes that there is capacity for a new child
// and that numChildren < 51
// and that keyByte is not already present
func (n *node51[T]) getPositionForInsertingKeyByte(keyByte byte) (position int) {
	if n.numChildren == 0 {
		return 0
	}
	lo, hi := 0, int(n.numChildren)
	for lo < hi {
		mid := (lo + hi) / 2
		if n.firstKeyByteChild[mid] < keyByte {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
