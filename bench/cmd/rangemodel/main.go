// Command rangemodel models the routing layer of step 6 (docs/redesign/step6-design.md): it builds, from the sorted
// keys of a corpus in the bench's build order, the tree of today (a page that overflows bursts into a byte node with
// one child per byte) and the tree of the proposal (byte nodes whose children may be ranges of bytes; a page that
// overflows splits in two at a byte boundary), and counts for every key what its lookup touches: byte nodes, the
// remainders a page search compares (walking the page as today, or by a binary search over prefix sums of the key
// lengths), and the 64-byte cache lines of the nodes and the page, the value's included. It counts, it does not time:
// speed comes only from rtcompare (MEASURING.md).
//
//	go run ./cmd/rangemodel -keys street,dirs,links,url -values single-value,natural -sizes 4096,16384,65536
//
// Simplifications, stated in the output's notes: a key is inserted with all its values at once (the bench adds them
// one after another, so a key's values may make a page burst earlier there); nothing is removed (the tree as built
// from the corpus, like the shape probe MKSHAPE); value sets of value overflows and prefix tails are not counted in
// the bytes; values are uint64 unless -str.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/page"
)

// item is a key with its values: their number and their bytes (8 a value for uint64, the lengths of the strings).
type item struct {
	key  []byte
	vals int
	vb   int
}

// kind of a child slot.
type kind uint8

const (
	kNil  kind = iota // an empty range: no object
	kPage             // a page (one key or many)
	kNode             // a byte node
	kBig              // a key in an object of its own (a value overflow)
)

type obj struct {
	k   kind
	p   *mpage
	n   *mnode
	big *item
}

// mpage is a page that holds its keys from position from on, in key order.
type mpage struct {
	from  int
	items []*item
}

// mnode is a byte node at path length depth with a common prefix of len(prefix) bytes; it branches on the byte at
// depth+len(prefix). Today's nodes have one start a child (the child of that byte); the proposal's child i covers the
// bytes [starts[i], starts[i+1]), and a node child covers exactly its start byte.
type mnode struct {
	depth  int
	prefix []byte
	end    obj
	starts []int
	kids   []obj
}

func (n *mnode) q() int { return n.depth + len(n.prefix) }

// model is one tree under one policy.
type model struct {
	ranges bool // the proposal; else today's burst
	limit  int  // the size above which a page of a range splits (the proposal); today: page.Largest
	str    bool // string values
	root   obj
}

func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
	return i
}

var sizes = []int{32, 64, 128, 256, 384, 512}

func classFor(need int) int {
	for _, s := range sizes {
		if need <= s {
			return s
		}
	}
	return -1
}

// need returns the bytes of a page of the items from position from, and false if they make no page (a limit of the
// page: a remainder, a key part, the number of values, or content beyond the largest class).
func (m *model) need(items []*item, from int) (int, bool) {
	if len(items) == 1 {
		it := items[0]
		r := len(it.key) - from
		need := page.Header + r + it.vb
		if m.str {
			need += it.vals
		}
		return need, r <= page.MaxKeyPart && it.vals <= page.MaxEntries && need <= page.Largest
	}
	cpl := min(lcp(items[0].key[from:], items[len(items)-1].key[from:]), page.MaxKeyPart)
	n, rem, vb := 0, 0, 0
	for _, it := range items {
		r := len(it.key) - from - cpl
		if r > page.MaxRemainder {
			return 0, false
		}
		rem += r
		n += it.vals
		vb += it.vb
	}
	if n > page.MaxEntries {
		return 0, false
	}
	var need int
	if m.str {
		need = page.NeedStrings(n, cpl, rem, vb)
	} else {
		need = page.NeedFixed[uint64](n, cpl, rem)
	}
	return need, need <= page.Largest
}

// fits reports whether the items make a page under this model's limit (one key: any page up to the largest class).
func (m *model) fits(items []*item, from int) bool {
	need, ok := m.need(items, from)
	if !ok {
		return false
	}
	return len(items) == 1 || need <= m.limit
}

