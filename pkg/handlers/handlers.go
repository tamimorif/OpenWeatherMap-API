package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{DB: db}
}
func (h *Handler) GetWeather(c *gin.Context) {
	//Task1
	// Use the Open Weather API to receive weather data for the given location
	//var cities []models.City
	//result := h.DB.Find(&cities)
	//if result.Error != nil {
	//	log.Println("DB error - cannot find tasks:", result.Error)
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "Internal Server Error",
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, cities)
	//Task2
	// Write the weather data to the response
}
