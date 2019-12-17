// Package events contains the basic primitives of the input and output events
package events

import (
	"fmt"
)

// BytesEvent holds all the unrecognized bytes from the input
type BytesEvent string

func (b BytesEvent) String() string {
	return fmt.Sprintf("BytesEvent: %s", []byte(b))
}

// Event holds the information of the event.
type Event interface {
	String() string
}

//https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/code/code_values

// KeyboardEvent represent a key press.
type KeyboardEvent struct {
	// Key Is the printibal character
	Key string
	// Ctrl is true if the Control key is pressed
	Ctrl bool
	// Alt is true if the Alt key is pressed
	Alt bool
	// Shift is true if the Shift key is pressed
	Shift bool
	// Code holds the name of the key. For example: Right, Escape or Enter
	Code string
}

func (ke KeyboardEvent) String() string {
	s := ke.Code
	if len(s) == 0 {
		s = ke.Key
	}

	if ke.Alt {
		s = "Alt+" + s
	}
	if ke.Ctrl {
		s = "Ctrl+" + s
	}

	return "KeyboardEvent: " + s
}

// WindowSizeEvent represents the window terminal size in characters
type WindowSizeEvent struct {
	Cols int
	Rows int
}

func (wse *WindowSizeEvent) String() string {
	return fmt.Sprintf("WindowSizeEvent: %dx%d", wse.Rows, wse.Cols)
}

// ErrorEvent holds an error produce while processing the input
type ErrorEvent string

func (e ErrorEvent) String() string {
	return "ErrorEvent: " + string(e)
}

// MouseEvent represent any action related with the mouse
type MouseEvent struct {
	Row, Col int
	Button   int
	Action   string
}

func (m *MouseEvent) String() string {
	return fmt.Sprintf("MouseEvent: (%d,%d) Button: %d Action: %s ", m.Row, m.Col, m.Button, m.Action)
}

// PasteEvent represent the intent of pasting content in the terminal
type PasteEvent string

func (p PasteEvent) String() string {

	return fmt.Sprintf("PasteEvent: %q", string(p))
}
