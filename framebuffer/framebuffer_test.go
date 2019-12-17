// Package output implements and area
package framebuffer

import (
	"github.com/guillermo/reacty/terminal/eachchange"
	"image/color"
	"testing"
)

func TestFramebuffer_Changes(t *testing.T) {

	fb := &Framebuffer{}

	equal := func(s string) {
		t.Helper()

		out := fb.Changes()
		if s != string(out) {
			t.Errorf("Expecting changes to be %q. Got: %q.", s, string(out))
		}
	}

	equal("")

	fb.Set(1, 1, c("a"))
	equal("\x1b[1;1Ha")
	equal("")

}

type c string

func (ch c) Content() string {
	return string(ch)
}

func (ch c) Background() color.Color {
	return nil
}

func (ch c) Foreground() color.Color {
	return nil
}

func (ch c) Bold() bool {
	return false
}

func (ch c) Faint() bool {
	return false
}

func (ch c) Italic() bool {
	return false
}

func (ch c) Underline() bool {
	return false
}

func (ch c) Blink() bool {
	return false
}

func (ch c) Inverse() bool {
	return false
}

func (ch c) Invisible() bool {
	return false
}

func (ch c) Crossed() bool {
	return false
}

func (ch c) Double() bool {
	return false
}
func (ch c) Equal(ch2 eachchange.Char) bool {
	if string(ch) != string(ch2.(c)) {
		return false
	}
	return true
}
