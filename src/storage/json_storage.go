package storage

import (
	"CLI-task-tracker/task"
	"encoding/json"
	"io"
	"os"
)

type JsonStorage struct {
	path string
}

func NewJsonStorage(path string) *JsonStorage {
	return &JsonStorage{path}
}

func (js *JsonStorage) Load() ([]*task.Task, error) {
	file, err := os.Open(js.path)

	if err != nil {
		if os.IsNotExist(err) {
			return []*task.Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []*task.Task{}, nil
	}

	var tasks []*task.Task

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (js *JsonStorage) Save(tasks []*task.Task) error {
	file, err := os.Create(js.path)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}
