package ansi

import (
	"github.com/guillermo/reacty/terminal/area"
	col "image/color"
	"testing"
)

type char struct {
	c                                     string
	fg                                    col.Color
	bg                                    col.Color
	bold, faint, italic, underline, blink bool
	inverse, invisible, crossed, double   bool
}

func (c char) Content() string {
	return c.c
}
func (c char) Background() col.Color {
	return c.bg
}

func (c char) Foreground() col.Color {
	return c.fg
}

func (c char) Bold() bool {
	return c.bold
}
func (c char) Faint() bool {
	return c.faint
}
func (c char) Italic() bool {
	return c.italic
}
func (c char) Underline() bool {
	return c.underline
}
func (c char) Blink() bool {
	return c.blink
}
func (c char) Inverse() bool {
	return c.inverse
}
func (c char) Invisible() bool {
	return c.invisible
}
func (c char) Crossed() bool {
	return c.crossed
}
func (c char) Double() bool {
	return c.double
}

func is(t *testing.T, a *area.Area, exp string) {
	t.Helper()
	out := Sequence(a)
	if string(out) != exp {
		t.Errorf("Expecting %q, Got: %q", exp, string(out))
	}
}

func TestArea(t *testing.T) {
	a := &area.Area{}

	is := func(exp string) {
		t.Helper()
		out := Sequence(a)
		if string(out) != exp {
			t.Errorf("Expecting %q, Got: %q", exp, string(out))
		}
	}

	a.Set(2, 2, char{c: "!"})
	is("\x1b[2;2H!")

	a.Set(2, 2, char{c: "😎", bold: true})
	is("\x1b[2;2H\x1b[1m😎")

	a.Set(2, 2, char{c: "😎", bold: true, bg: col.White})
	is("\x1b[2;2H\x1b[1m\x1b[48;2;255;255;255m😎")
	a.Set(2, 2, char{c: "😎", bold: true, fg: col.White})
	is("\x1b[2;2H\x1b[1m\x1b[38;2;255;255;255m😎")

	a.Set(2, 3, char{c: "h", bold: true})
	is("\x1b[2;2H\x1b[1m\x1b[38;2;255;255;255m😎\x1b[2;3H\x1b[39mh")
	a.Set(2, 2, char{c: "e"})
	is("\x1b[2;2He\x1b[1mh")
	a.Set(2, 2, char{})
	is("\x1b[2;2H \x1b[2;3H\x1b[1mh")
	a.Set(2, 2, nil)
	is("\x1b[2;3H\x1b[1mh")
	/*
		sb.Set(1,1,':',Bold,Underline,Foreground(Red),Background(Blue))
		// Order is relevant
		sb.shouldOutput(t, "\x1b[1;1H\x1b[0;1;4m\x1b[38;2;255;0;0m\x1b[38;2;0;0;255m:")
		sb.Set(1,2,'O',Bold,Underline,Foreground(Red),Background(Blue))
		sb.shouldOutput(t, "O")
		sb.Set(2,1,'O',Underline,Foreground(Red),Background(Blue))
		sb.shouldOutput(t, "\x1b[0;4mO")
	*/
}

func TestBold(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", bold: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[1mH\x1b[0me")
}

func TestFaint(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", faint: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[2mH\x1b[0me")
}

func TestItalic(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", italic: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[3mH\x1b[0me")
}

func TestUnderline(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", underline: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[4mH\x1b[0me")
}
func TestBlink(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", blink: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[5mH\x1b[0me")
}
func TestInverse(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", inverse: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[7mH\x1b[0me")
}
func TestInvisible(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", invisible: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[8mH\x1b[0me")
}
func TestCrossed(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", crossed: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[9mH\x1b[0me")
}
func TestDouble(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H", double: true})
	a.Set(1, 2, char{c: "e"})
	is(t, a, "\x1b[1;1H\x1b[21mH\x1b[0me")
}

func TestAllAttributes(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, char{c: "H",
		bold:      true,
		faint:     true,
		italic:    true,
		underline: true,
		blink:     true,
		inverse:   true,
		crossed:   true,
		double:    true,
	})
	a.Set(1, 2, char{c: "e", bold: true})
	a.Set(1, 3, char{c: "l"})
	is(t, a, "\x1b[1;1H\x1b[1;2;3;4;5;7;9;21mH\x1b[0;1me\x1b[0ml")
}
