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
		service.GetTodoList(tasks, app)
		break

	case config.MarkInProgress:
		target := service.MutateTodo(&tasks, service.Type(config.MarkInProgress), *input)
		utils.ApplyChanges(tasks, "task.json")
		color.Green("Start task with ID : %d", target)
		break

	case config.MarkDone:
		target := service.MutateTodo(&tasks, service.Type(config.MarkDone), *input)
		utils.ApplyChanges(tasks, "task.json")
		color.Green("Done task complete with ID : %d", target)
		break

	case config.MarkTodo:
		target := service.MutateTodo(&tasks, service.Type(config.MarkTodo), *input)
		utils.ApplyChanges(tasks, "task.json")
		color.Green("Done task complete with ID : %d", target)
		break

	default:
		fmt.Printf("task management command not found\n")
	}
}
