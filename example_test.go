package multimap

import (
	"fmt"

	set3 "github.com/TomTonic/Set3"
)

func Example_basicUsage() {
	mm := New[int]()
	// Use FromString to obtain normalized keys from user strings
	mm.AddValue(FromString("Alice"), 1)
	mm.AddValue(FromString("Bob"), 2)

	fmt.Println(mm.NumberOfKeys())
	// Output:
	// 2
}

func Example_rangeQuery() {
	mm := New[int]()
	mm.AddValue(FromString("a"), 1)
	mm.AddValue(FromString("b"), 2)
	mm.AddValue(FromString("c"), 3)

	set := mm.ValuesBetweenInclusive(FromString("a"), FromString("b"))
	// set is a *set3.Set3[int]; for documentation we print whether it contains 1 and 2
	fmt.Println(set.Equals(set3.From(1, 2)))
	// Output:
	// true
}

func Example_iterators() {
	mm := NewOrdered[string]() // unsynchronized; iteration in ascending key order
	mm.AddValue(FromString("b"), "second")
	mm.AddValue(FromString("a"), "first")
	mm.AddValue(FromString("c"), "third")

	for v := range mm.ValuesFromInclusiveSeq(FromString("b")) {
		fmt.Println(v)
	}
	// Output:
	// second
	// third
}
