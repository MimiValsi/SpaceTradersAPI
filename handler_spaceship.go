package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Location struct {
	Data struct {
		SystemSymbol string `json:"systemSymbol,omitempty"`
		Symbol       string `json:"symbol,omitempty"`
		Type         string `json:"type,omitempty"`
		X            int32  `json:"x,omitempty"`
		Y            int32  `json:"y,omitempty"`

		Orbitals []struct {
			Symbol string `json:"symbol,omitempty"`
		} `json:"orbitals,omitempty"`

		Traits []struct {
			Symbol      string `json:"symbol,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"traits,omitempty"`

		Modifiers []struct {
			Symbol      string `json:"symbol,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		}

		Chart []struct {
			SubmittedBy string    `json:"submittedBy,omitempty"`
			SubmittedOn time.Time `json:"submittedOn,omitempty"`
		} `json:"chart,omitempty"`

		Faction []struct {
			Symbol string `json:"symbol,omitempty"`
		} `json:"faction,omitempty"`

		IsUnderConstruction bool `json:"isUnderConstruction,omitempty"`
	} `json:"data,omitempty"`
}

func (app *application) handlerSpaceshipLocation(w http.ResponseWriter, r *http.Request) {
	resp, err := fetchWaypointLocation(app.agent.Data.Headquarters)
	if err != nil {
		log.Printf("Fetch waypoint location error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(app.location)
	if err != nil {
		log.Printf("json location decoder: %s\n", err)
		return
	}

	fmt.Printf("location: %+v\n", app.location)
}
