package input

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func init() {
	enableDebug = false
}

//https://jsfiddle.net/qfL9z4rv/4/
var inputTests = []struct {
	in  []byte
	out string
}{
	{[]byte("❤"), "❤"},
	{[]byte("a"), "a"},
	{[]byte(" "), "Space"},
	{[]byte{127}, "Backspace"},
	{[]byte{0x1}, "Ctrl+A"},
	{[]byte{0x3}, "Ctrl+C"},
	{[]byte{0x1A}, "Ctrl+Z"},

	{[]byte{0x1c}, "Ctrl+\\"},
	{[]byte{0x1d}, "Ctrl+]"},

	{[]byte{0x1B}, "Escape"},
	{[]byte("\x1bA"), "Alt+A"},
	{[]byte("\x1b "), "Alt+Space"},

	{[]byte{0x1A}, "Ctrl+Z"},
	//	{[]byte{0x27}, "Ctrl+C"}, //Wrong?
	{[]byte{32}, "Space"},

	{[]byte("\033["), "Alt+["},
	{[]byte("\033[A"), "Up"},
	{[]byte("\033[C"), "Right"},
	{[]byte("\033[B"), "Down"},
	{[]byte("\033[D"), "Left"},
	{[]byte("\033a"), "Alt+a"},
	{[]byte("\033A"), "Alt+A"},

	{[]byte("\033[1;5A"), "Ctrl+Up"},
	{[]byte("\033[1;5B"), "Ctrl+Down"},
	{[]byte("\033[1;5C"), "Ctrl+Right"},
	{[]byte("\033[1;5D"), "Ctrl+Left"},
	{[]byte{32}, "Space"},
	{[]byte("\033OP"), "F1"},
	{[]byte("\033OQ"), "F2"},
	{[]byte("\033OR"), "F3"},
	{[]byte("\033OS"), "F4"},

	{[]byte("\033[15~"), "F5"},
	//	{[]byte("\033[16~"), "F5"},
	{[]byte("\033[17~"), "F6"},
	{[]byte("\033[18~"), "F7"},
	{[]byte("\033[19~"), "F8"},
	{[]byte("\033[20~"), "F9"},
	{[]byte("\033[21~"), "F10"},
	{[]byte("\033[23~"), "F11"},
	{[]byte("\033[24~"), "F12"},
	{[]byte("\033[5~"), "PageUp"},
	{[]byte("\033[6~"), "PageDown"},

	{[]byte("\033[H"), "Home"},
	{[]byte("\033[F"), "End"},

	{[]byte("\033[2~"), "Insert"},
	{[]byte("\033[3~"), "Delete"},
	{[]byte{0177}, "Backspace"},
	{[]byte("\x1b[8;24;80t"), "WindowSizeEvent: 24x80"},
	/*

		/*
			//
				// Xterm
				{[]byte("\x1b[32;1;1M"), "MouseDown LEFT 1x1"},
				{[]byte("\x1b[35;1;1M"), "MouseUp 1x1"},

				{[]byte("\x1b[33;1;1M"), "MouseDown MIDDLE 1x1"},
				{[]byte("\x1b[35;1;1M"), "MouseUp 1x1"},

				{[]byte("\x1b[34;1;1M"), "MouseDown RIGHT 1x1"},
				{[]byte("\x1b[35;1;1M"), "MouseUp 1x1"},

				{[]byte("\x1b[97;1;1M"), "ScrollDown 1x1"},
				{[]byte("\x1b[96;1;1M"), "ScrollUp 1x1"},

				// Other
				{[]byte("\x1b[<0;1;1M"), "MouseDown LEFT 1x1"},
				{[]byte("\x1b[<0;1;1m"), "MouseUp LEFT 1x1"},
				{[]byte("\x1b[<1;1;1M"), "MouseDown MIDDLE 1x1"},
				{[]byte("\x1b[<1;1;1m"), "MouseUp MIDDLE 1x1"},
				{[]byte("\x1b[<2;1;1M"), "MouseDown RIGHT 1x1"},
				{[]byte("\x1b[<2;1;1m"), "MouseUp RIGHT 1x1"},

				{[]byte("\x1b[<65;1;1M"), "ScrollDown at 1x1"},
				{[]byte("\x1b[<64;1;1M"), "ScrollUp at 1x1"},

				{[]byte("\033[4;660;1001t"), "WindowResizeEvent: 10x10"},
				{[]byte("\033[<0;1;1M"), "MouseUp: 1x1"},
				{[]byte("\033[<0;1;1m"), "MouseDown: 1x1"},
				{[]byte("\033[200~hello\033[201~"), "PasteEvent: \"hello\""},


	*/
}

