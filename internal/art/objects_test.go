package art

import (
	"bytes"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
	"github.com/TomTonic/multimap/internal/vpage"
)

// allocated returns the size of the blocks Go allocates for the objects that mk
// creates: the smallest distance between two of many objects, which neighbours
// in one span have. (The runtime's allocation counters are lumpy, since they
// count a span at a time.)
func allocated(mk func() unsafe.Pointer) int {
	const n = 512
	keep := make([]unsafe.Pointer, n)
	addrs := make([]uintptr, n)
	for i := range keep {
		keep[i] = mk()
		addrs[i] = uintptr(keep[i])
	}
	slices.Sort(addrs)
	gap := int(addrs[n-1] - addrs[0])
	for i := 1; i < n; i++ {
		gap = min(gap, int(addrs[i]-addrs[i-1]))
	}
	runtime.KeepAlive(keep)
	return gap
}

// TestBlock makes sure that the cache-line statistic of the multimap's index
// reads the memory the Go program really uses: for an object of a given size
// it says which block the allocator takes, and where in the block the object
// starts. It belongs to the object statistic of the ART behind multimap.Ordered
// (docs/redesign), which reports how well the objects fill cache lines. The
// table of size classes is a copy of the runtime's, so the test compares it
// with what the runtime of the Go that runs it allocates, for objects with and
// without pointers, at and one step above every class.
func TestBlock(t *testing.T) {
	t.Run("rounds up to the size classes the runtime has", func(t *testing.T) {
		for _, c := range goClasses {
			if c >= 16 && c <= maxMeasured {
				// without pointers, from 16 bytes on (smaller ones share tiny blocks)
				for _, tc := range []struct{ size, want int }{{c, c}, {c + 1, next(c)}} {
					got := allocated(func() unsafe.Pointer {
						return reflect.New(reflect.ArrayOf(tc.size, reflect.TypeFor[byte]())).UnsafePointer()
					})
					if b, off := Block(tc.size, false); got != tc.want || b != tc.want || off != 0 {
						t.Errorf("%d bytes without pointers: runtime %d, Block %d at %d, want %d at 0", tc.size, got, b, off, tc.want)
					}
				}
			}
		}
	})
	t.Run("adds the malloc header to objects with pointers above 512 bytes", func(t *testing.T) {
		for _, c := range goClasses {
			if c > maxMeasured {
				break
			}
			// with pointers: the object of the largest size that fits the class, and the next size up
			fit, over := c, c+8
			if c > maxNoHeader {
				fit = c - mallocHeader
			}
			for _, tc := range []struct{ size, want, off int }{{fit, c, 0}, {over, next(c), 0}} {
				if tc.size > maxNoHeader {
					tc.off = mallocHeader
				}
				got := allocated(func() unsafe.Pointer {
					return reflect.New(reflect.ArrayOf(tc.size/8, reflect.TypeFor[*int]())).UnsafePointer()
				})
				if b, off := Block(tc.size, true); got != tc.want || b != tc.want || off != tc.off {
					t.Errorf("%d bytes with pointers: runtime %d, Block %d at %d, want %d at %d", tc.size, got, b, off, tc.want, tc.off)
				}
			}
		}
	})
	t.Run("rounds big objects up to whole pages", func(t *testing.T) {
		if b, off := Block(8193, false); b != 16384 || off != 0 {
			t.Errorf("8193 bytes: %d at %d, want 16384 at 0", b, off)
		}
	})
}

// maxMeasured is the largest class the test compares with the runtime: bigger
// objects come few to a span, too few for neighbours.
const maxMeasured = 4096

// next returns the size class after c, or the next whole page.
func next(c int) int {
	i := slices.Index(goClasses[:], c)
	if i == len(goClasses)-1 {
		return c + 8192
	}
	return goClasses[i+1]
}

