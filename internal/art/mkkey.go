package art

import (
	"bytes"
	"slices"
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/page"
	"github.com/TomTonic/multimap/internal/swar"
)

// In a map of strings or of pointer-free fixed-size values (Map.mk), entries share a
// multi-key page (internal/page, docs/redesign/step5-mkmv-design.md): the entries of a
// subtree whose content fits 512 bytes, in key order, with the bytes their keys share
// stored once, each entry with one value or several. A page hangs below a byte node like a
// single-key page and holds its keys from the path length of its place on; when a node is
// put above it or goes away, it loses or gains the front bytes of its common prefix (Skip,
// Prepend), so that its keys always begin at its path length.
//
// The tree builds a page where two entries meet (pair), adds values to it (reach), and
// builds the subtree again when a page is full (burst): build, from the entries in key
// order, makes a page of the entries that fit one, else a byte node on the next byte in
// which they differ and a subtree for each byte. An entry alone below a byte is a single-key
// page; an entry with more values than a page of 512 bytes holds is a value overflow.

func asMKStr(h *header) *page.Str      { return (*page.Str)(unsafe.Pointer(h)) }
func asMKFix(h *header) *page.Fixed    { return (*page.Fixed)(unsafe.Pointer(h)) }
func mkStrHdr(p *page.Str) *header     { return (*header)(unsafe.Pointer(p)) }
func mkFixHdr(p *page.Fixed) *header   { return (*header)(unsafe.Pointer(p)) }
func strOf[T comparable](v T) string   { return *(*string)(unsafe.Pointer(&v)) }
func fromStr[T comparable](s string) T { return *(*T)(unsafe.Pointer(&s)) }
func asStrSlice[T comparable](v []T) [][]byte {
	out := make([][]byte, len(v))
	for i := range v {
		out[i] = view(strOf(v[i]))
	}
	return out
}

// item is an entry of a subtree that build turns into pages and nodes: its key from
// the path length of the subtree on, its value, or all its values if it has
// several.
type item[T comparable] struct {
	rest  []byte
	val   T
	multi []T // the values if there are several, else nil
}

// values returns the values of the entry.
func (it *item[T]) values() []T {
	if it.multi != nil {
		return it.multi
	}
	return []T{it.val}
}

// leafFor returns the single-key page, or the value overflow, of the entry it, whose key part is its rest.
func (m *Map[T]) leafFor(it *item[T]) *singleKeyHead {
	if it.multi == nil {
		switch m.flat {
		case 3:
			if p := page.NewStr(it.rest, view(strOf(it.val))); p != nil {
				return skHead(p)
			}
		case 1:
			if p := page.NewFixed(it.rest, it.val, m.ptr); p != nil {
				return fixedHead(p)
			}
		}
	} else if p := m.onePage(it); p != nil {
		return p
	}
	l := newValueOverflow(it.rest, set3.EmptyWithCapacity[T](uint32(2*len(it.values()))))
	for _, v := range it.values() {
		overflowAdd(l, v)
	}
	return l
}

// onePage returns the page of the entry it, which has several values, or nil if they do not fit one.
func (m *Map[T]) onePage(it *item[T]) *singleKeyHead {
	n := len(it.multi)
	rests := m.scrRests[:0]
	for range n {
		rests = append(rests, it.rest)
	}
	m.scrRests = rests
	defer clear(rests)
	if m.flat == 3 {
		return skHead(page.BuildStrings(rests, asStrSlice(it.multi))) // a nil page is a nil head
	}
	return fixedHead(page.BuildFixedOf(rests, it.multi, m.ptr))
}

