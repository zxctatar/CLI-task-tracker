package main

import (
	"CLI-task-tracker/command"
	"CLI-task-tracker/manager"
	"CLI-task-tracker/storage"
	"CLI-task-tracker/validator"
	"log"
)

func main() {
	jsonStor := storage.NewJsonStorage("tasks.json")

	tasks, err := jsonStor.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	//TaskStorage
	taskStor := storage.NewTaskStorage(tasks)

	defer func() {
		jsonStor.Save(taskStor.GetTasks())
	}()

	//Validators
	notEmptyVal := validator.NotEmptyValidator{}
	emptyVal := validator.EmptyValudator{}
	minMaxOneVal := validator.NewMinMaxNumberValidator(1, 1) // for DeleteCommand, StartCommand, DoneCommand, CancelCommand
	minVal := validator.NewMinValidator(2)                   // for UpdateCommand

	//Commands
	add := command.NewAddCommand(taskStor, &notEmptyVal)
	exit := command.NewExitCommand(&emptyVal)
	list := command.NewListCommand(taskStor, &emptyVal)
	delete := command.NewDeleteCommand(taskStor, minMaxOneVal)
	update := command.NewUpdateCommand(taskStor, minVal)
	start := command.NewStartCommand(taskStor, minMaxOneVal)
	done := command.NewDoneCommand(taskStor, minMaxOneVal)
	cancel := command.NewCancelCommand(taskStor, minMaxOneVal)
	help := command.NewHelpCommand(&emptyVal)

	manager := manager.NewManager()
	manager.RegCommand(add)
	manager.RegCommand(update)
	manager.RegCommand(list)
	manager.RegCommand(delete)
	manager.RegCommand(start)
	manager.RegCommand(done)
	manager.RegCommand(cancel)
	manager.RegCommand(exit)
	manager.RegCommand(help)

	help.SetCommands(manager.GetCommandsSlice())

	manager.Run()
}
