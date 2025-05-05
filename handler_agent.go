package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Agent struct {
	Data struct {
		AccountID       string `json:"accountId,omitempty"`
		Symbol          string `json:"symbol,omitempty"`
		Headquarters    string `json:"headquarters,omitempty"`
		Credits         int64  `json:"credits,omitempty"`
		StartingFaction string `json:"startingFaction,omitempty"`
		ShipCount       int64  `json:"shipCount,omitempty"`
	} `json:"data,omitempty"`
}

// call clientData function to fetch information about the agent
func (c *apiCfg) handlerAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	resp, err := clientData(c.token)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	decoder := json.NewDecoder(resp.Body)
	agent := &Agent{}
	decoder.Decode(agent)

	fmt.Printf("agent: %+v\n", agent)
	w.Write(b)
}