// pageOf returns the multi-key page of the entries, or nil if they do not make one: they
// do not fit, or an entry or a value is beyond a limit of the page.
func (m *Map[T]) pageOf(items []item[T]) *header {
	n := len(items)
	cp := min(swar.Lcp(items[0].rest, items[n-1].rest), page.MaxKeyPart)
	slots, rem, vals := 0, -n*cp, 0
	for i := range items {
		it := &items[i]
		rem += len(it.rest)
		if it.multi == nil {
			slots++
			vals += len(strOfIf(m.flat == 3, it.val))
			continue
		}
		slots += len(it.multi)
		for _, v := range it.multi {
			vals += len(strOfIf(m.flat == 3, v))
		}
	}
	if m.flat == 3 { // 512 bytes hold no 255 slots, so the limit of n needs no check
		if page.NeedStrings(slots, cp, rem, vals) > 512 {
			return nil
		}
	} else if page.NeedFixed[T](slots, cp, rem) > 512 {
		return nil
	}
	rests := m.scrRests[:0]
	for i := range items {
		for range max(1, len(items[i].multi)) {
			rests = append(rests, items[i].rest)
		}
	}
	m.scrRests = rests
	defer clear(rests) // the scratch must not keep the keys alive
	if m.flat == 3 {
		vs := m.scrVals[:0]
		for i := range items {
			if items[i].multi == nil {
				vs = append(vs, view(strOf(items[i].val)))
				continue
			}
			for _, v := range items[i].multi {
				vs = append(vs, view(strOf(v)))
			}
		}
		m.scrVals = vs
		defer clear(vs)
		if p := page.BuildStrings(rests, vs); p != nil {
			return mkStrHdr(p)
		}
		return nil
	}
	vs := m.scrT[:0]
	for i := range items {
		if items[i].multi == nil {
			vs = append(vs, items[i].val)
		} else {
			vs = append(vs, items[i].multi...)
		}
	}
	m.scrT = vs
	defer clear(vs)
	if p := page.BuildFixedOf(rests, vs, m.ptr); p != nil {
		return mkFixHdr(p)
	}
	return nil
}

// strOfIf returns v as a string if str says T is string, else "".
func strOfIf[T comparable](str bool, v T) string {
	if str {
		return strOf(v)
	}
	return ""
}

// build returns the subtree for the entries, which are in key order, all different,
// and hold their keys from pathLen on: a single-key page for one entry, a multi-key
// page for entries that fit one, else a byte node on the next byte in which they
// differ, with a subtree for each byte (the entry whose key ends there is the
// node's end page). The items' rests are cut as the tree is built.
func (m *Map[T]) build(items []item[T], pathLen int) *header {
	n := len(items)
	if n == 1 {
		ev(evBuildSingleKey, 1)
		return singleKeyHdr(m.leafFor(&items[0]))
	}
	if h := m.pageOf(items); h != nil {
		ev(evBuildPage, n)
		return h
	}
	ev(evBuildNode, n)
	first, last := items[0].rest, items[n-1].rest
	end := swar.Lcp(first, last)
	nn := newNode(kN5, end)
	storePrefix(nn, first[:end])
	i := 0
	if len(first) == end { // the first key ends here: the end page of the node
		items[0].rest = first[end:]
		nn = setEndPage(nn, m.leafFor(&items[0]))
		i = 1
	}
	for i < n {
		b := items[i].rest[end]
		j := i + 1
		for j < n && items[j].rest[end] == b {
			j++
		}
		for k := i; k < j; k++ {
			items[k].rest = items[k].rest[end+1:]
		}
		nn, _ = addChild(nn, b, m.build(items[i:j], pathLen+end+1))
		i = j
	}
	return nn
}

// pageItems returns the entries of multi-key page n, with their whole keys from the
// path length of the page on, in key order, and room for one more. An entry with several
// values is one item with its values in multi: slices of one array, which has room for
// the values of the whole page.
func (m *Map[T]) pageItems(n *header) []item[T] {
	var items []item[T]
	var vals []T // the values of all entries, an entry's own next to each other
	lo := 0      // where the values of the last item begin in vals
	add := func(cp, rem []byte, buf *[]byte, v T, first bool) {
		if !first {
			vals = append(vals, v)
			items[len(items)-1].multi = vals[lo:len(vals):len(vals)]
			return
		}
		at := len(*buf)
		*buf = append(append(*buf, cp...), rem...)
		lo = len(vals)
		vals = append(vals, v)
		items = append(items, item[T]{rest: (*buf)[at:len(*buf):len(*buf)], val: v})
	}
	if m.flat == 3 {
		p := asMKStr(n)
		items, vals = make([]item[T], 0, p.Len()+1), make([]T, 0, p.Len())
		cp := p.CP()
		buf := make([]byte, 0, p.Len()*len(cp)+p.Used()) // room for every key: the keys are not longer than the page
		p.Each(func(rem, val []byte, first bool) bool {
			add(cp, rem, &buf, fromStr[T](string(val)), first)
			return true
		})
		return items
	}
	p := asMKFix(n)
	items, vals = make([]item[T], 0, p.Len()+1), make([]T, 0, p.Len())
	cp := p.CP()
	buf := make([]byte, 0, p.Len()*len(cp)+p.Used())
	p.Each(func(rem []byte, v T, first bool) bool {
		add(cp, rem, &buf, v, first)
		return true
	})
	return items
}

