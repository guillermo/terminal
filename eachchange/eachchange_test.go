package eachchange

import (
	"fmt"
	"testing"

	"github.com/guillermo/reacty/terminal/area"
)

type c string

func (ch c) Equal(ch2 Char) bool {
	return string(ch) == string(ch2.(c))
}

type change struct {
	Row, Col int
	ch1, ch2 string
}

type changes []change

func (c changes) shouldHave(t *testing.T, changes int) {
	t.Helper()
	if c == nil {
		t.Fatal("Expecting changes to exists")
	}
	if len(c) != changes {
		t.Fatalf("Expecting %d changes. Got: %d.", changes, len(c))
	}
}

func (c change) shouldBe(t *testing.T, row, col int, ch1, ch2 string) {
	t.Helper()
	if row != c.Row || col != c.Col || ch1 != c.ch1 || ch2 != c.ch2 {
		exp := change{row, col, ch1, ch2}
		t.Errorf("Expecting change to be %v. Got %v.", exp, c)
	}
}

func (c change) String() string {
	return fmt.Sprintf("[%d,%d] (%q=>%q)", c.Row, c.Col, c.ch1, c.ch2)
}

func Changes(a1, a2 *area.Area) changes {
	changes := changes{}
	EachChange(a1, a2, func(row, col int, ch1, ch2 area.Char) {
		chg := change{Row: row, Col: col}
		if ch1 == nil {
			chg.ch1 = "nil"
		} else {
			chg.ch1 = string(ch1.(c))
		}
		if ch2 == nil {
			chg.ch2 = "nil"
		} else {
			chg.ch2 = string(ch2.(c))
		}
		changes = append(changes, chg)
	})
	return changes
}

func TestEachChange(t *testing.T) {

	a1 := &area.Area{Rows: 2, Cols: 2}
	a2 := &area.Area{Rows: 2, Cols: 2}

	Changes(a1, a2).shouldHave(t, 0)
	a1.Set(1, 1, c("a"))
	chgs := Changes(a1, a2)
	chgs.shouldHave(t, 1)
	chgs[0].shouldBe(t, 1, 1, "a", "nil")

	a2.Set(1, 1, c("a"))
	chgs = Changes(a1, a2)
	chgs.shouldHave(t, 0)

	a2.Set(1, 1, c("b"))
	chgs = Changes(a1, a2)
	chgs.shouldHave(t, 1)
	chgs[0].shouldBe(t, 1, 1, "a", "b")

}

func shouldEqual(t *testing.T, a *area.Area, exp ...string) {
	rows, _ := a.Size()
	if rows != len(exp) {
		t.Fatalf("Expecint %d rows. Got %d.", len(exp), rows)
	}

	lines := make([]string, rows)
	a.Each(func(row, col int, ch area.Char) {
		if ch == nil {
			lines[row-1] += "#"
			return
		}
		lines[row-1] += string(ch.(c))
	})

	for i, line := range lines {
		if line != exp[i] {
			t.Errorf("Expecting line %d to be %q. Got %q", i, exp[i], line)
		}

	}

}

func TestDiff(t *testing.T) {
	a1 := &area.Area{Rows: 2, Cols: 2, Fixed: true}
	a2 := &area.Area{Rows: 2, Cols: 2, Fixed: true}

	diff := Diff(a1, a2)
	shouldEqual(t, diff, "##", "##")

	a1.Set(1, 1, c("a"))
	diff = Diff(a1, a2)
	shouldEqual(t, diff, "a#", "##")

	a2.Set(1, 1, c("a"))
	diff = Diff(a1, a2)
	shouldEqual(t, diff, "##", "##")

}
