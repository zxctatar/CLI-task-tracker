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
	description string
	usage string
	example string
}

func NewUpdateCommand(stor *storage.TaskStorage, val validator.Validator) *UpdateCommand {
	name := "update"
	description := "Update the description of a task."
	usage := "update <task id> <new task description>"
	example := "update 1 Buy bread"
	return &UpdateCommand{name, stor, val, description, usage, example}
}

func (uc *UpdateCommand) GetName() string {
	return uc.name
}

func (uc *UpdateCommand) GetValidator() validator.Validator {
	return uc.val
}

func (uc *UpdateCommand) GetDescription() string {
	return uc.description
}

func (uc *UpdateCommand) GetUsage() string {
	return uc.usage
}

func (uc *UpdateCommand) GetExample() string {
	return uc.example
}

func (uc *UpdateCommand) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("argument is not a number")
	}

	newDescription := strings.Join(args[1:], " ")

	formattedDateTime := utils.FormattedCurrentTime()

	err = uc.stor.UpdateTask(id, newDescription, formattedDateTime)

	return err
}