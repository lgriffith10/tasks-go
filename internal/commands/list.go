package commands

import (
	"fmt"
	"tasks_go/internal/json"
)

func List() {
	list := json.GetList()

	if len(list.Tasks) == 0 {
		fmt.Println("No tasks found, add a task to get started!")
	}

	for _, task := range list.Tasks {
		fmt.Println(task)
	}
}
