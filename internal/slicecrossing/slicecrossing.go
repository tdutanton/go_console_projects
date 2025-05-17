// Package slicecrossing provides functionality for working with integer slices,
// including parsing input, printing, and finding intersections between slices.
package slicecrossing

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Num represents a numeric value as a distinct type from built-in integers.
// This provides type safety when working with numeric slices.
type Num int

// NumSlice is a slice of Num values that provides additional methods
// for working with collections of numbers.
type NumSlice []Num

// ParseStringToSlice reads a line from the input reader and converts it
// to a NumSlice. The input should be space-separated integers.
//
// Parameters:
//   - reader: A buffered reader to read input from (typically os.Stdin)
//
// Returns:
//   - NumSlice: The parsed slice of numbers
//   - error: Returns an error if reading fails or contains invalid input
//
// Example input: "1 2 3 4\n" will return NumSlice{1, 2, 3, 4}
func ParseStringToSlice(reader *bufio.Reader) (NumSlice, error) {
	result := NumSlice{}
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("input error")
	}
	words := strings.Fields(input)
	for _, v := range words {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("invalid input")
		}
		result = append(result, Num(num))
	}
	return result, nil
}

// PrintNumSlice writes the contents of the NumSlice to the specified writer
// as a space-separated string followed by a newline.
//
// Parameters:
//   - writer: An io.Writer to output the results (typically os.Stdout)
//
// Example output for NumSlice{1, 2, 3}: "1 2 3\n"
func (n NumSlice) PrintNumSlice(writer io.Writer) {
	strs := make([]string, len(n))
	for i, num := range n {
		strs[i] = strconv.Itoa(int(num))
	}
	fmt.Fprintln(writer, strings.Join(strs, " "))
}

// GetCrossedIndices returns the intersection of two NumSlices, preserving
// the order of elements as they appear in the first slice (n1). Each element
// will appear only once in the result, even if it appears multiple times
// in either input slice.
//
// Parameters:
//   - n1: The first slice of numbers (order will be preserved)
//   - n2: The second slice of numbers for comparison
//
// Returns:
//   - NumSlice: The intersection of n1 and n2 in n1's original order
//   - int: The count of elements in the intersection
//
// Example:
//
//	n1: [1, 2, 3, 4], n2: [3, 4, 5, 6] → returns [3, 4], 2
func GetCrossedIndices(n1, n2 NumSlice) (NumSlice, int) {
	result := NumSlice{}
	map2 := make(map[Num]bool)
	for _, num := range n2 {
		map2[num] = true
	}
	for _, num := range n1 {
		if map2[num] {
			result = append(result, num)
			map2[num] = false
		}
	}
	return result, len(result)
}
