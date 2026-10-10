package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// linksDumpURL is the dump of Simple English Wikipedia of 2026-10-01; the
// files are named <linksDumpURL><table>.sql.gz. Wikimedia keeps a dump for
// about ten months only, so the built corpus is kept as a release asset of
// the repository (see keys/testdata/README.md).
const linksDumpURL = "https://dumps.wikimedia.org/simplewiki/20261001/simplewiki-20261001-"

// linkKey is a page of namespace 0 with the link targets of namespace 0 that
// it links to.
type linkKey struct {
	title    string
	redirect bool
	targets  []uint32 // ids in linktarget
}

// linkGraph is what the three tables of the dump say about the links between
// the pages of namespace 0.
type linkGraph struct {
	keys    []linkKey         // the pages that link to at least one target, ascending by title
	targets map[uint32]string // the titles of the targets of namespace 0 by their id
	exists  map[string]bool   // the titles of the pages of namespace 0
}

// links writes the corpus of page links: the titles of the link targets, then
// one line per page: its title, a tab and the comma-separated indexes of its
// targets in the list. Every page that links to a page of namespace 0 is in
// it, with all its links: the corpus is the wiki, not a sample of it. See
// keys/testdata/README.md.
func links(path string) error {
	g, err := readLinks(downloadTable)
	if err != nil {
		return err
	}
	titles, err := targetTitles(g)
	if err != nil {
		return err
	}
	fmt.Print(linkStats(g, titles))
	return writeGzip(path, func(w *strings.Builder) { formatLinks(w, g, titles) })
}

// downloadTable returns the unpacked SQL dump of a table of the Wikipedia
// dump; the caller closes it.
func downloadTable(table string) (io.ReadCloser, error) {
	body, err := get(linksDumpURL + table + ".sql.gz")
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(body)
	if err != nil {
		_ = body.Close()
		return nil, fmt.Errorf("%s: %w", table, err)
	}
	return struct {
		io.Reader
		io.Closer
	}{zr, body}, nil
}

// eachTableRow calls fn for every row of a table, which open provides. The
// tables are read one after the other, so that no download waits for minutes
// with its connection open.
func eachTableRow(open func(string) (io.ReadCloser, error), table string, fn func([]field) error) error {
	r, err := open(table)
	if err != nil {
		return err
	}
	defer func() { _ = r.Close() }()
	return eachRow(r, table, fn)
}

