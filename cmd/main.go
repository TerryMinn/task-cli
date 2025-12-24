package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	file, fileError := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)
	if fileError != nil {
		log.Fatal(fileError)
	}

	defer file.Close()

	raw, dataError := os.ReadFile("task.json")
	if dataError != nil {
		log.Fatal(dataError)
	}

	var tasks []Todo

	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &tasks); err != nil {
			log.Fatal(err)
		}
	}

	value := os.Args[2]

	_, conErr := strconv.Atoi(value)

	if conErr != nil {
		log.Fatal(conErr)
	}

	switch os.Args[1] {
	case "add":
		newTask := Todo{
			Id:          len(tasks) + 1,
			Description: value,
			Status:      todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, newTask)
		ApplyChanges(tasks, file)
		color.Green("Add new task complete with ID : %d", len(tasks))
		break
	case "update":
		break
	case "delete":

		break
	case "list":
		break
	case "mark-in-progress":
		break
	case "mark-done":
		break
	default:
	}
}
