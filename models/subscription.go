package models

import "time"

type Subscription struct {
	ID                uint      `gorm:"primaryKey;column:id_subscription" json:"id_subscription"`
	ClientID          uint      `gorm:"column:client_id" json:"client_id"`
	TrainingType      string    `gorm:"size:20;column:training_type" json:"training_type"`
	PersonalTrainerID *uint     `gorm:"column:personal_trainer_id" json:"personal_trainer_id"`
	StartDate         time.Time `gorm:"column:start_date" json:"start_date"`
	ExpiryDate        time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	MonthlyFee        float64   `gorm:"column:monthly_fee" json:"monthly_fee"`
	Status            string    `gorm:"default:active;column:status" json:"status"`
}
