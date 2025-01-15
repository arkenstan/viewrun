package pipeline

import (
	"viewrun/internal/ds"
)

type Directive struct {
	keyword string
	source  string
	pattern []string
	target  string
}

func ProcessDirective(command string) {

	// directive := Directive{
	// 	keyword: "",
	// 	source:  "",
	// 	pattern: []string{},
	// 	target:  "",
	// }

	// operatorStack := ds.Stack[rune]{}

	arrayStart := '['
	arrayEnd := ']'

	stringEnds := '"'

	escapeChar := '\\'

	words := []string{}

	runeBuffer := ""
	operatorStack := ds.Stack[rune]{}

	var strings []string

	for _, ch := range command {

		switch ch {

		case stringEnds:
			if operatorStack.Top() == escapeChar {
				runeBuffer += string(ch)

			} else if operatorStack.Top() == stringEnds {
				// pop operator
				operatorStack.Pop()
				if operatorStack.Top() == arrayStart {
				}
			} else if operatorStack.Top() == '\\' {
				// Escape string
			} else {
				operatorStack.Push('"')
			}

		case ' ', '\t':
			words = append(words, runeBuffer)
			runeBuffer = ""

		default:
			runeBuffer += string(ch)
			if operatorStack.Top() == escapeChar {
				operatorStack.Pop()
			}

		}
	}

}
