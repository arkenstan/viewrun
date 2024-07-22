package app

import (
	"fmt"
	"viewrun/internal/pipeline"
)

func App(pipelineName string, filePath string) {
	fmt.Println("Running pipeline:", pipelineName, "\n\nFile Path:", filePath)
	pipelineConfig, err := pipeline.LoadPipeline(filePath)

	if err != nil {
		fmt.Println("Unable to parse pipeline")
		return
	}

	fmt.Println("PIPELINE CONFIG: ", pipelineConfig)

}
