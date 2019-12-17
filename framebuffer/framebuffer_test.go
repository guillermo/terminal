// Package output implements and area
package framebuffer

import (
	"github.com/guillermo/terminal/char"
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

func c(a string) *char.Char {
	return &char.Char{Value: a}
}
