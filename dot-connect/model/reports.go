package model

// Bottle example
type Report  struct {
	LATITUDE   float64  `json:"latitude" example:"ex_latitude"`
	LONGTITUDE float64  `json:"longtitude" example:"ex_longtitude"`
	CONTENT string `json:"content"`
}

type PostReport  struct {
	LOCATION    string  `json:"location" example:"ex_location"`
	LOCATIONDETAIL   string  `json:"locationDetail" example:"ex_location_detail"`
	CONTENT string `json:"content"`
}