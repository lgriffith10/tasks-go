package repository

import (
	"encoding/json"
	"os"
	tasks "tasks_go/internal/domain"
)

type JsonTaskRepository struct {
	filepath string
}

func NewJsonTaskRepository(filepath string) *JsonTaskRepository {
	return &JsonTaskRepository{filepath: filepath}
}

func (jtr *JsonTaskRepository) Load() (*tasks.TaskList, error) {
	content, err := os.ReadFile(jtr.filepath)

	if err != nil {
		return nil, err
	}

	list := tasks.NewTaskList()
	err = json.Unmarshal(content, list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (jtr *JsonTaskRepository) Save(task *tasks.Task) error {
	list, err := jtr.Load()

	if err != nil {
		return err
	}

	list.Tasks = append(list.Tasks, *task)

	data, err := json.MarshalIndent(list, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(jtr.filepath, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
