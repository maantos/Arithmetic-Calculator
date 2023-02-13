package main

import (
	"errors"
	"strings"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		root     *ASTNode
		expected int
		err      error
	}{
		{&ASTNode{Operation: Digit, Value: 1}, 1, nil}, //"1"
		{&ASTNode{Operation: Div, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Digit, Value: 0}}, 0, errors.New("division by zero")},
		{&ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Digit, Value: 2}}, 3, nil},
		{&ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 4}, Right: &ASTNode{Operation: Minus, Left: &ASTNode{Operation: Digit, Value: 5}, Right: &ASTNode{Operation: Div, Left: &ASTNode{Operation: Digit, Value: 8}, Right: &ASTNode{Operation: Digit, Value: 2}}}}, 13, nil},                                                                                 //"4-5+8/2"
		{&ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Mult, Left: &ASTNode{Operation: Digit, Value: 2}, Right: &ASTNode{Operation: Digit, Value: 3}}}, 7, nil},                                                                                                                                                                //"1+2*3"
		{&ASTNode{Operation: Mult, Left: &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Digit, Value: 2}}, Right: &ASTNode{Operation: Digit, Value: 3}}, 9, nil},                                                                                                                                                                // "(1+2)*3"
		{&ASTNode{Operation: Minus, Left: &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 4}, Right: &ASTNode{Operation: Mult, Left: &ASTNode{Operation: Digit, Value: 5}, Right: &ASTNode{Operation: Minus, Left: &ASTNode{Operation: Digit, Value: 7}, Right: &ASTNode{Operation: Digit, Value: 3}}}}, Right: &ASTNode{Operation: Digit, Value: 2}}, 22, nil}, //"(4+5*(7-3))-2"
	}

	for _, test := range tests {
		result, err := Evaluate(test.root)
		if err != nil {
			if !strings.Contains(err.Error(), test.err.Error()) {
				t.Errorf("unexpected error message: %v, expected: %v", err, test.err)
			} else if result != test.expected {
				t.Errorf("unexpected result: %v, expected: %v", result, test.expected)
			}
		}
	}
}
