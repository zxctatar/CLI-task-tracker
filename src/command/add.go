package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/utils"
	"CLI-task-tracker/validator"
	"strings"
)

type AddCommand struct {
	name        string
	stor        *storage.TaskStorage
	val         validator.Validator
	description string
	usage       string
	example     string
}

func NewAddCommand(stor *storage.TaskStorage, val validator.Validator) *AddCommand {
	name := "add"
	description := "Add a new task."
	usage := "add <task description>"
	example := "add Buy milk"
	return &AddCommand{name, stor, val, description, usage, example}
}

func (ac *AddCommand) GetName() string {
	return ac.name
}

func (ac *AddCommand) GetValidator() validator.Validator {
	return ac.val
}

func (ac *AddCommand) GetDescription() string {
	return ac.description
}

func (ac *AddCommand) GetUsage() string {
	return ac.usage
}

func (ac *AddCommand) GetExample() string {
	return ac.example
}

func (ac *AddCommand) Execute(args []string) error {
	formattedDateTime := utils.FormattedCurrentTime()

	var title string = strings.Join(args, " ")

	ac.stor.AddTask(title, formattedDateTime)

	return nil
}
