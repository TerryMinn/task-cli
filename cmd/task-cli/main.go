package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TerryMinn/task-cli/internal/config"
	"github.com/TerryMinn/task-cli/internal/models"
	"github.com/TerryMinn/task-cli/internal/service"
	"github.com/TerryMinn/task-cli/internal/utils"
	"github.com/fatih/color"
)

func main() {

	input := utils.GetUserInput(func() {
		service.MainHelper()
	})

	app := &config.Application{
		Operation:     config.Operation(input.Op),
		ListOperation: config.Operation(input.Value),
		FilePath:      "task.json",
		Version:       "1.0.0",
	}

	// 1. This phrase will create json file when file is not exist
	file, fileError := os.OpenFile(app.FilePath, os.O_RDWR|os.O_CREATE, 0666)
	if fileError != nil {
		log.Fatal(fileError)
	}

	defer file.Close()

	// 2. Read the json data
	raw, rawError := io.ReadAll(file)
	if rawError != nil {
		log.Fatal(rawError)
	}

	// 3. parse into the task struct
	var tasks []models.Todo

	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &tasks); err != nil {
			log.Fatal(err)
		}
	}

	switch app.Operation {
	case config.Add:
		// adding new todo with slice pointer and apply change function
		service.AddNewTodo(&tasks, *input)
		utils.ApplyChanges(tasks, app.FilePath)
		color.Green("Add new task complete with ID : %d", len(tasks))
		break

	case config.Update:
		// update todo with slice pointer and apply change function
		target := service.MutateTodo(&tasks, service.Update, *input)
		utils.ApplyChanges(tasks, "task.json")
		color.Green("Update task complete with ID : %d", target)
		break

	case config.Delete:
		//delete todo with slice pointer and apply change function
		target := service.MutateTodo(&tasks, service.Delete, *input)
		utils.ApplyChanges(tasks, "task.json")
		color.Green("Delete task complete with ID : %d", target)
		break

	case config.List:
		fmt.Printf("%-5s %-20s %-10s\n", "ID", "Description", "Status")
		fmt.Println("-------------------------------------------")

		if len(os.Args) > 2 {
			goto OPTION
		}

		for _, task := range tasks {
			fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, utils.StatusChecker(task.Status))
		}

	OPTION:

		switch app.ListOperation {
		case config.Todo:
			for _, task := range tasks {
				if task.Status == models.TASK {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, utils.StatusChecker(task.Status))
				}
			}
			break
		case config.InProgress:
			for _, task := range tasks {
				if task.Status == models.IN_PROGRESS {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, utils.StatusChecker(task.Status))
				}
			}
			break
		case config.Done:
			for _, task := range tasks {
				if task.Status == models.DONE {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, utils.StatusChecker(task.Status))
				}
			}
			break
		default:
			fmt.Println("Error: invalid argument. Usage: task-task-cli <file>")
			os.Exit(1)
		}
		break

	case "mark-in-progress":
		utils.IndexFinderOld(func(target int) {
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Status = 1
				}
			}
			utils.ApplyChanges(tasks, "task.json")
			color.Green("Start task complete with ID : %d", target)
		})
		break

	case "mark-done":
		utils.IndexFinderOld(func(target int) {
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Status = 2
				}
			}
			utils.ApplyChanges(tasks, "task.json")
			color.Green("Done task complete with ID : %d", target)
		})
		break

	default:
		fmt.Printf("task management command not found\n")
	}
}
