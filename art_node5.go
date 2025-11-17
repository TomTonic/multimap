package multimap

// designed to be exactly 64 bytes in size so one node5 exactly
// fits into one cache line (most x86/x64 and arm64 CPUs)
type node5[T comparable] struct {
	node[T]
	numChildren       byte
	firstChildKeyByte [5]byte
	child             [5]*node[T]
}

func (n *node5[T]) currentChildCount() uint32 {
	return uint32(n.numChildren)
}

func (n *node5[T]) hasCapacityForChild() bool {
	return n.numChildren < 5
}
