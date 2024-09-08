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

	currentBuffer := ""

	directives := [][]string{}

	directive := []string{}

	multiLine := false

	inString := false

	operatorStack := []string{}
	commandString := ""

	for scanner.Scan() {

		switch scanner.Text() {

		case "\n":

		case "\\":
			operatorStack = append(operatorStack, "\\")

		case "\"":
			if operatorStack[len(operatorStack)-1] == "\\" {
				operatorStack = operatorStack[:len(operatorStack)-1]
			} else {
				operatorStack = append(operatorStack, "\"")
			}

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
			currentBuffer += scanner.Text()
			multiLine = false
		}

	}

	for _, buff := range directives {
		// ProcessDirective(buff)
		fmt.Println("COMMAND", buff, len(buff))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return nil, nil
}
