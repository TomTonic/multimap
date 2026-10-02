//go:build race

package art

// underRace is true when the tests run with the race detector, which slows the
// tree code down by a factor of about 25: tests that fill big key sets use
// smaller ones, which still drive the tree through every node, page and leaf
// kind.
const underRace = true
