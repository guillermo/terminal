package area

import (
	"github.com/guillermo/terminal/char"
	"testing"
)

func c(s string) *char.Char {
	return &char.Char{Value: s}
}

func (a *Area) isEqual(t *testing.T, expectation ...string) {
	t.Helper()
	actual := make([]string, a.Rows)
	// Get strings
	a.Each(func(r, c int, ch char.Charer) {
		if ch == nil {
			actual[r-1] += " "
		} else {
			actual[r-1] += ch.Content()
		}
	})
	// Compare size
	if len(actual) != len(expectation) {
		t.Errorf("Expected %d lines, Got %d", len(expectation), len(actual))
	}
	// Compare content
	lines := len(actual)
	if len(expectation) > lines {
		lines = len(expectation)
	}
	for i := 0; i < lines; i++ {
		if i >= len(actual) {
			// Actual and not expecation
			t.Errorf("Expected line %d to be %q. Got nothing.", i+1, expectation[i])
			continue
		}
		if i >= len(expectation) {
			t.Errorf("Expected line %d to not exists. Got %q.", i+1, actual[i])
			continue
		}
		if actual[i] != expectation[i] {
			t.Errorf("Expected line %d to be %q. Got: %q.", i+1, expectation[i], actual[i])
		}
	}
}

func TestArea_Set(t *testing.T) {
	a := &Area{}
	a.isEqual(t)
	err := a.Set(2, 2, c("a"))
	if err != nil {
		t.Error(err)
	}
	rows, cols := a.Size()
	if rows != 2 || cols != 2 {
		t.Error("Expected a size of 2x2. Got:", rows, cols)
	}
	a.isEqual(t, "  ", " a")

	a = &Area{Rows: 2, Cols: 2, Fixed: true}
	err = a.Set(0, 0, c("a"))
	if err == nil {
		t.Error("Expected an error")
	}
	err = a.Set(1, 0, c("a"))
	if err == nil {
		t.Error("Expected an error")
	}
	a.isEqual(t, "  ", "  ")
	a.Set(1, 1, c("a"))
	a.isEqual(t, "a ", "  ")

	err = a.Set(2, 3, c("a"))
	if err == nil {
		t.Error("Expected an error while setting an out of bands char")
	}
	err = a.Set(3, 2, c("a"))
	if err == nil {
		t.Error("Expected an error while setting an out of bands char")
	}

	a.Set(1, 2, c("b"))
	a.isEqual(t, "ab", "  ")

}

func TestArea_Get(t *testing.T) {
	a := &Area{}
	af := &Area{Fixed: true, Rows: 2, Cols: 2}

	// Smaller than 1
	_, err := a.Get(0, 1)
	if err == nil {
		t.Error("Expecting an error")
	}
	_, err = a.Get(1, 0)
	if err == nil {
		t.Error("Expecting an error")
	}

	// It should not fail
	ch, err := a.Get(1, 1)
	if err != nil {
		t.Error(err)
	}
	if ch != nil {
		t.Error(ch)
	}

	// It should fail if outside bounds
	_, err = af.Get(2, 3)
	if err == nil {
		t.Error("It should fail in a fixed area")
	}
	_, err = af.Get(3, 2)
	if err == nil {
		t.Error("It should fail in a fixed area")
	}

	// It should return the valu
	a.Set(10, 10, c("h"))
	ch, err = a.Get(10, 10)
	if err != nil {
		t.Error("It should not fail")
	}
	if ch == nil || ch.Content() != "h" {
		t.Error("Expecting to get the same. Got", ch)
	}
	ch, err = a.Get(10, 11)
	if err != nil {
		t.Error("It should not fail")
	}
	if ch != nil {
		t.Error("Expecting to get the same. Got")
	}
}
