package multimap

import "math/bits"

// bitfield256 is a bitfield to store the presence of keys in a node.
// It is a bitfield of 256 bits, so it is stored in 4 uint64.
type bitfield256 [4]uint64

func (p *bitfield256) get(index byte) bool {
	return ((*p)[index>>6] & (1 << (index & 0x3F))) != 0
}

func (p *bitfield256) set(index byte) {
	(*p)[index>>6] |= (1 << (index & 0x3F))
}

func (p *bitfield256) clear(index byte) {
	(*p)[index>>6] &= ^(1 << (index & 0x3F))
}

func (p *bitfield256) totalBitCount() byte {
	var count byte
	for i := range 4 {
		count += byte(bits.OnesCount64((*p)[i]))
	}
	return count
}
