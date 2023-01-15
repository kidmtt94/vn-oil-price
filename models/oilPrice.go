package models

// OilPrice model
type OilPrice struct {
	Name         string `json:"name"`
	AreaOnePrice string `json:"area-1"`
	AreaTwoPrice string `json:"area-2"`
}
