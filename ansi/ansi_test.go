package ansi

import (
	"github.com/guillermo/terminal/area"
	"github.com/guillermo/terminal/char"
	"image/color"
	"testing"
)

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

	a.Set(2, 2, &char.Char{Value: "!"})
	is("\x1b[2;2H!")

	a.Set(2, 2, &char.Char{Value: "😎", IsBold: true})
	is("\x1b[2;2H\x1b[1m😎")

	a.Set(2, 2, &char.Char{Value: "😎", IsBold: true, BackgroundColor: color.White})
	is("\x1b[2;2H\x1b[1m\x1b[48;2;255;255;255m😎")
	a.Set(2, 2, &char.Char{Value: "😎", IsBold: true, ForegroundColor: color.White})
	is("\x1b[2;2H\x1b[1m\x1b[38;2;255;255;255m😎")

	a.Set(2, 3, &char.Char{Value: "h", IsBold: true})
	is("\x1b[2;2H\x1b[1m\x1b[38;2;255;255;255m😎\x1b[2;3H\x1b[39mh")
	a.Set(2, 2, &char.Char{Value: "e"})
	is("\x1b[2;2He\x1b[1mh")
	a.Set(2, 2, &char.Char{})
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
	a.Set(1, 1, &char.Char{Value: "H", IsBold: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[1mH\x1b[0me")
}

func TestFaint(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsFaint: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[2mH\x1b[0me")
}

func TestItalic(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsItalic: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[3mH\x1b[0me")
}

func TestUnderline(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsUnderline: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[4mH\x1b[0me")
}
func TestBlink(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsBlink: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[5mH\x1b[0me")
}
func TestInverse(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsInverse: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[7mH\x1b[0me")
}
func TestInvisible(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsInvisible: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[8mH\x1b[0me")
}
func TestCrossed(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsCrossed: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[9mH\x1b[0me")
}
func TestDouble(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H", IsDouble: true})
	a.Set(1, 2, &char.Char{Value: "e"})
	is(t, a, "\x1b[1;1H\x1b[21mH\x1b[0me")
}

func TestAllAttributes(t *testing.T) {
	a := &area.Area{}
	a.Set(1, 1, &char.Char{Value: "H",
		IsBold:      true,
		IsFaint:     true,
		IsItalic:    true,
		IsUnderline: true,
		IsBlink:     true,
		IsInverse:   true,
		IsCrossed:   true,
		IsDouble:    true,
	})
	a.Set(1, 2, &char.Char{Value: "e", IsBold: true})
	a.Set(1, 3, &char.Char{Value: "l"})
	is(t, a, "\x1b[1;1H\x1b[1;2;3;4;5;7;9;21mH\x1b[0;1me\x1b[0ml")
}
