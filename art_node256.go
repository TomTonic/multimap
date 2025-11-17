package multimap

type node256[T comparable] struct {
	node[T]
	numChildren uint16
	child       [256]*node[T]
}

func (n *node256[T]) currentChildCount() uint32 {
	return uint32(n.numChildren)
}

func (n *node256[T]) hasCapacityForChild() bool {
	return n.numChildren < 256
}
