package page

import "bytes"

// Bounds are the bounds of a scan as a page of the many-key form sees them: the keys of the bounds from the
// page's path length on (the bytes before are the path to the page, which agrees with a bound wherever the
// bound is set). A side that is not set does not limit the page: the page lies wholly beyond that bound.
type Bounds struct {
	Lo, Hi         []byte
	HasLo, HasHi   bool
	LoIncl, HiIncl bool
}

// in reports, for the key cp+rem of a first slot, whether it is below the lower bound (skip) or above the upper
// bound (over: the scan is done).
func (b *Bounds) in(cp, rem []byte) (skip, over bool) {
	if b.HasLo {
		if c := cmpKey(cp, rem, b.Lo); c < 0 || (c == 0 && !b.LoIncl) {
			skip = true
		}
	}
	if b.HasHi {
		if c := cmpKey(cp, rem, b.Hi); c > 0 || (c == 0 && !b.HiIncl) {
			return skip, true
		}
	}
	return skip, false
}

// cmpKey compares the key a+b (the key part, then a remainder) with bound, without joining them.
func cmpKey(a, b, bound []byte) int {
	n := min(len(a), len(bound))
	if c := bytes.Compare(a[:n], bound[:n]); c != 0 {
		return c
	}
	if len(a) > len(bound) {
		return 1
	}
	return bytes.Compare(b, bound[len(a):])
}

// ScanValues calls yield with the values of the keys of p, a page of the many-key form, that lie within b, in
// key order, and reports false once the scan is over: yield returned false, or a key above the upper bound was
// reached. It is the tree's range scan of a page: one call a value, and a plain loop over the values when the page
// lies wholly inside the bounds (docs/redesign/scan-design.md).
func (p *Fixed) ScanValues[T comparable](b *Bounds, yield func(T) bool) bool {
	m := p.mem()
	n, l := int(p.n), p.cpl()
	vs := valuesIn[T](m, n)
	if !b.HasLo && !b.HasHi {
		for _, v := range vs {
			if !yield(v) {
				return false
			}
		}
		return true
	}
	kl := Header + l
	cp := m[Header:kl]
	off := kl + n
	skip := false
	for i, rl := range m[kl : kl+n] {
		if rl != Further {
			var over bool
			if skip, over = b.in(cp, m[off:off+int(rl)]); over {
				return false
			}
			off += int(rl)
		}
		if !skip && !yield(vs[i]) {
			return false
		}
	}
	return true
}

// ScanValues calls yield with the values of the keys of p, a page of the many-key form, that lie within b, as
// strings, in key order, and reports false once the scan is over (see Fixed.ScanValues). The strings are copies.
func (p *Str) ScanValues(b *Bounds, yield func(string) bool) bool {
	m := p.mem()
	n, l := int(p.n), p.cpl()
	kl := Header + l
	vl := kl + n
	s := len(m) - sum(m[vl:vl+n])
	if !b.HasLo && !b.HasHi {
		for _, x := range m[vl : vl+n] {
			if !yield(string(m[s : s+int(x)])) {
				return false
			}
			s += int(x)
		}
		return true
	}
	cp := m[Header:kl]
	off := vl + n
	skip := false
	for i, rl := range m[kl : kl+n] {
		if rl != Further {
			var over bool
			if skip, over = b.in(cp, m[off:off+int(rl)]); over {
				return false
			}
			off += int(rl)
		}
		x := int(m[vl+i])
		if !skip && !yield(string(m[s:s+x])) {
			return false
		}
		s += x
	}
	return true
}
