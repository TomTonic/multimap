package multimap

// Key helpers used only by the archived ART implementation.

// LongestCommonPrefix returns the length of the longest common prefix of a and b.
// If either a or b is nil, one of both is empty or they have no common prefix, 0 is returned.
// The returned length is in bytes. The value is always between 0 and min(len(a), len(b)).
func LongestCommonPrefix(a, b Key) uint {
	i := uint(0)
	la := uint(len(a))
	lb := uint(len(b))
	for i < la && i < lb && a[i] == b[i] {
		i++
	}
	return i
}

// LongestCommonPrefixAsKey returns the longest common prefix of a and b as a Key.
// If either a or b is nil, one of both is empty or they have no common prefix,
// an empty (zero-length) Key is returned (not nil). The returned Key is an independent
// copy and can thus be safely mutated by the caller.
func LongestCommonPrefixAsKey(a, b Key) Key {
	lcp := LongestCommonPrefix(a, b)
	k := make([]byte, lcp)
	copy(k, a[:lcp])
	return Key(k)
}

func (k *Key) append(key Key) {
	*k = append(*k, key...)
}
