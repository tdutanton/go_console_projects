package slicecrossing

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestPrintWords(t *testing.T) {
	words := NumSlice{44, 55, 66}

	var buf bytes.Buffer
	words.PrintNumSlice(&buf)

	expected := "44 55 66\n"
	if buf.String() != expected {
		t.Errorf("Got: %q\nWant: %q", buf.String(), expected)
	}
}

func TestParseStringToSlice(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    NumSlice
		expectErr   bool
		errContains string
	}{
		{
			name:     "valid single number",
			input:    "42\n",
			expected: NumSlice{42},
		},
		{
			name:     "valid multiple numbers",
			input:    "1 2 3 4 5\n",
			expected: NumSlice{1, 2, 3, 4, 5},
		},
		{
			name:     "valid with leading/trailing spaces",
			input:    "  10  20  30  \n",
			expected: NumSlice{10, 20, 30},
		},
		{
			name:        "invalid non-number input",
			input:       "abc\n",
			expectErr:   true,
			errContains: "invalid input",
		},
		{
			name:        "mixed valid and invalid",
			input:       "1 2 abc 4\n",
			expectErr:   true,
			errContains: "invalid input",
		},
		{
			name:     "empty input",
			input:    "\n",
			expected: NumSlice{},
		},
		{
			name:     "whitespace only",
			input:    "   \n",
			expected: NumSlice{},
		},
		{
			name:        "reader error (no newline)",
			input:       "1 2 3",
			expectErr:   true,
			errContains: "input error",
		},
		{
			name:        "invalid number format",
			input:       "1.23\n",
			expectErr:   true,
			errContains: "invalid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			result, err := ParseStringToSlice(reader)

			if tt.expectErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.errContains)) {
					t.Errorf("Expected error to contain %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Fatalf("Expected length %d, got %d", len(tt.expected), len(result))
			}

			for i, num := range result {
				if num != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], num)
				}
			}
		})
	}
}

func TestGetCrossedIndices(t *testing.T) {
	tests := []struct {
		name        string
		n1          NumSlice
		n2          NumSlice
		expected    NumSlice
		expectedLen int
	}{
		{
			name:        "simple intersection",
			n1:          NumSlice{1, 2, 3, 4},
			n2:          NumSlice{3, 4, 5, 6},
			expected:    NumSlice{3, 4},
			expectedLen: 2,
		},
		{
			name:        "no intersection",
			n1:          NumSlice{1, 2, 3},
			n2:          NumSlice{4, 5, 6},
			expected:    NumSlice{},
			expectedLen: 0,
		},
		{
			name:        "empty first slice",
			n1:          NumSlice{},
			n2:          NumSlice{1, 2, 3},
			expected:    NumSlice{},
			expectedLen: 0,
		},
		{
			name:        "empty second slice",
			n1:          NumSlice{1, 2, 3},
			n2:          NumSlice{},
			expected:    NumSlice{},
			expectedLen: 0,
		},
		{
			name:        "both slices empty",
			n1:          NumSlice{},
			n2:          NumSlice{},
			expected:    NumSlice{},
			expectedLen: 0,
		},
		{
			name:        "duplicates in first slice",
			n1:          NumSlice{1, 2, 2, 3},
			n2:          NumSlice{2, 3, 4},
			expected:    NumSlice{2, 3},
			expectedLen: 2,
		},
		{
			name:        "duplicates in second slice",
			n1:          NumSlice{1, 2, 3},
			n2:          NumSlice{2, 2, 3, 3},
			expected:    NumSlice{2, 3},
			expectedLen: 2,
		},
		{
			name:        "preserve order from first slice",
			n1:          NumSlice{3, 1, 4, 2},
			n2:          NumSlice{1, 2, 5},
			expected:    NumSlice{1, 2},
			expectedLen: 2,
		},
		{
			name:        "negative numbers",
			n1:          NumSlice{-1, -2, -3},
			n2:          NumSlice{-2, -3, -4},
			expected:    NumSlice{-2, -3},
			expectedLen: 2,
		},
		{
			name:        "mixed positive and negative",
			n1:          NumSlice{1, -2, 3},
			n2:          NumSlice{-2, 3, -4},
			expected:    NumSlice{-2, 3},
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, length := GetCrossedIndices(tt.n1, tt.n2)

			// Check length first
			if length != tt.expectedLen {
				t.Errorf("Expected length %d, got %d", tt.expectedLen, length)
			}

			// Check slice contents
			if len(result) != len(tt.expected) {
				t.Fatalf("Expected result length %d, got %d", len(tt.expected), len(result))
			}

			for i, num := range result {
				if num != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], num)
				}
			}
		})
	}
}
