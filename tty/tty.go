// Package tty implements the os part of the terminal
package tty

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tredoe/term/sys"
)

// TTY is a basic structure to control the terminal state
type TTY struct {
	Fd                  int
	oldState, lastState sys.Termios
}

// SaveTTYState records the current low level terminal state
func (t *TTY) SaveTTYState() error {
	if err := sys.Getattr(t.Fd, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Getattr", err)
	}
	t.oldState = t.lastState
	return nil
}

// Restore returns the tty to the state save with SaveTTYState
func (t *TTY) Restore() error {
	if err := sys.Setattr(t.Fd, sys.TCSANOW, &t.oldState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	t.lastState = t.oldState
	return nil
}

// TermSize  returns the terminal size
func (t *TTY) TermSize() (rows, cols int, err error) {
	ws := sys.Winsize{}
	if err := sys.GetWinsize(t.Fd, &ws); err != nil {
		return 0, 0, err
	}

	return int(ws.Row), int(ws.Col), nil
}

// OnResize listen of the SIGWINCH event (trigger by terminal apps after the terminal changes its sizes).
// The callback cbk is also called the first time this function is called even if not SIGWINCH is called.
func (t *TTY) OnResize(cbk func(rows, cols int)) error {
	sync := func() error {
		rows, cols, err := t.TermSize()
		if err != nil {
			return err
		}
		cbk(rows, cols)
		return nil
	}

	sigChan := make(chan (os.Signal))
	signal.Notify(sigChan, syscall.SIGWINCH)
	go func() {
		for range sigChan {
			sync()
		}
	}()
	return sync()
}

// RawMode sets the terminal into RawMode
func (t *TTY) RawMode() error {
	// Input modes - no break, no CR to NL, no NL to CR, no carriage return,
	// no strip char, no start/stop output control, no parity check.
	t.lastState.Iflag &^= (sys.BRKINT | sys.IGNBRK | sys.ICRNL | sys.INLCR |
		sys.IGNCR | sys.ISTRIP | sys.IXON | sys.PARMRK)

	// Output modes - disable post processing.
	t.lastState.Oflag &^= sys.OPOST

	// Local modes - echoing off, canonical off, no extended functions,
	// no signal chars (^Z,^C).
	t.lastState.Lflag &^= (sys.ECHO | sys.ECHONL | sys.ICANON | sys.IEXTEN | sys.ISIG)

	// Control modes - set 8 bit chars.
	t.lastState.Cflag &^= (sys.CSIZE | sys.PARENB)
	t.lastState.Cflag |= sys.CS8

	// Control chars - set return condition: min number of bytes and timer.
	// We want read to return every single byte, without timeout.
	t.lastState.Cc[sys.VMIN] = 1 // Read returns when one char is available.
	t.lastState.Cc[sys.VTIME] = 0

	// Put the terminal in raw mode after flushing
	if err := sys.Setattr(t.Fd, sys.TCSAFLUSH, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	return nil
}
