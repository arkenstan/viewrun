package pipeline

type Directive struct {
	keyword string
	source  string
	pattern []string
	target  string
}

func ProcessDirective(command []string) {

	for _, literal := range command {

		println("PROCESSING", literal)

	}

}
