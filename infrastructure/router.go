package infrastructure

import (
	"net/http"
	"weather-api/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(weatherService service.WeatherService) *gin.Engine {
	r := gin.Default()

	r.POST("/weather", func(c *gin.Context) {
		var data struct {
			Water int `json:"water"`
			Wind  int `json:"wind"`
		}

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		weather, status, err := weatherService.UpdateWeather(data.Water, data.Wind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"weather": weather,
			"status":  status,
		})
	})

	return r
}
