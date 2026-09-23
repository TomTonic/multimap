package multimap

import (
	"bytes"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// identical keys
	a := FromBytes([]byte{1, 2, 3, 4})
	b := FromBytes([]byte{1, 2, 3, 4})
	if got := LongestCommonPrefix(a, b); got != 4 {
		t.Fatalf("identical keys: got %d, want 4", got)
	}

	// partial match
	c := FromBytes([]byte{1, 2, 5, 6})
	if got := LongestCommonPrefix(a, c); got != 2 {
		t.Fatalf("partial match: got %d, want 2", got)
	}

	// no common prefix
	d := FromBytes([]byte{9, 8, 7})
	if got := LongestCommonPrefix(a, d); got != 0 {
		t.Fatalf("no common prefix: got %d, want 0", got)
	}

	// different lengths, shorter first
	e := FromBytes([]byte{1, 2})
	f := FromBytes([]byte{1, 2, 3, 4})
	if got := LongestCommonPrefix(e, f); got != 2 {
		t.Fatalf("different lengths (shorter first): got %d, want 2", got)
	}

	// different lengths, longer first
	if got := LongestCommonPrefix(f, e); got != 2 {
		t.Fatalf("different lengths (longer first): got %d, want 2", got)
	}

	// one empty key
	empty := FromBytes([]byte{})
	if got := LongestCommonPrefix(empty, a); got != 0 {
		t.Fatalf("empty vs non-empty: got %d, want 0", got)
	}
	if got := LongestCommonPrefix(a, empty); got != 0 {
		t.Fatalf("non-empty vs empty: got %d, want 0", got)
	}

	// both empty
	if got := LongestCommonPrefix(empty, empty); got != 0 {
		t.Fatalf("both empty: got %d, want 0", got)
	}

	// nil keys
	var nilKey Key = nil
	if got := LongestCommonPrefix(nilKey, a); got != 0 {
		t.Fatalf("nil vs non-empty: got %d, want 0", got)
	}
	if got := LongestCommonPrefix(nilKey, nilKey); got != 0 {
		t.Fatalf("both nil: got %d, want 0", got)
	}

	// single byte match
	g := FromBytes([]byte{5})
	h := FromBytes([]byte{5, 6, 7})
	if got := LongestCommonPrefix(g, h); got != 1 {
		t.Fatalf("single byte match: got %d, want 1", got)
	}

	// first byte differs
	i := FromBytes([]byte{1, 2, 3})
	j := FromBytes([]byte{2, 2, 3})
	if got := LongestCommonPrefix(i, j); got != 0 {
		t.Fatalf("first byte differs: got %d, want 0", got)
	}
}

func TestAppendInPlace(t *testing.T) {
	// Test basic append functionality
	k := FromBytes([]byte{1, 2, 3})
	toAppend := FromBytes([]byte{4, 5})
	k.append(toAppend)

	expected := []byte{1, 2, 3, 4, 5}
	if !bytes.Equal(k.Bytes(), expected) {
		t.Fatalf("appendInPlace failed: got %v, want %v", k.Bytes(), expected)
	}

	// Test appending to empty key
	var empty Key
	appendData := FromBytes([]byte{10, 20})
	empty.append(appendData)

	if !bytes.Equal(empty.Bytes(), []byte{10, 20}) {
		t.Fatalf("appendInPlace to empty key failed: got %v, want [10 20]", empty.Bytes())
	}

	// Test appending empty key
	k2 := FromBytes([]byte{7, 8, 9})
	emptyAppend := Key{}
	k2.append(emptyAppend)

	if !bytes.Equal(k2.Bytes(), []byte{7, 8, 9}) {
		t.Fatalf("appendInPlace empty key failed: got %v, want [7 8 9]", k2.Bytes())
	}

	// Test appending nil key
	k3 := FromBytes([]byte{1, 2})
	var nilKey Key
	k3.append(nilKey)

	if !bytes.Equal(k3.Bytes(), []byte{1, 2}) {
		t.Fatalf("appendInPlace nil key failed: got %v, want [1 2]", k3.Bytes())
	}

	// Test multiple appends
	k4 := FromBytes([]byte{1})
	k4.append(FromBytes([]byte{2}))
	k4.append(FromBytes([]byte{3, 4}))

	expected2 := []byte{1, 2, 3, 4}
	if !bytes.Equal(k4.Bytes(), expected2) {
		t.Fatalf("multiple appendInPlace failed: got %v, want %v", k4.Bytes(), expected2)
	}

	// Test that original appended key is not affected
	original := FromBytes([]byte{100, 200})
	target := FromBytes([]byte{1, 2})
	target.append(original)

	// Modify original to ensure independence
	originalBytes := original.Bytes()
	originalBytes[0] = 255

	expectedTarget := []byte{1, 2, 100, 200}
	if !bytes.Equal(target.Bytes(), expectedTarget) {
		t.Fatalf("appendInPlace should not be affected by original modification: got %v, want %v", target.Bytes(), expectedTarget)
	}
}

