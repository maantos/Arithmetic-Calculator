package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type Parser struct {
	Token      byte
	expression string
	index      int
}

// Grammar for expresion E -> T {+|- T}
func (p *Parser) parseE() *ASTNode {
	termNode := p.parseT()
	if termNode == nil {
		fmt.Println(errors.New("error while parsing term"))
		return nil
	}
	for {
		if p.Token == '+' {
			p.GetNextToken('+')
			term2Node := p.parseT()

			if term2Node == nil {
				fmt.Println(errors.New("error while parsing term"))
				return nil
			}

			termNode = NewNode("Plus", termNode, term2Node)
		} else if p.Token == '-' {
			p.GetNextToken('-')
			term2Node := p.parseT()

			if term2Node == nil {
				fmt.Println(errors.New("error while parsing term"))
				return nil
			}
			termNode = NewNode("Minus", termNode, term2Node)
		} else {
			return termNode
		}
	}
}

// Grammar for Term  T -> F {*|/ F}
func (p *Parser) parseT() *ASTNode {

	fNode := p.parseF()

	if fNode == nil {
		fmt.Println(errors.New("error while parsing factor"))
		return nil
	}
	for {
		if p.Token == '*' {
			p.GetNextToken('*')
			f2Node := p.parseF()
			if f2Node == nil {
				fmt.Println(errors.New("error while parsing factor"))
				return nil
			}
			fNode = NewNode("Mult", fNode, f2Node)

		} else if p.Token == '/' {
			p.GetNextToken('/')
			f2Node := p.parseF()
			if f2Node == nil {
				fmt.Println(errors.New("error while parsing factor"))
				return nil
			}
			fNode = NewNode("Div", fNode, f2Node)
		} else {
			return fNode
		}
	}
}

// Grammar for Factor  F -> (E) | Digit
func (p *Parser) parseF() *ASTNode {

	if unicode.IsNumber(rune(p.Token)) {
		x, _ := strconv.Atoi(string(p.Token))
		p.GetNextToken(p.Token)
		node := NewNode("Digit", nil, nil)
		node.Value = x
		return node

	} else if p.Token == '(' {
		p.GetNextToken('(')
		node := p.parseE()

		if node == nil {
			fmt.Println(errors.New("error while parsing expression"))
			return nil
		}

		if p.Token == ')' {
			p.GetNextToken(')')
			return node
		} else {
			fmt.Println(errors.New("closing bracket is missing"))
			return nil
		}
	}
	fmt.Println(errors.New("unexpected symbol"))
	return nil
}

func (p *Parser) GetNextToken(expected byte) {
	if p.index < len(p.expression)-1 {
		p.index++
		p.Token = p.expression[p.index]
	} else {
		p.Token = '\n'
	}
}

func (p *Parser) Parse(expression string) (*ASTNode, error) {
	p.index = 0
	p.expression = expression
	p.Token = p.expression[p.index]

	root := p.parseE()

	// if err != nil {
	// 	return -1, err
	// }

	if p.Token != '\n' {
		return nil, errors.New("error while Parsing")
	}

	return root, nil
}
