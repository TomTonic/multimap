package page

import (
	"testing"
	"unsafe"
)

type rec struct{ id, aux uint64 }

type onePtr struct{ p *rec }

type twoPtrs [2]*rec

// TestSupportedTypes covers which values a Fixed page takes, as a user of the map sees it: small values
// without pointers, and one word that is a pointer; not an empty value, a string, an interface or a
// struct of two words.
//
// The map asks once per map (Map.decide) and decides between a page, a page of pointers and the value
// overflow only; the page is made for the type afterwards without asking again.
//
// Expected: Supported and HoldsPointers answer as listed, for plain numbers, arrays and structs of
// them (also empty arrays and nested ones), and for a pointer, a struct of one pointer and unsafe.Pointer.
func TestSupportedTypes(t *testing.T) {
	type nested struct {
		a [2]uint8
		b struct{ f float32 }
	}
	for _, tc := range []struct {
		name string
		got  bool
		want bool
	}{
		{"uint64 is supported", Supported[uint64](), true},
		{"a struct of 16 bytes without pointers is supported", Supported[[2]uint64](), true},
		{"a struct of arrays and numbers is supported", Supported[nested](), true},
		{"an array of zero length is not (no size)", Supported[[0]int](), false},
		{"a value of 17 bytes is not", Supported[[17]byte](), false},
		{"an empty struct is not", Supported[struct{}](), false},
		{"a pointer is supported", Supported[*rec](), true},
		{"a struct of one pointer is supported", Supported[onePtr](), true},
		{"unsafe.Pointer is supported", Supported[unsafe.Pointer](), true},
		{"a string is not", Supported[string](), false},
		{"a struct of a number and a pointer is not", Supported[struct {
			n int
			p *rec
		}](), false},
		{"a struct of two pointers is not", Supported[twoPtrs](), false},
		{"an interface is not", Supported[any](), false},
		{"an array of a string is not", Supported[[1]string](), false},
	} {
		if tc.got != tc.want {
			t.Errorf("%s: got %v", tc.name, tc.got)
		}
	}
	if HoldsPointers[uint64]() || !HoldsPointers[*rec]() || HoldsPointers[string]() {
		t.Error("HoldsPointers is right for a pointer only")
	}
}

// TestAllocPtrEveryShape covers the typed objects of the pages of pointers: one for every size class
// and every number of words in front of the values, 166 in all.
//
// A map of pointers that changes a page in place relies on the object being of the type that marks
// exactly its first rawWords words as no pointers; a class or a number of words that no page has gets
// no object.
//
// Expected: an object for each of the 166 shapes, none for a class beyond the last, for no word in front
// and for as many words in front as the class has.
func TestAllocPtrEveryShape(t *testing.T) {
	shapes := 0
	for c, w := range [...]int{4, 8, 16, 32, 48, 64} {
		for j := 1; j < w; j++ {
			if allocPtr[*rec](c, j) == nil {
				t.Errorf("no object for class %d with %d words in front", c, j)
			}
			shapes++
		}
		if allocPtr[*rec](c, 0) != nil || allocPtr[*rec](c, w) != nil {
			t.Errorf("an object for class %d with no value or no word in front", c)
		}
	}
	if allocPtr[*rec](6, 1) != nil || shapes != 166 {
		t.Errorf("an object beyond the last class, or %d shapes", shapes)
	}
}
