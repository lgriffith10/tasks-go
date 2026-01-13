package services

import (
	"errors"
	"strings"
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
	if strings.Trim(name, " ") == "" {
		return errors.New("task name cannot be empty")
	}

	if ts.list.Exists(name) {
		return errors.New("task with the same name already exists")
	}

	task := domain.NewTask(name)

	err := ts.repo.Save(task)

	if err != nil {
		return err
	}

	list, err := ts.repo.Load()

	if err != nil {
		return err
	}

	ts.list = list

	return nil
}

func (ts *TaskService) GetTasks() *domain.TaskList {
	return ts.list
}

func (ts *TaskService) CompleteTask(name string) error {
	task, err := ts.list.GetByName(name)

	if err != nil {
		return err
	}

	task.MarkAsCompleted()

	err = ts.repo.Save(task)

	if err != nil {
		return err
	}

	_, err := ts.repo.Load()

	if err != nil {
		return err
	}
}
