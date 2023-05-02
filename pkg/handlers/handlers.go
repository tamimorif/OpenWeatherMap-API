package handlers

import (
	"OpenWeatherMap-API/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{DB: db}
}
func (h *Handler) GetOneWeather(c *gin.Context) {
	//Task1
	// Use the Open Weather API to receive weather data for the given location
	// Используйте Open Weather API для получения данных о погоде для заданного местоположения.
	var city models.City
	result := h.DB.Find(&city)
	if result.Error != nil {
		log.Println("DB error - cannot find City Weather:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, city)
	//Task2
	// Write the weather data to the response
}
