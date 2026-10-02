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
	"os"
	"slices"
	"sort"
	"strings"

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
	kindsF := fs.String("keys", "u64,str,uuid,email,url,path,street", "key kinds")
	n := fs.Int("n", 262144, "keys per kind (at most what the corpus holds)")
	chunk := fs.Int("chunk", 256, "keys below a range node")
	maxClass := fs.Int("maxclass", 2, "largest page class: 2 is 512 bytes, 3 is 1024")
	splitFill := fs.Int("splitfill", vpage.SplitFill, "percent a half of a split page may fill")
	byBytes := fs.Bool("splitbybytes", false, "split where the bytes are halved, not the count")
	if err := fs.Parse(args); err != nil {
		return err
	}
	vpage.MaxClass, vpage.SplitFill, vpage.SplitByBytes = *maxClass, *splitFill, *byBytes
	emit := func(format string, args ...any) error {
		_, err := fmt.Fprintf(w, format+"\n", args...)
		return err
	}
	if err := emit("chunk %d keys, largest class %d, split fill %d%%, split by bytes %v\n", *chunk, *maxClass, *splitFill, *byBytes); err != nil {
		return err
	}
	if err := emit("| kind | keys | suffix B | pages | keys/page | fill | uniform | page B/key | router B/key | total B/key | tail B/key | prefix B/key | too long |\n|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|"); err != nil {
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

// sharedPrefix returns how many bytes all suffixes of page p share: those of its
// first and last key.
func sharedPrefix(p *vpage.Page) int {
	var first, last [255]byte
	a := p.Key(0, &first)
	b := p.Key(p.Len()-1, &last)
	return lcp(a, b)
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

func measure(kind keys.Kind, n, chunk int) string {
	c := keys.Generate(kind, n, 0x5EED)
	sorted := keys.Sorted(c.Keys).B
	bounds := make([][]byte, 0, n/chunk+1) // the first key of each chunk
	bases := make([]int, 0, n/chunk+1)
	for i := 0; i < n; i += chunk {
		end := min(i+chunk, n)
		base := lcp(sorted[i], sorted[end-1]) // sorted: the first and last share what all share
		bounds = append(bounds, sorted[i])
		bases = append(bases, base)
	}
	runs := make([]vpage.Run, len(bounds))
	var suffixBytes, tooLong int
	for _, key := range c.Keys.B {
		ci := sort.Search(len(bounds), func(i int) bool { return string(bounds[i]) > string(key) }) - 1
		s := key[bases[ci]:]
		if err := runs[ci].Insert(s, 1); err != nil {
			tooLong++
			continue
		}
		suffixBytes += len(s)
	}
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
			prefix += (st.Keys - 1) * sharedPrefix(p)
		}
	}
	f := func(x int) float64 { return float64(x) / float64(keysIn) }
	return fmt.Sprintf("| %s | %d | %.1f | %d | %.1f | %.0f %% | %.0f %% | %.1f | %.1f | %.1f | %.1f | %.1f | %d |",
		kind, keysIn, f(suffixBytes), pages, f(keysIn)*float64(keysIn)/float64(pages), 100*float64(used)/float64(size),
		100*float64(uniform)/float64(pages), f(size), f(router), f(size+router), f(tails), f(prefix), tooLong)
}
