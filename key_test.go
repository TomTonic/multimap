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
	// Int32
	v32 := int32(0x01020304)
	k32 := FromInt32(v32)
	if len(k32) != 4 || k32[0] != 0x01 || k32[3] != 0x04 {
		t.Fatalf("FromInt32 produced wrong layout: %v", k32)
	}

	// Int64
	v64 := int64(0x0102030405060708)
	k64 := FromInt64(v64)
	if len(k64) != 8 || k64[0] != 0x01 || k64[7] != 0x08 {
		t.Fatalf("FromInt64 produced wrong layout: %v", k64)
	}

	// Check round-trip using binary.BigEndian
	if got := int64(binary.BigEndian.Uint64(k64.Bytes())); got != v64 {
		t.Fatalf("round-trip int64 mismatch: got=%#x want=%#x", got, v64)
	}
}

func TestUintBigEndianLayouts(t *testing.T) {
	u16 := uint16(0xABCD)
	k16 := FromUint16(u16)
	if len(k16) != 2 || k16[0] != 0xAB || k16[1] != 0xCD {
		t.Fatalf("FromUint16 wrong: %v", k16)
	}

	u64 := uint64(0x0102030405060708)
	k64 := FromUint64(u64)
	if len(k64) != 8 || k64[0] != 0x01 || k64[7] != 0x08 {
		t.Fatalf("FromUint64 wrong: %v", k64)
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
		if !(s1.LessThan(s2) || s2.LessThan(s1)) {
			t.Fatalf("expected one of %v or %v to be less", s1.Bytes(), s2.Bytes())
		}
	}
}
