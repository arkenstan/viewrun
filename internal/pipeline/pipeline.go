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

	for scanner.Scan() {

		switch scanner.Text() {

		case "\n":
			directive = append(directive, currentBuffer)
			currentBuffer = ""
			if !multiLine && len(directive) > 1 {
				directives = append(directives, directive)
				directive = nil
			}

		case " ", "\t":
			if currentBuffer != "" {
				fmt.Println("WORD | ", currentBuffer)
				directive = append(directive, currentBuffer)
				currentBuffer = ""
			}

		case "\\":
			multiLine = true

		default:
			currentBuffer += scanner.Text()
			multiLine = false

		}

	}

	for _, buff := range directives {
		fmt.Println("COMMAND", buff, len(buff))
	}

	// actualDirectives := []string{};

	// for

	// fmt.Println(knownBuffer)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return nil, nil
}