func (m *model) leaf(it *item, from int) obj {
	if _, ok := m.need([]*item{it}, from); !ok {
		return obj{k: kBig, big: it}
	}
	return obj{k: kPage, p: &mpage{from: from, items: []*item{it}}}
}

func itemsOf(o obj) []*item {
	switch o.k {
	case kPage:
		return o.p.items
	case kBig:
		return []*item{o.big}
	}
	return nil
}

func withItem(items []*item, it *item) []*item {
	out := make([]*item, 0, len(items)+1)
	i := sort.Search(len(items), func(i int) bool { return bytes.Compare(items[i].key, it.key) >= 0 })
	out = append(out, items[:i]...)
	out = append(out, it)
	return append(out, items[i:]...)
}

// --- today: burst into one child a byte ---

func (m *model) build(items []*item, depth int) obj {
	if len(items) == 1 {
		return m.leaf(items[0], depth)
	}
	if m.fits(items, depth) {
		return obj{k: kPage, p: &mpage{from: depth, items: items}}
	}
	end := lcp(items[0].key[depth:], items[len(items)-1].key[depth:])
	nd := &mnode{depth: depth, prefix: slices.Clone(items[0].key[depth : depth+end])}
	q := nd.q()
	i := 0
	if len(items[0].key) == q {
		nd.end = m.leaf(items[0], q)
		i = 1
	}
	for i < len(items) {
		b := items[i].key[q]
		j := i + 1
		for j < len(items) && items[j].key[q] == b {
			j++
		}
		nd.starts = append(nd.starts, int(b))
		nd.kids = append(nd.kids, m.build(slices.Clone(items[i:j]), q+1))
		i = j
	}
	return obj{k: kNode, n: nd}
}

// --- the proposal: ranges ---

// newNode lays out items, which share their keys up to depth, as a node with ranges.
func (m *model) newNode(items []*item, depth int) obj {
	end := lcp(items[0].key[depth:], items[len(items)-1].key[depth:])
	nd := &mnode{depth: depth, prefix: slices.Clone(items[0].key[depth : depth+end])}
	q := nd.q()
	if len(items[0].key) == q {
		nd.end = m.leaf(items[0], q)
		items = items[1:]
	}
	nd.starts, nd.kids = m.layout(items, q, int(items[0].key[q]), 256)
	return obj{k: kNode, n: nd}
}

