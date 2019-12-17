package ansi

import "strings"

type style uint

const (
	bold style = 1 << iota
	faint
	italic
	underline
	blink
	inverse
	invisible
	crossed
	double
)

func newStyle(ch Char) style {
	s := style(0)
	if ch.Bold() {
		s.Set(bold)
	}

	if ch.Faint() {
		s.Set(faint)
	}
	if ch.Italic() {
		s.Set(italic)
	}
	if ch.Underline() {
		s.Set(underline)
	}
	if ch.Blink() {
		s.Set(blink)
	}
	if ch.Inverse() {
		s.Set(inverse)
	}
	if ch.Invisible() {
		s.Set(invisible)
	}
	if ch.Crossed() {
		s.Set(crossed)
	}
	if ch.Double() {
		s.Set(double)
	}
	return s
}

func (s *style) codes() string {
	codes := []string{}
	if *s&bold != 0 {
		codes = append(codes, "1")
	}
	if *s&faint != 0 {
		codes = append(codes, "2")
	}
	if *s&italic != 0 {
		codes = append(codes, "3")
	}
	if *s&underline != 0 {
		codes = append(codes, "4")
	}
	if *s&blink != 0 {
		codes = append(codes, "5")
	}
	if *s&inverse != 0 {
		codes = append(codes, "7")
	}
	if *s&invisible != 0 {
		codes = append(codes, "8")
	}
	if *s&crossed != 0 {
		codes = append(codes, "9")
	}
	if *s&double != 0 {
		codes = append(codes, "21")
	}
	return strings.Join(codes, ";")
}

func (s *style) Set(newStyle style) {
	*s = *s | newStyle
}

func (s *style) Any() bool {
	return *s != 0
}

func (s *style) Normal() bool {
	return *s == 0
}
