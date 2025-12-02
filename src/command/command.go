package command

import "CLI-task-tracker/validator"

type Command interface {
	GetName() string
	Execute(args []string) error
	GetValidator() validator.Validator
}