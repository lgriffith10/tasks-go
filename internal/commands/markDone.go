package commands

import (
	"fmt"
	"slices"
	"tasks_go/internal/json"
	"tasks_go/internal/tasks"
)

func MarkDone() {
	var name string

	fmt.Println("Enter the name of the task to mark as done")
	fmt.Scanln(&name)

	if name == "" {
		fmt.Println("Task name cannot be empty")
		MarkDone()
	}

	list := json.GetList()

	index := slices.IndexFunc(list.Tasks, func(t tasks.Task) bool {
		return t.Name == name
	})

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	list.Tasks[index].ChangeDoneStatus()
	json.SaveList()
}
