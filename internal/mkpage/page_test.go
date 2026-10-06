package mkpage

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
	"testing"
)

// kv is an entry of the reference model: the key from the end of the path on, and
// its value.
type kv struct{ key, val string }

// model is the reference of the page tests: the entries in key order.
type model []kv

func (m model) find(key string) (int, bool) {
	i := sort.Search(len(m), func(i int) bool { return m[i].key >= key })
	return i, i < len(m) && m[i].key == key
}

// randKey returns a key that starts with prefix and goes on with up to 6 bytes of a
// small alphabet, so that keys collide, share bytes and end where others go on.
func randKey(r *rand.Rand, prefix string) string {
	var b strings.Builder
	b.WriteString(prefix)
	for range r.IntN(7) {
		b.WriteByte("ab\x00\xff"[r.IntN(4)])
	}
	return b.String()
}

func randVal(r *rand.Rand) string {
	return strings.Repeat("v", r.IntN(12)) + string(rune('0'+r.IntN(3)))
}

// checkPage compares a page with the model and checks its invariants: every entry in
// order with the right value, the used part within the object, nothing but zeros
// behind it.
func checkPage(t *testing.T, p *Page, m model, prefix string) {
	t.Helper()
	if p == nil {
		if len(m) != 0 {
			t.Fatalf("page is nil, model has %d entries", len(m))
		}
		return
	}
	if p.Len() != len(m) {
		t.Fatalf("Len %d, model %d", p.Len(), len(m))
	}
	i := 0
	cp := string(p.CP())
	p.Each(func(rem, val []byte, _ bool) bool {
		if got := cp + string(rem); got != m[i].key || string(val) != m[i].val {
			t.Fatalf("entry %d: %q -> %q, model %q -> %q", i, got, val, m[i].key, m[i].val)
		}
		i++
		return true
	})
	if !strings.HasPrefix(prefix, cp) && !strings.HasPrefix(cp, prefix) {
		t.Fatalf("common prefix %q unrelated to %q", cp, prefix)
	}
	if used := p.Used(); used > p.Size() {
		t.Fatalf("used %d exceeds the size %d", used, p.Size())
	} else if bytes.Count(p.mem()[used:], []byte{0}) != p.Size()-used {
		t.Fatalf("bytes behind the used part are not zero")
	}
	for _, e := range m {
		if v, ok := p.Get([]byte(e.key)); !ok || string(v) != e.val {
			t.Fatalf("Get(%q) = %q, %v; want %q", e.key, v, ok, e.val)
		}
	}
}

// expectInsert says what Add must answer for the model.
func expectInsert(m model, cp string, key, val string) Result {
	var r string
	if !strings.HasPrefix(key, cp) {
		return Outside
	}
	r = key[len(cp):]
	if i, ok := m.find(key); ok && m[i].val == val {
		return Present
	}
	if len(val) > MaxValue {
		return Full
	}
	rems, vals := len(r), len(val)
	for _, e := range m {
		rems += len(e.key) - len(cp)
		vals += len(e.val)
	}
	if len(r) > MaxRemainder || NeedStrings(len(m)+1, len(cp), rems, vals) > 512 {
		return Full
	}
	return Added
}

