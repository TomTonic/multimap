package art

import (
	"testing"

	"github.com/TomTonic/multimap/internal/page"
)

// TestPageTypesAreThoseOfTheTree pins the type bytes that internal/page and the tree must agree on.
//
// The page knows its form from its type byte (one-key types before the many-key types) and the tree finds
// pages by the ranges of the type bytes; a change on one side only would make the tree take a page for a node.
//
// Expected: the many-key pages begin at the tree's kMultiKey, the one-key pages at the first single-key page
// (the value overflow has the type before them), so the tree's ranges are the package's.
func TestPageTypesAreThoseOfTheTree(t *testing.T) {
	if page.TypeManyKeys != uint8(kMultiKey) || page.TypeOneKey != uint8(kValueOverflow)+2 {
		t.Fatalf("page: one-key types from %d, many-key types from %d; tree: single-key pages from %d, multi-key pages from %d", page.TypeOneKey, page.TypeManyKeys, kValueOverflow+2, kMultiKey)
	}
	if page.Classes != 6 || uint8(kLastMultiKey) != page.TypeManyKeys+2*5 || uint8(kLastSingleKey) != page.TypeOneKey+2*5 {
		t.Fatalf("the six classes of both forms do not end at the tree's last types: %d %d", kLastSingleKey, kLastMultiKey)
	}
}
