package art

import (
	"strconv"
	"testing"
)

// TestMergeBoundNeverAboveNeed: a map whose keys are removed merges its nodes back into pages exactly as before the
// cheaper merge tries (candidates-2026-10.md, W1). In the tree's merge check, the bound from the page heads
// (mergeBound) decides the clear refusals before the exact walk over the entries (mergeNeed). For every byte node of
// maps of fixed values, string values and pointer values, of keys that end at nodes and of keys with many values, the
// bound is never above the exact need, and wherever the bound refuses, the exact walk refuses too.
func TestMergeBoundNeverAboveNeed(t *testing.T) {
	keys := func(n, mod int) [][]byte {
		var ks [][]byte
		for i := range n {
			k := "k/" + strconv.Itoa(i*7919%mod)
			ks = append(ks, []byte(k), []byte(k+"/x"), []byte(k[:len(k)-1]))
		}
		return ks
	}
	nodes := 0
	check := func(name string, root *header, bound func(*header) (int, bool), need func(*header, []byte) (int, bool)) {
		var walk func(n *header)
		walk = func(n *header) {
			if n == nil || isPage(n.objType) {
				return
			}
			nodes++
			var buf [prefixBuf]byte
			lb, okb := bound(n)
			nd, okn := need(n, appendPrefix(buf[:0], n))
			if okn && (!okb || lb > nd) {
				t.Fatalf("%s: node of %d children: bound %d (%v) against need %d", name, n.count, lb, okb, nd)
			}
			if e := endPageOf(n); e != nil {
				walk(singleKeyHdr(e))
			}
			eachByteNode(n, func(_ byte, c *header) { walk(c) })
		}
		walk(root)
	}
	for _, mod := range []int{300, 5000, 100000} {
		var u Map[uint64]
		var s Map[string]
		var p Map[*rec]
		for i, k := range keys(2000, mod) {
			for v := range uint64(1 + i%3*i%40) {
				u.Add(k, v)
				s.Add(k, strconv.FormatUint(v, 10))
				p.Add(k, &rec{id: v})
			}
		}
		for i, k := range keys(2000, mod) {
			if i%3 == 0 {
				u.RemoveKey(k)
				s.RemoveKey(k)
				p.RemoveKey(k)
			}
		}
		check("uint64", u.t.root, u.mergeBound, u.mergeNeed)
		check("string", s.t.root, s.mergeBound, s.mergeNeed)
		check("pointer", p.t.root, p.mergeBound, p.mergeNeed)
	}
	var wide Map[uint64] // a 256-way node: no bound
	for b := range 256 {
		wide.Add([]byte{'w', byte(b), 'x'}, uint64(b))
		wide.Add([]byte{'w', byte(b), 'y'}, uint64(b))
	}
	check("256-way", wide.t.root, wide.mergeBound, wide.mergeNeed)
	if nodes < 100 {
		t.Fatalf("only %d nodes checked", nodes)
	}
}
