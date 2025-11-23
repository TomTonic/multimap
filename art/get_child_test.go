package art

import (
	"testing"

	mm "github.com/TomTonic/multimap"
)

// Test LeafNode.getChild scenarios

func TestLeafNode_getChild_exactMatch(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{1, 2, 3})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{1, 2, 3})

	result := leaf.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != leaf.asNode() {
		t.Fatalf("expected result to be the leaf node itself")
	}
}

func TestLeafNode_getChild_noMatch_differentKey(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{1, 2, 3})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{1, 2, 4}) // differs at position 2

	result := leaf.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected no match for different key, got %v", result)
	}
}

func TestLeafNode_getChild_searchKeyIsTruePrefix(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{1, 2, 3})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{1, 2}) // searchKey is prefix of lokalKey

	result := leaf.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when searchKey is true prefix of lokalKey, got %v", result)
	}
}

func TestLeafNode_getChild_delegateToChild(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{1, 2})

	childLeaf := &LeafNode[int]{}
	childLeaf.setNodeType(NodeTypeLeaf)
	childLeaf.setPrefix([]byte{3, 4})

	leaf.child = childLeaf.asNode()

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{1, 2, 3, 4}) // lokalKey of leaf is [1,2], continues to child

	result := leaf.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected delegation to child to succeed")
	}
	if result != childLeaf.asNode() {
		t.Fatalf("expected result to be the child node")
	}
}

func TestLeafNode_getChild_noChildForLongerKey(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{1, 2})
	// no child set

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{1, 2, 3}) // lokalKey is prefix of searchKey but no child

	result := leaf.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when no child for longer key, got %v", result)
	}
}

// Test Node64.getChild scenarios

func TestNode64_getChild_exactMatch(t *testing.T) {
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n64.setPrefix([]byte{5, 6})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{5, 6})

	result := n64.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != n64.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode64_getChild_noMatch_differentPrefix(t *testing.T) {
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n64.setPrefix([]byte{5, 6})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{5, 7}) // differs in prefix

	result := n64.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil for different prefix, got %v", result)
	}
}

func TestNode64_getChild_findChildByFirstKeyByte(t *testing.T) {
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n64.setPrefix([]byte{10})
	n64.numChildren = 2
	n64.firstKeyByte[0] = 20
	n64.firstKeyByte[1] = 30

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	// Child prefix must start with firstKeyByte value
	childNode.setPrefix([]byte{20})

	n64.child[0] = childNode

	currentPrefix := mm.Key([]byte{})
	// searchKey: parent prefix [10] + firstKeyByte[0]+child prefix [20]
	searchKey := mm.Key([]byte{10, 20})

	result := n64.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via firstKeyByte, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

func TestNode64_getChild_noMatchingChild(t *testing.T) {
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n64.setPrefix([]byte{10})
	n64.numChildren = 1
	n64.firstKeyByte[0] = 20

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{10, 99}) // no child for byte 99

	result := n64.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when no matching child, got %v", result)
	}
}

// Test Node128.getChild scenarios (similar structure to Node64)

func TestNode128_getChild_exactMatch(t *testing.T) {
	n128 := &Node128[int]{}
	n128.setNodeType(NodeType128)
	n128.setPrefix([]byte{7, 8})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{7, 8})

	result := n128.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != n128.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode128_getChild_findChildByFirstKeyByte(t *testing.T) {
	n128 := &Node128[int]{}
	n128.setNodeType(NodeType128)
	n128.setPrefix([]byte{11})
	n128.numChildren = 3
	n128.firstKeyByte[0] = 40
	n128.firstKeyByte[1] = 50
	n128.firstKeyByte[2] = 60

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	childNode.setPrefix([]byte{50}) // child prefix starts with firstKeyByte

	n128.child[1] = childNode

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{11, 50}) // matches firstKeyByte[1]=50

	result := n128.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via firstKeyByte, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

// Test Node256.getChild scenarios (similar structure to Node64/Node128)

