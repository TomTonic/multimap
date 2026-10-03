package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/TomTonic/multimap/internal/vpage"
)

// TestRun makes sure that the page fill experiment reports a row for every key
// kind it is asked for, with the figures the redesign's step 1 needs. It
// belongs to the page prototype (docs/redesign, PLAN step 1): integers must
// come out as uniform pages of 16 bytes a key at about 70% fill, strings as
// general pages, and a typo must be reported.
func TestRun(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "u64,uuid", "-n", "20000"}); err != nil {
		t.Fatal(err)
	}
	rows := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(rows) != 6 || !strings.HasPrefix(rows[4], "| u64 | 20000 |") || !strings.HasPrefix(rows[5], "| uuid | 20000 |") {
		t.Fatalf("rows:\n%s", out.String())
	}
	if c := strings.Split(rows[4], " | "); c[6] != "100 %" {
		t.Errorf("integer keys: %s of the pages are uniform, want all", c[6])
	}
	if c := strings.Split(rows[5], " | "); c[6] != "0 %" {
		t.Errorf("uuids: %s of the pages are uniform, want none", c[6])
	}
	if err := run(&out, []string{"-keys", "nope"}); err == nil {
		t.Error("an unknown key kind was accepted")
	}
	if err := run(&out, []string{"-nope"}); err == nil {
		t.Error("an unknown flag was accepted")
	}
}

// TestHelpers makes sure that the pieces of the page fill experiment measure
// what their names say. It belongs to the page prototype (docs/redesign, PLAN
// step 1).
func TestHelpers(t *testing.T) {
	if got := lcp([]byte("abcd"), []byte("abxy")); got != 2 {
		t.Errorf("lcp = %d, want 2", got)
	}
	if got := lcp([]byte("ab"), []byte("abxy")); got != 2 {
		t.Errorf("lcp of a prefix = %d, want 2", got)
	}
	// a suffix longer than a page holds is counted, not inserted
	m := &model{keys: [][]byte{make([]byte, 300), []byte("ab")}, bounds: [][]byte{{}}, bases: []int{0}}
	if bytes, tooLong := m.fill(); bytes != 2 || tooLong != 1 {
		t.Errorf("fill = %d bytes, %d too long; want 2 and 1", bytes, tooLong)
	}
	for pages, want := range map[int]int{1: 128, 8: 128, 9: 256, 24: 256, 25: 512, 56: 512, 57: 2112} {
		if got := routerBytes(pages); got != want {
			t.Errorf("routerBytes(%d) = %d, want %d", pages, got, want)
		}
	}
}

// TestTimings makes sure that the churn experiment of the page prefix reports,
// for each key kind, a row with the prefix off and one with it on, and counts
// the pages built anew. It belongs to the page prototype (docs/redesign, PLAN
// step 1), whose page prefix must not make `churn` dearer; the timings are for
// diagnosis only, so the test checks the shape of the report, not the numbers.
func TestTimings(t *testing.T) {
	defer func(a, b, c int) { vpage.MinPrefix, vpage.MinGain, vpage.PrefixSlack = a, b, c }(vpage.MinPrefix, vpage.MinGain, vpage.PrefixSlack)
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "path", "-n", "20000", "-timing", "2000"}); err != nil {
		t.Fatal(err)
	}
	rows := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(rows) != 6 || !strings.HasPrefix(rows[4], "| path | off |") || !strings.HasPrefix(rows[5], "| path | on |") {
		t.Fatalf("rows:\n%s", out.String())
	}
	if err := run(&out, []string{"-keys", "nope", "-timing", "10"}); err == nil {
		t.Error("an unknown key kind was accepted by the timing run")
	}
}

// TestLens makes sure that the model of the length-header page reports sizes a
// page of that layout can have. It belongs to the comparison of multi-key page
// layouts (docs/redesign, PLAN step 3): the sketched page holds as many entries as
// its header has length bytes, so integers with a header of 16 bytes and fixed
// values come out at 14 entries at most and below 14 on average, the bytes of
// header, prefix and remainders stay within the object, and bad arguments are
// reported.
func TestLens(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-layout", "lens", "-lenfixed", "-lenheader", "16", "-keys", "u64,uuid", "-n", "20000"}); err != nil {
		t.Fatal(err)
	}
	rows := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(rows) != 8 || !strings.HasPrefix(rows[6], "| u64 | 20000 |") || !strings.HasPrefix(rows[7], "| uuid | 20000 |") {
		t.Fatalf("rows:\n%s", out.String())
	}
	for _, row := range rows[6:] {
		c := strings.Split(row, " | ")
		var perPage, fill float64
		if _, err := fmt.Sscanf(c[4], "%f", &perPage); err != nil || perPage <= 1 || perPage > 14 {
			t.Errorf("%s: %q entries per page, want more than 1 and at most 14", c[0], c[4])
		}
		if _, err := fmt.Sscanf(c[5], "%f", &fill); err != nil || fill < 40 || fill > 100 {
			t.Errorf("%s: fill %q", c[0], c[5])
		}
	}
	for _, args := range [][]string{
		{"-layout", "lens", "-lenclasses", "x"},
		{"-layout", "lens", "-lenheader", "12"},
		{"-layout", "lens", "-keys", "nope"},
	} {
		if err := run(&out, args); err == nil {
			t.Errorf("%v was accepted", args)
		}
	}
}

// TestLensSize makes sure that the size of a page of the length-header layout
// adds up as the sketch says. It belongs to the comparison of multi-key page
// layouts (docs/redesign, PLAN step 3): header, common prefix, remainders and
// values, the prefix one byte shorter than the shortest suffix at most, and a
// page that does not fit is refused.
func TestLensSize(t *testing.T) {
	l := lensLayout{fixed: true, value: 8, maxHeader: 16, classes: []int{128, 256, 512}}
	for _, tc := range []struct {
		name string
		keys []string
		want lensSize
		ok   bool
	}{
		{"one entry has no prefix", []string{"abc"}, lensSize{8, 0, 3, 8 + 3 + 8}, true},
		{"entries share their prefix", []string{"abcX", "abcY"}, lensSize{8, 3, 2, 8 + 3 + 2 + 16}, true},
		{"the prefix leaves one byte of the shortest", []string{"ab", "abc"}, lensSize{8, 1, 3, 8 + 1 + 3 + 16}, true},
		{"seven entries need a header of 16", []string{"a", "b", "c", "d", "e", "f", "g"}, lensSize{16, 0, 7, 16 + 7 + 56}, true},
		{"more entries than the largest header", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}, lensSize{}, false},
		{"an empty suffix", []string{""}, lensSize{}, false},
		{"a remainder beyond 255 bytes", []string{strings.Repeat("x", 256)}, lensSize{}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			keys := make([][]byte, len(tc.keys))
			for i, k := range tc.keys {
				keys[i] = []byte(k)
			}
			got, ok := l.size(keys)
			if ok != tc.ok || got != tc.want {
				t.Errorf("size = %+v, %v; want %+v, %v", got, ok, tc.want, tc.ok)
			}
		})
	}
	if c := l.class(129); c != 256 {
		t.Errorf("class of 129 bytes: %d, want 256", c)
	}
	if c := l.class(513); c != 0 {
		t.Errorf("class of 513 bytes: %d, want none", c)
	}
}
