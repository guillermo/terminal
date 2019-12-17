// Package ansi converts a given area into a stream of ansi sequences to be dump into a terminal
package ansi

import (
	"bytes"
	"github.com/guillermo/reacty/terminal/area"
	col "image/color"
)

// Char represents the properties a char in an Area should have to be able to be represented in a terminal.
type Char interface {
	Content() string
	Background() col.Color
	Foreground() col.Color
	Bold() bool
	Faint() bool
	Italic() bool
	Underline() bool
	Blink() bool
	Inverse() bool
	Invisible() bool
	Crossed() bool
	Double() bool
}

// Sequence returns the ansi sequences to represent the area
// It will ignore any Char in the area that is nil
func Sequence(a *area.Area) []byte {
	rows, cols := a.Size()
	b := &bytes.Buffer{}
	c := &cursor{rows: rows, cols: cols}
	a.Each(func(Row, Col int, char area.Char) {
		if char == nil {
			return
		}
		b.Write(c.draw(Row, Col, char.(Char)))
	})
	return b.Bytes()
}