// layout returns the ranges that hold the items (their keys have their byte at q in [lo, hi)), covering [lo, hi):
// one page if they fit; else a group of one byte that cannot be a page gets that byte as a range of its own (a node,
// or a key in an object of its own), and the rest is cut in two at the byte boundary nearest the middle of the bytes.
func (m *model) layout(items []*item, q, lo, hi int) ([]int, []obj) {
	if len(items) == 0 {
		return []int{lo}, []obj{{k: kNil}}
	}
	if m.fits(items, q) {
		return []int{lo}, []obj{{k: kPage, p: &mpage{from: q, items: items}}}
	}
	var cuts []int // indexes where the byte at q changes
	for i := 1; i < len(items); i++ {
		if items[i].key[q] != items[i-1].key[q] {
			cuts = append(cuts, i)
		}
	}
	if len(cuts) == 0 { // one byte: a range of its own
		b := int(items[0].key[q])
		var starts []int
		var kids []obj
		if lo < b {
			starts, kids = append(starts, lo), append(kids, obj{k: kNil})
		}
		starts = append(starts, b)
		if len(items) == 1 {
			kids = append(kids, m.leaf(items[0], q)) // a key that makes no page
		} else {
			kids = append(kids, m.newNode(items, q+1))
		}
		if b+1 < hi {
			starts, kids = append(starts, b+1), append(kids, obj{k: kNil})
		}
		return starts, kids
	}
	total := 0
	for _, it := range items {
		total += len(it.key) - q + it.vb + 2
	}
	best, bestD, acc, at := cuts[0], 1<<62, 0, 0
	for _, c := range cuts {
		for at < c {
			acc += len(items[at].key) - q + items[at].vb + 2
			at++
		}
		if d := abs(2*acc - total); d < bestD {
			best, bestD = c, d
		}
	}
	mid := int(items[best].key[q])
	s1, k1 := m.layout(items[:best], q, lo, mid)
	s2, k2 := m.layout(items[best:], q, mid, hi)
	return append(s1, s2...), append(k1, k2...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// floor returns the index of the last start at most b, or -1.
func floor(starts []int, b int) int {
	return sort.Search(len(starts), func(i int) bool { return starts[i] > b }) - 1
}

// --- insert, both policies ---

func (m *model) insert(it *item) {
	slot, depth := &m.root, 0
	for {
		switch slot.k {
		case kNil:
			*slot = m.leaf(it, depth)
			return
		case kPage, kBig:
			items := withItem(itemsOf(*slot), it)
			if m.fits(items, depth) {
				*slot = obj{k: kPage, p: &mpage{from: depth, items: items}}
			} else if m.ranges {
				*slot = m.newNode(items, depth)
			} else {
				*slot = m.build(items, depth)
			}
			return
		}
		nd := slot.n
		if mis := lcp(it.key[depth:], nd.prefix); mis < len(nd.prefix) {
			m.splitPrefix(slot, it, depth, mis)
			return
		}
		q := nd.q()
		if len(it.key) == q {
			nd.end = m.leaf(it, q)
			return
		}
		b := int(it.key[q])
		if !m.ranges {
			i := sort.SearchInts(nd.starts, b)
			if i == len(nd.starts) || nd.starts[i] != b {
				nd.starts = slices.Insert(nd.starts, i, b)
				nd.kids = slices.Insert(nd.kids, i, m.leaf(it, q+1))
				return
			}
			slot, depth = &nd.kids[i], q+1
			continue
		}
		i := floor(nd.starts, b)
		if i < 0 { // below the first range: widen a first page or empty range, or put a new range in front
			if k := nd.kids[0].k; k == kNil || k == kPage {
				nd.starts[0] = b
				i = 0
			} else {
				nd.starts = slices.Insert(nd.starts, 0, b)
				nd.kids = slices.Insert(nd.kids, 0, obj{k: kNil})
				i = 0
			}
		}
		kid := &nd.kids[i]
		hi := 256
		if i+1 < len(nd.starts) {
			hi = nd.starts[i+1]
		}
		switch kid.k {
		case kNode:
			slot, depth = kid, q+1
			continue
		case kNil:
			*kid = m.leaf(it, q)
			if kid.k == kBig && hi-nd.starts[i] > 1 { // a key in an object of its own gets a range of one byte
				m.replaceRange(nd, i, []*item{it}, q, hi)
			}
			return
		}
		items := withItem(itemsOf(*kid), it)
		if kid.k == kPage && m.fits(items, q) {
			kid.p = &mpage{from: q, items: items}
			return
		}
		m.replaceRange(nd, i, items, q, hi)
		return
	}
}

// replaceRange lays out range i of nd anew with items.
func (m *model) replaceRange(nd *mnode, i int, items []*item, q, hi int) {
	s, k := m.layout(items, q, nd.starts[i], hi)
	nd.starts = slices.Replace(nd.starts, i, i+1, s...)
	nd.kids = slices.Replace(nd.kids, i, i+1, k...)
}

// splitPrefix puts a node above node *slot, whose common prefix the key of it leaves after mis bytes.
func (m *model) splitPrefix(slot *obj, it *item, depth, mis int) {
	old := slot.n
	ob := int(old.prefix[mis])
	top := &mnode{depth: depth, prefix: slices.Clone(old.prefix[:mis])}
	old.depth, old.prefix = depth+mis+1, slices.Clone(old.prefix[mis+1:])
	q := top.q()
	if len(it.key) == q {
		top.end = m.leaf(it, q)
		top.starts, top.kids = []int{ob}, []obj{{k: kNode, n: old}}
		if m.ranges && ob < 255 { // a node child covers its byte only
			top.starts, top.kids = append(top.starts, ob+1), append(top.kids, obj{k: kNil})
		}
		*slot = obj{k: kNode, n: top}
		return
	}
	nb := int(it.key[q])
	if !m.ranges {
		top.starts, top.kids = []int{ob}, []obj{{k: kNode, n: old}}
		i := sort.SearchInts(top.starts, nb)
		top.starts = slices.Insert(top.starts, i, nb)
		top.kids = slices.Insert(top.kids, i, m.leaf(it, q+1))
		*slot = obj{k: kNode, n: top}
		return
	}
	newPage := m.leaf(it, q)
	if nb < ob {
		top.starts = []int{nb, ob}
		top.kids = []obj{newPage, {k: kNode, n: old}}
		if ob < 255 {
			top.starts, top.kids = append(top.starts, ob+1), append(top.kids, obj{k: kNil})
		}
	} else {
		top.starts = []int{ob, ob + 1}
		top.kids = []obj{{k: kNode, n: old}, newPage}
	}
	if newPage.k == kBig { // a key of its own object needs a range of one byte
		for i, s := range top.starts {
			if s == nb || (top.kids[i].k == kBig) {
				hi := 256
				if i+1 < len(top.starts) {
					hi = top.starts[i+1]
				}
				if hi-s > 1 {
					m.replaceRange(top, i, []*item{it}, q, hi)
				}
				break
			}
		}
	}
	*slot = obj{k: kNode, n: top}
}

// --- counting ---

type counts struct {
	keys, lookups                  int
	nodesPassed                    int
	cmpLin, cmpBin                 int
	linesLin, linesBin             int
	pages, onePages, bigs, nodes   int
	pageKeys                       []int
	nodeBytes, pageBytes, bigBytes int
	// by node class (N5, N12, N26, N58, N256): nodes passed, and those whose child slot lies outside the node's
	// first line of 64 and of 128 bytes (a second dependent line inside the node); prefix tails read
	visits, second64, second128 [5]int
	tails                       int
}

// classIdx returns the index of the node class of capacity cap in counts.visits.
func classIdx(cap int) int {
	switch cap {
	case 5:
		return 0
	case 12:
		return 1
	case 26:
		return 2
	case 58:
		return 3
	}
	return 4
}

// slotLine counts node visit of class cap whose child slot is at offset off.
func (c *counts) slotLine(cap, off int) {
	k := classIdx(cap)
	c.visits[k]++
	if off+8 > 64 {
		c.second64[k]++
	}
	if off+8 > 128 {
		c.second128[k]++
	}
}

func nodeCap(c int) (cap, size int) {
	switch {
	case c <= 5:
		return 5, 64
	case c <= 12:
		return 12, 128
	case c <= 26:
		return 26, 256
	case c <= 58:
		return 58, 512
	}
	return 257, 2080
}

// slotOff returns the offset of child slot i (or of the end page, i = -1) in a node of class cap holding
// slots children; b is the byte (the 256-way node indexes by it).
func slotOff(cap, i, b int) int {
	if i < 0 {
		i = cap - 1
		b = 256
	}
	switch cap {
	case 5:
		return 24 + 8*i
	case 12:
		return 32 + 8*i
	case 26, 58:
		return 48 + 8*i
	}
	return 24 + 8*b
}

func (m *model) census(o obj, c *counts) {
	switch o.k {
	case kPage:
		need, _ := m.need(o.p.items, o.p.from)
		s := classFor(need)
		c.pages++
		c.pageBytes += s
		c.pageKeys = append(c.pageKeys, len(o.p.items))
		if len(o.p.items) == 1 {
			c.onePages++
		}
	case kBig:
		c.bigs++
		c.bigBytes += 64
	case kNode:
		n := o.n
		c.nodes++
		cnt := len(n.starts)
		if n.end.k != kNil {
			cnt++
		}
		_, size := nodeCap(cnt)
		if m.ranges && size == 2080 {
			size += 32 // the bitmap of the range starts
		}
		c.nodeBytes += size
		m.census(n.end, c)
		for _, k := range n.kids {
			m.census(k, c)
		}
	}
}

// lookup counts what the lookup of it touches.
func (m *model) lookup(it *item, c *counts) {
	o, depth := m.root, 0
	lines := 0
	for o.k == kNode {
		n := o.n
		c.nodesPassed++
		cnt := len(n.starts)
		if n.end.k != kNil {
			cnt++
		}
		cp, _ := nodeCap(cnt)
		ls := map[int]bool{0: true}
		if len(n.prefix) > 12 {
			lines++ // the tail of the prefix
			c.tails++
		}
		q := n.q()
		if len(it.key) == q {
			ls[slotOff(cp, -1, 0)/64] = true
			c.slotLine(cp, slotOff(cp, -1, 0))
			lines += len(ls)
			o, depth = n.end, q
			break
		}
		b := int(it.key[q])
		var i int
		if m.ranges {
			i = floor(n.starts, b)
		} else {
			i = sort.SearchInts(n.starts, b)
		}
		ls[slotOff(cp, i, n.starts[i])/64] = true
		c.slotLine(cp, slotOff(cp, i, n.starts[i]))
		lines += len(ls)
		if n.kids[i].k == kNode {
			if n.starts[i] != b {
				panic(fmt.Sprintf("model: the key %q reaches the node of byte %d with byte %d", it.key, n.starts[i], b))
			}
			depth = q + 1
		} else {
			depth = q
			if !m.ranges {
				depth = q + 1
			}
		}
		o = n.kids[i]
	}
	switch o.k {
	case kBig:
		c.cmpLin++
		c.cmpBin++
		lin := 1 + (len(it.key)-depth+4)/64
		c.linesLin += lines + lin
		c.linesBin += lines + lin
	case kPage:
		ll, lb, cl, cb := m.pageLookup(o.p, it)
		c.cmpLin += cl
		c.cmpBin += cb
		c.linesLin += lines + ll
		c.linesBin += lines + lb
	default:
		panic("lookup of a key that is not there: " + string(it.key))
	}
	c.lookups++
}

// pageLookup returns the cache lines a lookup of it in page p touches and the remainders it compares, walking the page
// (lin) and by a binary search over the prefix sums of the key lengths (bin).
func (m *model) pageLookup(p *mpage, it *item) (linesLin, linesBin, cmpLin, cmpBin int) {
	need, _ := m.need(p.items, p.from)
	size := classFor(need)
	j := sort.Search(len(p.items), func(i int) bool { return bytes.Compare(p.items[i].key, it.key) >= 0 })
	after := 0 // value bytes of the items after j
	for _, x := range p.items[j+1:] {
		after += x.vb
	}
	vEnd := size - after
	vStart := vEnd - it.vb
	set := func(lo, hi int, s map[int]bool) {
		for l := lo / 64; l <= (max(hi, lo+1)-1)/64; l++ {
			s[l] = true
		}
	}
	if len(p.items) == 1 {
		s := map[int]bool{}
		r := len(it.key) - p.from
		set(0, page.Header+r, s)
		if m.str {
			set(page.Header+r, page.Header+r+it.vals, s)
		}
		set(vStart, vEnd, s)
		return len(s), len(s), 1, 1
	}
	cpl := min(lcp(p.items[0].key[p.from:], p.items[len(p.items)-1].key[p.from:]), page.MaxKeyPart)
	n := 0
	slotOf := make([]int, len(p.items))
	remOff := make([]int, len(p.items)+1)
	for i, x := range p.items {
		slotOf[i] = n
		n += x.vals
	}
	kl := page.Header + cpl
	r0 := kl + n
	if m.str {
		r0 += n
	}
	remOff[0] = r0
	for i, x := range p.items {
		remOff[i+1] = remOff[i] + len(x.key) - p.from - cpl
	}
	vals := func(s map[int]bool) {
		if m.str {
			set(kl+n+slotOf[j], kl+2*n, s)
		}
		set(vStart, vEnd, s)
	}
	// walk: the key lengths up to j's slot, the remainders up to j's end
	sl := map[int]bool{0: true}
	set(kl, kl+slotOf[j]+1, sl)
	set(r0, remOff[j+1], sl)
	vals(sl)
	// binary search: all key lengths (the prefix sums), the probed remainders
	sb := map[int]bool{0: true}
	set(kl, kl+n, sb)
	lo, hi, cb := 0, len(p.items), 0
	for lo < hi {
		mid := (lo + hi) / 2
		cb++
		set(remOff[mid], remOff[mid+1], sb)
		c := bytes.Compare(p.items[mid].key, it.key)
		if c == 0 {
			break
		}
		if c < 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	vals(sb)
	return len(sl), len(sb), j + 1, cb
}

// --- the cases ---

func corpusItems(kind keys.Kind, n int, profile string, str bool) []*item {
	c := keys.Generate(kind, n, 0x5EED)
	var counts [][]uint64
	if c.Natural != nil {
		counts = c.Natural[:n]
	} else {
		vals, offs := keys.Values(n, 0xFA11)
		for i := range n {
			counts = append(counts, vals[offs[i]:offs[i+1]])
		}
	}
	items := make([]*item, n)
	for i := range n {
		vs := counts[i]
		if profile == "single-value" {
			vs = vs[:1]
		}
		it := &item{key: c.Keys.B[i], vals: len(vs)}
		for _, v := range vs {
			if str && c.Names != nil {
				it.vb += len(c.Names[v-1])
			} else if str {
				it.vb += len(strconv.FormatUint(v, 10))
			} else {
				it.vb += 8
			}
		}
		items[i] = it
	}
	return items
}

func run(w io.Writer, args []string) error {
	fs := flag.NewFlagSet("rangemodel", flag.ContinueOnError)
	kindsF := fs.String("keys", "street,dirs,links,url", "key kinds")
	valuesF := fs.String("values", "single-value,natural", "value profiles")
	sizesF := fs.String("sizes", "4096,16384,65536", "numbers of keys")
	limitsF := fs.String("limits", "512,384,256", "page sizes above which a page of a range splits (the proposal)")
	str := fs.Bool("str", false, "string values (the names of the corpus) instead of uint64")
	packF := fs.String("pack", "", "instead of the table of the trees: the byte nodes packed into objects of at most these bytes (0: as today), for today's pages and for ranges split above 512")
	nodeLines := fs.Bool("nodelines", false, "instead of the table of the trees: the byte nodes a lookup passes by class, and how many of them read the child slot from a second line (64- and 128-byte lines), today's tree")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *packF != "" {
		return runPack(w, *kindsF, *valuesF, *sizesF, *packF, *str)
	}
	if *nodeLines {
		return runNodeLines(w, *kindsF, *valuesF, *sizesF, *str)
	}
	var limits []int
	for _, s := range strings.Split(*limitsF, ",") {
		l, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		limits = append(limits, l)
	}
	if _, err := fmt.Fprintln(w, "| case | tree | nodes a lookup | compares walk / binary | lines walk / binary | pages | one-key | own objects | keys a page mean / median / p10 / p90 | objects a key | bytes a key |\n|---|---|--:|--:|--:|--:|--:|--:|---|--:|--:|"); err != nil {
		return err
	}
	for _, ks := range strings.Split(*kindsF, ",") {
		kind := keys.Kind(ks)
		if !keys.Available(kind) {
			if _, err := fmt.Fprintf(w, "| %s | not available (see keys/testdata/README.md) |\n", ks); err != nil {
				return err
			}
			continue
		}
		for _, profile := range strings.Split(*valuesF, ",") {
			for _, ns := range strings.Split(*sizesF, ",") {
				n, err := strconv.Atoi(ns)
				if err != nil {
					return err
				}
				items := corpusItems(kind, n, profile, *str)
				models := []*model{{limit: page.Largest, str: *str}}
				for _, l := range limits {
					models = append(models, &model{ranges: true, limit: l, str: *str})
				}
				for _, m := range models {
					for _, it := range items {
						m.insert(it)
					}
					var c counts
					c.keys = n
					m.census(m.root, &c)
					for _, it := range items {
						m.lookup(it, &c)
					}
					name := "today (burst)"
					if m.ranges {
						name = fmt.Sprintf("ranges, split above %d", m.limit)
					}
					slices.Sort(c.pageKeys)
					pk := c.pageKeys
					sum := 0
					for _, k := range pk {
						sum += k
					}
					q := func(f float64) int { return pk[min(len(pk)-1, int(f*float64(len(pk))))] }
					l := float64(c.lookups)
					if _, err := fmt.Fprintf(w, "| %s %s %d | %s | %.2f | %.2f / %.2f | %.2f / %.2f | %d | %d | %d | %.1f / %d / %d / %d | %.3f | %.1f |\n",
						kind, profile, n, name, float64(c.nodesPassed)/l, float64(c.cmpLin)/l, float64(c.cmpBin)/l,
						float64(c.linesLin)/l, float64(c.linesBin)/l, c.pages, c.onePages, c.bigs,
						float64(sum)/float64(len(pk)), q(0.5), q(0.1), q(0.9),
						float64(c.pages+c.nodes+c.bigs)/float64(n), float64(c.nodeBytes+c.pageBytes+c.bigBytes)/float64(n)); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func runPack(w io.Writer, kindsF, valuesF, sizesF, packF string, str bool) error {
	var limits []int
	for _, s := range strings.Split(packF, ",") {
		l, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		limits = append(limits, l)
	}
	if _, err := fmt.Fprintln(w, "| case | packing | objects a lookup | lines a lookup | routing objects | routing B a key | all B a key | objects a key |\n|---|---|--:|--:|--:|--:|--:|--:|"); err != nil {
		return err
	}
	for _, ks := range strings.Split(kindsF, ",") {
		kind := keys.Kind(ks)
		if !keys.Available(kind) {
			continue
		}
		for _, profile := range strings.Split(valuesF, ",") {
			for _, ns := range strings.Split(sizesF, ",") {
				n, err := strconv.Atoi(ns)
				if err != nil {
					return err
				}
				items := corpusItems(kind, n, profile, str)
				for _, m := range []*model{{limit: page.Largest, str: str}, {ranges: true, limit: page.Largest, str: str}} {
					for _, it := range items {
						m.insert(it)
					}
					tree := "today's pages"
					if m.ranges {
						tree = "range pages"
					}
					if err := m.packReport(w, fmt.Sprintf("%s %s %d, %s", kind, profile, n, tree), items, limits); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// runNodeLines writes, for today's tree of every case, the byte nodes a lookup passes by class, those that read their
// child slot from a second line of the node (64- and 128-byte lines), and the prefix tails read: the second dependent
// lines inside the nodes that descent-analysis.md found to be 16 to 19 % of the descent's time on the PC.
func runNodeLines(w io.Writer, kindsF, valuesF, sizesF string, str bool) error {
	if _, err := fmt.Fprintln(w, "| case | objects a lookup | nodes N5 / N12 / N26 / N58 / N256 a lookup | second lines a lookup, 64 B: N12 / N26 / N58 / N256 = all | 128 B: all | prefix tails a lookup |\n|---|--:|---|---|--:|--:|"); err != nil {
		return err
	}
	for _, ks := range strings.Split(kindsF, ",") {
		kind := keys.Kind(ks)
		if !keys.Available(kind) {
			continue
		}
		for _, profile := range strings.Split(valuesF, ",") {
			for _, ns := range strings.Split(sizesF, ",") {
				n, err := strconv.Atoi(ns)
				if err != nil {
					return err
				}
				items := corpusItems(kind, n, profile, str)
				m := &model{limit: page.Largest, str: str}
				for _, it := range items {
					m.insert(it)
				}
				var c counts
				for _, it := range items {
					m.lookup(it, &c)
				}
				l := float64(c.lookups)
				nodes, s64, s128 := 0, 0, 0
				for k := range c.visits {
					nodes += c.visits[k]
					s64 += c.second64[k]
					s128 += c.second128[k]
				}
				if _, err := fmt.Fprintf(w, "| %s %s %d | %.2f | %.2f / %.2f / %.2f / %.2f / %.2f | %.2f / %.2f / %.2f / %.2f = %.2f | %.2f | %.2f |\n",
					kind, profile, n, float64(nodes)/l+1,
					float64(c.visits[0])/l, float64(c.visits[1])/l, float64(c.visits[2])/l, float64(c.visits[3])/l, float64(c.visits[4])/l,
					float64(c.second64[1])/l, float64(c.second64[2])/l, float64(c.second64[3])/l, float64(c.second64[4])/l, float64(s64)/l,
					float64(s128)/l, float64(c.tails)/l); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "rangemodel:", err)
		os.Exit(2)
	}
}
