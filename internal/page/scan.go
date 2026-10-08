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

// in reports, for the key cp+rem of a first value, whether it is below the lower bound (skip) or above the upper
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

// valueRange returns the values [start, end) whose keys lie within b, for a page of the many-key form whose key lengths
// begin at kl and remainders at rem, and whether a key above the upper bound follows (over: the scan is done
// after this page). The keys are in order, so the values below the lower bound come first and those above the
// upper bound last: the lower bound is compared until the first key at or above it, the upper bound from there on.
func valueRange(m []byte, kl, n, rem int, b *Bounds) (start, end int, over bool) {
	cp := m[Header:kl]
	off := rem
	start = n
	for i, rl := range m[kl : kl+n] {
		if rl == Further {
			continue
		}
		r := m[off : off+int(rl)]
		off += int(rl)
		if start == n {
			if b.HasLo {
				if c := cmpKey(cp, r, b.Lo); c < 0 || (c == 0 && !b.LoIncl) {
					continue
				}
			}
			start = i
			if !b.HasHi {
				return start, n, false
			}
		}
		if c := cmpKey(cp, r, b.Hi); c > 0 || (c == 0 && !b.HiIncl) {
			return start, i, true
		}
	}
	return start, n, false
}

// ValuesIn returns the values of the keys of p that lie within b, in key order, as a slice of the page (it aliases
// p: valid until the map changes), and whether a key above the upper bound follows. A page of the one-key form
// holds one key, which the caller has checked: all its values. It is the tree's range scan of a page
// (docs/redesign/scan-design.md): the scan hands the slice to its caller's loop as it is.
func (p *Fixed) ValuesIn[T comparable](b *Bounds) ([]T, bool) {
	m := p.mem()
	n, l := int(p.currentValues), p.cpl()
	vs := valuesIn[T](m, n)
	if p.one() || !b.HasLo && !b.HasHi {
		return vs, false
	}
	start, end, over := valueRange(m, Header+l, n, Header+l+2*n, b)
	return vs[start:end], over
}

// AppendStrings appends the values of the keys of p that lie within b, in key order, to dst as strings (copies),
// and reports whether a key above the upper bound follows. A page of the one-key form gives all its values (see
// Fixed.ValuesIn).
func (p *Str) AppendStrings(dst []string, b *Bounds) ([]string, bool) {
	m := p.mem()
	n, l := int(p.currentValues), p.cpl()
	vl := Header + l // the value lengths of the one-key form follow the key part
	start, end, over := 0, n, false
	if !p.one() {
		vl += 2 * n
		if b.HasLo || b.HasHi {
			start, end, over = valueRange(m, Header+l, n, vl+n, b)
		}
	}
	s := len(m) - sum(m[vl+start:vl+n])
	for _, x := range m[vl+start : vl+end] {
		dst = append(dst, string(m[s:s+int(x)]))
		s += int(x)
	}
	return dst, over
}

// AppendKeys appends the keys within b of p, a page of the many-key form, each as pre, the key part and the
// remainder, to buf, and the end of each key in buf to ends; it reports whether a key above the upper bound
// follows. It is the range scan of the keys of a page.
func (p *head) AppendKeys(buf []byte, ends []int, pre []byte, b *Bounds, str bool) ([]byte, []int, bool) {
	m := p.mem()
	n, l := int(p.currentValues), p.cpl()
	kl := Header + l
	rem := kl + 2*n
	if str {
		rem += n
	}
	cp := m[Header:kl]
	for _, rl := range m[kl : kl+n] {
		if rl == Further {
			continue
		}
		r := m[rem : rem+int(rl)]
		rem += int(rl)
		skip, past := b.in(cp, r)
		if past {
			return buf, ends, true
		}
		if !skip {
			buf = append(append(append(buf, pre...), cp...), r...)
			ends = append(ends, len(buf))
		}
	}
	return buf, ends, false
}
