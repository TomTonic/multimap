package art

import (
	"bytes"
	"testing"
)

func TestNode_getNodeType(t *testing.T) {
	tests := []struct {
		name     string
		meta     uint8
		expected NodeType
	}{
		{"NodeTypeLeaf", 0x00, NodeTypeLeaf},
		{"NodeType64", 0x10, NodeType64},
		{"NodeType128", 0x20, NodeType128},
		{"NodeType256", 0x30, NodeType256},
		{"NodeType512", 0x40, NodeType512},
		{"NodeType1024", 0x50, NodeType1024},
		{"FullNodeType", 0x60, FullNodeType},
		{"NodeType with prefix bits set", 0x2F, NodeType128},
		{"Maximum node type value", 0xF0, NodeType(15)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[int]{meta: tt.meta}
			result := node.getNodeType()
			if result != tt.expected {
				t.Errorf("getNodeType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNode_setNodeType(t *testing.T) {
	tests := []struct {
		name        string
		initialMeta uint8
		setType     NodeType
	}{
		{"set to NodeType64 from zero", 0x00, NodeType64},
		{"set to NodeType1024 preserves lower nibble", 0x0F, NodeType1024},
		{"override upper nibble from 0xFF to FullNodeType", 0xFF, FullNodeType},
		{"set to max nibble value", 0x2A, NodeType(15)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[int]{meta: tt.initialMeta}
			ret := node.setNodeType(tt.setType)
			if ret != node {
				t.Fatalf("setNodeType did not return the same node pointer")
			}
			expected := (tt.initialMeta & 0x0F) | (uint8(tt.setType) << 4)
			if node.meta != expected {
				t.Errorf("meta = 0x%02X, want 0x%02X", node.meta, expected)
			}
			if node.getNodeType() != tt.setType {
				t.Errorf("getNodeType() = %v, want %v", node.getNodeType(), tt.setType)
			}
		})
	}
}

func TestNode_getPrefixLen(t *testing.T) {
	tests := []struct {
		name     string
		meta     uint8
		expected uint8
	}{
		{"zero meta yields zero prefix length", 0x00, 0},
		{"lower nibble max (0x0F)", 0x0F, 15},
		{"upper nibble set only (0xF0)", 0xF0, 0},
		{"mixed nibbles (0x2A)", 0x2A, 10},
		{"all bits set (0xFF)", 0xFF, 15},
		{"random upper with different lower (0x9C)", 0x9C, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[int]{meta: tt.meta}
			got := node.getPrefixLen()
			if got != tt.expected {
				t.Fatalf("getPrefixLen() = %d, want %d (meta=0x%02X)", got, tt.expected, tt.meta)
			}
		})
	}
}

func TestNode_getPrefix_empty(t *testing.T) {
	node := &Node[int]{}
	got := node.getPrefix()
	if len(got) != 0 {
		t.Fatalf("expected empty prefix, got len=%d", len(got))
	}
}

func TestNode_getPrefix_copyAndLength(t *testing.T) {
	prefix := []byte{1, 2, 3, 4, 5}
	node := &Node[int]{}
	node.setPrefix(prefix)

	got := node.getPrefix()
	if !bytes.Equal(got, prefix) {
		t.Fatalf("getPrefix() = %v, want %v", got, prefix)
	}

	got[0] = 99
	got2 := node.getPrefix()
	if got2[0] != prefix[0] {
		t.Fatalf("mutation of returned slice affected node localPrefix; got2[0]=%d want=%d", got2[0], prefix[0])
	}
}

func TestNode_getPrefix_longerThanLocalAndRespectsMeta(t *testing.T) {
	long := make([]byte, 256)
	for i := range long {
		long[i] = byte(i)
	}

	node := &Node[int]{}
	node.setPrefix(long)
	l := node.getPrefixLen()

	got := node.getPrefix()
	if len(got) != int(l) {
		t.Fatalf("len(getPrefix()) = %d, want %d", len(got), l)
	}
	if !bytes.Equal(got, long[:l]) {
		t.Fatalf("getPrefix() content mismatch; got=%v want=%v", got, long[:l])
	}

	node.meta = (node.meta & 0xF0) | 3
	trunc := node.getPrefix()
	if len(trunc) != 3 || !bytes.Equal(trunc, long[:3]) {
		t.Fatalf("after manual meta change getPrefix() = %v (len=%d), want %v (len=3)", trunc, len(trunc), long[:3])
	}
}

// New tests for setPrefix

func TestNode_setPrefix_empty(t *testing.T) {
	node := &Node[int]{meta: 0xA0} // upper nibble preset
	ret := node.setPrefix(nil)
	if ret != node {
		t.Fatalf("setPrefix did not return same pointer")
	}
	if node.getPrefixLen() != 0 {
		t.Fatalf("expected prefix len 0, got %d", node.getPrefixLen())
	}
	if node.meta&0xF0 != 0xA0 {
		t.Fatalf("upper nibble modified, meta=0x%02X", node.meta)
	}
	if len(node.getPrefix()) != 0 {
		t.Fatalf("expected empty stored prefix slice")
	}
	// ensure internal buffer bytes are zeroed
	for i := 0; i < maxLocalPrefixLen; i++ {
		if node.localPrefix[i] != 0 {
			t.Fatalf("expected localPrefix[%d] == 0 after empty setPrefix, got %d", i, node.localPrefix[i])
		}
	}
}

func TestNode_setPrefix_basic(t *testing.T) {
	prefix := []byte{7, 8, 9}
	node := &Node[int]{meta: 0x50}
	node.setPrefix(prefix)

	if node.meta&0x0F != uint8(len(prefix)) {
		t.Fatalf("meta lower nibble=%d want=%d", node.meta&0x0F, len(prefix))
	}
	if node.meta&0xF0 != 0x50 {
		t.Fatalf("upper nibble changed; meta=0x%02X", node.meta)
	}
	got := node.getPrefix()
	if !bytes.Equal(got, prefix) {
		t.Fatalf("stored prefix %v want %v", got, prefix)
	}
	// mutation safety
	prefix[0] = 99
	if node.getPrefix()[0] != 7 {
		t.Fatalf("mutation of source slice affected stored prefix")
	}
	// ensure remaining bytes after prefix are zeroed
	for i := len(prefix); i < maxLocalPrefixLen; i++ {
		if node.localPrefix[i] != 0 {
			t.Fatalf("expected localPrefix[%d] == 0 after setPrefix, got %d", i, node.localPrefix[i])
		}
	}
}

func TestNode_setPrefix_truncates(t *testing.T) {
	orig := make([]byte, maxLocalPrefixLen+10)
	for i := range orig {
		orig[i] = byte(i + 1)
	}
	node := &Node[int]{meta: 0xE0}
	node.setPrefix(orig)

	if node.getPrefixLen() != uint8(maxLocalPrefixLen&0x0F) {
		t.Fatalf("prefix len=%d want=%d", node.getPrefixLen(), uint8(maxLocalPrefixLen&0x0F))
	}
	got := node.getPrefix()
	if len(got) != maxLocalPrefixLen {
		t.Fatalf("stored length=%d want=%d", len(got), maxLocalPrefixLen)
	}
	if !bytes.Equal(got, orig[:maxLocalPrefixLen]) {
		t.Fatalf("truncated content mismatch")
	}
	if node.meta&0xF0 != 0xE0 {
		t.Fatalf("upper nibble changed; meta=0x%02X", node.meta)
	}
	// when truncated to maxLocalPrefixLen, there should be no zero-padding
	for i := 0; i < maxLocalPrefixLen; i++ {
		if node.localPrefix[i] == 0 {
			t.Fatalf("expected non-zero byte at localPrefix[%d] after truncating setPrefix", i)
		}
	}
}

func TestNode_setPrefix_exactMax(t *testing.T) {
	orig := make([]byte, maxLocalPrefixLen)
	for i := range orig {
		orig[i] = byte(i)
	}
	node := &Node[int]{meta: 0x20}
	node.setPrefix(orig)

	if node.getPrefixLen() != uint8(maxLocalPrefixLen&0x0F) {
		t.Fatalf("prefix len=%d want=%d", node.getPrefixLen(), uint8(maxLocalPrefixLen&0x0F))
	}
	if !bytes.Equal(node.getPrefix(), orig) {
		t.Fatalf("expected full prefix stored")
	}
	if node.meta&0xF0 != 0x20 {
		t.Fatalf("upper nibble changed; meta=0x%02X", node.meta)
	}
}

func TestNode_setPrefix_overwritesPrevious(t *testing.T) {
	first := []byte{1, 2, 3, 4}
	second := []byte{9, 8}
	node := &Node[int]{}
	node.setPrefix(first)
	node.setPrefix(second)

	if node.getPrefixLen() != uint8(len(second)) {
		t.Fatalf("prefix len=%d want=%d", node.getPrefixLen(), len(second))
	}
	if !bytes.Equal(node.getPrefix(), second) {
		t.Fatalf("expected overwritten prefix %v got %v", second, node.getPrefix())
	}
	// ensure bytes after new shorter prefix are zeroed
	for i := len(second); i < maxLocalPrefixLen; i++ {
		if node.localPrefix[i] != 0 {
			t.Fatalf("expected localPrefix[%d] == 0 after overwrite, got %d", i, node.localPrefix[i])
		}
	}
}

func TestNode_setPrefix_metaLowerNibbleMask(t *testing.T) {
	// Ensure lower nibble is masked correctly (only 0x0F bits kept)
	// Use length > 0x0F if maxLocalPrefixLen allows; fabricate slice accordingly.
	long := make([]byte, maxLocalPrefixLen)
	for i := range long {
		long[i] = byte(i)
	}
	node := &Node[int]{meta: 0x90}
	node.setPrefix(long)
	if (node.meta & 0x0F) != uint8(min(len(long), maxLocalPrefixLen))&0x0F {
		t.Fatalf("meta lower nibble=%d want=%d", node.meta&0x0F, uint8(min(len(long), maxLocalPrefixLen))&0x0F)
	}
	if node.meta&0xF0 != 0x90 {
		t.Fatalf("upper nibble changed; meta=0x%02X", node.meta)
	}
}

func TestNode_getMaxChildren(t *testing.T) {
	tests := []struct {
		name     string
		nodeType NodeType
		expected uint
	}{
		{"NodeTypeLeaf", NodeTypeLeaf, maxChildrenLeaf},
		{"NodeType64", NodeType64, maxChildrenNode64},
		{"NodeType128", NodeType128, maxChildrenNode128},
		{"NodeType256", NodeType256, maxChildrenNode256},
		{"NodeType512", NodeType512, maxChildrenNode512},
		{"NodeType1024", NodeType1024, maxChildrenNode1024},
		{"FullNodeType", FullNodeType, maxChildrenFullNode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[int]{}
			node.setNodeType(tt.nodeType)
			result := node.getMaxChildren()
			if result != tt.expected {
				t.Errorf("getMaxChildren() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNode_getMaxChildren_panic(t *testing.T) {
	node := &Node[int]{meta: 0xF0} // NodeType(15) - invalid type
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("getMaxChildren() did not panic for unknown node type")
		}
	}()
	node.getMaxChildren()
}

// Tests for hasValue, addValue, removeValue

func TestNode_hasValue_emptyNode(t *testing.T) {
	node := &Node[int]{}
	if node.hasValue() {
		t.Fatalf("expected hasValue() = false for new node with nil value")
	}
}

func TestNode_hasValue_afterAddValue(t *testing.T) {
	node := &Node[int]{}
	node.addValue(42)
	if !node.hasValue() {
		t.Fatalf("expected hasValue() = true after addValue")
	}
}

func TestNode_hasValue_emptySet(t *testing.T) {
	node := &Node[int]{}
	node.addValue(10)
	node.removeValue(10)
	// After removing the only value, the set should be nil
	if node.hasValue() {
		t.Fatalf("expected hasValue() = false after removing the only value")
	}
}

func TestNode_addValue_single(t *testing.T) {
	node := &Node[int]{}
	node.addValue(100)

	if node.value == nil {
		t.Fatalf("expected value set to be non-nil after addValue")
	}
	if node.value.Size() != 1 {
		t.Fatalf("expected set size = 1, got %d", node.value.Size())
	}
	if !node.value.Contains(100) {
		t.Fatalf("expected set to contain value 100")
	}
}

func TestNode_addValue_multiple(t *testing.T) {
	node := &Node[int]{}
	values := []int{10, 20, 30, 40}

	for _, v := range values {
		node.addValue(v)
	}

	if node.value.Size() != uint32(len(values)) {
		t.Fatalf("expected set size = %d, got %d", len(values), node.value.Size())
	}

	for _, v := range values {
		if !node.value.Contains(v) {
			t.Fatalf("expected set to contain value %d", v)
		}
	}
}

func TestNode_addValue_duplicate(t *testing.T) {
	node := &Node[int]{}
	node.addValue(50)
	node.addValue(50)
	node.addValue(50)

	// Set should deduplicate
	if node.value.Size() != 1 {
		t.Fatalf("expected set size = 1 after adding duplicates, got %d", node.value.Size())
	}
	if !node.value.Contains(50) {
		t.Fatalf("expected set to contain value 50")
	}
}

func TestNode_addValue_differentTypes(t *testing.T) {
	// Test with string type
	nodeStr := &Node[string]{}
	nodeStr.addValue("hello")
	nodeStr.addValue("world")

	if nodeStr.value.Size() != uint32(2) {
		t.Fatalf("expected string set size = 2, got %d", nodeStr.value.Size())
	}
	if !nodeStr.value.Contains("hello") || !nodeStr.value.Contains("world") {
		t.Fatalf("expected string set to contain 'hello' and 'world'")
	}
}

func TestNode_removeValue_single(t *testing.T) {
	node := &Node[int]{}
	node.addValue(100)
	node.removeValue(100)

	// After removing the only value, the set should be nil
	if node.value != nil {
		t.Fatalf("expected value to be nil after removing the only element")
	}
	if node.hasValue() {
		t.Fatalf("expected hasValue() = false after removing the only value")
	}
}

func TestNode_removeValue_fromMultiple(t *testing.T) {
	node := &Node[int]{}
	node.addValue(10)
	node.addValue(20)
	node.addValue(30)

	node.removeValue(20)

	if node.value.Size() != uint32(2) {
		t.Fatalf("expected set size = 2 after removal, got %d", node.value.Size())
	}
	if node.value.Contains(20) {
		t.Fatalf("expected value 20 to be removed from set")
	}
	if !node.value.Contains(10) || !node.value.Contains(30) {
		t.Fatalf("expected values 10 and 30 to remain in set")
	}
	if !node.hasValue() {
		t.Fatalf("expected hasValue() = true after partial removal")
	}
}

func TestNode_removeValue_nonExistent(t *testing.T) {
	node := &Node[int]{}
	node.addValue(100)
	node.addValue(200) // Add two values to ensure size > 1

	// Remove a value that doesn't exist
	node.removeValue(999)

	// Should not affect the existing values
	if node.value == nil {
		t.Fatalf("expected value set to remain non-nil after removing non-existent value")
	}
	if node.value.Size() != uint32(2) {
		t.Fatalf("expected set size = 2 after removing non-existent value, got %d", node.value.Size())
	}
	if !node.value.Contains(100) || !node.value.Contains(200) {
		t.Fatalf("expected values 100 and 200 to remain in set")
	}
}

func TestNode_removeValue_fromEmpty(t *testing.T) {
	node := &Node[int]{}

	// Should not panic
	node.removeValue(42)

	if node.value != nil {
		t.Fatalf("expected value to remain nil after removing from empty node")
	}
}

func TestNode_removeValue_afterClear(t *testing.T) {
	node := &Node[int]{}
	node.addValue(1)
	node.addValue(2)

	// Remove all values one by one
	node.removeValue(1)
	if node.value.Size() != uint32(1) {
		t.Fatalf("expected size 1 after first removal, got %d", node.value.Size())
	}

	node.removeValue(2)
	// After removing the last value, set should be nil
	if node.value != nil {
		t.Fatalf("expected value to be nil after removing all values")
	}
}

func TestNode_addRemoveSequence(t *testing.T) {
	node := &Node[int]{}

	// Add, remove, add again
	node.addValue(99)
	if !node.hasValue() {
		t.Fatalf("expected hasValue() = true after first add")
	}

	node.removeValue(99)
	if node.hasValue() {
		t.Fatalf("expected hasValue() = false after remove")
	}

	node.addValue(88)
	if !node.hasValue() {
		t.Fatalf("expected hasValue() = true after second add")
	}
	if !node.value.Contains(88) {
		t.Fatalf("expected set to contain value 88")
	}
	if node.value.Contains(99) {
		t.Fatalf("expected old value 99 to not be in set")
	}
}

// Round-trip casting tests for all node casting helpers
func TestCastingHelpers_RoundTrip(t *testing.T) {
	// LeafNode
	leaf := &LeafNode[int]{}
	leaf.setNodeType(NodeTypeLeaf)
	n := leaf.asNode()
	if got := n.asLeaf(); got != leaf {
		t.Fatalf("LeafNode round-trip failed: got %p want %p", got, leaf)
	}

	// Node64
	n64 := &Node64[int]{}
	n64.setNodeType(NodeType64)
	n = n64.asNode()
	if got := n.asNode64(); got != n64 {
		t.Fatalf("Node64 round-trip failed: got %p want %p", got, n64)
	}

	// Node128
	n128 := &Node128[int]{}
	n128.setNodeType(NodeType128)
	n = n128.asNode()
	if got := n.asNode128(); got != n128 {
		t.Fatalf("Node128 round-trip failed: got %p want %p", got, n128)
	}

	// Node256
	n256 := &Node256[int]{}
	n256.setNodeType(NodeType256)
	n = n256.asNode()
	if got := n.asNode256(); got != n256 {
		t.Fatalf("Node256 round-trip failed: got %p want %p", got, n256)
	}

	// Node512
	n512 := &Node512[int]{}
	n512.setNodeType(NodeType512)
	n = n512.asNode()
	if got := n.asNode512(); got != n512 {
		t.Fatalf("Node512 round-trip failed: got %p want %p", got, n512)
	}

	// Node1024
	n1024 := &Node1024[int]{}
	n1024.setNodeType(NodeType1024)
	n = n1024.asNode()
	if got := n.asNode1024(); got != n1024 {
		t.Fatalf("Node1024 round-trip failed: got %p want %p", got, n1024)
	}

	// FullNode
	fn := &FullNode[int]{}
	fn.setNodeType(FullNodeType)
	n = fn.asNode()
	if got := n.asFullNode(); got != fn {
		t.Fatalf("FullNode round-trip failed: got %p want %p", got, fn)
	}
}

func TestCastingHelpers_PanicOnWrongType(t *testing.T) {
	tests := []struct {
		name string
		meta uint8
		fn   func(n *Node[int])
	}{
		{"asLeaf_on_Node64", uint8(NodeType64) << 4, func(n *Node[int]) { _ = n.asLeaf() }},
		{"asNode64_on_Leaf", uint8(NodeTypeLeaf) << 4, func(n *Node[int]) { _ = n.asNode64() }},
		{"asNode128_on_Node64", uint8(NodeType64) << 4, func(n *Node[int]) { _ = n.asNode128() }},
		{"asNode256_on_Node128", uint8(NodeType128) << 4, func(n *Node[int]) { _ = n.asNode256() }},
		{"asNode512_on_Node256", uint8(NodeType256) << 4, func(n *Node[int]) { _ = n.asNode512() }},
		{"asNode1024_on_Node512", uint8(NodeType512) << 4, func(n *Node[int]) { _ = n.asNode1024() }},
		{"asFullNode_on_Node1024", uint8(NodeType1024) << 4, func(n *Node[int]) { _ = n.asFullNode() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node[int]{meta: tt.meta}
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected panic for %s but none occurred", tt.name)
				}
			}()
			tt.fn(node)
		})
	}
}
