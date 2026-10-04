// Command skmodel is the memory model of the single-key page (SKMV) for
// string values on the real data sets, for the design note of step 3.1
// (docs/redesign/step3-skmv-design.md). It builds no tree: it takes the keys of
// a corpus in key order, cuts each one where a tree of byte nodes (lazy
// expansion: the page hangs at the first byte that tells its key from its
// neighbours) would cut it, and adds up what a page of each key would hold:
// header, remainder, and the values with their length bytes. It prints, per
// data set and value profile, the bytes a key takes in pages of the size
// classes the plan names (128, 256, 384, 512) and in pages of the size classes
// of Go's allocator, to see what a page per key costs.
//
//	go run ./cmd/skmodel
//
// It takes the largest corpus of each data set (keys.Capacity).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
)

// headerBytes is the fixed part of the page of the design note: kind, number
// of values, length of the remainder (3 bytes; -header sets another).
var headerBytes = 3

// grid is the size classes of the plan; goClasses are Go's size classes up to 512.
var (
	grid      = []int{128, 256, 384, 512}
	goClasses = []int{16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512}
)

var nodesF = flag.Bool("nodes", false, "also build the model of the byte nodes and print the whole tree")
var header = flag.Int("header", 3, "bytes of the header of a page")
var sets = flag.Bool("sets", false, "also compare sets of size classes")
var detail = flag.Bool("detail", false, "also print the distribution of the page contents and examples")

func main() {
	flag.Parse()
	headerBytes = *header
	w := os.Stdout
	fmt.Fprintln(w, "| data | values | keys | rem. B | values a key | value B | content B | page B (grid 128..512) | page B (Go classes) | keys over 512 B | keys up to 64 B content |")
	fmt.Fprintln(w, "|---|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|")
	for _, kind := range []keys.Kind{keys.Street, keys.Dirs} {
		for _, unique := range []bool{false, true} {
			row(kind, keys.Capacity(kind), unique)
		}
	}
}

type entry struct {
	key  []byte
	vals []uint64
}

func es2(es []entry) []entry { return es }

func row(kind keys.Kind, n int, unique bool) {
	c := keys.Generate(kind, n, 0x5EED)
	if c.Natural == nil {
		fmt.Fprintf(os.Stderr, "skmodel: %s has no natural values\n", kind)
		return
	}
	n = len(c.Keys.B)
	es := make([]entry, n)
	for i, k := range c.Keys.B {
		vs := c.Natural[i]
		if unique {
			vs = vs[:1]
		}
		es[i] = entry{k, vs}
	}
	sort.Slice(es, func(i, j int) bool { return bytes.Compare(es[i].key, es[j].key) < 0 })
	var rem, nv, vb, content, gridB, goB, over, small float64
	sizes := make([]int, n)
	rems := make([]int, n)
	for i, e := range es {
		l := 0 // the longest common prefix with a neighbour
		if i > 0 {
			l = max(l, lcp(e.key, es[i-1].key))
		}
		if i+1 < n {
			l = max(l, lcp(e.key, es[i+1].key))
		}
		r := max(len(e.key)-l-1, 0) // the byte the page hangs at is consumed by its node
		size := headerBytes + r
		vbytes := 0
		for _, v := range e.vals {
			s := valueLen(c, v)
			size += 1 + s
			vbytes += s
		}
		sizes[i], rems[i] = size, r
		rem += float64(r)
		nv += float64(len(e.vals))
		vb += float64(vbytes)
		content += float64(size)
		if size > 512 {
			over++
			size = 512 // the rest is a value set; the page is full
		}
		if size <= 64 {
			small++
		}
		gridB += float64(classFor(grid, size))
		goB += float64(classFor(goClasses, size))
	}
	f := float64(n)
	name := "multi"
	if unique {
		name = "unique"
	}
	if *nodesF {
		wholeTree(kind, unique, c, es)
	}
	if *sets {
		classSets(kind, unique, sizes)
	}
	if *detail {
		histogram(kind, unique, sizes)
		examples(c, es2(es), sizes, rems)
	}
	fmt.Printf("| %s | %s | %d | %.1f | %.2f | %.1f | %.1f | %.1f | %.1f | %.2f %% | %.0f %% |\n", kind, name, n, rem/f, nv/f, vb/f, content/f, gridB/f, goB/f, 100*over/f, 100*small/f)
}

