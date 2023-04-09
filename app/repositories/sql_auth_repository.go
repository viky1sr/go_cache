package repositories

import (
	"database/sql"

	"github.com/viky1sr/go_cache.git/app/models"
)

type SQLAuthRepository struct {
	Db *sql.DB
}

func NewSQLAuthRepository(db *sql.DB) *SQLAuthRepository {
	return &SQLAuthRepository{
		Db: db,
	}
}

func (r *SQLAuthRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password FROM users WHERE email = ?"

	err := r.Db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *SQLAuthRepository) CreateToken(userID int64, token string) error {
	query := "INSERT INTO user_tokens (user_id, token) VALUES (?, ?)"

	_, err := r.Db.Exec(query, userID, token)
	if err != nil {
		return err
	}

	return nil
}
