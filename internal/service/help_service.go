package service

import (
	"fmt"
	"strings"

	"github.com/TerryMinn/task-cli/internal/config"
	"github.com/fatih/color"
)

func MainHelper() {
	fmt.Println(color.GreenString("Task-Cli"), color.RedString("v1.0.0"))
	fmt.Println(strings.Repeat(" ", 50))
	fmt.Println("Available commands:")
	fmt.Println(strings.Repeat(" ", 50))
	fmt.Printf(" %-5s %-20s %-10s", "", config.Add, "Add new todo task\n")
	fmt.Printf(" %-5s %-20s %-20s", "", config.Update, "Update your task with specific ID\n")
	fmt.Printf(" %-5s %-20s %-20s", "", config.Delete, "Delete your task with specific ID\n")
	fmt.Printf(" %-5s %-20s %-20s", "", config.MarkInProgress, "Change your task status to in progress with specific ID\n")
	fmt.Printf(" %-5s %-20s %-20s", "", config.MarkDone, "Change your task status to done with specific ID\n")
	fmt.Printf(" %-5s %-20s %-20s", "", config.List, "Check your task list\n")
	fmt.Printf(" %-5s %-20s %-20s %-20s", "", config.List, config.Todo, "Check your todo task list\n")
	fmt.Printf(" %-5s %-20s %-20s %-20s", "", config.List, config.InProgress, "Check your in progress task list\n")
	fmt.Printf(" %-5s %-20s %-20s %-20s", "", config.List, config.Done, "Check your done task list\n")
	fmt.Println(strings.Repeat(" ", 50))
	fmt.Println(strings.Repeat(" ", 50))
	color.Green(" ♥ This project is Task Tracker milestone from roadmap.sh:\n")
	fmt.Println("https://roadmap.sh/projects/task-tracker")
}
