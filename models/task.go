package main

type task struct{
  TaskType string `json:"task"`
  Sources  []string `json:"sources"`
  Target string `json:"target"`
  Options map[string]string `json:"options"`
}
