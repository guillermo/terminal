package ansi

import (
	"image/color"
)

// termColor holds the RGB information
type termColor struct {
	R, G, B uint8
}

func fromColor(c color.Color) termColor {
	r, g, b, _ := c.RGBA()
	padding := uint32(8)
	return termColor{
		uint8(r >> padding),
		uint8(g >> padding),
		uint8(b >> padding),
	}
}
