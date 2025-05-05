package main

import (
	"fmt"
	"net/http"
)

// Bridge to fetch agent data
// Send un request to "my/agent" api
// The body close is done from the function caller
func clientData(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://api.spacetraders.io/v2/my/agent", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	return client.Do(req)
}
