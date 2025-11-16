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
