package config

type Operation string

const (
	Add            Operation = "add"
	Delete         Operation = "delete"
	Update         Operation = "update"
	List           Operation = "list"
	Done           Operation = "done"
	InProgress     Operation = "in-progress"
	Todo           Operation = "todo"
	MarkInProgress Operation = "mark-in-progress"
	MarkDone       Operation = "mark-done"
	MarkTodo       Operation = "mark-todo"
)

type Application struct {
	Operation
	ListOperation Operation
	FilePath      string
	Version       string
}

func (app Application) GetVersion() string {
	return app.Version
}
