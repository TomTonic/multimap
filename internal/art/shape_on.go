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
	return b.String()
}
