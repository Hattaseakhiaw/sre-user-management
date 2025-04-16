package main

import (
	"fmt"
	"log"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	"github.com/Hattaseakhiaw/sre-user-management/backend/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load ENV config
	config.LoadConfig()

	// Connect to DB
	if err := db.ConnectPostgres(); err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")

	// Start Gin HTTP server
	router := gin.Default()

	// Health check endpoint
	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
