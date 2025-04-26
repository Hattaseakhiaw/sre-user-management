
func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	// ... healthz route

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)
}
