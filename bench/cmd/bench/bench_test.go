package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
)

// TestPairsFor makes sure the benchmark compares exactly what a user choosing
// a multimap wants compared, and skips what cannot be timed sensibly. It
// covers the scenario plan of the benchmark driver: Ordered against every
// other candidate in every operation, range queries on the scanning
// candidates and builds only up to their size limits.
func TestPairsFor(t *testing.T) {
	ops := []string{"valuesFor", "valuesBetween", "churn", "build"}
	tests := []struct {
		name  string
		n     int
		count int
		skip  []pair
	}{
		{name: "compares ordered with all three others in every operation", n: 4096, count: 12},
		{
			name: "drops scanning range queries and builds for large scenarios", n: 1 << 20, count: 7,
			skip: []pair{{"valuesBetween", ordered, hashed}, {"valuesBetween", ordered, mapSets}, {"build", ordered, btreeSets}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairsFor(tt.n, ops, 1<<16, 1<<16)
			if len(got) != tt.count {
				t.Errorf("%d pairs, want %d: %v", len(got), tt.count, got)
			}
			for _, p := range got {
				if p.a != ordered || p.b == ordered {
					t.Errorf("not ordered against another candidate: %v", p)
				}
			}
			for _, p := range tt.skip {
				if slices.Contains(got, p) {
					t.Errorf("unexpected %v", p)
				}
			}
		})
	}
}

// TestWorkloads makes sure the index workloads behave like a database index
// and never ask a multimap for something impossible. It covers the build and
// churn workloads of the benchmark driver: every insertion adds a value the
// key does not hold, every deletion removes a value inserted earlier,
// insertions and deletions interleave, the ratio of insertions to final
// values is as requested, and the multimap ends up holding exactly the corpus.
func TestWorkloads(t *testing.T) {
	const n = 3000
	vals, offs := keys.Values(n, 0xFA11)
	corpus := map[kv]bool{}
	for i := range n {
		for _, v := range vals[offs[i]:offs[i+1]] {
			corpus[kv{uint32(i), v}] = true
		}
	}
	for _, r := range []float64{1, 1.5, 2, 3} {
		t.Run(fmt.Sprintf("build with ratio %v", r), func(t *testing.T) {
			checkWorkload(t, buildWorkload(n, vals, offs, r, 1), map[kv]bool{}, corpus, r*float64(len(vals)))
		})
		t.Run(fmt.Sprintf("churn with ratio %v", r), func(t *testing.T) {
			checkWorkload(t, churnWorkload(n, vals, r, 1), maps.Clone(corpus), corpus, (r-1)*float64(len(vals)))
		})
	}
}

// checkWorkload replays ms on start and checks it against the expectations
// of TestWorkloads.
func checkWorkload(t *testing.T, ms []mutation, start, end map[kv]bool, inserts float64) {
	t.Helper()
	state, adds, switches := start, 0, 0
	for i, m := range ms {
		p := kv{m.key, m.val}
		if m.del != state[p] {
			t.Fatalf("mutation %d: del=%v but present=%v", i, m.del, state[p])
		}
		if m.del {
			delete(state, p)
		} else {
			state[p] = true
			adds++
		}
		if i > 0 && m.del != ms[i-1].del {
			switches++
		}
	}
	if !maps.Equal(state, end) {
		t.Errorf("ends with %d pairs, want the corpus of %d", len(state), len(end))
	}
	if math.Abs(float64(adds)-inserts) > 1 {
		t.Errorf("%d insertions, want %.0f", adds, inserts)
	}
	if inserts > float64(len(end)) && switches < adds/100 {
		t.Errorf("only %d switches between inserting and deleting in %d mutations", switches, len(ms))
	}
}

// TestSpeedup makes sure the summary states speed the way a reader expects:
// "2.00×" when A is twice as fast, "0.50×" when it is half as fast. It covers
// the conversion of rtcompare's relative difference (1 - timeA/timeB) in the
// benchmark driver's summary table.
func TestSpeedup(t *testing.T) {
	for _, tt := range []struct{ delta, want float64 }{{0.5, 2}, {0, 1}, {-1, 0.5}} {
		if got := speedup(tt.delta); math.Abs(got-tt.want) > 1e-12 {
			t.Errorf("speedup(%v) = %v, want %v", tt.delta, got, tt.want)
		}
	}
}

// TestMedian makes sure the summary tables report the typical process, not
// an outlier. It covers the median helper of the benchmark driver for odd
// and even counts and for no values at all.
func TestMedian(t *testing.T) {
	if got := median([]float64{3, 1, 2}); got != 2 {
		t.Errorf("odd: %v", got)
	}
	if got := median([]float64{4, 1, 3, 2}); got != 2.5 {
		t.Errorf("even: %v", got)
	}
	if got := median(nil); !math.IsNaN(got) {
		t.Errorf("empty: %v", got)
	}
}