// TestPageAgainstModel keeps a multi-key page of strings and a plain sorted list
// side by side through thousands of random inserts and removes, from many starting
// pages, and compares them after every step.
//
// For a user of the multimap this is the guarantee that a page holding many keys
// with one value each answers exactly like the simple structure it stands for:
// every key finds its value, nothing is lost or invented, the keys stay in order.
// The page is the multi-key page of the redesign (step 4.1); the tree will hold
// many of them.
//
// Expected: after every operation the result is the one the model predicts, the
// entries are equal, the layout invariants hold, and a page that grows or shrinks
// keeps its entries.
func TestPageAgainstModel(t *testing.T) {
	for seed := range uint64(40) {
		r := rand.New(rand.NewPCG(seed, 7))
		prefix := randKey(r, "")
		prefix = prefix[:min(len(prefix), r.IntN(3))]
		var m model
		for range 1 + r.IntN(4) {
			k := randKey(r, prefix)
			if i, ok := m.find(k); !ok {
				m = slices.Insert(m, i, kv{k, randVal(r)})
			}
		}
		rests := make([][]byte, len(m))
		vals := make([][]byte, len(m))
		for i, e := range m {
			rests[i], vals[i] = []byte(e.key), []byte(e.val)
		}
		p := BuildStrings(rests, vals)
		cp := string(p.CP())
		checkPage(t, p, m, prefix)
		for step := range 600 {
			key := randKey(r, cp)
			if r.IntN(20) == 0 {
				key = randKey(r, "")
			}
			var val string
			switch r.IntN(30) {
			case 0:
				val = strings.Repeat("L", MaxValue+1)
			case 1:
				val = strings.Repeat("M", 60+r.IntN(100))
			default:
				val = randVal(r)
			}
			if r.IntN(3) > 0 {
				q, rm := p.Remove([]byte(key), []byte(val))
				ok := rm != Absent
				i, found := m.find(key)
				want := found && strings.HasPrefix(key, cp) && m[i].val == val
				if ok != want {
					t.Fatalf("seed %d step %d: Remove(%q, %q) = %v, want %v", seed, step, key, val, ok, want)
				}
				if want {
					m = slices.Delete(m, i, i+1)
				}
				p = q
				if p == nil {
					if len(m) != 0 {
						t.Fatalf("seed %d step %d: page gone, model has %d entries", seed, step, len(m))
					}
					break
				}
			} else {
				if i, ok := m.find(key); ok { // these models hold one value a key
					val = m[i].val
				}
				want := expectInsert(m, cp, key, val)
				q, res := p.Add([]byte(key), []byte(val))
				if res != want {
					t.Fatalf("seed %d step %d: Insert(%q, %q) = %v, want %v", seed, step, key, val, res, want)
				}
				if res == Added {
					i, _ := m.find(key)
					m = slices.Insert(m, i, kv{key, val})
				} else if q != p {
					t.Fatalf("seed %d step %d: Insert returned another page for %v", seed, step, res)
				}
				p = q
			}
			checkPage(t, p, m, prefix)
		}
	}
}

// TestPageLimits shows which sets of entries make a page and which do not.
//
// The tree asks BuildStrings whether the entries of a subtree fit one page before it
// decides between a page and a byte node. A page has to refuse what it cannot hold
// (too many entries, remainders or values that are too long, more than 512 bytes)
// and take what it can, at the border.
//
// Expected: nil for no entry, a different number of values, more than 255 entries, a
// remainder of 256 bytes, a value of 255 bytes, 513 bytes of content, and a page
// for the largest sets that fit.
func TestPageLimits(t *testing.T) {
	b := func(ss ...string) [][]byte {
		out := make([][]byte, len(ss))
		for i, s := range ss {
			out[i] = []byte(s)
		}
		return out
	}
	many := make([][]byte, 256)
	manyV := make([][]byte, 256)
	for i := range many {
		many[i] = []byte{byte(i)}
		manyV[i] = nil
	}
	tests := []struct {
		name string
		keys [][]byte
		vals [][]byte
		want bool
	}{
		{"no entry", nil, nil, false},
		{"values and keys differ in number", b("a", "b"), b("x"), false},
		{"256 entries", many, manyV, false},
		{"255 entries of one byte each", many[:255], manyV[:255], false}, // 3+255+255+255 = 768 bytes
		{"a remainder of 255 bytes (the length of a further value)", b("a", "a"+strings.Repeat("x", 255)), b("1", "2"), false},
		{"a remainder of 254 bytes", b("a", "a"+strings.Repeat("x", 254)), b("1", "2"), true},
		{"a value of 255 bytes", b("a", "b"), b("1", strings.Repeat("v", 255)), false},
		{"a value of 254 bytes", b("a", "b"), b("1", strings.Repeat("v", 254)), true},
		{"content of 513 bytes", b("a", "b"), b(strings.Repeat("v", 254), strings.Repeat("v", 250)), false},
		{"content of 512 bytes", b("a", "b"), b(strings.Repeat("v", 254), strings.Repeat("v", 249)), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildStrings(tt.keys, tt.vals) != nil; got != tt.want {
				t.Errorf("BuildStrings built a page: %v, want %v", got, tt.want)
			}
		})
	}
	if got := NeedStrings(2, 1, 4, 6); got != 3+2+1+4+2+6 {
		t.Errorf("NeedStrings = %d", got)
	}
}

