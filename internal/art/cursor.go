package art

import (
	"bytes"
	"math/bits"
	"unsafe"

	set3 "github.com/TomTonic/Set3"

	"github.com/TomTonic/multimap/internal/page"
	"github.com/TomTonic/multimap/internal/swar"
)

// Cursor walks the keys or values of a map within bounds in ascending key order, one page at a time
// (docs/redesign/scan-design.md). NextPage does the walk of the tree without recursion, on a stack of the
// byte nodes on the way, and leaves the values (or the keys) of the next page that lies within the bounds in
// Vals (or Keys); the caller's loop over them is a plain loop, so that a range scan makes no call for a
// value:
//
//	var c Cursor[T]
//	c.Init(m, b, false)
//	for c.NextPage() {
//		for _, v := range c.Vals { ... }
//	}
//
// Vals may alias a page of the map and Keys a buffer of the cursor: both are valid until the next NextPage, and
// the map must not change during the walk.
type Cursor[T comparable] struct {
	Vals []T           // the values of the current page within the bounds
	Set  *set3.Set3[T] // the values of the current leaf if it is a value overflow (Vals is empty then)
	Keys [][]byte      // the keys of the current page within the bounds (keys only)

	m      *Map[T]
	b      Bounds // a copy: a pointer kept by the cursor would make the caller's bounds escape to the heap
	keys   bool   // the walk reports keys, not values
	state  uint8  // curStart, curRunning, curLast, curDone
	depth  int
	frames [32]frame
	more   []frame // the frames beyond the 32 of a deep tree

	scr  []T    // the values of a page that are not an array in it: strings, a value overflow's set
	path []byte // the path of the walk (keys only)
	kbuf []byte // the bytes of the keys of the current page (keys only)
	ends []int  // the end of each key in kbuf
	pb   page.Bounds
}

// frame is a byte node on the stack of a cursor: the walk visits its children with bytes from next to hiB. It is
// kept small (24 bytes), since the cursor holds 32 of them and is cleared for every walk.
type frame struct {
	n        *header
	pathLen  int32  // the path length of n's children, n's common prefix included and their own byte not
	next     uint16 // the next child byte to look at (256: none left)
	i        uint8  // the position of that child in the child array of a 5-, 12-, 26- or 58-way node
	loB, hiB byte
	lo, hi   bool // n lies on the path of the lower and of the upper bound
}

const (
	curStart   = iota // nothing walked yet
	curRunning        // walking
	curLast           // the page given last was the last one within the bounds
	curDone           // the walk is over
)

// Init prepares c for a walk of m within b: of the keys if keys is set, else of the values.
func (c *Cursor[T]) Init(m *Map[T], b *Bounds, keys bool) {
	c.m, c.b, c.keys, c.state, c.depth = m, *b, keys, curStart, 0
}

// NextPage moves to the next page (or leaf) with keys within the bounds and leaves its values in Vals (or its keys in
// Keys); it reports false when there is none. It is the slow path of the walk: once a page, not once a value.
func (c *Cursor[T]) NextPage() bool {
	c.Set = nil
	switch c.state {
	case curStart:
		c.state = curRunning
		if c.enter(c.m.t.root, 0, c.b.HasFrom, c.b.HasTo) {
			return true
		}
	case curLast:
		c.state = curDone
	}
	for c.state == curRunning && c.depth > 0 {
		f := c.frame(c.depth - 1)
		k, child, ok := f.nextChild()
		if !ok {
			c.depth--
			if f.hi { // on To's path, everything after the child for hiB is above To
				c.state = curDone
			}
			continue
		}
		if c.keys {
			c.path = append(c.path[:f.pathLen], k)
		}
		if c.enter(child, int(f.pathLen)+1, f.lo && k == f.loB, f.hi && k == f.hiB) {
			return true
		}
	}
	c.state = curDone
	c.Vals, c.Keys = nil, nil
	return false
}

// frame returns the frame at depth i.
func (c *Cursor[T]) frame(i int) *frame {
	if i < len(c.frames) {
		return &c.frames[i]
	}
	return &c.more[i-len(c.frames)]
}

