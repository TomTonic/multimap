package page

import "bytes"

// valueStart returns where the value of slot pos begins in the object m of a Str page: the values
// end with the object, so it is the end less the lengths of the values from pos on.
func (la *lay) valueStart(m []byte, pos int) int {
	return len(m) - sum(m[la.vl+pos:la.vl+la.n])
}

// BuildStrings returns a page for the values vals of the keys rests (the keys from the end of the
// path on, in key order: a key with several values appears once for each, one after the other), or
// nil if they do not fit: no value, more than MaxEntries, a remainder beyond MaxRemainder (many keys), a
// value beyond MaxValue or content beyond the largest class. If all keys are equal the page has the
// one-key form; else the common prefix of the page is the longest one all keys share (up to MaxKeyPart).
// The tree calls it when pages meet and when it builds the subtree of a burst.
func BuildStrings(rests, vals [][]byte) *Str {
	n := len(rests)
	if n == 0 || n > MaxEntries || len(vals) != n {
		return nil
	}
	one := string(rests[0]) == string(rests[n-1])
	l := min(lcp(rests[0], rests[n-1]), MaxKeyPart)
	if one {
		l = len(rests[0])
	}
	need := Header + l + n
	vb := 0
	if !one {
		need += n
	}
	for i, r := range rests {
		if !one && !further(rests, i) {
			if len(r)-l > MaxRemainder {
				return nil
			}
			need += len(r) - l
		}
		if len(vals[i]) > MaxValue {
			return nil
		}
		vb += len(vals[i])
	}
	c := classFor(need + vb)
	if c < 0 {
		return nil
	}
	p := (*Str)(allocRaw(c))
	p.setHead(!one, c, n, l, 0)
	m := p.mem()
	copy(m[Header:], rests[0][:l])
	la := p.lay(true)
	off, vo := la.rem, len(m)-vb
	for i, r := range rests {
		switch {
		case one:
		case further(rests, i):
			m[la.kl+i] = Further
		default:
			m[la.kl+i] = uint8(len(r) - l)
			off += copy(m[off:], r[l:])
		}
		m[la.vl+i] = uint8(len(vals[i]))
		vo += copy(m[vo:], vals[i])
	}
	return p
}

// further reports whether rests[i] is a further value of the key of rests[i-1].
func further(rests [][]byte, i int) bool {
	return i > 0 && string(rests[i]) == string(rests[i-1])
}

// slotOf finds the first slot of the key rest in a Str page, or reports that it is not there.
func (p *Str) slotOf(m []byte, la lay, rest []byte) (int, bool) {
	if !la.many {
		return 0, bytes.Equal(m[Header:Header+la.l], rest)
	}
	if p.Match(rest) < la.l {
		return 0, false
	}
	pos, _, found := la.locate(m, rest[la.l:])
	return pos, found
}

// Used returns the bytes of the page that hold keys: where the remainders end. The values are at the
// end of the object.
func (p *Str) Used() int {
	la := p.lay(true)
	return la.keyEnd(p.mem())
}

// Get returns the first value of the key rest (the key from the end of the path on), or false if the
// page does not hold it. The slice aliases the page.
func (p *Str) Get(rest []byte) ([]byte, bool) {
	m := p.mem()
	n, l := int(p.n), p.cpl()
	pos, vl := 0, Header+l // the value lengths of the one-key form follow the key part
	if p.one() {
		if string(m[Header:Header+l]) != string(rest) {
			return nil, false
		}
	} else {
		if lcp(m[Header:Header+l], rest) < l {
			return nil, false
		}
		var found bool
		if pos, _, found = locate(m, Header+l, n, Header+l+n+n, rest[l:]); !found {
			return nil, false
		}
		vl += n
	}
	s := len(m) - sum(m[vl+pos:vl+n])
	return m[s : s+int(m[vl+pos])], true
}

