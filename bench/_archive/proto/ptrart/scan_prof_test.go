package ptrart

import (
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
)

// BenchmarkScan100 exists for profiling only (go test -cpuprofile); the
// comparisons that decide anything are made with rtcompare in cmd/compare.
func BenchmarkScan100(b *testing.B) {
	for _, kind := range []keys.Kind{keys.U64, keys.Str} {
		b.Run(string(kind), func(b *testing.B) {
			c := keys.Generate(kind, 1<<20, 0x5EED)
			t := &Tree{}
			for i, k := range c.Keys.B {
				t.Put(k, uint32(i))
			}
			var acc uint64
			j := 0
			b.ResetTimer()
			for range b.N {
				cnt := 0
				t.Scan(c.Hits.B[j], func(_ []byte, v uint32) bool { acc += uint64(v); cnt++; return cnt < 100 })
				j = (j + 1) % len(c.Hits.B)
			}
			_ = acc
		})
	}
}
