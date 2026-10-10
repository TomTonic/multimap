package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// field is one column of a row of a SQL dump: a number, NULL or a quoted
// string, which is already unescaped.
type field struct {
	b      []byte
	quoted bool
}

// eachRow calls fn with the columns of every row of the INSERT statements for
// table in the SQL dump r (the format of mysqldump and the Wikimedia dumps:
// "INSERT INTO `table` VALUES" and then rows in brackets, separated by commas).
// It exists so that mkcorpora needs no SQL library for three tables: titles are
// quoted strings with backslash escapes, which a split on commas would get
// wrong. The columns are only valid until fn returns. A row stands on a line
// of its own, as in the dumps; several rows on one line work as well.
func eachRow(r io.Reader, table string, fn func([]field) error) error {
	br := bufio.NewReaderSize(r, 1<<20)
	prefix := []byte("INSERT INTO `" + table + "` VALUES")
	var p rowParser
	inInsert := false
	for {
		line, err := br.ReadSlice('\n')
		if errors.Is(err, bufio.ErrBufferFull) {
			return errors.New("a line of the dump is longer than 1 MiB")
		}
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if bytes.HasPrefix(line, prefix) {
			inInsert, line = true, line[len(prefix):]
		}
		if inInsert {
			if perr := p.rows(line, fn); perr != nil {
				return perr
			}
			inInsert = !bytes.HasSuffix(bytes.TrimRight(line, " \r\n"), []byte(";"))
		}
		if err != nil {
			return nil // io.EOF
		}
	}
}

// rowParser holds the buffers that parsing reuses from row to row.
type rowParser struct {
	cols    []field
	scratch []byte
	ends    []int    // for each column of the row its end in scratch if quoted, else -1
	raws    [][]byte // for each column of the row its text if not quoted
}

// rows parses the rows on one line of a dump and calls fn for each.
func (p *rowParser) rows(s []byte, fn func([]field) error) error {
	for {
		s = bytes.TrimLeft(s, " ,\r\n")
		if len(s) == 0 || s[0] == ';' {
			return nil
		}
		if s[0] != '(' {
			return fmt.Errorf("unexpected %q in a row of the dump", s[:min(20, len(s))])
		}
		p.cols, p.scratch, p.ends, p.raws = p.cols[:0], p.scratch[:0], p.ends[:0], p.raws[:0]
		var err error
		if s, err = p.row(s[1:]); err != nil {
			return err
		}
		if err := fn(p.cols); err != nil {
			return err
		}
	}
}

// row parses the columns of one row, after its opening bracket, and returns
// what follows the closing bracket. String columns are copied unescaped into
// p.scratch; the columns are cut at the end, as scratch may move while it grows.
func (p *rowParser) row(s []byte) ([]byte, error) {
	for {
		if len(s) == 0 {
			return nil, errors.New("a row of the dump has no closing bracket")
		}
		var err error
		if s[0] == '\'' {
			s, err = p.quoted(s[1:])
			if err != nil {
				return nil, err
			}
			p.ends, p.raws = append(p.ends, len(p.scratch)), append(p.raws, nil)
		} else {
			i := bytes.IndexAny(s, ",)")
			if i < 0 {
				return nil, errors.New("a row of the dump has no closing bracket")
			}
			p.ends, p.raws, s = append(p.ends, -1), append(p.raws, s[:i]), s[i:]
		}
		if len(s) == 0 {
			return nil, errors.New("a row of the dump has no closing bracket")
		}
		closing := s[0] == ')'
		s = s[1:]
		if closing {
			break
		}
	}
	start := 0
	for i, end := range p.ends {
		if end < 0 {
			p.cols = append(p.cols, field{b: p.raws[i]})
			continue
		}
		p.cols = append(p.cols, field{b: p.scratch[start:end:end], quoted: true})
		start = end
	}
	return s, nil
}

// quoted appends the string that starts after an opening quote to p.scratch,
// unescaped, and returns what follows the closing quote.
func (p *rowParser) quoted(s []byte) ([]byte, error) {
	for len(s) > 0 {
		c := s[0]
		switch {
		case c == '\'' && len(s) > 1 && s[1] == '\'': // a doubled quote
			p.scratch, s = append(p.scratch, '\''), s[2:]
		case c == '\'':
			return s[1:], nil
		case c == '\\' && len(s) > 1:
			p.scratch, s = append(p.scratch, unescape(s[1])), s[2:]
		case c == '\\':
			return nil, errors.New("a string of the dump ends in a backslash")
		default:
			p.scratch, s = append(p.scratch, c), s[1:]
		}
	}
	return nil, errors.New("a string of the dump has no closing quote")
}

// unescape returns the byte that the backslash escape \e stands for in MySQL
// strings; any other escaped byte (\' \" \\ \% \_ ...) stands for itself.
func unescape(e byte) byte {
	switch e {
	case '0':
		return 0
	case 'n':
		return '\n'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case 'b':
		return '\b'
	case 'Z':
		return 26
	}
	return e
}
