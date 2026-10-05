package art

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
)

// The benchmarks of step 3.5.1 (docs/redesign/step3-fixed-design.md): the single-key page
// for fixed-size values, skpage.Fixed (with pointers: *rec), standalone and not as
// part of the tree. It was measured against the flat leaf and the typed leaf that
// it replaced; those results are in bench/results-layout/step3-fixed and the code of
// the two leaves is that of commit a689e03. Kept as the baseline for the profile of the page. The
// entries are those of the real data sets street and dirs as a tree of byte
// nodes cuts them (testdata, made by bench/cmd/skmodel -histogram): the
// length of the key's remainder and the number of its values. Only the entries
// that both candidates hold in an object of their own are measured; the keys that
// overflow are the value overflow's.

// rec is the record that pointer values point to, an object of its own.
type rec struct{ id, aux uint64 }

// entry is one key of a data set: the length of its remainder and its number of values.
type entry struct{ rem, n int }

// loadEntries returns the entries of the data set name in a fixed random order.
func loadEntries(tb testing.TB, name string) []entry {
	tb.Helper()
	f, err := os.Open("testdata/entries-" + name + ".txt")
	if err != nil {
		tb.Fatal(err)
	}
	defer func() { _ = f.Close() }()
	var out []entry
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "#") {
			continue
		}
		var rem, n, count int
		if _, err := fmt.Sscan(sc.Text(), &rem, &n, &count); err != nil {
			tb.Fatal(err)
		}
		for range count {
			out = append(out, entry{rem, n})
		}
	}
	rand.New(rand.NewSource(7)).Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
	return out
}

// pages is what a candidate does with the object of one key, which it hands
// around as an unsafe.Pointer.
type pages[T comparable] struct {
	name   string
	make   func(rem int, v T) unsafe.Pointer
	add    func(p unsafe.Pointer, v T) unsafe.Pointer // the object after adding v (the same if it had room or v was there)
	remove func(p unsafe.Pointer, v T) unsafe.Pointer // the object after removing v, which is there and not the only value
	values func(p unsafe.Pointer) []T
	size   func(p unsafe.Pointer, rem int) int // the bytes Go allocates for it
	fits   func(e entry) bool                  // whether the entry is held in an object of its own, with room for one more value
}

func fixedPages[T comparable]() pages[T] {
	key := make([]byte, 600)
	return pages[T]{
		name: "fixed page",
		make: func(rem int, v T) unsafe.Pointer { return unsafe.Pointer(skpage.NewFixed(key[:rem], rem, v)) },
		add: func(p unsafe.Pointer, v T) unsafe.Pointer {
			q, _ := (*skpage.Fixed)(p).Add(v)
			return unsafe.Pointer(q)
		},
		remove: func(p unsafe.Pointer, v T) unsafe.Pointer {
			q, _ := (*skpage.Fixed)(p).Remove(v)
			return unsafe.Pointer(q)
		},
		values: func(p unsafe.Pointer) []T { return (*skpage.Fixed)(p).Values[T]() },
		size:   func(p unsafe.Pointer, rem int) int { return (*skpage.Fixed)(p).Size() },
		fits: func(e entry) bool {
			var z T
			a := int(unsafe.Alignof(z))
			off := (skpage.Header + e.rem + a - 1) &^ (a - 1)
			return e.n+1 <= (512-off)/int(unsafe.Sizeof(z))
		},
	}
}

// fixedWorld is a data set built with one candidate.
type fixedWorld[T comparable] struct {
	entries []entry
	objects []unsafe.Pointer
	val     func(e, k int) T
	order   []int // the entries in random order
	picks   []int // random numbers for the choice of a value
}

// buildWorld builds the common entries of data with candidate c; common says
// which entries every candidate holds.
func buildWorld[T comparable](tb testing.TB, c pages[T], data []entry, common func(entry) bool, val func(e, k int) T) *fixedWorld[T] {
	w := &fixedWorld[T]{val: val}
	for _, e := range data {
		if common(e) {
			w.entries = append(w.entries, e)
		}
	}
	w.objects = make([]unsafe.Pointer, len(w.entries))
	for i, e := range w.entries {
		p := c.make(e.rem, val(i, 0))
		for k := 1; k < e.n; k++ {
			p = c.add(p, val(i, k))
		}
		if got := len(c.values(p)); got != e.n {
			tb.Fatalf("%s: entry %d holds %d values, want %d", c.name, i, got, e.n)
		}
		w.objects[i] = p
	}
	r := rand.New(rand.NewSource(11))
	w.order = r.Perm(len(w.entries))
	w.picks = make([]int, 1<<16)
	for i := range w.picks {
		w.picks[i] = r.Int()
	}
	return w
}

