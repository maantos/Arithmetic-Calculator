package main

import (
	"fmt"
)

func main() {

	p := Parser{}
	expression := "-10"
	result, err := p.Parse(expression)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nEquation = %d\n", result)
}
