package main

import (
	"errors"
)

type Operation int

const (
	Plus Operation = iota
	Minus
	Mult
	Div
	Digit
)

type ASTNode struct {
	Operation Operation
	Value     int
	Left      *ASTNode
	Right     *ASTNode
}

func NewNode(operation Operation, left, right *ASTNode) *ASTNode {
	return &ASTNode{
		Operation: operation,
		Value:     0,
		Left:      left,
		Right:     right,
	}
}

func Evaluate(node *ASTNode) (int, error) {

	if node == nil {
		return 0, errors.New("node doesnt exist")
	}

	if node.Operation == Digit {
		return node.Value, nil
	}

	leftNode, err := Evaluate(node.Left)

	if err != nil {
		return 0, err
	}

	rightNode, err := Evaluate(node.Right)

	if err != nil {
		return 0, err
	}

	var result int
	switch node.Operation {
	case Plus:
		result = leftNode + rightNode
	case Minus:
		result = leftNode - rightNode
	case Mult:
		result = leftNode * rightNode
	case Div:
		if rightNode == 0 {
			return 0, errors.New("division by zero")
		}
		result = leftNode / rightNode
	default:
		return 0, errors.New("unknown operation")
	}

	return result, nil
}

// func printTree(node *ASTNode) string {

// 	if node == nil {
// 		return ""
// 	}

// 	var operation rune
// 	switch node.Operation {
// 	case Digit:
// 		return strconv.Itoa(node.Value)
// 	case Plus:
// 		operation = '+'
// 	case Minus:
// 		operation = '-'
// 	case Mult:
// 		operation = '*'
// 	case Div:
// 		operation = '/'
// 	}
// 	return fmt.Sprintf("(" + printTree(node.Left) + string(operation) + printTree(node.Right) + ")")
// }
