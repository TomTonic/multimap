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
	for i := range w {
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

// PrefixLen is the number of path bytes an inner node keeps in its header.
const PrefixLen = 12

// MatchPrefix reports whether key[depth:] can pass a node whose compressed path
// has length plen and whose first min(plen, PrefixLen) bytes are stored in
// prefix. Bytes beyond those are not checked here; the caller compares them
// with the rest of the path, which its node keeps elsewhere.
// plen must be > 0.
func MatchPrefix(prefix *[PrefixLen]byte, plen int, key []byte, depth int) bool {
	rest := len(key) - depth
	if rest < plen {
		return false
	}
	m := min(plen, PrefixLen)
	if m <= 8 && rest >= 8 {
		return Match8(prefix, m, key, depth)
	}
	return matchLong(prefix, m, key, depth)
}

// Match8 is the common case of MatchPrefix, small enough to inline into the
// descent: it reports whether a path of plen <= 8 bytes matches a key that has
// at least 8 bytes from depth on. It returns false whenever either condition
// does not hold, so a caller falls back to MatchPrefix then.
func Match8(prefix *[PrefixLen]byte, plen int, key []byte, depth int) bool {
	if plen > 8 || len(key)-depth < 8 {
		return false
	}
	x := binary.LittleEndian.Uint64(key[depth:]) ^ binary.LittleEndian.Uint64(prefix[:8])
	return x&(^uint64(0)>>(64-8*plen)) == 0
}

// matchLong compares the first m bytes of prefix with key[depth:], which has
// at least m bytes, where Match8 does not apply.
func matchLong(prefix *[PrefixLen]byte, m int, key []byte, depth int) bool {
	if len(key)-depth >= PrefixLen {
		k := key[depth:]
		x := binary.LittleEndian.Uint64(k) ^ binary.LittleEndian.Uint64(prefix[:8])
		y := binary.LittleEndian.Uint32(k[8:]) ^ binary.LittleEndian.Uint32(prefix[8:])
		return x == 0 && y&(^uint32(0)>>(32-8*(m-8))) == 0
	}
	for i := range m {
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
