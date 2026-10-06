package art

import (
	"fmt"
	"testing"
)

// BenchmarkMapAddPresent measures the descent of Add and the call into the page, which find the
// value there already (nothing changes): what upsert, reach and Page.Add cost without a change
// of the tree, for keys that share multi-key pages and for keys that have a single-key page each.
func BenchmarkMapAddPresent(b *testing.B) {
	for _, tc := range []struct {
		name string
		n    int
		long bool
	}{{"pages", 4096, false}, {"single-key pages", 64, true}} {
		var m Map[uint64]
		keys := make([][]byte, tc.n)
		for i := range keys {
			keys[i] = fmt.Appendf(nil, "street-%05d", i*7%tc.n)
			if tc.long { // a value too long for a page makes a single-key page of each key
				keys[i] = fmt.Appendf(nil, "%02d-%0300d", i, i)
			}
			m.Add(keys[i], uint64(i))
		}
		b.Run(tc.name, func(b *testing.B) {
			for i := range b.N {
				m.Add(keys[i%tc.n], uint64(i%tc.n))
			}
		})
	}
}
