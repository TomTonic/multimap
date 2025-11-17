package multimap

import (
	"testing"

	set3 "github.com/TomTonic/Set3"
)

func TestPutSizeAndContains(t *testing.T) {
	mm := New[int]()
	if mm.NumberOfKeys() != 0 {
		t.Fatalf("new map should be empty")
	}

	mm.AddValue(FromString("k1"), 1)
	if mm.NumberOfKeys() != 1 {
		t.Fatalf("expected size 1, got %d", mm.NumberOfKeys())
	}
	if !mm.ContainsKey(FromString("k1")) {
		t.Fatalf("expected ContainsKey(k1) true")
	}

	// putting another value for same key must not increase Size
	mm.AddValue(FromString("k1"), 2)
	if mm.NumberOfKeys() != 1 {
		t.Fatalf("expected size still 1 after adding second value to same key, got %d", mm.NumberOfKeys())
	}

	// add another key
	mm.AddValue(FromString("k2"), 3)
	if mm.NumberOfKeys() != 2 {
		t.Fatalf("expected size 2 after adding k2, got %d", mm.NumberOfKeys())
	}
}

func TestKeysAndRemoveKey(t *testing.T) {
	mm := New[string]()
	mm.AddValue(FromString("a"), "v1")
	mm.AddValue(FromString("b"), "v2")

	keys := mm.AllKeys()
	if len(keys) != int(mm.NumberOfKeys()) {
		t.Fatalf("Keys length %d does not match Size %d", len(keys), mm.NumberOfKeys())
	}

	// remove a key
	mm.RemoveKey(FromString("a"))
	if mm.ContainsKey(FromString("a")) {
		t.Fatalf("expected a to be removed")
	}
	if mm.NumberOfKeys() != 1 {
		t.Fatalf("expected size 1 after removing a, got %d", mm.NumberOfKeys())
	}
}

func TestClear(t *testing.T) {
	mm := New[int]()
	mm.AddValue(FromString("x"), 1)
	mm.AddValue(FromString("y"), 2)
	if mm.NumberOfKeys() == 0 {
		t.Fatalf("expected non-empty before Clear")
	}
	mm.Clear()
	if mm.NumberOfKeys() != 0 {
		t.Fatalf("expected size 0 after Clear, got %d", mm.NumberOfKeys())
	}
	if len(mm.AllKeys()) != 0 {
		t.Fatalf("expected no keys after Clear")
	}
}

func TestRangeQueryDoesNotPanic(t *testing.T) {
	mm := New[int]()
	mm.AddValue(FromString("a"), 1)
	mm.AddValue(FromString("b"), 2)
	mm.AddValue(FromString("c"), 3)

	// Basic sanity: these should not panic and return a non-nil set pointer
	if mm.ValuesBetweenInclusive(FromString("a"), FromString("b")) == nil {
		t.Fatalf("ValuesBetweenInclusive returned nil")
	}
	if mm.ValuesBetweenExclusive(FromString("a"), FromString("c")) == nil {
		t.Fatalf("ValuesBetweenExclusive returned nil")
	}
	if mm.ValuesFromInclusive(FromString("b")) == nil {
		t.Fatalf("ValuesFromInclusive returned nil")
	}
	if mm.ValuesToExclusive(FromString("b")) == nil {
		t.Fatalf("ValuesToExclusive returned nil")
	}
}

func TestRangeQueriesReturnExpectedSets(t *testing.T) {
	mm := New[int]()
	mm.AddValue(FromString("a"), 1)
	mm.AddValue(FromString("b"), 2)
	mm.AddValue(FromString("c"), 3)
	mm.AddValue(FromString("d"), 4)

	// a..c inclusive => 1,2,3
	res := mm.ValuesBetweenInclusive(FromString("a"), FromString("c"))
	want := set3.From(1, 2, 3)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(a,c) returned unexpected set")
	}

	// a..c exclusive => only b => 2
	res = mm.ValuesBetweenExclusive(FromString("a"), FromString("c"))
	want = set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("BetweenExclusive(a,c) returned unexpected set")
	}

	// from b inclusive => b,c,d => 2,3,4
	res = mm.ValuesFromInclusive(FromString("b"))
	want = set3.From(2, 3, 4)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(b) returned unexpected set")
	}

	// to c inclusive => a,b,c => 1,2,3
	res = mm.ValuesToInclusive(FromString("c"))
	want = set3.From(1, 2, 3)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(c) returned unexpected set")
	}

	// from b exclusive => c,d => 3,4
	res = mm.ValuesFromExclusive(FromString("b"))
	want = set3.From(3, 4)
	if !res.Equals(want) {
		t.Fatalf("FromExclusive(b) returned unexpected set")
	}

	// to c exclusive => a,b => 1,2
	res = mm.ValuesToExclusive(FromString("c"))
	want = set3.From(1, 2)
	if !res.Equals(want) {
		t.Fatalf("ToExclusive(c) returned unexpected set")
	}
}

