package command

import (
	"CLI-task-tracker/validator"
	"errors"
)

var ErrExit error = errors.New("exit")

type ExitCommand struct {
	name        string
	val         validator.Validator
	description string
	usage       string
	example     string
}

func NewExitCommand(val validator.Validator) *ExitCommand {
	name := "exit"
	description := "Exit the application."
	usage := "exit"
	example := "exit"
	return &ExitCommand{name, val, description, usage, example}
}

func (ec *ExitCommand) GetName() string {
	return ec.name
}

func (ec *ExitCommand) GetValidator() validator.Validator {
	return ec.val
}

func (ec *ExitCommand) GetDescription() string {
	return ec.description
}

func (ec *ExitCommand) GetUsage() string {
	return ec.usage
}

func (ec *ExitCommand) GetExample() string {
	return ec.example
}

func (ec *ExitCommand) Execute(args []string) error {
	return ErrExit
}
