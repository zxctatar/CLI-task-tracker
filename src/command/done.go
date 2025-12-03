package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"errors"
	"strconv"
)

type DoneCommand struct {
	name        string
	stor        *storage.TaskStorage
	val         validator.Validator
	description string
	usage       string
	example     string
}

func NewDoneCommand(stor *storage.TaskStorage, val validator.Validator) *DoneCommand {
	name := "done"
	description := "Mark a task as done."
	usage := "done <task id>"
	example := "done 1"
	return &DoneCommand{name, stor, val, description, usage, example}
}

func (dc *DoneCommand) GetName() string {
	return dc.name
}

func (dc *DoneCommand) GetValidator() validator.Validator {
	return dc.val
}

func (dc *DoneCommand) GetDescription() string {
	return dc.description
}

func (dc *DoneCommand) GetUsage() string {
	return dc.usage
}

func (dc *DoneCommand) GetExample() string {
	return dc.example
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