func TestNode256_getChild_exactMatch(t *testing.T) {
	n256 := &Node256[int]{}
	n256.setNodeType(NodeType256)
	n256.setPrefix([]byte{9, 10})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{9, 10})

	result := n256.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != n256.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode256_getChild_findChildByFirstKeyByte(t *testing.T) {
	n256 := &Node256[int]{}
	n256.setNodeType(NodeType256)
	n256.setPrefix([]byte{12})
	n256.numChildren = 2
	n256.firstKeyByte[0] = 70
	n256.firstKeyByte[1] = 80

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	childNode.setPrefix([]byte{80}) // child prefix starts with firstKeyByte

	n256.child[1] = childNode

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{12, 80}) // matches firstKeyByte[1]=80

	result := n256.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via firstKeyByte, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

// Test Node512.getChild scenarios (with bitmap and binary search)

func TestNode512_getChild_exactMatch(t *testing.T) {
	n512 := &Node512[int]{}
	n512.setNodeType(NodeType512)
	n512.setPrefix([]byte{13, 14})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{13, 14})

	result := n512.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != n512.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode512_getChild_bitmapNotSet(t *testing.T) {
	n512 := &Node512[int]{}
	n512.setNodeType(NodeType512)
	n512.setPrefix([]byte{15})
	// bitmap is empty, no children

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{15, 99}) // no child for byte 99

	result := n512.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when bitmap not set, got %v", result)
	}
}

func TestNode512_getChild_findChildViaBitmapAndBinarySearch(t *testing.T) {
	n512 := &Node512[int]{}
	n512.setNodeType(NodeType512)
	n512.setPrefix([]byte{16})
	n512.numChildren = 3
	n512.firstKeyByte[0] = 100
	n512.firstKeyByte[1] = 110
	n512.firstKeyByte[2] = 120
	n512.bitmap.Set(100)
	n512.bitmap.Set(110)
	n512.bitmap.Set(120)

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	childNode.setPrefix([]byte{110}) // child prefix starts with firstKeyByte

	n512.child[1] = childNode

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{16, 110}) // matches sorted firstKeyByte[1]=110

	result := n512.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via bitmap and binary search, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

func TestNode512_binarySearchKeyByte_found(t *testing.T) {
	n512 := &Node512[int]{}
	n512.numChildren = 5
	n512.firstKeyByte[0] = 10
	n512.firstKeyByte[1] = 20
	n512.firstKeyByte[2] = 30
	n512.firstKeyByte[3] = 40
	n512.firstKeyByte[4] = 50

	pos := n512.binarySearchKeyByte(30)
	if pos != 2 {
		t.Fatalf("expected position 2, got %d", pos)
	}
}

func TestNode512_binarySearchKeyByte_notFound(t *testing.T) {
	n512 := &Node512[int]{}
	n512.numChildren = 3
	n512.firstKeyByte[0] = 10
	n512.firstKeyByte[1] = 20
	n512.firstKeyByte[2] = 30

	pos := n512.binarySearchKeyByte(25)
	if pos != -1 {
		t.Fatalf("expected -1 for not found, got %d", pos)
	}
}

// Test Node1024.getChild scenarios (similar to Node512)

func TestNode1024_getChild_exactMatch(t *testing.T) {
	n1024 := &Node1024[int]{}
	n1024.setNodeType(NodeType1024)
	n1024.setPrefix([]byte{17, 18})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{17, 18})

	result := n1024.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != n1024.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode1024_getChild_bitmapNotSet(t *testing.T) {
	n1024 := &Node1024[int]{}
	n1024.setNodeType(NodeType1024)
	n1024.setPrefix([]byte{19})
	// bitmap is empty, no children

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{19, 99}) // no child for byte 99

	result := n1024.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when bitmap not set, got %v", result)
	}
}

func TestNode1024_getChild_findChildViaBitmapAndBinarySearch(t *testing.T) {
	n1024 := &Node1024[int]{}
	n1024.setNodeType(NodeType1024)
	n1024.setPrefix([]byte{21})
	n1024.numChildren = 4
	n1024.firstKeyByte[0] = 150
	n1024.firstKeyByte[1] = 160
	n1024.firstKeyByte[2] = 170
	n1024.firstKeyByte[3] = 180
	n1024.bitmap.Set(150)
	n1024.bitmap.Set(160)
	n1024.bitmap.Set(170)
	n1024.bitmap.Set(180)

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	childNode.setPrefix([]byte{170}) // child prefix starts with firstKeyByte

	n1024.child[2] = childNode

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{21, 170}) // matches sorted firstKeyByte[2]=170

	result := n1024.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via bitmap and binary search, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

