// Package main for analyzing word frequency in text.
package main

import (
	"bufio"
	"fmt"
	"os"

	"smart_utilities/internal/wordfreq"
)

// Main function for analyzing word frequency in text.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a string with some words with space between it: ")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	k, err := wordfreq.ParseK("Input number of unique words to show: ", reader, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res := wordfreq.GetResultWordsSlice(s, k)
	fmt.Print("Result: ")
	res.PrintWords(os.Stdout)
}
