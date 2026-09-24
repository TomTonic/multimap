// Package layout perturbs where a benchmark process places its fixtures in
// memory, so that separate processes sample different heap layouts instead of
// repeating one.
//
// Why: for large pointer-heavy fixtures the heap layout of a process (span and
// page placement, cache-set conflicts, TLB reach) shifts a comparison by
// several percentage points, and building fixtures in a fixed order biases
// every process the same way. rtcompare's interval covers only the noise
// within one process. Running each comparison in several processes, each with
// its own -layoutseed, turns layout into a random effect that cmd/summarize
// can average over. Seed 0 leaves the layout alone.
package layout

import (
	"flag"
	"math/rand/v2"
)

var seed = flag.Uint64("layoutseed", 0, "perturb the heap layout with this seed (0: leave it alone)")

var (
	rng  *rand.Rand
	keep []any // spacers stay alive so that later allocations land elsewhere
)

// Seed returns the -layoutseed flag; 0 means no perturbation. Commands record
// it with every result so that summaries can tell processes apart.
func Seed() uint64 { return *seed }

func random() *rand.Rand {
	if rng == nil {
		rng = rand.New(rand.NewPCG(*seed, 0x1a7047))
	}
	return rng
}

// Spacer allocates, for every small object size, a random number of objects
// with and without pointers (Go keeps the two in separate spans), plus one
// pointer-free block of up to 4 MiB, and keeps them all alive. Allocations
// made afterwards start at different offsets within their spans and pages.
// Call it after flag.Parse, before building each fixture. It does nothing when
// the seed is 0. Every call draws new sizes, so successive fixtures are shifted
// differently. The spacers add a few MiB of live heap.
func Spacer() {
	if *seed == 0 {
		return
	}
	r := random()
	for size := 16; size <= 4096; size += 16 {
		for range r.IntN(8) {
			keep = append(keep, make([]*byte, size/8))
		}
		for range r.IntN(8) {
			keep = append(keep, make([]byte, size))
		}
	}
	keep = append(keep, make([]byte, r.IntN(4<<20)))
}

// Shuffle puts xs, typically the fixture builders, into a random order derived
// from the seed. It does nothing when the seed is 0, which keeps the order the
// command was written with.
func Shuffle[T any](xs []T) {
	if *seed == 0 {
		return
	}
	random().Shuffle(len(xs), func(i, j int) { xs[i], xs[j] = xs[j], xs[i] })
}