func TestNode1024_binarySearchKeyByte_found(t *testing.T) {
	n1024 := &Node1024[int]{}
	n1024.numChildren = 6
	n1024.firstKeyByte[0] = 10
	n1024.firstKeyByte[1] = 20
	n1024.firstKeyByte[2] = 30
	n1024.firstKeyByte[3] = 40
	n1024.firstKeyByte[4] = 50
	n1024.firstKeyByte[5] = 60

	pos := n1024.binarySearchKeyByte(40)
	if pos != 3 {
		t.Fatalf("expected position 3, got %d", pos)
	}
}

func TestNode1024_binarySearchKeyByte_notFound(t *testing.T) {
	n1024 := &Node1024[int]{}
	n1024.numChildren = 3
	n1024.firstKeyByte[0] = 10
	n1024.firstKeyByte[1] = 20
	n1024.firstKeyByte[2] = 30

	pos := n1024.binarySearchKeyByte(15)
	if pos != -1 {
		t.Fatalf("expected -1 for not found, got %d", pos)
	}
}

// Test FullNode.getChild scenarios (direct index via bitmap)

func TestFullNode_getChild_exactMatch(t *testing.T) {
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	fn.setPrefix([]byte{22, 23})

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{22, 23})

	result := fn.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected exact match to return the node, got nil")
	}
	if result != fn.asNode() {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestFullNode_getChild_bitmapNotSet(t *testing.T) {
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	fn.setPrefix([]byte{24})
	fn.child = &[256]*Node[int]{} // allocate external array
	// bitmap is empty, no children

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{24, 99}) // no child for byte 99

	result := fn.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when bitmap not set, got %v", result)
	}
}

func TestFullNode_getChild_directIndex(t *testing.T) {
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	fn.setPrefix([]byte{25})
	fn.child = &[256]*Node[int]{} // allocate external array
	fn.bitmap.Set(200)

	childLeaf := &LeafNode[int]{}
	childNode := childLeaf.asNode()
	childNode.setNodeType(NodeTypeLeaf)
	childNode.setPrefix([]byte{200}) // child prefix starts with indexed byte

	fn.child[200] = childNode

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{25, 200}) // direct index via nextKeyByte=200

	result := fn.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected to find child via direct index, got nil")
	}
	if result != childNode {
		t.Fatalf("expected result to be the child leaf")
	}
}

func TestFullNode_getChild_bitmapSetButChildNil(t *testing.T) {
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	fn.setPrefix([]byte{26})
	fn.child = &[256]*Node[int]{} // allocate external array
	fn.bitmap.Set(100)
	// child[100] is nil

	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{26, 100, 101})

	result := fn.getChild(currentPrefix, searchKey)
	if result != nil {
		t.Fatalf("expected nil when child pointer is nil, got %v", result)
	}
}

// Test Node.getChild dispatch

func TestNode_getChild_dispatchToLeaf(t *testing.T) {
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	leaf.setPrefix([]byte{30})

	n := leaf.asNode()
	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{30})

	result := n.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected dispatch to LeafNode to succeed")
	}
	if result != n {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode_getChild_dispatchToNode64(t *testing.T) {
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n64.setPrefix([]byte{31})

	n := n64.asNode()
	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{31})

	result := n.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected dispatch to Node64 to succeed")
	}
	if result != n {
		t.Fatalf("expected result to be the node itself")
	}
}

func TestNode_getChild_dispatchToFullNode(t *testing.T) {
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	fn.setPrefix([]byte{32})

	n := fn.asNode()
	currentPrefix := mm.Key([]byte{})
	searchKey := mm.Key([]byte{32})

	result := n.getChild(currentPrefix, searchKey)
	if result == nil {
		t.Fatalf("expected dispatch to FullNode to succeed")
	}
	if result != n {
		t.Fatalf("expected result to be the node itself")
	}
}
