package vpage

import (
	"math/rand/v2"
	"testing"
)

// touch reads one byte of each line of the page after its first two, as soon as
// the header says how many there are: the loads do not depend on each other, so
// they overlap. The result feeds a condition that never holds, so that they are
// not optimized away.
func (p *Page) touch() bool {
	m := p.mem()
	x := m[64]
	for off := 128; off < len(m); off += 64 {
		x |= m[off]
	}
	return x == 0xff && m[0] == 0xfe && m[1] == 0xfd
}

// BenchmarkCold separates what a cold lookup waits for: the work around the
// page (finding the key and its page), reading the first line of the page, all
// its lines, and a full Get.
func BenchmarkCold(b *testing.B) {
	w := build(3_000_000, shapes["u64"])
	order := rand.New(rand.NewPCG(7, 8)).Perm(len(w.keys))
	run := func(name string, f func(p *Page, k []byte) uint64) {
		b.Run(name, func(b *testing.B) {
			var acc uint64
			for i := 0; b.Loop(); i++ {
				j := order[i%len(order)]
				acc += f(w.pages[w.page[j]], w.keys[j])
			}
			sink += acc
		})
	}
	run("nothing but the key and the page index", func(p *Page, k []byte) uint64 { return uint64(len(k)) })
	run("pointer of the page", func(p *Page, k []byte) uint64 {
		if p == nil {
			return 1
		}
		return 0
	})
	run("header only", func(p *Page, k []byte) uint64 { return uint64(p.class) })
	run("every line", func(p *Page, k []byte) uint64 {
		if p.touch() {
			return 1
		}
		return 0
	})
	run("Get", func(p *Page, k []byte) uint64 { v, _ := p.Get(k); return v })
}
