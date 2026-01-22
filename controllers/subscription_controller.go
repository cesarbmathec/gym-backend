package controllers

import (
	"net/http"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
)

type SubscriptionInput struct {
	ClientID          uint    `json:"client_id" binding:"required"`
	TrainingType      string  `json:"training_type" binding:"required"`
	PersonalTrainerID *uint   `json:"personal_trainer_id"`
	MonthlyFee        float64 `json:"monthly_fee" binding:"required"`
	StartDate         string  `json:"start_date" binding:"required"`
}

func CreateSubscription(c *gin.Context) {
	var input SubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscription := models.Subscription{
		ClientID:          input.ClientID,
		TrainingType:      input.TrainingType,
		PersonalTrainerID: input.PersonalTrainerID,
		StartDate:         parseDate(input.StartDate),
		ExpiryDate:        parseDate(input.StartDate).AddDate(0, 1, 0), // 1 mes
		MonthlyFee:        input.MonthlyFee,
		Status:            "active",
	}

	if err := config.DB.Create(&subscription).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating subscription"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": subscription})
}

func GetSubscriptions(c *gin.Context) {
	var subs []models.Subscription
	config.DB.Find(&subs)
	c.JSON(http.StatusOK, gin.H{"data": subs})
}

func GetOverdueSubscriptions(c *gin.Context) {
	var overdue []models.Subscription
	config.DB.Where("expiry_date <= NOW() AND status = ?", "active").Find(&overdue)
	c.JSON(http.StatusOK, gin.H{"data": overdue})
}
