package main

import (
	"fmt"
)

func main() {

	p := Parser{}
	expression := "(4+5*(7-3))-2"
	root, err := p.Parse(expression)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(root.print())

	fmt.Println(Evaluate(root))

	// fmt.Printf("\nEquation = %d\n", result)
}
