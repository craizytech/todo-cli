package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/craizytech/todo-cli/task"
)

// TestInMemoryStore_Add tests the Add method of InMemoryStore
func TestInMemoryStore_Add(t *testing.T) {
	tests := []struct {
		name       string
		task       task.Task
		setupTasks []task.Task
		wantErr    error
		wantTasks  []task.Task
	}{
		{
			name:       "Add valid task",
			task:       task.Task{Name: "Task 1"},
			setupTasks: []task.Task{},
			wantErr:    nil,
			wantTasks:  []task.Task{{Name: "Task 1", Done: false}},
		},
		{
			name:       "Add empty task name",
			task:       task.Task{Name: ""},
			setupTasks: []task.Task{},
			wantErr:    ErrTaskEmptyName,
			wantTasks:  []task.Task{},
		},
		{
			name:       "Add duplicate task",
			task:       task.Task{Name: "Task 1"},
			setupTasks: []task.Task{{Name: "Task 1", Done: false}},
			wantErr:    ErrTaskExists,
			wantTasks:  []task.Task{{Name: "Task 1", Done: false}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initializing store with setup tasks
			store := &InMemoryStore{tasks: tt.setupTasks}

			// call add
			err := store.Add(tt.task)

			// check error
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Add() error = %v, want %v", err, tt.wantErr)
			}

			//check tasks
			if !reflect.DeepEqual(store.tasks, tt.wantTasks) {
				t.Errorf("Add() tasks = %v, want %v", store.tasks, tt.wantTasks)
			}
		})
	}
}

// TestInMemoryStore_Delete tests the Delete method of InMemoryStore
func TestInMemoryStore_Delete(t *testing.T) {
	tests := []struct {
		name       string
		deleteName string
		setupTasks []task.Task
		wantErr    error
		wantTasks  []task.Task
	}{
		{
			name:       "Delete existing task",
			deleteName: "Task 1",
			setupTasks: []task.Task{{Name: "Task 1"}, {Name: "Task 2"}},
			wantErr:    nil,
			wantTasks:  []task.Task{{Name: "Task 2"}},
		},
		{
			name:       "Delete non-existing task",
			deleteName: "Task 3",
			setupTasks: []task.Task{{Name: "Task 1"}},
			wantErr:    ErrTaskNotFound,
			wantTasks:  []task.Task{{Name: "Task 1"}},
		},
		{
			name:       "Delete from empty store",
			deleteName: "Task 1",
			setupTasks: []task.Task{},
			wantErr:    ErrTaskNotFound,
			wantTasks:  []task.Task{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &InMemoryStore{tasks: tt.setupTasks}
			err := store.Delete(tt.deleteName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(store.tasks, tt.wantTasks) {
				t.Errorf("Delete() tasks = %v, want %v", store.tasks, tt.wantTasks)
			}
		})
	}

}

// TestInMemoryStore_MarkAsDone tests the MarkAsDone method of InMemoryStore
func TestInMemoryStore_MarkAsDone(t *testing.T) {
	tests := []struct {
		name       string
		markName   string
		setupTasks []task.Task
		wantErr    error
		wantTasks  []task.Task
	}{
		{
			name:       "Mark existing task as done",
			markName:   "Task 1",
			setupTasks: []task.Task{{Name: "Task 1", Done: false}},
			wantErr:    nil,
			wantTasks:  []task.Task{{Name: "Task 1", Done: true}},
		},
		{
			name:       "Mark non-existing task",
			markName:   "Task 2",
			setupTasks: []task.Task{{Name: "Task 1", Done: false}},
			wantErr:    ErrTaskNotFound,
			wantTasks:  []task.Task{{Name: "Task 1", Done: false}},
		},
		{
			name:       "Mark task in empty store",
			markName:   "Task 1",
			setupTasks: []task.Task{},
			wantErr:    ErrTaskNotFound,
			wantTasks:  []task.Task{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &InMemoryStore{tasks: tt.setupTasks}
			err := store.MarkAsDone(tt.markName)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("MarkAsDone() error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(store.tasks, tt.wantTasks) {
				t.Errorf("MarkAsDone() tasks = %v, want %v", store.tasks, tt.wantTasks)
			}
		})
	}
}

// TestInMemoryStore_List tests the List method of InMemoryStore
func TestInMemoryStore_List(t *testing.T) {
	tests := []struct {
		name       string
		setupTasks []task.Task
		wantTasks  []task.Task
	}{
		{
			name:       "List multiple tasks",
			setupTasks: []task.Task{{Name: "Task 1"}, {Name: "Task 2", Done: true}},
			wantTasks:  []task.Task{{Name: "Task 1"}, {Name: "Task 2", Done: true}},
		},
		{
			name:       "List empty store",
			setupTasks: []task.Task{},
			wantTasks:  []task.Task{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &InMemoryStore{tasks: tt.setupTasks}
			got := store.List()
			if !reflect.DeepEqual(got, tt.wantTasks) {
				t.Errorf("List() = %v, want %v", got, tt.wantTasks)
			}
		})
	}
}

// TestInMemoryStore_Exists tests the Exists method of InMemoryStore
func TestInMemoryStore_Exists(t *testing.T) {
	tests := []struct {
		name       string
		checkName  string
		setupTasks []task.Task
		want       bool
	}{
		{
			name:       "Task exists",
			checkName:  "Task 1",
			setupTasks: []task.Task{{Name: "Task 1"}, {Name: "Task 2"}},
			want:       true,
		},
		{
			name:       "Task does not exist",
			checkName:  "Task 3",
			setupTasks: []task.Task{{Name: "Task 1"}},
			want:       false,
		},
		{
			name:       "Empty store",
			checkName:  "Task 1",
			setupTasks: []task.Task{},
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &InMemoryStore{tasks: tt.setupTasks}
			got := store.Exists(tt.checkName)
			if got != tt.want {
				t.Errorf("Exists(%q) = %v, want %v", tt.checkName, got, tt.want)
			}
		})
	}
}
