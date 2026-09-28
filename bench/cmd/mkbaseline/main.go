// Command mkbaseline copies package multimap as of a git commit into
// package baseline of the bench, so that the bench can compare Ordered head
// to head with that earlier version of itself:
//
//	go run ./cmd/mkbaseline -ref main   # from the bench directory
//	go run -tags baseline ./cmd/bench -vs baseline
//
// It takes the library's Go files without tests (the module root and
// internal/), rewrites their imports of the library's own packages to the
// copy's, and records the commit in the constant BaselineRef. The copy is
// generated, not versioned (see .gitignore).
package main

import (
	"archive/tar"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	module = "github.com/TomTonic/multimap/"
	copyTo = module + "bench/baseline/"
)

func main() {
	ref := flag.String("ref", "main", "git commit, branch or tag to copy")
	out := flag.String("out", "baseline", "directory of the copy, replaced as a whole")
	flag.Parse()
	if err := run(*ref, *out); err != nil {
		fmt.Fprintln(os.Stderr, "mkbaseline:", err)
		os.Exit(1)
	}
}

func run(ref, out string) error {
	hash, err := git("rev-parse", "--short", ref+"^{commit}")
	if err != nil {
		return err
	}
	archive, err := git("archive", "--format=tar", ref)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(out); err != nil {
		return err
	}
	files, err := library(tar.NewReader(strings.NewReader(archive)))
	if err != nil {
		return err
	}
	files["ref.go"] = fmt.Sprintf("package multimap\n\n// BaselineRef is the commit this copy was made from.\nconst BaselineRef = %q\n",
		ref+" "+strings.TrimSpace(hash))
	for name, src := range files {
		path := filepath.Join(out, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
			return err
		}
	}
	fmt.Printf("copied %d files of %s (%s) to %s\n", len(files)-1, ref, strings.TrimSpace(hash), out)
	return nil
}

// git runs git on the repository around the bench and returns its output.
func git(args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", ".."}, args...)...) //nolint:gosec // fixed command, arguments from the command line
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	b, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, stderr.String())
	}
	return string(b), nil
}

// library returns the library's Go files without tests from the archive r,
// by path, with their imports rewritten (see rewrite).
func library(r *tar.Reader) (map[string]string, error) {
	files := map[string]string{}
	for {
		h, err := r.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		dir, name := filepath.Split(h.Name)
		if !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") ||
			dir != "" && !strings.HasPrefix(dir, "internal/") {
			continue
		}
		b, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		files[h.Name] = rewrite(string(b))
	}
	if len(files) == 0 {
		return nil, errors.New("the archive holds no Go files of the library")
	}
	return files, nil
}

// rewrite points the imports of the library's own packages in src to the
// copy's.
func rewrite(src string) string {
	return strings.ReplaceAll(src, `"`+module, `"`+copyTo)
}
