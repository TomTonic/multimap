// Command objstat reports, for every benchmark case, how the objects of the
// ordered multimap fill cache lines: how many objects the index has, how many
// bytes it takes per key, how many objects are not a multiple of 64 or of 128
// bytes, and how many cross more 64-byte lines than their size needs.
//
// These are the rules of docs/redesign/STRATEGY.md, and the figures are the
// acceptance criteria of the redesign steps (PLAN.md). The index is built from
// the corpus of the case, as the benchmark builds it, and counted afterwards,
// not after churn:
//
//	go run ./cmd/objstat                                  # every case of the release suite
//	go run ./cmd/objstat -keys u64,uuid -values unique    # a subset
//	go run ./cmd/objstat -sizes 4096,16384 -max=false     # smaller sizes only
//
// With -entries it prints a second table: for the entries with several values
// (the leaves of a tree with pages), how many would fit one object of 512 bytes
// that holds the entry inline (PLAN step 3, the single-key page), when its
// header has room for N values; see entryRow.
//
// The output is a Markdown table, one row per case. What it counts:
//
//   - block: the Go size class an object occupies, with the 8-byte malloc
//     header of objects with pointers above 512 bytes (art.Block). The tool
//     counts the index only: not what a value set of a set leaf allocates for
//     itself (2% of the leaves of multi maps), nor the keys and values the
//     caller owns.
//   - not x64, not x128: the share of objects whose block is not a multiple of
//     64 or of 128 bytes.
//   - line overflow: the expected share of objects that touch more 64-byte
//     lines than their size needs, averaged over the positions a block of its
//     size takes in a span (the real addresses are not looked at).
//   - mix: the share of each kind of object.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/art"
)

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "objstat:", err)
		os.Exit(1)
	}
}

