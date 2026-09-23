package multimap

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"sync"
	"testing"

	set3 "github.com/TomTonic/Set3"
)

// implementations lists every way to obtain a MultiMap, so that each test
// runs against all of them.
func implementations[T comparable]() []struct {
	name string
	mk   func() MultiMap[T]
	sync bool
} {
	return []struct {
		name string
		mk   func() MultiMap[T]
		sync bool
	}{
		{"New", New[T], true},
		{"Ordered", func() MultiMap[T] { return NewOrdered[T]() }, false},
		{"Hashed", func() MultiMap[T] { return NewHashed[T]() }, false},
		{"SynchronizedHashed", func() MultiMap[T] { return Synchronized[T](NewHashed[T]()) }, true},
	}
}

func forEachImpl[T comparable](t *testing.T, f func(t *testing.T, mm MultiMap[T])) {
	for _, impl := range implementations[T]() {
		t.Run(impl.name, func(t *testing.T) { f(t, impl.mk()) })
	}
}

func forEachSynchronized[T comparable](t *testing.T, f func(t *testing.T, mm MultiMap[T])) {
	for _, impl := range implementations[T]() {
		if impl.sync {
			t.Run(impl.name, func(t *testing.T) { f(t, impl.mk()) })
		}
	}
}

// refMM is the trivially correct multimap the implementations are compared
// with.
type refMM map[string]map[int]bool

func (r refMM) values(ok func(k string) bool) *set3.Set3[int] {
	s := set3.Empty[int]()
	for k, vs := range r {
		if ok(k) {
			for v := range vs {
				s.Add(v)
			}
		}
	}
	return s
}