// push puts a frame on the stack and returns it, to be filled in place.
func (c *Cursor[T]) push() *frame {
	c.depth++
	if c.depth <= len(c.frames) {
		return &c.frames[c.depth-1]
	}
	if i := c.depth - 1 - len(c.frames); i < len(c.more) {
		return &c.more[i]
	}
	c.more = append(c.more, frame{})
	return &c.more[len(c.more)-1]
}

// enter starts the walk of the subtree n at pathLen, which lies on the path of the lower bound if lo is set and
// of the upper bound if hi is set: a page or leaf within the bounds is reported at once (true), a byte node goes
// on the stack after its end page has been reported. It sets the state to curDone when the subtree lies above the
// upper bound, and to curLast when the page reported is the last within the bounds.
func (c *Cursor[T]) enter(n *header, pathLen int, lo, hi bool) bool {
	if n == nil {
		return false
	}
	b := &c.b
	if isPage(n.objType) {
		if isSingleKey(n.objType) {
			return c.leaf(asSingleKey(n), pathLen, lo, hi)
		}
		return c.page(n, pathLen, lo, hi)
	}
	pl := n.prefixLen()
	if (lo || hi) && pl > 0 {
		if lo {
			rest := b.From[pathLen:]
			if m, ch := prefixLcp(n, pl, rest); m < pl {
				if m < len(rest) && ch < rest[m] {
					return false // the whole subtree lies below From
				}
				lo = false // the whole subtree lies above From
			}
		}
		if hi {
			rest := b.To[pathLen:]
			if m, ch := prefixLcp(n, pl, rest); m < pl {
				if m == len(rest) || ch > rest[m] {
					c.state = curDone // the whole subtree lies above To
					return false
				}
				hi = false // the whole subtree lies below To
			}
		}
	}
	if c.keys {
		c.path = appendPrefix(c.path[:pathLen], n)
	}
	pathLen += pl
	// The end page's key is the path to n. On From's path it is below From unless the path is From itself; on
	// To's path it is below To unless the path is To itself, in which case it is the last key in range.
	endPageIsFrom := lo && pathLen == len(b.From)
	endPageIsTo := hi && pathLen == len(b.To)
	if endPageIsFrom {
		lo = false
	}
	found := false
	if t := endPageOf(n); t != nil && !lo && (!endPageIsFrom || b.FromIncl) && (!endPageIsTo || b.ToIncl) {
		found = c.leaf(t, pathLen, false, false)
	}
	if endPageIsTo {
		c.state = curDone
		if found {
			c.state = curLast
		}
		return found
	}
	var loB, hiB byte = 0, 255
	if lo {
		loB = b.From[pathLen]
	}
	if hi {
		hiB = b.To[pathLen]
	}
	touchChildren(n, loB, hiB, leafTail)
	f := c.push()
	f.n, f.pathLen, f.next, f.i, f.loB, f.hiB, f.lo, f.hi = n, int32(pathLen), uint16(loB), uint8(firstChild(n, loB)), loB, hiB, lo, hi
	return found
}

// leaf reports leaf l at pathLen if its key lies within the bounds (on a bound's path the key agrees with the bound
// up to pathLen, so its key part decides): its values, or its key. It sets the state to curDone if the key lies above
// the upper bound.
func (c *Cursor[T]) leaf(l *singleKeyHead, pathLen int, lo, hi bool) bool {
	b := &c.b
	if lo {
		if r := bytes.Compare(l.stored(), b.From[pathLen:]); r < 0 || (r == 0 && !b.FromIncl) {
			return false
		}
	}
	if hi {
		if r := bytes.Compare(l.stored(), b.To[pathLen:]); r > 0 || (r == 0 && !b.ToIncl) {
			c.state = curDone
			return false
		}
	}
	if c.keys {
		c.kbuf = append(append(c.kbuf[:0], c.path[:pathLen]...), l.stored()...)
		c.Keys = append(c.Keys[:0], c.kbuf)
		return true
	}
	switch {
	case l.isValueOverflow():
		c.Vals, c.Set = nil, *overflowSetOf[T](l)
	case c.m.flat == 3:
		c.Vals, _ = c.strings(asSK(l), &page.Bounds{})
	default:
		c.Vals, _ = asFixed(l).ValuesIn[T](&page.Bounds{})
	}
	return true
}

