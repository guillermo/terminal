package main

import (
	"fmt"
	"github.com/guillermo/terminal"
	"github.com/guillermo/terminal/char"
	"github.com/guillermo/terminal/events"
	"image/color"
	"time"
)

func main2() {
	// An empty Terminal is a valid one
	term := &terminal.Terminal{}

	err := term.Open()
	if err != nil {
		panic(err)
	}

	// Always restore the terminal to the previous state
	defer term.Close()

	// Get the size of the terminal
	rows, cols := term.Size()

	// Rows and Cols start with 1
	term.Set(1, 1, char.C("H"))
	term.Set(rows, cols, char.C("i"))
	term.Sync()

	for {
		// Listen for events
		e := term.NextEvent()
		if ke, ok := e.(events.KeyboardEvent); ok {
			if ke.Key == "q" {
				// EXIT
				break
			}
		}
	}
}

func main() {

	term := &terminal.Terminal{
		DefaultChar: &char.Char{BackgroundColor: color.White, ForegroundColor: color.Black},
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

func c(a string) *char.Char {
	return &char.Char{Value: a}
}
