package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"tasks_go/internal/tasks"
)

var list tasks.TaskList

func GetList() *tasks.TaskList {
	return &list
}

func CheckIfJsonExists() {
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

func ReadJson() {
	file, err := os.ReadFile("tasks.json")

	if err != nil {
		log.Fatal("Could not read file")
	}

	err = json.Unmarshal(file, &list)

	if err != nil {
		if err.Error() != "unexpected end of JSON input" {
			log.Fatal(err.Error())
		}
	}
}

func SaveList() {
	jsonData, err := json.MarshalIndent(&list, "", "  ")

	if err != nil {
		log.Fatal("Could not format data to JSON")
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)

	if err != nil {
		log.Fatal("Could not write to file")
	}
}
