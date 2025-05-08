package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	addr = ":3000"
)

type application struct {
	token    string
	agent    *Agent
	location *Location
}

func main() {
	godotenv.Load()

	jwt := os.Getenv("JWT")
	if jwt == "" {
		log.Fatalln("Must set JWT")
	}

	app := application{
		token:    jwt,
		agent:    new(Agent),
		location: new(Location),
	}

	server := &http.Server{
		Addr:    addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
