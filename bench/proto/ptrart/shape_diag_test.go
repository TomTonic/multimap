package ptrart

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
)

// children lists n's children in key order, hiding the four node layouts.
func children(n *header) []*header {
	switch n.kind {
	case kN4:
		x := asN4(n)
		return x.child[:x.count]
	case kN11:
		x := asN11(n)
		return x.child[:x.count]
	case kN57:
		x := asN57(n)
		return x.child[:x.count]
	case kN256:
		var out []*header
		for _, c := range asN256(n).child {
			if c != nil {
				out = append(out, c)
			}
		}
		return out
	}
	return nil
}

// anyKey returns some key below n, to tell which key segment n branches in.
func anyKey(n *header) []byte {
	for n.kind != kLeaf {
		if n.term != nil {
			return n.term.key()
		}
		n = children(n)[0]
	}
	return asLeaf(n).key()
}

// TestShapeDiag explains why lookups and scans cost more for some key
// distributions than for others: it reports how deep the tree gets and how
// many children its nodes have. It belongs to the benchmark prototypes and
// measures the pointer ART on the 1M-key u64 and str corpora, printing node
// counts and average fanout per tree level and, for str keys, per key segment
// (tenant/category/item/number). It asserts nothing and runs only when
// SHAPE_DIAG is set, because building both trees takes several seconds.
func TestShapeDiag(t *testing.T) {
	if os.Getenv("SHAPE_DIAG") == "" {
		t.Skip("set SHAPE_DIAG=1 to print the tree shape report")
	}
	for _, kind := range []keys.Kind{keys.U64, keys.Str} {
		c := keys.Generate(kind, 1<<20, 0x5EED)
		tr := &Tree{}
		for i, k := range c.Keys.B {
			tr.Put(k, uint32(i))
		}
		inner, leaves, depthSum := 0, 0, 0
		byLevel := map[int][2]int{}   // node level -> {nodes, children}
		bySegment := map[int][2]int{} // key segment -> {nodes, children}
		var walk func(n *header, byteDepth, level int)
		walk = func(n *header, byteDepth, level int) {
			if n.kind == kLeaf {
				leaves++
				depthSum += level
				return
			}
			inner++
			pos := byteDepth + int(n.plen)
			ch := children(n)
			fan := len(ch)
			v := byLevel[level]
			byLevel[level] = [2]int{v[0] + 1, v[1] + fan}
			seg := 0
			if kind == keys.Str {
				seg = bytes.Count(anyKey(n)[:pos], []byte("/"))
			}
			s := bySegment[seg]
			bySegment[seg] = [2]int{s[0] + 1, s[1] + fan}
			for _, x := range ch {
				walk(x, pos+1, level+1)
			}
		}
		walk(tr.root, 0, 0)
		fmt.Printf("%s: inner=%d (%.2f/key) leaves=%d avg leaf depth=%.2f\n",
			kind, inner, float64(inner)/float64(leaves), leaves, float64(depthSum)/float64(leaves))
		for l := 0; l < 12; l++ {
			if v, ok := byLevel[l]; ok {
				fmt.Printf("  level %d: %7d nodes, avg fanout %.1f\n", l, v[0], float64(v[1])/float64(v[0]))
			}
		}
		if kind == keys.Str {
			for s := 0; s < 4; s++ {
				v := bySegment[s]
				fmt.Printf("  segment %d: %7d nodes, avg fanout %.1f\n", s, v[0], float64(v[1])/float64(v[0]))
			}
		}
	}
}
