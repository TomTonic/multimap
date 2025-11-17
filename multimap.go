// Package multimap provides a simple, thread-safe multi-map keyed by Key objects.
// The default implementation is array-based. This file defines the generic
// MultiMap interface and constructors allowing future alternative implementations.
//
// Keys are compared using `Key.LessThan`, which performs a byte-wise lexicographic
// comparison of the underlying `[]byte` representation. Range queries and ordering
// semantics follow that byte-wise comparison; mixing different key encodings
// (for example UTF-8 strings and encoded integers) may yield ordering that is
// counter-intuitive for numeric or locale-aware comparisons.
//
// Concurrency: all methods of any MultiMap implementation are safe for concurrent
// use by multiple goroutines.
package multimap

import (
	set3 "github.com/TomTonic/Set3"
)

// MultiMap defines the behavior of a multi-map from Keys to a set of values.
// Implementations must clone Keys on insertion and return cloned value sets so
// callers cannot mutate internal state. Implementations must be safe for
// concurrent use by multiple goroutines.
type MultiMap[T comparable] interface {

	// AddValue adds value to the set at key. If the key does not exist in the MultiMap it will be
	// added and the according set will be created. The provided Key is cloned before insertion;
	// changes to the caller's Key after calling AddValue will not affect the stored key.
	AddValue(key Key, value T)

	// ContainsKey checks whether the MultiMap contains the specified key.
	ContainsKey(key Key) bool

	// ValuesFor returns the set of values associated with key. If the key does not exist
	// or if it exists but no values are stored for this key, it returns an empty set.
	// This function always returns a non-nil set. The result set is an independent
	// copy and can thus be safely mutated by the caller without affecting the MultiMap.
	ValuesFor(key Key) *set3.Set3[T]

	// ValuesBetweenInclusive returns a set with all values whose keys are between from and to,
	// including values stored for from and to. Comparisons use `Key.LessThan` (byte-wise
	// lexicographic). It is irrelevant whether the from and to keys exist in the MultiMap.
	// If `from` is greater than `to` according to `Key.LessThan`, the result is an empty set.
	// If no key exists between from and to or if no values are stored for these keys, this function
	// returns an empty set. This function always returns a non-nil set. The result set is an
	// independent copy and can thus be safely mutated by the caller.
	ValuesBetweenInclusive(from, to Key) *set3.Set3[T]

	// ValuesBetweenExclusive returns a set with all values whose keys are between from and to,
	// excluding values stored for from and to. Comparisons use `Key.LessThan` (byte-wise
	// lexicographic). It is irrelevant whether the from and to keys exist in the MultiMap.
	// If `from` is greater than `to` the result is an empty set. If no key exists between from
	// and to or if no values are stored for these keys, this function returns an empty set.
	// This function always returns a non-nil set. The result set is an independent copy and
	// can thus be safely mutated by the caller.
	ValuesBetweenExclusive(from, to Key) *set3.Set3[T]

	// ValuesFromInclusive returns a set with all values whose keys are greater than or equal to from.
	// Comparisons use `Key.LessThan`. It is irrelevant whether the from key exists in the MultiMap.
	// If no key exists after from or if no values are stored for these keys, this function returns
	// an empty set. This function always returns a non-nil set. The result set is an independent
	// copy and can thus be safely mutated by the caller.
	ValuesFromInclusive(from Key) *set3.Set3[T]

	// ValuesFromExclusive returns a set with all values whose keys are strictly greater than from.
	// Comparisons use `Key.LessThan`. It is irrelevant whether the from key exists in the MultiMap.
	// If no key exists after from or if no values are stored for these keys, this function returns
	// an empty set. This function always returns a non-nil set. The result set is an independent
	// copy and can thus be safely mutated by the caller.
	ValuesFromExclusive(from Key) *set3.Set3[T]

	// ValuesToInclusive returns a set with all values whose keys are less than or equal to to.
	// Comparisons use `Key.LessThan`. It is irrelevant whether the to key exists in the MultiMap.
	// If no key exists before to or if no values are stored for these keys, this function returns
	// an empty set. This function always returns a non-nil set. The result set is an independent
	// copy and can thus be safely mutated by the caller.
	ValuesToInclusive(to Key) *set3.Set3[T]

	// ValuesToExclusive returns a set with all values whose keys are strictly less than to.
	// Comparisons use `Key.LessThan`. It is irrelevant whether the to key exists in the MultiMap.
	// If no key exists before to or if no values are stored for these keys, this function returns
	// an empty set. This function always returns a non-nil set. The result set is an independent
	// copy and can thus be safely mutated by the caller.
	ValuesToExclusive(to Key) *set3.Set3[T]

	// AllValues returns a set with all values currently stored in the multi-map.
	// If no values are stored, this function returns an empty set. This function always returns a non-nil set.
	// The result set is an independent copy and can thus be safely mutated by the caller.
	AllValues() *set3.Set3[T]

	// NumberOfKeys returns the number of keys currently stored in the map.
	NumberOfKeys() uint64

	// AllKeys returns a slice with all keys currently stored in the map. Returned
	// keys are clones and can be safely mutated by the caller. The order of keys is
	// implementation-defined and should not be relied upon.
	AllKeys() []Key

	// RemoveValue removes value v from the set of values at key. Removing a non-existent
	// key or value is a no-op. If the set becomes empty the key may be removed.
	RemoveValue(key Key, v T)

	// RemoveKey removes the key and its associated set of values from the MultiMap.
	// Removing a non-existent key is a no-op.
	RemoveKey(key Key)

	// Clear removes all keys and values from the MultiMap.
	Clear()
}

// New returns a new MultiMap using the default array-based implementation.
func New[T comparable]() MultiMap[T] { return NewArrayBased[T]() }

// NewArrayBased explicitly constructs a MultiMap backed by the array-based implementation.
func NewArrayBased[T comparable]() MultiMap[T] { return newArrayBased[T]() }
