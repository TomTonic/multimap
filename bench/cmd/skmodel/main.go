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
	pageGrid  = []int{32, 64, 128, 256, 384, 512}              // the classes of the single-key page
	flatGrid  = []int{32, 48, 64, 96, 128, 192, 256, 384, 512} // the classes of the flat leaf of today
	goClasses = []int{16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512}
)

var overflowF = flag.Bool("overflow", false, "also describe the keys whose content does not fit a page (the value overflow)")
var nodesF = flag.Bool("nodes", false, "also build the model of the byte nodes and print the whole tree")
var header = flag.Int("header", 3, "bytes of the header of a page")
var sets = flag.Bool("sets", false, "also compare sets of size classes")
var valuesF = flag.String("values", "string", "the values: string (bytes with a length byte), words (fixed 8 bytes, no length byte), words-len (8 bytes with a length byte), pointers (fixed 8 bytes, the remainder padded to whole words)")
var ovBytesF = flag.Float64("ovbytes", 0, "bytes a value takes in the value set of the value overflow (0: the estimate for strings, 24 up to 64 values and 40 beyond; measured with cmd/ovbench: Set3 of uint64 16.3, of pointers 18)")
var histF = flag.String("histogram", "", "write the entries of each data set, as lines `remainder values count`, to files entries-<data>.txt in this directory (the data of the microbenchmark of step 3.5: internal/art/testdata)")
var detail = flag.Bool("detail", false, "also print the distribution of the page contents and examples")

func main() {
	flag.Parse()
	headerBytes = *header
	w := os.Stdout
	fmt.Fprintln(w, "| data | values | keys | rem. B | values a key | value B | content B | page B (grid 128..512) | page B (Go classes) | keys over 512 B | keys up to 64 B content | page B (grid 32..512) |")
	fmt.Fprintln(w, "|---|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|")
	for _, kind := range []keys.Kind{keys.Street, keys.Dirs} {
		for _, singleValue := range []bool{false, true} {
			row(kind, keys.Capacity(kind), singleValue)
		}
	}
}

type entry struct {
	key  []byte
	vals []uint64
}

func es2(es []entry) []entry { return es }