// readLinks builds the link graph from the dumps of the tables page,
// linktarget and pagelinks, which open provides by table name (in the layout
// of MediaWiki 1.43 and later: the pagelinks refer to linktarget rows instead
// of holding titles).
func readLinks(open func(string) (io.ReadCloser, error)) (*linkGraph, error) {
	g := &linkGraph{targets: map[uint32]string{}, exists: map[string]bool{}}
	// linktarget: (lt_id, lt_namespace, lt_title)
	err := eachTableRow(open, "linktarget", func(r []field) error {
		id, ns, err := numbers(r, 3)
		if err == nil && ns == 0 {
			g.targets[uint32(id)] = string(r[2].b)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("linktarget: %w", err)
	}
	// page: (page_id, page_namespace, page_title, page_is_redirect, ...)
	type pageRow struct {
		title    string
		redirect bool
	}
	pages := map[uint32]pageRow{}
	err = eachTableRow(open, "page", func(r []field) error {
		id, ns, err := numbers(r, 4)
		if err != nil || ns != 0 {
			return err
		}
		redirect, err := strconv.ParseUint(string(r[3].b), 10, 8)
		pages[uint32(id)] = pageRow{string(r[2].b), redirect == 1}
		g.exists[string(r[2].b)] = true
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("page: %w", err)
	}
	// pagelinks: (pl_from, pl_from_namespace, pl_target_id), a link per row
	var pairs []uint64 // page id << 32 | target id
	err = eachTableRow(open, "pagelinks", func(r []field) error {
		from, ns, err := numbers(r, 3)
		if err != nil || ns != 0 {
			return err
		}
		target, err := strconv.ParseUint(string(r[2].b), 10, 32)
		if _, ok := g.targets[uint32(target)]; err != nil || !ok {
			return err // a target of another namespace
		}
		if _, ok := pages[uint32(from)]; !ok {
			return fmt.Errorf("page %d of namespace 0 links but is not in the page table", from)
		}
		pairs = append(pairs, from<<32|target)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("pagelinks: %w", err)
	}
	slices.Sort(pairs)
	for i := 0; i < len(pairs); {
		from := pairs[i] >> 32
		k := linkKey{title: pages[uint32(from)].title, redirect: pages[uint32(from)].redirect}
		for ; i < len(pairs) && pairs[i]>>32 == from; i++ {
			k.targets = append(k.targets, uint32(pairs[i]))
		}
		g.keys = append(g.keys, k)
	}
	sort.Slice(g.keys, func(i, j int) bool { return g.keys[i].title < g.keys[j].title })
	return g, nil
}

// numbers returns the integers in the first two columns of a row of a table
// that has at least want columns, which every table here starts with: the id
// of the page or target and its namespace, or the id of the linking page and
// its namespace.
func numbers(r []field, want int) (id uint64, ns int64, err error) {
	if len(r) < want {
		return 0, 0, fmt.Errorf("a row has %d columns, want at least %d", len(r), want)
	}
	if id, err = strconv.ParseUint(string(r[0].b), 10, 64); err != nil {
		return 0, 0, err
	}
	ns, err = strconv.ParseInt(string(r[1].b), 10, 64) // namespaces can be negative
	return id, ns, err
}

// targetTitles returns the titles of all targets that the pages link to,
// ascending. A title that would break the line format is an error.
func targetTitles(g *linkGraph) ([]string, error) {
	seen := map[uint32]bool{}
	var titles []string
	for _, k := range g.keys {
		if err := checkTitle(k.title); err != nil {
			return nil, err
		}
		for _, id := range k.targets {
			if !seen[id] {
				seen[id] = true
				if err := checkTitle(g.targets[id]); err != nil {
					return nil, err
				}
				titles = append(titles, g.targets[id])
			}
		}
	}
	slices.Sort(titles)
	return titles, nil
}

// checkTitle rejects titles that are not UTF-8 or hold a tab or a line break,
// which the format of the corpus file cannot hold.
func checkTitle(t string) error {
	if !utf8.ValidString(t) || t == "" || strings.ContainsAny(t, "\t\r\n") {
		return fmt.Errorf("the title %q cannot be stored in the corpus", t)
	}
	return nil
}

// formatLinks writes the number of target titles, the titles one per line in
// ascending order, then one line per key: its title, a tab and the ascending,
// comma-separated indexes of its targets in the list of titles.
func formatLinks(w *strings.Builder, g *linkGraph, titles []string) {
	index := make(map[string]int, len(titles))
	w.WriteString(strconv.Itoa(len(titles)) + "\n")
	for i, t := range titles {
		index[t] = i
		w.WriteString(t + "\n")
	}
	var ids []int
	for _, k := range g.keys {
		ids = ids[:0]
		for _, t := range k.targets {
			ids = append(ids, index[g.targets[t]])
		}
		slices.Sort(ids)
		w.WriteString(k.title)
		for i, id := range ids {
			if i == 0 {
				w.WriteByte('\t')
			} else {
				w.WriteByte(',')
			}
			w.WriteString(strconv.Itoa(id))
		}
		w.WriteByte('\n')
	}
}

// linkStats returns what the README reports about the corpus: the number of
// keys and values, how many values the keys hold, and how many of the values
// stand in large sets.
func linkStats(g *linkGraph, titles []string) string {
	var w strings.Builder
	bounds := []int{1, 2, 4, 8, 16, 64, 256, 1 << 30}
	labels := []string{"1", "2", "3-4", "5-8", "9-16", "17-64", "65-256", "257+"}
	keysIn, valsIn := make([]int, len(bounds)), make([]int, len(bounds))
	counts := make([]int, len(g.keys))
	total, keyBytes, redirects := 0, 0, 0
	for i, k := range g.keys {
		n := len(k.targets)
		counts[i], total, keyBytes = n, total+n, keyBytes+len(k.title)
		if k.redirect {
			redirects++
		}
		b, _ := slices.BinarySearch(bounds, n)
		keysIn[b]++
		valsIn[b] += n
	}
	slices.Sort(counts)
	red := 0
	for _, t := range titles {
		if !g.exists[t] {
			red++
		}
	}
	nk := float64(len(g.keys))
	fmt.Fprintf(&w, "links: %d keys, %d values, %d target titles (%d of them no page: red links)\n",
		len(g.keys), total, len(titles), red)
	fmt.Fprintf(&w, "links: values a key: mean %.1f, median %d, max %d; key length mean %.1f B; redirects %.1f %% of the keys\n",
		float64(total)/nk, counts[len(counts)/2], counts[len(counts)-1], float64(keyBytes)/nk, 100*float64(redirects)/nk)
	w.WriteString("links: values a key   keys     values\n")
	for i, l := range labels {
		fmt.Fprintf(&w, "links: %-13s %5.1f %%  %5.1f %%\n", l, 100*float64(keysIn[i])/nk, 100*float64(valsIn[i])/float64(total))
	}
	big := 0
	for _, v := range valsIn[6:] {
		big += v
	}
	small := keysIn[0] + keysIn[1] + keysIn[2] + keysIn[3]
	fmt.Fprintf(&w, "links: keys with 1 / 2-8 / 9-64 / 65+ values: %.1f / %.1f / %.1f / %.1f %%; values in sets of 65+: %.1f %%\n",
		100*float64(keysIn[0])/nk, 100*float64(small-keysIn[0])/nk, 100*float64(keysIn[4]+keysIn[5])/nk, 100*float64(keysIn[6]+keysIn[7])/nk,
		100*float64(big)/float64(total))
	return w.String()
}
