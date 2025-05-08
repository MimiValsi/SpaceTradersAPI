package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Waypoint struct {
	Data struct {
		Symbol       string `json:"symbol"`
		Type         string `json:"type"`
		SystemSymbol string `json:"systemSymbol"`
		X            int32  `json:"x"`
		Y            int32  `json:"y"`

		Orbitals []struct {
			Symbol string `json:"symbol"`
		} `json:"orbitals"`

		Orbits string `json:"orbits,omitempty"` // optional

		Faction []struct {
			Symbol string `json:"symbol"`
		} `json:"faction"`

		Traits []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"traits"`

		Modifiers []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"modifiers"`

		Chart []struct {
			WaypointSymbol string    `json:"waypointSymbol"`
			SubmittedBy    string    `json:"submittedBy,omitempty"`
			SubmittedOn    time.Time `json:"submittedOn,omitempty"`
		} `json:"chart,omitempty"` // optional

		IsUnderConstruction bool `json:"isUnderConstruction"` // optional
	} `json:"data"`
}

func (app *application) handlerSpaceshipLocation(w http.ResponseWriter, r *http.Request) {
	resp, err := fetchWaypointLocation(app.agent.Data.Headquarters)
	if err != nil {
		log.Printf("Fetch waypoint location error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&app.location)
	if err != nil {
		log.Printf("json location decoder: %s\n", err)
		return
	}

	fmt.Printf("location: %+v\n", app.location)
}
