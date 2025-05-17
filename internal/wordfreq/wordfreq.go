// Package wordfreq provides utilities for analyzing word frequency in text.
//
// It includes functions to:
//   - Parse user input for a desired number of top words (ParseK),
//   - Count word frequencies in a given string (GetWordsMap),
//   - Sort words by frequency (getWordsSlice),
//   - Return the top-k most frequent words (GetResultWordsSlice).
//
// The Word type is a custom string type used as a map key for clarity and consistency.
package wordfreq

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// Word represents a single word as a distinct type.
type Word string

// WordSlice represents a slice of words as a distinct type.
type WordSlice []Word

// wordsCount returns the number of words in the given string.
// It splits the string using whitespace as the delimiter.
func wordsCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// ParseK prompts the user (via writer) with a message (s) and reads input from reader,
// repeatedly asking until a valid positive integer is entered.
// Returns the integer value or an error if reading fails.
func ParseK(s string, reader *bufio.Reader, writer io.Writer) (int, error) {
	fmt.Fprint(writer, s)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, errors.New("input error")
		}
		input = strings.TrimSpace(input)
		v, err := strconv.Atoi(input)
		if err != nil || v <= 0 {
			fmt.Fprint(writer, "Invalid input. Please give me correct K value: ")
			continue
		}
		return v, nil
	}
}

// GetWordsMap takes a string, splits it into words, and returns a map
// where each key is a Word and each value is the number of occurrences of that word.
func GetWordsMap(s string) map[Word]int {
	result := map[Word]int{}
	words := strings.Fields(s)
	for _, v := range words {
		result[Word(v)]++
	}
	return result
}

// getWordsSlice takes a map of Words to their frequencies and returns a slice
// of Words sorted in descending order of frequency.
func getSortedWords(m map[Word]int) WordSlice {
	result := make([]Word, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		return m[result[i]] > m[result[j]]
	})
	return result
}

// GetResultWordsSlice returns the first k words from the slice w.
// If k is greater than the length of w, it returns all words in w.
func GetResultWordsSlice(s string, k int) WordSlice {
	m := GetWordsMap(s)
	sorted := getSortedWords(m)
	result := WordSlice{}
	if k >= len(sorted) {
		result = sorted[:]
	} else if k < len(sorted) && len(sorted) > 0 {
		result = sorted[:k]
	}
	return result
}

// PrintWords method for WordSlice to print all words in the slice
func (w WordSlice) PrintWords(writer io.Writer) {
	strs := make([]string, len(w))
	for i, word := range w {
		strs[i] = string(word)
	}
	fmt.Fprintln(writer, strings.Join(strs, " "))
}
