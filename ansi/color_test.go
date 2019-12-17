package ansi

import (
	"image/color"
	"testing"
)

func TestColor(t *testing.T) {
	white := fromColor(color.White)
	if white.R != 255 ||
		white.G != 255 ||
		white.B != 255 {
		t.Fatal("Expected Black to equal Black")
	}
}
