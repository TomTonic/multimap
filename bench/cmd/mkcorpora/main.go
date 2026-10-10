// Command mkcorpora downloads the real-world key corpora of the benchmark and
// writes them to keys/testdata, where package keys embeds them. It exists so
// that anyone can check where the data comes from and rebuild it byte for
// byte; the benchmark itself never touches the network.
//
//	go run ./cmd/mkcorpora                # all corpora, from the bench directory
//	go run ./cmd/mkcorpora hosts streets  # only these
//
// Four corpora:
//   - streets.tsv.gz: German street names and the localities that have a
//     street of that name, from the OpenPLZ API data, which is an extract of
//     OpenStreetMap (ODbL 1.0, see keys/testdata/README.md).
//   - paths.txt.gz: a deterministic sample of the file paths in the packages
//     of Debian 12 "bookworm" (main, all and amd64), from the archive's
//     Contents files as of the snapshot below.
//   - hosts.txt.gz: the most popular host names of the Tranco list below,
//     including subdomains, in rank order (see keys/testdata/README.md).
//   - links.tsv.gz: all pages of Simple English Wikipedia that link to
//     articles and the articles they link to, from the database dump of
//     2026-10-01 (CC BY-SA 4.0, see keys/testdata/README.md). It is written to
//     cache/, which git ignores, and not to keys/testdata. Rebuilding it
//     downloads 155 MB; it prints the statistics of the corpus.
package main

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	streetsURL = "https://raw.githubusercontent.com/openpotato/openplzapi.data/c26389b30573ab2738d0b14fd350f143aa0a1251/src/de/osm/streets.updated.csv"
	debianURL  = "https://snapshot.debian.org/archive/debian/20260901T000000Z/dists/bookworm/main/"
	// trancoURL is Tranco list 8P48V with subdomains, generated on
	// 2026-09-27 from the ranks of 2026-08-29 to 2026-09-27; a list ID is
	// permanent.
	trancoURL = "https://tranco-list.eu/download/8P48V/1000000"
	// pathSample is how many paths the sample keeps: enough for 256K keys
	// with as many miss keys (see keys.Generate), with some to spare.
	pathSample = 600_000
	// hostCount is how many hosts the host corpus keeps, from the top; the
	// url kind draws hosts by rank and rarely reaches further down.
	hostCount = 200_000
)

func main() {
	all := map[string]func(string) error{"streets": streets, "paths": paths, "hosts": hosts, "links": links}
	names := os.Args[1:]
	if len(names) == 0 {
		names = []string{"streets", "paths", "hosts", "links"}
	}
	for _, name := range names {
		build, ok := all[name]
		if !ok {
			fail(fmt.Errorf("unknown corpus %q; want streets, paths, hosts or links", name))
		}
		ext := ".txt.gz"
		if name == "streets" || name == "links" {
			ext = ".tsv.gz"
		}
		dir := filepath.Join("keys", "testdata")
		if name == "links" { // too big for the repository: see keys/testdata/README.md
			dir = "cache"
		}
		if err := os.MkdirAll(dir, 0o755); err != nil { //nolint:gosec // a directory of corpus files, nothing secret
			fail(err)
		}
		if err := build(filepath.Join(dir, name+ext)); err != nil {
			fail(err)
		}
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "mkcorpora:", err)
	os.Exit(1)
}

// userAgent names this tool to the servers it downloads from; Wikimedia
// refuses the default agent of Go's HTTP client (403) and asks clients for a
// descriptive one with a way to reach their authors.
const userAgent = "multimap-mkcorpora/1.0 (https://github.com/TomTonic/multimap)"

// get returns the body of url; the caller closes it.
func get(url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get %s: %w", url, err)
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req) //nolint:gosec // fixed URLs of this tool
	if err != nil {
		return nil, fmt.Errorf("get %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("get %s: %s", url, resp.Status)
	}
	return resp.Body, nil
}

