package multimap

import (
	"bytes"
	"testing"
)

func TestNode_AppendLocalPrefix(t *testing.T) {
	tests := []struct {
		name      string
		prefixLen byte
		prefix    [maxPrefixLen]byte
		parentKey Key
		expected  Key
	}{
		{
			name:      "empty parent key and empty prefix",
			prefixLen: 0,
			prefix:    [maxPrefixLen]byte{},
			parentKey: Key{},
			expected:  Key{},
		},
		{
			name:      "empty parent key with prefix",
			prefixLen: 3,
			prefix:    [maxPrefixLen]byte{'a', 'b', 'c'},
			parentKey: Key{},
			expected:  Key{'a', 'b', 'c'},
		},
		{
			name:      "parent key with empty prefix",
			prefixLen: 0,
			prefix:    [maxPrefixLen]byte{},
			parentKey: Key{'x', 'y', 'z'},
			expected:  Key{'x', 'y', 'z'},
		},
		{
			name:      "parent key with prefix",
			prefixLen: 2,
			prefix:    [maxPrefixLen]byte{'a', 'b'},
			parentKey: Key{'x', 'y'},
			expected:  Key{'x', 'y', 'a', 'b'},
		},
		{
			name:      "max prefix length",
			prefixLen: maxPrefixLen,
			prefix:    [maxPrefixLen]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'},
			parentKey: Key{'x'},
			expected:  Key{'x', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'},
		},
		{
			name:      "single byte parent and prefix",
			prefixLen: 1,
			prefix:    [maxPrefixLen]byte{'z'},
			parentKey: Key{'a'},
			expected:  Key{'a', 'z'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node[int]{
				prefixLen: tt.prefixLen,
				prefix:    tt.prefix,
			}

			result := n.appendLocalPrefix(tt.parentKey)

			if !bytes.Equal(result, tt.expected) {
				t.Errorf("appendLocalPrefix() = %v, want %v", result, tt.expected)
			}

			// Verify the result length is correct
			expectedLen := len(tt.parentKey) + int(tt.prefixLen)
			if len(result) != expectedLen {
				t.Errorf("appendLocalPrefix() length = %d, want %d", len(result), expectedLen)
			}
		})
	}
}
