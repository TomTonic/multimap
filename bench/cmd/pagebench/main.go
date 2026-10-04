// Command pagebench compares the two candidates for the multi-key page of the
// redesign (docs/redesign, PLAN step 3) on the real data sets: the page of
// internal/vpage (a directory of tags, 16-byte slots, one 8-byte value per key)
// and the page of internal/lpage (a header of lengths, values of any length, or
// of one width for numbers), with the keys of the corpora street and dirs and
// the first value of each key (the single-value profile) as the value: the number of
// the locality or the file name, or its name as a string.
//
//	go run ./cmd/pagebench -keys street,dirs -validation 10
//
// A run of pages is built from the keys in their random order, as the tree
// builds them below a range node, and the timed part is only the page's own
// lookup: which page holds a key is worked out before. The point lookups of a
// pair of layouts are compared with rtcompare, interleaved; the table at the end
// says how much memory the pages take per key. It is a diagnosis of the layouts
// in isolation, not a claim of speed (the tree's own benchmarks, cmd/bench, are).
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/multimap/internal/lpage"
	"github.com/TomTonic/multimap/internal/vpage"
	"github.com/TomTonic/rtcompare"
)

func main() {
	kinds := flag.String("keys", "street,dirs", "key kinds with natural values")
	n := flag.Int("n", 0, "keys per kind (0: all the corpus allows)")
	flag.Parse() // also reads -repeats and -validation of package rtopt
	if err := run(os.Stdout, *kinds, *n); err != nil {
		fmt.Fprintln(os.Stderr, "pagebench:", err)
		os.Exit(1)
	}
}

var sink uint64

// data is the entries of a data set that fit every layout: the keys in their
// random order with their first value as a number and as a name.
type data struct {
	keys    [][]byte
	nums    []uint64
	names   [][]byte
	absent  [][]byte // keys that are not there
	skipped int      // entries too big for a page
}

func load(kind keys.Kind, n int) data {
	c := keys.Generate(kind, min(n, keys.Capacity(kind)), 0x5EED)
	var d data
	for i, k := range c.Keys.B {
		v := c.Natural[i][0]
		name := []byte(c.Names[v-1])
		if len(k) > 250 || len(k)+len(name) > 500 {
			d.skipped++
			continue
		}
		d.keys, d.nums, d.names = append(d.keys, k), append(d.nums, v), append(d.names, name)
	}
	for _, k := range c.Misses.B {
		if len(k) <= 250 {
			d.absent = append(d.absent, k)
		}
	}
	return d
}

// layout is one candidate: how to build it, and a lookup of the page that holds
// the key as the timed part sees it.
type layout struct {
	name  string
	pages int
	bytes int // object sizes of the pages
	get   func(i int, absent bool) uint64
}

func (d data) order() []int {
	rng := rtcompare.NewDPRNG(7)
	idx := make([]int, len(d.keys))
	for i := range idx {
		idx[i] = i
	}
	for i := len(idx) - 1; i > 0; i-- {
		j := int(rng.Uint64() % uint64(i+1))
		idx[i], idx[j] = idx[j], idx[i]
	}
	return idx
}

func vpageLayout(d data, order []int) layout {
	var run vpage.Run
	for i, k := range d.keys {
		if err := run.Insert(k, d.nums[i]); err != nil {
			panic(err)
		}
	}
	hitP, missP := make([]*vpage.Page, len(d.keys)), make([]*vpage.Page, len(d.absent))
	for i, k := range d.keys {
		hitP[i] = run.PageFor(k)
	}
	for i, k := range d.absent {
		missP[i] = run.PageFor(k)
	}
	l := layout{name: "directory of tags (vpage), numbers"}
	for _, p := range run.Pages() {
		l.pages++
		l.bytes += p.Size()
	}
	l.get = func(i int, absent bool) uint64 {
		if absent {
			j := i % len(d.absent)
			v, _ := missP[j].Get(d.absent[j])
			return v
		}
		j := order[i%len(order)]
		v, _ := hitP[j].Get(d.keys[j])
		return v
	}
	return l
}

