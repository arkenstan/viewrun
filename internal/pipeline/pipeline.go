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

	for scanner.Scan() {
		fmt.Println("SCANNER | ", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return nil, nil

}
