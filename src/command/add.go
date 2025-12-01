package command

import (
	"CLI-task-tracker/storage"
	"strings"
	"time"
)

type AddCommand struct {
	name string
	stor *storage.TaskStorage
}

func NewAddCommand(storage *storage.TaskStorage) *AddCommand {
	name := "add"
	return &AddCommand{name, storage}
}

func (ac *AddCommand) GetName() string {
	return ac.name
}

func (ac *AddCommand) Execute(args []string) error {
	currentTime := time.Now()
	format := "02.01.2006 15:04"
	formattedDateTime := currentTime.Format(format)

	var title string = strings.Join(args, " ")

	ac.stor.AddTask(title, formattedDateTime)

	return nil
}