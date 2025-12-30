package utils

import (
	"os"
)

type Input struct {
	Op     string
	Option string
	Value  string
}

func GetUserInput(cb func()) *Input {
	input := Input{}

	if len(os.Args) == 3 {
		input.Op = os.Args[1]
		input.Value = os.Args[2]
	} else if len(os.Args) == 4 {
		input.Op = os.Args[1]
		input.Option = os.Args[2]
		input.Value = os.Args[3]
	} else {
		cb()
	}

	return &input
}
