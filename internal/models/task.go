package models

import (
	"time"

	"github.com/TerryMinn/task-cli/internal/config"
)

type Callback func(int)

type Todo struct {
	Id          int              `json:"id"`
	Description string           `json:"description"`
	Status      config.Operation `json:"status"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}
