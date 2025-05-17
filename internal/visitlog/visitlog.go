// Package visitlog provides a simple in-memory database for tracking patient visits to doctors.
// It supports adding visits, retrieving full visit history, and fetching the last visit for a patient.
// The package is designed for CLI interaction with commands like "save", "gethistory", and "getlastvisit".
package visitlog

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// Patient represents a patient's full name as a string.
type Patient string

// Doctor represents a doctor's specialization as a string.
type Doctor string

// Visit contains information about a patient's visit to a doctor,
// including the doctor's specialization and the date of the visit.
type Visit struct {
	DocSpec Doctor
	Date    time.Time
}

// VisitHistory stores a map of patients to their visit records.
// It is the core in-memory database for the package.
type VisitHistory struct {
	visits map[Patient][]Visit
}

// UserNotFoundError is returned when a requested patient has no visit records.
type UserNotFoundError struct{}

// Error returns the error message for UserNotFoundError.
func (e *UserNotFoundError) Error() string {
	return "user not found"
}

// Predefined errors for input validation and system operations.
var (
	ErrEmptyInput      = errors.New("empty input")          // Returned when input is empty.
	ErrInputError      = errors.New("input error")          // Generic input error (e.g., read failure).
	ErrValidInputError = errors.New("incorrect value")      // Returned for invalid commands/inputs.
	ErrIncorrectDate   = errors.New("incorrect date value") // Returned for malformed dates.
	ErrExit            = errors.New("exit")                 // Signals an exit request.
)

// ErrUserNotFound is a singleton instance of UserNotFoundError.
var ErrUserNotFound = &UserNotFoundError{}

// commands - all of "good" commands
var commands = map[string]bool{
	"save":         true,
	"gethistory":   true,
	"getlastvisit": true,
	"exit":         true,
}

// parseString reads a line of input from the reader after displaying a prompt.
// Returns the trimmed input or an error if input is empty or unreadable.
func parseString(s string, reader *bufio.Reader, writer io.Writer) (string, error) {
	fmt.Fprint(writer, s)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Join(ErrInputError, err)
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", ErrEmptyInput
	}
	return input, nil
}

// parseCommand reads and validates a user command (e.g., "save", "gethistory").
// Returns the lowercase command or an error if invalid.
func parseCommand(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Join(ErrInputError, err)
	}
	input = strings.TrimSpace(strings.ToLower(input))
	if _, ok := commands[input]; !ok {
		return "", ErrValidInputError
	}
	return input, nil
}

// parsePatient prompts for and reads a patient's name.
// Wraps parseString with a standardized prompt.
func parsePatient(reader *bufio.Reader, writer io.Writer) (string, error) {
	return parseString("Input patient full name: ", reader, writer)
}

// parseDoctor prompts for and reads a doctor's specialization.
// Wraps parseString with a standardized prompt.
func parseDoctor(reader *bufio.Reader, writer io.Writer) (string, error) {
	return parseString("Input doctor's specialization: ", reader, writer)
}

// parseDate prompts for and parses a date in "YYYY-MM-DD" format.
// Returns a time.Time or an error if parsing fails.
func parseDate(reader *bufio.Reader, writer io.Writer) (time.Time, error) {
	input, err := parseString("Input date in format\"YYYY-MM-DD\": ", reader, writer)
	if err != nil {
		return time.Time{}, errors.Join(ErrInputError, err)
	}
	visitDate, err := time.Parse("2006-01-02", input)
	if err != nil {
		return time.Time{}, errors.Join(ErrIncorrectDate, err)
	}
	return visitDate, nil
}

// addVisit records a new visit for a patient.
// Requires patient name, doctor specialization, and date.
// Initializes the visits map if nil.
func (v *VisitHistory) addVisit(reader *bufio.Reader, writer io.Writer) error {
	if v.visits == nil {
		v.visits = make(map[Patient][]Visit)
	}
	p, err := parsePatient(reader, writer)
	if err != nil {
		return err
	}
	d, err := parseDoctor(reader, writer)
	if err != nil {
		return err
	}
	t, err := parseDate(reader, writer)
	if err != nil {
		return err
	}
	v.visits[Patient(p)] = append(v.visits[Patient(p)], Visit{Doctor(d), t})
	return nil
}

