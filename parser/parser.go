package parser

type Node struct {
	Value      string
	Next       *Node
	Children   []*Node
	Attributes map[string]bool
}

func parse() {

}