// page reports the keys of multi-key page n at pathLen that lie within the bounds, if any: their values, or the keys.
func (c *Cursor[T]) page(n *header, pathLen int, lo, hi bool) bool {
	b := &c.b
	c.pb = page.Bounds{HasLo: lo, HasHi: hi, LoIncl: b.FromIncl, HiIncl: b.ToIncl}
	if lo {
		c.pb.Lo = b.From[pathLen:]
	}
	if hi {
		c.pb.Hi = b.To[pathLen:]
	}
	var over bool
	found := false
	switch {
	case c.keys:
		c.kbuf, c.ends = c.kbuf[:0], c.ends[:0]
		h := (*page.Fixed)(unsafe.Pointer(n))
		c.kbuf, c.ends, over = h.AppendKeys(c.kbuf, c.ends, c.path[:pathLen], &c.pb, c.m.flat == 3)
		c.Keys = c.Keys[:0]
		start := 0
		for _, e := range c.ends {
			c.Keys = append(c.Keys, c.kbuf[start:e])
			start = e
		}
		found = len(c.Keys) > 0
	case c.m.flat == 3:
		c.Vals, over = c.strings(asMKStr(n), &c.pb)
		found = len(c.Vals) > 0
	default:
		c.Vals, over = asMKFix(n).ValuesIn[T](&c.pb)
		found = len(c.Vals) > 0
	}
	if over {
		c.state = curDone
		if found {
			c.state = curLast
		}
	}
	return found
}

// strings returns the values of string page p within pb as the cursor's scratch slice of T (T is string), and
// whether a key above the upper bound follows.
func (c *Cursor[T]) strings(p *page.Str, pb *page.Bounds) ([]T, bool) {
	s := (*[]string)(unsafe.Pointer(&c.scr))
	var over bool
	*s, over = p.AppendStrings((*s)[:0], pb)
	return c.scr, over
}

// firstChild returns the position in the child array of byte node n of the first child with a byte of at least lo
// (5-, 12-, 26- and 58-way nodes; the 256-way node is indexed by the byte).
func firstChild(n *header, lo byte) int {
	switch n.objType {
	case kN26, kN58:
		bm, _ := bitmapOf(n)
		return swar.Rank(bm, lo)
	case kN256:
		return 0
	}
	keys, _ := sorted(n)
	i := 0
	for i < len(keys) && keys[i] < lo {
		i++
	}
	return i
}

// nextChild returns the next child of the frame's node with a byte up to hiB, and moves past it: the walk goes
// through the children of a node one after the other, so each step is a step, not a search.
func (f *frame) nextChild() (byte, *header, bool) {
	n, next, hi := f.n, int(f.next), int(f.hiB)
	switch n.objType {
	case kN26, kN58:
		bm, child := bitmapOf(n)
		for w := next >> 6; next <= hi && w < 4; w++ {
			set := bm[w]
			if w == next>>6 {
				set &= ^uint64(0) << (next & 63)
			}
			if set != 0 {
				k := w<<6 + bits.TrailingZeros64(set)
				if k > hi {
					break
				}
				f.next = uint16(k + 1)
				f.i++
				return byte(k), child[f.i-1], true
			}
		}
		f.next = uint16(hi + 1)
		return 0, nil, false
	case kN256:
		x := asN256(n)
		for k := next; k <= hi; k++ {
			if x.child[k] != nil {
				f.next = uint16(k + 1)
				return byte(k), x.child[k], true
			}
		}
		f.next = uint16(hi + 1)
		return 0, nil, false
	}
	keys, child := sorted(n)
	if i := int(f.i); i < len(keys) && keys[i] <= f.hiB {
		f.i++
		return keys[i], child[i], true
	}
	return 0, nil, false
}
