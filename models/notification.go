package models

import "time"

type Notification struct {
	ID               uint      `gorm:"primaryKey;column:id_notification" json:"id_notification"`
	Title            string    `gorm:"not null;size:200;column:title" json:"title"`
	Message          string    `gorm:"type:text;column:message" json:"message"`
	NotificationType string    `gorm:"size:30;column:notification_type" json:"notification_type"`
	ClientID         *uint     `gorm:"column:client_id" json:"client_id"`
	SentAt           time.Time `gorm:"column:sent_at" json:"sent_at"`
	IsRead           bool      `gorm:"default:false;column:is_read" json:"is_read"`
}
