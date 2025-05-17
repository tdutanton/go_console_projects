// Package calculator provides basic arithmetic operations
// and utilities to parse operands and operators from user input
package calculator

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// operators map with available math operators
var operators = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

// add left + right, error = nil
func add(left, right float64) (float64, error) {
	result := left + right
	return math.Round(result*1000) / 1000, nil
}

// sub left - right, error = nil
func sub(left, right float64) (float64, error) {
	result := left - right
	return math.Round(result*1000) / 1000, nil
}

// mult left * right, error = nil
func mult(left, right float64) (float64, error) {
	result := left * right
	return math.Round(result*1000) / 1000, nil
}

// div left / right
// left / 0 -> return error
func div(left, right float64) (float64, error) {
	if right == 0 {
		return 0, errors.New("unfortunately you can't divide by zero :-(")
	}
	result := left / right
	return math.Round(result*1000) / 1000, nil
}

// stringLength length of string
func stringLength(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

// CreateOperation returns one of available math function
// dependens on input math operand
func CreateOperation(left float64, r rune, right float64) (float64, error) {
	switch r {
	case '+':
		return add(left, right)
	case '-':
		return sub(left, right)
	case '*':
		return mult(left, right)
	case '/':
		return div(left, right)
	default:
		return 0, errors.New("unknown operation")
	}
}

// ParseOperator get math operator from string.
// String 's' is the prompt message shown to the user in the command line
func ParseOperator(s string, reader *bufio.Reader, writer io.Writer) (rune, error) {
	fmt.Fprint(writer, s)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, errors.New("input error")
		}
		input = strings.TrimSpace(input)
		if stringLength(input) != 1 || !operators[rune(input[0])] {
			fmt.Fprint(writer, "Invalid input. Please give me correct operator (+, -, *, /): ")
			continue
		}
		return rune(input[0]), nil
	}
}

// ParseOperand get operand (left or right) from string.
// String 's' is the prompt message shown to the user in the command line
func ParseOperand(s string, reader *bufio.Reader, writer io.Writer) (float64, error) {
	fmt.Fprint(writer, s)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, errors.New("input error")
		}
		input = strings.TrimSpace(input)
		v, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprint(writer, "Invalid input. Please give me correct operand: ")
			continue
		}
		return v, nil
	}
}
