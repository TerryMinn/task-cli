package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TerryMinn/task-cli/cmd"
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

	var tasks []cmd.Todo

	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &tasks); err != nil {
			log.Fatal(err)
		}
	}

	value := os.Args[2]

	switch os.Args[1] {
	case "add":
		newTask := cmd.Todo{
			Id:          len(tasks) + 1,
			Description: value,
			Status:      0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, newTask)
		cmd.ApplyChanges(tasks, "task.json")
		color.Green("Add new task complete with ID : %d", len(tasks))
		break
	case "update":
		cmd.IndexFinder(value, func(target int) {
			updateValue := os.Args[3]
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Description = updateValue
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Update task complete with ID : %d", target)
		})
		break
	case "delete":
		cmd.IndexFinder(value, func(target int) {
			for i, task := range tasks {

				if task.Id == target {
					tasks = append(tasks[:i], tasks[i+1:]...)
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Delete task complete with ID : %d", target)
		})
		break
	case "list":
		fmt.Printf("%-5s %-20s %-10s\n", "ID", "Description", "Status")
		fmt.Println("-------------------------------------------")
		break
	case "mark-in-progress":
		break
	case "mark-done":
		break
	default:
		fmt.Printf("task management command not found\n")
	}
}
