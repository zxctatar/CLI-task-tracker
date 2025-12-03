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
}

func NewListCommand(stor *storage.TaskStorage, val validator.Validator) *ListCommand {
	name := "list"
	return &ListCommand{name, stor, val}
}

func (lc *ListCommand) GetName() string {
	return lc.name
}

func (lc *ListCommand) GetValidator() validator.Validator {
	return lc.val
}

func (lc *ListCommand) Execute(args []string) error {
	tasks := lc.stor.GetTasks()

	if len(tasks) == 0 {
		return errors.New("no tasks")
	}

	idLen, titleLen, createdTimeLen, lastUpdateTimeLen, statLen := getFieldSizes(tasks)

	fmt.Println()

	fmt.Printf(
		"%-*s | %-*s | %-*s | %-*s | %-*s\n",
		idLen, "ID",
		titleLen, "Title",
		createdTimeLen, "CreatedTime",
		statLen, "Status",
		lastUpdateTimeLen, "LastUpdate",
	)

	totalWidth := idLen + titleLen + createdTimeLen + statLen + lastUpdateTimeLen + 12
	fmt.Println(strings.Repeat("-", totalWidth))

	for _, t := range tasks {
		fmt.Printf(
			"%-*d | %-*s | %-*s | %-*s | %-*s\n",
			idLen, t.GetId(),
			titleLen, t.GetTitle(),
			createdTimeLen, t.GetCreatedTime(),
			statLen, t.GetStatus().String(),
			lastUpdateTimeLen, t.GetLastUpdateTime(),
		)
	}

	fmt.Println()

	return nil
}

func getFieldSizes(tasks []*task.Task) (idLen, titleLen, createdTimeLen, lastUpdateTimeLen, statLen int){
	idLen = 2
	titleLen = 5
	createdTimeLen = 10
	lastUpdateTimeLen = 10
	statLen = 12

	for _, t := range tasks {
		idStr := strconv.Itoa(t.GetId())
		if idLen < len(idStr) {
			idLen = len(idStr)
		}
		
		titleRun := []rune(t.GetTitle())

		if titleLen < len(titleRun) {
			titleLen = len(titleRun)
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