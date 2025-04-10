package main

import (
	"fmt"
	"log"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	"github.com/Hattaseakhiaw/sre-user-management/backend/pkg/db"
)

func main() {
	config.LoadConfig()
	if err := db.ConnectPostgres(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")
}
