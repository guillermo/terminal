// Package ansi converts a given area into a stream of ansi sequences to be dump into a terminal
package ansi

import (
	"bytes"
	"github.com/guillermo/terminal/area"
	"github.com/guillermo/terminal/char"
)

// Sequence returns the ansi sequences to represent the area
// It will ignore any Char in the area that is nil
func Sequence(a *area.Area) []byte {
	rows, cols := a.Size()
	b := &bytes.Buffer{}
	c := &cursor{rows: rows, cols: cols}
	a.Each(func(Row, Col int, char char.Charer) {
		if char == nil {
			return
		}
		b.Write(c.draw(Row, Col, char))
	})
	return b.Bytes()
}
