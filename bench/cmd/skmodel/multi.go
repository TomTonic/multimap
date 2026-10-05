package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/mkpage"
)

// The memory model of the multi-key page (MKSV) for the design note of step 4
// (docs/redesign/step4-mksv-design.md). It builds the tree the way inserts build
// it, which for a tree without deletions does not depend on the order: a subtree
// whose entries all have one value and whose content fits the largest page is
// one multi-key page, hung at the byte that tells its keys from the keys beside
// it; a subtree that does not fit gets a byte node on the next byte in which its
// keys differ, and each child is a subtree in turn (the page "bursts"). A subtree
// with a single entry is a single-key page. An entry with several values is a
// single-key page, and the subtree above it cannot be a multi-key page.

var multiF = flag.Bool("multi", false, "model the multi-key page: the tree with multi-key pages against the tree with single-key pages only")
var mkGridF = flag.String("mkgrid", "128,256,384,512", "size classes of the multi-key page")

type mkModel struct {
	es        []entry
	c         keys.Corpus
	mkClasses []int
	usePages  bool

	nodes, nodeBytes   int
	skPages, skBytes   int
	mkPages, mkBytes   int
	mkEntries, mkSlack int // entries held by multi-key pages; their pages' bytes minus their content
	realBytes, realOff int // bytes of the real pages (internal/mkpage) built for the same entries; pages the model fits and the package does not
	shape              map[int]int
}

// valueBytes returns what the values of entry i take in a page: the bytes of the values and,
// for strings, a length byte each.
func (m *mkModel) valueBytes(i int) int {
	n := 0
	for _, v := range m.es[i].vals {
		switch *valuesF {
		case "words", "pointers":
			n += 8
		case "words-len":
			n += 9
		default:
			n += 1 + valueLen(m.c, v)
		}
	}
	return n
}

// skSize is the content of a single-key page for entry i whose remainder is rem bytes.
func (m *mkModel) skSize(i, rem int) int {
	size := headerBytes + rem
	if *valuesF == "pointers" {
		size = (size + 7) &^ 7
	}
	return size + m.valueBytes(i)
}

// mkSize is the content of a multi-key page for the entries lo..hi-1 below depth d:
// head (type, count, length of the common prefix), a remainder length per entry, the common
// prefix once, then remainder, value length and value of each entry (strings), or all
// remainders, padding to a word and the values (fixed-size values).
func (m *mkModel) mkSize(lo, hi, d int) (size, cp int) {
	first, last := m.es[lo].key, m.es[hi-1].key
	cp = lcp(first[d:], last[d:])
	n := hi - lo
	rems := 0
	vals := 0
	for i := lo; i < hi; i++ {
		rems += len(m.es[i].key) - d - cp
		vals += m.valueBytes(i)
	}
	switch *valuesF {
	case "words", "pointers", "words-len":
		size = 3 + n + cp + rems
		size = (size + 7) &^ 7
		size += vals
	default:
		size = 3 + n + cp + rems + vals // vals counts the value length byte of each entry
	}
	return size, cp
}

func (m *mkModel) build(lo, hi, d int) {
	if hi-lo == 1 {
		m.single(lo, len(m.es[lo].key)-d)
		return
	}
	if m.usePages && m.allSingle(lo, hi) {
		if size, _ := m.mkSize(lo, hi, d); size <= m.mkClasses[len(m.mkClasses)-1] {
			class := classFor(m.mkClasses, size)
			m.mkPages++
			m.mkBytes += class
			m.mkEntries += hi - lo
			m.mkSlack += class - size
			m.shape[min(hi-lo, 30)]++
			m.realPage(lo, hi, d, class)
			return
		}
	}
	first, last := m.es[lo].key, m.es[hi-1].key
	end := d + lcp(first[d:], last[d:])
	children := 0
	if len(first) == end { // the first key ends here: the end page of the node
		m.single(lo, 0)
		children++
		lo++
	}
	for i := lo; i < hi; {
		j := i + 1
		for j < hi && m.es[j].key[end] == m.es[i].key[end] {
			j++
		}
		children++
		m.build(i, j, end+1)
		i = j
	}
	t := &trie{nodes: map[int]int{}}
	t.node(children, end-d)
	m.nodes += t.count
	m.nodeBytes += t.bytes
}

func (m *mkModel) allSingle(lo, hi int) bool {
	for i := lo; i < hi; i++ {
		if len(m.es[i].vals) != 1 {
			return false
		}
	}
	return true
}