// TestPageInsertResults shows each answer an insert can give, and that only a page
// that really changed is a different page.
//
// The tree acts on the answer: Added and AddedValue need nothing, Present nothing, Full
// means a burst, Outside a byte node above the page.
//
// Expected: one case for every Result, with the page and its entries unchanged where
// nothing was added.
func TestPageInsertResults(t *testing.T) {
	fresh := func() *Page {
		return BuildStrings([][]byte{[]byte("abc1"), []byte("abc2")}, [][]byte{[]byte("x"), []byte("y")})
	}
	p := fresh()
	if got := string(p.CP()); got != "abc" {
		t.Fatalf("common prefix %q", got)
	}
	tests := []struct {
		name string
		key  string
		val  string
		want Result
	}{
		{"adds a key under the prefix", "abc3", "z", Added},
		{"adds the key that ends at the prefix", "abc", "z", Added},
		{"knows an entry that is there", "abc1", "x", Present},
		{"adds a second value to a key", "abc1", "other", AddedValue},
		{"refuses a key outside the prefix", "abd1", "x", Outside},
		{"refuses a key shorter than the prefix", "ab", "x", Outside},
		{"refuses a value of 255 bytes", "abc4", strings.Repeat("v", 255), Full},
		{"refuses a second value of 255 bytes", "abc1", strings.Repeat("v", 255), Full},
		{"refuses a remainder of 256 bytes", "abc" + strings.Repeat("k", 256), "x", Full},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := fresh()
			q, res := p.Add([]byte(tt.key), []byte(tt.val))
			if res != tt.want {
				t.Fatalf("Insert = %v, want %v", res, tt.want)
			}
			if res != Added && res != AddedValue && (q != p || p.Len() != 2) {
				t.Errorf("the page changed without an addition")
			}
			if v, ok := q.Get([]byte(tt.key)); res == Added && (!ok || string(v) != tt.val) {
				t.Errorf("Get after Added = %q, %v", v, ok)
			}
		})
	}
	if _, ok := p.Get([]byte("zzz")); ok {
		t.Error("Get found a key outside the prefix")
	}
	if _, ok := p.Get([]byte("abc9")); ok {
		t.Error("Get found a key that is not there")
	}
	if q, rm := p.Remove([]byte("abd1"), []byte("x")); rm != Absent || q != p {
		t.Error("Remove took a key outside the prefix")
	}
	if q, rm := p.Remove([]byte("abc9"), []byte("x")); rm != Absent || q != p {
		t.Error("Remove took a key that is not there")
	}
	if q, rm := p.Remove([]byte("abc1"), []byte("y")); rm != Absent || q != p {
		t.Error("Remove took an entry with another value")
	}
	if q, rm := BuildStrings([][]byte{[]byte("k")}, [][]byte{[]byte("v")}).Remove([]byte("k"), []byte("v")); rm != Gone || q != nil {
		t.Error("Remove of the only entry")
	}
	if p.Class() != 0 || p.Size() != 32 {
		t.Errorf("class %d, size %d", p.Class(), p.Size())
	}
}

// TestPageFullAtTheLimits fills a page with entries of one byte until it refuses.
//
// A page that reaches 512 bytes must say Full and stay as it was, so that the tree
// can burst it without losing an entry.
//
// Expected: the entry that does not fit is refused, the page is then of the largest
// class and holds the entries it had.
func TestPageFullAtTheLimits(t *testing.T) {
	p := BuildStrings([][]byte{{0}, {255}}, [][]byte{nil, nil})
	for i := 1; i < 255; i++ {
		q, res := p.Add([]byte{byte(i)}, nil)
		if res == Full {
			// each entry takes the length of its remainder, one byte, and its value length
			if p.Size() != 512 || p.Used() != 3+p.Len()*3 || p.Used() > 512 || p.Used() < 512-3 {
				t.Fatalf("Full at %d entries, %d bytes used of %d", p.Len(), p.Used(), p.Size())
			}
			return
		}
		p = q
	}
	t.Fatal("never full")
}

