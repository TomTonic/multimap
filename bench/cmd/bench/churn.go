package main

import (
	"strings"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/rtcompare"
)

// An index workload adds and removes values the way a database index sees
// them: bursts of 1-16 insertions alternate with bursts of deletions of
// values inserted earlier. Every workload is computed before timing starts
// and replayed from an array, so that the timed loop does nothing but read
// the next mutation and apply it.

// kv is one key-value pair of a workload; key indexes the fixture's churn
// keys: the corpus keys first, then extra keys that only the workload uses.
type kv struct {
	key uint32
	val uint64
}

// mutation is one step of a workload: add or remove one value of one key.
type mutation struct {
	val uint64
	key uint32
	del bool
}

// transientBit marks the values a workload inserts and deletes again. They
// are running row IDs with the top bit set; corpus values have it clear, so
// the two never collide and every mutation really changes the multimap.
const transientBit = 1 << 63

// buildWorkload builds the whole corpus from empty while inserting r-1 times
// as many transient values in between and deleting them again: r times as
// many insertions as values in the end. The multimap ends up holding exactly
// the corpus. With unique, no key ever holds more than one value (see
// keyPool).
func buildWorkload(n int, vals []uint64, offs []int, r float64, seed uint64, unique bool) []mutation {
	rng := rtcompare.NewDPRNG(seed)
	ins := make([]kv, 0, int(r*float64(len(vals)))+1)
	for i := range n {
		for _, v := range vals[offs[i]:offs[i+1]] {
			ins = append(ins, kv{uint32(i), v})
		}
	}
	t := transients(int((r-1)*float64(len(vals))), n, &rng)
	ins = append(ins, t...)
	shuffleKV(ins, &rng)
	return interleave(ins, max(1, len(t)/8), newPool(n, unique), &rng)
}

// churnWorkload is one cycle of steady-state churn on a multimap that holds
// the corpus: (r-1) times as many transient values as the corpus holds are
// inserted and deleted again, about an eighth of the corpus's values alive
// at a time. After the cycle the multimap holds exactly the corpus again, so
// the cycle can be replayed endlessly. With unique, no key ever holds more
// than one value (see keyPool).
func churnWorkload(n int, vals []uint64, r float64, seed uint64, unique bool) []mutation {
	rng := rtcompare.NewDPRNG(seed)
	t := transients(int((r-1)*float64(len(vals))), n, &rng)
	return interleave(t, max(1, len(vals)/8), newPool(n, unique), &rng)
}

// transients returns count pairs with fresh values: half of them on random
// corpus keys, whose value sets then grow and shrink, half on n/2 extra keys,
// which appear and disappear as whole keys. A keyPool, if any, overrides
// these keys when the values are inserted.
func transients(count, n int, rng *rtcompare.DPRNG) []kv {
	extra := max(1, n/2)
	out := make([]kv, count)
	for i := range out {
		k := uint32(rng.Uint64() % uint64(n))
		if rng.Uint64()&1 == 1 {
			k = uint32(n) + uint32(rng.Uint64()%uint64(extra))
		}
		out[i] = kv{k, transientBit | uint64(i)}
	}
	return out
}

// keyPool hands out the extra keys of a workload on unique values, the churn
// keys n to 2n-1: a transient value takes a random free key when it is
// inserted and returns it when it is deleted. Every transient insertion thus
// creates a key and every deletion removes one, and no key ever holds two
// values, as in an index on a unique column.
type keyPool struct{ free []uint32 }

// newPool returns the pool of the n extra keys, or nil for a workload whose
// keys may hold several values.
func newPool(n int, unique bool) *keyPool {
	if !unique {
		return nil
	}
	p := &keyPool{free: make([]uint32, n)}
	for i := range p.free {
		p.free[i] = uint32(n + i)
	}
	return p
}

func (p *keyPool) take(rng *rtcompare.DPRNG) uint32 {
	if len(p.free) == 0 {
		panic("keyPool: more transient values alive than extra keys; lower -ratio")
	}
	i := int(rng.Uint64() % uint64(len(p.free)))
	k := p.free[i]
	p.free[i] = p.free[len(p.free)-1]
	p.free = p.free[:len(p.free)-1]
	return k
}

func (p *keyPool) put(k uint32) { p.free = append(p.free, k) }

// interleave turns the insertions ins, in this order, into bursts of 1-16
// insertions, each followed by a burst of deletions drawn uniformly from the
// transient values inserted so far and not yet deleted. A deletion can
// therefore never precede its insertion. Burst sizes keep about target
// transient values alive when all insertions are transient; once ins is
// used up, the remaining transient values are deleted in bursts, too. With a
// pool, transient values take their keys from it (see keyPool).
func interleave(ins []kv, target int, pool *keyPool, rng *rtcompare.DPRNG) []mutation {
	var live []kv
	out := make([]mutation, 0, 2*len(ins))
	del := func(d int) {
		for range d {
			i := int(rng.Uint64() % uint64(len(live)))
			out = append(out, mutation{live[i].val, live[i].key, true})
			if pool != nil {
				pool.put(live[i].key)
			}
			live[i] = live[len(live)-1]
			live = live[:len(live)-1]
		}
	}
	for len(ins) > 0 {
		a := min(1+int(rng.Uint64()%16), len(ins))
		for _, p := range ins[:a] {
			if p.val&transientBit != 0 {
				if pool != nil {
					p.key = pool.take(rng)
				}
				live = append(live, p)
			}
			out = append(out, mutation{p.val, p.key, false})
		}
		ins = ins[a:]
		// On average 8.5 deletions per burst when len(live) == target, the
		// mean number of insertions.
		dmax := len(live) * 17 / target
		del(min(int(rng.Uint64()%uint64(dmax+1)), len(live)))
	}
	for len(live) > 0 {
		del(min(1+int(rng.Uint64()%16), len(live)))
	}
	return out
}

func shuffleKV(s []kv, rng *rtcompare.DPRNG) {
	for i := len(s) - 1; i > 0; i-- {
		j := int(rng.Uint64() % uint64(i+1))
		s[i], s[j] = s[j], s[i]
	}
}

// The apply functions replay mutations on one concrete candidate type, so
// that no interface call distorts the comparison.

func applyOrdered(m *multimap.Ordered[uint64], k [][]byte, ms []mutation) {
	for _, x := range ms {
		if x.del {
			m.RemoveValue(k[x.key], x.val)
		} else {
			m.AddValue(k[x.key], x.val)
		}
	}
}

func applyHashed(m *multimap.Hashed[uint64], k [][]byte, ms []mutation) {
	for _, x := range ms {
		if x.del {
			m.RemoveValue(k[x.key], x.val)
		} else {
			m.AddValue(k[x.key], x.val)
		}
	}
}

func applyBtree(m *btreeMM, k []string, ms []mutation) {
	for _, x := range ms {
		if x.del {
			btreeRemove(m, k[x.key], x.val)
		} else {
			btreeAdd(m, k[x.key], x.val)
		}
	}
}

func applyBtreeMap(m *btreeMap, k []string, ms []mutation) {
	for _, x := range ms {
		if x.del {
			m.Delete(k[x.key])
		} else {
			m.Set(strings.Clone(k[x.key]), x.val)
		}
	}
}

func applyMap(m mapMM, k []string, ms []mutation) {
	for _, x := range ms {
		if x.del {
			mapRemove(m, k[x.key], x.val)
		} else {
			mapAdd(m, k[x.key], x.val)
		}
	}
}
