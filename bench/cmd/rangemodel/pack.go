package main

import (
	"fmt"
	"io"
	"sort"
)

// Packing several small byte nodes into one object (the question B of step6-design.md, section 8): an object of at
// most `limit` bytes takes a byte node and as many of its descendant byte nodes as fit, the heaviest subtree first.
// Inside the object a level refers to a packed child by one byte; children outside it (pages, objects of their own,
// nodes not packed) are 8-byte pointers in an array behind the levels. A level costs 2 bytes (count, prefix length),
// its prefix, one byte a child (its byte or range start) and one byte a packed child. A node of more than 26 children
// stays a node of its own, as today. This is a static view of the finished tree: whether the packing can be kept up
// while keys come and go is not modelled.

type level struct {
	comp *comp
	off  int // offset of the level in its object
	size int
	ext  map[int]int // child index (-1: the end page) -> slot of its pointer, for children outside the object
}

type comp struct {
	members []*mnode
	size    int // bytes, before rounding to a class
	ptrBase int
	nptr    int
	plain   bool // one node as today
}

// keysBelow counts the keys of the subtree o.
func keysBelow(o obj, memo map[*mnode]int) int {
	switch o.k {
	case kPage:
		return len(o.p.items)
	case kBig:
		return 1
	case kNode:
		if v, ok := memo[o.n]; ok {
			return v
		}
		k := keysBelow(o.n.end, memo)
		for _, c := range o.n.kids {
			k += keysBelow(c, memo)
		}
		memo[o.n] = k
		return k
	}
	return 0
}

func children(n *mnode) int {
	c := len(n.starts)
	if n.end.k != kNil {
		c++
	}
	return c
}

// pack returns the level of every byte node of the tree when objects take at most limit bytes (0: no packing).
func (m *model) pack(limit int) (map[*mnode]*level, []*comp) {
	levels := map[*mnode]*level{}
	var comps []*comp
	memo := map[*mnode]int{}
	var tops []*mnode
	if m.root.k == kNode {
		tops = append(tops, m.root.n)
	}
	for len(tops) > 0 {
		top := tops[len(tops)-1]
		tops = tops[:len(tops)-1]
		c := &comp{members: []*mnode{top}}
		comps = append(comps, c)
		in := map[*mnode]bool{top: true}
		if limit > 0 && children(top) <= 26 {
			// greedy: the heaviest packable node child of a member that still fits
			size := 4 + 2 + len(top.prefix) + children(top) + 8*children(top)
			for {
				var best *mnode
				bestK := -1
				for _, mb := range c.members {
					for _, k := range mb.kids {
						if k.k == kNode && !in[k.n] && children(k.n) <= 26 {
							add := 2 + len(k.n.prefix) + children(k.n) + 8*children(k.n) - 8 + 1
							if size+add <= limit {
								if kk := keysBelow(k, memo); kk > bestK {
									best, bestK = k.n, kk
								}
							}
						}
					}
				}
				if best == nil {
					break
				}
				size += 2 + len(best.prefix) + children(best) + 8*children(best) - 8 + 1
				c.members = append(c.members, best)
				in[best] = true
			}
		}
		c.plain = len(c.members) == 1
		// the layout: levels in packing order, then the pointers
		off := 4
		for _, mb := range c.members {
			lv := &level{comp: c, off: off, ext: map[int]int{}}
			internal := 0
			for _, k := range mb.kids {
				if k.k == kNode && in[k.n] {
					internal++
				}
			}
			lv.size = 2 + len(mb.prefix) + children(mb) + internal
			off += lv.size
			levels[mb] = lv
		}
		c.ptrBase = off
		for _, mb := range c.members {
			lv := levels[mb]
			if mb.end.k != kNil {
				lv.ext[-1] = c.nptr
				c.nptr++
			}
			for i, k := range mb.kids {
				if k.k == kNil {
					continue
				}
				if k.k == kNode && in[k.n] {
					continue
				}
				lv.ext[i] = c.nptr
				c.nptr++
				if k.k == kNode {
					tops = append(tops, k.n)
				}
			}
		}
		c.size = c.ptrBase + 8*c.nptr
	}
	return levels, comps
}

func compClass(c *comp) int {
	if c.plain {
		_, s := nodeCap(children(c.members[0]))
		return s
	}
	for _, s := range []int{128, 256, 512} {
		if c.size <= s {
			return s
		}
	}
	return ((c.size + 511) / 512) * 512
}

type packCounts struct {
	objects, lines, lookups int
	nodeBytes, comps        int
}

// lookupPacked counts the objects a lookup of it loads one after the other and the cache lines it touches.
func (m *model) lookupPacked(it *item, levels map[*mnode]*level, pc *packCounts) {
	o := m.root
	var cur *comp
	seen := map[int]bool{}
	flush := func() {
		pc.lines += len(seen)
		clear(seen)
	}
	mark := func(lo, hi int) {
		for l := lo / 64; l <= (max(hi, lo+1)-1)/64; l++ {
			seen[l] = true
		}
	}
	for o.k == kNode {
		n := o.n
		lv := levels[n]
		if lv.comp != cur {
			flush()
			cur = lv.comp
			pc.objects++
			mark(0, 4)
		}
		q := n.q()
		idx := -1
		if len(it.key) > q {
			b := int(it.key[q])
			if m.ranges {
				idx = floor(n.starts, b)
			} else {
				idx = sort.SearchInts(n.starts, b)
			}
		}
		if cur.plain {
			cp, _ := nodeCap(children(n))
			mark(0, 16)
			if idx >= 0 {
				o2 := slotOff(cp, idx, n.starts[idx])
				mark(o2, o2+8)
			} else {
				o2 := slotOff(cp, -1, 0)
				mark(o2, o2+8)
			}
		} else {
			mark(lv.off, lv.off+lv.size)
			if slot, ok := lv.ext[idx]; ok {
				p := cur.ptrBase + 8*slot
				mark(p, p+8)
			}
		}
		if idx < 0 {
			o = n.end
		} else {
			o = n.kids[idx]
		}
	}
	flush()
	pc.objects++ // the page or the object of its own
	switch o.k {
	case kPage:
		ll, _, _, _ := m.pageLookup(o.p, it)
		pc.lines += ll
	case kBig:
		pc.lines++
	}
	pc.lookups++
}

// packReport prints, for each limit, the objects and lines a lookup and the bytes of the routing objects a key.
func (m *model) packReport(w io.Writer, name string, items []*item, limits []int) error {
	for _, limit := range limits {
		levels, comps := m.pack(limit)
		var pc packCounts
		for _, c := range comps {
			pc.nodeBytes += compClass(c)
		}
		pc.comps = len(comps)
		for _, it := range items {
			m.lookupPacked(it, levels, &pc)
		}
		var c counts
		m.census(m.root, &c)
		label := "no packing"
		if limit > 0 {
			label = fmt.Sprintf("packed into %d B", limit)
		}
		l := float64(pc.lookups)
		multi := 0
		for _, c := range comps {
			if !c.plain {
				multi++
			}
		}
		if _, err := fmt.Fprintf(w, "| %s | %s | %.2f | %.2f | %d (%d packed) | %.1f | %.1f | %.3f |\n", name, label,
			float64(pc.objects)/l, float64(pc.lines)/l, pc.comps, multi, float64(pc.nodeBytes)/float64(len(items)),
			float64(pc.nodeBytes+c.pageBytes+c.bigBytes)/float64(len(items)),
			float64(pc.comps+c.pages+c.bigs)/float64(len(items))); err != nil {
			return err
		}
	}
	return nil
}
