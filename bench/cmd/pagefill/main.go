// Command pagefill measures what the page of the redesign (docs/redesign, PLAN
// step 1; package internal/vpage) costs per key for the keys of each benchmark
// corpus.
//
// It models the tree around the pages: the sorted keys are cut into chunks of
// -chunk keys, the part of a key that all keys of its chunk share is stripped
// (the tree holds it in a node), and the stripped suffixes of a chunk are
// inserted into a run of pages in the random order of the corpus, which splits
// full pages the way random inserts do. A chunk stands for a range node; its
// run for the pages below it. The bytes of the range nodes are added from the
// number of pages of each chunk, in the sizes of branch node-pages
// (rnode.go: 128, 256, 512 or 2112 bytes for up to 8, 24, 56 or 256 pages).
//
//	go run ./cmd/pagefill -keys u64,uuid -n 262144
//
// The output is a Markdown table, one row per key kind: the bytes of pages and
// of range nodes per key, the fill of the pages (what their keys need, of
// their size), how many keys a page holds, which share of the pages is
// uniform, the tails that the heads do not hold, and the bytes a key would save
// if its page held the prefix all its keys share once (prefix B/key; the page
// does not do this yet).
package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"runtime"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/vpage"
)

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "pagefill:", err)
		os.Exit(1)
	}
}

func run(w io.Writer, args []string) error {
	fs := flag.NewFlagSet("pagefill", flag.ContinueOnError)
	kindsF := fs.String("keys", "u64,str,uuid,email,url,path,street,dirs", "key kinds")
	n := fs.Int("n", 262144, "keys per kind (at most what the corpus holds)")
	chunk := fs.Int("chunk", 256, "keys below a range node")
	maxClass := fs.Int("maxclass", 2, "largest page class: 2 is 512 bytes, 3 is 1024")
	splitFill := fs.Int("splitfill", vpage.SplitFill, "percent a half of a split page may fill")
	byBytes := fs.Bool("splitbybytes", false, "split where the bytes are halved, not the count")
	minPrefix := fs.Int("minprefix", vpage.MinPrefix, "shortest prefix a page stores once (above 255: none)")
	minGain := fs.Int("mingain", vpage.MinGain, "bytes a prefix must save, after what it takes")
	slack := fs.Int("slack", vpage.PrefixSlack, "bytes a rebuild keeps less than the keys share")
	layout := fs.String("layout", "vpage", "the page to model: vpage (step 2) or lens (the sketch with a header of lengths, PLAN step 3)")
	lenFixed := fs.Bool("lenfixed", false, "lens: all values have one length, so the header has no length byte per value")
	lenValue := fs.Int("lenvalue", 8, "lens: bytes of a value")
	lenHeader := fs.Int("lenheader", 24, "lens: largest header in bytes, a multiple of 8")
	lenClasses := fs.String("lenclasses", "128,256,512", "lens: the object sizes, comma separated")
	timing := fs.Int("timing", 0, "instead of the table, time build and churn (this many operations) with the prefix off and on")
	if err := fs.Parse(args); err != nil {
		return err
	}
	vpage.MaxClass, vpage.SplitFill, vpage.SplitByBytes = *maxClass, *splitFill, *byBytes
	vpage.MinPrefix, vpage.MinGain, vpage.PrefixSlack = *minPrefix, *minGain, *slack
	emit := func(format string, args ...any) error {
		_, err := fmt.Fprintf(w, format+"\n", args...)
		return err
	}
	if err := emit("chunk %d keys, largest class %d, split fill %d%%, split by bytes %v, min prefix %d, min gain %d, slack %d\n", *chunk, *maxClass, *splitFill, *byBytes, *minPrefix, *minGain, *slack); err != nil {
		return err
	}
	if *timing > 0 {
		return timings(w, strings.Split(*kindsF, ","), min(*n, 262144), *chunk, *timing)
	}
	if *layout == "lens" {
		return runLens(w, strings.Split(*kindsF, ","), *n, *chunk, *lenFixed, *lenValue, *lenHeader, *lenClasses)
	}
	if err := emit("| kind | keys | suffix B | pages | keys/page | fill | uniform | page B/key | router B/key | total B/key | tail B/key | prefix B/page | too long |\n|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|"); err != nil {
		return err
	}
	for _, name := range strings.Split(*kindsF, ",") {
		kind := keys.Kind(name)
		if !slices.Contains(keys.Kinds, kind) {
			return fmt.Errorf("unknown key kind %q", name)
		}
		if err := emit("%s", measure(kind, min(*n, keys.Capacity(kind)), *chunk)); err != nil {
			return err
		}
	}
	return nil
}

