package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type StartCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
	description string
	usage string
	example string
}

func NewStartCommand(stor *storage.TaskStorage, val validator.Validator) *StartCommand {
	name := "start"
	description := "Start a task."
	usage := "start <task id>"
	example := "start 1"
	return &StartCommand{name, stor, val, description, usage, example}
}

func (sc *StartCommand) GetName() string {
	return sc.name
}

func (sc *StartCommand) GetValidator() validator.Validator {
	return sc.val
}

func (sc *StartCommand) GetDescription() string {
	return sc.description
}

func (sc *StartCommand) GetUsage() string {
	return sc.usage
}

func (sc *StartCommand) GetExample() string {
	return sc.example
}

func (sc *StartCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("argument is not a number")
	}

	formattedDateTime := utils.FormattedCurrentTime()

	err = sc.stor.StartTask(id, formattedDateTime)

	return err
}