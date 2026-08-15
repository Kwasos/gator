package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	lookup, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("the following command does not exist: %s", cmd.name)
	}
	return lookup(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("username is requred")
	}

	err := s.Config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Printf("set user: %s\n", cmd.args[0])
	return nil
}
