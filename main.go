package main

import (
	"flag"
	"os"
	"viewrun/cmd"
)

var (
  pipeline string
  filePath string
)

func main(){

  pwd, err := os.Getwd()

  if err != nil{
    panic( err)
  }

  defaultPipeline := pwd + "/.viewrun"

  flag.StringVar(&pipeline, "pipeline", "", "Pipeline name to run")
  flag.StringVar(&filePath, "file", defaultPipeline, "Pipeline config file")
  flag.Parse()

  

  app.App(pipeline, filePath)
}
