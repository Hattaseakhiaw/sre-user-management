package db

import (
	"fmt"
	"log"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectPostgres() error {
	conf := config.AppConfig

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName,
	)

	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}

	log.Println("✅ Connected to PostgreSQL")
	return nil
}

func GetDB() *sqlx.DB {
	return DB
}
