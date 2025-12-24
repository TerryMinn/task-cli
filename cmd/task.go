package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Status int

const (
	todo = iota
	inProgress
	done
)

type Todo struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ApplyChanges(newTodo []Todo, file *os.File) {
	jsonData, err := json.MarshalIndent(newTodo, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := file.WriteString(string(jsonData)); err != nil {
		log.Fatal(err)
	}

}
