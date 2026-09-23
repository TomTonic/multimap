package art

import (
	mm "github.com/TomTonic/multimap"
)

type keyComparison int32

const (
	cmpUnclear          keyComparison = 0b0000_0111 // this one is private on purpose
	CmpLessThan         keyComparison = 0b0000_0001
	CmpLessThanEqual    keyComparison = 0b0000_0011
	CmpEqual            keyComparison = 0b0000_0010
	CmpGreaterThanEqual keyComparison = 0b0000_0110
	CmpGreaterThan      keyComparison = 0b0000_0100
)

// This function compares the next Position in two Keys a and b.
// It assumes that all previous positions were equal.
// It returns a comparisonResult indicating the relation between a and b
// at the given position.
// If either a or b has no byte at nextByteToCompare while the other has
// (i.e. one Key is a prefix of the other), the shorter Key is considered
// less than the longer Key (CmpLessThan resp. CmpGreaterThan).
// If both Keys have no byte at nextByteToCompare (i.e. both Keys are
// equal up to this position and have the same length), the Keys is
// considered equal (CmpEqual).
// If both Keys have a byte at nextByteToCompare and these bytes are unequal,
// the result indicates which Key is less/greater (CmpLessThan resp.
// CmpGreaterThan).
// If both Keys have a byte at nextByteToCompare and these bytes are equal,
// the result is cmpUnclear (further comparison is needed).
// nextByteToCompare is zero-based.
// Note: This function does not return CmpLessThanEqual or CmpGreaterThanEqual.
func compareNextPos(a, b mm.Key, nextByteToCompare int) keyComparison {
	idx := int(nextByteToCompare)
	la, lb := len(a), len(b)

	hasA := idx < la
	hasB := idx < lb

	switch {
	case !hasA && !hasB:
		// both exhausted at this position -> equal
		return CmpEqual
	case !hasA && hasB:
		// a is shorter (prefix) -> a less than b
		return CmpLessThan
	case hasA && !hasB:
		// b is shorter (prefix) -> a greater than b
		return CmpGreaterThan
	default:
		// both have a byte at idx -> compare bytes
		ba := a[idx]
		bb := b[idx]
		if ba < bb {
			return CmpLessThan
		}
		if ba > bb {
			return CmpGreaterThan
		}
		// bytes equal -> further comparison needed
		return cmpUnclear
	}
}
