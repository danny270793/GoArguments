package goarguments_test

import (
	"fmt"
	"testing"

	"github.com/danny270793/goarguments"
)

func TestActionNotFound(t *testing.T) {
	program := "call"
	arguments := []string{"runner", program}
	commands := []goarguments.Command{}
	_, err := goarguments.ParseCommand(commands, arguments[1:])
	errorObtained := err.Error()
	expectedError := fmt.Sprintf("action \"%s\" is invalid", program)
	if errorObtained != expectedError {
		t.Errorf("expected error %s, but %s was found", expectedError, errorObtained)
	}
}

func TestActionFound(t *testing.T) {
	program := "call"
	arguments := []string{"runner", program}
	commands := []goarguments.Command{
		{
			Name:        program,
			Description: "",
			Arguments:   map[string]string{},
			Commands:    []goarguments.Command{},
			Callback:    func(m map[string]string) {},
		},
	}
	chain, err := goarguments.ParseCommand(commands, arguments[1:])
	if err != nil {
		t.Errorf("unexpected error found: %s", err.Error())
	}
	if len(chain) != 1 {
		t.Errorf("chain must contains 1 callback but %d was found", len(chain))
	}
	chain[0].Callback(map[string]string{})
}

func TestChainSize(t *testing.T) {
	program := "program"
	subprogram := "subprogram"
	arguments := []string{"runner", program, subprogram}
	commands := []goarguments.Command{
		{
			Name:        program,
			Description: "",
			Arguments:   map[string]string{},
			Commands: []goarguments.Command{
				{
					Name:        subprogram,
					Description: "",
					Arguments:   map[string]string{},
					Commands:    []goarguments.Command{},
					Callback:    func(m map[string]string) {},
				},
			},
			Callback: func(m map[string]string) {},
		},
	}
	chain, err := goarguments.ParseCommand(commands, arguments[1:])
	if err != nil {
		t.Errorf("unexpected error found: %s", err.Error())
	}
	if len(chain) != 2 {
		t.Errorf("chain must contains 2 callback but %d was found", len(chain))
	}
}
