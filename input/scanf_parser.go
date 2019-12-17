package input

import (
	"fmt"
	"github.com/guillermo/reacty/terminal/events"
)

//https://play.golang.org/p/dudzyAar-Wa
func scanParser(data []rune) (e events.Event, n int, more bool) {

	var a, b, c int
	n, err := fmt.Sscanf(string(data), "\x1b[%d;%d;%dt", &a, &b, &c)
	if err == nil {
		return &events.WindowSizeEvent{Rows: b, Cols: c}, len(data), false
	}
	errMsg := err.Error()
	if errMsg == "input does not match format" ||
		errMsg == "expected integer" {
		return nil, 0, false
	}
	if errMsg == "unexpected EOF" ||
		errMsg == "EOF" {
		return nil, 0, true
	}
	panic(err)

}

func scanMouseDownParser(data []rune) (e events.Event, n int, more bool) {

	var a, col, row int
	n, err := fmt.Sscanf(string(data), "\x1b[<%d;%d;%dm", &a, &col, &row)
	if err == nil {
		return &events.MouseEvent{
			Col:    col,
			Row:    row,
			Button: a,
			Action: "Press",
		}, len(data), false
	}
	errMsg := err.Error()
	if errMsg == "input does not match format" ||
		errMsg == "expected integer" {
		return nil, 0, false
	}
	if errMsg == "unexpected EOF" ||
		errMsg == "EOF" {
		return nil, 0, true
	}
	panic(err)

}

func scanMouseReleaseParser(data []rune) (e events.Event, n int, more bool) {

	var a, col, row int
	n, err := fmt.Sscanf(string(data), "\x1b[<%d;%d;%dM", &a, &col, &row)
	if err == nil {
		me := &events.MouseEvent{Col: col, Row: row}
		switch a {
		case 64:
			me.Action = "SCROLLUP"
			me.Button = 5
		case 65:
			me.Action = "SCROLLDOWN"
			me.Button = 4
		case 0, 1, 2:
			me.Action = "Release"
			me.Button = a
		}
		return me, len(data), false
	}
	errMsg := err.Error()
	if errMsg == "input does not match format" ||
		errMsg == "expected integer" {
		return nil, 0, false
	}
	if errMsg == "unexpected EOF" ||
		errMsg == "EOF" {
		return nil, 0, true
	}
	panic(err)

}

func scanPasteParser(data []rune) (e events.Event, n int, more bool) {

	var content string
	n, err := fmt.Sscanf(string(data), "\x1b[200~%s\x1b[201~", &content)
	if err == nil {
		return events.PasteEvent(content), len(data), false
	}
	errMsg := err.Error()
	if errMsg == "input does not match format" ||
		errMsg == "expected integer" {
		return nil, 0, false
	}
	if errMsg == "unexpected EOF" ||
		errMsg == "EOF" {
		return nil, 0, true
	}
	panic(err)

}
