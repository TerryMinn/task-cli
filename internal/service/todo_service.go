package service

import (
	"time"

	"github.com/TerryMinn/task-cli/internal/models"
	"github.com/TerryMinn/task-cli/internal/utils"
)

type Type string

const (
	Update Type = "update"
	Delete Type = "delete"
)

func AddNewTodo(tasks *[]models.Todo, input utils.Input) {
	newTask := models.Todo{
		Id:          len(*tasks) + 1,
		Description: input.Value,
		Status:      models.TASK,
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
		} else {
			*tasks = append((*tasks)[:target], (*tasks)[target+1:]...)
		}
	})
	return cbTarget
}
