package models

import "time"

type Card struct {
	CardId          string    `json:"cardId"`
	CreatedAt       time.Time `json:"createdAt"`
	State           string    `json:"state"`
	Title           string    `json:"title"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	SubDistrictCode string    `json:"subDistrictCode"`
	CreateBy        string    `json:"createdBy"`
}
