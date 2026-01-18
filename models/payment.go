package models

import "time"

type Payment struct {
	ID             uint      `gorm:"primaryKey;column:id_payment" json:"id_payment"`
	SubscriptionID uint      `gorm:"column:subscription_id" json:"subscription_id"`
	Amount         float64   `gorm:"column:amount" json:"amount"`
	PaymentDate    time.Time `gorm:"column:payment_date" json:"payment_date"`
	PaymentMethod  string    `gorm:"size:20;column:payment_method" json:"payment_method"`
}
