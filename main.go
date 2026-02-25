package main

import (
	"fmt"
	"log"

	"github.com/ardatak1992/gator_blog_agg/internal/config"
)

func main() {

	conf, err := config.Read()
	if err != nil {
		log.Fatalf("Can't create config: %s", err.Error())
	}

	conf.SetUser("arda1222")

	conf, err = config.Read()
	if err != nil {
		log.Fatalf("Can't create config: %s", err.Error())
	}

	fmt.Printf("dbURL: %s\nusername: %s\n", conf.DbURL, conf.CurrentUserName)

}