func TestLongestCommonPrefixAsKey_Basic(t *testing.T) {
	t.Run("IdenticalKeys_ReturnsCopy", func(t *testing.T) {
		a := FromBytes([]byte{1, 2, 3, 4})
		b := FromBytes([]byte{1, 2, 3, 4})
		got := LongestCommonPrefixAsKey(a, b)

		if got == nil {
			t.Fatalf("expected non-nil Key for identical non-empty inputs")
		}
		if !bytes.Equal(got.Bytes(), a.Bytes()) {
			t.Fatalf("unexpected lcp: got %v want %v", got.Bytes(), a.Bytes())
		}

		// ensure returned key is an independent copy
		gb := got.Bytes()
		gb[0] = 9
		if a.Bytes()[0] == 9 {
			t.Fatalf("modifying LCP result affected original: original=%v got=%v", a.Bytes(), gb)
		}
		if b.Bytes()[0] == 9 {
			t.Fatalf("modifying LCP result affected original: original=%v got=%v", b.Bytes(), gb)
		}
	})

	t.Run("PartialMatch_ReturnsPrefixCopy", func(t *testing.T) {
		a := FromBytes([]byte{1, 2, 3})
		b := FromBytes([]byte{1, 2, 9})
		got := LongestCommonPrefixAsKey(a, b)

		want := []byte{1, 2}
		if got == nil {
			t.Fatalf("expected non-nil Key for partial match")
		}
		if !bytes.Equal(got.Bytes(), want) {
			t.Fatalf("unexpected lcp: got %v want %v", got.Bytes(), want)
		}

		// independence: modifying returned should not affect a
		gb := got.Bytes()
		gb[1] = 99
		if a.Bytes()[1] == 99 {
			t.Fatalf("modifying LCP result affected original: original=%v got=%v", a.Bytes(), gb)
		}
	})

	t.Run("NoCommonPrefix_ReturnsEmptyNonNil", func(t *testing.T) {
		a := FromBytes([]byte{1})
		b := FromBytes([]byte{2})
		got := LongestCommonPrefixAsKey(a, b)

		if got == nil {
			t.Fatalf("expected non-nil empty Key when no common prefix")
		}
		if !got.IsEmpty() {
			t.Fatalf("expected empty Key when no common prefix, got len=%d", len(got))
		}
	})

	t.Run("NilAndEmptyInputs", func(t *testing.T) {
		var nilKey Key = nil
		non := FromBytes([]byte{1})

		// nil vs non-nil
		got1 := LongestCommonPrefixAsKey(nilKey, non)
		if got1 == nil || !got1.IsEmpty() {
			t.Fatalf("expected non-nil empty Key for nil vs non-empty, got %#v", got1)
		}

		// both nil
		got2 := LongestCommonPrefixAsKey(nilKey, nilKey)
		if got2 == nil || !got2.IsEmpty() {
			t.Fatalf("expected non-nil empty Key for nil vs nil, got %#v", got2)
		}

		// both empty (non-nil empty slices)
		empty := FromBytes([]byte{})
		got3 := LongestCommonPrefixAsKey(empty, empty)
		if got3 == nil || !got3.IsEmpty() {
			t.Fatalf("expected non-nil empty Key for empty vs empty, got %#v", got3)
		}
	})
}
