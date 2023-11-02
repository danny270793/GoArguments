package main

import (
	"fmt"
	"os"

	"github.com/danny270793/goarguments"
)

func main() {
	commands := []goarguments.Command{
		{
			Name: "container", Description: "Container actions", Arguments: map[string]string{"--driver": "docker"}, Commands: []goarguments.Command{
				{
					Name: "manage", Description: "Manage containers", Arguments: map[string]string{}, Commands: []goarguments.Command{
						{
							Name: "export", Description: "Export container images to filesystem", Arguments: map[string]string{"--folder": "./images"}, Commands: []goarguments.Command{}, Callback: func(arguments map[string]string) {
								fmt.Println(arguments)
							},
						},
						{
							Name: "import", Description: "Import container images from filesystem", Arguments: map[string]string{"--folder": "./images"}, Commands: []goarguments.Command{}, Callback: func(arguments map[string]string) {
								fmt.Println(arguments)
							},
						},
					}, Callback: func(arguments map[string]string) {
						fmt.Println(arguments)
					},
				},
			}, Callback: func(arguments map[string]string) {
				fmt.Println(arguments)
			},
		},
		{
			Name: "version", Description: "Show version of the program", Arguments: map[string]string{}, Commands: []goarguments.Command{}, Callback: func(arguments map[string]string) {
				fmt.Println("version 1.0.0")
			},
		},
	}

	chain, err := goarguments.ParseCommand(commands, os.Args[1:])
	if err != nil {
		fmt.Printf("%s\n\n", err)
		goarguments.ShowHelp(commands)
		os.Exit(1)
	}

	for _, item := range chain {
		item.Callback(item.Arguments)
	}
}
