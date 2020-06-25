package models

import "time"

type User struct {
	User            string    `json:"user"`
	CreatedAt       time.Time `json:"createdAt"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	SubDistrictCode string    `json:"subDistrictCode"`
}
