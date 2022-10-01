package models

// OilPrice model
type OilPrice struct {
	Name         string `json:"name"`
	AreaOnePrice string `json:"areaOnePrice"`
	AreaTwoPrice string `json:"areaTwoPrice"`
}
