package goarguments

import (
	"errors"
	"fmt"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Arguments   map[string]string
	Commands    []Command
	Callback    func(map[string]string)
}

func repeat(character string, quantity int) string {
	value := ""
	for index := 0; index < quantity; index++ {
		value += character
	}
	return value
}

func ShowHelp(commands []Command) {
	fmt.Printf("Usage:\n")
	showHelp(commands, 1)
}

func showHelp(commands []Command, level int) {
	for _, command := range commands {
		fmt.Printf("%s%s: %s\n", repeat("  ", level), command.Name, command.Description)
		for key, value := range command.Arguments {
			fmt.Printf("%s%s: %s\n", repeat("  ", level+1), key, value)
		}
		showHelp(command.Commands, level+1)
	}
}

func ParseCommand(commands []Command, arguments []string) ([]Command, error) {
	chain, err := parseCommand(commands, arguments)
	if err != nil {
		return []Command{}, err
	}

	return reverseArray(chain), nil
}

func parseCommand(commands []Command, arguments []string) ([]Command, error) {
	argument := arguments[0]
	commandFound := false
	for _, command := range commands {
		if argument == command.Name {
			commandFound = true
			for index, subargument := range arguments[1:] {
				if strings.HasPrefix(subargument, "--") {
					keyValue := strings.SplitN(subargument, "=", 2)

					key := keyValue[0]
					if _, exists := command.Arguments[key]; !exists {
						return []Command{}, errors.New(fmt.Sprintf("argument \"%s\" is invalid on command \"%s\"", key, argument))
					}

					if len(keyValue) == 2 {
						command.Arguments[key] = keyValue[1]
					} else {
						command.Arguments[key] = "true"
					}
				} else {
					chain, err := ParseCommand(command.Commands, arguments[index+1:])
					if err != nil {
						return []Command{}, errors.New(fmt.Sprintf("action \"%s\" is invalid", subargument))
					}
					chain = append(chain, command)
					return chain, nil
				}
			}
		}

		if commandFound {
			return []Command{command}, nil
		}
	}
	return []Command{}, errors.New(fmt.Sprintf("action \"%s\" is invalid", argument))
}

func reverseArray(commands []Command) []Command {
	length := len(commands)
	reversed := make([]Command, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = commands[j]
	}

	return reversed
}
