// Package multimap provides a simple, thread-safe multi-map keyed by Key objects.
// The default implementation is array-based. This file defines the generic
// MultiMap interface and constructors allowing future alternative implementations.
//
// Concurrency: all exported methods are safe for concurrent use by multiple goroutines.
package multimap

import (
	set3 "github.com/TomTonic/Set3"
)

// MultiMap defines the behavior of a multi-map from Keys to a set of values.
// Implementations must clone Keys on insertion and return cloned value sets so
// callers cannot mutate internal state.
type MultiMap[T comparable] interface {
	PutValue(key Key, v T)
	RemoveValue(key Key, v T)
	ContainsKey(key Key) bool
	RemoveKey(key Key)
	GetValuesFor(key Key) *set3.Set3[T]
	GetAllValues() *set3.Set3[T]
	GetValuesBetweenInclusive(from, to Key) *set3.Set3[T]
	GetValuesBetweenExclusive(from, to Key) *set3.Set3[T]
	GetValuesFromInclusive(from Key) *set3.Set3[T]
	GetValuesToInclusive(to Key) *set3.Set3[T]
	GetValuesFromExclusive(from Key) *set3.Set3[T]
	GetValuesToExclusive(to Key) *set3.Set3[T]
	Size() uint64
	Keys() []Key
	Clear()
}

// New returns a new MultiMap using the default array-based implementation.
func New[T comparable]() MultiMap[T] { return NewArrayBased[T]() }

// NewArrayBased explicitly constructs a MultiMap backed by the array-based implementation.
func NewArrayBased[T comparable]() MultiMap[T] { return newArrayBased[T]() }
