package storage

import (
	"CLI-task-tracker/task"
	"errors"
	"slices"
)

type TaskStorage struct {
	tasks  []*task.Task
	nextId int
}

func NewTaskStorage(tasks []*task.Task) *TaskStorage {
	nextId := findLastId(tasks)
	return &TaskStorage{tasks, nextId}
}

func (ts *TaskStorage) AddTask(title string, dateTime string) {
	newTask := task.NewTask(ts.nextId, title, dateTime, dateTime, task.NotStarted)
	ts.tasks = append(ts.tasks, newTask)
	ts.nextId++
}

func (ts *TaskStorage) DeleteTask(id int) error {
	for index, t := range ts.tasks {
		if t.GetId() == id {
			ts.tasks = slices.Delete(ts.tasks, index, index+1)
			return nil
		}
	}
	return errors.New("task with this id not found")
}

func (ts *TaskStorage) UpdateTask(id int, newDescription string, newLastUpdateTime string) error {
	t, err := ts.findById(id)

	if err != nil {
		return err
	}

	t.SetDescription(newDescription)
	t.SetLastUpdateTime(newLastUpdateTime)

	return nil
}

func (ts *TaskStorage) StartTask(id int, newLastUpdateTime string) error {
	t, err := ts.findById(id)

	if err != nil {
		return err
	}

	t.StartTask()
	t.SetLastUpdateTime(newLastUpdateTime)

	return nil
}

func (ts *TaskStorage) DoneTask(id int, newLastUpdateTime string) error {
	t, err := ts.findById(id)

	if err != nil {
		return err
	}

	t.DoneTask()
	t.SetLastUpdateTime(newLastUpdateTime)

	return nil
}

func (ts *TaskStorage) CancelTask(id int, newLastUpdateTime string) error {
	t, err := ts.findById(id)

	if err != nil {
		return err
	}

	t.CancelTask()
	t.SetLastUpdateTime(newLastUpdateTime)

	return nil
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

func (ts *TaskStorage) findById(id int) (*task.Task, error) {
	for _, t := range ts.tasks {
		if t.GetId() == id {
			return t, nil
		}
	}
	return nil, errors.New("task with this id not found")
}
