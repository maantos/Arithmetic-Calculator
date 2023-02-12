package main

import (
	"testing"
)

func TestParser(t *testing.T) {

	tests := []struct {
		input    string
		expected *ASTNode
	}{
		{"", nil},
		{" 1", nil},
		{"10+1", nil},
		{"-10", nil},
		{"1", &ASTNode{Operation: Digit, Value: 1}},
		{"1+2", &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Digit, Value: 2}}},
		{"1+2*3", &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Mult, Left: &ASTNode{Operation: Digit, Value: 2}, Right: &ASTNode{Operation: Digit, Value: 3}}}},
		{"(1+2)*3", &ASTNode{Operation: Mult, Left: &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 1}, Right: &ASTNode{Operation: Digit, Value: 2}}, Right: &ASTNode{Operation: Digit, Value: 3}}},
		{"(4+5*(7-3))-2", &ASTNode{Operation: Minus, Left: &ASTNode{Operation: Plus, Left: &ASTNode{Operation: Digit, Value: 4}, Right: &ASTNode{Operation: Mult, Left: &ASTNode{Operation: Digit, Value: 5}, Right: &ASTNode{Operation: Minus, Left: &ASTNode{Operation: Digit, Value: 7}, Right: &ASTNode{Operation: Digit, Value: 3}}}}, Right: &ASTNode{Operation: Digit, Value: 2}}},
	}

	for _, test := range tests {
		p := NewParser(test.input)
		result, _ := p.Parse()

		if !equal(result, test.expected) {
			t.Errorf("Parse(%q) = %v, want %v.", test.input, result, test.expected)
		}
	}
}

func equal(a, b *ASTNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Operation != b.Operation || a.Value != b.Value {
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