// EachValue calls fn with every value of the key rest, in the order they came in, until fn returns
// false, and reports whether the page holds the key.
func (p *Str) EachValue(rest []byte, fn func(val []byte) bool) bool {
	m := p.mem()
	la := p.lay(true)
	pos, ok := p.slotOf(m, la, rest)
	if !ok {
		return false
	}
	end := la.n
	if la.many {
		end = la.runEnd(m, pos)
	}
	s := la.valueStart(m, pos)
	for i := pos; i < end; i++ {
		l := int(m[la.vl+i])
		if !fn(m[s : s+l]) {
			return true
		}
		s += l
	}
	return true
}

// hasValue returns the slot among pos..end-1 whose value is val, or -1.
func (la *lay) hasValue(m []byte, pos, end int, val []byte) int {
	s := la.valueStart(m, pos)
	for i := pos; i < end; i++ {
		l := int(m[la.vl+i])
		if l == len(val) && string(m[s:s+l]) == string(val) {
			return i
		}
		s += l
	}
	return -1
}

// Add adds the value val to the key rest and returns the page that holds the result, which is p
// itself unless the content no longer fits p's class or the page of one key meets another key, and
// says what happened (see Result). val is copied. A key that is there gets the value behind its others.
func (p *Str) Add(rest, val []byte) (*Str, Result) {
	m := p.mem()
	la := p.lay(true)
	var slot, ro int
	var kl byte
	var r []byte
	res := AddedValue
	e := la.rem
	switch {
	case !la.many:
		if !bytes.Equal(m[Header:Header+la.l], rest) {
			return p.pairWith(rest, val)
		}
		if la.hasValue(m, 0, la.n, val) >= 0 {
			return p, Present
		}
		if len(val) > MaxValue {
			return p, Full
		}
		slot = la.n
	default:
		if p.Match(rest) < la.l {
			return p, Outside
		}
		r = rest[la.l:]
		pos, off, found := la.locate(m, r)
		ro, slot = off-la.rem, pos
		if found {
			end := la.runEnd(m, pos)
			if la.hasValue(m, pos, end, val) >= 0 {
				return p, Present
			}
			if len(val) > MaxValue {
				return p, Full
			}
			slot, kl, r = end, Further, nil
			e = la.keyEndFrom(m, pos, off)
		} else {
			if len(r) > MaxRemainder || len(val) > MaxValue {
				return p, Full
			}
			kl, res = uint8(len(r)), Added
			e = la.keyEndFrom(m, pos, off)
		}
	}
	vb := sum(m[la.vl : la.vl+la.n])
	need := e + b2i(la.many) + 1 + len(r) + vb + len(val)
	q := p
	if need > p.Size() {
		c := classFor(need)
		if c < 0 {
			return p, Full
		}
		q = p.regrow(c, e, vb)
		m = q.mem()
	}
	// the values area: the values of the slots before the new one move down by its length
	tail := sum(m[la.vl+slot : la.vl+la.n])
	front := len(m) - vb
	copy(m[front-len(val):len(m)-tail-len(val)], m[front:len(m)-tail])
	copy(m[len(m)-tail-len(val):], val)
	insertSlot(m, &la, e, slot, ro, kl, uint8(len(val)), r)
	q.n++
	return q, res
}

// regrow returns a new page of class c with the key area (the first e bytes) and the values (vb
// bytes) of p.
func (p *Str) regrow(c, e, vb int) *Str {
	q := (*Str)(allocRaw(c))
	q.setHead(!p.one(), c, int(p.n), p.cpl(), 0)
	pm, qm := p.mem(), q.mem()
	copy(qm[Header:e], pm[Header:e])
	copy(qm[len(qm)-vb:], pm[len(pm)-vb:])
	return q
}

// pairWith makes the page of the two keys that a one-key page and a new key rest make, in the
// many-key form: both keys with the values of the one-key page and val. It returns the page, or p and
// Full if they do not fit.
func (p *Str) pairWith(rest, val []byte) (*Str, Result) {
	m := p.mem()
	la := p.lay(true)
	key := m[Header : Header+la.l]
	var rb, vb [9][]byte
	rests, vals := rb[:0], vb[:0]
	first := bytes.Compare(rest, key) < 0
	if first {
		rests, vals = append(rests, rest), append(vals, val)
	}
	s := la.valueStart(m, 0)
	for i := range la.n {
		l := int(m[la.vl+i])
		rests, vals = append(rests, key), append(vals, m[s:s+l])
		s += l
	}
	if !first {
		rests, vals = append(rests, rest), append(vals, val)
	}
	if q := BuildStrings(rests, vals); q != nil {
		return q, Added
	}
	return p, Full
}

