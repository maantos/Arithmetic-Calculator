package main

import (
	"fmt"
	"strconv"
)

var Operations = []string{"Plus", "Minus", "Mult", "Div", "Digit"}

type ASTNode struct {
	Operation string
	Value     int
	Left      *ASTNode
	Right     *ASTNode
}

func Evaluate(node *ASTNode) int {

	if node.Operation == "Digit" {
		return node.Value
	}

	leftNode := Evaluate(node.Left)
	rightNode := Evaluate(node.Right)

	var result int
	switch node.Operation {
	case "Plus":
		result = leftNode + rightNode
	case "Minus":
		result = leftNode - rightNode
	case "Mult":
		result = leftNode * rightNode
	case "Div":
		result = leftNode / rightNode
	}

	return result
}

func (a *ASTNode) print() string {

	var operation rune
	switch a.Operation {
	case "Digit":
		return strconv.Itoa(a.Value)
	case "Plus":
		operation = '+'
	case "Minus":
		operation = '-'
	case "Mult":
		operation = '*'
	case "Div":
		operation = '/'
	}
	return fmt.Sprintf("(" + a.Left.print() + string(operation) + a.Right.print() + ")")
}

func NewNode(operation string, left, right *ASTNode) *ASTNode {
	return &ASTNode{
		Operation: operation,
		Value:     0,
		Left:      left,
		Right:     right,
	}
}

// func (n *Node) getLeft() Node {
// 	return *n.Left
// }
// func (n *Node) getRight() Node {
// 	return *n.Right
// }
// func (n *Node) getOperation() string {
// 	return n.Operation
// }

// func (n *Node) print() {

// }
