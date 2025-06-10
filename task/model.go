package task

type Task struct {
	Name string
	Done bool
}

func (t *Task) MarkAsDone() {
	t.Done = true
}
