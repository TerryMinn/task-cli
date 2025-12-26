package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/TerryMinn/task-cli/cmd"
	"github.com/TerryMinn/task-cli/lib"
	"github.com/fatih/color"
)

func main() {
	file, fileError := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)
	if fileError != nil {
		log.Fatal(fileError)
	}

	defer file.Close()

	raw, dataError := io.ReadAll(file)
	if dataError != nil {
		log.Fatal(dataError)
	}

	var tasks []cmd.Todo

	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &tasks); err != nil {
			log.Fatal(err)
		}
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: missing argument. Usage: task-cli <file>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		value := os.Args[2]
		newTask := cmd.Todo{
			Id:          len(tasks) + 1,
			Description: value,
			Status:      cmd.TASK,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, newTask)
		cmd.ApplyChanges(tasks, "task.json")
		color.Green("Add new task complete with ID : %d", len(tasks))
		break

	case "update":
		value := os.Args[2]
		cmd.IndexFinder(value, func(target int) {
			updateValue := os.Args[3]
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Description = updateValue
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Update task complete with ID : %d", target)
		})
		break

	case "delete":
		value := os.Args[2]
		cmd.IndexFinder(value, func(target int) {
			for i, task := range tasks {

				if task.Id == target {
					tasks = append(tasks[:i], tasks[i+1:]...)
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Delete task complete with ID : %d", target)
		})
		break

	case "list":
		if len(os.Args) > 2 {
			goto OPTION
		}

		fmt.Printf("%-5s %-20s %-10s\n", "ID", "Description", "Status")
		fmt.Println("-------------------------------------------")
		for _, task := range tasks {
			fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, lib.StatusChecker(task.Status))
		}

	OPTION:
		fmt.Printf("%-5s %-20s %-10s\n", "ID", "Description", "Status")
		fmt.Println("-------------------------------------------")
		switch os.Args[2] {
		case "todo":
			for _, task := range tasks {
				if task.Status == cmd.TASK {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, lib.StatusChecker(task.Status))
				}
			}
			break
		case "in-progress":
			for _, task := range tasks {
				if task.Status == cmd.IN_PROGRESS {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, lib.StatusChecker(task.Status))
				}
			}
			break
		case "done":
			for _, task := range tasks {
				if task.Status == cmd.DONE {
					fmt.Printf("%-5d %-20s %-10s\n", task.Id, task.Description, lib.StatusChecker(task.Status))
				}
			}
			break
		default:
			fmt.Println("Error: invalid argument. Usage: task-cli <file>")
			os.Exit(1)
		}
		break

	case "mark-in-progress":
		value := os.Args[2]
		cmd.IndexFinder(value, func(target int) {
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Status = 1
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Start task complete with ID : %d", target)
		})
		break

	case "mark-done":
		value := os.Args[2]
		cmd.IndexFinder(value, func(target int) {
			for i, task := range tasks {
				if task.Id == target {
					tasks[i].Status = 2
				}
			}
			cmd.ApplyChanges(tasks, "task.json")
			color.Green("Done task complete with ID : %d", target)
		})
		break

	default:
		fmt.Printf("task management command not found\n")
	}
}
