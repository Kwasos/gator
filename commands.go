package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Kwasos/gator/internal/database"
	"github.com/google/uuid"
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
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println("error getting user:", err)
		os.Exit(1)
	}

	err = s.Config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Printf("set user: %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("username is required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		fmt.Println("error creating user:", err)
		os.Exit(1)
	}
	fmt.Printf("User created successfully\n")

	err = s.Config.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Printf("User created: %+v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println("error deleting users:", err)
		os.Exit(1)
	}
	fmt.Println("Deleted users successfully")
	return nil
}
