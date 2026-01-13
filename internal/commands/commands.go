package commands

import (
	"fmt"
	"tasks_go/internal/services"
)

type Command struct {
	service *services.TaskService
}

func NewCommand(service *services.TaskService) *Command {
	return &Command{
		service: service,
	}
}

func (c *Command) AddTask() error {
	var name string

	fmt.Println("Enter task name:")
	fmt.Scanln(&name)

	err := c.service.AddTask(name)
	if err != nil {
		return fmt.Errorf("Could not add error: %v", err.Error())
	}

	return nil
}

func (c *Command) List() error {
	tasks := c.service.GetTasks()

	for _, task := range tasks.Tasks {
		fmt.Println(task)
	}

	return nil
}

// func (c *Command) CompleteTask() error {
// 	tasks := c.service.GetTasks()

// 	var name string
// 	fmt.Println("Enter task name:")
// 	fmt.Scanln(&name)

// 	// index := slices.IndexFunc(tasks, func(t Task) bool {
// 	// 	return name == t.Name
// 	// })

// 	// if index == -1 {
// 	// 	return fmt.Errorf("Task not found")
// 	// }
// }
