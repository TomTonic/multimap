package multimap

import (
	"encoding/binary"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// Key is an alias for a byte slice used as a map key representation.
// Use the provided constructors to build Keys from primitive types or
// normalized strings.
//
// Integer encoding policy
// -----------------------
// All integer constructors produce an 8-byte big-endian representation
// (most-significant byte first). To ensure consistent, order-preserving
// comparisons across signed and unsigned types and across different
// integer widths, every integer constructor adds an offset of `1<<63`
// before encoding the numeric value. For signed constructors the value
// is first converted to `int64`, for unsigned constructors it is treated
// as `uint64`; in both cases the offset is added and the resulting
// unsigned 64-bit value is written big-endian into the Key.
//
// This mapping has two useful properties:
//   - Lexicographic byte-wise comparison of Keys corresponds to numeric
//     ordering of the original values (taking signedness into account).
//   - Values produced from different source widths are comparable (for
//     example `FromInt32(x)` equals `FromInt64(x)` for the same numeric x).
//
// Note: Because both signed and unsigned constructors add the same
// offset, `FromInt64(0)` equals `FromUint64(0)`. The smallest `int64`
// value (`math.MinInt64`) maps to `0` and negative signed values compare
// before zero/positive values as expected for numeric ordering.
type Key []byte

// FromBytes returns a copy of the provided byte slice as a Key. If b is
// nil this returns an empty (zero-length) Key (not nil).
func FromBytes(b []byte) Key {
	if b == nil {
		return []byte{}
	}
	kb := make([]byte, len(b))
	copy(kb, b)
	return Key(kb)
}

// FromString returns a Key produced from the provided string after
// normalizing it to Unicode NFC. The resulting Key contains the UTF-8
// encoding of the normalized string. (FromString does not alter case or
// trim spaces.)
func FromString(s string) Key {
	s = norm.NFC.String(s) // normalize to NFC
	return FromBytes([]byte(s))
}

// FromInt converts an `int` to an 8-byte big-endian Key. The signed integer
// range is shifted by adding 1<<63 so that negative values compare before
// positive values when Keys are compared lexicographically.
func FromInt(i int) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	u := uint64(int64(i)) + offset
	binary.BigEndian.PutUint64(b[:], u)
	return FromBytes(b[:])
}

// FromInt64 converts an int64 to an 8-byte big-endian Key. The value is
// shifted by 1<<63 so that lexicographic key order matches numeric order.
func FromInt64(i int64) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	u := uint64(i) + offset
	binary.BigEndian.PutUint64(b[:], u)
	return FromBytes(b[:])
}

// FromInt32 converts an int32 to an 8-byte big-endian Key (value is encoded
// into 64 bits). The signed value is shifted by 1<<63 for order-preserving
// behavior across widths.
func FromInt32(i int32) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	u := uint64(int64(i)) + offset
	binary.BigEndian.PutUint64(b[:], u)
	return FromBytes(b[:])
}

// FromInt16 converts an int16 to an 8-byte big-endian Key (value is encoded
// into 64 bits). The signed value is shifted by 1<<63 for order-preserving
// behavior across widths.
func FromInt16(i int16) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	u := uint64(int64(i)) + offset
	binary.BigEndian.PutUint64(b[:], u)
	return FromBytes(b[:])
}

// FromInt8 converts an int8 to an 8-byte big-endian Key (value is encoded
// into 64 bits). The signed value is shifted by 1<<63 for order-preserving
// behavior across widths.
func FromInt8(i int8) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	u := uint64(int64(i)) + offset
	binary.BigEndian.PutUint64(b[:], u)
	return FromBytes(b[:])
}

// FromUint converts a uint to an 8-byte big-endian Key (MSB first).
func FromUint(u uint) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	binary.BigEndian.PutUint64(b[:], uint64(u)+offset)
	return FromBytes(b[:])
}

