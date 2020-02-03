package main

import (
	"fmt"
	"log"

	"github.com/Liberkeys/api-skeleton/infrastructure/application"
)

func main() {
	fmt.Println("Starting API server...")

	// Dependencies registration into object
	app := application.New()
	server := app.Server()
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
