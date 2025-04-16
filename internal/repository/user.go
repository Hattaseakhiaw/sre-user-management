package repository

import (
	"database/sql"
	"time"

	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(username, hashedPassword string) (int64, error) {
	query := "INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	err := r.DB.QueryRow(query, username, hashedPassword, time.Now()).Scan(&id)
	return id, err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, created_at FROM users WHERE username = $1"
	row := r.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
