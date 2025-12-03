package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type DeleteCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
}

func NewDeleteCommand(stor *storage.TaskStorage, val validator.Validator) *DeleteCommand {
	name := "delete"
	return &DeleteCommand{name, stor, val}
}

func (dc *DeleteCommand) GetName() string {
	return dc.name
}

func (dc *DeleteCommand) GetValidator() validator.Validator {
	return dc.val
}

func (dc *DeleteCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("argument is not a number")
	}
	dc.stor.DeleteTask(id)
	return err
}