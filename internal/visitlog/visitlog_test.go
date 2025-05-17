package visitlog

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"
)

// Example:
//   history := &VisitHistory{}
//   err := history.ActionDB(bufio.NewReader(os.Stdin), os.Stdout)
//   if errors.Is(err, visitlog.ErrExit) {
//       os.Exit(0)
//   }

func Test_addVisit(t *testing.T) {
	input := "John Doe\nCardiology\n2024-12-01\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	history := &VisitHistory{}
	err := history.addVisit(reader, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(history.visits) != 1 {
		t.Errorf("expected 1 patient, got %d", len(history.visits))
	}
	if len(history.visits["John Doe"]) != 1 {
		t.Errorf("expected 1 visit, got %d", len(history.visits["John Doe"]))
	}
}

func Test_getAllVisits_userNotFound(t *testing.T) {
	input := "Unknown User\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	history := &VisitHistory{}
	_, err := history.getAllVisits(reader, &output)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func Test_getLastVisit(t *testing.T) {
	history := &VisitHistory{
		visits: map[Patient][]Visit{
			"Jane Doe": {
				{DocSpec: "Cardiology", Date: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)},
				{DocSpec: "Cardiology", Date: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
		},
	}
	input := "Jane Doe\nCardiology\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	visit, err := history.getLastVisit(reader, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if visit.Date.Year() != 2025 {
		t.Errorf("expected year 2025, got %d", visit.Date.Year())
	}
}

func Test_handleCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasErr   bool
	}{
		{"valid save", "Save\n", "save", false},
		{"valid gethistory", "GetHistory\n", "gethistory", false},
		{"invalid then valid", "badcommand\nsave\n", "save", false},
		{"empty input", "\n", "", true},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			h := &VisitHistory{}
			reader := bufio.NewReader(strings.NewReader(d.input))
			var output bytes.Buffer

			cmd, err := h.handleCommand(reader, &output)

			if d.hasErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !d.hasErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cmd != d.expected {
				t.Errorf("expected %s, got %s", d.expected, cmd)
			}
		})
	}
}

func Test_handleSave_success(t *testing.T) {
	input := "John Doe\nNeurology\n2024-12-05\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	h := &VisitHistory{}
	err := h.handleSave(reader, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(h.visits) != 1 {
		t.Errorf("expected 1 patient, got %d", len(h.visits))
	}
	if len(h.visits["John Doe"]) != 1 {
		t.Errorf("expected 1 visit for John Doe")
	}
}

func Test_handleSave_invalidDate(t *testing.T) {
	input := "Jane Doe\nOncology\ninvalid-date\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	h := &VisitHistory{}
	err := h.handleSave(reader, &output)
	if err != nil && !errors.Is(err, ErrIncorrectDate) {
		t.Errorf("expected ErrIncorrectDate, got %v", err)
	}
}

func Test_handleAllVisits_notFound(t *testing.T) {
	input := "Unknown Patient\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	h := &VisitHistory{}
	err := h.handleAllVisits(reader, &output)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func Test_handleAllVisits_success(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{
			"Anna Smith": {
				{DocSpec: "Pediatrics", Date: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC)},
			},
		},
	}
	input := "Anna Smith\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.handleAllVisits(reader, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(output.String(), "Pediatrics") {
		t.Errorf("expected output to contain 'Pediatrics'")
	}
}

func Test_ActionDB_save_then_gethistory(t *testing.T) {
	input := strings.Join([]string{
		"save",
		"Max Payne",
		"Surgery",
		"2025-01-01",
		"gethistory",
		"Max Payne",
		"exit",
	}, "\n") + "\n"

	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	h := &VisitHistory{}
	for {
		err := h.ActionDB(reader, &output)
		if err != nil {
			break
		}
	}
	if !strings.Contains(output.String(), "Surgery") {
		t.Errorf("expected to find Surgery in output")
	}
}

func Test_handleLastVisit_success(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{
			"Alice Wonder": {
				{DocSpec: "Cardiology", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
				{DocSpec: "Cardiology", Date: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
				{DocSpec: "Dentistry", Date: time.Date(2023, 5, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
	}
	input := "Alice Wonder\nCardiology\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.handleLastVisit(reader, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outStr := output.String()
	if !strings.Contains(outStr, "2025-01-01") || !strings.Contains(outStr, "Cardiology") {
		t.Errorf("expected latest Cardiology visit in output, got: %s", outStr)
	}
}

func Test_handleLastVisit_noDoctorMatch(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{
			"Bob Brown": {
				{DocSpec: "Neurology", Date: time.Date(2024, 6, 6, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	input := "Bob Brown\nSurgery\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.handleLastVisit(reader, &output)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got: %v", err)
	}
	if !strings.Contains(output.String(), "user not found") {
		t.Errorf("expected error message about user not found")
	}
}

func Test_handleLastVisit_userNotFound(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{},
	}

	input := "Nonexistent Patient\nAny\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.handleLastVisit(reader, &output)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got: %v", err)
	}
	if !strings.Contains(output.String(), "user not found") {
		t.Errorf("expected error message about user not found")
	}
}

func TestActionDB_getLastVisit_success(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{
			"Test Patient": {
				{DocSpec: "Therapy", Date: time.Date(2023, 3, 10, 0, 0, 0, 0, time.UTC)},
				{DocSpec: "Therapy", Date: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	input := "getlastvisit\nTest Patient\nTherapy\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.ActionDB(reader, &output)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	outStr := output.String()
	if !strings.Contains(outStr, "2024-03-10") || !strings.Contains(outStr, "Therapy") {
		t.Errorf("expected correct visit info in output, got: %s", outStr)
	}
}

func TestActionDB_getLastVisit_userNotFound(t *testing.T) {
	h := &VisitHistory{
		visits: map[Patient][]Visit{},
	}

	input := "getlastvisit\nUnknown Patient\nTherapy\n"
	reader := bufio.NewReader(strings.NewReader(input))
	var output bytes.Buffer

	err := h.ActionDB(reader, &output)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got: %v", err)
	}
	if !strings.Contains(output.String(), "user not found") {
		t.Errorf("expected 'user not found' message, got: %s", output.String())
	}
}

func TestParseString_InputError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	writer := &bytes.Buffer{}

	_, err := parseString("Prompt: ", reader, writer)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrInputError) {
		t.Errorf("expected ErrInputError, got: %v", err)
	}
}

func TestParseString_EmptyInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("\n"))
	writer := &bytes.Buffer{}

	_, err := parseString("Prompt: ", reader, writer)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrEmptyInput) {
		t.Errorf("expected ErrEmptyInput, got: %v", err)
	}
}

func TestParseDate_InputErrorFromParseString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	writer := &bytes.Buffer{}

	_, err := parseDate(reader, writer)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrInputError) {
		t.Errorf("expected ErrInputError, got: %v", err)
	}
}
