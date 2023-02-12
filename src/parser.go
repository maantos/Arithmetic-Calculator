package main

import (
	"errors"
	"strconv"
	"unicode"
)

type Parser struct {
	token      byte
	expression string
	index      int
}

// Grammar for expresion E -> T {+|- T}
func (p *Parser) parseE() (*ASTNode, error) {
	termNode, err := p.parseT()
	if err != nil {
		//fmt.Println(errors.New("error while parsing term"))
		return nil, err
	}
	for {
		if p.token == '+' {
			p.GetNextToken()
			term2Node, err := p.parseT()

			if err != nil {
				//fmt.Println(errors.New("error while parsing term"))
				return nil, err
			}

			termNode = NewNode(Plus, termNode, term2Node)
		} else if p.token == '-' {
			p.GetNextToken()
			term2Node, err := p.parseT()

			if err != nil {
				return nil, err
			}
			termNode = NewNode(Minus, termNode, term2Node)
		} else {
			return termNode, nil
		}
	}
}

// Grammar for Term  T -> F {*|/ F}
func (p *Parser) parseT() (*ASTNode, error) {

	fNode, err := p.parseF()

	if err != nil {
		//fmt.Println(errors.New("error while parsing factor"))
		return nil, err
	}
	for {
		if p.token == '*' {
			p.GetNextToken()
			f2Node, err := p.parseF()
			if err != nil {
				//fmt.Println(errors.New("error while parsing factor"))
				return nil, err
			}
			fNode = NewNode(Mult, fNode, f2Node)

		} else if p.token == '/' {
			p.GetNextToken()
			f2Node, err := p.parseF()
			if err != nil {
				//fmt.Println(errors.New("error while parsing factor"))
				return nil, err
			}
			fNode = NewNode(Div, fNode, f2Node)
		} else {
			return fNode, nil
		}
	}
}

// Grammar for Factor  F -> (E) | Digit
func (p *Parser) parseF() (*ASTNode, error) {

	if unicode.IsNumber(rune(p.token)) {
		x, err := strconv.Atoi(string(p.token))
		if err != nil {
			//fmt.Errorf("error while converting digit to int: %v", err)
			return nil, errors.New("error while converting digit to int")
		}
		p.GetNextToken()
		node := NewNode(Digit, nil, nil)
		node.Value = x
		return node, nil

	} else if p.token == '(' {
		p.GetNextToken()
		node, err := p.parseE()

		if err != nil {
			return nil, errors.New("error while parsing expression")
		}

		if p.token == ')' {
			p.GetNextToken()
			return node, nil
		} else {
			return nil, errors.New("closing bracket is missing")
		}
	}
	return nil, errors.New("unexpected or missing symbol")
}

func (p *Parser) GetNextToken() {
	if p.index < len(p.expression)-1 {
		p.index++
		p.token = p.expression[p.index]
	} else {
		p.token = '\n'
	}
}

func (p *Parser) Parse() (*ASTNode, error) {

	p.GetNextToken()
	root, err := p.parseE()

	if err != nil {
		return nil, err //errors.New("error parsing expression")
	}

	if p.token != '\n' {
		return nil, errors.New("didnt consume all tokens while parsing")
	}

	return root, nil
}

func NewParser(expression string) *Parser {
	return &Parser{
		expression: expression,
		index:      -1,
	}
}
