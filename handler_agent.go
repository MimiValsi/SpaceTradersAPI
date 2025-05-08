package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Agent struct {
	Data struct {
		AccountID       string `json:"accountId,omitempty"`
		Symbol          string `json:"symbol"`
		Headquarters    string `json:"headquarters"`
		Credits         int64  `json:"credits"`
		StartingFaction string `json:"startingFaction"`
		ShipCount       int64  `json:"shipCount"`
	} `json:"data"`
}

// call clientData function to fetch information about the agent
func (app *application) handlerFetchAgentData(w http.ResponseWriter, r *http.Request) {
	resp, err := fetchAgentData(app.token)
	if err != nil {
		log.Printf("Fetch agent data error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&app.agent)
	if err != nil {
		log.Printf("json agent decoder: %s\n", err)
		return
	}

	fmt.Printf("agent: %+v\n", app.agent)
}
