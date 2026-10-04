package main

import "github.com/TomTonic/multimap/internal/art"

// trie is the byte-node tree of the model: lazy expansion, so a key hangs
// below the first node that tells it from its neighbours, with the common prefix
// of the keys below a node stored in the node (pessimistic common prefixes,
// as in internal/art).
type trie struct {
	keys  [][]byte
	pages func(i, remainder int) // called for every page: key i, remainder length
	nodes map[int]int            // node bytes (Go blocks) by node size class
	count int
	bytes int
}

// build adds the subtree of keys[lo:hi], which agree in their first d bytes.
func (t *trie) build(lo, hi, d int) {
	if hi-lo == 1 {
		t.pages(lo, len(t.keys[lo])-d)
		return
	}
	first, last := t.keys[lo], t.keys[hi-1]
	end := d + lcp(first[d:], last[d:])
	children := 0
	if len(first) == end { // the first key ends here: the end page of the node
		t.pages(lo, 0)
		children++
		lo++
	}
	for i := lo; i < hi; {
		j := i + 1
		for j < hi && t.keys[j][end] == t.keys[i][end] {
			j++
		}
		children++
		t.build(i, j, end+1)
		i = j
	}
	t.node(children, end-d)
}

// node counts a node with the given number of child slots and prefix length,
// as in internal/art: 64 bytes up to 5 slots, 128 up to 12, 256 up to 26, 512
// up to 58, else 2080, plus a tail for a prefix beyond 12 bytes.
func (t *trie) node(slots, prefix int) {
	size := 2080
	switch {
	case slots <= 5:
		size = 64
	case slots <= 12:
		size = 128
	case slots <= 26:
		size = 256
	case slots <= 58:
		size = 512
	}
	switch p := prefix - 12; {
	case p <= 0:
	case p <= 16:
		size += 16
	case p <= 48:
		size += 48
	case p <= 112:
		size += 112
	default:
		size += 16 // a string header; the bytes are a separate object, not counted
	}
	block, _ := art.Block(size, true)
	t.nodes[size]++
	t.count++
	t.bytes += block
}
