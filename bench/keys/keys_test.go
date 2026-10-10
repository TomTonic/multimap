package keys

import (
	"bytes"
	"maps"
	"slices"
	"strings"
	"testing"
)

// TestGenerate makes sure every key kind gives the benchmark what it relies
// on: distinct keys, misses that are really absent, hits that are the keys
// again in another order, and the same corpus for the same seed. It covers
// the synthetic generators and the real-world corpora of package keys, and
// that only the kinds with real values (street, dirs, links) carry natural values.
func TestGenerate(t *testing.T) {
	for _, kind := range Kinds {
		t.Run(string(kind), func(t *testing.T) {
			const n = 3000
			c := Generate(kind, n, 42)
			seen := map[string]bool{}
			for _, k := range c.Keys.S {
				if seen[k] || len(k) == 0 || len(k) > maxKeyLen {
					t.Fatalf("key %q is repeated, empty or too long", k)
				}
				seen[k] = true
			}
			for _, k := range c.Misses.S {
				if seen[k] {
					t.Fatalf("miss %q is a key", k)
				}
			}
			hits := map[string]bool{}
			for _, k := range c.Hits.S {
				hits[k] = true
			}
			if len(c.Keys.S) != n || len(c.Misses.S) != n || len(hits) != n || !maps.Equal(hits, seen) {
				t.Fatalf("%d keys, %d misses, %d distinct hits; want %d each, hits = keys", len(c.Keys.S), len(c.Misses.S), len(hits), n)
			}
			if again := Generate(kind, n, 42); !slices.Equal(again.Keys.S, c.Keys.S) {
				t.Fatal("the same seed gave another corpus")
			}
			if natural := kind == Street || kind == Dirs || kind == Links; (c.Natural != nil) != natural || (c.Names != nil) != natural {
				t.Fatalf("natural values: %v, names: %v", c.Natural != nil, c.Names != nil)
			}
			for i, vs := range c.Natural {
				if len(vs) == 0 || slices.Contains(vs, 0) || slices.Max(vs) > uint64(len(c.Names)) {
					t.Fatalf("key %q has the values %v; want at least one, none of them 0, none beyond the %d names", c.Keys.S[i], vs, len(c.Names))
				}
			}
		})
	}
}

// TestCapacity makes sure the benchmark learns how far a real-world corpus
// goes before it asks for more keys than there are. It covers Capacity and
// the check in Generate: synthetic kinds have no limit, corpus kinds hold
// what their files hold, and asking for more fails loudly.
func TestCapacity(t *testing.T) {
	if Capacity(UUID) < 1<<30 {
		t.Error("synthetic kinds must have no limit")
	}
	for _, kind := range []Kind{Path, Street, Dirs, Links} {
		if c := Capacity(kind); c < 50_000 || c > 1_000_000 {
			t.Errorf("%s: capacity %d", kind, c)
		}
	}
	defer func() {
		if recover() == nil {
			t.Error("Generate beyond the capacity must panic")
		}
	}()
	Generate(Street, Capacity(Street)+1, 1)
}

// TestPrefixEnd makes sure a prefix search over text keys finds exactly the
// keys that start with the typed characters. It covers Text and PrefixEnd on
// the street corpus, whose names share many prefixes.
func TestPrefixEnd(t *testing.T) {
	if Text(U64) || !Text(Street) {
		t.Fatal("only integer keys are not text")
	}
	c := Generate(Street, 20000, 3)
	for _, p := range []string{"Haup", "Bahn", "Am M", "Zu", "Ö"} {
		end := PrefixEnd([]byte(p))
		for _, k := range c.Keys.B {
			in := bytes.Compare(k, []byte(p)) >= 0 && bytes.Compare(k, end) <= 0
			if in != strings.HasPrefix(string(k), p) {
				t.Fatalf("key %q, prefix %q: in range %v", k, p, in)
			}
		}
	}
}

// TestPrefix makes sure a prefix search asks what a user would: the typed
// start of a name, or the directory of a path or URL. It covers Prefix for
// every text kind.
func TestPrefix(t *testing.T) {
	for _, tt := range []struct {
		kind      Kind
		key, want string
	}{
		{Street, "Hauptstr.", "Haup"},
		{Email, "ab@x.de", "ab@x"},
		{Street, "Ax", "Ax"},
		{Links, "List_of_cities", "List"},
		{Path, "/usr/share/doc/README", "/usr/share/doc/"},
		{URL, "https://a.example/items/12?ref=1", "https://a.example/items/"},
		{Path, "noslash", "nosl"},
	} {
		if got := Prefix(tt.kind, []byte(tt.key)); string(got) != tt.want {
			t.Errorf("Prefix(%s, %q) = %q, want %q", tt.kind, tt.key, got, tt.want)
		}
	}
}