// lpageLayout builds the length-header pages with the values vals; with copy the
// lookup turns the value into a string, as a map of strings has to give it out.
func lpageLayout(name string, d data, order []int, vals [][]byte, copyOut bool) layout {
	var run lpage.Run
	for i, k := range d.keys {
		if err := run.Insert(k, vals[i]); err != nil {
			panic(err)
		}
	}
	hitP, missP := make([]*lpage.Page, len(d.keys)), make([]*lpage.Page, len(d.absent))
	for i, k := range d.keys {
		hitP[i] = run.PageFor(k)
	}
	for i, k := range d.absent {
		missP[i] = run.PageFor(k)
	}
	l := layout{name: name}
	for _, p := range run.Pages() {
		l.pages++
		l.bytes += p.Size()
	}
	weigh := func(v []byte) uint64 {
		if copyOut {
			s := string(v)
			return uint64(len(s))
		}
		if len(v) >= 8 {
			return binary.LittleEndian.Uint64(v)
		}
		return uint64(len(v))
	}
	l.get = func(i int, absent bool) uint64 {
		if absent {
			j := i % len(d.absent)
			v, _ := missP[j].Get(d.absent[j])
			return weigh(v)
		}
		j := order[i%len(order)]
		v, _ := hitP[j].Get(d.keys[j])
		return weigh(v)
	}
	return l
}

func candidate(l layout, absent bool) rtcompare.Candidate {
	return rtcompare.Candidate{Name: l.name, Batch: func(n uint64) {
		var acc uint64
		for i := range int(n) {
			acc += l.get(i, absent)
		}
		sink += acc
	}}
}

// run writes the comparison for the key kinds named in kinds (comma separated),
// with n keys of each, or all the corpus holds if n is 0.
func run(w io.Writer, kinds string, n int) error {
	var sb strings.Builder // written to w at the end: a failed write is then reported once
	for _, name := range strings.Split(kinds, ",") {
		kind := keys.Kind(name)
		if kind != keys.Street && kind != keys.Dirs {
			return fmt.Errorf("%q has no natural values (street or dirs)", name)
		}
		count := n
		if count == 0 {
			count = keys.Capacity(kind)
		}
		d := load(kind, count)
		order := d.order()
		numbers := make([][]byte, len(d.nums))
		for i, v := range d.nums {
			numbers[i] = binary.LittleEndian.AppendUint64(nil, v)
		}
		a := vpageLayout(d, order)
		bn := lpageLayout("length header (lpage), numbers", d, order, numbers, false)
		bs := lpageLayout("length header (lpage), names", d, order, d.names, false)
		bc := lpageLayout("length header (lpage), names copied to strings", d, order, d.names, true)
		fmt.Fprintf(&sb, "## %s: %d keys (%d too big for a page, %d absent keys)\n\n", kind, len(d.keys), d.skipped, len(d.absent))
		for _, pr := range []struct {
			title  string
			x, y   layout
			absent bool
		}{
			{"point lookups of present keys: tags against lengths, numbers", a, bn, false},
			{"point lookups of absent keys: tags against lengths, numbers", a, bn, true},
			{"point lookups of present keys: numbers against names, both in length-header pages", bn, bs, false},
			{"point lookups of present keys: names as bytes against names copied to strings", bs, bc, false},
		} {
			ca, cb := candidate(pr.x, pr.absent), candidate(pr.y, pr.absent)
			rep, err := rtcompare.Compare(ca, cb, rtopt.Options(ca, cb, false, 1))
			if err != nil {
				return err
			}
			fmt.Fprintf(&sb, "### %s\nA = %s; B = %s\n\n%s\n\n", pr.title, pr.x.name, pr.y.name, rep)
		}
		fmt.Fprintf(&sb, "| layout | pages | keys per page | page bytes per key |\n|---|--:|--:|--:|\n")
		for _, l := range []layout{a, bn, bs} {
			fmt.Fprintf(&sb, "| %s | %d | %.1f | %.1f |\n", l.name, l.pages, float64(len(d.keys))/float64(l.pages), float64(l.bytes)/float64(len(d.keys)))
		}
		fmt.Fprintln(&sb)
	}
	fmt.Fprintf(&sb, "checksum %d\n", sink) // the batches' results stay observable
	_, err := io.WriteString(w, sb.String())
	return err
}
