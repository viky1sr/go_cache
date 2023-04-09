package migrations

import (
	"database/sql"
)

// Up function for users table migration
func UpUsers(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE users (
			id bigint identity PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL UNIQUE,
			password VARCHAR(100) NOT NULL
		);`)
	if err != nil {
		return err
	}

	return nil
}

// Down function for users table migration
func DownUsers(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return err
	}

	return nil
}
