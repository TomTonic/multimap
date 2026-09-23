package multimap

import (
	"testing"
)

func TestNode5_GetChild(t *testing.T) {
	// Helper to create a Key from string
	key := func(s string) Key {
		return Key(s)
	}

	t.Run("exact match returns node", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "abc")

		result := n.getChild(key(""), key("abc"))
		if result != n.asNode() {
			t.Error("Expected exact match to return the node itself")
		}
	})

	t.Run("search key is prefix of local key returns nil", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 5
		copy(n.prefix[:], "abcde")

		result := n.getChild(key(""), key("abc"))
		if result != nil {
			t.Error("Expected nil when search key is prefix of local key")
		}
	})

	t.Run("keys differ returns nil", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "abc")

		result := n.getChild(key(""), key("def"))
		if result != nil {
			t.Error("Expected nil when keys differ")
		}
	})

	t.Run("finds matching child", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "abc")

		child := &nodeLeaf[string]{}
		child.ntype = Leaf
		child.prefixLen = 1
		copy(child.prefix[:], "e")

		n.numChildren = 1
		n.firstKeyByteChild[0] = byte('d')
		n.child[0] = child.asNode()

		result := n.getChild(key(""), key("abcde"))
		if result == nil {
			t.Error("Expected to find matching child")
		}
	})

	t.Run("no matching child returns nil", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "abc")

		child := &nodeLeaf[string]{}
		child.ntype = Leaf
		n.numChildren = 1
		n.firstKeyByteChild[0] = byte('d')
		n.child[0] = child.asNode()

		result := n.getChild(key(""), key("abcx"))
		if result != nil {
			t.Error("Expected nil when no matching child found")
		}
	})

	t.Run("with current prefix", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "def")

		result := n.getChild(key("abc"), key("abcdef"))
		if result != n.asNode() {
			t.Error("Expected exact match with current prefix")
		}
	})

	t.Run("nil child pointer skipped", func(t *testing.T) {
		n := &node5[string]{}
		n.ntype = Node5
		n.prefixLen = 3
		copy(n.prefix[:], "abc")

		n.numChildren = 2
		n.firstKeyByteChild[0] = byte('d')
		n.child[0] = nil
		n.firstKeyByteChild[1] = byte('e')

		child := &nodeLeaf[string]{}
		child.ntype = Leaf
		n.child[1] = child.asNode()

		result := n.getChild(key(""), key("abcd"))
		if result != nil {
			t.Error("Expected nil when child pointer is nil")
		}
	})
}
