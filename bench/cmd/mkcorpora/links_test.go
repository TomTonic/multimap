package main

import (
	"errors"
	"io"
	"slices"
	"strings"
	"testing"
)

// miniDump opens the tables of a small wiki: the pages Alpha (links to Beta,
// Gamma and the not existing Red), Beta (links to Alpha only) and Moved (a
// redirect to Alpha), a talk page that links to Alpha, and links to a target
// of another namespace.
func miniDump(table string) (io.ReadCloser, error) {
	dumps := map[string]string{
		"page": "INSERT INTO `page` VALUES\n" +
			"(1,0,'Alpha',0,0,0.5,'20260101000000',NULL,10,100,'wikitext',NULL),\n" +
			"(2,0,'Beta',0,0,0.5,'20260101000000',NULL,10,100,'wikitext',NULL),\n" +
			"(3,0,'Gamma',0,0,0.5,'20260101000000',NULL,10,100,'wikitext',NULL),\n" +
			"(4,0,'Moved',1,0,0.5,'20260101000000',NULL,10,100,'wikitext',NULL),\n" +
			"(5,1,'Alpha',0,0,0.5,'20260101000000',NULL,10,100,'wikitext',NULL);\n",
		"linktarget": "INSERT INTO `linktarget` VALUES\n" +
			"(10,0,'Alpha'),\n(11,0,'Beta'),\n(12,0,'Gamma'),\n(13,0,'Red'),\n(14,1,'Alpha'),\n(15,0,'Don\\'t');\n",
		"pagelinks": "INSERT INTO `pagelinks` VALUES\n" +
			"(1,0,11),\n(1,0,12),\n(1,0,13),\n(1,0,14),\n(2,0,10),\n(4,0,10),\n(5,1,10);\n",
	}
	return io.NopCloser(strings.NewReader(dumps[table])), nil
}

// TestReadLinks makes sure the page links corpus is made of what a reader of
// Wikipedia calls a link from article to article: only pages and targets of the
// main namespace count, a redirect is a page like any other, and a page without
// a link to the main namespace is no key. It covers readLinks on a tiny dump
// that has all of these cases.
func TestReadLinks(t *testing.T) {
	g, err := readLinks(miniDump)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, k := range g.keys {
		var ts []string
		for _, id := range k.targets {
			ts = append(ts, g.targets[id])
		}
		got = append(got, k.title+map[bool]string{false: "", true: "(redirect)"}[k.redirect]+">"+strings.Join(ts, ","))
	}
	want := []string{"Alpha>Beta,Gamma,Red", "Beta>Alpha", "Moved(redirect)>Alpha"}
	if !slices.Equal(got, want) {
		t.Errorf("keys %q, want %q", got, want)
	}
	if !g.exists["Beta"] || g.exists["Red"] || len(g.targets) != 5 {
		t.Errorf("existing pages %v, %d targets of namespace 0; want Beta but not Red, 5", g.exists, len(g.targets))
	}
}

// TestReadLinksFailures makes sure that data that contradicts itself stops the
// corpus builder. It covers the errors of readLinks: a link from a page that the
// page table does not know, damaged rows in each of the three tables, and a
// table that cannot be opened.
func TestReadLinksFailures(t *testing.T) {
	replace := func(table, dump string) func(string) (io.ReadCloser, error) {
		return func(name string) (io.ReadCloser, error) {
			if name == table {
				return io.NopCloser(strings.NewReader(dump)), nil
			}
			return miniDump(name)
		}
	}
	tests := []struct {
		name string
		open func(string) (io.ReadCloser, error)
	}{
		{"fails on a link from an unknown page", replace("pagelinks", "INSERT INTO `pagelinks` VALUES\n(99,0,10);\n")},
		{"fails on a damaged pagelinks row", replace("pagelinks", "INSERT INTO `pagelinks` VALUES\n(1,0,'x');\n")},
		{"fails on a pagelinks row with a namespace that is no number", replace("pagelinks", "INSERT INTO `pagelinks` VALUES\n(1,'x',10);\n")},
		{"fails on a damaged linktarget row", replace("linktarget", "INSERT INTO `linktarget` VALUES\n(1);\n")},
		{"fails on a damaged page row", replace("page", "INSERT INTO `page` VALUES\n(1,0,'A','x');\n")},
		{"fails on a page row with too few columns", replace("page", "INSERT INTO `page` VALUES\n(1,0,'A');\n")},
		{"fails on a table that cannot be opened", func(string) (io.ReadCloser, error) { return nil, errors.New("offline") }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := readLinks(tt.open); err == nil {
				t.Error("no error")
			}
		})
	}
}

// TestFormatLinks makes sure the corpus file says what the README promises:
// every page that links is a key, in ascending order, each with all its links as
// ascending indexes into the ascending list of target titles, and a title that
// cannot be stored stops the builder. It covers targetTitles, formatLinks,
// linkStats and checkTitle on the tiny dump.
func TestFormatLinks(t *testing.T) {
	g, err := readLinks(miniDump)
	if err != nil {
		t.Fatal(err)
	}
	titles, err := targetTitles(g)
	if err != nil {
		t.Fatal(err)
	}
	var w strings.Builder
	formatLinks(&w, g, titles)
	const want = "4\nAlpha\nBeta\nGamma\nRed\n" + "Alpha\t1,2,3\nBeta\t0\nMoved\t0\n"
	if w.String() != want {
		t.Errorf("file %q, want %q", w.String(), want)
	}
	stats := linkStats(g, titles)
	for _, s := range []string{"3 keys, 5 values, 4 target titles (1 of them no page", "max 3", "redirects 33.3 %"} {
		if !strings.Contains(stats, s) {
			t.Errorf("statistics lack %q:\n%s", s, stats)
		}
	}

	for _, bad := range []string{"", "a\tb", "a\nb", "a\xffb"} {
		g.targets[10] = bad
		if _, err := targetTitles(g); err == nil {
			t.Errorf("target title %q was accepted", bad)
		}
	}
	g.targets[10] = "Alpha"
	g.keys[0].title = "x\ty"
	if _, err := targetTitles(g); err == nil {
		t.Error("a key title with a tab was accepted")
	}
}
