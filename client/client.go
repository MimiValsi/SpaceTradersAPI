package client

import "net/http"

// Client package that check for reset updates, ...

func CheckAgent(token string) error {
	req, err := http.NewRequest("GET", "https://api.spacetraders.io/v2/my/agent")
	if err != nil {
		return err
	}

	req.Header.Add("Authorization:", "Bearer "+token)

}
