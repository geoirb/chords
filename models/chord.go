package models

import (
	"strings"
)

type Chord struct {
	Symbol string `json:"chord"`
	Index  int    `json:"position"`
}

type Line struct {
	CleanLine string  `json:"line"`
	Сhords    []Chord `json:"chords"`
	Index     int     `json:"idx"`
}

func getChords(line string) (resultLine string, lineChords []Chord) {
	resultLine = line

	for begin, s := range line {
		if s == '[' {
			end := strings.Index(line[begin:], "]")
			chord := Chord{
				Symbol: line[begin+1 : begin+end],
				Index:  begin + 1,
			}

			lineChords = append(lineChords, chord)
			resultLine = strings.Replace(resultLine, line[begin:begin+end+1], "", 1)
		}
	}

	return
}

func CreateJSON(text string) (textJSON []Line) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lineJSON := Line{
			Index: i,
		}
		line = strings.TrimSpace(line)
		lineJSON.CleanLine, lineJSON.Сhords = getChords(line)
		textJSON = append(textJSON, lineJSON)
	}

	return
}
