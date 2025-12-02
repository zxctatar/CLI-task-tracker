package command

import "CLI-task-tracker/validator"

var ErrExit error

type ExitCommand struct {
	name string
	val validator.Validator
}

func NewExitCommand(val validator.Validator) *ExitCommand {
	name := "exit"
	return &ExitCommand{name, val}
}

func (ec *ExitCommand) GetName() string {
	return ec.name
}

func (ec *ExitCommand) GetValidator() validator.Validator {
	return ec.val
}

func (ec *ExitCommand) Execute(args []string) error {
	return ErrExit
}