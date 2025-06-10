package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/craizytech/todo-cli/store"
	"github.com/craizytech/todo-cli/task"
	"github.com/craizytech/todo-cli/ui"
)

// entry point of our application

func main() {
	store := store.NewInMemoryStore()

	for {
		ui.PrintMenu()
		choiceStr := ui.GetInput("Enter choice: ")

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Enter a number")
			continue
		}

		switch choice {
		case 1:
			name := ui.GetInput("Task name: ")
			err := store.Add(task.Task{Name: name})
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Task added successfully")
			}
		case 2:
			name := ui.GetInput("Task name to mark as done: ")
			err := store.MarkAsDone(name)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Task marked as done")
			}
		case 3:
			name := ui.GetInput("Task name to delete: ")
			err := store.Delete(name)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Task deleted")
			}
		case 4:
			tasks := store.List()
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				continue
			}
			fmt.Println("Tasks:")
			for _, task := range tasks {
				status := " "
				if task.Done {
					status = "✓"
				}
				fmt.Printf("- [%s] %s\n", status, task.Name)
			}
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}

	}
}
