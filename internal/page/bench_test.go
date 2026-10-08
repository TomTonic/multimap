package page

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

// pageSet returns 4,096 sets of n sorted keys, 5 to 12 bytes of lower-case letters after a common
// prefix of 3 (what the remainders of street names look like), and a value of 10 bytes each: more pages
// than the L1 cache holds, so that a lookup does not find its page in a loop of the predictor, and fewer
// than the L2 cache.
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

var sink uint64

// BenchmarkPage measures what the tree asks of a many-key page of strings: finding a key that is
// there, one that is not, adding and removing an entry in a page that has room, and building a page
// from its entries, for pages of 3, 7 and 20 entries (the model of street has 6.5 on average, up to 29).
func BenchmarkPage(b *testing.B) {
	for _, n := range []int{3, 7, 20} {
		keys, vals := pageSet(n)
		pages := make([]*Str, len(keys))
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
				q, res := p.Add(k, vals[0][0])
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
				q, res := p.Add(k, uint64(99), false)
				if res == Added {
					p, _ = q.Remove(k, uint64(99), false)
					pages[i&4095] = p
				}
			}
		})
	}
}

// BenchmarkFixedPointers is BenchmarkFixed for values that are pointers: a page of pointers changes in
// place while the byte area and the free values hold (docs/redesign/step5-one-page.md).
func BenchmarkFixedPointers(b *testing.B) {
	for _, n := range []int{3, 7, 20} {
		keys, _ := pageSet(n)
		pages := make([]*Fixed, len(keys))
		for i := range pages {
			vs := make([]*uint64, len(keys[i]))
			for j := range vs {
				vs[j] = &ptrPool[j]
			}
			pages[i] = BuildFixed(keys[i], vs)
		}
		b.Run(fmt.Sprintf("n=%d/get hit", n), func(b *testing.B) {
			for i := range b.N {
				j := i & 4095
				if _, ok := pages[j].Get[*uint64](keys[j][i%len(keys[j])]); !ok {
					b.Fatal("lost")
				}
			}
		})
		b.Run(fmt.Sprintf("n=%d/insert and remove", n), func(b *testing.B) {
			k := []byte("prem")
			for i := range b.N {
				p := pages[i&4095]
				q, res := p.Add(k, &ptrPool[63], true)
				if res == Added {
					p, _ = q.Remove(k, &ptrPool[63], true)
					pages[i&4095] = p
				}
			}
		})
	}
}

// BenchmarkFixedScan measures what a range scan asks of a page of one-word values per value: Each, which
// calls a function with every remainder and value (what the tree's scanPage does).
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
				values += int(p.currentValues)
			}
			b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(values), "ns/value")
			sink = sum
		})
	}
}

// BenchmarkPageScan is BenchmarkFixedScan for values of variable length.
func BenchmarkPageScan(b *testing.B) {
	for _, n := range []int{7, 20} {
		keys, vals := pageSet(n)
		pages := make([]*Str, len(keys))
		for i := range pages {
			pages[i] = BuildStrings(keys[i], vals[i])
		}
		b.Run(fmt.Sprintf("n=%d/Each", n), func(b *testing.B) {
			var sum uint64
			values := 0
			for i := range b.N {
				p := pages[i&255]
				p.Each(func(rem, v []byte, _ bool) bool { sum += uint64(len(v) + len(rem)); return true })
				values += int(p.currentValues)
			}
			b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(values), "ns/value")
			sink = sum
		})
	}
}

// BenchmarkOneKey measures the one-key form, which holds the values of one key (what the single-key
// page did): Get, adding a value and removing it again, for a key with 3 values, strings and `uint64`.
func BenchmarkOneKey(b *testing.B) {
	keys, vals := pageSet(1)
	one := make([]*Str, 4096)
	onef := make([]*Fixed, 4096)
	for i := range one {
		k := keys[i][0]
		one[i] = BuildStrings([][]byte{k, k, k}, [][]byte{vals[i][0], []byte("1111111111"), []byte("2222222222")})
		onef[i] = BuildFixed([][]byte{k, k, k}, []uint64{1, 2, 3})
	}
	b.Run("strings/get hit", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			if _, ok := one[j].Get(keys[j][0]); !ok {
				b.Fatal("lost")
			}
		}
	})
	b.Run("strings/add and remove a value", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			q, res := one[j].Add(keys[j][0], []byte("3333333333"))
			if res == AddedValue {
				one[j], _ = q.Remove(keys[j][0], []byte("3333333333"))
			}
		}
	})
	b.Run("uint64/get hit", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			if _, ok := onef[j].Get[uint64](keys[j][0]); !ok {
				b.Fatal("lost")
			}
		}
	})
	b.Run("uint64/add and remove a value", func(b *testing.B) {
		for i := range b.N {
			j := i & 4095
			q, res := onef[j].Add(keys[j][0], uint64(99), false)
			if res == AddedValue {
				onef[j], _ = q.Remove(keys[j][0], uint64(99), false)
			}
		}
	})
}
