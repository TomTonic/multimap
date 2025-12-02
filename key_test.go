package multimap

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestFromBytesCopies(t *testing.T) {
	src := []byte{1, 2, 3}
	k := FromBytes(src)
	src[0] = 9
	if bytes.Equal(k.Bytes(), src) {
		t.Fatalf("FromBytes did not copy input: got %v, want original unaffected %v", k.Bytes(), src)
	}
}

func TestFromBytesNilProducesEmpty(t *testing.T) {
	k := FromBytes(nil)
	if !k.IsEmpty() {
		t.Fatalf("FromBytes(nil) expected empty key")
	}
	if got := k.Bytes(); got == nil {
		// k.Bytes returns nil only for nil Key; FromBytes(nil) returns empty slice
		t.Fatalf("FromBytes(nil) expected empty slice, got nil")
	}
}

func TestFromStringNormalization(t *testing.T) {
	// 'ä' can be U+00E4 or 'a' + U+0308
	precomposed := "\u00E4"
	decomposed := "a\u0308"
	p := FromString(precomposed)
	d := FromString(decomposed)
	if !bytes.Equal(p.Bytes(), d.Bytes()) {
		t.Fatalf("normalization mismatch: %v vs %v", p.Bytes(), d.Bytes())
	}
}

func TestIntBigEndianLayouts(t *testing.T) {
	// Verify 64-bit encoding and round-trip decode for signed ints
	const offset = uint64(1) << 63

	v32 := int32(0x01020304)
	k32 := FromInt32(v32)
	if len(k32) != 8 {
		t.Fatalf("FromInt32 should produce 8 bytes, got %d", len(k32))
	}
	// decode by subtracting offset
	got32 := int32(int64(binary.BigEndian.Uint64(k32.Bytes()) - offset))
	if got32 != v32 {
		t.Fatalf("round-trip int32 mismatch: got=%#x want=%#x", got32, v32)
	}

	v64 := int64(0x0102030405060708)
	k64 := FromInt64(v64)
	if len(k64) != 8 {
		t.Fatalf("FromInt64 should produce 8 bytes, got %d", len(k64))
	}
	got64 := int64(binary.BigEndian.Uint64(k64.Bytes()) - offset)
	if got64 != v64 {
		t.Fatalf("round-trip int64 mismatch: got=%#x want=%#x", got64, v64)
	}

	// small widths should match 64-bit encoding for the same value
	if !FromInt32(5).Equal(FromInt64(5)) {
		t.Fatalf("FromInt32 and FromInt64 should produce identical keys for same value")
	}
}

func TestUintBigEndianLayouts(t *testing.T) {
	const offset = uint64(1) << 63
	u16 := uint16(0xABCD)
	k16 := FromUint16(u16)
	if len(k16) != 8 {
		t.Fatalf("FromUint16 should produce 8 bytes, got %d", len(k16))
	}
	// round-trip: read 64-bit then cast
	got16 := uint16(binary.BigEndian.Uint64(k16.Bytes()) - offset)
	if got16 != u16 {
		t.Fatalf("round-trip uint16 mismatch: got=%#x want=%#x", got16, u16)
	}

	u64 := uint64(0x0102030405060708)
	k64 := FromUint64(u64)
	if len(k64) != 8 {
		t.Fatalf("FromUint64 should produce 8 bytes, got %d", len(k64))
	}
	if binary.BigEndian.Uint64(k64.Bytes()) != u64+offset {
		t.Fatalf("FromUint64 produced wrong encoding")
	}

	// small-width unsigned equals 64-bit encoding for same value
	if !FromUint16(0x1234).Equal(FromUint64(0x1234)) {
		t.Fatalf("FromUint16 and FromUint64 should produce identical keys for same value")
	}
}

func TestFromRuneUTF8(t *testing.T) {
	r := '€' // U+20AC, three-byte UTF-8
	k := FromRune(r)
	if !bytes.Equal(k.Bytes(), []byte(string(r))) {
		t.Fatalf("FromRune produced wrong UTF-8: %v", k.Bytes())
	}
}

func TestStringFormatting(t *testing.T) {
	k := FromBytes([]byte{0x01, 0xAB, 0x00})
	if k.String() != "[01,AB,00]" {
		t.Fatalf("String() formatted incorrectly: %s", k.String())
	}
}

