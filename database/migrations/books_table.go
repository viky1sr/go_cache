package migrations

import (
	"database/sql"
)

// Up function for books table migration
func UpBooks(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE books (
			id bigint identity PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			author VARCHAR(50) NOT NULL,
			year INTEGER NOT NULL
		);`)
	if err != nil {
		return err
	}

	return nil
}

// Down function for books table migration
func DownBooks(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS books;`)
	if err != nil {
		return err
	}

	return nil
}