// appendValue gives items[i] a further value v.
func appendValue[T comparable](items []item[T], v T, i int) []item[T] {
	it := &items[i]
	if it.multi == nil {
		it.multi = []T{it.val}
	}
	it.multi = append(it.multi, v)
	return items
}

// reach is called by upsert: the descent for key ended at multi-key page n.
func (m *Map[T]) reach(loc **header, n *header, key []byte, pathLen int) **header {
	rest := key[pathLen:]
	var q *header
	var res page.Result
	if m.flat == 3 {
		p := asMKStr(n)
		if mis := p.Match(rest); mis < p.PrefixLen() {
			return m.outside(loc, n, key, pathLen, mis)
		}
		r, x := p.Add(rest, view(strOf(m.cur)))
		q, res = mkStrHdr(r), x
	} else {
		p := asMKFix(n)
		if mis := p.Match(rest); mis < p.PrefixLen() {
			return m.outside(loc, n, key, pathLen, mis)
		}
		r, x := p.Add(rest, m.cur, m.ptr)
		q, res = mkFixHdr(r), x
	}
	switch res {
	case page.Added:
		ev(evAdded, 1)
		*loc = q
		m.t.size++
	case page.AddedValue:
		ev(evAddedValue, 1)
		*loc = q
	case page.Full: // the page bursts
		fresh := !m.pageHas(n, pathLen, key)
		items := m.itemsWithNew(n, rest)
		ev(evBurst, len(items))
		*loc = m.build(items, pathLen)
		if fresh {
			m.t.size++
		}
	}
	return nil
}

// itemsWithNew returns the entries of multi-key page n and the new value rest -> cur, in key
// order: a further value of its entry if the page has the key, else a new entry.
func (m *Map[T]) itemsWithNew(n *header, rest []byte) []item[T] {
	items := m.pageItems(n)
	at := len(items)
	for i := range items {
		c := bytes.Compare(items[i].rest, rest)
		if c == 0 {
			return appendValue(items, m.cur, i)
		}
		if c > 0 {
			at = i
			break
		}
	}
	items = append(items, item[T]{})
	copy(items[at+1:], items[at:])
	items[at] = item[T]{rest: rest, val: m.cur}
	return items
}

// outside handles a new key that leaves the common prefix of multi-key page n after mis
// bytes: if it fits the page together with the page's entries, the page is built again with
// a shorter common prefix (Widen); if not, a byte node goes above the page (abovePage).
func (m *Map[T]) outside(loc **header, n *header, key []byte, pathLen, mis int) **header {
	rest := key[pathLen:]
	var q *header
	if m.flat == 3 {
		if p := asMKStr(n).Widen(rest, view(strOf(m.cur))); p != nil {
			q = mkStrHdr(p)
		}
	} else if p := asMKFix(n).Widen(rest, m.cur, m.ptr); p != nil {
		q = mkFixHdr(p)
	}
	if q == nil {
		return m.abovePage(loc, n, key, pathLen, mis)
	}
	ev(evWiden, 1)
	*loc = q
	m.t.size++
	return nil
}

// abovePage puts a byte node above multi-key page n, whose common prefix key leaves after
// mis bytes: the node takes those bytes, the page the next one as its byte and loses them,
// and the new key gets a single-key page below the node. It returns that page's slot.
func (m *Map[T]) abovePage(loc **header, n *header, key []byte, pathLen, mis int) **header {
	rest := key[pathLen:]
	var b byte
	ev(evAbove, 1)
	if m.flat == 3 {
		p := asMKStr(n)
		b = p.CP()[mis]
		p.Skip(mis + 1)
	} else {
		p := asMKFix(n)
		b = p.CP()[mis]
		p.Skip(mis + 1)
	}
	nn := newNode(kN5, mis)
	storePrefix(nn, rest[:mis])
	h, _ := addChild(nn, b, n)
	tail := rest[mis:]
	h, slot := attachAt(h, tail, m.newLeaf(tail[min(1, len(tail)):]))
	*loc = h
	m.t.size++
	return slot
}

