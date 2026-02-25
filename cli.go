package main

import (
	"fmt"

	"github.com/ardatak1992/gator_blog_agg/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	commandHandler, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", cmd.name)
	}

	err := commandHandler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
