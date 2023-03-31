package model

// Bottle example
type ReportForm  struct {
	LATITUDE   float64  `json:"latitude" example:"37.55327189658719"`
	LONGTITUDE float64  `json:"longtitude" example:"126.97232991836047"`
	CONTENT string `json:"content"`
}