// TestPageSkipAndPrepend moves a page down and up in the tree: a byte node arrives
// above it and takes bytes of its common prefix, or goes away and gives them back.
//
// The page keeps its entries through both; what changes is how much of every key
// it spells itself. Prepend may need a larger object, or refuse when the prefix
// or the content would be too long.
//
// Expected: after Skip the entries read as the keys without their first bytes;
// Prepend restores them, grows the class when needed, and refuses at 255 bytes of
// prefix or 512 bytes of content.
func TestPageSkipAndPrepend(t *testing.T) {
	p := BuildStrings([][]byte{[]byte("abcdef1"), []byte("abcdef2")}, [][]byte{[]byte("x"), []byte("y")})
	p.Skip(2)
	if got := string(p.CP()); got != "cdef" {
		t.Fatalf("after Skip: %q", got)
	}
	if v, ok := p.Get([]byte("cdef2")); !ok || string(v) != "y" {
		t.Fatalf("Get after Skip: %q, %v", v, ok)
	}
	p.Skip(4)
	if p.PrefixLen() != 0 || p.Used() != Header+2+(1+1+1)+(1+1+1) {
		t.Fatalf("after Skip(4): prefix %d used %d", p.PrefixLen(), p.Used())
	}
	q := p.Prepend([]byte("abcd"))
	if q != p {
		t.Fatalf("Prepend allocated although the class holds it")
	}
	if v, ok := q.Get([]byte("abcd1")); !ok || string(v) != "x" {
		t.Fatalf("Get after Prepend: %q, %v", v, ok)
	}
	big := q.Prepend([]byte(strings.Repeat("P", 40)))
	if big == nil || big == q || big.Size() <= q.Size() {
		t.Fatalf("Prepend of 40 bytes: %v", big)
	}
	if v, ok := big.Get([]byte(strings.Repeat("P", 40) + "abcd2")); !ok || string(v) != "y" {
		t.Fatalf("Get after the larger Prepend: %q, %v", v, ok)
	}
	if big.Prepend([]byte(strings.Repeat("P", MaxPrefix))) != nil {
		t.Error("Prepend beyond MaxPrefix")
	}
	fat := BuildStrings([][]byte{[]byte("1"), []byte("2")}, [][]byte{[]byte(strings.Repeat("v", 240)), []byte(strings.Repeat("w", 240))})
	if fat.Prepend([]byte(strings.Repeat("P", 40))) != nil {
		t.Error("Prepend beyond the largest class")
	}
}

// TestPageShrinksWithHysteresis removes entries from a full page and shows that it
// changes class only when the content fills half of the smaller one.
//
// A page whose entries come and go around a class border must not copy itself with
// every operation (the lesson of step 3.5).
//
// Expected: the page keeps its class while its content is more than half of the next
// smaller class, shrinks when it is at most half of it, and a page at the border
// does not change its object when an entry is added and removed again.
func TestPageShrinksWithHysteresis(t *testing.T) {
	p := BuildStrings([][]byte{{0}, {255}}, [][]byte{[]byte("vvvvvvv"), []byte("vvvvvvv")})
	for i := 1; i < 40; i++ {
		var res Result
		if p, res = p.Add([]byte{byte(i)}, []byte("vvvvvvv")); res != Added {
			t.Fatalf("entry %d: %v", i, res)
		}
	}
	if p.Size() != 512 {
		t.Fatalf("size %d", p.Size())
	}
	classes := []int{p.Size()}
	for i := 39; i >= 1; i-- {
		var rm Removal
		if p, rm = p.Remove([]byte{byte(i)}, []byte("vvvvvvv")); rm == Absent {
			t.Fatalf("entry %d not removed", i)
		}
		if c := classFor(2 * p.Used()); c >= 0 && c < p.class() {
			t.Fatalf("content %d in %d bytes should have shrunk", p.Used(), p.Size())
		}
		if p.Size() != classes[len(classes)-1] {
			classes = append(classes, p.Size())
		}
	}
	if !slices.Equal(classes, []int{512, 384, 256, 128, 64}) {
		t.Errorf("classes passed: %v", classes)
	}
	size := p.Size()
	for range 5 {
		q, res := p.Add([]byte{9}, []byte("vvvvvvv"))
		if res != Added {
			t.Fatalf("insert: %v", res)
		}
		r, rm := q.Remove([]byte{9}, []byte("vvvvvvv"))
		ok := rm != Absent
		if !ok || r.Size() != size {
			t.Fatalf("remove after insert: %v, size %d, was %d", ok, r.Size(), size)
		}
		p = r
	}
}

