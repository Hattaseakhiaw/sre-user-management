package repository

import (
	"github.com/Hattaseakhiaw/sre-user-management/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password) 
			  VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	return r.db.QueryRowx(query, user.Username, user.Email, user.Password).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	)
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password, created_at, updated_at
			  FROM users WHERE email = $1`

	err := r.db.Get(user, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
