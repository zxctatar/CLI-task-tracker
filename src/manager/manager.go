package manager

import (
	"CLI-task-tracker/command"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Manager struct {
	commandsMap map[string]command.Command
	commandSlice []command.Command
}

func NewManager() *Manager {
	return &Manager{map[string]command.Command{}, []command.Command{}}
}

func (m *Manager) RegCommand(com command.Command) {
	m.commandsMap[com.GetName()] = com
	m.commandSlice = append(m.commandSlice, com)
}

func (m *Manager) GetCommandsMap() map[string]command.Command {
	return m.commandsMap
}

func (m *Manager) GetCommandsSlice() []command.Command {
	return m.commandSlice
}

func (m *Manager) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	stop := false

	for !stop {
		fmt.Print("--> ")

		if !scanner.Scan() {
			fmt.Println("bad input")
			continue
		}

		str := scanner.Text()
		trimStr := strings.Fields(str)

		if len(trimStr) == 0 {
			continue
		}

		selectCommand := m.commandsMap[trimStr[0]]

		if selectCommand == nil {
			fmt.Println("unknown command")
			continue
		}

		err := selectCommand.GetValidator().Check(trimStr[1:])

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = selectCommand.Execute(trimStr[1:])

		if err != nil {
			if errors.Is(err, command.ErrExit) {
				stop = true
				continue
			}
			fmt.Println(err.Error())
		}
	}
}