// FromUint64 converts a uint64 to an 8-byte big-endian Key (MSB first).
func FromUint64(u uint64) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	binary.BigEndian.PutUint64(b[:], u+offset)
	return FromBytes(b[:])
}

// FromUint32 converts a uint32 to an 8-byte big-endian Key (value encoded into 64 bits).
func FromUint32(u uint32) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	binary.BigEndian.PutUint64(b[:], uint64(u)+offset)
	return FromBytes(b[:])
}

// FromUint16 converts a uint16 to an 8-byte big-endian Key (value encoded into 64 bits).
func FromUint16(u uint16) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	binary.BigEndian.PutUint64(b[:], uint64(u)+offset)
	return FromBytes(b[:])
}

// FromUint8 converts a uint8 to an 8-byte big-endian Key (value encoded into 64 bits).
func FromUint8(u uint8) Key {
	var b [8]byte
	const offset = uint64(1) << 63
	binary.BigEndian.PutUint64(b[:], uint64(u)+offset)
	return FromBytes(b[:])
}

// FromByte is an alias for FromUint8 and produces an 8-byte representation.
func FromByte(b byte) Key { return FromUint8(uint8(b)) }

// FromRune converts a rune to its UTF-8 encoding as a Key.
func FromRune(r rune) Key {
	// encode rune to UTF-8 bytes
	var buf [4]byte
	n := utf8EncodeRune(buf[:], r)
	return FromBytes(buf[:n])
}

// Bytes returns a copy of the Key as a byte slice.
func (k Key) Bytes() []byte {
	if k == nil {
		return nil
	}
	b := make([]byte, len(k))
	copy(b, k)
	return b
}

// Clone returns an independent copy of the Key. If k is nil, Clone returns nil.
func (k Key) Clone() Key {
	if k == nil {
		return nil
	}
	kb := make([]byte, len(k))
	copy(kb, k)
	return Key(kb)
}

// String returns the Key as a string consisting of uppercase hex tuples per byte,
// separated by commas and surrounded by `[]` (e.g. `[01,AB,00]`).
func (k Key) String() string {
	if len(k) == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteByte('[')
	const hex = "0123456789ABCDEF"
	for i, b := range k {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteByte(hex[b>>4])
		sb.WriteByte(hex[b&0x0F])
	}
	sb.WriteByte(']')
	return sb.String()
}

// Equal reports whether k and other have the same contents.
func (k Key) Equal(other Key) bool {
	if len(k) != len(other) {
		return false
	}
	for i := range k {
		if k[i] != other[i] {
			return false
		}
	}
	return true
}

// LessThan reports whether k is lexicographically less than other.
func (k Key) LessThan(other Key) bool {
	for i := 0; i < len(k) && i < len(other); i++ {
		if k[i] < other[i] {
			return true
		} else if k[i] > other[i] {
			return false
		}
	}
	return len(k) < len(other)
}

// IsEmpty returns whether the Key is empty or nil.
func (k Key) IsEmpty() bool { return len(k) == 0 }

// helper: encode rune to utf8 into buf, return length
func utf8EncodeRune(buf []byte, r rune) int {
	switch {
	case r <= 0x7F:
		buf[0] = byte(r)
		return 1
	case r <= 0x7FF:
		buf[0] = 0xC0 | byte(r>>6)
		buf[1] = 0x80 | byte(r)&0x3F
		return 2
	case r <= 0xFFFF:
		buf[0] = 0xE0 | byte(r>>12)
		buf[1] = 0x80 | byte(r>>6)&0x3F
		buf[2] = 0x80 | byte(r)&0x3F
		return 3
	default:
		buf[0] = 0xF0 | byte(r>>18)
		buf[1] = 0x80 | byte(r>>12)&0x3F
		buf[2] = 0x80 | byte(r>>6)&0x3F
		buf[3] = 0x80 | byte(r)&0x3F
		return 4
	}
}
