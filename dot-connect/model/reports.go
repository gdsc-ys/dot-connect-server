package model

// Bottle example
type Report  struct {
	Latitude   float64  `json:"latitude" example:"ex_latitude"`
	Longtitude float64  `json:"longtitude" example:"ex_longtitude"`
	CONTENT string `json:"content"`

	// IMAGE string `json:"image"`
}

type PostReport  struct {
	LOCATION    string  `json:"location" example:"ex_location"`
	LOCATIONDETAIL   string  `json:"locationDetail" example:"ex_location_detail"`
	CONTENT string `json:"content"`
}