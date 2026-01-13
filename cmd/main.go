package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"tasks_go/internal/commands"
	"tasks_go/internal/repository"
	"tasks_go/internal/services"
)

var commandFlag = flag.String("command", "", "Command to execute")

func main() {
	flag.Parse()

	checkIfJsonExists()

	repository := repository.NewJsonTaskRepository("tasks.json")
	taskService, err := services.NewTaskService(repository)

	if err != nil {
		log.Fatalf("Could not create task service: %v", err.Error())
	}

	commands := commands.NewCommand(taskService)

	switch *commandFlag {
	case "list":
		commands.List()
	case "add":
		err = commands.AddTask()
	case "done":
		// commands.MarkDone()
	case "delete":
		// commands.Delete()
	default:
		fmt.Println("No valid command provided")
	}

	if err != nil {
		log.Fatalf("Error executing command %v: [%v]", *commandFlag, err.Error())
	}
}

func checkIfJsonExists() {
	projectDirectory, err := filepath.Abs(".")

	if err != nil {
		log.Fatal("Could not get project directory")
	}

	_, err = os.Stat(path.Join(projectDirectory, "tasks.json"))

	if os.IsNotExist(err) {
		createFile()
	} else if err != nil {
		log.Fatal(err.Error())
	}
}

func createFile() {
	_, err := os.Create("tasks.json")

	if err != nil {
		log.Fatal("Could not create file")
	}

	fmt.Println("File created successfully")
}
