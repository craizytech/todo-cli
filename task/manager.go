package task

type Store interface {
	Add(task Task) error
	Delete(name string) error
	MarkDone(name string) error
	List() []Task
	Exists(name string) bool
}
