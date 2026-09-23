// Package multimap provides multimaps keyed by Key objects: maps in which a
// key holds a set of values instead of a single one.
//
// # Implementations
//
//   - Ordered is an adaptive radix tree. Range queries touch only the keys in
//     range, and every iteration visits keys in ascending order.
//   - Hashed is a Go map. Point operations on long string keys are faster
//     than in Ordered, but every range query scans all keys, and iteration
//     order is unspecified. Like every Go map it does not shrink its memory
//     after deletions.
//
// Both keep the values of a key in a compact container: up to three values
// inline, then a plain array, then a hash set.
//
// # Concurrency
//
// Ordered and Hashed are not synchronized. Any number of goroutines may read
// concurrently, but a write must not run concurrently with any other access;
// callers provide that exclusion themselves. Synchronized wraps any MultiMap
// with a sync.RWMutex, so that all methods are safe for concurrent use. New
// returns a synchronized Ordered multimap.
//
// # Sets and iterators
//
// Every read method comes in two forms. The plain form returns a *set3.Set3
// or a []Key that is an independent copy: it costs an allocation, but the
// caller may keep and modify it, and on a synchronized multimap it holds no
// lock once returned. The Seq form returns an iterator (iter.Seq) over the
// stored data without copying. It yields the values of each key in turn, so a
// value stored under several keys is yielded once per key, whereas the set
// forms hold it once. On a synchronized multimap, a Seq iterator holds the
// read lock for the whole loop: the loop body must not call any method of the
// same multimap (that can deadlock), and a long loop delays writers. Use the
// set forms when the loop body needs the multimap or runs long.
//
// # Key order
//
// Keys are compared byte-wise lexicographically, see Key.LessThan. Range
// queries and ordering follow that comparison. Mixing key encodings (for
// example UTF-8 strings and encoded integers) may order keys in ways that
// differ from numeric or locale-aware expectations.
package multimap

import (
	"iter"

	set3 "github.com/TomTonic/Set3"
)

// MultiMap is the behavior all multimaps of this package share: a map from
// Keys to sets of values. Implementations clone keys on insertion, so callers
// may reuse their Key buffers.
//
// Whether the methods are safe for concurrent use depends on the
// implementation; see the package documentation.
type MultiMap[T comparable] interface {

	// AddValue adds value to the set at key. If the key does not exist in the
	// MultiMap it is added. The provided Key is cloned before insertion;
	// changes to the caller's Key afterwards do not affect the stored key.
	AddValue(key Key, value T)

	// ContainsKey reports whether the MultiMap contains the key.
	ContainsKey(key Key) bool

	// ValuesFor returns the set of values associated with key, or an empty set
	// if the key does not exist. The result is never nil and is an independent
	// copy that the caller may modify.
	ValuesFor(key Key) *set3.Set3[T]

	// ValuesForSeq iterates over the values associated with key, in
	// unspecified order. It yields nothing if the key does not exist.
	ValuesForSeq(key Key) iter.Seq[T]

	// ValuesBetweenInclusive returns a set with all values whose keys are
	// between from and to, including values stored for from and to. The
	// bounds need not exist as keys. If from is greater than to, the result is
	// empty. The result is never nil and is an independent copy.
	ValuesBetweenInclusive(from, to Key) *set3.Set3[T]

	// ValuesBetweenInclusiveSeq iterates over the values of the keys that
	// ValuesBetweenInclusive covers, key by key.
	ValuesBetweenInclusiveSeq(from, to Key) iter.Seq[T]

	// ValuesBetweenExclusive returns a set with all values whose keys are
	// strictly between from and to. The bounds need not exist as keys. The
	// result is never nil and is an independent copy.
	ValuesBetweenExclusive(from, to Key) *set3.Set3[T]

	// ValuesBetweenExclusiveSeq iterates over the values of the keys that
	// ValuesBetweenExclusive covers, key by key.
	ValuesBetweenExclusiveSeq(from, to Key) iter.Seq[T]

	// ValuesFromInclusive returns a set with all values whose keys are greater
	// than or equal to from. The result is never nil and is an independent
	// copy.
	ValuesFromInclusive(from Key) *set3.Set3[T]

	// ValuesFromInclusiveSeq iterates over the values of the keys that
	// ValuesFromInclusive covers, key by key.
	ValuesFromInclusiveSeq(from Key) iter.Seq[T]

	// ValuesFromExclusive returns a set with all values whose keys are
	// strictly greater than from. The result is never nil and is an
	// independent copy.
	ValuesFromExclusive(from Key) *set3.Set3[T]

	// ValuesFromExclusiveSeq iterates over the values of the keys that
	// ValuesFromExclusive covers, key by key.
	ValuesFromExclusiveSeq(from Key) iter.Seq[T]

	// ValuesToInclusive returns a set with all values whose keys are less than
	// or equal to to. The result is never nil and is an independent copy.
	ValuesToInclusive(to Key) *set3.Set3[T]

	// ValuesToInclusiveSeq iterates over the values of the keys that
	// ValuesToInclusive covers, key by key.
	ValuesToInclusiveSeq(to Key) iter.Seq[T]

	// ValuesToExclusive returns a set with all values whose keys are strictly
	// less than to. The result is never nil and is an independent copy.
	ValuesToExclusive(to Key) *set3.Set3[T]

	// ValuesToExclusiveSeq iterates over the values of the keys that
	// ValuesToExclusive covers, key by key.
	ValuesToExclusiveSeq(to Key) iter.Seq[T]

	// AllValues returns a set with all values currently stored. The result is
	// never nil and is an independent copy.
	AllValues() *set3.Set3[T]

	// AllValuesSeq iterates over all values, key by key.
	AllValuesSeq() iter.Seq[T]

	// NumberOfKeys returns the number of keys currently stored. A key whose
	// last value is removed is removed as well.
	NumberOfKeys() uint64

	// AllKeys returns a slice with clones of all keys. Ordered returns them in
	// ascending order; for other implementations the order is unspecified.
	AllKeys() []Key

	// AllKeysSeq iterates over all keys, in the same order as AllKeys. The
	// yielded Keys are views into the multimap: they must not be modified or
	// retained after the loop step (clone them to keep them).
	AllKeysSeq() iter.Seq[Key]

	// RemoveValue removes value from the set at key. Removing a non-existent
	// key or value is a no-op. When the set becomes empty, the key is removed.
	RemoveValue(key Key, value T)

	// RemoveKey removes the key and all its values. Removing a non-existent
	// key is a no-op.
	RemoveKey(key Key)

	// Clear removes all keys and values.
	Clear()
}

// New returns a MultiMap that is safe for concurrent use: an Ordered multimap
// wrapped by Synchronized. Use NewOrdered or NewHashed directly when the
// caller already serializes access or needs no locking.
func New[T comparable]() MultiMap[T] { return Synchronized[T](NewOrdered[T]()) }

// collect gathers the values an iterator yields into a new set.
func collect[T comparable](seq iter.Seq[T]) *set3.Set3[T] {
	s := set3.Empty[T]()
	for v := range seq {
		s.Add(v)
	}
	return s
}
