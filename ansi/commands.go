package ansi

import (
	"fmt"
)

var commands = map[string]string{
	// SMCUP will enable the alternative buffer
	"SMCUP":         "\x1b[?1049h\x1b[22;0;0t", // Activate alternate buffer
	"RMCUP":         "\x1b[?1049l\x1b[23;0;0t", // Undo ^^
	"HIDECURSOR":    "\x1b[?25l",
	"SHOWCURSOR":    "\x1b[?25h",
	"SOFTRESET":     "\x1b[>!p",
	"ENABLEMOUSE":   "\x1b[?1000;1006;1015h",
	"DISABLEMOUSE":  "\x1b[?1000;1006;1015l",
	"CLEAR":         "\x1b\x0c",
	"CURSORUP":      "\x1bA",
	"CURSORDOWN":    "\x1bB",
	"CURSORRIGHT":   "\x1bC",
	"CURSORLEFT":    "\x1bD",
	"CURSORHOME":    "\x1bH",
	"ERASEBELOW":    "\x1b[0J",
	"ERASEABOVE":    "\x1b[1J",
	"ERASEALL":      "\x1b[2J",
	"GETWINDOWSIZE": "\x1b[18t",
	"CURSORUPN":     "\x1b[%dA",
	"GOTO":          "\x1b[%d;%dH", // RowxColumn starting on 1
	"ENABLEPASTE":   "\x1b[?2004h", // Enable paste
	"DISABLEPASTE":  "\x1b[?2004l", // Disable paste
	"CHARSTYLE":     "\x1b[%sm",
	"BGCOLOR":       "\x1b[48;2;%d;%d;%dm",
	"RESETBGCOLOR":  "\x1b[49m",
	"FGCOLOR":       "\x1b[38;2;%d;%d;%dm",
	"RESETFGCOLOR":  "\x1b[39m",
}

// Sequence generate the given command with the provided arguments
func seq(name string, args ...interface{}) []byte {
	c := commands[name]
	if len(args) == 0 {
		return []byte(c)
	}
	return []byte(fmt.Sprintf(string(c), args...))
}
