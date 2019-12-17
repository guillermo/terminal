package assert

import (
	"testing"
)

func TestTArea(t *testing.T) {
	a := TArea{T: t}
	a.AssertSize(0, 0)
	a.AssertLines()

	a.SetChar(1, 1, "a")
	a.AssertSize(1, 1)
	a.AssertLines("a")

	a.SetChar(2, 2, "b")
	a.AssertSize(2, 2)
	a.AssertLines("a ", " b")

	a.AssertNotFixed()
	a.Fixed = true
	a.AssertFixed()
}