// TestPageEachAndEqual covers iterating with an early stop and comparing pages.
//
// The tree's scans stop when the caller's callback says so; the tests compare pages
// built differently.
//
// Expected: Each stops at the first false and reports it; Equal ignores the class,
// finds a page with another key, another value or another number of entries
// different, and sees through a different split of the key into prefix and remainder.
func TestPageEachAndEqual(t *testing.T) {
	build := func(keys []string, vals ...string) *Page {
		rs, vs := make([][]byte, len(keys)), make([][]byte, len(keys))
		for i := range keys {
			rs[i], vs[i] = []byte(keys[i]), []byte(vals[i])
		}
		return BuildStrings(rs, vs)
	}
	a := build([]string{"ab", "ac"}, "1", "2")
	n := 0
	if a.Each(func(rem, val []byte, _ bool) bool { n++; return false }) || n != 1 {
		t.Errorf("Each did not stop: %d", n)
	}
	tests := []struct {
		name string
		b    *Page
		want bool
	}{
		{"the same entries", build([]string{"ab", "ac"}, "1", "2"), true},
		{"the same entries with a shorter prefix", func() *Page {
			p := build([]string{"ab", "ac"}, "1", "2")
			p.Skip(0)
			return p
		}(), true},
		{"another key", build([]string{"ab", "ad"}, "1", "2"), false},
		{"another value", build([]string{"ab", "ac"}, "1", "9"), false},
		{"fewer entries", build([]string{"ab"}, "1"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(a, tt.b); got != tt.want {
				t.Errorf("Equal = %v, want %v", got, tt.want)
			}
		})
	}
}

// FuzzPage runs the page against the model on operations chosen by the fuzzer.
//
// It covers what random keys of a small alphabet may not: any byte value, any
// lengths, in any order.
//
// Expected: the page and the model agree after every operation.
func FuzzPage(f *testing.F) {
	f.Add([]byte("abc\x00x1abd\x00y2\x01abc"))
	f.Fuzz(func(t *testing.T, data []byte) {
		var p *Page
		var m model
		for len(data) >= 3 {
			op, kl, vl := data[0]%3, int(data[1])%9, int(data[2])%9
			data = data[3:]
			if len(data) < kl+vl {
				return
			}
			key, val := string(data[:kl]), string(data[kl:kl+vl])
			data = data[kl+vl:]
			i, found := m.find(key)
			switch {
			case p == nil && op != 1:
				p = BuildStrings([][]byte{[]byte(key)}, [][]byte{[]byte(val)})
				m = model{{key, val}}
			case p == nil:
			case op == 1:
				q, rm := p.Remove([]byte(key), []byte(val))
				ok := rm != Absent
				if want := found && strings.HasPrefix(key, string(p.CP())) && m[i].val == val; ok != want {
					t.Fatalf("Remove %q: %v, want %v", key, ok, want)
				} else if want {
					m = slices.Delete(m, i, i+1)
				}
				p = q
			default:
				if found { // this model holds one value a key
					val = m[i].val
				}
				want := expectInsert(m, string(p.CP()), key, val)
				q, res := p.Add([]byte(key), []byte(val))
				if res != want {
					t.Fatalf("Insert %q: %v, want %v", key, res, want)
				}
				if res == Added {
					m = slices.Insert(m, i, kv{key, val})
				}
				p = q
			}
			if p == nil {
				m = nil
			}
			checkPage(t, p, m, "")
		}
	})
}

