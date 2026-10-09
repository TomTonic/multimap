package art

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/internal/page"
)

// mostPageKeys returns the most keys a multi-key page of the subtree n holds, and 0 if it has none.
func mostPageKeys(n *header) int {
	if n == nil {
		return 0
	}
	if isPage(n.objType) {
		if isMultiKey(n.objType) {
			return (*page.Fixed)(unsafe.Pointer(n)).Keys()
		}
		return 0
	}
	most := 0
	if e := endPageOf(n); e != nil {
		most = mostPageKeys(singleKeyHdr(e))
	}
	eachByteNode(n, func(_ byte, c *header) { most = max(most, mostPageKeys(c)) })
	return most
}

// TestMaxKeys shows that a map whose pages are limited to a number of keys keeps its contents and its limit.
//
// The autotune of the keys per page (docs/redesign/autotune-design.md) lets a map hold at most maxKeys keys in a
// page, from 1 (every key its own page) to the page's own limits; experiment exp-maxkeys measures the curve. Every
// way a page gains keys (an add, a pair of two keys, a key outside the common prefix, a burst, a merge) must keep
// the limit, and the map must behave as without it.
//
// Expected: for limits 1, 2 and 4, with values of uint64, pointers and strings, the map agrees with the reference
// through growth and deletion (againstReference, uniqueAgainstReference), its invariants hold, and no multi-key page
// holds more keys than the limit (none at all for 1).
func TestMaxKeys(t *testing.T) {
	str := func(v uint64) string { return fmt.Sprint("value ", v) }
	for _, k := range []int{1, 2, 4} {
		for name, keys := range keySets() {
			if name == "paths-over-64k" {
				continue // the long paths have no pages that the limit would change; they are slow
			}
			t.Run(fmt.Sprintf("%s/maxKeys=%d", name, k), func(t *testing.T) {
				check := func(n *header) {
					t.Helper()
					if got := mostPageKeys(n); got > k || (k == 1 && got != 0) {
						t.Fatalf("a multi-key page holds %d keys, the limit is %d", got, k)
					}
				}
				u := &Map[uint64]{maxKeys: k}
				againstReference(t, keys, u, id)
				check(u.t.root)
				v := &Map[uint64]{maxKeys: k}
				uniqueAgainstReference(t, keys, v, id)
				check(v.t.root)
				p := &Map[*rec]{maxKeys: k}
				againstReference(t, keys, p, ptrValue)
				check(p.t.root)
				s := &Map[string]{maxKeys: k}
				againstReference(t, keys, s, str)
				check(s.t.root)
				w := &Map[string]{maxKeys: k}
				uniqueAgainstReference(t, keys, w, str)
				check(w.t.root)
				var g Map[uint64] // grown only: the pages are fullest
				g.SetMaxKeys(k)
				for i, key := range keys {
					g.Add(key, uint64(i))
				}
				checkInvariants(t, &g.t)
				check(g.t.root)
			})
		}
	}
}
