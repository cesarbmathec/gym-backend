package controllers

import (
	"net/http"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
)

// Para rankings públicos
func GetCrossfitRanking(c *gin.Context) {
	var rankings []models.CrossfitPerformance
	config.DB.
		Select("client_id, AVG(score) as avg_score").
		Where("score IS NOT NULL").
		Group("client_id").
		Order("avg_score DESC").
		Limit(10).
		Find(&rankings)
	c.JSON(http.StatusOK, gin.H{"data": rankings})
}
