package art

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"slices"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// recPool returns a function that gives the same *rec for the same i, for the tests of maps of
// pointers: a value is the pointer, and its identity is what the map compares.
func recPool() func(i int) *rec {
	pool := map[int]*rec{}
	return func(i int) *rec {
		if pool[i] == nil {
			pool[i] = &rec{id: uint64(i)}
		}
		return pool[i]
	}
}

// pageValues returns how many values the multi-key pages of m hold in all.
func pageValues[T comparable](m *Map[T]) (values int) {
	m.Objects(func(o Object) {
		if o.Label == "multi-key page" {
			values += o.Values
		}
	})
	return values
}

// holds fails the test unless key has exactly the values want, in any order.
func holds[T comparable](t *testing.T, m *Map[T], key []byte, want []T) {
	t.Helper()
	got := valuesOf(m, key)
	if len(got) != len(want) {
		t.Fatalf("key %q holds %v, want %v", key, got, want)
	}
	for _, v := range want {
		if !slices.Contains(got, v) {
			t.Fatalf("key %q holds %v, want %v", key, got, want)
		}
	}
}

// TestMultiKeyPagesHoldSeveralValues makes sure that keys with a few values share pages
// like keys with one.
//
// A user who maps names to a handful of values each (the typical multimap: most keys have
// one to four values) gets one object for many keys, not a single-key page for each key
// that has a second value (docs/redesign/step5-mkmv-design.md).
//
// Expected, for strings and for uint64 values: after adding 300 keys with one to four values
// each, most keys are in multi-key pages, the pages hold more values than keys, every key has
// exactly its values, the map counts 300 keys, and a scan reports each key once and each
// value once.
func TestMultiKeyPagesHoldSeveralValues(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runSeveralValues(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runSeveralValues(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runSeveralValues(t, recPool()) })
}

func runSeveralValues[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	keys := mkKeys(300)
	want := make([][]T, len(keys))
	total := 0
	for i, k := range keys {
		for j := range 1 + i%4 {
			m.Add(k, val(i*10+j))
			want[i] = append(want[i], val(i*10+j))
			total++
		}
	}
	checkInvariants(t, &m.t)
	if m.Len() != 300 {
		t.Fatalf("Len = %d, want 300", m.Len())
	}
	pages, inPages := pageCount(&m)
	if pages == 0 || inPages < 250 {
		t.Fatalf("%d multi-key pages hold %d of 300 keys", pages, inPages)
	}
	if v := pageValues(&m); v <= inPages {
		t.Fatalf("the pages hold %d values for %d keys, want more", v, inPages)
	}
	for i, k := range keys {
		holds(t, &m, k, want[i])
	}
	nKeys, nValues := 0, 0
	var last []byte
	m.Range(&Bounds{}, func(k []byte) bool {
		if last != nil && string(k) <= string(last) {
			t.Fatalf("scan: key %q after %q", k, last)
		}
		last = append(last[:0], k...)
		nKeys++
		return true
	})
	m.RangeValues(&Bounds{}, func(T) bool { nValues++; return true })
	if nKeys != 300 || nValues != total {
		t.Fatalf("scan: %d keys and %d values, want 300 and %d", nKeys, nValues, total)
	}
}

// TestMultiKeyPageBurstOnFurtherValue makes sure that a further value for a key that a full
// page cannot take bursts the page without counting the key twice.
//
// A user who keeps adding values to one key in an index of many keys must see the same
// number of keys all the time, whichever object holds the values.
//
// Expected: three keys share a page; 150 values added to one of them bring it through the
// multi-key page, the single-key page and the value overflow, with three keys and exactly
// the values added at every step.
func TestMultiKeyPageBurstOnFurtherValue(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		runBurst(t, func(i int) string { return fmt.Sprint("value-", i, strings.Repeat("x", 20)) })
	})
	t.Run("uint64", func(t *testing.T) { runBurst(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runBurst(t, recPool()) })
}

