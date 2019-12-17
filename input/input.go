// Package input converts a io.Reader (normally a tty) into a sequence of events.
// It have basic support for xtermm and vt100 escape sequences
package input

//go:generate sequencegenerator -file sequences.go
// install with ``go install github.com/guillermo/reacty/cmd/sequencegenerator``

import (
	"bytes"
	"fmt"
	"github.com/guillermo/reacty/terminal/events"
	"io"
	"time"
)

var enableDebug bool

func debug(args ...interface{}) {
	if enableDebug {
		fmt.Println(args...)
	}
}

const inputWait = time.Millisecond * 2

// Input represent an input device. It converts a stream of bytes into a collection of events.
type Input struct {
	Input  io.Reader
	runes  <-chan (rune)
	Events <-chan (events.Event)
	data   []rune
	err    error

	currentEvent     events.Event
	currentEventSize int
}

// Open starts reading bytes from the reader and publishing events in Events.
// Once io.EOF is found, the channel will be closed.
func (i *Input) Open() {
	runesChan := make(chan (rune), 1024)
	eventsChan := make(chan (events.Event), 1024)

	i.runes = runesChan
	i.Events = eventsChan

	go i.runeLoop(runesChan)
	go i.eventLoop(eventsChan)
}

func (i *Input) runeLoop(c chan (rune)) {
	buf := make([]byte, 4096)
	for {
		n, err := i.Input.Read(buf)
		for _, ch := range bytes.Runes(buf[:n]) {
			c <- ch
		}
		if err == io.EOF {
			close(c)
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (i *Input) eventLoop(c chan (events.Event)) {

	for {
		if len(i.data) == 0 {
			i.readRune()
			if i.err != nil {
				break
			}
		}
		event, eventSize := i.genEvent()
		if i.err != nil {
			break
		}
		if event == nil {
			break
		}
		c <- event
		i.data = i.data[eventSize:]
	}
	close(c)
}

func (i *Input) genEvent() (e events.Event, size int) {

	for more := true; more; {
		pEvent, pSize, pMore := parsers.parse(i.data)
		more = pMore
		if pEvent != nil {
			e = pEvent
			size = pSize
		}
		if e == nil {
			i.readRune()
			continue
		}
		if more && i.readRuneTimeout() {
			break
		}
	}
	if e == nil {
		e = events.BytesEvent(string(i.data[0]))
		size = 1
	}
	return
}

func buildCharEvent(ch string) events.Event {
	return &events.KeyboardEvent{
		Key:  ch,
		Code: ch,
	}
}

const (
	firstASCIIChar = '!' //041
	lastASCIIChar  = '~' //026
)

func isUtf8(ch rune) bool {
	return ch > 127
}

func isASCII(ch rune) bool {
	return ch >= firstASCIIChar && ch <= lastASCIIChar
}

func (i *Input) readRune() {
	ch, ok := <-i.runes
	if !ok {
		i.err = io.EOF
	}
	i.data = append(i.data, ch)
}

// readRuneTimeout returns true in case there was a timeout
func (i *Input) readRuneTimeout() (timeout bool) {
	select {
	case ch, ok := <-i.runes:
		if !ok {
			i.err = io.EOF
		}
		i.data = append(i.data, ch)
		return false
	case <-time.After(inputWait):
		return true
	}
}
