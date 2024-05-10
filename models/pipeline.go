package main

type PipelineContext struct{
  Source []string
  config map[string]string
}

type Task struct{
  TaskType string `json:"task"`
  Sources  []string `json:"sources"`
  Target string `json:"target"`
  Options map[string]string `json:"options"`
}

type PipelineConfig struct{
  Pipelines map[string]Task `json:"pipelines"`
}

