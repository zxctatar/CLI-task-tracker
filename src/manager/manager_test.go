package manager

import (
	"CLI-task-tracker/command"
	"testing"
)

func TestRegCommand(t *testing.T) {
	manager := NewManager()

	tests := []struct{
		name string
		waitLen int
		com command.Command
	}{
		{"exit", 1, command.NewExitCommand(nil)},
		{"delete", 2, command.NewDeleteCommand(nil, nil)},
		{"add", 3, command.NewAddCommand(nil, nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager.RegCommand(tt.com)

			if len(manager.GetCommandsMap()) != tt.waitLen {
				t.Errorf("expected map len %d, got %d", len(manager.GetCommandsMap()), tt.waitLen)
			}

			if len(manager.GetCommandsSlice()) != tt.waitLen {
				t.Errorf("expected slice len %d, got %d", len(manager.GetCommandsSlice()), tt.waitLen)
			}
		})
	}
}