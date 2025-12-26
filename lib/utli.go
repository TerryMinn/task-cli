package lib

import (
	"github.com/TerryMinn/task-cli/cmd"
)

func StatusChecker(status cmd.Status) string {
	if status == cmd.TASK {
		return "Todo"
	} else if status == cmd.IN_PROGRESS {
		return "In Progress"
	} else {
		return "Done"
	}
}