func row(kind keys.Kind, n int, singleValue bool) {
	c := keys.Generate(kind, n, 0x5EED)
	if c.Natural == nil {
		fmt.Fprintf(os.Stderr, "skmodel: %s has no natural values\n", kind)
		return
	}
	n = len(c.Keys.B)
	es := make([]entry, n)
	for i, k := range c.Keys.B {
		vs := c.Natural[i]
		if singleValue {
			vs = vs[:1]
		}
		es[i] = entry{k, vs}
	}
	sort.Slice(es, func(i, j int) bool { return bytes.Compare(es[i].key, es[j].key) < 0 })
	var rem, nv, vb, content, gridB, goB, over, small, pageB float64
	longRem, multi := 0, 0
	hist := map[[2]int]int{}
	var multiGrid, multiFlat float64
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
		if *valuesF == "pointers" {
			size = (size + 7) &^ 7 // the values start at a word
		}
		vbytes := 0
		for _, v := range e.vals {
			s, lenByte := valueLen(c, v), 1
			switch *valuesF {
			case "words", "pointers":
				s, lenByte = 8, 0
			case "words-len":
				s = 8
			}
			size += lenByte + s
			vbytes += s
		}
		sizes[i], rems[i] = size, r
		hist[[2]int{r, len(e.vals)}]++
		if r > 58 {
			longRem++
		}
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
		pageB += float64(classFor(pageGrid, size))
		if len(e.vals) > 1 {
			multi++
			multiGrid += float64(classFor(pageGrid, size))
			multiFlat += float64(classFor(flatGrid, size))
		}
		gridB += float64(classFor(grid, size))
		goB += float64(classFor(goClasses, size))
	}
	fmt.Fprintf(os.Stderr, "skmodel: %s: %d keys (%.3f %%) have a remainder above 58 bytes (the longest that a typed page of 8 words holds)\n", kind, longRem, 100*float64(longRem)/float64(n))
	fmt.Fprintf(os.Stderr, "skmodel: %s: %d keys (%.1f %%) have several values: a page of the grid 32..512 takes %.1f B on average, a flat leaf of today's classes %.1f B\n", kind, multi, 100*float64(multi)/float64(n), multiGrid/float64(max(multi, 1)), multiFlat/float64(max(multi, 1)))
	if *histF != "" && !singleValue {
		writeHistogram(*histF, kind, hist)
	}
	f := float64(n)
	name := "natural"
	if singleValue {
		name = "single-value"
	}
	if *nodesF {
		wholeTree(kind, singleValue, c, es)
	}
	if *overflowF && !singleValue {
		overflow(kind, c, es)
	}
	if *sets {
		classSets(kind, singleValue, sizes)
	}
	if *detail {
		histogram(kind, singleValue, sizes)
		examples(c, es2(es), sizes, rems)
	}
	fmt.Printf("| %s | %s | %d | %.1f | %.2f | %.1f | %.1f | %.1f | %.1f | %.2f %% | %.0f %% | %.1f |\n", kind, name, n, rem/f, nv/f, vb/f, content/f, gridB/f, goB/f, 100*over/f, 100*small/f, pageB/f)
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
func histogram(kind keys.Kind, singleValue bool, sizes []int) {
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
	name := "natural"
	if singleValue {
		name = "single-value"
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
func classSets(kind keys.Kind, singleValue bool, sizes []int) {
	name := "natural"
	if singleValue {
		name = "single-value"
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
func wholeTree(kind keys.Kind, singleValue bool, c keys.Corpus, es []entry) {
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
		if *valuesF == "pointers" {
			size = (size + 7) &^ 7
		}
		for _, v := range es[i].vals {
			switch *valuesF {
			case "words", "pointers":
				size += 8
			case "words-len":
				size += 9
			default:
				size += 1 + valueLen(c, v)
			}
		}
		contentB += min(size, 512)
		if size > 512 { // a value overflow of 64 bytes (key area of up to 50 bytes) and a value set
			n := len(es[i].vals)
			ovKeys++
			ovVals += n
			pageB += 64 - classFor(classes, 512) // the value overflow replaces the page
			switch {
			case *ovBytesF > 0:
				ovBytes += int(*ovBytesF * float64(n))
			case n <= 64:
				ovBytes += 24 * n // an array of string headers, grown by halves
			default:
				ovBytes += 40 * n // a hash set of string headers
			}
		}
		pageB += classFor(classes, min(size, 512))
	}
	t.build(0, len(ks), 0)
	n := float64(len(ks))
	name := "natural"
	if singleValue {
		name = "single-value"
	}
	fmt.Printf("\n%s, %s: the whole tree (byte nodes of internal/art and pages of 32, 64, 128, 256, 384, 512)\n\n| nodes | node B/key | page B/key | total B/key |\n|--:|--:|--:|--:|\n| %d (%.2f a key) | %.1f | %.1f | %.1f |\n\nvalue overflow (content over 512 B): %d keys (%.2f %%) with %d values (%.0f %% of all); their value sets, estimated at 24 B a value up to 64 values and 40 B beyond: %.1f B/key. Total with them: %.1f B/key.\n", kind, name, t.count, float64(t.count)/n, float64(t.bytes)/n, float64(pageB)/n, float64(t.bytes+pageB)/n, ovKeys, 100*float64(ovKeys)/n, ovVals, 100*float64(ovVals)/float64(allVals), float64(ovBytes)/n, float64(t.bytes+pageB+ovBytes)/n)
}

// overflow describes the keys whose content is above 512 bytes: how many values
// they have, how long the values are, and what share of all values and value
// bytes they hold.
func overflow(kind keys.Kind, c keys.Corpus, es []entry) {
	bounds := []int{16, 64, 256, 1024, 4096, 1 << 30}
	type bucket struct{ keys, vals, bytes, maxVals int }
	bs := make([]bucket, len(bounds))
	allVals, allBytes := 0, 0
	for _, e := range es {
		n, b := len(e.vals), 0
		for _, v := range e.vals {
			b += valueLen(c, v)
		}
		allVals += n
		allBytes += b
		if headerBytes+b+n > 512 { // the rest of the content is small
			i := sort.SearchInts(bounds, n)
			bs[i].keys++
			bs[i].vals += n
			bs[i].bytes += b
			bs[i].maxVals = max(bs[i].maxVals, n)
		}
	}
	fmt.Printf("\n%s, natural: the keys with more than 512 bytes of content\n\n| values of the key | keys | values | share of all values | value bytes | average value | most values |\n|---|--:|--:|--:|--:|--:|--:|\n", kind)
	lo := 0
	for i, b := range bs {
		if b.keys == 0 {
			continue
		}
		hi := fmt.Sprint(bounds[i])
		if i == len(bounds)-1 {
			hi = "more"
		}
		fmt.Printf("| %d to %s | %d | %d | %.1f %% | %d | %.1f B | %d |\n", lo+1, hi, b.keys, b.vals, 100*float64(b.vals)/float64(allVals), b.bytes, float64(b.bytes)/float64(b.vals), b.maxVals)
		lo = bounds[i]
	}
	fmt.Printf("\nall values: %d, value bytes %d (%.1f B a value)\n", allVals, allBytes, float64(allBytes)/float64(allVals))
}

// writeHistogram writes the entries of a data set, one line `remainder values count` for each
// pair that occurs, in order, to entries-<kind>.txt in dir.
func writeHistogram(dir string, kind keys.Kind, hist map[[2]int]int) {
	pairs := make([][2]int, 0, len(hist))
	for k := range hist {
		pairs = append(pairs, k)
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	var b strings.Builder
	fmt.Fprintf(&b, "# remainder length, values, number of keys: the entries of %s as a tree of byte nodes cuts them (bench/cmd/skmodel -histogram)\n", kind)
	for _, k := range pairs {
		fmt.Fprintf(&b, "%d %d %d\n", k[0], k[1], hist[k])
	}
	if err := os.WriteFile(fmt.Sprintf("%s/entries-%s.txt", dir, kind), []byte(b.String()), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "skmodel:", err)
		os.Exit(1)
	}
}
