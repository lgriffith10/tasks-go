package commands

import (
	"fmt"
	"tasks_go/internal/json"
	"tasks_go/internal/tasks"
)

func Add() {
	list := json.GetList()

	var name string

	fmt.Println("Enter the name of the task")
	fmt.Scanln(&name)

	if name == "" {
		fmt.Println("Task name cannot be empty")
		Add()
	}

	if list.Exists(name) {
		fmt.Println("Task with this name already exists")
		Add()
	}

	newTask := tasks.NewTask(name)

	list.Tasks = append(list.Tasks, *newTask)

	json.SaveList()
}
