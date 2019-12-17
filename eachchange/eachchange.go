// Package eachchange will compare two areas
package eachchange

import (
	"github.com/guillermo/terminal/area"
	"github.com/guillermo/terminal/char"
)

// EachChange compares two areas and invoke the callback every time there is
// a difference. It only covers the area of area1.
func EachChange(area1, area2 *area.Area, fn func(Row, Col int, a1Char, a2Char char.Charer)) {
	area1.Each(func(Row, Col int, a1ch char.Charer) {
		a2ch, err := area2.Get(Row, Col)
		// Both empty
		if err == nil && a1ch == nil && a2ch == nil {
			return
		}

		// Call in case
		// * There is some error (as we don't know for sure)
		// * If any of them is nil
		if err != nil ||
			a1ch == nil || a2ch == nil {
			fn(Row, Col, a1ch, a2ch)
			return
		}

		if a1ch.Equal(a2ch) {
			return
		}
		fn(Row, Col, a1ch, a2ch)
	})
}

// Diff returns a new fixed area of the size of area1 with all the Char of area1
// that are not equal in area2
func Diff(area1, area2 *area.Area) *area.Area {
	rows, cols := area1.Size()
	changes := &area.Area{Fixed: true, Rows: rows, Cols: cols}
	EachChange(area1, area2, func(Row, Col int, a1Char, a2Char char.Charer) {
		changes.Set(Row, Col, a1Char)
	})
	return changes
}