type FakeIO chan ([]byte)

func (buf FakeIO) Read(d []byte) (int, error) {
	data, ok := <-buf
	if !ok {
		return 0, io.EOF
	}
	copy(d, data)
	return len(data), nil
}

func TestUnknown(t *testing.T) {
	buf := make(chan ([]byte))
	input := &Input{Input: FakeIO(buf)}
	input.Open()

	buf <- []byte("\033[❓")
	close(buf)

	event, ok := <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed")
	}
	eName := event.String()
	if eName != "KeyboardEvent: Alt+[" {
		t.Fatal("Expected a keyboard event. Got:", eName)
	}

	event, ok = <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed earlier than expected")
	}
	eName = event.String()
	if eName != "KeyboardEvent: ❓" {
		t.Fatal("Expected a keyboard event. Got:", eName)
	}

	event, ok = <-input.Events
	if ok {
		t.Fatal("I should get io.EOF")
	}

}

func TestReiszeEvent(t *testing.T) {
	buf := make(chan ([]byte))
	input := &Input{Input: FakeIO(buf)}
	input.Open()

	s := "\x1b[8;24;80t"
	buf <- []byte(s)
	close(buf)

	event, ok := <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed")
	}
	eName := event.String()
	exp := "WindowSizeEvent: 24x80"
	if eName != exp {
		t.Fatal("Expected a ", exp, " Got:", eName)
	}

	event, ok = <-input.Events
	if ok {
		t.Fatal("I should get io.EOF")
	}

}

func TestSingle(t *testing.T) {
	buf := make(chan ([]byte))
	input := &Input{Input: FakeIO(buf)}
	input.Open()

	buf <- []byte{'a'}

	event, ok := <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed")
	}
	expectation := "KeyboardEvent: a"
	if event.String() != expectation {
		t.Errorf("expected %q, got %q.", expectation, event.String())
	}

	buf <- []byte("❤")
	event, ok = <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed")
	}
	expectation = "KeyboardEvent: ❤"
	if event.String() != expectation {
		t.Errorf("expected %q, got %q.", expectation, event.String())
	}

	buf <- []byte(" ")
	event, ok = <-input.Events
	if !ok {
		t.Fatal("Test():The channel was closed")
	}
	expectation = "KeyboardEvent: Space"
	if event.String() != expectation {
		t.Errorf("expected %q, got %q.", expectation, event.String())
	}

	close(buf)

}

func TestIndividualInput(t *testing.T) {
	for _, test := range inputTests {
		t.Run(fmt.Sprintf("%X - (%s)", test.in, string(test.in)), func(t *testing.T) {

			buf := make(chan ([]byte))
			input := &Input{Input: FakeIO(buf)}
			input.Open()

			buf <- test.in

			event, ok := <-input.Events
			if !ok {
				t.Fatal("The channel was closed")
			}
			s := event.String()
			if strings.HasPrefix(s, kbEventString) {
				s = s[len(kbEventString):]
			}
			if s != test.out {
				t.Errorf("expected %q, got %q.", test.out, s)
			}
			close(buf)

		})
	}
}

const kbEventString = "KeyboardEvent: "

func TestStreamInput(t *testing.T) {

	buf := make(chan ([]byte))
	input := &Input{Input: FakeIO(buf)}
	input.Open()
	for _, test := range inputTests {
		t.Run(fmt.Sprintf("%X - (%s)", test.in, string(test.in)), func(t *testing.T) {
			buf <- test.in
			event, ok := <-input.Events
			if !ok {
				t.Fatal("The channel was closed")
			}
			s := event.String()
			if strings.HasPrefix(s, kbEventString) {
				s = s[len(kbEventString):]
			}
			if s != test.out {
				t.Errorf("expected %q, got %q.", test.out, s)
			}

		})
	}
	close(buf)
}
