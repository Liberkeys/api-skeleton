package main

import (
	"fmt"
	"log"

	"github.com/Liberkeys/api-skeleton/adapters/http"
	"github.com/Liberkeys/api-skeleton/infrastructure/application"
)

func main() {
	fmt.Println("Starting API server...")

	// Dependencies registration into object
	ctx, err := application.NewContext(application.ContextModeServer)
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServer(ctx)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
