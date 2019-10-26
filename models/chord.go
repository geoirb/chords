package models

import (
	"fmt"
	"strings"
)

type Chord struct {
	Symbol string
	Line   int
	Index  int
}

func getChords(i int, line string) (lineChords []Chord) {
	fmt.Println(line)
	resultLine := line

	for begin, s := range line {
		if s == '[' {
			end := strings.Index(line[begin:], "]")
			chord := Chord{
				Symbol: line[begin+1 : begin+end],
				Line:   i,
				Index:  begin + 1,
			}

			fmt.Println(chord)

			// lineChords = append(lineChords, chord)
			resultLine = strings.Replace(resultLine, line[begin:begin+end+1], "", 1)
		}
	}

	fmt.Println(resultLine)

	return
}

func CreateText(text string) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		getChords(i, line)
	}
}
