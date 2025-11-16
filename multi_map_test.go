package multimap

import (
	"testing"

	set3 "github.com/TomTonic/Set3"
)

func TestPutSizeAndContains(t *testing.T) {
	mm := New[int]()
	if mm.Size() != 0 {
		t.Fatalf("new map should be empty")
	}

	mm.PutValue(FromString("k1"), 1)
	if mm.Size() != 1 {
		t.Fatalf("expected size 1, got %d", mm.Size())
	}
	if !mm.ContainsKey(FromString("k1")) {
		t.Fatalf("expected ContainsKey(k1) true")
	}

	// putting another value for same key must not increase Size
	mm.PutValue(FromString("k1"), 2)
	if mm.Size() != 1 {
		t.Fatalf("expected size still 1 after adding second value to same key, got %d", mm.Size())
	}

	// add another key
	mm.PutValue(FromString("k2"), 3)
	if mm.Size() != 2 {
		t.Fatalf("expected size 2 after adding k2, got %d", mm.Size())
	}
}

func TestKeysAndRemoveKey(t *testing.T) {
	mm := New[string]()
	mm.PutValue(FromString("a"), "v1")
	mm.PutValue(FromString("b"), "v2")

	keys := mm.Keys()
	if len(keys) != int(mm.Size()) {
		t.Fatalf("Keys length %d does not match Size %d", len(keys), mm.Size())
	}

	// remove a key
	mm.RemoveKey(FromString("a"))
	if mm.ContainsKey(FromString("a")) {
		t.Fatalf("expected a to be removed")
	}
	if mm.Size() != 1 {
		t.Fatalf("expected size 1 after removing a, got %d", mm.Size())
	}
}

func TestClear(t *testing.T) {
	mm := New[int]()
	mm.PutValue(FromString("x"), 1)
	mm.PutValue(FromString("y"), 2)
	if mm.Size() == 0 {
		t.Fatalf("expected non-empty before Clear")
	}
	mm.Clear()
	if mm.Size() != 0 {
		t.Fatalf("expected size 0 after Clear, got %d", mm.Size())
	}
	if len(mm.Keys()) != 0 {
		t.Fatalf("expected no keys after Clear")
	}
}

func TestRangeQueryDoesNotPanic(t *testing.T) {
	mm := New[int]()
	mm.PutValue(FromString("a"), 1)
	mm.PutValue(FromString("b"), 2)
	mm.PutValue(FromString("c"), 3)

	// Basic sanity: these should not panic and return a non-nil set pointer
	if mm.GetValuesBetweenInclusive(FromString("a"), FromString("b")) == nil {
		t.Fatalf("GetValuesBetweenInclusive returned nil")
	}
	if mm.GetValuesBetweenExclusive(FromString("a"), FromString("c")) == nil {
		t.Fatalf("GetValuesBetweenExclusive returned nil")
	}
	if mm.GetValuesFromInclusive(FromString("b")) == nil {
		t.Fatalf("GetValuesFromInclusive returned nil")
	}
	if mm.GetValuesToExclusive(FromString("b")) == nil {
		t.Fatalf("GetValuesToExclusive returned nil")
	}
}

func TestRangeQueriesReturnExpectedSets(t *testing.T) {
	mm := New[int]()
	mm.PutValue(FromString("a"), 1)
	mm.PutValue(FromString("b"), 2)
	mm.PutValue(FromString("c"), 3)
	mm.PutValue(FromString("d"), 4)

	// a..c inclusive => 1,2,3
	res := mm.GetValuesBetweenInclusive(FromString("a"), FromString("c"))
	want := set3.From(1, 2, 3)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(a,c) returned unexpected set")
	}

	// a..c exclusive => only b => 2
	res = mm.GetValuesBetweenExclusive(FromString("a"), FromString("c"))
	want = set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("BetweenExclusive(a,c) returned unexpected set")
	}

	// from b inclusive => b,c,d => 2,3,4
	res = mm.GetValuesFromInclusive(FromString("b"))
	want = set3.From(2, 3, 4)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(b) returned unexpected set")
	}

	// to c inclusive => a,b,c => 1,2,3
	res = mm.GetValuesToInclusive(FromString("c"))
	want = set3.From(1, 2, 3)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(c) returned unexpected set")
	}

	// from b exclusive => c,d => 3,4
	res = mm.GetValuesFromExclusive(FromString("b"))
	want = set3.From(3, 4)
	if !res.Equals(want) {
		t.Fatalf("FromExclusive(b) returned unexpected set")
	}

	// to c exclusive => a,b => 1,2
	res = mm.GetValuesToExclusive(FromString("c"))
	want = set3.From(1, 2)
	if !res.Equals(want) {
		t.Fatalf("ToExclusive(c) returned unexpected set")
	}
}

