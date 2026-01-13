package repository

import tasks "tasks_go/internal/domain"

type TaskRepository interface {
	Save(task *tasks.Task) error
	// FindByName(name string) (*tasks.Task, error)
	// Delete(name string) error
	Load() (*tasks.TaskList, error)
}
