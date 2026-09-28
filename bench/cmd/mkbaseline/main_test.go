package main

import (
	"archive/tar"
	"bytes"
	"maps"
	"slices"
	"testing"
)

// TestLibrary makes sure a head-to-head bench compares against exactly the
// library of the earlier commit. It covers cmd/mkbaseline, which copies the
// library out of a git archive: it keeps the Go files of the module root and
// of internal/ without their tests, leaves out everything else (the bench
// among it), and points the copy's imports of the library's own packages to
// the copy, while other imports stay as they are.
func TestLibrary(t *testing.T) {
	var buf bytes.Buffer
	w := tar.NewWriter(&buf)
	for name, body := range map[string]string{
		"ordered.go":             `import "github.com/TomTonic/multimap/internal/art"`,
		"ordered_test.go":        "test",
		"internal/art/lookup.go": `import "github.com/TomTonic/multimap/internal/swar"; import "github.com/TomTonic/Set3"`,
		"bench/cmd/bench/x.go":   "bench",
		"README.md":              "readme",
	} {
		if err := w.WriteHeader(&tar.Header{Name: name, Mode: 0o644, Size: int64(len(body))}); err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	files, err := library(tar.NewReader(&buf))
	if err != nil {
		t.Fatal(err)
	}
	if got := slices.Sorted(maps.Keys(files)); !slices.Equal(got, []string{"internal/art/lookup.go", "ordered.go"}) {
		t.Fatalf("copied %v, want the library's Go files without tests", got)
	}
	if want := `import "github.com/TomTonic/multimap/bench/baseline/internal/swar"; import "github.com/TomTonic/Set3"`; files["internal/art/lookup.go"] != want {
		t.Fatalf("rewrote imports to %q, want %q", files["internal/art/lookup.go"], want)
	}
	if _, err := library(tar.NewReader(bytes.NewReader(nil))); err == nil {
		t.Fatal("an archive without the library gave no error")
	}
}
