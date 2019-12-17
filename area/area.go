package area

import (
	"fmt"
)

// Char represent a given Character in an area
type Char interface{}

// Area represents a rectangular area of Chars.
// An empty Area is a valid one.
type Area struct {
	Rows, Cols int
	content    [][]Char
	Fixed      bool
}

// Size returns the current Size.
func (a *Area) Size() (rows, cols int) {
	return a.Rows, a.Cols
}

// SetSize changes the current area size an sets the area to Fixed.
func (a *Area) SetSize(rows, cols int) {
	a.Rows = rows
	a.Cols = cols
	a.Fixed = true
}

// Each iterates over each column and row. The rows and the columns starts with 1
func (a *Area) Each(fn func(Row, Col int, char Char)) {
	for r := 0; r < a.Rows; r++ {
		for c := 0; c < a.Cols; c++ {
			if len(a.content) <= r {
				fn(r+1, c+1, nil)
				continue
			}
			row := a.content[r]
			if len(row) <= c {
				fn(r+1, c+1, nil)
				continue
			}
			fn(r+1, c+1, row[c])
		}
	}
}

// Set will change the given Char in the row,col position.
// If the area is fixed, it will return an error if a char is being set outside the area
// A row or col smaller than 1 will also return an error
func (a *Area) Set(row, col int, c Char) error {
	if row <= 0 || (a.Fixed && row > a.Rows) {
		return fmt.Errorf("Invalid Row %d", row)
	}
	if col <= 0 || (a.Fixed && col > a.Cols) {
		return fmt.Errorf("Invalid Col %d", col)
	}

	for len(a.content) < row {
		a.content = append(a.content, make([]Char, col))
	}
	for len(a.content[row-1]) < col {
		a.content[row-1] = append(a.content[row-1], nil)
	}
	a.content[row-1][col-1] = c
	if a.Cols < col {
		a.Cols = col
	}
	if a.Rows < row {
		a.Rows = row
	}
	return nil
}

// Get will return the character at the given position.
// A row or col smaller than 1 will return nil and error.
// For a fixed area a row and col bigger than the actual size will also return an error
func (a *Area) Get(Row, Col int) (Char, error) {
	if Row <= 0 || (a.Fixed && Row > a.Rows) {
		return nil, fmt.Errorf("Invalid Row %d", Row)
	}
	if Col <= 0 || (a.Fixed && Col > a.Cols) {
		return nil, fmt.Errorf("Invalid Col %d", Col)
	}
	if len(a.content) < Row {
		return nil, nil
	}
	row := a.content[Row-1]
	if len(row) < Col {
		return nil, nil
	}
	return row[Col-1], nil
}
