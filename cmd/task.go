package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

type Status int

const (
	TASK = iota
	IN_PROGRESS
	DONE
)

type Callback func(int)

type Todo struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ApplyChanges(newTodo []Todo, filename string) {
	jsonData, err := json.MarshalIndent(newTodo, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	// Clean up (truncate + overwrite)
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		log.Fatal(err)
	}
}

func IndexFinder(value string, cb Callback) {
	if target, conErr := strconv.Atoi(value); conErr != nil {
		log.Fatal(conErr)
	} else {
		cb(target)
	}
}
