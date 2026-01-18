package models

import (
	"time"
)

type Client struct {
	ID                uint      `gorm:"primaryKey;column:id" json:"id"`
	IdentityCard      string    `gorm:"unique;not null;size:15;column:identity_card" json:"identity_card"`
	Username          string    `gorm:"unique;not null;size:50;column:username" json:"username"`
	PasswordHash      string    `gorm:"not null;size:255;column:password_hash" json:"-"`
	FullName          string    `gorm:"not null;size:100;column:full_name" json:"full_name"`
	Email             string    `gorm:"unique;not null;size:100;column:email" json:"email"`
	Phone             string    `gorm:"size:20;column:phone" json:"phone"`
	Weight            float64   `gorm:"column:weight" json:"weight"`
	Height            float64   `gorm:"column:height" json:"height"`
	BloodType         string    `gorm:"size:5;column:blood_type" json:"blood_type"`
	MedicalConditions string    `gorm:"type:text;column:medical_conditions" json:"medical_conditions"`
	BirthDate         time.Time `gorm:"column:birth_date" json:"birth_date"`
	IsActive          bool      `gorm:"default:true;column:is_active" json:"is_active"`
	IsAdmin           bool      `gorm:"default:false;column:is_admin" json:"is_admin"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
}
