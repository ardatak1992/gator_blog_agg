package main

import (
	"log"
	"os"

	"github.com/ardatak1992/gator_blog_agg/internal/config"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Enter a command")
	}

	conf, err := config.Read()
	if err != nil {
		log.Fatal("Error occuiered")
	}

	st := state{
		&conf,
	}

	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	enteredCommand := command{os.Args[1], os.Args[2:]}

	err = commands.run(&st, enteredCommand)
	if err != nil {
		log.Fatal(err)
	}

}
