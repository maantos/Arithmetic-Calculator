package main

import (
	"errors"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected any
		err      error
	}{
		{"(4+5*(7-3))-2", 22, nil},
		{"1+2*3", 7, nil},
		{"(1+2)*3", 9, nil},
		{"-10", -1, errors.New("error while parsing T")},
		{"12+1", -1, errors.New("Error while Parsing")},
		{"1 + 2 (", -1, errors.New("error while Parsing")},
	}
	p := &Parser{}
	for _, test := range tests {
		result, err := p.Parse(test.input)
		if err != test.err {
			t.Errorf("parse(%q) = %d, %v, expected %d, %v", test.input, result, err, test.expected, test.err)
		}
	}

}
