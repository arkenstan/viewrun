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

	directive := Directive{}

	operatorStack := ds.Stack[rune]{}

	for _, ch := range command {

		topOperator := operatorStack.Top()

		switch ch {

		case '"':
			if topOperator == '"' {
				// String end
			} else {
				operatorStack = append(operatorStack, '"')
			}

		case '[':
			if topOperator != '"' {
				operatorStack = append(operatorStack)
			} else {
			}

		case ' ', '\t':
			if topOperator == '"' {
				// In string
			} else {

			}

		case '\\':
			// Escape next character

		case '#':
			if topOperator != '"' {
				break
			}

		default:

		}

	}

}