// pageHas reports whether multi-key page n holds key, whose keys begin at pathLen.
func (m *Map[T]) pageHas(n *header, pathLen int, key []byte) bool {
	if m.flat == 3 {
		_, ok := asMKStr(n).Get(key[pathLen:])
		return ok
	}
	_, ok := asMKFix(n).Get[T](key[pathLen:])
	return ok
}

// pageEach calls yield for the values of key in multi-key page n, in the order they came in,
// until it returns false.
func (m *Map[T]) pageEach(n *header, pathLen int, key []byte, yield func(T) bool) {
	if m.flat == 3 {
		asMKStr(n).EachValue(key[pathLen:], func(val []byte) bool { return yield(fromStr[T](string(val))) })
		return
	}
	asMKFix(n).EachValue(key[pathLen:], yield)
}

// pageRemove removes the value v of the entry of key from multi-key page n, whose keys begin
// at pathLen, or all values of the entry if all is set. It returns the number of entries the
// page has left (counted up to mergeBelow+1), or 0 if it removed nothing. A page left with
// one entry becomes a single-key page, and a page that shrinks takes a smaller object.
func (m *Map[T]) pageRemove(n *header, pathLen int, key []byte, v T, all bool) int {
	rest := key[pathLen:]
	cur, removed := n, false
	for {
		q, res := m.removeFrom(cur, rest, v, all)
		if res == page.Absent {
			break
		}
		removed, cur = true, q
		if res == page.Gone {
			m.t.size--
			break
		}
		if !all {
			break
		}
	}
	if !removed {
		return 0
	}
	var left int
	if m.flat == 3 {
		left = asMKStr(cur).KeysUpTo(mergeBelow + 1)
	} else {
		left = asMKFix(cur).KeysUpTo(mergeBelow + 1)
	}
	ev(evPageRemove, left)
	if left == 1 { // the page has made itself a single-key page, in place
		ev(evToSingleKey, 1)
	}
	if cur != n {
		ev(evShrink, left)
		*m.t.findSlot(key) = cur
	}
	return left
}

// removeFrom removes the value v of the entry rest from multi-key page n, or, if first is set,
// the first value of the entry, whatever it is. It returns the page that holds the rest.
func (m *Map[T]) removeFrom(n *header, rest []byte, v T, first bool) (*header, page.Removal) {
	if m.flat == 3 {
		p := asMKStr(n)
		val := view(strOf(v))
		if first {
			got, ok := p.Get(rest)
			if !ok {
				return n, page.Absent
			}
			val = got
		}
		r, res := p.Remove(rest, val)
		return mkStrHdr(r), res
	}
	p := asMKFix(n)
	if first {
		got, ok := p.Get[T](rest)
		if !ok {
			return n, page.Absent
		}
		v = got
	}
	r, res := p.Remove(rest, v, m.ptr)
	return mkFixHdr(r), res
}

// rekeyPage is rekey for multi-key page l, which moves up: it takes front, the bytes the node above
// it no longer holds, in front of its common prefix. It returns nil if they do not fit.
func (m *Map[T]) rekeyPage(l *singleKeyHead, front []byte) *singleKeyHead {
	h := singleKeyHdr(l)
	if m.flat == 3 {
		if q := asMKStr(h).Prepend(front); q != nil {
			return asSingleKey(mkStrHdr(q))
		}
		ev(evPrependNo, 1)
		return nil
	}
	if q := asMKFix(h).Prepend[T](front, m.ptr); q != nil {
		return asSingleKey(mkFixHdr(q))
	}
	ev(evPrependNo, 1)
	return nil
}

// cmpParts compares the bytes of a followed by those of b with bound, without joining them.
func cmpParts(a, b, bound []byte) int {
	n := min(len(a), len(bound))
	if c := bytes.Compare(a[:n], bound[:n]); c != 0 {
		return c
	}
	if len(a) > len(bound) {
		return 1
	}
	return bytes.Compare(b, bound[len(a):])
}

