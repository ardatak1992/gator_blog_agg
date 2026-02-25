package main

import "fmt"

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("no username entered")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	println("User has been set.")

	return nil
}
