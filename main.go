package main

import (
	"fmt"
	"os"

	"github.com/Kwasos/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	appState := state{Config: &cfg}

	cmds := commands{
		commands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("missing argument")
		os.Exit(1)
	}

	commandName := os.Args[1]
	argsSlice := os.Args[2:]
	fullCommand := command{name: commandName, args: argsSlice}
	err = cmds.run(&appState, fullCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
