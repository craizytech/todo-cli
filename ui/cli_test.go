package ui

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestPrintMenu tests the PrintMenu function by capturing its output
func TestPrintMenu(t *testing.T) {
	// Redirect stdout ton capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// call PrintMenu
	PrintMenu()

	//Restore stdout and read captured output
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	// expected output
	expected := `
	== ToDo CLI ==
	
	1. Add a new Task
	2. Mark Task as Done
	3. Delete Task
	4. List All Tasks
	5. Exit
	`

	// normalize newlines for cross-platform compatibility
	got := strings.TrimSpace(strings.ReplaceAll(buf.String(), "\r\n", "\n"))
	expected = strings.TrimSpace(strings.ReplaceAll(expected, "\r\n", "\n"))

	if got != expected {
		t.Errorf("PrintMenu() output =\n%q\nwant\n%q", got, expected)
	}

}

// TestGetInput tests the GetInputfunction by simulating user input.
func TestGetInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal input\n",
			input:    "Test input\n",
			expected: "Test input",
		},
		{
			name:     "Input with spaces",
			input:    "   Test input   \n",
			expected: "Test input",
		},
		{
			name:     "Empty input",
			input:    "\n",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a pipe for stdin
			r, w, _ := os.Pipe()
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = r

			// write input to the pipe
			w.WriteString(tt.input)
			w.Close()

			got := GetInput("Enter something: ")
			if got != tt.expected {
				t.Errorf("GetInput() %q, want %q", got, tt.expected)
			}

		})
	}
}
