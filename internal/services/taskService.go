package services

import (
	"tasks_go/internal/domain"
	"tasks_go/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
	list *domain.TaskList
}

func NewTaskService(repo repository.TaskRepository) (*TaskService, error) {
	list, err := repo.Load()

	if err != nil {
		return nil, err
	}

	return &TaskService{
		repo: repo,
		list: list,
	}, nil
}

func (ts *TaskService) AddTask(name string) error {
	task := domain.NewTask(name)

	err := ts.repo.Save(task)

	if err != nil {
		return err
	}

	return nil
}
