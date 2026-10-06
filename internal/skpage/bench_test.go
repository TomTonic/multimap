package skpage

import (
	"math/rand/v2"
	"testing"
)

// BenchmarkOneKey measures the single-key page for what internal/page's BenchmarkOneKey measures of its one-key
// form (docs/redesign/step5-one-page.md): the lookup of the key and of its first value, and adding a value and
// removing it again, for a key of 8 to 15 bytes with 3 values, strings and `uint64`, in 4,096 pages.
func BenchmarkOneKey(b *testing.B) {
	r := rand.New(rand.NewPCG(1, 2))
	keys := make([][]byte, 4096)
	one := make([]*Page, 4096)
	onef := make([]*Fixed, 4096)
	for i := range keys {
		k := []byte("pre")
		for range 5 + r.IntN(8) {
			k = append(k, 'a'+byte(r.IntN(26)))
		}
		keys[i] = k
		one[i] = Build(k, len(k), [][]byte{[]byte("0123456789"), []byte("1111111111"), []byte("2222222222")})
		onef[i] = BuildFixed(k, len(k), []uint64{1, 2, 3})
	}
	first := []byte("0123456789")
	b.Run("strings/get hit", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			if !one[j].Match(keys[j]) || !one[j].Has(first) {
				b.Fatal("lost")
			}
		}
	})
	b.Run("strings/add and remove a value", func(b *testing.B) {
		v := []byte("3333333333")
		for i := range b.N {
			j := i & 4095
			q, res := one[j].Add(v)
			if res == Added {
				one[j], _ = q.Remove(v)
			}
		}
	})
	b.Run("uint64/get hit", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			if !onef[j].Match(keys[j]) || !onef[j].Has(uint64(1)) {
				b.Fatal("lost")
			}
		}
	})
	b.Run("uint64/add and remove a value", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			q, res := onef[j].Add(uint64(99))
			if res == Added {
				onef[j], _ = q.Remove(uint64(99))
			}
		}
	})
}
