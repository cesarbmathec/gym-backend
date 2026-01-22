package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ClientInput struct {
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

// CRUD Clientes
func RegisterClient(c *gin.Context) {
	var input ClientInput
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Client created", "data": client})
}

func GetAllClients(c *gin.Context) {
	var clients []models.Client
	config.DB.Find(&clients)
	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func GetClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func UpdateClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	var input ClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&client).Updates(models.Client{
		FullName:          input.FullName,
		Email:             input.Email,
		Phone:             input.Phone,
		Weight:            input.Weight,
		Height:            input.Height,
		BloodType:         input.BloodType,
		MedicalConditions: input.MedicalConditions,
		IsActive:          input.Weight > 0, // Auto-active si tiene peso
	})
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func DeleteClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.DB.Delete(&models.Client{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted"})
}

func GetClientByID(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func parseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, dateStr)
	return t
}
