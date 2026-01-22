package controllers

import (
	"net/http"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Si client_id es null, se considera una noticia general para todos
	config.DB.Create(&notification)
	c.JSON(http.StatusCreated, gin.H{"message": "Noticia publicada correctamente"})
}

func GetActiveNews(c *gin.Context) {
	var news []models.Notification
	// Buscamos las notificaciones tipo 'event' o 'payment_change'
	config.DB.Where("notification_type IN ?", []string{"event", "payment_change", "competition"}).Find(&news)
	c.JSON(http.StatusOK, news)
}
