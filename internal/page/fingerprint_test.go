package page

import (
	"math/rand/v2"
	"testing"
)

// TestFingerprint pins what a fingerprint is made of.
//
// A user looks keys up in a map; the page that holds them finds a key by comparing its fingerprint byte with the
// stored ones first, and compares the key itself only if one is equal. For that the fingerprint must depend on
// everything that tells two remainders apart: their bytes (the last 14 of them) and their length.
//
// Expected: the function agrees with a reference that builds the block of 16 bytes one byte at a time (length of
// the remainder in two bytes, then the last up to 14 bytes right-aligned) for 10,000 random remainders of 0 to 40
// bytes; remainders that differ only in a leading zero byte (only in their length) get different fingerprints for
// at least one of 256 byte values; remainders that differ only in a byte before the last 14 get the same one.
func TestFingerprint(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 2))
	for range 10000 {
		r := make([]byte, rng.IntN(41))
		for i := range r {
			r[i] = byte(rng.Uint32())
		}
		var block [16]byte
		block[0], block[1] = byte(len(r)>>8), byte(len(r))
		for i := 1; i <= 14 && i <= len(r); i++ {
			block[16-i] = r[len(r)-i]
		}
		var w1, w2 uint64
		for i := 7; i >= 0; i-- {
			w1, w2 = w1<<8|uint64(block[i]), w2<<8|uint64(block[8+i])
		}
		if got, want := Fingerprint(r), byte(wh64(w2, wh64(w1, 0))); got != want {
			t.Fatalf("Fingerprint(%v) = %d, reference %d", r, got, want)
		}
	}

	differs := false
	for c := range 256 {
		differs = differs || Fingerprint([]byte{byte(c)}) != Fingerprint([]byte{0, byte(c)})
	}
	if !differs {
		t.Error("the fingerprint does not depend on the length of the remainder")
	}
	long := []byte("0123456789abcdefghij")
	other := []byte("X123456789abcdefghij") // differs in the first byte, which lies before the last 14
	if Fingerprint(long) != Fingerprint(other) {
		t.Error("the fingerprint depends on a byte before the last 14")
	}
}
