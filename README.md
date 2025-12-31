# 📝 Task Tracker (Go)

A simple **Task Tracker** project built with **Go** as part of the [roadmap.sh Task Tracker challenge](https://roadmap.sh/projects/task-tracker). This project helps users manage tasks efficiently with basic CRUD operations via a CLI interface.

## 🚀 Installation

To install and run this CLI project, you need to have Go installed on your system (version 1.16 or higher).

### Prerequisites
- Go 1.16+
- Git

### Installation Steps

Install on local device: 
   ```bash
    go install github.com/TerryMinn/task-cli/cmd/task-cli@latest
   ```

or

Install from source:
1. Clone the repository:
   ```bash
   git clone https://github.com/TerryMinn/task-cli
   cd task-cli
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o task-cli ./cmd/task-cli
   ```

4. Alternatively, you can install it directly using Go:
   ```bash
   go install ./cmd/task-cli
   ```

## 🛠️ How to Run

After installation, you can run the CLI in the following ways:

1. If built locally:
   ```bash
   ./task-cli [command]
   ```

2. If installed via `go install`:
   ```bash
   task-cli [command]
   ```

### Available Commands
- `task-cli add "Task description"` - Add a new task
- `task-cli list` - List all tasks
- `task-cli list mark-in-progress` - List all in progress tasks
- `task-cli list mark-todo` - List all todo tasks
- `task-cli list mark-done` - List all done tasks
- `task-cli update [task-id]` - Mark a task as complete
- `task-cli delete [task-id]` - Delete a task
- `task-cli mark-in-progress [task-id]` - Change in progress status
- `task-cli mark-done [task-id]` - Change done status

## 📁 Folder Structure

```
task-cli/
├── cmd/
│   └── task-cli/           # Main application entry point
├── internal/               # Internal packages (not importable by other projects)
├── pkg/                    # Public packages (importable by other projects)
├── e2e/                    # End-to-end tests
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── Makefile                # Build and automation commands
├── README.md               # Project documentation
└── .gitignore             # Git ignore rules
```

- **cmd/**: Contains the main application. Each subdirectory in cmd/ is a separate command.
- **internal/**: Contains private application and library code that is not intended to be used by other projects.
- **pkg/**: Contains public library code that can be imported and used by other projects.
- **e2e/**: Contains end-to-end tests to verify the complete functionality of the application.