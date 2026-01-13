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
	err := c.service.AddTask("teststs")

	if err != nil {
		return fmt.Errorf("Could not add error: %v", err.Error())
	}

	return nil
}
