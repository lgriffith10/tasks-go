package main

import (
	"flag"
	"fmt"
	"tasks_go/internal/commands"
	"tasks_go/internal/json"
)

var commandFlag = flag.String("command", "", "Command to execute")

func main() {
	flag.Parse()

	json.CheckIfJsonExists()
	json.ReadJson()

	switch *commandFlag {
	case "list":
		commands.List()
	case "add":
		commands.Add()
	case "done":
		commands.MarkDone()
	case "delete":
		commands.Delete()
	default:
		fmt.Println("No valid command provided")
	}
}
