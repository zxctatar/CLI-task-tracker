package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type CancelCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
	description string
	usage string
	example string
}

func NewCancelCommand(stor *storage.TaskStorage, val validator.Validator) *CancelCommand {
	name := "cancel"
	description := "Cancel a task and mark it as not started."
	usage := "cancel <task id>"
	example := "cancel 1"
	return &CancelCommand{name, stor, val, description, usage, example}
}

func (cc *CancelCommand) GetName() string {
	return cc.name
}

func (cc *CancelCommand) GetValidator() validator.Validator {
	return cc.val
}

func (cc *CancelCommand) GetDescription() string {
	return cc.description
}

func (cc *CancelCommand) GetUsage() string {
	return cc.usage
}

func (cc *CancelCommand) GetExample() string {
	return cc.example
}

func (cc *CancelCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("argument is not a number")
	}

	formattedDateTime := utils.FormattedCurrentTime()

	err = cc.stor.CancelTask(id, formattedDateTime)

	return err
}