package main

import (
	"fmt"
	"net/http"
)

// The client side of this app will be written here.
// Still need to find a good convention for the function names...

// Basic function that send a request to api spacetraders
// that fetch the agent data and return a *http.Response.
func fetchAgentData(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://api.spacetraders.io/v2/my/agent", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	return client.Do(req)
}

func fetchWaypointLocation(waypointSymbol string) (*http.Response, error) {

	// sector -> X1
	// systemSymbol -> X1-DF55
	// waypointSymbol -> X1-DF55-A1

	// 1st param -> :systemSymbol
	// 2nd param -> :waypointSymbol
	url := fmt.Sprintf("https://api.spacetraders.io/v2/systems/%s/waypoints/%s", waypointSymbol[:7], waypointSymbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	return client.Do(req)
}
