package page

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"unsafe"
)

// frontPage returns a copy of the many-key page of uint64 values p with its values moved right behind the
// byte area (the first multiple of 8 behind the remainders) instead of at the end of the object: the other
// layout of step 5.5b, to read.
func frontPage(p *Fixed) (*Fixed, int) {
	m := p.mem()
	n := p.Len()
	vs := (p.Used() + 7) &^ 7
	q := (*Fixed)(allocRaw(p.class()))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(q)), Header), m[:Header]) // the head first: mem needs the class
	copy(q.mem(), m)
	qm := q.mem()
	clear(qm[p.Used():])
	copy(qm[vs:], m[len(m)-n*8:])
	return q, vs
}

// getFront is Get of a page whose values start at vs.
func getFront(p *Fixed, vs int, rest []byte) (uint64, bool) {
	m := p.mem()
	n, l := int(p.currentValues), p.cpl()
	if lcp(m[Header:Header+l], rest) < l {
		return 0, false
	}
	pos, _, found := find(m, Header+l, n, Header+l+2*n, rest[l:])
	if !found {
		return 0, false
	}
	return *(*uint64)(unsafe.Pointer(&m[vs+pos*8])), true
}

// BenchmarkPlacement measures the one question of the placement of the values (docs/redesign/step5-one-page.md,
// decision 3): a Get hit in a page whose values are at the end of the object against a page whose values begin
// right behind the byte area, with the pages in the cache (4,096 pages) and out of it (512 Ki pages of
// 128 bytes or more, 64 MiB and more, visited in random order).
func BenchmarkPlacement(b *testing.B) {
	for _, n := range []int{3, 7, 20} {
		for _, cold := range []bool{false, true} {
			count := 4096
			if cold {
				count = 1 << 19
			}
			keys, _ := pageSet(n)
			end := make([]*Fixed, count)
			front := make([]*Fixed, count)
			vss := make([]int, count)
			for i := range end {
				ks := keys[i&4095]
				vals := make([]uint64, len(ks))
				for j := range vals {
					vals[j] = uint64(j)
				}
				end[i] = BuildFixed(ks, vals)
				front[i], vss[i] = frontPage(end[i])
			}
			order := rand.New(rand.NewPCG(3, 4)).Perm(count)
			name := "hot"
			if cold {
				name = "cold"
			}
			b.Run(fmt.Sprintf("n=%d/%s/end", n, name), func(b *testing.B) {
				for i := range b.N {
					j := order[i%count]
					ks := keys[j&4095]
					if _, ok := end[j].Get[uint64](ks[i%len(ks)]); !ok {
						b.Fatal("lost")
					}
				}
			})
			b.Run(fmt.Sprintf("n=%d/%s/front", n, name), func(b *testing.B) {
				for i := range b.N {
					j := order[i%count]
					ks := keys[j&4095]
					if _, ok := getFront(front[j], vss[j], ks[i%len(ks)]); !ok {
						b.Fatal("lost")
					}
				}
			})
		}
	}
}
