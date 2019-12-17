package main

import (
	"fmt"
	"github.com/guillermo/terminal"
	"github.com/guillermo/terminal/char"
	"github.com/guillermo/terminal/events"
	"image/color"
	"os"
	"time"
)

func main() {

	term := &terminal.Terminal{
		Input:       os.Stdin,
		Output:      os.Stdout,
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
