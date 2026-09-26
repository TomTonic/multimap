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
				if len(vs) == 0 {
					t.Fatalf("street %q has no localities", c.Keys.S[i])
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
