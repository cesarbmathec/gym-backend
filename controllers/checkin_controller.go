package controllers

import (
	"net/http"
	"time"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
)

type CheckInInput struct {
	ClientID  uint  `json:"client_id" binding:"required"`
	TrainerID *uint `json:"trainer_id"`
}

func CreateCheckIn(c *gin.Context) {
	var input CheckInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkin := models.CheckIn{
		ClientID:    input.ClientID,
		CheckInTime: time.Now(),
		TrainerID:   input.TrainerID,
	}

	config.DB.Create(&checkin)
	c.JSON(http.StatusCreated, gin.H{"message": "Check-in registered", "data": checkin})
}

func GetTodayCheckIns(c *gin.Context) {
	var checkins []models.CheckIn
	config.DB.Where("DATE(check_in_time) = CURRENT_DATE").Find(&checkins)
	c.JSON(http.StatusOK, gin.H{"data": checkins})
}
