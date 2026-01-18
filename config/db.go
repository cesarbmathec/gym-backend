package config

import (
	"log"

	"github.com/cesarbmathec/gym-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=gym_user password=123456 dbname=gym_backend port=5432 sslmode=disable TimeZone=America/Caracas"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	DB = database
	log.Println("✅ Conexión a PostgreSQL exitosa")
}

func MigrateTables() {
	modelsList := []interface{}{
		&models.Client{},
		&models.Trainer{},
		&models.Subscription{},
		&models.Payment{},
		&models.CheckIn{},
		&models.CrossfitPerformance{},
		&models.BodybuildingPerformance{},
		&models.Chat{},
		&models.ChatParticipant{},
		&models.ChatMessage{},
		&models.Notification{},
	}

	for _, model := range modelsList {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Printf("⚠️ Error migrando modelo %T: %v", model, err)
		}
	}
	log.Println("✅ TODAS las tablas migradas correctamente")
}
