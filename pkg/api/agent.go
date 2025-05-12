package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/model"
)

func (c *Client) MyAgent() error {
	url := url.URL{Path: "my/agent"}

	uri := c.BaseURI.ResolveReference(&url)
	fmt.Printf("uri: %s\n", uri.String())
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return err
	}

	req.Header = *c.Header
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("header: %v\n", req.Header)

	decoder := json.NewDecoder(resp.Body)
	agent := model.Agent{}
	if err = decoder.Decode(&agent); err != nil {
		return err
	}

	fmt.Printf("%+v\n", agent)

	return nil
}
