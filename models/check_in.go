package models

import "time"

type CheckIn struct {
	ID          uint      `gorm:"primaryKey;column:id_check_in" json:"id_check_in"`
	ClientID    uint      `gorm:"column:client_id" json:"client_id"`
	CheckInTime time.Time `gorm:"column:check_in_time" json:"check_in_time"`
	TrainerID   *uint     `gorm:"column:trainer_id" json:"trainer_id"`
}
