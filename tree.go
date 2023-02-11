package main

var Operations = []string{"Add", "Subtract", "Mult", "Div", "Digit"}

type Node struct {
	Operation string
	Value     int
	Left      *Node
	Right     *Node
}

func NewNode(operation string, left, right *Node) *Node {
	return &Node{
		Operation: operation,
		Left:      left,
		Right:     right,
	}
}

func (n *Node) getLeft() Node {
	return *n.Left
}
func (n *Node) getRight() Node {
	return *n.Right
}
func (n *Node) getOperation() string {
	return n.Operation
}

// func (n *Node) print() {

// }
