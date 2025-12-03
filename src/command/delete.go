package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type DeleteCommand struct {
	name        string
	stor        *storage.TaskStorage
	val         validator.Validator
	description string
	usage       string
	example     string
}

func NewDeleteCommand(stor *storage.TaskStorage, val validator.Validator) *DeleteCommand {
	name := "delete"
	description := "Delete task."
	usage := "delete <task id>"
	example := "delete 1"
	return &DeleteCommand{name, stor, val, description, usage, example}
}

func (dc *DeleteCommand) GetName() string {
	return dc.name
}

func (dc *DeleteCommand) GetValidator() validator.Validator {
	return dc.val
}

func (dc *DeleteCommand) GetDescription() string {
	return dc.description
}

func (dc *DeleteCommand) GetUsage() string {
	return dc.usage
}

func (dc *DeleteCommand) GetExample() string {
	return dc.example
}

func (dc *DeleteCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("argument is not a number")
	}
	dc.stor.DeleteTask(id)
	return err
}