// scanPage visits the entries of multi-key page n, whose keys begin at pathLen, within b:
// lo and hi say whether the page lies on the path of the lower and the upper bound, as for a
// leaf. It calls fn with the key (kb must be set) or, if fn is nil, yield with the value, and
// reports false once the scan is over.
func (m *Map[T]) scanPage(n *header, pathLen int, b *Bounds, lo, hi bool, kb *keyBuf, fn func(key []byte) bool, yield func(T) bool) bool {
	over, in := false, false // in: the key of the value before is within the bounds
	visit := func(cp, rem []byte, v T, first bool) bool {
		if first {
			in = true
			if lo {
				if c := cmpParts(cp, rem, b.From[pathLen:]); c < 0 || (c == 0 && !b.FromIncl) {
					in = false
				}
			}
			if in && hi {
				if c := cmpParts(cp, rem, b.To[pathLen:]); c > 0 || (c == 0 && !b.ToIncl) {
					over = true
					return false
				}
			}
		}
		if !in {
			return true
		}
		if fn != nil {
			if !first {
				return true
			}
			kb.key = append(append(append(kb.key[:0], kb.path[:pathLen]...), cp...), rem...)
			return fn(kb.key)
		}
		return yield(v)
	}
	var done bool
	if m.flat == 3 {
		p := asMKStr(n)
		cp := p.CP()
		done = p.Each(func(rem, val []byte, first bool) bool {
			var v T
			if fn == nil {
				v = fromStr[T](string(val))
			}
			return visit(cp, rem, v, first)
		})
	} else {
		p := asMKFix(n)
		cp := p.CP()
		done = p.Each(func(rem []byte, v T, first bool) bool { return visit(cp, rem, v, first) })
	}
	return done && !over
}

// mergeLimit is the most entries that mergeFits adds up before it gives up: a page of 512
// bytes holds that many only if its entries are a few bytes each. mergeChildren is the most
// byte children a node may have for tryMerge to look at it, the capacity of a 12-way node:
// every child holds an entry at least, and more than twelve of them rarely fit one page, while
// looking at every child of a wide node on every removal costs more than the merge saves.
const (
	mergeLimit    = 64
	mergeChildren = 12
)

// mergeBelow is how few entries a page may have left after the removal of one of its entries
// for the merge of the nodes above it to be tried. Pages in use hold some ten entries and
// their siblings as many, so that nearly every try fails (96 % in the benchmark's churn of
// street names, docs/redesign/step4-probe.md); the page that is nearly empty is where
// a merge can succeed. In the churn, trying only then took a third off the time per
// operation, and a tree that had half of its keys removed value by value took 3 % more memory.
const mergeBelow = 2

// mergeFits reports whether the children of byte node n, which has prefix pre, are all pages
// (value overflows excepted) whose entries fit half of the largest page together (mergeFill), without building anything: it
// adds up the sizes the merged page would take. The node branches, so the common prefix of the
// merged page is pre, exactly.
func (m *Map[T]) mergeFits(n *header, pre []byte) bool {
	slots, keys, sumRest, sumVal, ok := 0, 0, 0, 0, true
	collect := func(extra int, c *header) {
		if !ok || !isPage(c.objType) {
			ok = false
			return
		}
		if isMultiKey(c.objType) {
			if m.flat == 3 {
				p := asMKStr(c)
				cp := p.PrefixLen()
				p.Each(func(rem, val []byte, first bool) bool {
					if first {
						keys++
						sumRest += len(pre) + extra + cp + len(rem)
					}
					sumVal += len(val)
					return true
				})
				slots += p.Len()
			} else {
				p := asMKFix(c)
				cp := p.PrefixLen()
				p.Each(func(rem []byte, _ T, first bool) bool {
					if first {
						keys++
						sumRest += len(pre) + extra + cp + len(rem)
					}
					return true
				})
				slots += p.Len()
			}
		} else {
			l := asSingleKey(c)
			if l.isValueOverflow() {
				ok = false
				return
			}
			sumRest += len(pre) + extra + len(l.stored())
			if m.flat == 3 {
				asSK(l).Each(func(_, val []byte, _ bool) bool { sumVal += len(val); return true })
			}
			slots += int(l.n)
			keys++
		}
		ok = ok && slots <= mergeLimit
	}
	if e := endPageOf(n); e != nil {
		collect(0, singleKeyHdr(e))
	}
	eachByteNode(n, func(_ byte, c *header) { collect(1, c) })
	if !ok || keys < 2 {
		return false
	}
	cp := min(len(pre), page.MaxKeyPart)
	if m.flat == 3 {
		return page.NeedStrings(slots, cp, sumRest-keys*cp, sumVal) <= mergeFill
	}
	return page.NeedFixed[T](slots, cp, sumRest-keys*cp) <= mergeFill
}

