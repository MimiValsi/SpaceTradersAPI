package agent

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/database"
	"github.com/MimiValsi/SpaceTradersAPI/pkg/model"
)

type Agent struct {
	Token string
	db    *database.DB
}

// The client side of this app will be written here.
// Still need to find a good convention for the function names...

func FetchAgentData(c *Agent) (*model.Agent, error) {
	req, err := http.NewRequest("GET", "https://api.spacetraders.io/v2/my/agent", nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch agent data: %w", err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	agent := model.Agent{}
	err = decoder.Decode(&agent)
	if err != nil {
		return nil, fmt.Errorf("failed to decode agent: %w", err)
	}

	return &agent, nil
}

func FetchWaypointLocation(waypointSymbol string) (*http.Response, error) {

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