func TestEqualAndIsEmpty(t *testing.T) {
	a := FromBytes([]byte{1, 2, 3})
	b := FromBytes([]byte{1, 2, 3})
	c := FromBytes([]byte{1, 2})
	if !a.Equal(b) {
		t.Fatalf("Equal expected true for identical contents")
	}
	if a.Equal(c) {
		t.Fatalf("Equal expected false for different contents")
	}
	if !FromBytes(nil).IsEmpty() || !Key(nil).IsEmpty() {
		// Key(nil).IsEmpty uses len==0 so also true
		t.Fatalf("IsEmpty behavior unexpected")
	}
}

func TestCloneCreatesIndependentCopy(t *testing.T) {
	orig := FromBytes([]byte{1, 2, 3})
	clone := orig.Clone()
	if !orig.Equal(clone) {
		t.Fatalf("clone should be equal to original: orig=%v clone=%v", orig.Bytes(), clone.Bytes())
	}
	// modify clone and ensure orig is unaffected
	cloneBytes := clone.Bytes()
	cloneBytes[0] = 9
	if orig.Bytes()[0] == 9 {
		t.Fatalf("modifying clone affected original: orig=%v clone=%v", orig.Bytes(), cloneBytes)
	}

	// clone of Key(nil) should be nil
	var nk Key = nil
	if nk.Clone() != nil {
		t.Fatalf("Clone of nil Key expected nil")
	}
}

func TestLessThan(t *testing.T) {
	a := FromBytes([]byte{1, 2, 3})
	b := FromBytes([]byte{1, 2, 4})
	if !a.LessThan(b) {
		t.Fatalf("expected %v < %v", a.Bytes(), b.Bytes())
	}
	if b.LessThan(a) {
		t.Fatalf("expected %v not < %v", b.Bytes(), a.Bytes())
	}

	// differing at first byte
	x := FromBytes([]byte{0x00})
	y := FromBytes([]byte{0xFF})
	if !x.LessThan(y) {
		t.Fatalf("expected %v < %v", x.Bytes(), y.Bytes())
	}

	// prefix: shorter is less
	p := FromBytes([]byte{1, 2})
	q := FromBytes([]byte{1, 2, 0})
	if !p.LessThan(q) {
		t.Fatalf("expected prefix %v < %v", p.Bytes(), q.Bytes())
	}

	// equal keys -> not less
	if a.LessThan(a) {
		t.Fatalf("expected %v not < itself", a.Bytes())
	}

	// empty vs non-empty
	var empty Key = nil
	non := FromBytes([]byte{0})
	if !empty.LessThan(non) {
		t.Fatalf("expected empty < non-empty")
	}
	if non.LessThan(empty) {
		t.Fatalf("expected non-empty not < empty")
	}

	// unicode bytes compare by underlying UTF-8
	s1 := FromString("a")
	s2 := FromString("ä")
	// depending on normalization, their byte values should compare; ensure relation is consistent
	if s1.Equal(s2) {
		// unlikely, but if equal then neither is less
		if s1.LessThan(s2) || s2.LessThan(s1) {
			t.Fatalf("unexpected less relation for equal strings")
		}
	} else {
		// exactly one should be less
		if !s1.LessThan(s2) && !s2.LessThan(s1) {
			t.Fatalf("expected one of %v or %v to be less", s1.Bytes(), s2.Bytes())
		}
	}
}

func TestSignedOrderingAcrossWidths(t *testing.T) {
	vals := []int64{-2, -1, 0, 1, 2}
	// ensure that ordering compares numerically regardless of source width
	for i := 0; i < len(vals); i++ {
		for j := 0; j < len(vals); j++ {
			a := FromInt8(int8(vals[i]))
			b := FromInt64(vals[j])
			want := vals[i] < vals[j]
			if a.LessThan(b) != want {
				t.Fatalf("ordering mismatch: %d < %d expected %v", vals[i], vals[j], want)
			}
		}
	}
}

func TestUnsignedWidthConsistency(t *testing.T) {
	u16 := uint16(0xABCD)
	if !FromUint16(u16).Equal(FromUint64(uint64(u16))) {
		t.Fatalf("unsigned widths produced different keys for same numeric value")
	}
	if !FromInt16(int16(-1)).LessThan(FromInt16(int16(0))) {
		t.Fatalf("signed ordering within 16-bit failed")
	}
}

