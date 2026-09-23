package art

import (
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	mm "github.com/TomTonic/multimap"
)

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

func (n *Node[T]) GetNodeType() NodeType {
	nt := NodeType((n.meta >> 4) & 0x0F)
	return NodeType(nt)
}

func (n *Node[T]) setNodeType(nt NodeType) *Node[T] {
	n.meta = (n.meta & 0x0F) | (uint8(nt) << 4)
	return n
}

func (n *Node[T]) GetPrefixLen() uint8 {
	return n.meta & 0x0F
}

func (n *Node[T]) GetPrefix() []byte {
	l := n.GetPrefixLen()
	k := make([]byte, l)
	copy(k, n.localPrefix[:l])
	return k
}

func (n *Node[T]) appendLocalPrefixTo(parentKey mm.Key) mm.Key {
	l := n.GetPrefixLen()
	result := make([]byte, len(parentKey)+int(l))
	copy(result, parentKey)
	copy(result[len(parentKey):], n.localPrefix[:l])
	return result
}

func (n *Node[T]) setPrefix(prefix []byte) *Node[T] {
	l := min(len(prefix), maxLocalPrefixLen)
	n.meta = (n.meta & 0xF0) | (uint8(l) & 0x0F)
	// copy the provided bytes
	copy(n.localPrefix[:], prefix[:l])
	// ensure remaining bytes are zeroed so no stale data remains
	for i := l; i < maxLocalPrefixLen; i++ {
		n.localPrefix[i] = 0
	}
	return n
}

func (n *Node[T]) GetMaxChildren() uint {
	switch n.GetNodeType() {
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
		panic("unknown node type: " + string((byte)(n.GetNodeType())))
	}
}

func (n *Node[T]) HasValue() bool {
	return n.value != nil && n.value.Size() > 0
}

func (n *Node[T]) AddValue(val T) {
	if n.value == nil {
		n.value = set3.Empty[T]()
	}
	n.value.Add(val)
}

func (n *Node[T]) RemoveValue(val T) {
	if n.value != nil {
		if n.value.Size() <= 1 {
			n.value = nil
		} else {
			n.value.Remove(val)
		}
	}
}

func (n *Node[T]) GetValues() *set3.Set3[T] {
	if n.value == nil {
		return set3.Empty[T]()
	}
	return n.value.Clone()
}

// Type casting helpers

func (n *Node[T]) asLeaf() *LeafNode[T] {
	if n.GetNodeType() != NodeTypeLeaf {
		panic("node is not of kind Leaf but of kind " + n.GetNodeType().String())
	}
	return (*LeafNode[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode64() *Node64[T] {
	if n.GetNodeType() != NodeType64 {
		panic("node is not of kind Node64 but of kind " + n.GetNodeType().String())
	}
	return (*Node64[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode128() *Node128[T] {
	if n.GetNodeType() != NodeType128 {
		panic("node is not of kind Node128 but of kind " + n.GetNodeType().String())
	}
	return (*Node128[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode256() *Node256[T] {
	if n.GetNodeType() != NodeType256 {
		panic("node is not of kind Node256 but of kind " + n.GetNodeType().String())
	}
	return (*Node256[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode512() *Node512[T] {
	if n.GetNodeType() != NodeType512 {
		panic("node is not of kind Node512 but of kind " + n.GetNodeType().String())
	}
	return (*Node512[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asNode1024() *Node1024[T] {
	if n.GetNodeType() != NodeType1024 {
		panic("node is not of kind Node1024 but of kind " + n.GetNodeType().String())
	}
	return (*Node1024[T])(unsafe.Pointer(n))
}

func (n *Node[T]) asFullNode() *FullNode[T] {
	if n.GetNodeType() != FullNodeType {
		panic("node is not of kind FullNode but of kind " + n.GetNodeType().String())
	}
	return (*FullNode[T])(unsafe.Pointer(n))
}

func (n *LeafNode[T]) asNode() *Node[T] {
	return (*Node[T])(unsafe.Pointer(n))
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
