package tasks

import "slices"

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: []Task{},
	}
}

func (tl *TaskList) Exists(name string) bool {
	index := slices.IndexFunc(tl.Tasks, func(t Task) bool {
		return t.Name == name
	})

	return index != -1
}
