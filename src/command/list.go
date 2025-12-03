package command

import (
	"CLI-task-tracker/storage"
	"CLI-task-tracker/task"
	"CLI-task-tracker/validator"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ListCommand struct {
	name string
	stor *storage.TaskStorage
	val validator.Validator
	description string
	usage string
	example string
}

func NewListCommand(stor *storage.TaskStorage, val validator.Validator) *ListCommand {
	name := "list"
	description := "Show all tasks."
	usage := "list"
	example := "list"
	return &ListCommand{name, stor, val, description, usage, example}
}

func (lc *ListCommand) GetName() string {
	return lc.name
}

func (lc *ListCommand) GetValidator() validator.Validator {
	return lc.val
}

func (lc *ListCommand) GetDescription() string {
	return lc.description
}

func (lc *ListCommand) GetUsage() string {
	return lc.usage
}

func (lc *ListCommand) GetExample() string {
	return lc.example
}

func (lc *ListCommand) Execute(args []string) error {
	tasks := lc.stor.GetTasks()

	if len(tasks) == 0 {
		return errors.New("no tasks")
	}

	idLen, descriptionLen, createdTimeLen, lastUpdateTimeLen, statLen := getListFieldSizes(tasks)

	fmt.Println()

	fmt.Printf(
		"%-*s | %-*s | %-*s | %-*s | %-*s\n",
		idLen, "ID",
		descriptionLen, "Description",
		createdTimeLen, "CreatedTime",
		statLen, "Status",
		lastUpdateTimeLen, "LastUpdate",
	)

	totalWidth := idLen + descriptionLen + createdTimeLen + statLen + lastUpdateTimeLen + 12
	fmt.Println(strings.Repeat("-", totalWidth))

	for _, t := range tasks {
		fmt.Printf(
			"%-*d | %-*s | %-*s | %-*s | %-*s\n",
			idLen, t.GetId(),
			descriptionLen, t.GetDescription(),
			createdTimeLen, t.GetCreatedTime(),
			statLen, t.GetStatus().String(),
			lastUpdateTimeLen, t.GetLastUpdateTime(),
		)
	}

	fmt.Println()

	return nil
}

func getListFieldSizes(tasks []*task.Task) (idLen, descriptionLen, createdTimeLen, lastUpdateTimeLen, statLen int){
	idLen = 2
	descriptionLen = 11
	createdTimeLen = 10
	lastUpdateTimeLen = 10
	statLen = 12

	for _, t := range tasks {
		idStr := strconv.Itoa(t.GetId())
		if idLen < len(idStr) {
			idLen = len(idStr)
		}
		
		descriptionRun := []rune(t.GetDescription())

		if descriptionLen < len(descriptionRun) {
			descriptionLen = len(descriptionRun)
		}

		dateTimeRun := []rune(t.GetCreatedTime())

		if createdTimeLen < len(dateTimeRun) {
			createdTimeLen = len(dateTimeRun)
		}

		lastUpdateRun := []rune(t.GetLastUpdateTime())

		if lastUpdateTimeLen < len(lastUpdateRun) {
			lastUpdateTimeLen = len(lastUpdateRun)
		}
	}

	return
}