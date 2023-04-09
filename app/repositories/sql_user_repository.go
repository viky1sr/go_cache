package repositories

import (
	"database/sql"
	"errors"
	"github.com/viky1sr/go_cache.git/app/models"
)

type SqlUserRepository struct {
	Db *sql.DB
}

func NewSqlUserRepository(db *sql.DB) UserRepository {
	return &SqlUserRepository{
		Db: db,
	}
}

func (repo *SqlUserRepository) GetAllUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)

	rows, err := repo.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *SqlUserRepository) GetUserByID(id int) (*models.User, error) {
	user := new(models.User)

	row := repo.Db.QueryRow("SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return user, nil
}

func (repo *SqlUserRepository) FindByEmail(email string) (bool, error) {
	user := new(models.User)

	row := repo.Db.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	// Jika email sudah ada di database, return false dan pesan error
	return true, errors.New("email already registered")
}

func (repo *SqlUserRepository) CreateUser(user *models.User) error {
	tx, err := repo.Db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO users (name, email, password) OUTPUT INSERTED.id VALUES (?, ?, ?);`
	var id uint
	err = tx.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	user.ID = id

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (repo *SqlUserRepository) UpdateUser(id int, user *models.User) error {
	query := `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`

	_, err := repo.Db.Exec(query, user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SqlUserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := repo.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
