package database

import (
	"context"
	"fmt"

	"github.com/MimiValsi/SpaceTradersAPI/pkg/model"
)

// AccountID       string `json:"accountId,omitempty"`
// Symbol          string `json:"symbol"`
// Headquarters    string `json:"headquarters"`
// Credits         int64  `json:"credits"`
// StartingFaction string `json:"startingFaction"`
// ShipCount       int64  `json:"shipCount"`
func (db *DB) SaveAgent(agent *model.Agent) error {
	sql := `
INSERT INTO agent(account_id, symbol, headquarters, credits, starting_faction, ship_count) VALUES($1, $2, $3, $4, $5, $6);
	`
	_, err := db.conn.Exec(context.Background(), sql, agent.Data.AccountID, agent.Data.Symbol, agent.Data.Headquarters, agent.Data.Credits, agent.Data.StartingFaction, agent.Data.ShipCount)
	if err != nil {
		return fmt.Errorf("couldn't insert agent data: %w", err)
	}

	return nil
}