// Remove removes the value val of the key rest and says whether it was there (Removed) and
// whether the key went with it (Gone, its last value); a key that has no value left is gone. It
// returns the page that holds the rest: p itself, a page of a smaller class once the content fills at
// most ShrinkLimit of it, or nil if the value was the only one (the page is gone). A many-key page that
// is left with one key becomes a one-key page in place, if the key fits the key part.
func (p *Str) Remove(rest, val []byte) (*Str, Removal) {
	m := p.mem()
	la := p.lay(true)
	var slot, ro, remLen int
	res := Removed
	e := la.rem
	if !la.many {
		if !bytes.Equal(m[Header:Header+la.l], rest) {
			return p, Absent
		}
		if slot = la.hasValue(m, 0, la.n, val); slot < 0 {
			return p, Absent
		}
		if la.n == 1 {
			return nil, Gone
		}
	} else {
		if p.Match(rest) < la.l {
			return p, Absent
		}
		r := rest[la.l:]
		pos, off, found := la.locate(m, r)
		if !found {
			return p, Absent
		}
		end := la.runEnd(m, pos)
		if slot = la.hasValue(m, pos, end, val); slot < 0 {
			return p, Absent
		}
		e = la.keyEndFrom(m, pos, off)
		if slot == pos {
			if end > pos+1 { // the next value takes the key's place: its slot becomes the key's
				m[la.kl+slot+1] = uint8(len(r))
			} else { // (a many-key page has two keys at least)
				remLen, res = len(r), Gone
				ro = off - la.rem
			}
		}
	}
	vb := sum(m[la.vl : la.vl+la.n])
	vlen := int(m[la.vl+slot])
	start := la.valueStart(m, slot)
	front := len(m) - vb
	copy(m[front+vlen:start+vlen], m[front:start])
	clear(m[front : front+vlen])
	removeSlot(m, &la, e, slot, ro, remLen)
	p.n--
	e -= b2i(la.many) + 1 + remLen
	vb -= vlen
	if la.many && res == Gone && oneKeyLeft(m, la.kl, int(p.n)) {
		e = toOneKey(&p.head, m, p.lay(true), e)
	}
	if c := shrinkClass(p.class(), e+vb); c < p.class() {
		return p.regrow(c, e, vb), res
	}
	return p, res
}

// keys returns the number of keys.
func (p *head) keys() int {
	if p.one() {
		return 1
	}
	return p.KeysUpTo(MaxEntries)
}

// Keys returns the number of keys of the page.
func (p *head) Keys() int { return p.keys() }

// KeysUpTo returns the number of keys of the page, or limit if it has at least that many: the scan of
// the key lengths stops there.
func (p *head) KeysUpTo(limit int) int {
	if p.one() {
		return min(1, limit)
	}
	m := p.mem()
	lo := Header + p.cpl()
	k := 0
	for _, rl := range m[lo : lo+int(p.n)] {
		if k += b2i(rl != Further); k >= limit {
			break
		}
	}
	return k
}

// Each calls fn with the remainder (after the key part), the value and whether it is the first value of
// its key, for every value in key order until fn returns false, and reports whether it ran to
// completion. The slices alias the page; the remainder of a further value is that of its key, and
// that of the one-key form is empty (the key is the key part). The key of an entry is the path, then CP,
// then the remainder.
func (p *Str) Each(fn func(rem, val []byte, first bool) bool) bool {
	m := p.mem()
	la := p.lay(true)
	s := la.valueStart(m, 0)
	if !la.many {
		for i := range la.n {
			l := int(m[la.vl+i])
			if !fn(nil, m[s:s+l], i == 0) {
				return false
			}
			s += l
		}
		return true
	}
	off := la.rem
	var rem []byte
	for i, rl := range m[la.kl : la.kl+la.n] {
		first := rl != Further
		if first {
			rem = m[off : off+int(rl)]
			off += int(rl)
		}
		l := int(m[la.vl+i])
		if !fn(rem, m[s:s+l], first) {
			return false
		}
		s += l
	}
	return true
}

