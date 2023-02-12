package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Result struct {
	value      int
	expression string
}

type ErrorResult struct {
	err        error
	expression string
}

func main() {
	file, err := os.Open("../expressions.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := make(chan Result)
	errors := make(chan ErrorResult)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		p := NewParser(line)
		counter++
		go func() {
			root, err := p.Parse()
			if err != nil {
				errResult := &ErrorResult{err: err, expression: line}
				errors <- *errResult
				return
			}
			result, err := Evaluate(root)
			if err != nil {
				errResult := &ErrorResult{err: err, expression: line}
				errors <- *errResult
				return
			}
			r := &Result{value: result, expression: line}
			results <- *r
		}()
	}

	for i := 0; i < counter; i++ {
		select {
		case result := <-results:
			fmt.Printf("equation: %s with value %d\n", result.expression, result.value)
			//fmt.Printf("equation=%d",result.value)
		case err := <-errors:
			fmt.Printf("equation: %s with Error: %s\n", err.expression, err.err)
			//fmt.Println(err)
		}
	}
}
