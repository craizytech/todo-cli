package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintMenu() {
	fmt.Println(`
	== ToDo CLI ==
	
	1. Add a new Task
	2. Mark Task as Done
	3. Delete Task
	4. List All Tasks
	5. Exit
	`)
}

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
