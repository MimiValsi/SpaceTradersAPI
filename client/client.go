package client

import (
	"fmt"
	"net/http"
)

// The client side of this app will be written here.
// Still need to find a good convention for the function names...

// Basic function that send a request to api spacetraders
// that fetch the agent data and return a *http.Response.
func FetchClientData(token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://api.spacetraders.io/v2/my/agent", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	return client.Do(req)
}
