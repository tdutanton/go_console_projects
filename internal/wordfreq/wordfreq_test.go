package wordfreq

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_wordsCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"one word", "hello", 1},
		{"two words", "hello go", 2},
		{"empty", "", 0},
		{"empty space empty", " ", 0},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			if got := wordsCount(d.input); got != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, got)
			}
		})
	}
}

func TestParseK(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     int
		expErr       bool
		expReaderErr bool
		errContains  string
	}{
		{"two", "2\n", 2, false, false, ""},
		{"float", "2.5\n", 0, true, false, "Invalid input"},
		{"not number", "gogogo\n", 0, true, false, "Invalid input"},
		{"not just number", "4go\n", 0, true, false, "Invalid input"},
		{"invalid reader", "444", 0, false, true, "input error"},
		{"negative K", "-5\n", 0, true, false, "Invalid input"},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(d.input))
			var output bytes.Buffer
			got, err := ParseK("Enter K value (words limit): ", reader, &output)
			if !d.expErr && got != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, got)
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

func TestGetWordsMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
		expEmpty bool
	}{
		{"three two one", "aa bb cc aa bb aa aa", map[string]int{"aa": 3, "bb": 2, "cc": 1}, false},
		{"empty", "", map[string]int{}, true},
		{"empty space empty", " ", map[string]int{}, true},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			got := GetWordsMap(d.input)
			if len(got) != len(d.expected) {
				t.Errorf("Got %d, expected %d", len(got), len(d.expected))
			}
			if !d.expEmpty {
				if got["aa"] != 3 && d.expected["aa"] != 3 {
					t.Errorf("Got %d, expected %d", got["aa"], d.expected["aa"])
				}
			}
			if d.expEmpty {
				if len(got) != len(d.expected) && len(got) > 0 {
					t.Errorf("Got %d, expected %d", len(got), len(d.expected))
				}
			}
		})
	}
}

func Test_getSortedWords(t *testing.T) {
	tests := []struct {
		name     string
		input    map[Word]int
		expected WordSlice
	}{
		{"three two one", map[Word]int{"aa": 1, "bb": 2, "cc": 3}, WordSlice{"cc", "bb", "aa"}},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			got := getSortedWords(d.input)
			if len(got) != len(d.expected) {
				t.Errorf("Got %d, expected %d", len(got), len(d.expected))
			}
		})
	}
}

func TestGetResultWordsSlice(t *testing.T) {
	tests := []struct {
		name       string
		inputWords string
		inputK     int
		expected   WordSlice
	}{
		{"standart case", "aa bb cc", 2, WordSlice{"aa", "bb"}},
		{"empty slice", "", 3, WordSlice{}},
		{"big K", "aa bb cc", 5, WordSlice{"aa", "bb", "cc"}},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			got := GetResultWordsSlice(d.inputWords, d.inputK)
			if len(got) != len(d.expected) {
				t.Errorf("Got %d, expected %d", len(got), len(d.expected))
			}
		})
	}
}

func TestPrintWords(t *testing.T) {
	words := WordSlice{"apple", "banana", "cherry"}

	var buf bytes.Buffer
	words.PrintWords(&buf)

	expected := "apple banana cherry\n"
	if buf.String() != expected {
		t.Errorf("Got: %q\nWant: %q", buf.String(), expected)
	}
}
