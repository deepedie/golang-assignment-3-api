package repository

import (
	"weather-api/models"

	"gorm.io/gorm"
)

type WeatherRepository interface {
	UpdateWeather(weather *models.Weather) error
}

type weatherRepository struct {
	db *gorm.DB
}

func NewWeatherRepository(db *gorm.DB) WeatherRepository {
	return &weatherRepository{db: db}
}

func (r *weatherRepository) UpdateWeather(weather *models.Weather) error {
	return r.db.Create(weather).Error
}
