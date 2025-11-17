package multimap

// Designed to be exactly 512 bytes in size so one node51 exactly
// fits into 8 cache lines of most x86/x64 and arm64 CPUs. Due to
// prefetching and burst reads (DDR memory), a node51 should be loaded
// in only a couple of nanoseconds.
// the member order is
type node51[T comparable] struct {
	node[T]
	numChildren       byte
	firstChildKeyByte [51]byte
	presence          bitfield256
	child             [51]*node[T]
}

func (n *node51[T]) currentChildCount() uint32 {
	return uint32(n.numChildren)
}

func (n *node51[T]) hasCapacityForChild() bool {
	return n.numChildren < 51
}
