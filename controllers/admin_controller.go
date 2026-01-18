package controllers

import (
	"net/http"
	"strconv"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"

	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetAllClients(c *gin.Context) {
	var clients []models.Client
	if err := config.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func GetClientByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}

func UpdateClient(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	var input models.Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&client).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func DeleteClient(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	if err := config.DB.Delete(&models.Client{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting client"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

func RegisterClient(c *gin.Context) {
	var input struct {
		IdentityCard      string  `json:"identity_card" binding:"required"`
		Username          string  `json:"username" binding:"required"`
		Password          string  `json:"password" binding:"required,min=6"`
		FullName          string  `json:"full_name" binding:"required"`
		Email             string  `json:"email" binding:"required,email"`
		Phone             string  `json:"phone"`
		Weight            float64 `json:"weight"`
		Height            float64 `json:"height"`
		BloodType         string  `json:"blood_type"`
		MedicalConditions string  `json:"medical_conditions"`
		BirthDate         string  `json:"birth_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create client
	client := models.Client{
		IdentityCard:      input.IdentityCard,
		Username:          input.Username,
		PasswordHash:      string(bytes),
		FullName:          input.FullName,
		Email:             input.Email,
		Phone:             input.Phone,
		Weight:            input.Weight,
		Height:            input.Height,
		BloodType:         input.BloodType,
		MedicalConditions: input.MedicalConditions,
		BirthDate:         parseDate(input.BirthDate),
	}

	if err := config.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client already exists or database error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Client registered successfully",
		"client":  client,
	})
}

func parseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
