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

func getChords(i int, line string) {
	fmt.Println(line)
	resultLine := line
	for begin, s := range line {
		if s == '[' {
			end := strings.Index(line[begin+1:], "]")
			fmt.Printf("%c %c\n", line[begin], line[end])
			resultLine = resultLine[:begin] + resultLine[end+1:]
		}
	}

	// fmt.Println(resultLine)
}

func CreateText(text string) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		getChords(i, line)
	}
}
