package main

import (
	"fmt"
	"log"

	"github.com/alexparco/rest-api-todo/config"
	"github.com/alexparco/rest-api-todo/routes"
)

func main() {
	config, err := config.NewConfig("./config.yml")
	if err != nil {
		log.Fatalf("File error %e", err)
	}

	r := routes.NewRouter(config)
	r.Run(fmt.Sprintf(":%s", config.Port))
}
