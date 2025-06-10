package store

import (
	"errors"
	"strings"

	"github.com/craizytech/todo-cli/task"
)

type InMemoryStore struct {
	tasks []task.Task
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{tasks: []task.Task{}}
}

func (s *InMemoryStore) Add(t task.Task) error {
	if len(strings.TrimSpace(t.Name)) == 0 {
		return errors.New("task name cannot be empty")
	}

	if s.Exists(t.Name) {
		return errors.New("task already exists")
	}

	s.tasks = append(s.tasks, t)
	return nil
}

func (s *InMemoryStore) Delete(name string) error {
	for i, t := range s.tasks {
		if t.Name == name {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *InMemoryStore) MarkAsDone(name string) error {
	for i := range s.tasks {
		if s.tasks[i].Name == name {
			s.tasks[i].MarkAsDone()
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *InMemoryStore) List() []task.Task {
	return s.tasks
}

func (s *InMemoryStore) Exists(name string) bool {
	for _, task := range s.tasks {
		if task.Name == name {
			return true
		}
	}
	return false
}
