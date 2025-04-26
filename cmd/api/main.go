package main

import (
	"log"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	handler "github.com/Hattaseakhiaw/sre-user-management/backend/internal/handlers"
	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/repository"
	service "github.com/Hattaseakhiaw/sre-user-management/backend/internal/services"
	"github.com/Hattaseakhiaw/sre-user-management/backend/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadConfig()

	// Connect to DB
	if err := db.ConnectPostgres(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	// Create Repository, Service, Handler
	userRepo := repository.NewUserRepository(db.DB)
	authService := service.NewAuthService(userRepo, []byte("your-secret-key"))
	authHandler := handler.NewAuthHandler(authService)

	// Setup Gin
	r := gin.Default()

	// Routes
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Auth Routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
