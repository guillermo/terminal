package input

import (
	"github.com/guillermo/reacty/terminal/events"
	"strings"
)

type parser func(data []rune) (e events.Event, eventSize int, more bool)

func sequenceParser(data []rune) (e events.Event, eventSize int, more bool) {

	var lastSeq *seq

	for _, current := range sequences {
		if current.seq == string(data) {
			a := current
			lastSeq = &a
			continue
		}
		if !more && strings.HasPrefix(current.seq, string(data)) {
			more = true
		}
	}
	if lastSeq == nil {
		return
	}
	return lastSeq.event, len(data), more

}

func utf8Parser(data []rune) (e events.Event, eventSize int, more bool) {
	if len(data) == 1 && isUtf8(data[0]) {
		e = &events.KeyboardEvent{Key: string(data[0]), Code: string(data[0])}
		eventSize = 1
	}
	return
}

type multiParser []parser

func (p multiParser) parse(data []rune) (e events.Event, n int, more bool) {
	for _, parser := range p {
		pe, pn, pmore := parser(data)
		if e == nil {
			// Assign the first occurrance
			e = pe
			n = pn
		}
		if pmore {
			more = true
		}
	}

	return
}

var parsers multiParser = []parser{
	utf8Parser,
	sequenceParser,
	scanParser,
	scanMouseDownParser,
	scanMouseReleaseParser,
	scanPasteParser,
}