// Skip drops the first k bytes of the key part (k at most its length): the tree has put a byte node
// above the page that consumes them. It happens in place.
func (p *Str) Skip(k int) {
	m := p.mem()
	e := p.Used()
	copy(m[Header:], m[Header+k:e])
	clear(m[e-k : e])
	p.setCpl(p.cpl() - k)
}

// Prepend returns the page with pre in front of the key part: the page moves up in the tree, when
// the node above it goes away. It is p itself if the content still fits, else a page of a larger class;
// nil if the content would exceed the largest class.
func (p *Str) Prepend(pre []byte) *Str {
	m := p.mem()
	la := p.lay(true)
	e := la.keyEnd(m)
	vb := sum(m[la.vl : la.vl+la.n])
	nk := e + len(pre)
	c := classFor(nk + vb)
	if c < 0 {
		return nil
	}
	q := p
	if c > p.class() {
		q = p.regrow(c, e, vb)
		m = q.mem()
	}
	copy(m[Header+len(pre):nk], m[Header:e])
	copy(m[Header:], pre)
	q.setCpl(la.l + len(pre))
	return q
}

// Widen adds the entry rest -> val, a key that leaves the common prefix of the many-key page after
// mis = Match(rest) < PrefixLen() bytes, and shortens the common prefix to those mis bytes: the page is
// built again, with the rest of the old prefix in front of every remainder. It returns the new page, or
// nil if the entry does not go in (the content would exceed the largest class, or a remainder or the
// value is too long); p is then unchanged. The tree calls it where it would otherwise put a byte node
// above the page.
func (p *Str) Widen(rest, val []byte) *Str {
	m := p.mem()
	la := p.lay(true)
	mis := p.Match(rest)
	d := la.l - mis
	newRem := rest[mis:]
	if len(newRem) > MaxRemainder || len(val) > MaxValue {
		return nil
	}
	e := la.keyEnd(m)
	vb := sum(m[la.vl : la.vl+la.n])
	heads, longest := 0, 0
	for _, rl := range m[la.kl : la.kl+la.n] {
		if rl != Further {
			heads++
			longest = max(longest, int(rl))
		}
	}
	if longest+d > MaxRemainder {
		return nil
	}
	nk := Header + mis + 2*(la.n+1) + (e - la.rem) + heads*d + len(newRem)
	c := classFor(nk + vb + len(val))
	if c < 0 {
		return nil
	}
	first := len(newRem) == 0 || newRem[0] < p.CP()[mis]
	at := 0
	if !first {
		at = la.n
	}
	q := (*Str)(allocRaw(c))
	q.setHead(true, c, la.n+1, mis, 0)
	out := q.mem()
	copy(out[Header:], m[Header:Header+mis])
	nl := q.lay(true)
	for i := range la.n {
		rl := m[la.kl+i]
		if rl != Further {
			rl += uint8(d)
		}
		out[nl.kl+i+b2i(i >= at)] = rl
		out[nl.vl+i+b2i(i >= at)] = m[la.vl+i]
	}
	out[nl.kl+at] = uint8(len(newRem))
	out[nl.vl+at] = uint8(len(val))
	off := nl.rem
	if first {
		off += copy(out[off:], newRem)
	}
	extra := p.CP()[mis:]
	src := la.rem
	for i := range la.n {
		if rl := int(m[la.kl+i]); rl != Further {
			off += copy(out[off:], extra)
			off += copy(out[off:], m[src:src+rl])
			src += rl
		}
	}
	if !first {
		copy(out[off:], newRem)
	}
	nvb := vb + len(val)
	vo := len(out) - nvb
	if first {
		vo += copy(out[vo:], val)
		copy(out[vo:], m[len(m)-vb:])
	} else {
		vo += copy(out[vo:], m[len(m)-vb:])
		copy(out[vo:], val)
	}
	return q
}
