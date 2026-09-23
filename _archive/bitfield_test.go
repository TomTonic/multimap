package multimap

import "testing"

func TestBitfieldGetSetClear(t *testing.T) {
	var b bitfield256

	indices := []byte{0, 63, 64, 127, 128, 191, 192, 255}
	// initially all bits should be clear
	for _, i := range indices {
		if b.get(i) {
			t.Fatalf("bit %d should be clear initially", i)
		}
	}

	// set and verify
	for _, i := range indices {
		b.set(i)
		if !b.get(i) {
			t.Fatalf("bit %d should be set after set()", i)
		}
	}

	// some other bits should remain clear
	for _, i := range []byte{1, 2, 60, 65, 129, 254} {
		if b.get(i) {
			t.Fatalf("bit %d should remain clear", i)
		}
	}

	// clear and verify
	for _, i := range indices {
		b.clear(i)
		if b.get(i) {
			t.Fatalf("bit %d should be clear after clear()", i)
		}
	}
}

func TestBitfieldTotalBitCount(t *testing.T) {
	var b bitfield256

	if got := b.totalBitCount(); got != 0 {
		t.Fatalf("expected count 0 on new bitfield, got %d", got)
	}

	// set a few bits (including a duplicate set) and check count
	b.set(10)
	b.set(20)
	b.set(10) // duplicate, should not increase count
	if got := b.totalBitCount(); got != 2 {
		t.Fatalf("expected count 2 after setting two distinct bits, got %d", got)
	}

	// add more bits and check
	b.set(0)
	b.set(255)
	if got := b.totalBitCount(); got != 4 {
		t.Fatalf("expected count 4 after adding two more bits, got %d", got)
	}

	// clear one and check
	b.clear(20)
	if got := b.totalBitCount(); got != 3 {
		t.Fatalf("expected count 3 after clearing one bit, got %d", got)
	}
}

func TestBitfieldMultipleOperations(t *testing.T) {
	var b bitfield256

	// stress same index many times
	for i := 0; i < 10; i++ {
		b.set(42)
	}
	if !b.get(42) {
		t.Fatalf("bit 42 should be set")
	}
	if got := b.totalBitCount(); got != 1 {
		t.Fatalf("expected count 1 after repeatedly setting same bit, got %d", got)
	}

	// clear and ensure count goes to zero
	b.clear(42)
	if b.get(42) {
		t.Fatalf("bit 42 should be clear after clear()")
	}
	if got := b.totalBitCount(); got != 0 {
		t.Fatalf("expected count 0 after clearing last bit, got %d", got)
	}
}
