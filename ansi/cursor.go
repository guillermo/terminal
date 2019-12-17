package ansi

import (
	"bytes"
	col "image/color"

	"github.com/mattn/go-runewidth"
)

type cursor struct {
	rows  int
	cols  int
	row   int
	col   int
	fg    *termColor
	bg    *termColor
	style style
}

func (c *cursor) draw(Row, Col int, ch Char) []byte {
	b := &bytes.Buffer{}
	b.Write(c.move(Row, Col))
	b.Write(c.setStyle(ch))
	b.Write(c.setFgColor(ch.Foreground()))
	b.Write(c.setBgColor(ch.Background()))
	b.Write(c.writeChar(ch.Content()))
	return b.Bytes()
}

func (c *cursor) move(Row, Col int) []byte {
	b := &bytes.Buffer{}
	if c.row != Row || c.col != Col {
		b.Write(seq("GOTO", Row, Col))
		c.row = Row
		c.col = Col
	}
	return b.Bytes()
}

func (c *cursor) setStyle(ch Char) []byte {
	s := newStyle(ch)
	if s == c.style {
		return nil
	}
	defer func() {
		c.style = s
	}()
	// There is a change
	if s.Normal() {
		// Reset
		return seq("CHARSTYLE", "0")
	}
	if c.style.Normal() {
		return seq("CHARSTYLE", s.codes())
	}
	return seq("CHARSTYLE", "0;"+s.codes())
}

func (c *cursor) setFgColor(color col.Color) []byte {
	if c.fg == nil && color == nil {
		return nil
	}
	if color == nil {
		c.fg = nil
		return seq("RESETFGCOLOR")
	}

	fg := fromColor(color)
	if c.fg == &fg {
		return nil
	}

	c.fg = &fg
	return seq("FGCOLOR", fg.R, fg.G, fg.B)
}

func (c *cursor) setBgColor(color col.Color) []byte {
	if c.bg == nil && color == nil {
		return nil
	}
	if color == nil {
		c.bg = nil
		return seq("RESETBGCOLOR")
	}

	bg := fromColor(color)
	if c.bg == &bg {
		return nil
	}

	c.bg = &bg
	return seq("BGCOLOR", bg.R, bg.G, bg.B)
}

func (c *cursor) advance() {
	c.col++
	if c.col > c.cols {
		c.col = 1
		c.row++
	}
}

func (c *cursor) writeChar(s string) []byte {
	b := &bytes.Buffer{}

	// Print the character
	runes := []rune(s)
	if len(runes) == 0 || !isPrintable(runes[0]) {
		runes = []rune{' '}
	}
	b.Write([]byte(string(runes)))

	size := runewidth.StringWidth(s)
	for i := 0; i < size; i++ {
		c.advance()
	}

	return b.Bytes()
}

const (
	firstASCIIChar = '!' //041
	lastASCIIChar  = '~' //026
)

func isPrintable(ch rune) bool {
	// Is utf8
	if ch > 127 {
		return true
	}
	// Is printable ASCII
	if ch >= firstASCIIChar && ch <= lastASCIIChar {
		return true
	}
	return false
}
