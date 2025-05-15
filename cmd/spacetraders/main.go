package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/agent"
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

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatalln("Must set DB_URL")
	}

	conn, err := database.InitDB(db_url)
	if err != nil {
		log.Fatalf("couldn't init connection: %s", err)
	}

	agent := agent.Agent{
		Token: jwt,
		Db:    conn,
	}

	client := api.NewClient(context.Background(), agent.Db)
	if err != nil {
		log.Fatalf("New client: %s", err)
	}

	symbol := "MIMIVALSI"
	faction := "AEGIS"
	if err = client.Register(symbol, faction); err != nil {
		log.Fatalf("Couldn't register client")
	}

	fmt.Printf("%+v\n", agent)

}
