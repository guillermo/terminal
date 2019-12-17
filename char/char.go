// Package char represent a Char in a terminal
package char

import (
	"image/color"
)

// Charer represents the properties a char in an Area should have to be able to be represented in a terminal.
type Charer interface {
	Content() string
	Background() color.Color
	Foreground() color.Color
	Bold() bool
	Faint() bool
	Italic() bool
	Underline() bool
	Blink() bool
	Inverse() bool
	Invisible() bool
	Crossed() bool
	Double() bool
}

// C returns a char with the value as ch. It is just syntax sugar.
func C(ch string) *Char {
	return &Char{Value: ch}
}

// Char is a concrete implementation of the Charer interface.
type Char struct {
	Value           string
	ForegroundColor color.Color
	BackgroundColor color.Color
	IsBold          bool
	IsFaint         bool
	IsItalic        bool
	IsUnderline     bool
	IsBlink         bool
	IsInverse       bool
	IsInvisible     bool
	IsCrossed       bool
	IsDouble        bool
}

// Content returns the content of the char
func (c *Char) Content() string {
	return c.Value
}

// Background returns the backgroundColor of the char
func (c *Char) Background() color.Color {
	return c.BackgroundColor
}

// Foreground returns the foreground color of the char
func (c *Char) Foreground() color.Color {
	return c.ForegroundColor
}

// Bold returns true if the char is in bold
func (c *Char) Bold() bool {
	return c.IsBold
}

// Faint returns true if the char is in faint
func (c *Char) Faint() bool {
	return c.IsFaint
}

// Italic returns true if the char is in Italic
func (c *Char) Italic() bool {
	return c.IsItalic
}

// Underline returns true if the char is Underline
func (c *Char) Underline() bool {
	return c.IsUnderline
}

// Blink returns true if the char should be blinking
func (c *Char) Blink() bool {
	return c.IsBlink
}

// Inverse returns true if the char should be reverse
func (c *Char) Inverse() bool {
	return c.IsInverse
}

// Invisible returns true if the char should be invisible
func (c *Char) Invisible() bool {
	return c.IsInvisible
}

// Crossed returns true if the char should be crossed
func (c *Char) Crossed() bool {
	return c.IsCrossed
}

// Double returns true if the char should be double underline
func (c *Char) Double() bool {
	return c.IsDouble
}

func equalColor(c1, c2 color.Color) bool {
	if c1 == nil && c2 == nil {
		return true
	}
	if c1 == nil || c2 == nil {
		return false
	}

	r, g, b, a := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	if r != r2 ||
		g != g2 ||
		b != b2 ||
		a != a2 {
		return false
	}

	return true
}

// Equal returns true if c1 and c2 are equal
func Equal(c, c2 Charer) bool {
	if c == nil && c2 == nil {
		return true
	}

	if c == nil || c2 == nil {
		return false
	}

	if c.Content() != c2.Content() ||
		c.Bold() != c2.Bold() ||
		c.Faint() != c2.Faint() ||
		c.Italic() != c2.Italic() ||
		c.Underline() != c2.Underline() ||
		c.Blink() != c2.Blink() ||
		c.Inverse() != c2.Inverse() ||
		c.Invisible() != c2.Invisible() ||
		c.Crossed() != c2.Crossed() ||
		c.Double() != c2.Double() {
		return false
	}

	if !equalColor(c.Background(), c2.Background()) {
		return false
	}
	if !equalColor(c.Foreground(), c2.Foreground()) {
		return false
	}

	return true
}
