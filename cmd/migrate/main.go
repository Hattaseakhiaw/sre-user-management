package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	"github.com/Hattaseakhiaw/sre-user-management/backend/pkg/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()

	err := db.ConnectPostgres()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	conn := db.GetDB() // ต้องมีฟังก์ชันนี้ใน pkg/db เพื่อ return *sqlx.DB
	applyMigrations(conn)
}

func applyMigrations(db *sqlx.DB) {
	migrationFile := "migrations/001_create_users_table.up.sql"

	content, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("❌ Failed to read migration file: %v", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatalf("❌ Failed to apply migration: %v", err)
	}

	fmt.Println("✅ Migration applied successfully!")
}
