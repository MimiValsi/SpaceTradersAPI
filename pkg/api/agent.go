package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/model"
)

func (c *Client) Register(symbol, faction string) error {
	url := url.URL{Path: "/register"}
	uri := c.BaseURI.ResolveReference(&url)

	// c.Header.Set("Content-Type", "application/json")

	// c.Data = &model.Data{
	// 	Symbol:  "MIMIVALSI",
	// 	Faction: "AEGIS",
	// 	Email:   "miguel@dasilvaf.net",
	// }
	body := fmt.Sprintf("{\"symbol\": \"%s\", \"faction\": \"%s\"}", symbol, faction)

	req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	// req.Header = *c.Header
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	agent := model.Player{}
	if err = decoder.Decode(&agent); err != nil {
		return err
	}

	fmt.Printf("%+v\n", agent)

	return nil
}

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
