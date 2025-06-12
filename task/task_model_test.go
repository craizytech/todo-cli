package task

import "testing"

// Testing the MarkDone method of the Task struct
func TestTask_MarkAsDone(t *testing.T) {
	// creating a task
	task := Task{Name: "Test Task", Done: false}

	// calling MarkAsDone on the task
	task.MarkAsDone()

	// verify that Done on the task is now true
	if !task.Done {
		t.Errorf("MarkAsDone failed: expected Done to be true, got false")
	}

	// verifying that task name remained unchanged
	if task.Name != "Test Task" {
		t.Errorf("MarkAsDone modified Name: expected %q, got %q", "Test Task", task.Name)
	}
}
