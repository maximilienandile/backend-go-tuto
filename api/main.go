package main

import (
	"log"

	"github.com/maximilienandile/backend-go-tuto/internal/server"
)

func main() {
	myServer, err := server.New(server.Config{
		Port: 9090,
	})
	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}
	err = myServer.Run()
	if err != nil {
		log.Fatalf("impossible to start the server: %s", err)
	}
}
