package calculator

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// ExampleCreateOperation
func ExampleCreateOperation() {
	result, _ := CreateOperation(3, '+', 5)
	fmt.Println(result)
	// Output: 8
}

func Test_add(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 float64
		expected   float64
		err        error
	}{
		{"goodInt", 2, 2, 4, nil},
		{"goodFloat", 1.15, 2.23, 3.38, nil},
		{"truncateFloat", 1.1234, 2.1234, 3.247, nil},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			if got, _ := add(d.num1, d.num2); got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 float64
		expected   float64
		err        error
	}{
		{"goodInt", 2, 2, 0, nil},
		{"goodFloat", 1.15, 2.23, -1.08, nil},
		{"truncateFloat", 3.8976, 2.2453, 1.652, nil},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			if got, _ := sub(d.num1, d.num2); got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
		})
	}
}

func Test_mult(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 float64
		expected   float64
		err        error
	}{
		{"goodInt", 2, 2, 4, nil},
		{"goodFloat", 1.15, 2.23, 2.565, nil},
		{"truncateFloat", 3.8976, 2.2453, 8.751, nil},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			if got, _ := mult(d.num1, d.num2); got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
		})
	}
}

func Test_div(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 float64
		expected   float64
		expErr     bool
		errMsg     string
	}{
		{"goodInt", 2, 2, 1, false, ""},
		{"goodFloat", 6.2, 2, 3.1, false, ""},
		{"truncateFloat", 3.8976, 3, 1.299, false, ""},
		{"devZero", 4, 0, 0, true, "unfortunately you can't divide by zero :-("},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			got, err := div(d.num1, d.num2)
			if d.expErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if err.Error() != d.errMsg {
					t.Errorf("Expected %s, got %s", d.errMsg, err.Error())
				}
			}
			if got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
		})
	}
}

func Test_stringLength(t *testing.T) {
	tests := []struct {
		name     string
		argSring string
		expected int
	}{
		{"positiveString", "Hello", 5},
		{"zeroString", "", 0},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			if got := stringLength(d.argSring); got != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, got)
			}
		})
	}
}

func TestCreateOperation(t *testing.T) {
	tests := []struct {
		name       string
		num1, num2 float64
		op         rune
		expected   float64
		expErr     bool
		errMsg     string
	}{
		{"add", 2, 2, '+', 4, false, ""},
		{"sub", 2, 2, '-', 0, false, ""},
		{"mult", 2, 3, '*', 6, false, ""},
		{"div", 8, 2, '/', 4, false, ""},
		{"unknown", 5, 2, '&', 0, true, "unknown operation"},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			got, err := CreateOperation(d.num1, d.op, d.num2)
			if d.expErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if err.Error() != d.errMsg {
					t.Errorf("Expected %s, got %s", d.errMsg, err.Error())
				}
			}
			if got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
		})
	}
}

func TestParseOperator(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     rune
		expErr       bool
		expReaderErr bool
		errContains  string
	}{
		{
			name:     "valid operator +",
			input:    "+\n",
			expected: '+',
		},
		{
			name:     "valid operator -",
			input:    "-\n",
			expected: '-',
		},
		{
			name:        "invalid operator (multiple chars)",
			input:       "++\n+\n",
			expErr:      true,
			errContains: "Invalid input",
		},
		{
			name:        "invalid operator (wrong char)",
			input:       "a\n+\n",
			expErr:      true,
			errContains: "Invalid input",
		},
		{
			name:         "invalid reader",
			input:        "+++",
			expReaderErr: true,
			errContains:  "input error",
		},
	}
	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(d.input))
			var output bytes.Buffer
			got, err := ParseOperator("Enter operator: ", reader, &output)
			if !d.expErr && got != d.expected {
				t.Errorf("Expected %c, got %c", d.expected, got)
			}
			if d.expErr {
				outStr := output.String()
				if !strings.Contains(outStr, "Invalid input") {
					t.Errorf("Expected error message for invalid input, got: %s", outStr)
				}
			}
			if d.expReaderErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if err.Error() != d.errContains {
					t.Errorf("Expected %s, got %s", d.errContains, err.Error())
				}
			}
		})
	}
}

func TestParseOperand(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     float64
		expErr       bool
		expReaderErr bool
		errContains  string
	}{
		{
			name:     "valid left operand 0",
			input:    "0\n",
			expected: 0,
		},
		{
			name:     "valid right operand 0",
			input:    "0\n",
			expected: 0,
		},
		{
			name:        "invalid left operand",
			input:       "45f\n45\n",
			expErr:      true,
			errContains: "Invalid input",
		},
		{
			name:        "invalid right operand",
			input:       "47a\n47\n",
			expErr:      true,
			errContains: "Invalid input",
		},
		{
			name:         "invalid reader",
			input:        "47a",
			expReaderErr: true,
			errContains:  "input error",
		},
	}
	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(d.input))
			var output bytes.Buffer
			got, err := ParseOperand("Enter operand: ", reader, &output)
			if !d.expErr && got != d.expected {
				t.Errorf("Expected %f, got %f", d.expected, got)
			}
			if d.expErr {
				outStr := output.String()
				if !strings.Contains(outStr, "Invalid input") {
					t.Errorf("Expected error message for invalid input, got: %s", outStr)
				}
			}
			if d.expReaderErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if err.Error() != d.errContains {
					t.Errorf("Expected %s, got %s", d.errContains, err.Error())
				}
			}
		})
	}
}
