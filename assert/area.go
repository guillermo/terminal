package assert

import (
	"github.com/guillermo/terminal/area"
	"github.com/guillermo/terminal/char"
	"testing"
)

// TArea is a testing Area
type TArea struct {
	area.Area
	T *testing.T
}

// SetChar is a shorthand to set a specific Char
func (a *TArea) SetChar(Row, Col int, ch string) {
	a.Set(Row, Col, &char.Char{Value: ch})
}

// AssertSize checks if the Area have the size given by rows and cols
func (a *TArea) AssertSize(rows, cols int) {
	a.T.Helper()
	r, c := a.Size()
	if r != rows {
		a.T.Errorf("Expecting %d rows. Got %d.", rows, r)
	}
	if c != cols {
		a.T.Errorf("Expecing %d cols. Got %d", cols, c)
	}
}

// AssertLines checks if the area.Lines maches the expectation
func (a *TArea) AssertLines(expectation ...string) {
	a.T.Helper()
	actual := a.Lines()

	// Compare size
	if len(actual) != len(expectation) {
		a.T.Errorf("Expected %d lines, Got %d", len(expectation), len(actual))
	}
	// Compare content
	lines := len(actual)
	if len(expectation) > lines {
		lines = len(expectation)
	}
	for i := 0; i < lines; i++ {
		if i >= len(actual) {
			// Actual and not expecation
			a.T.Errorf("Expected line %d to be %q. Got nothing.", i+1, expectation[i])
			continue
		}
		if i >= len(expectation) {
			a.T.Errorf("Expected line %d to not exists. Got %q.", i+1, actual[i])
			continue
		}
		if actual[i] != expectation[i] {
			a.T.Errorf("Expected line %d to be %q. Got: %q.", i+1, expectation[i], actual[i])
		}
	}

}

// Lines return the area as a slice of strings. Each string have the same characters as columns the area.
func (a *TArea) Lines() []string {
	rows, _ := a.Size()
	actual := make([]string, rows)
	// Get strings
	a.Each(func(r, c int, ch char.Charer) {
		if ch == nil {
			actual[r-1] += " "
		} else {
			actual[r-1] += ch.Content()
		}
	})
	return actual
}

// AssertFixed checks if the area is fixed
func (a *TArea) AssertFixed() {
	if !a.Fixed {
		a.T.Error("Expecting the area to be Fixed.")
	}
}

// AssertNotFixed checks if the area is not fixed
func (a *TArea) AssertNotFixed() {
	if a.Fixed {
		a.T.Error("Expecting the area to not be Fixed.")
	}
}