// TestPageWiden adds keys that leave the common prefix of a page and checks that the page
// takes them in, with a shorter prefix, exactly where a model puts them.
//
// The tree does this where it would otherwise make a byte node and a single-key page: a key that
// differs early from its neighbours joins their page if they fit. The page must keep every
// entry, put the new one in order (below all or above all, since it leaves the prefix), and
// refuse what does not fit.
//
// Expected: for pages with prefixes of 0 to 4 bytes and new keys that differ at every position, below
// and above, including one that ends inside the prefix, the widened page equals the model; a page
// that is full, a remainder that would exceed 255 bytes and a value of 255 bytes give nil.
func TestPageWiden(t *testing.T) {
	for _, cpl := range []int{1, 2, 4} {
		for mis := range cpl {
			for _, kind := range []string{"below", "above", "ends"} {
				t.Run(fmt.Sprint(cpl, "/", mis, "/", kind), func(t *testing.T) {
					cp := strings.Repeat("m", cpl)
					m := model{{cp + "a", "1"}, {cp + "bb", "2"}, {cp + "c", "3"}}
					rests, vals := make([][]byte, 3), make([][]byte, 3)
					for i, e := range m {
						rests[i], vals[i] = []byte(e.key), []byte(e.val)
					}
					p := BuildStrings(rests, vals)
					if p.PrefixLen() != cpl {
						t.Fatalf("prefix %d", p.PrefixLen())
					}
					key := cp[:mis]
					switch kind {
					case "below":
						key += "a" // 'a' < 'm'
					case "above":
						key += "z"
					}
					q := p.Widen([]byte(key), []byte("new"))
					if q == nil {
						t.Fatal("Widen refused")
					}
					i, _ := m.find(key)
					m = slices.Insert(m, i, kv{key, "new"})
					if q.PrefixLen() != mis {
						t.Errorf("prefix %d, want %d", q.PrefixLen(), mis)
					}
					checkPage(t, q, m, "")
				})
			}
		}
	}
	long := strings.Repeat("v", 200)
	full := BuildStrings([][]byte{[]byte("ma"), []byte("mb")}, [][]byte{[]byte(long), []byte(long)})
	tests := []struct {
		name string
		p    *Page
		key  string
		val  string
	}{
		{"a page that is full", full, "z", long},
		{"a value of 255 bytes", full, "z", strings.Repeat("v", 255)},
		{"a new remainder of 256 bytes", full, "z" + strings.Repeat("k", 255), "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.p.Widen([]byte(tt.key), []byte(tt.val)) != nil {
				t.Error("Widen took the entry")
			}
		})
	}
}

// TestFixedWiden is TestPageWiden for a page of uint64 values.
//
// Expected: the widened page equals the model for every position and side; a full page and a
// remainder of more than 255 bytes give nil.
func TestFixedWiden(t *testing.T) {
	for _, cpl := range []int{1, 2, 4} {
		for mis := range cpl {
			for _, kind := range []string{"below", "above", "ends"} {
				t.Run(fmt.Sprint(cpl, "/", mis, "/", kind), func(t *testing.T) {
					cp := strings.Repeat("m", cpl)
					m := fixedModel[uint64]{{cp + "a", 1}, {cp + "bb", 2}, {cp + "c", 3}}
					p := BuildFixed([][]byte{[]byte(cp + "a"), []byte(cp + "bb"), []byte(cp + "c")}, []uint64{1, 2, 3})
					key := cp[:mis]
					switch kind {
					case "below":
						key += "a"
					case "above":
						key += "z"
					}
					q := p.Widen([]byte(key), uint64(9))
					if q == nil {
						t.Fatal("Widen refused")
					}
					i, _ := m.find(key)
					m = slices.Insert(m, i, struct {
						key string
						val uint64
					}{key, 9})
					if q.PrefixLen() != mis {
						t.Errorf("prefix %d, want %d", q.PrefixLen(), mis)
					}
					checkFixed(t, q, m)
				})
			}
		}
	}
	rests := make([][]byte, 50)
	vals := make([]uint64, 50)
	for i := range rests {
		rests[i] = []byte{'m', byte('A' + i)}
	}
	full := BuildFixed(rests, vals)
	if full == nil || full.PrefixLen() != 1 || full.Widen([]byte("z"), 1) != nil {
		t.Error("Widen took an entry into a full page")
	}
	if full.Widen([]byte("z"+strings.Repeat("k", 256)), 1) != nil {
		t.Error("Widen took a remainder of 256 bytes")
	}
}
