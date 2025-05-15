package model

type Data struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
	Email   string `json:"email"`
}

type Player struct {
	Data struct {
		Token    string   `json:"token"`
		Agent    Agent    `json:"agent"`
		Faction  Faction  `json:"faction"`
		Contract Contract `json:"contract"`
		Ships    []Ship   `json:"ships"`
	} `json:"data"`
}
