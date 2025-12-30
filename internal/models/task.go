package models

import (
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
