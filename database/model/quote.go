package model

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type Quotes struct {
	Author string   `json:"author"`
	Quotes []string `json:"quotes"`
}