// valueLen returns the length of the string value number v: the name where the
// corpus has one, else 16 hex digits (the bench's string values).
func valueLen(c keys.Corpus, v uint64) int {
	if v >= 1 && v <= uint64(len(c.Names)) {
		return len(c.Names[v-1])
	}
	return 16
}

func lcp(a, b []byte) int {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i
}

// classFor returns the smallest class that holds size bytes (the largest if none).
func classFor(classes []int, size int) int {
	for _, c := range classes {
		if size <= c {
			return c
		}
	}
	return classes[len(classes)-1]
}

// histogram prints how the contents of the pages (header, remainder, values
// with their length bytes) spread over the size classes: the share of keys,
// the bytes of content, and the bytes a page per key takes in the grid
// 128..512 and in Go's classes.
func histogram(kind keys.Kind, unique bool, sizes []int) {
	bounds := []int{16, 32, 48, 64, 96, 128, 192, 256, 384, 512, 1 << 30}
	type bucket struct{ keys, content, grid, gocls int }
	bs := make([]bucket, len(bounds))
	for _, s := range sizes {
		i := sort.SearchInts(bounds, s)
		b := &bs[i]
		b.keys++
		b.content += s
		c := min(s, 512)
		b.grid += classFor(grid, c)
		b.gocls += classFor(goClasses, c)
	}
	name := "multi"
	if unique {
		name = "unique"
	}
	fmt.Printf("\n%s, %s: content of the page of a key\n\n| content bytes | keys | share | content B/key | grid 128..512 B/key | Go classes B/key |\n|---|--:|--:|--:|--:|--:|\n", kind, name)
	lo := 0
	for i, b := range bs {
		if b.keys == 0 {
			continue
		}
		hi := fmt.Sprint(bounds[i])
		if i == len(bounds)-1 {
			hi = "more"
		}
		k := float64(b.keys)
		fmt.Printf("| %d to %s | %d | %.1f %% | %.1f | %d | %.1f |\n", lo+1, hi, b.keys, 100*k/float64(len(sizes)), float64(b.content)/k, b.grid/b.keys, float64(b.gocls)/k)
		lo = bounds[i]
	}
}

// examples prints keys at quantiles of the content size, with their remainder
// and their values, as the page would hold them.
func examples(c keys.Corpus, es []entry, sizes, rems []int) {
	order := make([]int, len(sizes))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(a, b int) bool { return sizes[order[a]] < sizes[order[b]] })
	fmt.Printf("\nexamples (key; remainder bytes; values; content bytes):\n\n")
	for _, q := range []float64{0.05, 0.25, 0.5, 0.75, 0.9, 0.99} {
		i := order[int(q*float64(len(order)-1))]
		var vs []string
		for _, v := range es[i].vals {
			if v >= 1 && v <= uint64(len(c.Names)) {
				vs = append(vs, c.Names[v-1])
			} else {
				vs = append(vs, fmt.Sprintf("%016x", v))
			}
		}
		if len(vs) > 6 {
			vs = append(vs[:6], fmt.Sprintf("... (%d values)", len(es[i].vals)))
		}
		fmt.Printf("- p%.0f: %q; remainder %d B; %q; content %d B\n", 100*q, es[i].key, rems[i], vs, sizes[i])
	}
}

// classSet is a named set of size classes for the comparison of -sets.
type classSet struct {
	name    string
	classes []int
}

