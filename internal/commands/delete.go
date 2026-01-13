package commands

import (
	"fmt"
	"slices"
	"tasks_go/internal/json"
	"tasks_go/internal/tasks"
)

func Delete() {
	list := json.GetList()

	var name string

	fmt.Println("Enter the name of the task to delete")
	fmt.Scanln(&name)

	if name == "" {
		fmt.Println("Task name cannot be empty")
		Delete()
	}

	index := slices.IndexFunc(list.Tasks, func(t tasks.Task) bool {
		return t.Name == name
	})

	if index == -1 {
		fmt.Println("Task not found")
		Delete()
	}

	list.Tasks = append(list.Tasks[:index], list.Tasks[index+1:]...)
	json.SaveList()
}