// mergeFill is the most bytes a merged page may need: half of the largest page, the size at which a page bursts.
// A page bursts only beyond the largest class (below it, it moves from class to class with the shrink rule of
// internal/page), so the distance to the next burst is measured from the largest page whatever class the merged
// page lands in, and half of it leaves room for as many entries again. A merge into a page that is nearly full
// made the same subtrees burst and merge back and forth (8 bursts and 8 merges of some 25 entries in 1,000
// operations of a small single-value tree); with half of it as the limit there are none, and the churn of small
// trees is 10 to 40 % faster (docs/redesign/review-2026-10.md, E6). The same half brings a value overflow back
// into a page (page.BackFits).
const mergeFill = page.Largest / 2

// leafItem returns the entry of single-key page l, which is not a value overflow, with its
// values, for the key rest.
func (m *Map[T]) leafItem(l *singleKeyHead, rest []byte) item[T] {
	it := item[T]{rest: rest}
	add := func(v T) {
		if l.n == 1 {
			it.val = v
			return
		}
		if it.multi == nil {
			it.val, it.multi = v, make([]T, 0, l.n)
		}
		it.multi = append(it.multi, v)
	}
	if m.flat == 3 {
		asSK(l).Each(func(_, val []byte, _ bool) bool { add(fromStr[T](string(val))); return true })
		return it
	}
	asFixed(l).Each(func(_ []byte, v T, _ bool) bool { add(v); return true })
	return it
}

// tryMerge replaces the byte node at *loc, whose path begins at pathLen, by one multi-key page
// if every child is a page and their entries together fit half of the largest page (mergeFits).
// It reports whether it did. Only a map with multi-key pages calls it.
func (m *Map[T]) tryMerge(loc **header, pathLen int) bool {
	n := *loc
	if n.count > mergeChildren {
		return false
	}
	var buf [prefixBuf]byte
	pre := appendPrefix(buf[:0], n)
	ev(evMergeTry, int(n.count))
	if !m.mergeFits(n, pre) {
		return false
	}
	var items []item[T]
	collect := func(b int, c *header) {
		front := append(make([]byte, 0, len(pre)+1), pre...)
		if b >= 0 {
			front = append(front, byte(b))
		}
		if isMultiKey(c.objType) {
			for _, it := range m.pageItems(c) {
				it.rest = append(slices.Clone(front), it.rest...)
				items = append(items, it)
			}
			return
		}
		l := asSingleKey(c)
		items = append(items, m.leafItem(l, append(front, l.stored()...)))
	}
	if e := endPageOf(n); e != nil {
		collect(-1, singleKeyHdr(e))
	}
	eachByteNode(n, func(b byte, c *header) { collect(int(b), c) })
	// mergeFits allowed at most mergeFill bytes, which keeps every remainder, value and slot count within the
	// page's own limits (each of them alone would take more): pageOf builds the page.
	ev(evMergeOK, len(items))
	*loc = m.pageOf(items)
	return true
}

// mergeUp tries to merge the nodes on the path of key, from the lowest up, after a removal
// below them: a node becomes a page when its children are pages that fit one together, and
// then its parent may do the same. It reports whether *loc is now a page.
func (m *Map[T]) mergeUp(loc **header, key []byte, pathLen int) bool {
	n := *loc
	if n == nil || isPage(n.objType) {
		return true
	}
	pl := n.prefixLen()
	if pl > 0 && (pathLen+pl > len(key) || !prefixMatches(n, pl, key, pathLen)) {
		return m.tryMerge(loc, pathLen) // the key's path left this node's: a collapse below it gave it more bytes
	}
	d := pathLen + pl
	var c **header
	if d == len(key) {
		c = endPageSlot(n)
	} else {
		c = findLoc(n, key[d])
	}
	if c != nil && *c != nil && !m.mergeUp(c, key, d+1) {
		return false
	}
	if n.count == 1 && endPageOf(n) == nil {
		if _, only := onlyChild(n); isSingleKey(only.objType) {
			// a page that could not move up shrank to one key: the single-key page can
			*loc = collapse(n, key, pathLen, d, m.rekey)
			return isPage((*loc).objType)
		}
	}
	return m.tryMerge(loc, pathLen)
}