// getAllVisits retrieves all visits for a given patient.
// Returns ErrUserNotFound if the patient has no records.
func (v *VisitHistory) getAllVisits(reader *bufio.Reader, writer io.Writer) ([]Visit, error) {
	if v.visits == nil {
		return nil, ErrUserNotFound
	}
	p, err := parsePatient(reader, writer)
	if err != nil {
		return nil, err
	}
	visits, exists := v.visits[Patient(p)]
	if !exists {
		return nil, ErrUserNotFound
	}
	return visits, nil
}

// getLastVisit fetches the most recent visit for a patient to a specific doctor.
// Returns ErrUserNotFound if no matching visits exist.
func (v *VisitHistory) getLastVisit(reader *bufio.Reader, writer io.Writer) (Visit, error) {
	if v.visits == nil {
		return Visit{}, ErrUserNotFound
	}
	visits, err := v.getAllVisits(reader, writer)
	if err != nil {
		return Visit{}, err
	}
	var lastVisit Visit
	found := false
	d, err := parseDoctor(reader, writer)
	if err != nil {
		return Visit{}, err
	}
	for _, visit := range visits {
		if visit.DocSpec == Doctor(d) {
			if !found || visit.Date.After(lastVisit.Date) {
				lastVisit = visit
				found = true
			}
		}
	}
	if !found {
		return Visit{}, ErrUserNotFound
	}
	return lastVisit, nil
}

// printVisit formats and writes a single visit to the writer.
func printVisit(v Visit, writer io.Writer) {
	fmt.Fprintln(writer, v.DocSpec, " ", v.Date.Format("2006-01-02"))
}

// printAllVisits prints all visits in a slice using printVisit.
func printAllVisits(v []Visit, writer io.Writer) {
	for _, visit := range v {
		printVisit(visit, writer)
	}
}

// handleCommand displays the command menu and reads user input.
// Repeats on invalid input until a valid command is received.
func (v *VisitHistory) handleCommand(reader *bufio.Reader, writer io.Writer) (string, error) {
	fmt.Fprintln(writer, "--------------")
	fmt.Fprintln(writer, "Input command:\n- Save\n- GetHistory\n- GetLastVisit\n- Exit:")
	var s string
	for {
		sParsed, err := parseCommand(reader)
		if err != nil {
			if errors.Is(err, ErrValidInputError) {
				fmt.Fprint(writer, "Invalid input. Please give me correct command: ")
				continue
			}
			return "", err
		}
		s = sParsed
		break
	}
	return s, nil
}

// handleSave orchestrates the "save" command flow.
// Handles errors like incorrect dates or empty inputs.
func (v *VisitHistory) handleSave(reader *bufio.Reader, writer io.Writer) error {
	err := v.addVisit(reader, writer)
	if err != nil {
		if errors.Is(err, ErrIncorrectDate) {
			fmt.Fprintln(writer, "Incorrect input date value")
		}
		if errors.Is(err, ErrEmptyInput) {
			fmt.Fprintln(writer, "Input empty string")
		}
		return err
	}
	return nil
}

// handleAllVisits orchestrates the "gethistory" command flow.
// Prints all visits or a "not found" message.
func (v *VisitHistory) handleAllVisits(reader *bufio.Reader, writer io.Writer) error {
	visits, err := v.getAllVisits(reader, writer)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Fprintln(writer, err.Error())
		}
		return err
	}
	printAllVisits(visits, writer)
	return nil
}

// handleLastVisit orchestrates the "getlastvisit" command flow.
// Prints the last visit or a "not found" message.
func (v *VisitHistory) handleLastVisit(reader *bufio.Reader, writer io.Writer) error {
	visit, err := v.getLastVisit(reader, writer)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Fprintln(writer, err.Error())
		}
		return err
	}
	printVisit(visit, writer)
	return nil
}

// ActionDB is the main entry point for the package.
// Processes user commands (save/gethistory/getlastvisit/exit) and delegates to handlers.
// Returns ErrExit to signal a graceful termination.
func (v *VisitHistory) ActionDB(reader *bufio.Reader, writer io.Writer) error {
	s, err := v.handleCommand(reader, writer)
	if err != nil {
		return err
	}
	switch s {
	case "save":
		err := v.handleSave(reader, writer)
		if err != nil {
			return err
		}
	case "gethistory":
		err := v.handleAllVisits(reader, writer)
		if err != nil {
			return err
		}
	case "getlastvisit":
		err := v.handleLastVisit(reader, writer)
		if err != nil {
			return err
		}
	case "exit":
		return ErrExit
	}
	return nil
}
