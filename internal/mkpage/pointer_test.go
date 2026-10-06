package mkpage

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"slices"
	"sync/atomic"
	"testing"
	"time"
)

// ptrRec is a value of the pointer pages of the tests: a finalizer says when the collector
// has found it unreachable.
type ptrRec struct{ id int }

var ptrDead [1 << 12]atomic.Bool

func newPtrRec(id int) *ptrRec {
	r := &ptrRec{id}
	runtime.SetFinalizer(r, func(r *ptrRec) { ptrDead[r.id].Store(true) })
	return r
}

// settle runs the collector until the finalizers of what it has found are done.
func settle() {
	for range 3 {
		runtime.GC()
		time.Sleep(time.Millisecond)
	}
}

// TestFixedPointerPageKeepsItsValuesAlive shows that the values of a page of pointers stay
// reachable for the garbage collector through every change of the page.
//
// A user whose map holds *Record values must never see a record freed while the map holds it,
// whichever way the index moved it: a page of pointers is a typed object that the collector
// scans for its last words, and a move of the values as bytes would hide them (write barrier).
//
// Expected: after random adds, removes, widenings, Skip and Prepend, with a collection after every
// few steps, no value that the page holds has been finalized and Get returns the very object
// that was added; a value that was removed and not held elsewhere is freed (the control that
// finalizers work).
func TestFixedPointerPageKeepsItsValuesAlive(t *testing.T) {
	type entry struct {
		key string
		ids []int
	}
	for seed := range uint64(6) {
		r := rand.New(rand.NewPCG(seed, 5))
		next := int(seed) * 600
		var model []entry
		find := func(k string) (int, bool) {
			return slices.BinarySearchFunc(model, k, func(e entry, k string) int {
				switch {
				case e.key < k:
					return -1
				case e.key > k:
					return 1
				}
				return 0
			})
		}
		rests := [][]byte{[]byte("key-a"), []byte("key-b")}
		vals := []*ptrRec{newPtrRec(next), newPtrRec(next + 1)}
		model = []entry{{"key-a", []int{next}}, {"key-b", []int{next + 1}}}
		next += 2
		p := BuildFixed(rests, vals)
		for step := range 500 {
			key := fmt.Sprintf("key-%c%c", 'a'+r.IntN(8), 'a'+r.IntN(2)*r.IntN(3))
			i, ok := find(key)
			switch op := r.IntN(10); {
			case op < 6:
				q, res := p.Add([]byte(key), newPtrRec(next), true)
				p = q
				switch res {
				case Added:
					model = slices.Insert(model, i, entry{key, []int{next}})
				case AddedValue:
					model[i].ids = append(model[i].ids, next)
				}
				next++
			case op < 9 && ok:
				id := model[i].ids[r.IntN(len(model[i].ids))]
				v := ptrDeref(t, p, key, id)
				q, rm := p.Remove([]byte(key), v, true)
				if q == nil {
					t.Fatalf("seed %d: the page went away", seed)
				}
				p = q
				model[i].ids = slices.DeleteFunc(model[i].ids, func(x int) bool { return x == id })
				if rm == Gone {
					model = slices.Delete(model, i, i+1)
				}
			case op == 9 && r.IntN(3) == 0:
				p.Skip[*ptrRec](2)
				p = p.Prepend[*ptrRec]([]byte("ke"), true)
			}
			if step%7 == 0 {
				settle()
				for _, e := range model {
					for _, id := range e.ids {
						if ptrDead[id].Load() {
							t.Fatalf("seed %d step %d: value %d of key %q was freed while the page holds it", seed, step, id, e.key)
						}
						ptrDeref(t, p, e.key, id)
					}
				}
			}
		}
	}
	// the control: a value that nothing holds is found by the collector
	func() { _ = newPtrRec(len(ptrDead) - 1) }()
	for range 100 {
		if settle(); ptrDead[len(ptrDead)-1].Load() {
			return
		}
	}
	t.Error("a value that nothing holds was never freed: the test cannot see a freed value")
}

// ptrDeref returns the value of key with the id, among the values of the page.
func ptrDeref(t *testing.T, p *Fixed, key string, id int) *ptrRec {
	t.Helper()
	var found *ptrRec
	p.EachValue([]byte(key), func(v *ptrRec) bool {
		if v.id == id {
			found = v
		}
		return true
	})
	if found == nil {
		t.Fatalf("key %q does not hold the value %d", key, id)
	}
	return found
}
