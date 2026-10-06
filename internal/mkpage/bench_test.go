package mkpage

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"

	"github.com/TomTonic/multimap/internal/skpage"
)

// pageSet returns 4,096 sets of n sorted keys, 5 to 12 bytes of lower-case letters
// after a common prefix of 3 (what the remainders of street names look like), and a
// value of 10 bytes each: more pages than the L1 cache holds, so that a lookup does
// not find its page in a loop of the predictor, and fewer than the L2 cache.
func pageSet(n int) (keys [][][]byte, vals [][][]byte) {
	r := rand.New(rand.NewPCG(1, 2))
	for range 4096 {
		var ks [][]byte
		for len(ks) < n {
			k := []byte("pre")
			for range 5 + r.IntN(8) {
				k = append(k, 'a'+byte(r.IntN(26)))
			}
			ks = append(ks, k)
		}
		slices.SortFunc(ks, func(a, b []byte) int { return strings.Compare(string(a), string(b)) })
		ks = slices.CompactFunc(ks, func(a, b []byte) bool { return string(a) == string(b) })
		vs := make([][]byte, len(ks))
		for i := range vs {
			vs[i] = []byte("0123456789")
		}
		keys, vals = append(keys, ks), append(vals, vs)
	}
	return keys, vals
}

// BenchmarkPage measures what the tree asks of a multi-key page of strings: finding
// a key that is there, one that is not, adding and removing an entry in a page that
// has room, and building a page from its entries, for pages of 3, 7 and 20 entries
// (the model of street has 6.5 on average, up to 29). It also measures the single-key
// page of skpage for the one comparison it must answer for a key: Match.
func BenchmarkPage(b *testing.B) {
	for _, n := range []int{3, 7, 20} {
		keys, vals := pageSet(n)
		pages := make([]*Page, len(keys))
		for i := range pages {
			pages[i] = BuildStrings(keys[i], vals[i])
		}
		miss := []byte("prezzzzzzzz")
		b.Run(fmt.Sprintf("n=%d/get hit", n), func(b *testing.B) {
			for i := range b.N {
				j := i & 4095
				if _, ok := pages[j].Get(keys[j][i%len(keys[j])]); !ok {
					b.Fatal("lost")
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/get miss", n), func(b *testing.B) {
			for i := range b.N {
				if _, ok := pages[i&4095].Get(miss); ok {
					b.Fatal("found")
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/insert and remove", n), func(b *testing.B) {
			k := []byte("prem")
			for i := range b.N {
				p := pages[i&4095]
				q, res := p.Insert(k, vals[0][0])
				if res == Added {
					p, _ = q.Remove(k, vals[0][0])
					pages[i&4095] = p
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/build", n), func(b *testing.B) {
			for i := range b.N {
				if BuildStrings(keys[i&4095], vals[i&4095]) == nil {
					b.Fatal("nil")
				}
			}
		})
	}
	// the single-key page: the key is cut after the byte of its node, one entry
	sk := make([]*skpage.Page, 4096)
	keys, vals := pageSet(1)
	for i := range sk {
		sk[i] = skpage.New(keys[i][0], len(keys[i][0]), vals[i][0])
	}
	b.Run("single-key page/match", func(b *testing.B) {
		for i := range b.N {
			if !sk[i&4095].Match(keys[i&4095][0]) {
				b.Fatal("lost")
			}
		}
	})
}

// BenchmarkFixed is BenchmarkPage for values of one word (uint64).
func BenchmarkFixed(b *testing.B) {
	for _, n := range []int{3, 7, 20} {
		keys, _ := pageSet(n)
		pages := make([]*Fixed, len(keys))
		for i := range pages {
			vs := make([]uint64, len(keys[i]))
			for j := range vs {
				vs[j] = uint64(j)
			}
			pages[i] = BuildFixed(keys[i], vs)
		}
		miss := []byte("prezzzzzzzz")
		b.Run(fmt.Sprintf("n=%d/get hit", n), func(b *testing.B) {
			for i := range b.N {
				j := i & 4095
				if _, ok := pages[j].Get[uint64](keys[j][i%len(keys[j])]); !ok {
					b.Fatal("lost")
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/get miss", n), func(b *testing.B) {
			for i := range b.N {
				if _, ok := pages[i&4095].Get[uint64](miss); ok {
					b.Fatal("found")
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/insert and remove", n), func(b *testing.B) {
			k := []byte("prem")
			for i := range b.N {
				p := pages[i&4095]
				q, res := p.Insert(k, uint64(99))
				if res == Added {
					p, _ = q.Remove(k, uint64(99))
					pages[i&4095] = p
				}
			}
		})
	}
}

// BenchmarkFixedScan measures what a range scan asks of a page of one-word values per value:
// Each, which calls a function with every remainder and value (what the tree's scanPage does),
// and a loop over the same arrays without a call per entry, the floor a tuned scan can reach
// (docs/redesign/step5-mkmv-design.md).
func BenchmarkFixedScan(b *testing.B) {
	for _, n := range []int{7, 20} {
		keys, _ := pageSet(n)
		pages := make([]*Fixed, len(keys))
		for i := range pages {
			vs := make([]uint64, len(keys[i]))
			for j := range vs {
				vs[j] = uint64(j)
			}
			pages[i] = BuildFixed(keys[i], vs)
		}
		b.Run(fmt.Sprintf("n=%d/Each", n), func(b *testing.B) {
			var sum uint64
			values := 0
			for i := range b.N {
				p := pages[i&255]
				p.Each(func(rem []byte, v uint64, _ bool) bool { sum += v + uint64(len(rem)); return true })
				values += int(p.n)
			}
			b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(values), "ns/value")
			sink = sum
		})
		b.Run(fmt.Sprintf("n=%d/loop", n), func(b *testing.B) {
			var sum uint64
			values := 0
			for i := range b.N {
				p := pages[i&255]
				m := p.mem()
				cnt := int(p.n)
				vs := p.vs(8)
				lo := Header + p.cpl()
				for j := range cnt {
					sum += *valueAt[uint64](m, vs, j) + uint64(m[lo+j])
				}
				values += cnt
			}
			b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(values), "ns/value")
			sink = sum
		})
	}
}

var sink uint64
