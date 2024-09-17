package models

import "time"

type Weather struct {
	ID        uint      `json:"id"`
	Water     int       `json:"water"`
	Wind      int       `json:"wind"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WeatherStatus struct {
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}