func runBurst[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	a, b, c := []byte("street-a"), []byte("street-b"), []byte("street-c")
	for i, k := range [][]byte{a, b, c} {
		m.Add(k, val(1000+i))
	}
	if pages, keys := pageCount(&m); pages != 1 || keys != 3 {
		t.Fatalf("setup: %d pages hold %d keys, want 1 page with 3", pages, keys)
	}
	var want []T
	for i := range 150 {
		m.Add(b, val(i))
		want = append(want, val(i))
		if m.Len() != 3 {
			t.Fatalf("after %d values: Len = %d, want 3", i+1, m.Len())
		}
		holds(t, &m, b, append([]T{val(1001)}, want...))
		holds(t, &m, a, []T{val(1000)})
		holds(t, &m, c, []T{val(1002)})
	}
	checkInvariants(t, &m.t)
}

// TestMultiKeyPagePairsWithSeveralValues makes sure that a key with several values and a new
// key make a page when they fit one.
//
// A user who adds a second key near a key that already has a few values gets the one object
// for both, as for two keys with one value each.
//
// Expected: a map with one key with three values, then a new key that sorts after it, and
// the same with a new key that sorts before it: one page with two keys and four values, every
// value found.
func TestMultiKeyPagePairsWithSeveralValues(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runPairs(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runPairs(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runPairs(t, recPool()) })
}

func runPairs[T comparable](t *testing.T, val func(i int) T) {
	for _, second := range []string{"street-2", "street-0"} {
		var m Map[T]
		k1, k2 := []byte("street-1"), []byte(second)
		for i := range 3 {
			m.Add(k1, val(i))
		}
		m.Add(k2, val(9))
		if pages, keys := pageCount(&m); pages != 1 || keys != 2 || pageValues(&m) != 4 {
			t.Fatalf("with %q: %d pages hold %d keys and %d values, want 1, 2 and 4", second, pages, keys, pageValues(&m))
		}
		holds(t, &m, k1, []T{val(0), val(1), val(2)})
		holds(t, &m, k2, []T{val(9)})
		checkInvariants(t, &m.t)
	}
}

// TestMultiKeyPageRemovesValuesAndKeys makes sure that removing values and keys from a page
// whose keys have several values counts the keys right and gives the page up when one key is
// left.
//
// A user who removes one value of a key must still see the key; a user who removes the last
// one must see it gone; the keys left in a page of a single key must be found as before.
//
// Expected: in a page of two keys, one with three values, removing two of the values keeps the
// key and the page; removing the third takes the key and leaves a single-key page for the other
// key; RemoveKey of a key with several values takes all of them at once, counts one key, and
// leaves the other key in a single-key page.
func TestMultiKeyPageRemovesValuesAndKeys(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runRemoves(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runRemoves(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runRemoves(t, recPool()) })
}

func runRemoves[T comparable](t *testing.T, val func(i int) T) {
	a, b := []byte("street-a"), []byte("street-b")
	fresh := func() *Map[T] {
		m := new(Map[T])
		for i := range 3 {
			m.Add(a, val(i))
		}
		m.Add(b, val(9))
		if pages, keys := pageCount(m); pages != 1 || keys != 2 {
			t.Fatalf("setup: %d pages hold %d keys, want 1 page with 2", pages, keys)
		}
		return m
	}
	m := fresh()
	m.Remove(a, val(1))
	m.Remove(a, val(7)) // not a value of the key
	m.Remove(a, val(0))
	if pages, keys := pageCount(m); pages != 1 || keys != 2 || m.Len() != 2 {
		t.Fatalf("after two values of three: %d pages hold %d keys, Len %d, want 1 page with 2 keys", pages, keys, m.Len())
	}
	holds(t, m, a, []T{val(2)})
	m.Remove(a, val(2))
	if pages, _ := pageCount(m); pages != 0 || m.Len() != 1 || m.Has(a) {
		t.Fatalf("after the last value: %d pages, Len %d, Has %v, want no page, 1 key, and no key", pages, m.Len(), m.Has(a))
	}
	holds(t, m, b, []T{val(9)})
	checkInvariants(t, &m.t)

	m = fresh()
	m.RemoveKey(a)
	m.RemoveKey(a) // not there any more
	if pages, _ := pageCount(m); pages != 0 || m.Len() != 1 || m.Has(a) {
		t.Fatalf("after RemoveKey: %d pages, Len %d, Has %v, want no page, 1 key, and no key", pages, m.Len(), m.Has(a))
	}
	holds(t, m, b, []T{val(9)})
	checkInvariants(t, &m.t)

	m = fresh()
	m.RemoveKey(b) // the key with one value: the other key keeps its three
	holds(t, m, a, []T{val(0), val(1), val(2)})
	if m.Len() != 1 {
		t.Fatalf("Len = %d, want 1", m.Len())
	}
	checkInvariants(t, &m.t)
}

