package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TerryMinn/task-cli/internal/config"
	"github.com/TerryMinn/task-cli/internal/models"
	"github.com/TerryMinn/task-cli/internal/utils"
)

type Type string

const (
	Update Type = "update"
	Delete Type = "delete"
	Prefix Type = "mark-"
)

func AddNewTodo(tasks *[]models.Todo, input utils.Input) {
	newTask := models.Todo{
		Id:          len(*tasks) + 1,
		Description: input.Value,
		Status:      config.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	*tasks = append(*tasks, newTask)
}

func MutateTodo(tasks *[]models.Todo, status Type, input utils.Input) int {
	var cbTarget int
	utils.IndexFinder(*tasks, func(target int) {
		cbTarget = target
		if status == Update {
			(*tasks)[target].Description = input.Value
		} else if strings.HasPrefix(string(status), string(Prefix)) {
			(*tasks)[target].Status = config.Operation(strings.TrimPrefix(string(status), string(Prefix)))
		} else {
			*tasks = append((*tasks)[:target], (*tasks)[target+1:]...)
		}
	})
	return cbTarget
}

func printOut(task models.Todo) {
	fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, utils.StatusChecker(task.Status))
}

func GetTodoList(tasks []models.Todo, app *config.Application) {
	fmt.Printf("%-5s %-20s %-10s\n", "ID", "Description", "Status")
	fmt.Println("-------------------------------------------")
	for _, task := range tasks {
		if len(os.Args) > 2 {
			if task.Status == config.Operation(strings.TrimPrefix(string(app.ListOperation), string(Prefix))) {
				printOut(task)
			}
		} else {
			printOut(task)
		}
	}
}
