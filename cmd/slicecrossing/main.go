// Package main for check common values in two integers slices.
package main

import (
	"bufio"
	"fmt"
	"os"

	"smart_utilities/internal/slicecrossing"
)

// Main function for check common values in two integers slices.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first slice of integers: ")
	s, err := slicecrossing.ParseStringToSlice(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Enter second slice of integers: ")
	s2, err := slicecrossing.ParseStringToSlice(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Result: ")
	resSlice, len := slicecrossing.GetCrossedIndices(s, s2)
	if len > 0 {
		resSlice.PrintNumSlice(os.Stdout)
	} else {
		fmt.Print("Empty intersection\n")
	}
}
