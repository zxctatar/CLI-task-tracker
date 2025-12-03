package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
	"strings"
)

type UpdateCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
}

func NewUpdateCommand(stor *storage.TaskStorage, val validator.Validator) *UpdateCommand {
	name := "update"
	return &UpdateCommand{name, stor, val}
}

func (uc *UpdateCommand) GetName() string {
	return uc.name
}

func (uc *UpdateCommand) GetValidator() validator.Validator {
	return uc.val
}

func (uc *UpdateCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("argument is not a number")
	}

	newTitle := strings.Join(args[1:], " ")

	formattedDateTime := utils.FormattedCurrentTime()

	err = uc.stor.UpdateTask(id, newTitle, formattedDateTime)

	return err
}