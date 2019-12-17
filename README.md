# terminal

[![Build status](https://github.com/guillermo/terminal/workflows/Go/badge.svg)](https://github.com/guillermo/terminal/actions?query=workflow%3AGo)
[![Godoc](https://godoc.org/github.com/guillermo/terminal?status.svg)](https://godoc.org/github.com/guillermo/terminal)
[![GODEV](https://img.shields.io/badge/godev-overview-5272B4.svg)](https://pkg.go.dev/github.com/guillermo/terminal?tab=overview)

Terminal exposes the terminal as a matrix of characters and a stream of events.

## Structure 

```asciiart
+------------------------------------------------------+
|                      terminal                        |
+-----+--------------+--------------------+------------+
      |              |                    |
      v              v                    v
+-----+-----+  +-----+-----+    +---------+-------------+
|    tty    |  |   input   |    |      framebuffer      |
+-----------+  +-----+-----+    +-----+-----------+-----+
                     |                |           |
                     v                |           v
               +-----+-----+          |     +-----+-----+
               |  events   |          |     |   ansi    |
               +-----------+          |     +-----+-----+
                                      |           |
                                      v           v
                                +-----+-----------+-----+
                                |         area          |
                                +-----------------------+

```

*terminal*: Package terminal exposes the terminal as a matrix of characters and a collections of events.

It have two main components the input and the output:

## Input

*events* Package events contains the basic primitives of the input events

*input* Package input converts a io.Reader (normally a tty) into a sequence of events.

*tty* Package tty implements the os part of the terminal.

## Output

*char* Package char represent a Char in a terminal.

*area* Package area implements a matrix of Characters.

*ansi* Package ansi converts a given area into a stream of ansi sequences to be dump into a terminal.

*framebuffer* Package framebuffer stores the current terminal state and return the ansi sequences require to transform the current state to the new one.






## Example

```go
package main

import (
    "github.com/guillermo/terminal"
    "github.com/guillermo/terminal/char"
    "github.com/guillermo/terminal/events"
)


func main() {
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
    // Without calling Sync no changes will be dump to the Output
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
```