func TestInt64Uint64MixedOrdering(t *testing.T) {
	if !FromInt64(int64(0)).Equal(FromUint64(uint64(0))) {
		t.Fatalf("unsigned and signed int produced different keys for same numeric value")
	}
	if !FromInt64(int64(-1)).LessThan(FromUint64(uint64(0))) {
		t.Fatalf("unsigned and signed int not correctly ordered")
	}
}
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
func TestLessThanOrEqual_Basic(t *testing.T) {
	a := FromBytes([]byte{1, 2, 3})
	b := FromBytes([]byte{1, 2, 4})

	if !a.LessThanOrEqual(b) {
		t.Fatalf("expected %v <= %v", a.Bytes(), b.Bytes())
	}
	if b.LessThanOrEqual(a) {
		t.Fatalf("expected %v not <= %v", b.Bytes(), a.Bytes())
	}
	if !a.LessThanOrEqual(a) {
		t.Fatalf("expected %v <= %v (equal keys)", a.Bytes(), a.Bytes())
	}
}

func TestLessThanOrEqual_PrefixAndLengths(t *testing.T) {
	// shorter prefix is <= longer when bytes equal up to the shorter length
	short := FromBytes([]byte{1, 2})
	long := FromBytes([]byte{1, 2, 0})

	if !short.LessThanOrEqual(long) {
		t.Fatalf("expected prefix %v <= %v", short.Bytes(), long.Bytes())
	}
	if long.LessThanOrEqual(short) {
		t.Fatalf("expected %v not <= %v (longer should not be <= shorter when not equal)", long.Bytes(), short.Bytes())
	}

	// differing at first byte
	x := FromBytes([]byte{0x00})
	y := FromBytes([]byte{0xFF})
	if !x.LessThanOrEqual(y) {
		t.Fatalf("expected %v <= %v", x.Bytes(), y.Bytes())
	}
	if y.LessThanOrEqual(x) {
		t.Fatalf("expected %v not <= %v", y.Bytes(), x.Bytes())
	}
}

func TestLessThanOrEqual_EmptyAndNil(t *testing.T) {
	var nilKey Key = nil
	empty := FromBytes(nil) // empty slice (len 0, non-nil)

	non := FromBytes([]byte{0})

	// nil and empty both have len 0, should be <= each other
	if !nilKey.LessThanOrEqual(empty) || !empty.LessThanOrEqual(nilKey) {
		t.Fatalf("expected nil and empty to be <= each other")
	}

	// empty/nil <= non-empty is true; reverse is false
	if !nilKey.LessThanOrEqual(non) || !empty.LessThanOrEqual(non) {
		t.Fatalf("expected empty/nil <= non-empty")
	}
	if non.LessThanOrEqual(nilKey) || non.LessThanOrEqual(empty) {
		t.Fatalf("expected non-empty not <= empty/nil")
	}

	// both nil
	if !nilKey.LessThanOrEqual(nilKey) {
		t.Fatalf("expected nil <= nil")
	}
}

func TestLessThanOrEqual_ConsistencyWithLessThanAndEqual(t *testing.T) {
	cases := []struct {
		a, b Key
	}{
		{FromBytes([]byte{1, 2, 3}), FromBytes([]byte{1, 2, 3})}, // equal
		{FromBytes([]byte{1, 2, 3}), FromBytes([]byte{1, 2, 4})}, // a < b
		{FromBytes([]byte{1, 2, 4}), FromBytes([]byte{1, 2, 3})}, // a > b
		{FromBytes([]byte{1, 2}), FromBytes([]byte{1, 2, 0})},    // prefix shorter < longer
		{FromBytes([]byte{1, 2, 0}), FromBytes([]byte{1, 2})},    // longer > shorter
		{FromBytes([]byte{}), FromBytes([]byte{0})},              // empty < non-empty
	}

	for _, c := range cases {
		lte := c.a.LessThanOrEqual(c.b)
		lt := c.a.LessThan(c.b)
		eq := c.a.Equal(c.b)
		if lte != (lt || eq) {
			t.Fatalf("inconsistency: a=%v b=%v: <==%v, <||==%v (lt=%v eq=%v)",
				c.a.Bytes(), c.b.Bytes(), lte, lt || eq, lt, eq)
		}
	}
}
