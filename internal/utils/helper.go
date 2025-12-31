package utils

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/TerryMinn/task-cli/internal/config"
	"github.com/TerryMinn/task-cli/internal/models"
)

func StatusChecker(status config.Operation) string {
	if status == config.Todo {
		return "Todo"
	} else if status == config.InProgress {
		return "In Progress"
	} else {
		return "Done"
	}
}

func ApplyChanges(newTodo []models.Todo, filename string) {
	jsonData, err := json.MarshalIndent(newTodo, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	// Clean up (truncate + overwrite)
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		log.Fatal(err)
	}
}

func IndexFinderOld(cb models.Callback) {
	value := os.Args[2]
	if num, conErr := strconv.Atoi(value); conErr != nil {
		log.Fatal(conErr)
	} else {
		cb(num)
	}
}

func IndexFinder(tasks []models.Todo, cb models.Callback) {
	value := os.Args[2]
	if num, conErr := strconv.Atoi(value); conErr != nil {
		log.Fatal(conErr)
	} else {
		var target int
		for _, t := range tasks {
			if t.Id == num {
				target = t.Id
			}
		}
		cb(target)
	}
}
