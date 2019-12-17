package main

import (
	"fmt"
	"github.com/guillermo/reacty/terminal"
	"github.com/guillermo/reacty/terminal/eachchange"
	"github.com/guillermo/reacty/terminal/events"
	"image/color"
	"os"
	"time"
)

func main() {

	term := &terminal.Terminal{
		Input:       os.Stdin,
		Output:      os.Stdout,
		DefaultChar: c(" "),
	}

	err := term.Open()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	var row, col int
	printf := func(format string, args ...interface{}) {
		s := fmt.Sprintf(format, args...)
		rows, cols := term.Size()
		for _, ch := range s {
			term.Set(row+1, col+1, c(string(ch)))
			col++
			if col >= cols {
				col = 0
				row++
			}
			if row >= rows {
				row = 0
			}

		}
		row++
		col = 0
		if row >= rows {
			row = 0
		}
		term.Sync()
	}

	for {
		e := term.NextEvent()
		printf("%v", e)
		if ke, ok := e.(events.KeyboardEvent); ok {
			if ke.Key == "q" {
				time.Sleep(time.Second / 2)
				break
			}
			if ke.Code == "C" && ke.Ctrl {
				time.Sleep(time.Second / 2)
				break
			}
		}

	}

}

type c string

func (ch c) Content() string {
	return string(ch)
}

func (ch c) Background() color.Color {
	return color.White
}

func (ch c) Foreground() color.Color {
	return color.Black
}

func (ch c) Bold() bool {
	return false
}

func (ch c) Faint() bool {
	return false
}

func (ch c) Italic() bool {
	return false
}

func (ch c) Underline() bool {
	return false
}

func (ch c) Blink() bool {
	return false
}

func (ch c) Inverse() bool {
	return false
}

func (ch c) Invisible() bool {
	return false
}

func (ch c) Crossed() bool {
	return false
}

func (ch c) Double() bool {
	return false
}
func (ch c) Equal(ch2 eachchange.Char) bool {
	if string(ch) != string(ch2.(c)) {
		return false
	}
	return true
}
