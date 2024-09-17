package service

import (
	"weather-api/models"
	"weather-api/repository"
)

type WeatherService interface {
	UpdateWeather(water, wind int) (*models.Weather, *models.WeatherStatus, error)
}

type weatherService struct {
	repo repository.WeatherRepository
}

func NewWeatherService(repo repository.WeatherRepository) WeatherService {
	return &weatherService{repo: repo}
}

func (s *weatherService) UpdateWeather(water, wind int) (*models.Weather, *models.WeatherStatus, error) {
	weather := &models.Weather{
		Water: water,
		Wind:  wind,
	}

	err := s.repo.UpdateWeather(weather)
	if err != nil {
		return nil, nil, err
	}

	status := &models.WeatherStatus{
		WaterStatus: getWaterStatus(water),
		WindStatus:  getWindStatus(wind),
	}

	return weather, status, nil
}

func getWaterStatus(water int) string {
	if water < 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	}
	return "Bahaya"
}

func getWindStatus(wind int) string {
	if wind < 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	}
	return "Bahaya"
}