// TestMultiKeyPageMergesWithSeveralValues makes sure that the merge of the pages below a node
// takes keys with several values along.
//
// A user who deletes keys from an index where most keys have a few values gets the pages back
// together as for keys with one value each, and loses no value in it.
//
// Expected: two sibling pages of six keys with two values each become one page of
// 6+mergeBelow keys, with all their values, when the removals leave mergeBelow keys in one of
// them.
func TestMultiKeyPageMergesWithSeveralValues(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runMerges(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runMerges(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runMerges(t, recPool()) })
}

func runMerges[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	key := func(side byte, i int) []byte { return fmt.Appendf(nil, "%c%02d%s", side, i, strings.Repeat("x", 40)) }
	for i := range 6 {
		for j := range 2 {
			m.Add(key('a', i), val(i*10+j))
			m.Add(key('b', i), val(100+i*10+j))
		}
	}
	if pages, keys := pageCount(&m); pages != 2 || keys != 12 {
		t.Fatalf("setup: %d pages hold %d keys, want 2 pages with 12", pages, keys)
	}
	for i := range 6 - mergeBelow {
		m.RemoveKey(key('a', i))
	}
	pages, keys := pageCount(&m)
	if pages != 1 || keys != 6+mergeBelow || pageValues(&m) != 2*(6+mergeBelow) {
		t.Fatalf("after the merge: %d pages hold %d keys and %d values, want 1 page with %d keys and %d values", pages, keys, pageValues(&m), 6+mergeBelow, 2*(6+mergeBelow))
	}
	for i := range 6 {
		holds(t, &m, key('b', i), []T{val(100 + i*10), val(100 + i*10 + 1)})
	}
	for i := 6 - mergeBelow; i < 6; i++ {
		holds(t, &m, key('a', i), []T{val(i * 10), val(i*10 + 1)})
	}
	checkInvariants(t, &m.t)
}

// TestMultiKeyPagesOfPointersKeepTheirValuesAlive makes sure that the values of a map of pointers
// stay reachable for the garbage collector while the map holds them, in the pages that every
// change of the map makes anew.
//
// A user whose map holds *Record values must never see a record freed that the map still has: a page
// of pointers is a typed object, and the tree copies pages when a key or a value is added or removed.
//
// Expected: after many adds and removes of values of 200 keys (with up to four values each), with a
// collection after every round, no value that the map holds has been finalized, and it still has
// the very objects it was given.
func TestMultiKeyPagesOfPointersKeepTheirValuesAlive(t *testing.T) {
	var dead [1 << 13]atomic.Bool
	var m Map[*rec]
	r := rand.New(rand.NewPCG(3, 4))
	keys := mkKeys(200)
	want := make([][]int, len(keys)) // the ids a key holds: the test keeps no pointer to them
	next := 0
	for round := range 40 {
		for range 150 {
			i := r.IntN(len(keys))
			if r.IntN(3) > 0 || len(want[i]) == 0 {
				v := &rec{id: uint64(next)}
				runtime.SetFinalizer(v, func(v *rec) { dead[v.id].Store(true) })
				m.Add(keys[i], v)
				want[i] = append(want[i], next)
				next++
				continue
			}
			j := r.IntN(len(want[i]))
			id := want[i][j]
			var victim *rec
			m.Each(keys[i], func(v *rec) bool {
				if int(v.id) == id {
					victim = v
				}
				return victim == nil
			})
			m.Remove(keys[i], victim)
			want[i] = slices.Delete(want[i], j, j+1)
		}
		for range 3 {
			runtime.GC()
			time.Sleep(time.Millisecond)
		}
		for i, ids := range want {
			got := 0
			m.Each(keys[i], func(v *rec) bool {
				if dead[v.id].Load() || !slices.Contains(ids, int(v.id)) {
					t.Fatalf("round %d: key %q holds %d, which is freed or not one of %v", round, keys[i], v.id, ids)
				}
				got++
				return true
			})
			if got != len(ids) {
				t.Fatalf("round %d: key %q holds %d values, want %d", round, keys[i], got, len(ids))
			}
		}
	}
}
