package tasks

import "github.com/google/uuid"

type Task struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Done bool      `json:"done"`
}

func NewTask(name string) *Task {
	return &Task{
		ID:   uuid.New(),
		Name: name,
		Done: false,
	}
}

func (t *Task) ChangeDoneStatus() {
	t.Done = !t.Done
}