// perKey returns, in key order, the value sets of the keys ok accepts: what a
// Seq iterator must yield group by group.
func (r refMM) perKey(ok func(k string) bool) [][]int {
	var keys []string
	for k := range r {
		if ok(k) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var out [][]int
	for _, k := range keys {
		var vs []int
		for v := range r[k] {
			vs = append(vs, v)
		}
		slices.Sort(vs)
		out = append(out, vs)
	}
	return out
}

// TestAgainstReference checks, for every implementation, that point reads,
// all range queries (as sets and as iterators) and key listings agree with a
// trivially correct multimap through random adds, value removals and key
// removals. For Ordered it also checks that iterators yield keys in
// ascending order.
func TestAgainstReference(t *testing.T) {
	forEachImpl(t, func(t *testing.T, mm MultiMap[int]) {
		_, ordered := mm.(*Ordered[int])
		ordered = ordered || isOrdered(mm)
		r := rand.New(rand.NewPCG(3, 4))
		ref := refMM{}
		keys := make([]Key, 400)
		for i := range keys {
			keys[i] = FromString(fmt.Sprintf("k/%d/%d", r.IntN(20), r.IntN(40)))
		}
		for round := range 5 {
			for range 2000 {
				k := keys[r.IntN(len(keys))]
				v := r.IntN(12)
				switch op := r.IntN(10); {
				case op < 6 || round == 0:
					mm.AddValue(k, v)
					if ref[string(k)] == nil {
						ref[string(k)] = map[int]bool{}
					}
					ref[string(k)][v] = true
				case op < 9:
					mm.RemoveValue(k, v)
					if vs := ref[string(k)]; vs != nil {
						delete(vs, v)
						if len(vs) == 0 {
							delete(ref, string(k))
						}
					}
				default:
					mm.RemoveKey(k)
					delete(ref, string(k))
				}
			}
			checkAgainst(t, mm, ref, keys, r, ordered)
		}
	})
}

// isOrdered reports whether the synchronized wrapper holds an Ordered map.
func isOrdered(mm MultiMap[int]) bool {
	s, ok := mm.(*synchronized[int])
	if !ok {
		return false
	}
	_, ordered := s.m.(*Ordered[int])
	return ordered
}

func checkAgainst(t *testing.T, mm MultiMap[int], ref refMM, keys []Key, r *rand.Rand, ordered bool) {
	t.Helper()
	if int(mm.NumberOfKeys()) != len(ref) {
		t.Fatalf("NumberOfKeys = %d, want %d", mm.NumberOfKeys(), len(ref))
	}
	for _, k := range keys {
		want := ref.values(func(x string) bool { return x == string(k) })
		if !mm.ValuesFor(k).Equals(want) || !collect(mm.ValuesForSeq(k)).Equals(want) {
			t.Fatalf("ValuesFor(%s) wrong", k)
		}
		if mm.ContainsKey(k) != (ref[string(k)] != nil) {
			t.Fatalf("ContainsKey(%s) wrong", k)
		}
	}
	all := func(string) bool { return true }
	checkQuery(t, "AllValues", mm.AllValues(), mm.AllValuesSeq(), ref, all, ordered)
	for range 50 {
		a, b := keys[r.IntN(len(keys))], keys[r.IntN(len(keys))]
		if r.IntN(3) == 0 {
			a = a[:r.IntN(len(a)+1)] // bounds that are not keys
		}
		ge := func(k string) bool { return k >= string(a) }
		gt := func(k string) bool { return k > string(a) }
		le := func(k string) bool { return k <= string(b) }
		lt := func(k string) bool { return k < string(b) }
		and := func(f, g func(string) bool) func(string) bool { return func(k string) bool { return f(k) && g(k) } }
		checkQuery(t, "BetweenInclusive", mm.ValuesBetweenInclusive(a, b), mm.ValuesBetweenInclusiveSeq(a, b), ref, and(ge, le), ordered)
		checkQuery(t, "BetweenExclusive", mm.ValuesBetweenExclusive(a, b), mm.ValuesBetweenExclusiveSeq(a, b), ref, and(gt, lt), ordered)
		checkQuery(t, "FromInclusive", mm.ValuesFromInclusive(a), mm.ValuesFromInclusiveSeq(a), ref, ge, ordered)
		checkQuery(t, "FromExclusive", mm.ValuesFromExclusive(a), mm.ValuesFromExclusiveSeq(a), ref, gt, ordered)
		checkQuery(t, "ToInclusive", mm.ValuesToInclusive(b), mm.ValuesToInclusiveSeq(b), ref, le, ordered)
		checkQuery(t, "ToExclusive", mm.ValuesToExclusive(b), mm.ValuesToExclusiveSeq(b), ref, lt, ordered)
	}
	gotKeys := mm.AllKeys()
	var seqKeys []string
	for k := range mm.AllKeysSeq() {
		seqKeys = append(seqKeys, string(k))
	}
	var wantKeys []string
	for k := range ref {
		wantKeys = append(wantKeys, k)
	}
	sort.Strings(wantKeys)
	var got []string
	for _, k := range gotKeys {
		got = append(got, string(k))
	}
	if ordered {
		if !slices.Equal(got, wantKeys) || !slices.Equal(seqKeys, wantKeys) {
			t.Fatalf("Ordered keys not in ascending order or wrong")
		}
	}
	slices.Sort(got)
	slices.Sort(seqKeys)
	if !slices.Equal(got, wantKeys) || !slices.Equal(seqKeys, wantKeys) {
		t.Fatalf("AllKeys/AllKeysSeq wrong")
	}
}

// checkQuery compares a set result with the reference and checks that the
// matching iterator yields exactly the per-key value groups (in key order
// for ordered multimaps, in any order otherwise).
func checkQuery(t *testing.T, name string, set *set3.Set3[int], seq func(func(int) bool), ref refMM, ok func(string) bool, ordered bool) {
	t.Helper()
	if !set.Equals(ref.values(ok)) {
		t.Fatalf("%s: set differs from reference", name)
	}
	var yielded []int
	seq(func(v int) bool { yielded = append(yielded, v); return true })
	groups := ref.perKey(ok)
	var flat []int
	for _, g := range groups {
		flat = append(flat, g...)
	}
	if len(yielded) != len(flat) {
		t.Fatalf("%s: iterator yielded %d values, want %d", name, len(yielded), len(flat))
	}
	if ordered {
		i := 0
		for _, g := range groups {
			if !slices.Equal(slices.Sorted(slices.Values(yielded[i:i+len(g)])), g) {
				t.Fatalf("%s: iterator groups out of key order", name)
			}
			i += len(g)
		}
		return
	}
	slices.Sort(yielded)
	slices.Sort(flat)
	if !slices.Equal(yielded, flat) {
		t.Fatalf("%s: iterator yielded wrong values", name)
	}
}

// TestIteratorStopsEarly checks that breaking out of a range loop over any
// iterator stops it (and, for synchronized multimaps, releases the lock, so
// that a write afterwards does not block).
func TestIteratorStopsEarly(t *testing.T) {
	forEachImpl(t, func(t *testing.T, mm MultiMap[int]) {
		for i := range 100 {
			mm.AddValue(FromInt(i), i)
			mm.AddValue(FromInt(i), -i-1)
		}
		n := 0
		for range mm.AllValuesSeq() {
			if n++; n == 5 {
				break
			}
		}
		for range mm.AllKeysSeq() {
			break
		}
		mm.AddValue(FromInt(1000), 1) // would deadlock if a read lock leaked
		if n != 5 || !mm.ContainsKey(FromInt(1000)) {
			t.Fatalf("iteration did not stop cleanly")
		}
	})
}

// TestRemoveLastValueRemovesKey checks that a key disappears once its last
// value is removed, in every implementation.
func TestRemoveLastValueRemovesKey(t *testing.T) {
	forEachImpl(t, func(t *testing.T, mm MultiMap[int]) {
		k := FromString("k")
		mm.AddValue(k, 1)
		mm.AddValue(k, 2)
		mm.RemoveValue(k, 1)
		mm.RemoveValue(k, 2)
		if mm.ContainsKey(k) || mm.NumberOfKeys() != 0 {
			t.Fatalf("key with no values left is still present")
		}
	})
}

// TestSynchronizedConcurrentUse checks that the synchronized multimaps can be
// read (including through iterators) and written from many goroutines at
// once without data races (run with -race) and without losing writes.
func TestSynchronizedConcurrentUse(t *testing.T) {
	forEachSynchronized(t, func(t *testing.T, mm MultiMap[int]) {
		var wg sync.WaitGroup
		for w := range 4 {
			wg.Go(func() {
				for i := range 500 {
					mm.AddValue(FromInt(w*1000+i), i)
				}
			})
			wg.Go(func() {
				for range 50 {
					for k := range mm.AllKeysSeq() {
						if len(k) != 8 {
							t.Errorf("torn key %v", k)
						}
					}
					_ = mm.ValuesBetweenInclusive(FromInt(0), FromInt(2000))
				}
			})
		}
		wg.Wait()
		if mm.NumberOfKeys() != 2000 {
			t.Fatalf("NumberOfKeys = %d after concurrent writes, want 2000", mm.NumberOfKeys())
		}
	})
}

// TestAllKeysReturnsCopies checks that the keys an iterator yields read
// the stored bytes, and that AllKeys returns copies a caller may modify.
func TestAllKeysReturnsCopies(t *testing.T) {
	forEachImpl(t, func(t *testing.T, mm MultiMap[int]) {
		mm.AddValue(FromString("abc"), 1)
		keys := mm.AllKeys()
		keys[0][0] = 'X'
		for k := range mm.AllKeysSeq() {
			if !bytes.Equal(k, []byte("abc")) {
				t.Fatalf("stored key changed through AllKeys result: %q", k)
			}
		}
	})
}
