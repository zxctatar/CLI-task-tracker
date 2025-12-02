package storage

import (
	"CLI-task-tracker/task"
)

type TaskStorage struct {
	tasks []*task.Task
	nextId int
}

func NewTaskStorage(tasks []*task.Task) *TaskStorage {
	nextId := findLastId(tasks)
	return &TaskStorage{tasks, nextId}
}

func (ts *TaskStorage) AddTask(title string, dateTime string) {
	newTask := task.NewTask(ts.nextId, title, dateTime, task.NotStarted)
	ts.tasks = append(ts.tasks, newTask)
	ts.nextId++
}

func (ts *TaskStorage) GetTasks() []*task.Task {
	return ts.tasks
}

func findLastId(tasks []*task.Task) int {
	var maxId int = 0

	for _, t := range tasks {
		if maxId < t.GetId() {
			maxId = t.GetId()
		}
	}

	return maxId + 1
} 