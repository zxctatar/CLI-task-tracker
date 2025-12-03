package storage

import (
	"CLI-task-tracker/task"
	"testing"
)

func TestAddTask(t *testing.T) {
	ts := NewTaskStorage(nil)

	tests := []struct{
		id int
		description string
		createdTime string
		lastUpdateTime string
		stat task.Status
	}{
		{1, "Buy milk", "01-01-1999 00:00", "01-01-1999 00:00", task.NotStarted},
		{2, "Walk dog", "02-11-2025 11:11", "02-11-2025 11:11", task.NotStarted},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ts.AddTask(tt.description, tt.createdTime)
			
			oneTask, _ := ts.findById(tt.id)

			if len(ts.GetTasks()) != tt.id {
				t.Errorf("expected %d task, got %d", tt.id, len(ts.GetTasks()))
			}

			if oneTask.GetId() != tt.id {
				t.Errorf("expected task with id %d, got %d", tt.id, oneTask.GetId())
			}

			if oneTask.GetDescription() != tt.description {
				t.Errorf("expected task with description '%s', got '%s'", tt.description, oneTask.GetDescription())
			}
			
			if oneTask.GetCreatedTime() != tt.createdTime {
				t.Errorf("expected task with created time '%s', got '%s'", tt.createdTime,oneTask.GetCreatedTime())
			}

			if oneTask.GetLastUpdateTime() != tt.lastUpdateTime {
				t.Errorf("expected task with last update time '%s', got '%s'", tt.lastUpdateTime, oneTask.GetLastUpdateTime())
			}

			if oneTask.GetStatus() != task.NotStarted {
				t.Errorf("expected task with status '%s', got '%s'", tt.stat.String(), oneTask.GetStatus().String())
			}
		})
	} 
}

func TestDeleteTask(t *testing.T) {
	ts := NewTaskStorage(nil)

	ts.AddTask("Buy milk", "01-01-1999 00:00")
	ts.AddTask("Walk dog", "01-01-1999 00:00")

	tests := []struct{
		name string
		id int
		len int
	}{
		{"Buy milk", 1, 1},
		{"Walk dog", 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts.DeleteTask(tt.id)
			tasks := ts.GetTasks()
			
			if len(tasks) != tt.len {
				t.Errorf("expected len %d, got %d", tt.len, len(tasks))
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	ts := NewTaskStorage(nil)

	ts.AddTask("Buy milk", "01-01-1999 00:00")
	ts.AddTask("Walk dog", "01-01-1999 00:00")

	tests := []struct{
		name string
		id int
		wantDescription string
		wantLastUpdate string
	}{
		{"Buy milk", 1, "Buy bread", "02-11-2025 11:11"},
		{"Walk dog", 2, "Go to school", "03-03-2005 12:34"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts.UpdateTask(tt.id, tt.wantDescription, tt.wantLastUpdate)

			oneTask, _ := ts.findById(tt.id)

			if oneTask.GetDescription() != tt.wantDescription {
				t.Errorf("expected task with description '%s', got '%s'", oneTask.Description, tt.wantDescription)
			}

			if oneTask.GetLastUpdateTime() != tt.wantLastUpdate {
				t.Errorf("expected task with last update time '%s', got '%s'", oneTask.LastUpdateTime, tt.wantLastUpdate)
			}
		})
	}
}

func TestChangeStatusTask(t *testing.T) {
	ts := NewTaskStorage(nil)

	ts.AddTask("Buy milk", "01-01-1999 00:00")
	ts.AddTask("Walk dog", "01-01-1999 00:00")

	tests := []struct{
		id int
		name string
		firstWantStatus task.Status
		secondWantStatus task.Status
		thirdWantStatus task.Status
		firstWantLastUpdate string
		secondWantLastUpdate string
		thirdWantLastUpdate string
	}{
		{1, "Buy milk", task.Started, task.Done, task.NotStarted, "01-01-1999 00:01", "01-01-1999 00:02", "01-01-1999 00:03"},
		{2, "Walk dog", task.Started, task.Done, task.NotStarted, "01-01-1999 00:01", "01-01-1999 00:02", "01-01-1999 00:03"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts.StartTask(tt.id, tt.firstWantLastUpdate)

			oneTask, _ := ts.findById(tt.id)

			if oneTask.GetStatus() != tt.firstWantStatus {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.Stat.String(), tt.firstWantStatus.String())
			}

			if oneTask.GetLastUpdateTime() != tt.firstWantLastUpdate {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.GetLastUpdateTime(), tt.firstWantLastUpdate)
			}

			ts.DoneTask(tt.id, tt.secondWantLastUpdate)

			if oneTask.GetStatus() != tt.secondWantStatus {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.Stat.String(), tt.secondWantStatus.String())
			}

			if oneTask.GetLastUpdateTime() != tt.secondWantLastUpdate {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.GetLastUpdateTime(), tt.secondWantLastUpdate)
			}

			ts.CancelTask(tt.id, tt.thirdWantLastUpdate)

			if oneTask.GetStatus() != tt.thirdWantStatus {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.Stat.String(), tt.thirdWantStatus.String())
			}

			if oneTask.GetLastUpdateTime() != tt.thirdWantLastUpdate {
				t.Errorf("expected task with status '%s', got '%s'", oneTask.GetLastUpdateTime(), tt.thirdWantLastUpdate)
			}
		})
	}
}