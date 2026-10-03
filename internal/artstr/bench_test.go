package artstr

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/TomTonic/multimap/internal/art"
)

// benchKeys returns n distinct text keys of the shape of street names and
// their values, names of about ten bytes.
func benchKeys(n int) (keys [][]byte, vals []string) {
	r := rand.New(rand.NewPCG(1, 2))
	w := []string{"Haupt", "Berg", "Linden", "Kirch", "Schul", "Bahnhof", "Garten", "Wald", "Wiesen", "Feld", "Markt", "Mühlen"}
	e := []string{"strasse", "weg", "platz", "allee", "gasse", "ring", "damm", "pfad"}
	seen := map[string]bool{}
	for len(keys) < n {
		k := fmt.Sprintf("%s%s %d", w[r.IntN(len(w))], e[r.IntN(len(e))], r.IntN(1<<20))
		if !seen[k] {
			seen[k] = true
			keys = append(keys, []byte(k))
			vals = append(vals, fmt.Sprintf("Ort%d-%s", r.IntN(1000), w[r.IntN(len(w))]))
		}
	}
	return keys, vals
}

// BenchmarkTree compares the tree with pages for strings (artstr) with the
// tree of today (art) in the operations of the benchmark driver: build, point
// lookups, a range scan and churn.
func BenchmarkTree(b *testing.B) {
	const n = 50000
	keys, vals := benchKeys(n)
	b.Run("build/art", func(b *testing.B) {
		for range b.N {
			var m art.Map[string]
			for i, k := range keys {
				m.Add(k, vals[i])
			}
		}
	})
	b.Run("build/artstr", func(b *testing.B) {
		for range b.N {
			var m Map[string]
			for i, k := range keys {
				m.Add(k, vals[i])
			}
		}
	})
	var a art.Map[string]
	var s Map[string]
	for i, k := range keys {
		a.Add(k, vals[i])
		s.Add(k, vals[i])
	}
	b.Run("get/art", func(b *testing.B) {
		var sink int
		for i := range b.N {
			a.Each(keys[i%n], func(v string) bool { sink += len(v); return true })
		}
		_ = sink
	})
	b.Run("get/artstr", func(b *testing.B) {
		var sink int
		for i := range b.N {
			s.Each(keys[i%n], func(v string) bool { sink += len(v); return true })
		}
		_ = sink
	})
	bd := art.Bounds{HasFrom: true, FromIncl: true}
	bs := Bounds{HasFrom: true, FromIncl: true}
	b.Run("range/art", func(b *testing.B) {
		var sink int
		for i := range b.N {
			bd.From = keys[i%n]
			c := 0
			a.RangeValues(&bd, func(v string) bool { sink += len(v); c++; return c < 100 })
		}
		_ = sink
	})
	b.Run("range/artstr", func(b *testing.B) {
		var sink int
		for i := range b.N {
			bs.From = keys[i%n]
			c := 0
			s.RangeValues(&bs, func(v string) bool { sink += len(v); c++; return c < 100 })
		}
		_ = sink
	})
	b.Run("churn/art", func(b *testing.B) {
		for i := range b.N {
			j := i % n
			a.Remove(keys[j], vals[j])
			a.Add(keys[j], vals[j])
		}
	})
	b.Run("churn/artstr", func(b *testing.B) {
		for i := range b.N {
			j := i % n
			s.Remove(keys[j], vals[j])
			s.Add(keys[j], vals[j])
		}
	})
}

// BenchmarkPairs compares the modes of the page tree on keys with several values:
// build and churn with one value per key and with up to four.
func BenchmarkPairs(b *testing.B) {
	const n = 50000
	keys, vals := benchKeys(n)
	for _, mode := range []struct {
		name string
		m    func() *Map[string]
	}{
		{"single", func() *Map[string] { return &Map[string]{} }},
		{"pairs", func() *Map[string] { return &Map[string]{Pairs: true} }},
		{"pairs-zc", func() *Map[string] { return &Map[string]{Pairs: true, ZeroCopy: true} }},
	} {
		fill := func(m *Map[string]) {
			for i, k := range keys {
				m.Add(k, vals[i])
				if i%5 == 0 {
					m.Add(k, vals[(i+1)%n]+"x")
				}
			}
		}
		b.Run("build/"+mode.name, func(b *testing.B) {
			for range b.N {
				fill(mode.m())
			}
		})
		m := mode.m()
		fill(m)
		b.Run("churn/"+mode.name, func(b *testing.B) {
			for i := range b.N {
				j := i % n
				m.Remove(keys[j], vals[j])
				m.Add(keys[j], vals[j])
			}
		})
	}
}
