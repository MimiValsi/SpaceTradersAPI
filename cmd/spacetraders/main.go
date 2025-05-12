package main

import (
	"context"
	"log"
	"os"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/api"
	"github.com/MimiValsi/SpaceTradersAPI/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	jwt := os.Getenv("JWT")
	if jwt == "" {
		log.Fatalln("Must set JWT")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalln("Must set DB_URL")
	}

	db, err := database.InitDB(dbURL)
	if err != nil {
		log.Fatalln(err)
	}

	c := api.NewClient(context.Background(), db, jwt)
	c.SetToken(jwt)

}
