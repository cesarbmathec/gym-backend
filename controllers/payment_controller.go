package controllers

import (
	"net/http"
	"time"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
)

type PaymentInput struct {
	SubscriptionID uint    `json:"subscription_id" binding:"required"`
	Amount         float64 `json:"amount" binding:"required"`
	PaymentMethod  string  `json:"payment_method" binding:"required"`
}

func CreatePayment(c *gin.Context) {
	var input PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := models.Payment{
		SubscriptionID: input.SubscriptionID,
		Amount:         input.Amount,
		PaymentDate:    time.Now(),
		PaymentMethod:  input.PaymentMethod,
	}

	// Actualizar suscripción a activa
	config.DB.Model(&models.Subscription{}).
		Where("id_subscription = ?", input.SubscriptionID).
		Update("status", "active")

	config.DB.Create(&payment)
	c.JSON(http.StatusCreated, gin.H{"message": "Payment registered", "data": payment})
}

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	config.DB.Find(&payments)
	c.JSON(http.StatusOK, gin.H{"data": payments})
}
