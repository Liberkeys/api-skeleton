package main

import (
	"fmt"
	"log"

	"github.com/Liberkeys/api-skeleton/adapters/http"
	"github.com/Liberkeys/api-skeleton/infrastructure/app"
)

func main() {
	fmt.Println("Starting API server...")

	// Dependencies registration into object
	ctx, err := app.NewContext(app.ContextModeServer)
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServer(ctx)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
