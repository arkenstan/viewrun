package pipeline

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	TaskType string            `json:"task"`
	Sources  []string          `json:"sources"`
	Target   string            `json:"target"`
	Options  map[string]string `json:"options"`
}

type PipelineConfig struct {
	Pipelines map[string][]Task `json:"pipelines"`
}

func (p PipelineConfig) GetPipeline(name string) []Task {
	return p.Pipelines[name]
}

func LoadPipeline(filePath string) (*PipelineConfig, error) {
	// fileData, err := os.Open(filePath)

	// if err != nil {
	// 	return nil, err
	// }

	// fileBytes, _ := io.ReadAll(fileData)
	// var pipelineConfig PipelineConfig
	// _ = json.Unmarshal(fileBytes, &pipelineConfig)

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanRunes)

	// currentBuffer := ""

	// directives := [][]string{}

	// directive := []string{}

	// multiLine := false

	// inString := false

	operatorStack := []string{}

	commands := []string{}

	commandString := ""

	for scanner.Scan() {

		fmt.Println("opStack: ", operatorStack, scanner.Text())

		switch scanner.Text() {

		case "\n":
			var topOperator string
			if len(operatorStack) > 0 {
				topOperator = operatorStack[len(operatorStack)-1]
			}

			fmt.Println("TOP OP:", topOperator)

			if topOperator == "\\" {
				operatorStack = operatorStack[:len(operatorStack)-1]
			} else if topOperator == "\"" {
				// Processing String
			} else {
				fmt.Println("COMMAND STRING: ", commandString)
				commands = append(commands, commandString)
				commandString = ""
			}

		case "\\":
			operatorStack = append(operatorStack, "\\")

		case "\"":

			var topOperator string
			if len(operatorStack) > 0 {
				topOperator = operatorStack[len(operatorStack)-1]
			}

			fmt.Println("TOP", topOperator)

			if topOperator == "" {
				operatorStack = append(operatorStack, "\"")
			} else if topOperator == "\\" {
				fmt.Println("HERE 2")
				operatorStack = operatorStack[:len(operatorStack)-1]
			} else if topOperator == "\"" {
				fmt.Println("HERE 1")
				operatorStack = operatorStack[:len(operatorStack)-1]
			}

			commandString += scanner.Text()

		// case "\n":
		// 	directive = append(directive, currentBuffer)
		// 	currentBuffer = ""
		// 	if !multiLine && len(directive) > 1 {
		// 		directives = append(directives, directive)
		// 		directive = nil
		// 	}

		// case " ", "\t":
		// 	if inString {
		// 		currentBuffer += scanner.Text()
		// 	} else if currentBuffer != "" && !inString {
		// 		fmt.Println("WORD | ", currentBuffer)
		// 		directive = append(directive, currentBuffer)
		// 		currentBuffer = ""
		// 	}

		// case "\"":
		// 	currentBuffer += "\""
		// 	if inString {
		// 		inString = false
		// 		if currentBuffer != "" {
		// 			directive = append(directive, currentBuffer)
		// 			currentBuffer = ""
		// 		}
		// 	} else {
		// 		inString = true
		// 	}

		// case "\\":
		// 	multiLine = true

		default:
			fmt.Println("Command String:", commandString)
			commandString += scanner.Text()
			// currentBuffer += scanner.Text()
			// multiLine = false
		}

	}

	fmt.Println(commands)

	for _, buff := range commands {
		// ProcessDirective(buff)
		fmt.Println("COMMAND", buff, len(buff))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return nil, nil
}
