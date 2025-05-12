package api

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/database"
)

type Client struct {
	BaseURI    *url.URL
	db         *database.DB
	Header     *http.Header
	httpClient *http.Client
}

func NewClient(ctx context.Context, db *database.DB, token string) *Client {
	baseURI, err := url.Parse("https://api.spacetraders.io/v2/")
	if err != nil {
		panic("failed to parse url")
	}
	client := &Client{
		BaseURI: baseURI,
		db:      db,
		Header: &http.Header{
			"Content-Type": {"application/json"},
		},
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	return client
}

func (c *Client) SetToken(token string) {
	c.Header.Set("Authorization", "Bearer "+token)
}
