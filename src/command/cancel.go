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
}

func NewCancelCommand(stor *storage.TaskStorage, val validator.Validator) *CancelCommand {
	name := "cancel"
	return &CancelCommand{name, stor, val}
}

func (cc *CancelCommand) GetName() string {
	return cc.name
}

func (cc *CancelCommand) GetValidator() validator.Validator {
	return cc.val
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