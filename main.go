package main

import (
	"weather-api/infrastructure"
	"weather-api/repository"
	"weather-api/service"
)

func main() {
	db := infrastructure.NewDatabase()
	weatherRepo := repository.NewWeatherRepository(db)
	weatherService := service.NewWeatherService(weatherRepo)
	r := infrastructure.SetupRouter(weatherService)

	r.Run(":8081")
}
