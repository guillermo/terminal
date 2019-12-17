// Package terminal exposes the terminal as a matrix of characters and a collections of events.
//
//	package main
//
//	import (
//		"github.com/guillermo/terminal"
//		"github.com/guillermo/terminal/char"
//		"github.com/guillermo/terminal/events"
//	)
//
// 	func main() {
// 		// An empty Terminal is a valid one
// 		term := &terminal.Terminal{}
//
// 		err := term.Open()
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		// Always restore the terminal to the previous state
// 		defer term.Close()
//
// 		// Get the size of the terminal
// 		rows, cols := term.Size()
//
// 		// Rows and Cols start with 1
// 		term.Set(1, 1, char.C("H"))
// 		term.Set(rows, cols, char.C("i"))
// 		term.Sync()
//
// 		for {
// 			// Listen for events
// 			e := term.NextEvent()
// 			if ke, ok := e.(events.KeyboardEvent); ok {
// 				if ke.Key == "q" {
// 					// EXIT
// 					break
// 				}
// 			}
// 		}
// 	}
package terminal

import (
	"io"

	"github.com/guillermo/terminal/char"
	"github.com/guillermo/terminal/events"
	"github.com/guillermo/terminal/framebuffer"
	"github.com/guillermo/terminal/input"
	"github.com/guillermo/terminal/tty"
	"os"
)

// Terminal holds the state of the current terminal
type Terminal struct {
	Input       io.Reader
	Output      io.Writer
	DefaultChar char.Charer
	events      chan (events.Event)
	tty         *tty.TTY
	fb          *framebuffer.Framebuffer
	input       input.Input
	rows, cols  int
}

// File is the interface used for getting the Fd() of the Output device.
// For example os.Stdout.Fd() returns the file descriptor needed to ask the OS for the terminal size.
type File interface {
	Fd() uintptr
}

// Open opens a terminal.
//
// If Output is a File it will try to open the terminal as a tty.
//
// If Output is nil, os.Stdout will be used.
//
// If Input is nil, os.Stdin will be used.
//
// If DefaultChar is present, it will clear the screen with the background Color of the DefaultChar
func (t *Terminal) Open() error {
	t.fb = &framebuffer.Framebuffer{}
	t.events = make(chan events.Event, 1024)
	if t.Input == nil {
		t.Input = os.Stdin
	}
	if t.Output == nil {
		t.Output = os.Stdout
	}

	if fd, ok := t.Input.(File); ok {
		t.tty = &tty.TTY{Fd: int(fd.Fd())}
		err := t.tty.OnResize(t.onResize)
		if err != nil {
			return err
		}
		if err := t.tty.SaveTTYState(); err != nil {
			return err
		}

		if err := t.tty.RawMode(); err != nil {
			return err
		}
	} else {
		t.send("GETWINDOWSIZE")
		t.rows = 25
		t.cols = 80
	}

	// Process input
	t.input.Input = t.Input
	t.input.Open()

	go t.forwardInputEvents()

	// Prepare terminal
	t.saveScreen()
	t.hideCursor()

	if t.DefaultChar != nil {
		t.Set(1, 1, t.DefaultChar)
		t.Sync()
		t.send("ERASEALL")
	}
	return nil
}

func (t *Terminal) onResize(rows, cols int) {
	// Set terminal size
	t.fb.SetSize(rows, cols)

	// Publish event
	t.sendWinSize(t.events, rows, cols)
}

// Close resets the terminal to the previous state
func (t *Terminal) Close() error {
	if t.tty != nil {
		t.tty.Restore()
	}
	t.restoreScreen()
	t.showCursor()
	//t.fb.Close()
	return nil
}

// Size returns the terminal size
func (t *Terminal) Size() (Rows, Cols int) {
	return t.fb.Size()
}

// Sync dump all the changes in the buffer.
func (t *Terminal) Sync() {
	t.write(t.fb.Changes())
}

// NextEvent return the next events.Event. If there are no events available it will block.
func (t *Terminal) NextEvent() events.Event {
	return <-t.events
}

// Set changes the character display in the given row/col.
func (t *Terminal) Set(row, col int, ch char.Charer) {
	t.fb.Set(row, col, ch)
}

// Send send a command to the output
func (t *Terminal) send(cmd string, args ...interface{}) {
	t.write(seq(cmd, args...))
}

func (t *Terminal) write(data []byte) (n int, err error) {
	n, err = t.Output.Write(data)
	if syncer, ok := t.Output.(interface{ Sync() error }); ok {
		syncer.Sync()
	}
	return
}

func (t *Terminal) forwardInputEvents() {
	for e := range t.input.Events {
		t.events <- e
	}
	panic("input was closed")
}

func (t *Terminal) saveScreen() {
	t.send("SMCUP")
}

func (t *Terminal) restoreScreen() {
	t.send("RMCUP")
}

func (t *Terminal) hideCursor() {
	t.send("HIDECURSOR")
}
func (t *Terminal) showCursor() {
	t.send("SHOWCURSOR")
}

func (t *Terminal) sendWinSize(c chan (events.Event), rows, cols int) {
	c <- &events.WindowSizeEvent{
		Cols: cols,
		Rows: rows,
	}
}
