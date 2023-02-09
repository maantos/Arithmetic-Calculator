package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Token types
const (
	Error             int = 0
	Plus                  = 1
	Minus                 = 2
	Mul                   = 3
	Div                   = 4
	EndOfText             = 5
	OpenParenthesis       = 6
	ClosedParenthesis     = 7
	Number                = 8
)

type Token struct {
	TokenType int
	Value     int
}

type Parser struct {
	Token      Token
	expression string
	index      int
}

func (p *Parser) Expression() {
	p.Term()
	p.Expression1()
}

func (p *Parser) Expression1() {
	switch p.Token.TokenType {
	case Plus:
		p.GetNextToken()
		p.Term()
		p.Expression1()
		break
	case Minus:
		p.GetNextToken()
		p.Term()
		p.Expression1()
		break
	}
}

func (p *Parser) Term() {
	p.Factor()
	p.Term1()
}

func (p *Parser) Term1() {
	switch p.Token.TokenType {
	case Mul:
		p.GetNextToken()
		p.Factor()
		p.Term1()
		break
	case Div:
		p.GetNextToken()
		p.Factor()
		p.Term1()
		break
	}
}

func (p *Parser) Factor() {
	switch p.Token.TokenType {
	case OpenParenthesis:
		p.GetNextToken()
		p.Expression()
		p.Match(')')
		break
	case Number:
		p.GetNextToken()
		break
	default:
		fmt.Errorf("Unexpected token %d", p.Token)
	}
}

func (p *Parser) Match(expected byte) {
	if p.expression[p.index-1] == expected {
		p.GetNextToken()
	} else {
		fmt.Errorf("match failed")
	}

}

func (p *Parser) GetNextToken() {

	//end of text check ???
	c := p.expression[p.index]
	if unicode.IsNumber(rune(c)) {
		p.Token.TokenType = Number
		p.Token.Value, _ = strconv.Atoi(string(c))
		p.index++
	}

	p.Token.TokenType = Error

	switch x := p.expression[p.index]; x {
	case '+':
		p.Token.TokenType = Plus
		break
	case '-':
		p.Token.TokenType = Minus
		break
	case '*':
		p.Token.TokenType = Mul
		break
	case '/':
		p.Token.TokenType = Div
		break
	case '(':
		p.Token.TokenType = OpenParenthesis
		break
	case ')':
		p.Token.TokenType = ClosedParenthesis
		break
	}

	if p.Token.TokenType != Error {
		p.index++
	} else {
		fmt.Errorf("Unexpected symbol on index %d", p.index)
	}
}

func main() {

	// if len(os.Args) > 1 {
	// 	fmt.Println(os.Args)
	// }

	fmt.Println("hello")
}
