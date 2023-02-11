package main

import (
	"errors"
	"strconv"
	"unicode"
)

type Parser struct {
	Token      byte
	expression string
	index      int
}

// Grammar for expresion E -> T {+|- T}
func (p *Parser) parseE() (int, error) {
	a, err := p.parseT()
	if err != nil {
		return -1, errors.New("error while parsing T")
	}
	for {
		if p.Token == '+' {
			p.GetNextToken('+')
			b, err := p.parseT()
			if err != nil {
				return -1, errors.New("error while parsing T")
			}
			a += b
		} else if p.Token == '-' {
			p.GetNextToken('-')
			b, err := p.parseT()
			if err != nil {
				return -1, errors.New("error while parsing T")
			}
			a -= b
		} else {
			return a, nil
		}
	}
}

// Grammar for Term  T -> F {*|/ F}
func (p *Parser) parseT() (int, error) {
	a, err := p.parseF()
	if err != nil {
		return -1, errors.New("error while parsing F")
	}
	for {
		if p.Token == '*' {
			p.GetNextToken('*')
			b, err := p.parseF()

			if err != nil {
				return -1, errors.New("error while parsing F")
			}

			a *= b
		} else if p.Token == '/' {
			p.GetNextToken('/')
			b, err := p.parseT()

			if err != nil {
				return -1, errors.New("error while parsing T")
			}

			a /= b
		} else {
			return a, nil
		}
	}
}

// Grammar for Factor  F -> (E) | Digit
func (p *Parser) parseF() (int, error) {

	if unicode.IsNumber(rune(p.Token)) {
		x, _ := strconv.Atoi(string(p.Token))
		p.GetNextToken(p.Token)
		return x, nil
	} else if p.Token == '(' {
		p.GetNextToken('(')
		a, err := p.parseE()
		if err != nil {
			return -1, errors.New("error while parsing E")
		}
		//error check
		if p.Token == ')' {
			p.GetNextToken(')')
			return a, nil
		} else {
			return -1, errors.New("missing closing bracket")
		}
	}
	return -1, errors.New("unexpected symbol")
}

func (p *Parser) GetNextToken(expected byte) {
	if p.index < len(p.expression)-1 {
		p.index++
		p.Token = p.expression[p.index]
	} else {
		p.Token = '\n'
	}
}

func (p *Parser) Parse(expression string) (int, error) {
	p.index = 0
	p.expression = expression
	p.Token = p.expression[p.index]

	result, err := p.parseE()

	if err != nil {
		return -1, err
	}

	if p.Token != '\n' {
		return -1, errors.New("error while Parsing")
	}

	return result, nil
}
