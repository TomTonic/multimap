//go:build mkstats

package art

import (
	"fmt"
	"strings"
)

// Shape describes the routing part of the tree for the questions of step 6 (docs/redesign/PLAN.md): how many
// byte nodes of each type, how full, how many carry a tail for their common prefix, how many have only pages
// below them, and how many nodes a lookup passes on its way to a key. Diagnosis only (build tag mkstats).
func (m *Map[T]) Shape() string {
	type typeStats struct{ count, children, tails, endPages, onlyPages, bytes int }
	var st [64]typeStats
	n5kids := [6]int{}
	depthKeys := map[int]int{}
	keys, plens := 0, [3]int{} // common prefix 0, 1..12, 13+
	var walk func(n *header, depth int)
	walk = func(n *header, depth int) {
		if isPage(n.objType) {
			k := 1
			if isMultiKey(n.objType) {
				k = m.multiKeyObject(n).Keys
			}
			depthKeys[depth] += k
			keys += k
			return
		}
		t := n.objType & objTypeMask
		s := &st[t]
		s.count++
		s.bytes += m.object(n).Size
		pl := n.prefixLen()
		switch {
		case pl == 0:
			plens[0]++
		case pl <= 12:
			plens[1]++
		default:
			plens[2]++
		}
		if tailClass(pl) != tailNone {
			s.tails++
		}
		kids, pages := 0, 0
		if e := endPageOf(n); e != nil {
			s.endPages++
			depthKeys[depth+1]++
			keys++
		}
		eachByteNode(n, func(_ byte, c *header) {
			kids++
			if isPage(c.objType) {
				pages++
			}
			walk(c, depth+1)
		})
		s.children += kids
		if pages == kids {
			s.onlyPages++
		}
		if t == kN5 {
			n5kids[min(kids, 5)]++
		}
	}
	if m.t.root == nil {
		return "empty tree\n"
	}
	walk(m.t.root, 0)
	sel := m.selectivity()
	var b strings.Builder
	fmt.Fprintf(&b, "| node | count | bytes | children (mean) | with tail | with end page | only pages below |\n|---|--:|--:|--:|--:|--:|--:|\n")
	for _, t := range []objType{kN5, kN12, kN26, kN58, kN256} {
		s := st[t]
		if s.count == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %d | %d | %.1f | %d | %d | %d |\n", objTypeLabels[t], s.count, s.bytes, float64(s.children)/float64(s.count), s.tails, s.endPages, s.onlyPages)
	}
	fmt.Fprintf(&b, "\nN5 by number of byte children (0..5): %v; common prefix of the nodes: none %d, 1-12 bytes %d, 13+ bytes %d\n", n5kids, plens[0], plens[1], plens[2])
	sum, maxD := 0, 0
	for d, k := range depthKeys {
		sum += d * k
		maxD = max(maxD, d)
	}
	fmt.Fprintf(&b, "nodes on the way to a key: mean %.2f, by depth:", float64(sum)/float64(max(1, keys)))
	for d := 0; d <= maxD; d++ {
		fmt.Fprintf(&b, " %d: %.1f %%", d, 100*float64(depthKeys[d])/float64(max(1, keys)))
	}
	b.WriteString("\n")
	b.WriteString(sel)
	return b.String()
}

// selectivity describes how well the first bytes of the remainders tell the keys of a multi-key page apart, for the
// question of a list of first bytes searched with SWAR (docs/redesign/review-2026-10.md, analysis b): for a key
// of a page, the number of keys of its page that share its first byte (or its first two bytes) — the keys a
// lookup would still compare in full — averaged over all keys of multi-key pages; and the share of the keys that
// are alone with their first byte(s).
func (m *Map[T]) selectivity() string {
	var keys, remBytes int
	var share, alone [3]int // by prefix length 1, 2 (index 1, 2)
	var pageKeys []int
	var visit func(n *header)
	visit = func(n *header) {
		if isPage(n.objType) {
			if !isMultiKey(n.objType) {
				return
			}
			var rems [][]byte
			each := func(rem []byte, first bool) {
				if first {
					rems = append(rems, rem)
				}
			}
			if m.flat == 3 {
				asMKStr(n).Each(func(rem, _ []byte, first bool) bool { each(rem, first); return true })
			} else {
				asMKFix(n).Each(func(rem []byte, _ T, first bool) bool { each(rem, first); return true })
			}
			pageKeys = append(pageKeys, len(rems))
			for _, r := range rems {
				keys++
				remBytes += len(r)
				for w := 1; w <= 2; w++ {
					same := 0
					for _, o := range rems {
						if string(o[:min(w, len(o))]) == string(r[:min(w, len(r))]) {
							same++
						}
					}
					share[w] += same
					if same == 1 {
						alone[w]++
					}
				}
			}
			return
		}
		if e := endPageOf(n); e != nil {
			visit(singleKeyHdr(e))
		}
		eachByteNode(n, func(_ byte, c *header) { visit(c) })
	}
	visit(m.t.root)
	if keys == 0 {
		return "no multi-key pages\n"
	}
	k := float64(keys)
	return fmt.Sprintf("first bytes of the remainders in multi-key pages (%d keys in %d pages, %.1f keys a page, remainder %.1f bytes): keys sharing the first byte with a key %.2f (alone %.0f %%), the first two bytes %.2f (alone %.0f %%)\n",
		keys, len(pageKeys), k/float64(len(pageKeys)), float64(remBytes)/k, float64(share[1])/k, 100*float64(alone[1])/k, float64(share[2])/k, 100*float64(alone[2])/k)
}
