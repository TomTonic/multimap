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
// that only street names carry natural values.
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
			if (c.Natural != nil) != (kind == Street) {
				t.Fatalf("natural values: %v", c.Natural != nil)
			}
			for i, vs := range c.Natural {
				if len(vs) == 0 || slices.Contains(vs, 0) {
					t.Fatalf("street %q has localities %v; want at least one, none of them 0", c.Keys.S[i], vs)
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
	for _, kind := range []Kind{Path, Street} {
		if c := Capacity(kind); c < 100_000 || c > 1_000_000 {
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
	for i, vs := range s.locs {
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
