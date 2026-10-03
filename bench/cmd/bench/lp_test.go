//go:build strvals

package main

import (
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
)

// TestOrderedLPAgrees makes sure the candidate with length-header pages answers
// like the ordered one before it is timed against it.
//
// A reader of the benchmark trusts that a faster candidate holds the same
// values. This covers the verification of the benchmark driver for the
// experimental candidate ordered-lpage on small corpora of several kinds of
// keys and both value profiles: points, ranges, prefixes, and the contents the
// build stream leaves behind.
func TestOrderedLPAgrees(t *testing.T) {
	st := stream{ratio: 1.5}
	for _, kind := range []keys.Kind{keys.U64, keys.Str, keys.Path, keys.Street} {
		for _, profile := range []string{multi, unique} {
			t.Run(string(kind)+"/"+profile, func(t *testing.T) {
				vsOnly = lpImpls
				defer func() { vsOnly = nil }()
				impls := implsFor(profile)
				if len(impls) != 5 || impls[1] != orderedLP || impls[4] != orderedLPMZ {
					t.Fatalf("candidates %v", impls)
				}
				f := newFixture(kind, 3000, profile, impls, st, func([]string) {})
				if err := f.verify(); err != nil {
					t.Fatal(err)
				}
				if err := f.verifyBuild(append([]string{ordered}, lpImpls...)...); err != nil {
					t.Fatal(err)
				}
				for _, l := range f.lps {
					if l.NumberOfKeys() != len(f.c.Keys.B) {
						t.Fatalf("%d keys, want %d", l.NumberOfKeys(), len(f.c.Keys.B))
					}
				}
			})
		}
	}
}
