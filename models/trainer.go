package models

type Trainer struct {
	ID           uint   `gorm:"primaryKey;column:id_trainer" json:"id_trainer"`
	IdentityCard string `gorm:"unique;not null;size:15;column:identity_card" json:"identity_card"`
	Username     string `gorm:"unique;not null;size:50;column:username" json:"username"`
	PasswordHash string `gorm:"not null;size:255;column:password_hash" json:"-"`
	FullName     string `gorm:"not null;size:100;column:full_name" json:"full_name"`
	Email        string `gorm:"unique;not null;size:100;column:email" json:"email"`
	Phone        string `gorm:"size:20;column:phone" json:"phone"`
	TrainerType  string `gorm:"size:20;default:general;column:trainer_type" json:"trainer_type"`
	IsActive     bool   `gorm:"default:true;column:is_active" json:"is_active"`
}