// TestStreetValues makes sure every street of the real-world corpus maps to
// localities the benchmark can tell from a missing key. Package keys numbers
// the localities of streets.tsv.gz for the street kind's natural values; the
// benchmark reads a value sum of zero as a missing key, so no street of the
// whole corpus may have the value 0 or no value at all.
func TestStreetValues(t *testing.T) {
	s := streetCorpus()
	for i, vs := range s.vals {
		if len(vs) == 0 || slices.Contains(vs, 0) {
			t.Fatalf("street %q has localities %v; want at least one, none of them 0", s.names[i], vs)
		}
	}
}

// TestURL makes sure the url kind looks like the web to the multimap: real
// host names, many of them, none dominating, and a directory to search by
// that never reaches into the query. It covers the url model of package keys
// (url.go) over the Tranco host corpus: every key is an http or https
// address of a host from the corpus, served with or without "www.", the most
// frequent host holds under 2% of the keys, and Prefix ends at a slash
// before any query.
func TestURL(t *testing.T) {
	known := map[string]bool{}
	for _, h := range hostCorpus() {
		known[h] = true
	}
	c := Generate(URL, 50_000, 5)
	perHost := map[string]int{}
	for _, k := range c.Keys.S {
		rest, ok := strings.CutPrefix(k, "https://")
		if !ok {
			rest, ok = strings.CutPrefix(k, "http://")
		}
		host, _, slash := strings.Cut(rest, "/")
		if !ok || !slash || !known[host] && !known[strings.TrimPrefix(host, "www.")] {
			t.Fatalf("key %q is no address of a known host", k)
		}
		perHost[host]++
		if p := string(Prefix(URL, []byte(k))); !strings.HasSuffix(p, "/") || strings.Contains(p, "?") {
			t.Fatalf("key %q: directory %q", k, p)
		}
	}
	top := slices.Max(slices.Collect(maps.Values(perHost)))
	if len(perHost) < 10_000 || top > len(c.Keys.S)/50 {
		t.Fatalf("%d hosts, the top one holds %d of %d keys", len(perHost), top, len(c.Keys.S))
	}
}

// TestProbes makes sure the point-query benchmarks look up keys in an order that
// a branch predictor cannot learn. It belongs to the benchmark's key corpora: a
// probe sequence as long as the corpus repeats every n lookups, and at a few
// thousand keys the predictor learns the cycle, which favours the code with
// fewer branches over the code with more (docs/redesign, step 0). A corpus
// therefore offers ProbeLen probes, random picks from its keys when it holds
// fewer, and its own hits when it holds at least that many.
func TestProbes(t *testing.T) {
	t.Run("gives ProbeLen random picks from the keys of a small corpus", func(t *testing.T) {
		for _, kind := range []Kind{U64, Str, Street} {
			c := Generate(kind, 100, 42)
			if len(c.Probes.B) != ProbeLen || len(c.Probes.S) != ProbeLen {
				t.Fatalf("%s: %d probes, want %d", kind, len(c.Probes.B), ProbeLen)
			}
			keys := map[string]int{}
			for _, k := range c.Keys.S {
				keys[k] = 0
			}
			for _, p := range c.Probes.S {
				n, ok := keys[p]
				if !ok {
					t.Fatalf("%s: probe %q is not a key", kind, p)
				}
				keys[p] = n + 1
			}
			for k, n := range keys {
				if n < ProbeLen/100/3 { // a third of the mean: all keys are probed about equally often
					t.Fatalf("%s: key %q is probed %d times of %d", kind, k, n, ProbeLen)
				}
			}
			if again := Generate(kind, 100, 42); !slices.Equal(again.Probes.S, c.Probes.S) {
				t.Fatalf("%s: the same seed gave other probes", kind)
			}
		}
	})
	t.Run("does not repeat within ProbeLen probes", func(t *testing.T) {
		c := Generate(U64, 4096, 7)
		for _, period := range []int{4096, 8192, 65536} { // the periods of a cycle over the keys
			if slices.Equal(c.Probes.S[:period], c.Probes.S[period:2*period]) {
				t.Errorf("the probes repeat after %d", period)
			}
		}
	})
	t.Run("is the order of the hits for a corpus of at least ProbeLen keys", func(t *testing.T) {
		c := Generate(U64, ProbeLen, 3)
		if len(c.Probes.S) != ProbeLen || !slices.Equal(c.Probes.S, c.Hits.S) {
			t.Errorf("%d probes, equal to the hits: %v", len(c.Probes.S), slices.Equal(c.Probes.S, c.Hits.S))
		}
	})
}