var classSets_ = []classSet{
	{"128, 256, 384, 512 (plan)", grid},
	{"64 + plan", []int{64, 128, 256, 384, 512}},
	{"32, 64 + plan", []int{32, 64, 128, 256, 384, 512}},
	{"32, 64, 96 + plan", []int{32, 64, 96, 128, 256, 384, 512}},
	{"32, 64, 96, 128, 192, 256, 384, 512", []int{32, 64, 96, 128, 192, 256, 384, 512}},
	{"16, 32, 48, 64 + plan", []int{16, 32, 48, 64, 128, 256, 384, 512}},
	{"Go's classes up to 512", goClasses},
}

// classSets prints, for each set of size classes, the bytes a page per key
// takes, how full the pages are, and the share of keys in each class.
func classSets(kind keys.Kind, unique bool, sizes []int) {
	name := "multi"
	if unique {
		name = "unique"
	}
	fmt.Printf("\n%s, %s: sets of size classes\n\n| size classes | page B/key | content B/key | fill | keys per class |\n|---|--:|--:|--:|---|\n", kind, name)
	for _, cs := range classSets_ {
		var total, content float64
		share := make([]int, len(cs.classes))
		for _, sz := range sizes {
			c := min(sz, 512)
			cl := classFor(cs.classes, c)
			total += float64(cl)
			content += float64(c)
			share[sort.SearchInts(cs.classes, cl)]++
		}
		f := float64(len(sizes))
		var parts []string
		for i, c := range cs.classes {
			if share[i] > 0 {
				parts = append(parts, fmt.Sprintf("%d: %.1f %%", c, 100*float64(share[i])/f))
			}
		}
		fmt.Printf("| %s | %.1f | %.1f | %.0f %% | %s |\n", cs.name, total/f, content/f, 100*content/total, strings.Join(parts, ", "))
	}
}

// wholeTree builds the trie of the model and prints the bytes a key takes in
// nodes and in pages (the classes 32, 64 and the plan's grid).
func wholeTree(kind keys.Kind, unique bool, c keys.Corpus, es []entry) {
	ks := make([][]byte, len(es))
	for i, e := range es {
		ks[i] = e.key
	}
	classes := []int{32, 64, 128, 256, 384, 512}
	var pageB, contentB, ovKeys, ovVals, ovBytes int
	allVals := 0
	for _, e := range es {
		allVals += len(e.vals)
	}
	t := &trie{keys: ks, nodes: map[int]int{}}
	t.pages = func(i, rem int) {
		size := headerBytes + rem
		for _, v := range es[i].vals {
			size += 1 + valueLen(c, v)
		}
		contentB += min(size, 512)
		if size > 512 { // a set leaf of 128 bytes and a value set of Go strings
			n := len(es[i].vals)
			ovKeys++
			ovVals += n
			pageB += 128 - classFor(classes, 512) // the set leaf replaces the page
			if n <= 64 {
				ovBytes += 24 * n // an array of string headers, grown by halves
			} else {
				ovBytes += 40 * n // a hash set of string headers
			}
		}
		pageB += classFor(classes, min(size, 512))
	}
	t.build(0, len(ks), 0)
	n := float64(len(ks))
	name := "multi"
	if unique {
		name = "unique"
	}
	fmt.Printf("\n%s, %s: the whole tree (byte nodes of internal/art and pages of 32, 64, 128, 256, 384, 512)\n\n| nodes | node B/key | page B/key | total B/key |\n|--:|--:|--:|--:|\n| %d (%.2f a key) | %.1f | %.1f | %.1f |\n\nvalue overflow (content over 512 B): %d keys (%.2f %%) with %d values (%.0f %% of all); their value sets, estimated at 24 B a value up to 64 values and 40 B beyond: %.1f B/key. Total with them: %.1f B/key.\n", kind, name, t.count, float64(t.count)/n, float64(t.bytes)/n, float64(pageB)/n, float64(t.bytes+pageB)/n, ovKeys, 100*float64(ovKeys)/n, ovVals, 100*float64(ovVals)/float64(allVals), float64(ovBytes)/n, float64(t.bytes+pageB+ovBytes)/n)
}
