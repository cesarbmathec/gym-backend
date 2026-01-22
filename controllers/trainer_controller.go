package controllers

import (
	"net/http"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type TrainerInput struct {
	IdentityCard string `json:"identity_card" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required,min=6"`
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	TrainerType  string `json:"trainer_type"`
}

func RegisterTrainer(c *gin.Context) {
	var input TrainerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	trainer := models.Trainer{
		IdentityCard: input.IdentityCard,
		Username:     input.Username,
		PasswordHash: string(bytes),
		FullName:     input.FullName,
		Email:        input.Email,
		TrainerType:  input.TrainerType,
	}

	config.DB.Create(&trainer)
	c.JSON(http.StatusCreated, gin.H{"data": trainer})
}

func GetAllTrainers(c *gin.Context) {
	var trainers []models.Trainer
	config.DB.Find(&trainers)
	c.JSON(http.StatusOK, gin.H{"data": trainers})
}

func GetTrainerByID(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer
	if err := config.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": trainer})
}

func UpdateTrainer(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer
	if err := config.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}

	var input TrainerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTrainer := models.Trainer{
		IdentityCard: input.IdentityCard,
		Username:     input.Username,
		FullName:     input.FullName,
		Email:        input.Email,
		TrainerType:  input.TrainerType,
	}

	config.DB.Model(&trainer).Updates(updatedTrainer)
	c.JSON(http.StatusOK, gin.H{"data": trainer})
}

func DeleteTrainer(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer
	if err := config.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}

	config.DB.Delete(&trainer)
	c.JSON(http.StatusOK, gin.H{"message": "Trainer deleted"})
}
