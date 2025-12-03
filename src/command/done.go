package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type DoneCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
}

func NewDoneCommand(stor *storage.TaskStorage, val validator.Validator) *DoneCommand {
	name := "done"
	return &DoneCommand{name, stor, val}
}

func (dc *DoneCommand) GetName() string {
	return dc.name
}

func (dc *DoneCommand) GetValidator() validator.Validator {
	return dc.val
}

func (dc *DoneCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("argument is not a number")
	}

	formattedDateTime := utils.FormattedCurrentTime()

	err = dc.stor.DoneTask(id, formattedDateTime)

	return err
}