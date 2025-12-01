package command

import "CLI-task-tracker/storage"

type Command interface {
	GetName() string
	Execute(stor *storage.TaskStorage, args []string) error
}