// run builds the cases the arguments select and writes the table to w.
func run(w io.Writer, args []string) error {
	fs := flag.NewFlagSet("objstat", flag.ContinueOnError)
	kindsF := fs.String("keys", "u64,str,uuid,email,url,path,street,dirs", "key kinds")
	valuesF := fs.String("values", "multi,unique", "value profiles: multi (a skewed number of values per key) and unique (one value per key)")
	strF := fs.Bool("strvals", true, "also measure every profile with string values (the bench's strvals build)")
	sizesF := fs.String("sizes", "4096,16384,262144,1048576", "numbers of keys")
	entriesF := fs.Bool("entries", false, "print the table of entries with several values (the single-key page statistic) after the object table")
	maxF := fs.Bool("max", true, "for a kind whose corpus holds fewer keys than a size, measure at the largest size the corpus allows (path, street)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	kinds, err := parseKinds(*kindsF)
	if err != nil {
		return err
	}
	sizes, err := parseInts(*sizesF)
	if err != nil {
		return err
	}
	var profiles []string
	for _, p := range strings.Split(*valuesF, ",") {
		if p != "multi" && p != "unique" {
			return fmt.Errorf("unknown value profile %q (multi or unique)", p)
		}
		profiles = append(profiles, p)
		if *strF {
			profiles = append(profiles, p+"-str")
		}
	}
	emit := func(line string) error { // one row at a time: a run with 1M keys takes minutes
		_, err := fmt.Fprintln(w, line)
		return err
	}
	if err := emit("| case | objects | block bytes per key | not x64 | not x128 | line overflow | mix |\n|---|--:|--:|--:|--:|--:|---|"); err != nil {
		return err
	}
	var entries []string
	for _, kind := range kinds {
		for _, profile := range profiles {
			for _, n := range sizesOf(kind, sizes, *maxF) {
				name := fmt.Sprintf("%s %s %s", kind, profile, human(n))
				s := measure(kind, profile, n)
				if err := s.verify(name, n); err != nil {
					return err
				}
				if err := emit(s.row(name, n)); err != nil {
					return err
				}
				entries = append(entries, s.entryRow(name, n))
			}
		}
	}
	if !*entriesF {
		return nil
	}
	if err := emit("\n" + entryHeader()); err != nil {
		return err
	}
	for _, row := range entries {
		if err := emit(row); err != nil {
			return err
		}
	}
	return nil
}

// sizesOf returns the sizes to measure for kind: those the corpus holds, and
// with max the corpus's own maximum in place of the sizes beyond it.
func sizesOf(kind keys.Kind, sizes []int, max bool) []int {
	var out []int
	capacity := keys.Capacity(kind)
	for _, n := range sizes {
		switch {
		case n <= capacity:
			out = append(out, n)
		case max && !slices.Contains(out, capacity):
			out = append(out, capacity)
		}
	}
	return out
}

func parseKinds(s string) ([]keys.Kind, error) {
	known := []keys.Kind{keys.U64, keys.Str, keys.UUID, keys.Email, keys.URL, keys.Path, keys.Street, keys.Dirs}
	var out []keys.Kind
	for _, name := range strings.Split(s, ",") {
		k := keys.Kind(name)
		if !slices.Contains(known, k) {
			return nil, fmt.Errorf("unknown key kind %q", name)
		}
		out = append(out, k)
	}
	return out, nil
}

func parseInts(s string) ([]int, error) {
	var out []int
	for _, f := range strings.Split(s, ",") {
		n, err := strconv.Atoi(f)
		if err != nil || n < 1 {
			return nil, fmt.Errorf("bad size %q", f)
		}
		out = append(out, n)
	}
	return out, nil
}

// human writes n as the benchmark does: 4K, 256K, 1M, or the number itself.
func human(n int) string {
	switch {
	case n%(1<<20) == 0:
		return strconv.Itoa(n>>20) + "M"
	case n%(1<<10) == 0:
		return strconv.Itoa(n>>10) + "K"
	}
	return strconv.Itoa(n)
}

// stat adds up the objects of one case.
type stat struct {
	objects, bytes int
	not64, not128  int
	overflow       float64 // in objects
	mix            map[string]int
	keysInObjects  int
	leaves         []leafInfo // the leaves, for the table of -entries
	valueBytes     int        // the bytes of one value, as the single-key page would store it
}

// leafInfo is what the entry statistic needs of a leaf.
type leafInfo struct{ remainder, values int }

// add counts one object.
func (s *stat) add(o art.Object) {
	block, offset := art.Block(o.Size, o.Pointers)
	s.objects++
	s.bytes += block
	s.keysInObjects += o.Keys
	if o.Values > 0 {
		s.leaves = append(s.leaves, leafInfo{o.Remainder, o.Values})
	}
	if block%64 != 0 {
		s.not64++
	}
	if block%128 != 0 {
		s.not128++
	}
	s.overflow += lineOverflow(o.Size, block, offset)
	if s.mix == nil {
		s.mix = map[string]int{}
	}
	s.mix[o.Label]++
}

// verify checks that the objects hold the n keys of the case: an object kind
// the statistic does not know would drop keys from its figures.
func (s *stat) verify(name string, n int) error {
	if s.keysInObjects != n {
		return fmt.Errorf("%s: the objects hold %d keys, the corpus has %d (art.Map.Objects misses an object kind?)", name, s.keysInObjects, n)
	}
	return nil
}

// lineOverflow returns the share of the positions a block of the given size
// takes, among the 64 a span has for it, at which an object of size bytes
// that starts offset bytes into the block touches more 64-byte lines than
// ceil(size/64). A span is aligned to 8 KiB and holds blocks one after the
// other, so block k starts at k*block; k running over 64 values visits every
// offset modulo 64 that block can have.
func lineOverflow(size, block, offset int) float64 {
	minLines := (size + 63) / 64
	over := 0
	for k := range 64 {
		start := (k*block + offset) % 64
		if (start+size-1)/64+1 > minLines {
			over++
		}
	}
	return float64(over) / 64
}

// row formats the table row of a case of n keys.
func (s *stat) row(name string, n int) string {
	pct := func(x float64) string { return strconv.FormatFloat(100*x/float64(s.objects), 'f', 1, 64) + " %" }
	labels := make([]string, 0, len(s.mix))
	for l := range s.mix {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	var mix []string
	for _, l := range labels {
		mix = append(mix, fmt.Sprintf("%s %.0f%%", l, 100*float64(s.mix[l])/float64(s.objects)))
	}
	return fmt.Sprintf("| %s | %d | %.1f | %s | %s | %s | %s |", name, s.objects,
		float64(s.bytes)/float64(n), pct(float64(s.not64)), pct(float64(s.not128)), pct(s.overflow), strings.Join(mix, "; "))
}

// measure builds the index of a case, and counts its objects.
func measure(kind keys.Kind, profile string, n int) *stat {
	unique := strings.HasPrefix(profile, "unique")
	if strings.HasSuffix(profile, "-str") {
		return build(kind, n, unique, func(v uint64) string { return fmt.Sprintf("%016x", v) })
	}
	return build(kind, n, unique, func(v uint64) uint64 { return v })
}

// build fills an index with the keys of the corpus and the values of the
// profile, as the benchmark's fixture does (cmd/bench: profileValues), and
// counts its objects.
func build[T comparable](kind keys.Kind, n int, unique bool, value func(uint64) T) *stat {
	c := keys.Generate(kind, n, 0x5EED)
	vals, offs := keys.Values(n, 0xFA11)
	if c.Natural != nil { // street names hold the localities they have
		vals, offs = nil, make([]int, n+1)
		for i, vs := range c.Natural[:n] {
			vals = append(vals, vs...)
			offs[i+1] = len(vals)
		}
	}
	var m art.Map[T]
	for i, key := range c.Keys.B {
		vs := vals[offs[i]:offs[i+1]]
		if unique {
			vs = vs[:1]
		}
		for _, v := range vs {
			m.Add(key, value(v))
		}
	}
	var s stat
	m.Objects(s.add)
	s.valueBytes = 8
	if str, ok := any(value(0)).(string); ok {
		s.valueBytes = len(str) // the bytes of the string, which a layout of byte strings stores inline
	}
	return &s
}

// entryNs are the numbers of values a header may have room for in the table of
// -entries.
var entryNs = []int{3, 4, 5, 6, 7, 8, 10, 12, 14, 16}

// pageBytes is the largest object of the single-key page, as long as it keeps
// the grid.
const pageBytes = 512

// entryHeader returns the header of the table of -entries.
func entryHeader() string {
	cols := []string{"case", "leaves", "entries with 2+ values (of all keys)", "of them: 2-4 values", "5-16", "17+", "remainder > 503"}
	for _, n := range entryNs {
		cols = append(cols, fmt.Sprintf("fits N=%d", n))
	}
	return "| " + strings.Join(cols, " | ") + " |\n|---|" + strings.Repeat("--:|", len(cols)-1)
}

// headerBytes is the size of the header of a single-key page that has room
// for lengths of n values: two bytes (type, remainder length) and one byte per
// value, rounded up to a multiple of 8 so that the data starts at a multiple
// of 8.
func headerBytes(n int) int { return (2 + n + 7) / 8 * 8 }

// entryRow formats the row of a case in the table of -entries: among the
// leaves with several values, how many would fit one object of 512 bytes, when
// the header has room for N values (fits N): the entry has at most N values
// and header, remainder and values add up to at most 512 bytes. The rest would
// take the fallback (a value set, or an object of its own beyond the grid).
// "remainder > 503" counts the entries whose remainder alone does not fit
// behind a header of 8 bytes: they would be oversized objects. The remainder
// is the one the leaf of the tree holds now, from its base on.
func (s *stat) entryRow(name string, n int) string {
	multi, small, mid, big, long := 0, 0, 0, 0, 0
	fits := make([]int, len(entryNs))
	for _, l := range s.leaves {
		if l.values < 2 {
			continue
		}
		multi++
		switch {
		case l.values <= 4:
			small++
		case l.values <= 16:
			mid++
		default:
			big++
		}
		if l.remainder > pageBytes-8 {
			long++
		}
		for i, w := range entryNs {
			if l.values <= w && headerBytes(w)+l.remainder+l.values*s.valueBytes <= pageBytes {
				fits[i]++
			}
		}
	}
	pct := func(x, of int) string {
		if of == 0 {
			return "-"
		}
		return strconv.FormatFloat(100*float64(x)/float64(of), 'f', 1, 64) + " %"
	}
	cells := []string{name, strconv.Itoa(len(s.leaves)), pct(multi, n), pct(small, multi), pct(mid, multi), pct(big, multi), pct(long, multi)}
	for _, f := range fits {
		cells = append(cells, pct(f, multi))
	}
	return "| " + strings.Join(cells, " | ") + " |"
}