func TestRangeWithNonexistentBoundaries(t *testing.T) {
	// map has keys b, d, f
	mm := New[int]()
	mm.AddValue(FromString("b"), 2)
	mm.AddValue(FromString("d"), 4)
	mm.AddValue(FromString("f"), 6)

	// query from 'c' to 'e' (neither endpoint exists) -> should include only 'd'
	res := mm.ValuesBetweenInclusive(FromString("c"), FromString("e"))
	want := set3.From(4)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(c,e) = unexpected set")
	}

	// exclusive between 'c' and 'f' -> should include d only (f excluded)
	res = mm.ValuesBetweenExclusive(FromString("c"), FromString("f"))
	want = set3.From(4)
	if !res.Equals(want) {
		t.Fatalf("BetweenExclusive(c,f) = unexpected set")
	}

	// from 'a' inclusive -> includes b,d,f
	res = mm.ValuesFromInclusive(FromString("a"))
	want = set3.From(2, 4, 6)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(a) = unexpected set")
	}

	// to 'e' inclusive -> includes b,d
	res = mm.ValuesToInclusive(FromString("e"))
	want = set3.From(2, 4)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(e) = unexpected set")
	}

	// to 'a' inclusive -> empty set (no keys <= 'a')
	res = mm.ValuesToInclusive(FromString("a"))
	want = set3.Empty[int]()
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(a) expected empty set")
	}

	// from 'z' inclusive -> empty set (no keys >= 'z')
	res = mm.ValuesFromInclusive(FromString("z"))
	want = set3.Empty[int]()
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(z) expected empty set")
	}
}

func TestRemoveValueAndValuesForClone(t *testing.T) {
	mm := New[int]()
	k := FromString("key")
	mm.AddValue(k, 1)
	mm.AddValue(k, 2)

	// remove one value
	mm.RemoveValue(k, 1)
	res := mm.ValuesFor(k)
	want := set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("after RemoveValue expected {2}, got unexpected set")
	}

	// returned set is a clone: modifying it should not affect stored data
	res.Add(999)
	res2 := mm.ValuesFor(k)
	if res2.Equals(set3.From(2, 999)) {
		t.Fatalf("modifying returned set should not affect stored set")
	}

	// removing a non-existent value should be no-op
	mm.RemoveValue(k, 42)
	if !mm.ValuesFor(k).Equals(want) {
		t.Fatalf("RemoveValue non-existent mutated set")
	}
}

func TestGetAllValuesAggregates(t *testing.T) {
	mm := New[int]()
	mm.AddValue(FromString("a"), 1)
	mm.AddValue(FromString("b"), 2)
	mm.AddValue(FromString("a"), 3)

	all := mm.AllValues()
	want := set3.From(1, 2, 3)
	if !all.Equals(want) {
		t.Fatalf("GetAllValues expected %v, got unexpected set", "{1,2,3}")
	}
}

func TestPutClonesKey(t *testing.T) {
	mm := New[int]()
	k := Key([]byte{0x61})
	mm.AddValue(k, 7)
	// mutate original key
	k[0] = 0x62
	keys := mm.AllKeys()
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
				mm.AddValue(FromString("k"), i*100+j)
			}
			done <- struct{}{}
		}(i)
	}
	// wait
	for i := 0; i < 10; i++ {
		<-done
	}
	// ensure no panic and some values present
	if mm.NumberOfKeys() == 0 {
		t.Fatalf("expected non-empty after concurrent puts")
	}
}

func TestRangeQueriesWithNegativeInts(t *testing.T) {
	mm := New[int]()
	mm.AddValue(FromInt64(-3), -3)
	mm.AddValue(FromInt64(-1), -1)
	mm.AddValue(FromInt64(0), 0)
	mm.AddValue(FromInt64(2), 2)

	// between -2 and 1 inclusive => should include -1 and 0
	res := mm.ValuesBetweenInclusive(FromInt64(-2), FromUint64(1))
	want := set3.From(-1, 0)
	if !res.Equals(want) {
		t.Fatalf("BetweenInclusive(-2,1) expected %v got %v", want, res)
	}

	// to 0 inclusive => -3, -1, 0
	res = mm.ValuesToInclusive(FromInt64(0))
	want = set3.From(-3, -1, 0)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(int64(0)) expected %v got %v", want, res)
	}

	// to 0 inclusive => -3, -1, 0
	res = mm.ValuesToInclusive(FromUint64(0))
	want = set3.From(-3, -1, 0)
	if !res.Equals(want) {
		t.Fatalf("ToInclusive(uint64(0)) expected %v got %v", want, res)
	}

	// from 0 exclusive => 2
	res = mm.ValuesFromExclusive(FromInt64(0))
	want = set3.From(2)
	if !res.Equals(want) {
		t.Fatalf("FromExclusive(0) expected %v got %v", want, res)
	}

	// from -4 inclusive => all values
	res = mm.ValuesFromInclusive(FromInt64(-4))
	want = set3.From(-3, -1, 0, 2)
	if !res.Equals(want) {
		t.Fatalf("FromInclusive(-4) expected %v got %v", want, res)
	}
}
