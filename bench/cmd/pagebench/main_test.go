package main

import (
	"bytes"
	"flag"
	"strings"
	"testing"
)

// TestRun makes sure that the page layout comparison reports what the redesign's
// step 3 needs from it: for each real data set the four comparisons of lookups
// and a table of the memory the pages take per key, and a typo in the kind of
// keys is reported. It belongs to the multi-key page candidates (docs/redesign),
// whose layouts the tool compares on the keys of street and dirs.
func TestRun(t *testing.T) {
	for name, v := range map[string]string{"validation": "2", "repeats": "11"} {
		old := flag.Lookup(name).Value.String()
		if err := flag.Set(name, v); err != nil {
			t.Fatal(err)
		}
		defer func() { _ = flag.Set(name, old) }()
	}
	var out bytes.Buffer
	if err := run(&out, "street", 2000); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"## street: ", "### point lookups of present keys: tags against lengths", "### point lookups of absent keys", "names copied to strings", "| directory of tags (vpage), numbers |", "| length header (lpage), names |"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("the output has no %q:\n%s", want, out.String())
		}
	}
	if err := run(&out, "uuid", 10); err == nil {
		t.Error("a kind without natural values was accepted")
	}
}
