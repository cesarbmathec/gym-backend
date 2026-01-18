package models

import "time"

type CrossfitPerformance struct {
	ID              uint      `gorm:"primaryKey;column:id_performance" json:"id_performance"`
	ClientID        uint      `gorm:"column:client_id" json:"client_id"`
	WODName         string    `gorm:"size:100;column:wod_name" json:"wod_name"`
	Reps            int       `gorm:"column:reps" json:"reps"`
	TimeSpent       string    `gorm:"type:interval;column:time_spent" json:"time_spent"`
	PerformanceDate time.Time `gorm:"column:performance_date" json:"performance_date"`
	Score           float64   `gorm:"column:score" json:"score"`
}

type BodybuildingPerformance struct {
	ID                uint      `gorm:"primaryKey;column:id_performance" json:"id_performance"`
	ClientID          uint      `gorm:"column:client_id" json:"client_id"`
	Weight            float64   `gorm:"column:weight" json:"weight"`
	BodyFatPercentage float64   `gorm:"column:body_fat_percentage" json:"body_fat_percentage"`
	ArmMeasurement    float64   `gorm:"column:arm_measurement" json:"arm_measurement"`
	MeasurementDate   time.Time `gorm:"column:measurement_date" json:"measurement_date"`
}
