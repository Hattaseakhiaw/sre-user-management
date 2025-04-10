package db

import (
	"database/sql"
	"fmt"

	"github.com/Hattaseakhiaw/sre-user-management/backend/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return DB.Ping()
}
