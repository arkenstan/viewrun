package pipeline

import (
	"encoding/json"
	"io"
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

func loadPipeline(filePath string) (*PipelineConfig, error) {
	fileData, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	fileBytes, _ := io.ReadAll(fileData)

	var pipelineConfig PipelineConfig

	_ = json.Unmarshal(fileBytes, &pipelineConfig)

	return &pipelineConfig, nil
}
