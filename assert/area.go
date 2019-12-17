package assert

import (
	"github.com/guillermo/terminal/area"
	"github.com/guillermo/terminal/char"
	"testing"
)

type tArea struct {
	area.Area
	t *testing.T
}

func (a *tArea) SetChar(Row, Col int, ch string) {
	a.Set(Row, Col, &char.Char{Value: ch})
}

func (a *tArea) AssertSize(rows, cols int) {
	a.t.Helper()
	r, c := a.Size()
	if r != rows {
		a.t.Errorf("Expecting %d rows. Got %d.", rows, r)
	}
	if c != cols {
		a.t.Errorf("Expecing %d cols. Got %d", cols, c)
	}
}

func (a *tArea) AssertLines(expectation ...string) {
	a.t.Helper()
	actual := a.Lines()

	// Compare size
	if len(actual) != len(expectation) {
		a.t.Errorf("Expected %d lines, Got %d", len(expectation), len(actual))
	}
	// Compare content
	lines := len(actual)
	if len(expectation) > lines {
		lines = len(expectation)
	}
	for i := 0; i < lines; i++ {
		if i >= len(actual) {
			// Actual and not expecation
			a.t.Errorf("Expected line %d to be %q. Got nothing.", i+1, expectation[i])
			continue
		}
		if i >= len(expectation) {
			a.t.Errorf("Expected line %d to not exists. Got %q.", i+1, actual[i])
			continue
		}
		if actual[i] != expectation[i] {
			a.t.Errorf("Expected line %d to be %q. Got: %q.", i+1, expectation[i], actual[i])
		}
	}

}

func (a *tArea) Lines() []string {
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

func (a *tArea) AssertFixed() {
	if !a.Fixed {
		a.t.Error("Expecting the area to be Fixed.")
	}
}

func (a *tArea) AssertNotFixed() {
	if a.Fixed {
		a.t.Error("Expecting the area to not be Fixed.")
	}
}
