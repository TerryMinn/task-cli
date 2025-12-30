package utils

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/TerryMinn/task-cli/internal/models"
)

func StatusChecker(status models.Status) string {
	if status == models.TASK {
		return "Todo"
	} else if status == models.IN_PROGRESS {
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
		for index, t := range tasks {
			if t.Id == num {
				target = index
			}
		}
		cb(target)
	}
}
