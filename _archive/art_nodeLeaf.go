package multimap

import set3 "github.com/TomTonic/Set3"

type nodeLeaf[T comparable] struct {
	node[T]
	data  *set3.Set3[T]
	child *node[T]
}

func (n *nodeLeaf[T]) currentChildCount() uint32 {
	if n.prefixLen > 0 && n.child != nil {
		return 1
	}
	return 0
}

func (n *nodeLeaf[T]) hasCapacityForChild() bool {
	return n.prefixLen == 0 && n.child == nil
}
