package main

import (
	"fmt"
)

type plugin struct{}


func (plg plugin) Exec(context PipelineContext) PipelineContext {


fmt.Println(context)

return context;

}
