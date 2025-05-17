// Package main for Visit Log.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/tdutanton/go_console_projects/internal/visitlog"
)

// Main function for Visit Log.
func main() {
	reader := bufio.NewReader(os.Stdin)
	v := visitlog.VisitHistory{}
	for {
		err := v.ActionDB(reader, os.Stdout)
		if errors.Is(err, visitlog.ErrExit) {
			fmt.Println("Good bye!")
			break
		}
	}
	os.Exit(0)
}
