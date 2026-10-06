package page

//go:generate go run ./gen

import (
	"reflect"
	"unsafe"
)

// Which T take a Fixed page, and how the page of a T is allocated: reflection runs once per map (Map.decide) and
// once per new page, never in an operation of the tree.

// fixedKind says how the page of a T is allocated.
type fixedKind int

const (
	fixedNone fixedKind = iota // not a page
	fixedRaw                   // no pointer: an array of words
	fixedPtr                   // one word that is a pointer: a typed object
)

// maxFixedValue is the largest T without a pointer that takes pages; larger
// values would leave too few per class.
const maxFixedValue = 16

// kindOf returns how the page of T is allocated, or fixedNone.
func kindOf[T comparable]() fixedKind {
	var z T
	switch w := unsafe.Sizeof(z); {
	case w == 0:
		return fixedNone
	case pointerFree(reflect.TypeFor[T]()):
		if w <= maxFixedValue {
			return fixedRaw
		}
	case w == 8 && unsafe.Alignof(z) == 8:
		return fixedPtr // 8 bytes with a pointer in them: the pointer is all of it
	}
	return fixedNone
}

// Supported reports whether values of type T go into a Fixed page: T of 1 to 16
// bytes without a pointer, or one word that is a pointer (a pointer, a map, a
// channel, a function, an unsafe.Pointer, a struct or array of one of them).
// Any other T, an interface or a struct with two words, has no page.
func Supported[T comparable]() bool { return kindOf[T]() != fixedNone }

// HoldsPointers reports whether the page of T is an object with pointers: T is
// a word that is a pointer. Such an object is scanned by the garbage
// collector and, in the tree's statistic, an object with pointers.
func HoldsPointers[T comparable]() bool { return kindOf[T]() == fixedPtr }

func pointerFree(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	case reflect.Array:
		return t.Len() == 0 || pointerFree(t.Elem())
	case reflect.Struct:
		for i := range t.NumField() {
			if !pointerFree(t.Field(i).Type) {
				return false
			}
		}
		return true
	}
	return false
}
