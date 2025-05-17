// Package main for Smart Calculator function.
package main

import (
	"bufio"
	"fmt"
	"os"

	"smart_utilities/internal/calculator"
)

// Main function for Smart Calculator function.
// Takes left operand, operator and right operand with ENTER after every one input
func main() {
	reader := bufio.NewReader(os.Stdin)
	left, _ := calculator.ParseOperand("Input left operand: ", reader, os.Stdout)
	operator, _ := calculator.ParseOperator("Input one of operations  + - * /: ", reader, os.Stdout)
	right, _ := calculator.ParseOperand("Input right operand: ", reader, os.Stdout)
	res, err := calculator.CreateOperation(left, operator, right)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Result:", res)
}
