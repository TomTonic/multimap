package multimap

import "unsafe"

const (
	// maxPrefixLen is maximum prefix length for internal nodes.
	maxPrefixLen = 12
)

// The type prefix used in the node to store the key prefix.
type prefix [maxPrefixLen]byte

type nodeType byte

func (k nodeType) String() string {
	return []string{"Abstract", "Leaf", "Node5", "Node51", "Node256"}[k+1]
}

const (
	Leaf    nodeType = 0
	Node5   nodeType = 1
	Node51  nodeType = 2
	Node256 nodeType = 3
)

// node is the base struct for all node types.
// it contains the common fields for all nodeX types.
type node[T comparable] struct {
	ntype     nodeType
	prefixLen byte   // length of the prefix
	prefix    prefix // prefix of the node
}

func (n *node[T]) isLeaf() bool {
	return n.ntype == Leaf
}

func (n *node[T]) getKind() nodeType {
	return n.ntype
}

func (n *node[T]) maxChildCount() uint32 {
	switch n.ntype {
	case Leaf:
		return 1
	case Node5:
		return 5
	case Node51:
		return 51
	case Node256:
		return 256
	default:
		return 0
	}
}

func (n *node[T]) asNode5() *node5[T] {
	if n.ntype != Node5 {
		panic("node is not of kind Node5 but of kind " + n.ntype.String())
	}
	return (*node5[T])(unsafe.Pointer(n))
}
func (n *node[T]) asNode51() *node51[T] {
	if n.ntype != Node51 {
		panic("node is not of kind Node51 but of kind " + n.ntype.String())
	}
	return (*node51[T])(unsafe.Pointer(n))
}
func (n *node[T]) asNode256() *node256[T] {
	if n.ntype != Node256 {
		panic("node is not of kind Node256")
	}
	return (*node256[T])(unsafe.Pointer(n))
}
func (n *node[T]) asLeaf() *nodeLeaf[T] {
	if n.ntype != Leaf {
		panic("node is not of kind Leaf")
	}
	return (*nodeLeaf[T])(unsafe.Pointer(n))
}

func (n *node[T]) getLocalKey(parentKey Key) Key {
	result := make([]byte, len(parentKey)+int(n.prefixLen))
	copy(result, parentKey)
	copy(result[len(parentKey):], n.prefix[:n.prefixLen])
	return result
}

type nodeOps[T comparable] interface {
	currentChildCount() uint32
	hasCapacityForChild() bool
	grow() *node[T]

	isReadyToShrink() bool
	shrink() *node[T]
}
