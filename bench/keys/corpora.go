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
//go:embed testdata/paths.txt.gz testdata/streets.tsv.gz testdata/hosts.txt.gz
var corpora embed.FS

// streets is the street corpus: names in ascending order and, for each, the
// localities that have a street of that name, as their index plus one: the
// benchmark takes a value sum of zero for a missing key.
type streets struct {
	names  []string
	locs   [][]uint64
	places []string // the names of the localities: value v is places[v-1]
}

// dirs is the directory corpus, derived from the file paths: the directories
// that hold at least one file of the sample, each with the names of its files
// as their number plus one (an interned file name has one number for every
// directory it occurs in).
type dirs struct {
	names []string   // "/usr/share/doc/foo/", with the closing slash
	files [][]uint64 // for each directory the numbers of its file names, in path order
	file  []string   // the file names: name v is file[v-1]
}

var (
	pathCorpus   = sync.OnceValue(func() []string { return readLines("testdata/paths.txt.gz") })
	dirCorpus    = sync.OnceValue(loadDirs)
	streetCorpus = sync.OnceValue(loadStreets)
	hostCorpus   = sync.OnceValue(func() []string { return readLines("testdata/hosts.txt.gz") })
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
	s := streets{names: make([]string, len(rows)), locs: make([][]uint64, len(rows)), places: lines[1 : 1+nl]}
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
			s.locs[i] = append(s.locs[i], v+1)
		}
	}
	return s
}

// loadDirs groups the file paths of the path corpus by directory. The paths are
// sorted, so the directories come in the order of their first file, and the
// result does not depend on anything but the corpus file.
func loadDirs() dirs {
	var d dirs
	index := map[string]int{}
	number := map[string]uint64{}
	for _, p := range pathCorpus() {
		i := strings.LastIndexByte(p, '/')
		dir, file := p[:i+1], p[i+1:]
		if file == "" {
			continue
		}
		di, ok := index[dir]
		if !ok {
			di = len(d.names)
			index[dir] = di
			d.names = append(d.names, dir)
			d.files = append(d.files, nil)
		}
		v, ok := number[file]
		if !ok {
			d.file = append(d.file, file)
			v = uint64(len(d.file))
			number[file] = v
		}
		d.files[di] = append(d.files[di], v)
	}
	return d
}

// fromList draws n keys and n misses from all without repetition; natural,
// if not nil, holds the natural values of every key of all, and names, if not
// nil, the strings those values stand for (value v is names[v-1]).
func fromList(all []string, natural [][]uint64, names []string, n int, rng *rtcompare.DPRNG) Corpus {
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
	return Corpus{Keys: Pack(keys), Hits: Pack(hits), Misses: Pack(misses), Natural: nat, Names: names}.withProbes(hits, rng)
}
