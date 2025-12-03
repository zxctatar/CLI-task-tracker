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
}

func NewStartCommand(stor *storage.TaskStorage, val validator.Validator) *StartCommand {
	name := "start"
	return &StartCommand{name, stor, val}
}

func (sc *StartCommand) GetName() string {
	return sc.name
}

func (sc *StartCommand) GetValidator() validator.Validator {
	return sc.val
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