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

type apiCfg struct {
	token string
}

func main() {
	godotenv.Load()

	jwt := os.Getenv("JWT")
	if jwt == "" {
		log.Fatalln("Must set JWT")
	}

	c := apiCfg{
		token: jwt,
	}

	server := &http.Server{
		Addr:    addr,
		Handler: c.routes(),
	}

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