func (m *mkModel) single(i, rem int) {
	m.skPages++
	m.skBytes += classFor(pageGrid, min(m.skSize(i, rem), 512))
}

func multiModel(kind keys.Kind, singleValue bool, c keys.Corpus, es []entry) {
	var grid []int
	for _, f := range strings.Split(*mkGridF, ",") {
		var v int
		fmt.Sscan(f, &v)
		grid = append(grid, v)
	}
	sort.Ints(grid)
	name := "natural"
	if singleValue {
		name = "single-value"
	}
	n := float64(len(es))
	fmt.Printf("\n%s, %s, values %s, multi-key page classes %v: the tree of byte nodes with and without multi-key pages\n\n", kind, name, *valuesF, grid)
	fmt.Println("| tree | single-key pages | multi-key pages (entries a page) | nodes | page B/key | node B/key | total B/key |")
	fmt.Println("|---|--:|--:|--:|--:|--:|--:|")
	for _, use := range []bool{false, true} {
		m := &mkModel{es: es, c: c, mkClasses: grid, usePages: use, shape: map[int]int{}}
		m.build(0, len(es), 0)
		label := "single-key pages only"
		if use {
			label = "with multi-key pages"
		}
		per := 0.0
		if m.mkPages > 0 {
			per = float64(m.mkEntries) / float64(m.mkPages)
		}
		fmt.Printf("| %s | %d | %d (%.1f) | %d | %.1f | %.1f | %.1f |\n", label, m.skPages, m.mkPages, per, m.nodes, float64(m.skBytes+m.mkBytes)/n, float64(m.nodeBytes)/n, float64(m.skBytes+m.mkBytes+m.nodeBytes)/n)
		if use && m.mkPages > 0 {
			fmt.Printf("\nentries held by multi-key pages: %.1f %%; slack of those pages: %.1f B/key of all keys; the same pages built with internal/mkpage: %d B (model %d B), %d of them of another size or refused; pages by number of entries (30: 30 and more):", 100*float64(m.mkEntries)/n, float64(m.mkSlack)/n, m.realBytes, m.mkBytes, m.realOff)
			var ks []int
			for k := range m.shape {
				ks = append(ks, k)
			}
			sort.Ints(ks)
			for _, k := range ks {
				fmt.Printf(" %d:%d", k, m.shape[k])
			}
			fmt.Println()
		}
	}
}

// u64Model models the keys of the bench's u64 profile (random 8-byte keys, one value each) for
// the case uint64 -> {*T}, 262,144 keys, values of one word.
func u64Model() {
	const n = 262144
	c := keys.Generate(keys.U64, n, 0x5EED)
	es := make([]entry, len(c.Keys.B))
	for i, k := range c.Keys.B {
		es[i] = entry{k, []uint64{uint64(i) + 1}}
	}
	sort.Slice(es, func(i, j int) bool { return string(es[i].key) < string(es[j].key) })
	multiModel(keys.U64, true, c, es)
}

// realPage builds the page of the model with internal/mkpage, checks every entry in it
// and counts its real size. The model and the package must agree on what fits and on
// the class: a difference is counted in realOff.
func (m *mkModel) realPage(lo, hi, d, class int) {
	rests := make([][]byte, hi-lo)
	for i := lo; i < hi; i++ {
		rests[i-lo] = m.es[i].key[d:]
	}
	var size int
	switch *valuesF {
	case "words", "pointers", "words-len":
		vals := make([]uint64, hi-lo)
		for i := lo; i < hi; i++ {
			vals[i-lo] = m.es[i].vals[0]
		}
		p := mkpage.BuildFixed(rests, vals)
		if p == nil {
			m.realOff++
			return
		}
		for i, r := range rests {
			if v, ok := p.Get[uint64](r); !ok || v != vals[i] {
				panic("mkpage.Fixed lost an entry")
			}
		}
		size = p.Size()
	default:
		vals := make([][]byte, hi-lo)
		for i := lo; i < hi; i++ {
			v := m.es[i].vals[0]
			if v >= 1 && v <= uint64(len(m.c.Names)) {
				vals[i-lo] = []byte(m.c.Names[v-1])
			} else {
				vals[i-lo] = []byte(fmt.Sprintf("%016x", v))
			}
		}
		p := mkpage.BuildStrings(rests, vals)
		if p == nil {
			m.realOff++
			return
		}
		for i, r := range rests {
			if v, ok := p.Get(r); !ok || string(v) != string(vals[i]) {
				panic("mkpage.Page lost an entry")
			}
		}
		size = p.Size()
	}
	if size != class {
		m.realOff++
	}
	m.realBytes += size
}
