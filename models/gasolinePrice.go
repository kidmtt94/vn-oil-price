package models

// Gasoline price model
type GasolinePrice struct {
	Location     string `json:"location"`
	Cylinder12KG string `json:"12-kg-cylinder"`
	Cylinder48KG string `json:"48-kg-cylinder"`
}
