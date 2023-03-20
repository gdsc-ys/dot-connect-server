package model

// Bottle example
type Report  struct {
	ID      int     `json:"id" example:"1"`
	LOCATION    string  `json:"location" example:"ex_location"`
	CONTENT string `json:"content"`
	IMAGE string `json:"image"`
}

type PostReport  struct {
	LOCATION    string  `json:"location" example:"ex_location"`
	LOCATIONLOCATION    string  `json:"locationDetail" example:"ex_location_detail"`
	CONTENT string `json:"content"`
}