// lcp returns the length of the longest common prefix of a and b.
func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	for i := range n {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

// routerBytes is the size of the range node over a chunk of pages.
func routerBytes(pages int) int {
	switch {
	case pages <= 8:
		return 128
	case pages <= 24:
		return 256
	case pages <= 56:
		return 512
	}
	return 2112
}

// model is the tree around the pages: the chunks of sorted keys, each with the
// base its keys share and a run of pages below it.
type model struct {
	keys   [][]byte // in the random order of the corpus
	bounds [][]byte // the first key of each chunk
	bases  []int
	runs   []vpage.Run
}

func newModel(kind keys.Kind, n, chunk int) *model {
	c := keys.Generate(kind, n, 0x5EED)
	sorted := keys.Sorted(c.Keys).B
	m := &model{keys: c.Keys.B}
	for i := 0; i < n; i += chunk {
		end := min(i+chunk, n)
		m.bounds = append(m.bounds, sorted[i])
		m.bases = append(m.bases, lcp(sorted[i], sorted[end-1])) // sorted: the first and last share what all share
	}
	m.runs = make([]vpage.Run, len(m.bounds))
	return m
}

// chunk returns the chunk key belongs to.
func (m *model) chunk(key []byte) int {
	return sort.Search(len(m.bounds), func(i int) bool { return string(m.bounds[i]) > string(key) }) - 1
}

// insert adds key; it fails for a suffix a page cannot hold.
func (m *model) insert(key []byte) error {
	ci := m.chunk(key)
	return m.runs[ci].Insert(key[m.bases[ci]:], 1)
}

// remove deletes key and reports whether it was there.
func (m *model) remove(key []byte) bool {
	ci := m.chunk(key)
	return m.runs[ci].Delete(key[m.bases[ci]:])
}

// fill inserts all keys in random order into empty runs.
func (m *model) fill() (suffixBytes, tooLong int) {
	m.runs = make([]vpage.Run, len(m.bounds))
	for _, key := range m.keys {
		if err := m.insert(key); err != nil {
			tooLong++
			continue
		}
		suffixBytes += len(key) - m.bases[m.chunk(key)]
	}
	return suffixBytes, tooLong
}

// timings prints, for each kind, what building the pages and then churning
// them (deleting a random key and putting it back) costs per operation with the
// prefix off and with it on, and how many pages were built anew. The settings
// alternate over three rounds and the fastest round counts. For diagnosis, not
// a claim of speed.
func timings(w io.Writer, kinds []string, n, chunk, ops int) error {
	if _, err := fmt.Fprintf(w, "| kind | prefix | build ns/key | churn ns/op | rebuilds/1000 ops | by shrinking | with a new prefix | pages |\n|---|---|--:|--:|--:|--:|--:|--:|\n"); err != nil {
		return err
	}
	on := vpage.MinPrefix
	for _, name := range kinds {
		kind := keys.Kind(name)
		if !slices.Contains(keys.Kinds, kind) {
			return fmt.Errorf("unknown key kind %q", name)
		}
		m := newModel(kind, min(n, keys.Capacity(kind)), chunk)
		type res struct {
			build, churn            time.Duration
			rebuilds, shrinks, newp int
			pages                   int
		}
		best := map[bool]res{}
		for round := range 3 {
			for _, prefix := range []bool{false, true} {
				vpage.MinPrefix = 256
				if prefix {
					vpage.MinPrefix = on
				}
				runtime.GC()
				t0 := time.Now()
				m.fill()
				build := time.Since(t0)
				r := rand.New(rand.NewPCG(uint64(round), 9))
				picks := make([][]byte, ops)
				for i := range picks {
					picks[i] = m.keys[r.IntN(len(m.keys))]
				}
				before := vpage.Counts
				t0 = time.Now()
				for _, k := range picks {
					if m.remove(k) {
						_ = m.insert(k)
					}
				}
				churn := time.Since(t0)
				pages := 0
				for i := range m.runs {
					pages += len(m.runs[i].Pages())
				}
				cur := res{build, churn, vpage.Counts.Rebuilds - before.Rebuilds, vpage.Counts.Shrinks - before.Shrinks, vpage.Counts.PrefixChanges - before.PrefixChanges, pages}
				if b, ok := best[prefix]; !ok || cur.build+cur.churn < b.build+b.churn {
					best[prefix] = cur
				}
			}
		}
		for _, prefix := range []bool{false, true} {
			b := best[prefix]
			if _, err := fmt.Fprintf(w, "| %s | %v | %.0f | %.0f | %.1f | %.1f | %.1f | %d |\n", kind, map[bool]string{false: "off", true: "on"}[prefix],
				float64(b.build.Nanoseconds())/float64(len(m.keys)), float64(b.churn.Nanoseconds())/float64(ops),
				1000*float64(b.rebuilds)/float64(ops), 1000*float64(b.shrinks)/float64(ops), 1000*float64(b.newp)/float64(ops), b.pages); err != nil {
				return err
			}
		}
	}
	vpage.MinPrefix = on
	return nil
}

func measure(kind keys.Kind, n, chunk int) string {
	m := newModel(kind, n, chunk)
	suffixBytes, tooLong := m.fill()
	runs := m.runs
	var pages, keysIn, size, used, tails, uniform, router, prefix int
	for i := range runs {
		ps := runs[i].Pages()
		router += routerBytes(len(ps))
		for _, p := range ps {
			st := p.Stats()
			pages++
			keysIn += st.Keys
			size += st.Size
			used += st.Used
			tails += st.Tails
			if st.Uniform {
				uniform++
			}
			prefix += p.PrefixLen()
		}
	}
	f := func(x int) float64 { return float64(x) / float64(keysIn) }
	return fmt.Sprintf("| %s | %d | %.1f | %d | %.1f | %.0f %% | %.0f %% | %.1f | %.1f | %.1f | %.1f | %.1f | %d |",
		kind, keysIn, f(suffixBytes), pages, f(keysIn)*float64(keysIn)/float64(pages), 100*float64(used)/float64(size),
		100*float64(uniform)/float64(pages), f(size), f(router), f(size+router), f(tails), float64(prefix)/float64(pages), tooLong)
}

// runLens writes the table of the lens layout (lens.go) for the key kinds.
func runLens(w io.Writer, kinds []string, n, chunk int, fixed bool, value, maxHeader int, classes string) error {
	l := lensLayout{fixed: fixed, value: value, maxHeader: maxHeader}
	for _, f := range strings.Split(classes, ",") {
		var c int
		if _, err := fmt.Sscanf(f, "%d", &c); err != nil || c < 64 {
			return fmt.Errorf("bad object size %q", f)
		}
		l.classes = append(l.classes, c)
	}
	slices.Sort(l.classes)
	if maxHeader < 8 || maxHeader%8 != 0 {
		return fmt.Errorf("the largest header must be a multiple of 8, not %d", maxHeader)
	}
	if _, err := fmt.Fprintf(w, "lens layout: fixed values %v, value %d B, largest header %d B (%d entries), classes %v\n\n%s\n", fixed, value, maxHeader, l.capacity(maxHeader), l.classes, lensTableHeader); err != nil {
		return err
	}
	for _, name := range kinds {
		kind := keys.Kind(name)
		if !slices.Contains(keys.Kinds, kind) {
			return fmt.Errorf("unknown key kind %q", name)
		}
		if _, err := fmt.Fprintln(w, lensRow(kind, min(n, keys.Capacity(kind)), chunk, l)); err != nil {
			return err
		}
	}
	return nil
}
