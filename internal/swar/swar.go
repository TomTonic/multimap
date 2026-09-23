// Package swar holds the branch-free byte search primitives of the ART nodes:
// SWAR ("SIMD within a register") search over 8 key bytes at a time, and rank
// queries on a 256-bit presence bitmap. All of them inline into their callers.
package swar

import (
	"encoding/binary"
	"math/bits"
)

const (
	lsb = 0x0101010101010101
	msb = 0x8080808080808080
)

// Index8 returns the position of the lowest byte in w (little-endian, so byte 0
// of the source array is position 0) that equals b, or 8 if there is none.
//
// The classic "has zero byte" trick can flag false positives, but only in bytes
// more significant than a genuine match (a borrow propagates upwards), so the
// lowest flagged byte is always exact.
func Index8(w uint64, b byte) int {
	x := w ^ (lsb * uint64(b))
	m := (x - lsb) &^ x & msb
	return bits.TrailingZeros64(m) >> 3
}

// Word loads eight key bytes as one little-endian word. The compiler fuses this
// into a single load.
func Word(p []byte) uint64 { return binary.LittleEndian.Uint64(p) }

// Rank returns the number of set bits in bm below position b, which is the
// index of b's child in a node whose children are stored in byte order.
func Rank(bm *[4]uint64, b byte) int {
	w := b >> 6
	r := bits.OnesCount64(bm[w&3] & (uint64(1)<<(b&63) - 1))
	for i := byte(0); i < w; i++ {
		r += bits.OnesCount64(bm[i])
	}
	return r
}

// Has reports whether bit b is set.
func Has(bm *[4]uint64, b byte) bool {
	return bm[(b>>6)&3]&(uint64(1)<<(b&63)) != 0
}

// Set sets bit b.
func Set(bm *[4]uint64, b byte) { bm[(b>>6)&3] |= uint64(1) << (b & 63) }

// MatchPrefix reports whether key[depth:] can pass a node whose compressed path
// has length plen and whose first min(plen, 8) bytes are stored in prefix. Bytes
// beyond the eighth are not checked here (optimistic path compression); the
// full-key comparison at the leaf catches a mismatch there. plen must be > 0.
func MatchPrefix(prefix *[8]byte, plen int, key []byte, depth int) bool {
	rest := len(key) - depth
	if rest < plen {
		return false
	}
	m := min(plen, 8)
	if rest >= 8 {
		x := binary.LittleEndian.Uint64(key[depth:]) ^ binary.LittleEndian.Uint64(prefix[:])
		return x&(^uint64(0)>>(64-8*m)) == 0
	}
	for i := 0; i < m; i++ {
		if key[depth+i] != prefix[i] {
			return false
		}
	}
	return true
}

// Lcp returns the length of the longest common prefix of a and b.
func Lcp(a, b []byte) int {
	n := min(len(a), len(b))
	i := 0
	for i+8 <= n {
		x := binary.LittleEndian.Uint64(a[i:]) ^ binary.LittleEndian.Uint64(b[i:])
		if x != 0 {
			return i + bits.TrailingZeros64(x)>>3
		}
		i += 8
	}
	for i < n && a[i] == b[i] {
		i++
	}
	return i
}
