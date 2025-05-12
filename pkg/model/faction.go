package model

type Faction struct {
	Data struct {
		Symbol       string   `json:"symbol"`
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		Headquarters string   `json:"headquartes"`
		Traits       []Traits `json:"traits"`
		IsRecruiting bool     `json:"isRecruiting"`
	} `json:"data"`
}

type Traits struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
