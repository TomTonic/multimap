package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/artstr"
	"github.com/TomTonic/multimap/internal/lpage"
)

func try(c keys.Corpus, vals []uint64, offs []int, n int, hdr int) (ok bool, at int) {
	lpage.MaxHeader = hdr
	m := artstr.Map[string]{Pairs: true}
	for i := range n {
		for _, v := range vals[offs[i]:offs[i+1]] {
			s := fmt.Sprintf("%016x", v)
			if c.Names != nil && v >= 1 && v <= uint64(len(c.Names)) {
				s = c.Names[v-1]
			}
			m.Add(c.Keys.B[i], s)
		}
	}
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	for i := 0; i < n; i += 2 {
		at = i
		m.RemoveKey(c.Keys.B[i])
	}
	return true, at
}

func main() {
	hdr, _ := strconv.Atoi(os.Args[1])
	c := keys.Generate(keys.Dirs, 86215, 0x5EED)
	offs := make([]int, 86216)
	var vals []uint64
	for i, vs := range c.Natural[:86215] {
		vals = append(vals, vs...)
		offs[i+1] = len(vals)
	}
	ok, at := try(c, vals, offs, 86215, hdr)
	fmt.Println("full:", ok, at)
	if ok {
		return
	}
	// shrink n by bisection
	lo, hi := 1, 86215
	for lo < hi {
		mid := (lo + hi) / 2
		if ok, _ := try(c, vals, offs, mid, hdr); ok {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	fmt.Println("smallest n that panics:", lo)
}