// TestObjectSizes makes sure that the object statistic of the index sees every
// type of object at the size Go really allocates it. It belongs to the object
// statistic of the ART behind multimap.Ordered (docs/redesign), which reports
// per benchmark case how the objects fill cache lines. For every type of node,
// page and leaf, in every size class, the test creates the object, asks the
// statistic for its size and compares the block the statistic derives with the
// number of bytes the runtime allocated.
func TestObjectSizes(t *testing.T) {
	var flat Map[uint64]
	flat.flat = 1
	var typed Map[string]
	typed.flat = 2
	var sets Map[uint64]
	sets.flat = -1
	check := func(t *testing.T, m interface{ object(*header) Object }, mk func() unsafe.Pointer) {
		t.Helper()
		got := allocated(mk)
		o := m.object((*header)(mk()))
		if b, _ := Block(o.Size, o.Pointers); b != got {
			t.Errorf("%s of %d bytes: Block %d, runtime allocated %d", o.Label, o.Size, b, got)
		}
	}
	t.Run("single-key pages", func(t *testing.T) {
		var pages Map[string]
		pages.flat = 3
		for _, n := range []int{20, 50, 100, 240, 254} { // 32, 64, 128, 256 and 384 bytes
			rest := bytes.Repeat([]byte("k"), n)
			check(t, &pages, func() unsafe.Pointer { return unsafe.Pointer(newSK(rest, 0)) })
		}
		long := bytes.Repeat([]byte("v"), 251)
		check(t, &pages, func() unsafe.Pointer { // 512 bytes
			return unsafe.Pointer(skHead(skpage.New(bytes.Repeat([]byte("k"), 254), 300, long)))
		})
	})
	t.Run("nodes, with and without a prefix tail", func(t *testing.T) {
		for k, label := range objTypeLabels {
			if label == "" {
				continue
			}
			for _, plen := range []int{0, 12, 13, 28, 29, 60, 61, 124} { // the ends of the tail classes
				path := make([]byte, plen)
				mk := func() unsafe.Pointer {
					n := newNode(objType(k), plen)
					storePrefix(n, path)
					return unsafe.Pointer(n)
				}
				check(t, &flat, mk)
			}
			// a path beyond the tail classes is a string in a 16-byte tail
			long := func() unsafe.Pointer { return unsafe.Pointer(newNode(objType(k), 200)) }
			size := int(fixedSize[k]) + 16
			if b, _ := Block(size, true); b != allocated(long) {
				t.Errorf("%s with a long path of %d bytes: Block %d, runtime allocated %d", label, size, b, allocated(long))
			}
		}
	})
	t.Run("pages of every class", func(t *testing.T) {
		for class := range 4 {
			check(t, &flat, func() unsafe.Pointer { return unsafe.Pointer(vpage.New(class, 0)) })
		}
	})
	t.Run("flat leaves of every class", func(t *testing.T) {
		for cls := 1; cls < len(flatSizes); cls++ {
			check(t, &flat, func() unsafe.Pointer { return unsafe.Pointer(allocFlat(uint8(cls))) })
		}
	})
	t.Run("typed leaves of every class and key area", func(t *testing.T) {
		for c := 1; c < len(typedCaps); c++ {
			for _, klen := range []int{0, 2, 3, 10, 11, 18, 19, 26, 27, 34, 35, 42, 43, 50, 51, 58} {
				check(t, &typed, func() unsafe.Pointer {
					l := allocTyped[string](uint8(c), klen)
					l.setRem(klen)
					return unsafe.Pointer(l)
				})
			}
		}
	})
	t.Run("value overflows of every key area, and with the key as a string", func(t *testing.T) {
		for _, n := range []int{0, 16, 17, 32, 33, 48, 49, 64, 65, 96, 97, 128, 129, 192, 193, maxInline} {
			key := make([]byte, n)
			check(t, &sets, func() unsafe.Pointer { return unsafe.Pointer(newSetLeafOf[uint64](key, n)) })
		}
		// the key of such a leaf is a string of its own; the leaf is the object whose address the test takes
		long := make([]byte, 300)
		check(t, &sets, func() unsafe.Pointer { return unsafe.Pointer(newSetLeaf[uint64](long, 0)) })
		var strs Map[string]
		strs.flat = -1
		check(t, &strs, func() unsafe.Pointer { return unsafe.Pointer(newSetLeaf[string](long, 0)) })
	})
}

// TestObjects makes sure that the object statistic of the index accounts for
// every key. It belongs to the object statistic of the ART behind
// multimap.Ordered (docs/redesign), which reports per benchmark case how the
// objects fill cache lines; an object type the statistic does not know would
// silently drop keys from its figures. The test fills maps of every leaf type
// with the corpora of the other tests and expects the keys of the objects to
// add up to the keys of the map, every object to have a label and a size, and
// an empty map to have no objects. Every leaf must also report how many values
// it holds and how long a key remainder, which the single-key page statistic
// of the bench (PLAN step 3) is made of.
func TestObjects(t *testing.T) {
	str := func(v uint64) string { return string(rune('a' + v%26)) }
	fill := func(t *testing.T, name string, keys [][]byte, second bool, add func(k []byte, v uint64), count func(func(Object)), length func() int) {
		t.Helper()
		for i, k := range keys {
			add(k, uint64(i))
			if second && i%3 == 0 {
				add(k, uint64(i)+1000) // a second value, which gives the key a leaf
			}
		}
		total, objects := 0, 0
		count(func(o Object) {
			total += o.Keys
			objects++
			if o.Label == "" || o.Size < 16 || o.Size%8 != 0 {
				t.Errorf("%s: object %+v is not an object of the tree", name, o)
			}
			if single := strings.HasSuffix(o.Label, "leaf") || o.Label == "single-key page" || o.Label == "value overflow"; single != (o.Values > 0) || !single && o.Remainder != 0 {
				t.Errorf("%s: object %+v: only a single-key page has values and a remainder", name, o)
			}
		})
		if total != length() {
			t.Errorf("%s: %d objects hold %d keys, the map %d", name, objects, total, length())
		}
	}
	for name, keys := range keySets() {
		var pages, mixed, flat, sets Map[uint64]
		flat.flat, sets.flat = 1, -1
		var typed Map[string]
		typed.flat = 2
		fill(t, name+"/pages", keys, false, func(k []byte, v uint64) { pages.Add(k, v) }, pages.Objects, pages.Len)
		fill(t, name+"/pages and leaves", keys, true, func(k []byte, v uint64) { mixed.Add(k, v) }, mixed.Objects, mixed.Len)
		fill(t, name+"/flat", keys, true, func(k []byte, v uint64) { flat.Add(k, v) }, flat.Objects, flat.Len)
		fill(t, name+"/typed", keys, true, func(k []byte, v uint64) { typed.Add(k, str(v)) }, typed.Objects, typed.Len)
		fill(t, name+"/sets", keys, true, func(k []byte, v uint64) { sets.Add(k, v) }, sets.Objects, sets.Len)
		var strPages Map[string]
		fill(t, name+"/single-key pages", keys, true, func(k []byte, v uint64) { strPages.Add(k, str(v)) }, strPages.Objects, strPages.Len)
		if name == "u64-dense" { // one value per integer key: pages below range nodes
			seen := map[string]bool{}
			pages.Objects(func(o Object) { seen[o.Label] = true })
			ranges := seen["R8"] || seen["R24"] || seen["R56"] || seen["R256"]
			if !seen["page"] || !ranges {
				t.Errorf("%s with one value per key has the objects %v, want pages and range nodes", name, seen)
			}
		}
	}
	var empty Map[uint64]
	empty.Objects(func(o Object) { t.Errorf("an empty map has the object %+v", o) })
}
