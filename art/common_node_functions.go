package art

import "unsafe"

type NodeType uint8

const (
	NodeTypeLeaf NodeType = 0
	NodeType64   NodeType = 1
	NodeType128  NodeType = 2
	NodeType256  NodeType = 3
	NodeType512  NodeType = 4
	NodeType1024 NodeType = 5
	FullNodeType NodeType = 6
)

func (nt NodeType) String() string {
	return []string{"Leaf", "Node64", "Node128", "Node256", "Node512", "Node1024", "FullNode"}[nt]
}

func (n *Node[T]) getNodeType() NodeType {
	nt := NodeType((n.meta >> 4) & 0x0F)
	return NodeType(nt)
}

func (n *Node[T]) setNodeType(nt NodeType) {
	n.meta = (n.meta & 0x0F) | (uint8(nt) << 4)
}

func (n *Node[T]) getPrefixLen() uint8 {
	return n.meta & 0x0F
}

func (n *Node[T]) setPrefixLen(plen uint8) {
	n.meta = (n.meta & 0xF0) | (plen & 0x0F)
}

func (n *Node[T]) getDirectChildren() uint {
	return uint(n.numChildren)
}

func (n *Node[T]) getMaxChildren() uint {
	switch n.getNodeType() {
	case NodeTypeLeaf:
		return maxChildrenLeaf
	case NodeType64:
		return maxChildrenNode64
	case NodeType128:
		return maxChildrenNode128
	case NodeType256:
		return maxChildrenNode256
	case NodeType512:
		return maxChildrenNode512
	case NodeType1024:
		return maxChildrenNode1024
	case FullNodeType:
		return maxChildrenFullNode
	default:
		return 0
	}
}

// Type casting helpers

func (n *Node[T]) asNode64() *Node64[T] {
	if n.getNodeType() != NodeType64 {
		panic("node is not of kind Node64 but of kind " + n.getNodeType().String())
	}
	return (*Node64[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode128() *Node128[T] {
	if n.getNodeType() != NodeType128 {
		panic("node is not of kind Node128 but of kind " + n.getNodeType().String())
	}
	return (*Node128[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode256() *Node256[T] {
	if n.getNodeType() != NodeType256 {
		panic("node is not of kind Node256 but of kind " + n.getNodeType().String())
	}
	return (*Node256[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode512() *Node512[T] {
	if n.getNodeType() != NodeType512 {
		panic("node is not of kind Node512 but of kind " + n.getNodeType().String())
	}
	return (*Node512[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode1024() *Node1024[T] {
	if n.getNodeType() != NodeType1024 {
		panic("node is not of kind Node1024 but of kind " + n.getNodeType().String())
	}
	return (*Node1024[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asFullNode() *FullNode[T] {
	if n.getNodeType() != FullNodeType {
		panic("node is not of kind FullNode but of kind " + n.getNodeType().String())
	}
	return (*FullNode[T])(unsafe.Pointer(n))
}

func (n *Node64[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}

func (n *Node128[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}

func (n *Node256[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}

func (n *Node512[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}

func (n *Node1024[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}

func (n *FullNode[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
}
