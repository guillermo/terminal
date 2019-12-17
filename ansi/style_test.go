package ansi

import "testing"

func TestStyle(t *testing.T) {
	s := newStyle(char{})
	if s != 0 {
		t.Error("Expectin an empty style")
	}
	if s.Any() {
		t.Error("Expecting no style")
	}

	s.Set(crossed)
	if s != crossed {
		t.Error("should be crossed")
	}

	s = newStyle(char{
		bold:      true,
		faint:     true,
		italic:    true,
		underline: true,
		blink:     true,
		inverse:   true,
		crossed:   true,
		double:    true,
	})

	if s != bold|faint|italic|underline|blink|inverse|crossed|double {
		t.Fatal(s)
	}
	if !s.Any() {
		t.Error("Expecting  style")
	}
	s = 0
	s.Set(bold)
	s.Set(faint)
	s.Set(italic)
	s.Set(underline)
	s.Set(blink)
	s.Set(inverse)
	s.Set(invisible)
	s.Set(crossed)
	s.Set(double)

	if s != bold|faint|italic|underline|blink|inverse|invisible|crossed|double {
		t.Fatal(s, bold|faint|italic|underline|blink|inverse|invisible|crossed|double)

	}
}
