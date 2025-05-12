package model

type Agent struct {
	Data struct {
		AccountID       string `json:"accountId,omitempty"`
		Symbol          string `json:"symbol"`
		Headquarters    string `json:"headquarters"`
		Credits         int64  `json:"credits"`
		StartingFaction string `json:"startingFaction"`
		ShipCount       int64  `json:"shipCount"`
	} `json:"data"`
}
