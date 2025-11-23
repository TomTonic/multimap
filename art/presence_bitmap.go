package art

// PresenceBitmap is a compact 256-bit presence map used by some node types.
// It is stored as four 64-bit words (little index: word 0 contains bits 0..63).
type PresenceBitmap [4]uint64

// Get reports whether the bit for index b (0..255) is set.
func (p *PresenceBitmap) Get(b byte) bool {
	word := b >> 6
	off := b & 0x3F
	return ((*p)[word] & (uint64(1) << off)) != 0
}

// Set marks the bit for index b (0..255).
func (p *PresenceBitmap) Set(b byte) {
	word := b >> 6
	off := b & 0x3F
	(*p)[word] |= uint64(1) << off
}

// Clear clears the bit for index b (0..255).
func (p *PresenceBitmap) Clear(b byte) {
	word := b >> 6
	off := b & 0x3F
	(*p)[word] &^= uint64(1) << off
}