// streets writes the street corpus: the number of localities, the localities
// one per line in ascending order, then one line per street name in
// ascending order with the indexes of its localities, tab-separated from the
// name and comma-separated among each other.
func streets(path string) error {
	body, err := get(streetsURL)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()
	r := csv.NewReader(body)
	r.FieldsPerRecord = -1
	if _, err := r.Read(); err != nil { // header
		return err
	}
	byName := map[string]map[string]bool{}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name := strings.TrimSpace(strings.Trim(strings.TrimSpace(rec[0]), `"`))
		loc := strings.TrimSpace(rec[2])
		if name == "" || loc == "" || strings.ContainsAny(name, "\t\n\r") || strings.ContainsAny(loc, "\n\r") {
			continue
		}
		if byName[name] == nil {
			byName[name] = map[string]bool{}
		}
		byName[name][loc] = true
	}
	locSet := map[string]bool{}
	for _, ls := range byName {
		for l := range ls {
			locSet[l] = true
		}
	}
	locs := slices.Sorted(func(yield func(string) bool) {
		for l := range locSet {
			if !yield(l) {
				return
			}
		}
	})
	idx := make(map[string]int, len(locs))
	for i, l := range locs {
		idx[l] = i
	}
	return writeGzip(path, func(w *strings.Builder) {
		w.WriteString(strconv.Itoa(len(locs)) + "\n")
		for _, l := range locs {
			w.WriteString(l + "\n")
		}
		for _, name := range sortedKeys(byName) {
			ids := make([]int, 0, len(byName[name]))
			for l := range byName[name] {
				ids = append(ids, idx[l])
			}
			slices.Sort(ids)
			w.WriteString(name)
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
	})
}

// paths writes a sample of pathSample distinct file paths, one per line in
// ascending order, each with a leading slash.
func paths(path string) error {
	seen := map[string]bool{}
	for _, f := range []string{"Contents-all.gz", "Contents-amd64.gz"} {
		if err := contents(debianURL+f, seen); err != nil {
			return err
		}
	}
	all := sortedKeys(seen)
	rng := rand.New(rand.NewPCG(2026, 9)) //nolint:gosec // a reproducible sample, not a secret
	rng.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })
	sample := all[:min(pathSample, len(all))]
	slices.Sort(sample)
	return writeGzip(path, func(w *strings.Builder) {
		for _, p := range sample {
			w.WriteString(p + "\n")
		}
	})
}

// contents adds the paths of one Contents file to seen. Each line is a path,
// whitespace, and the comma-separated packages that ship it; a path may
// contain spaces, the package list does not.
func contents(url string, seen map[string]bool) error {
	body, err := get(url)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()
	zr, err := gzip.NewReader(body)
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(zr)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		line := sc.Text()
		i := strings.LastIndexAny(line, " \t")
		if i <= 0 {
			continue
		}
		p := strings.TrimRight(line[:i], " \t")
		if p == "" || p == "FILE" {
			continue
		}
		seen["/"+p] = true
	}
	return sc.Err()
}

// validHost matches the host names the host corpus keeps: lowercase letters,
// digits, dashes and dots, at least two labels, at most 64 bytes. It drops
// the few service names with underscores and the long machine-made names of
// DNS infrastructure.
var validHost = regexp.MustCompile(`^[a-z0-9-]+(\.[a-z0-9-]+)+$`)

// hosts writes the first hostCount valid host names of the Tranco list, one
// per line in rank order.
func hosts(path string) error {
	body, err := get(trancoURL)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()
	r := csv.NewReader(body)
	var keep []string
	for len(keep) < hostCount {
		rec, err := r.Read()
		if err == io.EOF {
			return fmt.Errorf("%s: only %d valid hosts", trancoURL, len(keep))
		}
		if err != nil {
			return err
		}
		if h := rec[1]; len(h) <= 64 && validHost.MatchString(h) {
			keep = append(keep, h)
		}
	}
	return writeGzip(path, func(w *strings.Builder) {
		for _, h := range keep {
			w.WriteString(h + "\n")
		}
	})
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

func writeGzip(path string, write func(*strings.Builder)) error {
	var b strings.Builder
	write(&b)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	zw, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		_ = f.Close()
		return err
	}
	if _, err := io.WriteString(zw, b.String()); err != nil {
		_ = f.Close()
		return err
	}
	if err := zw.Close(); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}