// both is the entries that all candidates of a kind of value hold.
func both(fs ...func(entry) bool) func(entry) bool {
	return func(e entry) bool {
		for _, f := range fs {
			if !f(e) {
				return false
			}
		}
		return true
	}
}

// BenchmarkFixedPages times the operations of a key's object on the real entries of two
// data sets, for the page with pointers (the page with values of 8 bytes without a
// pointer was measured against the flat leaf, see above): hit (adding a value that is
// there, which is a scan and the check of a set), churn (adding and removing a
// fresh value, which moves the object at the border of a class), read (summing the
// values) and build (all values of every entry, per entry), on the first 4,096
// entries (hot: they stay in the cache) and on all.
func BenchmarkFixedPages(b *testing.B) {
	recs := make([]*rec, 1<<16)
	for i := range recs {
		recs[i] = &rec{id: uint64(i), aux: 3}
	}
	for _, data := range []string{"street", "dirs"} {
		entries := loadEntries(b, data)
		b.Run("pointer/"+data, func(b *testing.B) {
			benchPages(b, entries, []pages[*rec]{fixedPages[*rec]()},
				func(e, k int) *rec { return recs[(e*131+k*17)%len(recs)] }, func(v *rec) uint64 { return v.id ^ v.aux })
		})
	}
}

func benchPages[T comparable](b *testing.B, data []entry, cands []pages[T], val func(e, k int) T, weigh func(T) uint64) {
	var fs []func(entry) bool
	for _, c := range cands {
		fs = append(fs, c.fits)
	}
	common := both(fs...)
	for _, c := range cands {
		w := buildWorld(b, c, data, common, val)
		for _, scope := range []struct {
			name string
			n    int
		}{{"hot", 4096}, {"all", len(w.entries)}} {
			n := min(scope.n, len(w.entries))
			b.Run(c.name+"/"+scope.name+"/hit", func(b *testing.B) {
				for i := range b.N {
					e := w.order[i%n]
					k := w.picks[i&(len(w.picks)-1)] % w.entries[e].n
					w.objects[e] = c.add(w.objects[e], w.val(e, k))
				}
			})
			b.Run(c.name+"/"+scope.name+"/churn", func(b *testing.B) {
				for i := range b.N {
					e := w.order[i%n]
					fresh := w.val(e, 1<<15) // a value the entry does not hold
					p := c.add(w.objects[e], fresh)
					w.objects[e] = c.remove(p, fresh)
				}
			})
			b.Run(c.name+"/"+scope.name+"/read", func(b *testing.B) {
				var sum uint64
				for i := range b.N {
					for _, v := range c.values(w.objects[w.order[i%n]]) {
						sum += weigh(v)
					}
				}
				sinkSum += sum
			})
		}
		b.Run(c.name+"/build", func(b *testing.B) {
			var keep []unsafe.Pointer
			for i := range b.N {
				e := i % len(w.entries)
				p := c.make(w.entries[e].rem, w.val(e, 0))
				for k := 1; k < w.entries[e].n; k++ {
					p = c.add(p, w.val(e, k))
				}
				if i < 1<<16 {
					keep = append(keep, p)
				}
			}
			_ = keep
		})
	}
}

var sinkSum uint64

// TestFixedMemory prints, for the entries both candidates hold, the bytes Go
// allocates for the objects of an entry: the page against the
// page, which the design note predicts (docs/redesign/step3-fixed-design.md). It
// checks only that the numbers are there; run it with -v.
func TestFixedMemory(t *testing.T) {
	recs := make([]*rec, 1<<16)
	for i := range recs {
		recs[i] = &rec{id: uint64(i), aux: 3}
	}
	for _, data := range []string{"street", "dirs"} {
		entries := loadEntries(t, data)
		for _, single := range []bool{false, true} {
			es := entries
			if single {
				es = make([]entry, len(entries))
				for i, e := range entries {
					es[i] = entry{e.rem, 1}
				}
			}
			ptrs := func(c pages[*rec]) (float64, int) {
				return meanSize(t, c, es, fixedPages[*rec]().fits, func(e, k int) *rec { return recs[(e*131+k*17)%len(recs)] })
			}
			profile := map[bool]string{false: "real", true: "single-value"}[single]
			f, na := ptrs(fixedPages[*rec]())
			t.Logf("%s %s, *rec: page %.1f B per entry (%d of %d entries held by a page)", data, profile, f, na, len(es))
		}
	}
}

// meanSize returns the mean bytes of the objects of candidate c over the
// entries of es that common selects, and their number.
func meanSize[T comparable](t *testing.T, c pages[T], es []entry, common func(entry) bool, val func(e, k int) T) (float64, int) {
	w := buildWorld(t, c, es, common, val)
	total := 0
	for i, p := range w.objects {
		total += c.size(p, w.entries[i].rem)
	}
	return float64(total) / float64(len(w.objects)), len(w.objects)
}
