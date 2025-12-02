package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/validator"
	"strings"
	"time"
)

type AddCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
}

func NewAddCommand(stor *storage.TaskStorage, val validator.Validator) *AddCommand {
	name := "add"
	return &AddCommand{name, stor, val}
}

func (ac *AddCommand) GetName() string {
	return ac.name
}

func (ac *AddCommand) GetValidator() validator.Validator {
	return ac.val
}

func (ac *AddCommand) Execute(args []string) error {
	currentTime := time.Now()
	format := "02.01.2006 15:04"
	formattedDateTime := currentTime.Format(format)

	var title string = strings.Join(args, " ")

	ac.stor.AddTask(title, formattedDateTime)

	return nil
}