package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestRun makes sure that the object statistic of the benchmark's index reports
// the cases a developer asks for, in a table that can go straight into the
// results. It belongs to the cache-line analysis of the redesign (docs/redesign):
// the statistic tells, per benchmark case, how well the objects of the ART
// behind multimap.Ordered fill cache lines. Integer keys with one value each live
// in pages and must report no violation of the 64- and 128-byte rules, while
// the same keys with several values live in leaves and must report violations.
func TestRun(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "u64", "-sizes", "4096"}); err != nil {
		t.Fatal(err)
	}
	rows := map[string][]string{}
	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n")[2:] {
		cells := strings.Split(strings.Trim(line, "| "), " | ")
		rows[cells[0]] = cells
	}
	for _, name := range []string{"u64 natural 4K", "u64 single-value 4K", "u64 natural-str 4K", "u64 single-value-str 4K"} {
		if rows[name] == nil {
			t.Fatalf("no row for %q in\n%s", name, out.String())
		}
	}
	if len(rows) != 4 {
		t.Errorf("%d rows, want 4:\n%s", len(rows), out.String())
	}
	if r := rows["u64 single-value 4K"]; !strings.Contains(r[6], "flat leaf") {
		t.Errorf("integer keys with one value should live in flat leaves (the multi-key pages are out of the tree): %v", r)
	}
	if r := rows["u64 natural 4K"]; r[3] == "0.0 %" || !strings.Contains(r[6], "flat leaf") {
		t.Errorf("integer keys with several values should live in leaves, which are not all multiples of 64 bytes: %v", r)
	}
	if r := rows["u64 natural-str 4K"]; !strings.Contains(r[6], "single-key page") {
		t.Errorf("keys with several string values should live in single-key pages: %v", r)
	}
}

// TestRunMaximum makes sure that a size beyond what a real corpus holds is
// measured at the corpus's own maximum when asked to, and left out otherwise.
// It belongs to the cache-line analysis (docs/redesign), whose tables need a
// row for the street names although the corpus holds fewer keys than the
// largest benchmark size.
func TestRunMaximum(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want int // rows
	}{
		{"measures at the corpus maximum with -max", []string{"-keys", "street", "-values", "single-value", "-strvals=false", "-sizes", "1000000"}, 1},
		{"leaves the size out without -max", []string{"-keys", "street", "-values", "single-value", "-strvals=false", "-sizes", "1000000", "-max=false"}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := run(&out, tc.args); err != nil {
				t.Fatal(err)
			}
			if rows := strings.Count(out.String(), "\n") - 2; rows != tc.want {
				t.Errorf("%d rows, want %d:\n%s", rows, tc.want, out.String())
			}
		})
	}
}

// TestRunErrors makes sure that the tool says what is wrong with its
// arguments instead of printing an empty table. It belongs to the cache-line
// analysis (docs/redesign), and is what a developer sees after a typo.
func TestRunErrors(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want string
	}{
		{"rejects an unknown key kind", []string{"-keys", "nope"}, `unknown key kind "nope"`},
		{"rejects an unknown value profile", []string{"-values", "many"}, `unknown value profile "many"`},
		{"rejects a size that is not a number", []string{"-sizes", "4096,x"}, `bad size "x"`},
		{"rejects a size below one", []string{"-sizes", "0"}, `bad size "0"`},
		{"rejects an unknown flag", []string{"-nope"}, "flag provided but not defined"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(&out, tc.args)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error %v, want one that contains %q", err, tc.want)
			}
		})
	}
}

// TestLineOverflow makes sure that the statistic counts the cache lines an
// object touches the way the cache-line rules do (docs/redesign/STRATEGY.md):
// an object of 64 bytes in a 64-byte block fills one line and never overflows,
// one of 48 bytes crosses a line in half of the positions a span gives it, and
// a block that starts off the line boundary pushes a 64-byte object over two
// lines.
func TestLineOverflow(t *testing.T) {
	for _, tc := range []struct {
		name                string
		size, block, offset int
		want                float64
	}{
		{"a 64-byte object in a 64-byte block fills one line", 64, 64, 0, 0},
		{"a 32-byte object in a 32-byte block stays in its line", 32, 32, 0, 0},
		{"a 48-byte object crosses a line in half of its positions", 48, 48, 0, 0.5},
		{"a 96-byte object touches the two lines it needs", 96, 96, 0, 0},
		{"a 64-byte object 8 bytes into its block crosses a line", 64, 64, 8, 1},
		{"a 2080-byte object behind a malloc header needs the lines it needs", 2080, 2304, 8, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := lineOverflow(tc.size, tc.block, tc.offset); got != tc.want {
				t.Errorf("lineOverflow(%d, %d, %d) = %v, want %v", tc.size, tc.block, tc.offset, got, tc.want)
			}
		})
	}
}

// TestVerify makes sure that the statistic notices objects it does not know.
// It belongs to the cache-line analysis (docs/redesign): a table that
// silently misses a kind of object would flatter the index.
func TestVerify(t *testing.T) {
	s := &stat{keysInObjects: 99}
	if err := s.verify("case", 100); err == nil || !strings.Contains(err.Error(), "99 keys") {
		t.Errorf("verify = %v, want an error about 99 keys", err)
	}
	if err := s.verify("case", 99); err != nil {
		t.Errorf("verify = %v, want none", err)
	}
}

// TestHuman makes sure that sizes read like the benchmark's own: 4K, 256K, 1M,
// and a plain number where none of these fits (the corpus maxima).
func TestHuman(t *testing.T) {
	for n, want := range map[int]string{4096: "4K", 262144: "256K", 1048576: "1M", 212449: "212449", 2097152: "2M"} {
		if got := human(n); got != want {
			t.Errorf("human(%d) = %q, want %q", n, got, want)
		}
	}
}

// TestRunEntries makes sure that the statistic of entries with several values,
// which the redesign's step 3 needs to size the single-key page, reports for a
// benchmark case how many entries would fit an object of 512 bytes for
// headers of different capacity. It belongs to the cache-line analysis
// (docs/redesign). The skewed value profile has 3% of its keys with 17 or more
// values: those never fit, the other columns grow with the header's capacity,
// and with 8-byte values and short integer keys the bytes never decide.
func TestRunEntries(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "u64", "-values", "natural", "-strvals=false", "-sizes", "16384", "-entries"}); err != nil {
		t.Fatal(err)
	}
	_, table, ok := strings.Cut(out.String(), "\n\n")
	if !ok {
		t.Fatalf("no entry table in\n%s", out.String())
	}
	lines := strings.Split(strings.TrimSpace(table), "\n")
	if len(lines) != 3 {
		t.Fatalf("%d lines in the entry table, want a header, a rule and a row:\n%s", len(lines), table)
	}
	head := strings.Split(strings.Trim(lines[0], "| "), " | ")
	row := strings.Split(strings.Trim(lines[2], "| "), " | ")
	if len(head) != len(row) || row[0] != "u64 natural 16K" {
		t.Fatalf("header %v and row %v do not match", head, row)
	}
	last := -1.0
	for i, h := range head {
		if !strings.HasPrefix(h, "fits N=") {
			continue
		}
		var share float64
		if _, err := fmt.Sscanf(row[i], "%f %%", &share); err != nil {
			t.Fatalf("%s: %q: %v", h, row[i], err)
		}
		if share < last || share > 100-3 {
			t.Errorf("%s: %.1f %%, after %.1f %%: should not fall, and the 17+ values never fit", h, share, last)
		}
		last = share
	}
	if last < 50 {
		t.Errorf("only %.1f %% of the entries fit the widest header, want most of them", last)
	}
}
