package repository

import (
	"database/sql"

	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, user.Email, user.Password, user.Name)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	row := r.DB.QueryRow(`SELECT id, email, password, name FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
