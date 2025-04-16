var DB *sqlx.DB

func ConnectPostgres() error {
	cfg := config.AppConfig

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetDB() *sqlx.DB {
	return DB
}
