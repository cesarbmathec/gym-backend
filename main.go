package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/controllers"
	"github.com/cesarbmathec/gym-backend/middleware"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.MigrateTables()

	r := gin.Default()

	// 🔥 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8081", "http://127.0.0.1:5500", "*"}, // Flutter/Frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rutas públicas
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Gym Backend funcionando 🚀"})
	})
	public := r.Group("/public")
	{
		public.GET("/rankings/crossfit", controllers.GetCrossfitRanking)
		public.GET("/news", controllers.GetActiveNews)
		public.POST("/auth/login", controllers.Login)
		public.GET("/ws/:room", controllers.ChatWebSocket)
	}

	// Admin routes (TODO protegido)
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoutes.GET("/clients", controllers.GetAllClients)
		adminRoutes.GET("/clients/:id", controllers.GetClientByID)
		adminRoutes.POST("/clients", controllers.RegisterClient)
		adminRoutes.PUT("/clients/:id", controllers.UpdateClient)
		adminRoutes.DELETE("/clients/:id", controllers.DeleteClient)

		adminRoutes.POST("/trainers", controllers.RegisterTrainer)
		adminRoutes.GET("/trainers", controllers.GetAllTrainers)

		adminRoutes.POST("/subscriptions", controllers.CreateSubscription)
		adminRoutes.GET("/subscriptions", controllers.GetSubscriptions)
		adminRoutes.GET("/subscriptions/overdue", controllers.GetOverdueSubscriptions)

		adminRoutes.POST("/payments", controllers.CreatePayment)
		adminRoutes.GET("/payments", controllers.GetPayments)

		adminRoutes.POST("/checkins", controllers.CreateCheckIn)
		adminRoutes.GET("/checkins/today", controllers.GetTodayCheckIns)
	}

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
