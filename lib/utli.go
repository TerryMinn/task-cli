package lib

import (
	"github.com/TerryMinn/task-cli/cmd"
)

func StatusChecker(status cmd.Status) string {
	if status == 0 {
		return "Todo"
	} else if status == 1 {
		return "In Progress"
	} else {
		return "Done"
	}
}
