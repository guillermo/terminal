// Package framebuffer stores the current terminal state and return the ansi
// sequences require to transform the current state to the new one
package framebuffer

import (
	"github.com/guillermo/reacty/terminal/ansi"
	"github.com/guillermo/reacty/terminal/area"
	"github.com/guillermo/reacty/terminal/eachchange"
)

// Framebuffer is an Area that retains the state since the last time Changes was
// called. It is responsability of the implementator to call Changes()
// periodically (as of 60 frames per second) or sync manually after each change.
type Framebuffer struct {
	area.Area
	area2 area.Area
}

// Changes will return the ansi sequence
func (fb *Framebuffer) Changes() []byte {
	rows, cols := fb.Size()
	changes := &area.Area{Fixed: true, Rows: rows, Cols: cols}
	eachchange.EachChange(
		&fb.Area,
		&fb.area2,
		func(Row, Col int, a1Char, a2Char area.Char) {
			changes.Set(Row, Col, a1Char)
			fb.area2.Set(Row, Col, a1Char)
		})

	return ansi.Sequence(changes)
}
