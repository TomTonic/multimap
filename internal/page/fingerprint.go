package page

import (
	"encoding/binary"
	"math/bits"
)

// wh64 is hashing.WH64Det of github.com/TomTonic/Set3 (branch new_hasher): wyhash's mixing for one 64-bit word.
func wh64(val, seed uint64) uint64 {
	const m5, p1 = 0x1d8e4e27c47d124f, 0xf20a3e5e0b7b9731
	hi, lo := bits.Mul64(val^p1, bits.RotateLeft64(val, 32)^seed)
	hi, lo = bits.Mul64(m5^8, hi^lo)
	return hi ^ lo
}

// Fingerprint returns the fingerprint of a remainder r (the bytes of a key after the key part of its page): the
// lowest 8 bits of a wyhash mix of a block of 16 bytes, the length of r as two bytes (big-endian) and the last up
// to 14 bytes of r, right-aligned with zeros in front. A many-key page stores the fingerprint of every key's
// remainder in a list behind the key lengths, so that a search compares only the keys whose fingerprint equals the
// one searched for (docs/redesign/page-search-design.md). It is not stable across versions and never leaves
// memory; it is exported for the statistics of the tree's diagnostics (build tag mkstats), which must measure what
// the page stores.
func Fingerprint(r []byte) byte {
	var block [16]byte
	binary.BigEndian.PutUint16(block[:], uint16(len(r)))
	tail := r[max(0, len(r)-14):]
	copy(block[16-len(tail):], tail)
	w1, w2 := binary.LittleEndian.Uint64(block[:8]), binary.LittleEndian.Uint64(block[8:])
	return byte(wh64(w2, wh64(w1, 0)))
}
