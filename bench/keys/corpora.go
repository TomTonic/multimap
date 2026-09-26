package keys

import (
	"bufio"
	"compress/gzip"
	"embed"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/TomTonic/rtcompare"
)

// The real-world corpora, made by cmd/mkcorpora. See testdata/README.md for
// their sources and licenses.
//
//go:embed testdata/paths.txt.gz testdata/streets.tsv.gz
var corpora embed.FS

// streets is the street corpus: names in ascending order and, for each, the
// indexes of the localities that have a street of that name.
type streets struct {
	names []string
	locs  [][]uint64
}

var (
	pathCorpus   = sync.OnceValue(func() []string { return readLines("testdata/paths.txt.gz") })
	streetCorpus = sync.OnceValue(loadStreets)
)

// readLines returns the lines of a gzipped corpus file. The corpora are part
// of this package, so a broken one is a bug, not an input error.
func readLines(name string) []string {
	f, err := corpora.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	zr, err := gzip.NewReader(f)
	if err != nil {
		panic(fmt.Sprintf("%s: %v", name, err))
	}
	var out []string
	sc := bufio.NewScanner(zr)
	sc.Buffer(make([]byte, 1<<16), 1<<16)
	for sc.Scan() {
		out = append(out, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("%s: %v", name, err))
	}
	return out
}

// loadStreets parses streets.tsv.gz: the number of localities, the
// localities, then per street name the name, a tab and the comma-separated
// indexes of its localities.
func loadStreets() streets {
	lines := readLines("testdata/streets.tsv.gz")
	nl, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(fmt.Sprintf("streets: %v", err))
	}
	rows := lines[1+nl:]
	s := streets{names: make([]string, len(rows)), locs: make([][]uint64, len(rows))}
	for i, row := range rows {
		name, ids, ok := strings.Cut(row, "\t")
		if !ok {
			panic(fmt.Sprintf("streets: line %q", row))
		}
		s.names[i] = name
		for id := range strings.SplitSeq(ids, ",") {
			v, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("streets: line %q: %v", row, err))
			}
			s.locs[i] = append(s.locs[i], v)
		}
	}
	return s
}

// fromList draws n keys and n misses from all without repetition; natural,
// if not nil, holds the natural values of every key of all.
func fromList(all []string, natural [][]uint64, n int, rng *rtcompare.DPRNG) Corpus {
	if 2*n > len(all) {
		panic(fmt.Sprintf("a corpus of %d keys cannot give %d keys and as many misses", len(all), n))
	}
	idx := make([]int32, len(all))
	for i := range idx {
		idx[i] = int32(i)
	}
	for i := range 2 * n { // partial Fisher-Yates: the first 2n are a random sample
		j := i + int(rng.Uint64()%uint64(len(idx)-i))
		idx[i], idx[j] = idx[j], idx[i]
	}
	keys, misses := make([][]byte, n), make([][]byte, n)
	var nat [][]uint64
	if natural != nil {
		nat = make([][]uint64, n)
	}
	for i := range n {
		keys[i], misses[i] = []byte(all[idx[i]]), []byte(all[idx[n+i]])
		if natural != nil {
			nat[i] = natural[idx[i]]
		}
	}
	hits := append([][]byte(nil), keys...)
	shuffle(hits, rng)
	return Corpus{Keys: Pack(keys), Hits: Pack(hits), Misses: Pack(misses), Natural: nat}
}
