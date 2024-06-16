package app

import (
	"fmt"
	"viewrun/internal/pipeline"
)

func App(pipelineName string, filePath string) {

	fmt.Println("Running pipeline:", pipelineName, "\n\nFile Path:", filePath)

	pipeline := main.PipelineConfig{}
}
