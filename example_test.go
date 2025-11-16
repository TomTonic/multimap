package multimap

import (
	"fmt"

	set3 "github.com/TomTonic/Set3"
)

func Example_basicUsage() {
	mm := New[int]()
	// Use FromString to obtain normalized keys from user strings
	mm.PutValue(FromString("Alice"), 1)
	mm.PutValue(FromString("Bob"), 2)

	fmt.Println(mm.Size())
	// Output:
	// 2
}

func Example_rangeQuery() {
	mm := New[int]()
	mm.PutValue(FromString("a"), 1)
	mm.PutValue(FromString("b"), 2)
	mm.PutValue(FromString("c"), 3)

	set := mm.GetValuesBetweenInclusive(FromString("a"), FromString("b"))
	// set is a *set3.Set3[int]; for documentation we print whether it contains 1 and 2
	fmt.Println(set.Equals(set3.From(1, 2)))
	// Output:
	// true
}
