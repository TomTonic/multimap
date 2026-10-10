package main

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

// rowsOf collects the rows that eachRow reads for table from dump, every
// column as text, string columns in quotes.
func rowsOf(dump, table string) ([][]string, error) {
	var out [][]string
	err := eachRow(strings.NewReader(dump), table, func(r []field) error {
		row := make([]string, len(r))
		for i, f := range r {
			row[i] = string(f.b)
			if f.quoted {
				row[i] = "'" + row[i] + "'"
			}
		}
		out = append(out, row)
		return nil
	})
	return out, err
}

// TestEachRow makes sure that the corpus builder reads the Wikipedia SQL dumps
// the way a database would: titles with quotes, backslashes, commas and
// brackets come out as they were typed, and nothing is taken from tables or
// statements that are not asked for. It covers eachRow and its parser, which
// feed every row of the page, linktarget and pagelinks tables to mkcorpora.
func TestEachRow(t *testing.T) {
	const head = "INSERT INTO `t` VALUES\n"
	tests := []struct {
		name string
		dump string
		want [][]string
	}{
		{"reads numbers, NULL and strings of one row", head + "(1,0,'April',0.5,NULL);\n", [][]string{{"1", "0", "'April'", "0.5", "NULL"}}},
		{"reads one row a line", head + "(1,'a'),\n(2,'b'),\n(3,'c');\n", [][]string{{"1", "'a'"}, {"2", "'b'"}, {"3", "'c'"}}},
		{"reads several rows on one line", head + "(1,'a'),(2,'b');\n", [][]string{{"1", "'a'"}, {"2", "'b'"}}},
		{"reads rows on the line of the INSERT", "INSERT INTO `t` VALUES (1,'a'),(2,'b');\n", [][]string{{"1", "'a'"}, {"2", "'b'"}}},
		{"unescapes quotes and backslashes", head + `(1,'it\'s \"x\" \\ it''s');` + "\n", [][]string{{"1", `'it's "x" \ it's'`}}},
		{"unescapes control characters", head + `(1,'a\nb\tc\0d\rf\bg\Zh');` + "\n", [][]string{{"1", "'a\nb\tc\x00d\rf\bg\x1ah'"}}},
		{"keeps commas and brackets inside strings", head + "(1,'a,b),(c'),\n(2,'');\n", [][]string{{"1", "'a,b),(c'"}, {"2", "''"}}},
		{"keeps UTF-8 titles", head + "(1,'Zürich_𝗙𝗮');\n", [][]string{{"1", "'Zürich_𝗙𝗮'"}}},
		{"reads the last line without a line break", head + "(1,'a')", [][]string{{"1", "'a'"}}},
		{"ignores other statements and other tables", "-- comment\nCREATE TABLE `t` (\n  `a` int\n);\nINSERT INTO `u` VALUES\n(9,'no');\nINSERT INTO `t` VALUES\n(1,'a');\nLOCK TABLES `t`;\n(2,'no')\n", [][]string{{"1", "'a'"}}},
		{"reads every INSERT statement of the table", head + "(1,'a');\n" + head + "(2,'b');\n", [][]string{{"1", "'a'"}, {"2", "'b'"}}},
		{"finds no rows in a dump without any", "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rowsOf(tt.dump, "t")
			if err != nil || !slices.EqualFunc(got, tt.want, slices.Equal) {
				t.Errorf("rows %q, error %v; want %q", got, err, tt.want)
			}
		})
	}
}

// TestEachRowFailures makes sure that a damaged dump stops the corpus builder
// instead of giving it a corpus with silently wrong titles. It covers the error
// paths of eachRow: rows and strings that end too early, text where a row
// should start, a line that is too long, and an error of the caller's function.
func TestEachRowFailures(t *testing.T) {
	const head = "INSERT INTO `t` VALUES\n"
	stop := errors.New("stop")
	tests := []struct {
		name string
		dump string
	}{
		{"fails on a row without a closing bracket", head + "(1,'a'\n"},
		{"fails on a number without a closing bracket", head + "(1"},
		{"fails on a string without a closing quote", head + "(1,'a);\n"},
		{"fails on a string that ends in a backslash", head + "(1,'a\\"},
		{"fails on text where a row should start", head + "(1,'a'),\nfoo\n"},
		{"fails on a line longer than the buffer", head + "(1,'" + strings.Repeat("a", 2<<20) + "');\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if rows, err := rowsOf(tt.dump, "t"); err == nil {
				t.Errorf("rows %q without an error", rows)
			}
		})
	}
	t.Run("returns the error of the function for a row", func(t *testing.T) {
		err := eachRow(strings.NewReader(head+"(1);\n"), "t", func([]field) error { return stop })
		if !errors.Is(err, stop) {
			t.Errorf("error %v, want %v", err, stop)
		}
	})
	t.Run("returns the error of the reader", func(t *testing.T) {
		if err := eachRow(errReader{}, "t", func([]field) error { return nil }); err == nil {
			t.Error("a broken reader gave no error")
		}
	})
}

// errReader fails every read.
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errors.New("broken") }