func TestRangeWithNonexistentBoundaries(t *testing.T) {
	// map has keys b, d, f
	mm := New[int]()
	mm.PutValue(FromString("b"), 2)
	mm.PutValue(FromString("d"), 4)
	mm.PutValue(FromString("f"), 6)

	// query from 'c' to 'e' (neither endpoint exists) -> should include only 'd'
	res := mm.GetValuesBetweenInclusive(FromString("c"), FromString("e"))
	want := set3.From(4)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(c,e) = unexpected set")
	}

	// exclusive between 'c' and 'f' -> should include d only (f excluded)
	res = mm.GetValuesBetweenExclusive(FromString("c"), FromString("f"))
	want = set3.From(4)
	if !res.Equals(want) {
		t.Fatalf("BetweenExclusive(c,f) = unexpected set")
	}

	// from 'a' inclusive -> includes b,d,f
	res = mm.GetValuesFromInclusive(FromString("a"))
	want = set3.From(2, 4, 6)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(a) = unexpected set")
	}

	// to 'e' inclusive -> includes b,d
	res = mm.GetValuesToInclusive(FromString("e"))
	want = set3.From(2, 4)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(e) = unexpected set")
	}

	// to 'a' inclusive -> empty set (no keys <= 'a')
	res = mm.GetValuesToInclusive(FromString("a"))
	want = set3.Empty[int]()
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(a) expected empty set")
	}

	// from 'z' inclusive -> empty set (no keys >= 'z')
	res = mm.GetValuesFromInclusive(FromString("z"))
	want = set3.Empty[int]()
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(z) expected empty set")
	}
}

func TestRemoveValueAndGetValuesForClone(t *testing.T) {
	mm := New[int]()
	k := FromString("key")
	mm.PutValue(k, 1)
	mm.PutValue(k, 2)

	// remove one value
	mm.RemoveValue(k, 1)
	res := mm.GetValuesFor(k)
	want := set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("after RemoveValue expected {2}, got unexpected set")
	}

	// returned set is a clone: modifying it should not affect stored data
	res.Add(999)
	res2 := mm.GetValuesFor(k)
	if res2.Equals(set3.From(2, 999)) {
		t.Fatalf("modifying returned set should not affect stored set")
	}

	// removing a non-existent value should be no-op
	mm.RemoveValue(k, 42)
	if !mm.GetValuesFor(k).Equals(want) {
		t.Fatalf("RemoveValue non-existent mutated set")
	}
}

func TestGetAllValuesAggregates(t *testing.T) {
	mm := New[int]()
	mm.PutValue(FromString("a"), 1)
	mm.PutValue(FromString("b"), 2)
	mm.PutValue(FromString("a"), 3)

	all := mm.GetAllValues()
	want := set3.From(1, 2, 3)
	if !all.Equals(want) {
		t.Fatalf("GetAllValues expected %v, got unexpected set", "{1,2,3}")
	}
}

func TestPutClonesKey(t *testing.T) {
	mm := New[int]()
	k := Key([]byte{0x61})
	mm.PutValue(k, 7)
	// mutate original key
	k[0] = 0x62
	keys := mm.Keys()
	if len(keys) != 1 {
		t.Fatalf("expected one key")
	}
	if keys[0].Bytes()[0] != 0x61 {
		t.Fatalf("stored key was mutated when original key changed")
	}
}

func TestConcurrentPuts(t *testing.T) {
	mm := New[int]()
	done := make(chan struct{})
	// spawn writers
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 100; j++ {
				mm.PutValue(FromString("k"), i*100+j)
			}
			done <- struct{}{}
		}(i)
	}
	// wait
	for i := 0; i < 10; i++ {
		<-done
	}
	// ensure no panic and some values present
	if mm.Size() == 0 {
		t.Fatalf("expected non-empty after concurrent puts")
	}
}

func TestRangeQueriesWithNegativeInts(t *testing.T) {
	mm := New[int]()
	mm.PutValue(FromInt64(-3), -3)
	mm.PutValue(FromInt64(-1), -1)
	mm.PutValue(FromInt64(0), 0)
	mm.PutValue(FromInt64(2), 2)

	// between -2 and 1 inclusive => should include -1 and 0
	res := mm.GetValuesBetweenInclusive(FromInt64(-2), FromUint64(1))
	want := set3.From(-1, 0)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(-2,1) expected %v got %v", want, res)
	}

	// to 0 inclusive => -3, -1, 0
	res = mm.GetValuesToInclusive(FromInt64(0))
	want = set3.From(-3, -1, 0)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(int64(0)) expected %v got %v", want, res)
	}

	// to 0 inclusive => -3, -1, 0
	res = mm.GetValuesToInclusive(FromUint64(0))
	want = set3.From(-3, -1, 0)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(uint64(0)) expected %v got %v", want, res)
	}

	// from 0 exclusive => 2
	res = mm.GetValuesFromExclusive(FromInt64(0))
	want = set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("FromExclusive(0) expected %v got %v", want, res)
	}

	// from -4 inclusive => all values
	res = mm.GetValuesFromInclusive(FromInt64(-4))
	want = set3.From(-3, -1, 0, 2)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(-4) expected %v got %v", want, res)
	}
}
