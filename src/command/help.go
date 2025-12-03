package command

import (
	"CLI-task-tracker/validator"
	"errors"
	"fmt"
	"strings"
)

type HelpCommand struct {
	name string
	val validator.Validator
	commands []Command
	description string
	usage string
	example string
}

func NewHelpCommand(val validator.Validator) *HelpCommand {
	name := "help"
	description := "Show available commands."
	usage := "help"
	example := "help"

	return &HelpCommand{
		name: name,
		val: val, 
		description: description,
		usage: usage,
		example: example,
	}
}

func (hc *HelpCommand) SetCommands(commands []Command) {
	hc.commands = commands
}

func (hc *HelpCommand) GetName() string {
	return hc.name
}

func (hc *HelpCommand) GetValidator() validator.Validator {
	return hc.val
}

func (hc *HelpCommand) GetDescription() string {
	return hc.description
}

func (hc *HelpCommand) GetUsage() string {
	return hc.usage
}

func (hc *HelpCommand) GetExample() string {
	return hc.example
}

func (hc *HelpCommand) Execute(args []string) error {
	if len(hc.commands) == 0 {
		return errors.New("commands not found")
	}

	commandLen, descriptionLen, usageLen, exampleLen := hc.getHelpFieldSizes()

	fmt.Println()

	fmt.Printf(
		"%-*s | %-*s | %-*s | %-*s\n",
		commandLen, "Command",
		descriptionLen, "Description",
		usageLen, "Usage",
		exampleLen, "Example",
	)

	totalWidth := commandLen + descriptionLen + usageLen + exampleLen + 9
	fmt.Println(strings.Repeat("-", totalWidth))

	for _, c := range hc.commands {
		fmt.Printf(
			"%-*s | %-*s | %-*s | %-*s\n",
			commandLen, c.GetName(),
			descriptionLen, c.GetDescription(),
			usageLen, c.GetUsage(),
			exampleLen, c.GetExample(),
		)
	}

	fmt.Println()

	return nil
}

func (hc *HelpCommand) getHelpFieldSizes() (commandLen, descriptionLen, usageLen, exampleLen int) {
	commandLen = 7
	descriptionLen = 11
	usageLen = 5
	exampleLen = 7

	for _, c := range hc.commands {
		commandRun := []rune(c.GetName())

		if len(commandRun) > commandLen {
			commandLen = len(commandRun)
		}

		descriptionRun := []rune(c.GetDescription())

		if len(descriptionRun) > descriptionLen {
			descriptionLen = len(descriptionRun)
		}

		usageRun := []rune(c.GetUsage())

		if len(usageRun) > usageLen {
			usageLen = len(usageRun)
		}

		exampleRun := []rune(c.GetExample())

		if len(exampleRun) > exampleLen {
			exampleLen = len(exampleRun)
		}
	}

	return
}