// TestDirs makes sure the dirs kind is what its description says: a benchmark of
// entries with several values of variable length, made of real data. It
// covers the directory corpus (corpora.go) and Prefix: every key is a directory
// with a closing slash, every value is the name of a file in it, so that
// directory and name together are a path of the Debian sample, the names are
// strings of different lengths, most directories hold one file and a few hold
// many, and the prefix of a directory is the directory it lies in.
func TestDirs(t *testing.T) {
	paths := map[string]bool{}
	for _, p := range pathCorpus() {
		paths[p] = true
	}
	c := Generate(Dirs, 20_000, 7)
	single, many, shortest, longest := 0, 0, 1<<30, 0
	for i, k := range c.Keys.S {
		if !strings.HasSuffix(k, "/") {
			t.Fatalf("directory %q has no closing slash", k)
		}
		for _, v := range c.Natural[i] {
			name := c.Names[v-1]
			if !paths[k+name] {
				t.Fatalf("%q in %q is no path of the corpus", name, k)
			}
			shortest, longest = min(shortest, len(name)), max(longest, len(name))
		}
		switch n := len(c.Natural[i]); {
		case n == 1:
			single++
		case n > 100:
			many++
		}
	}
	if single < len(c.Keys.S)/3 || many == 0 || longest < 3*shortest {
		t.Errorf("%d of %d directories with one file, %d with more than 100, file names of %d to %d bytes", single, len(c.Keys.S), many, shortest, longest)
	}
	for in, want := range map[string]string{"/usr/share/doc/foo/": "/usr/share/doc/", "/bin/": "/", "/": ""} {
		if got := string(Prefix(Dirs, []byte(in))); got != want {
			t.Errorf("Prefix(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestLinks makes sure the links kind is what its description says: the page
// links of a wiki, where the typical page links to many others and a few to
// thousands. It covers the link corpus (corpora.go) and its use through
// Generate: the keys are page titles in ascending order, the values of a key are
// distinct titles of other pages in ascending order, about three in ten keys
// hold a single value (the redirects), about one in six holds 65 or more and
// those hold three of four values, and the biggest set holds thousands.
func TestLinks(t *testing.T) {
	all := linkCorpus()
	if !slices.IsSorted(all.names) || !slices.IsSorted(all.labels) || len(all.names) < 2*65536 {
		t.Fatalf("%d titles, sorted: %v, %d target titles sorted: %v; want at least %d titles in ascending order", len(all.names), slices.IsSorted(all.names), len(all.labels), slices.IsSorted(all.labels), 2*65536)
	}
	single, large, values, largeValues, biggest := 0, 0, 0, 0, 0
	for i, vs := range all.vals {
		if len(vs) == 0 || vs[len(vs)-1] > uint64(len(all.labels)) || vs[0] == 0 || strings.ContainsAny(all.names[i], " \t") {
			t.Fatalf("page %q has the values %v for %d target titles", all.names[i], vs, len(all.labels))
		}
		if !slices.IsSorted(vs) || len(slices.Compact(slices.Clone(vs))) != len(vs) {
			t.Fatalf("page %q links to %v; want distinct targets in ascending order", all.names[i], vs)
		}
		values += len(vs)
		biggest = max(biggest, len(vs))
		switch {
		case len(vs) == 1:
			single++
		case len(vs) >= 65:
			large++
			largeValues += len(vs)
		}
	}
	n := float64(len(all.vals))
	if s, l, v := float64(single)/n, float64(large)/n, float64(largeValues)/float64(values); s < 0.25 || s > 0.35 || l < 0.13 || l > 0.20 || v < 0.70 || v > 0.85 || biggest < 1000 {
		t.Errorf("%.1f%% of the pages with one link, %.1f%% with 65 or more holding %.1f%% of the links, the biggest page has %d", 100*s, 100*l, 100*v, biggest)
	}

	index := make(map[string]int, len(all.names))
	for i, name := range all.names {
		index[name] = i
	}
	c := Generate(Links, 20_000, 7)
	if len(c.Names) != len(all.labels) {
		t.Fatalf("%d names of values, want %d", len(c.Names), len(all.labels))
	}
	for i, k := range c.Keys.S {
		if !slices.Equal(c.Natural[i], all.vals[index[k]]) {
			t.Fatalf("page %q has the links %v in the corpus and %v in the benchmark", k, all.vals[index[k]], c.Natural[i])
		}
	}
}
