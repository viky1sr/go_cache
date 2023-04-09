package providers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/viky1sr/go_cache.git/app/repositories"
)

type SqlProvider struct {
	db *sql.DB
}

func NewSqlProvider() (*SqlProvider, error) {
	// Buat koneksi ke database
	db, err := sql.Open("mssql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Cek koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to database")

	return &SqlProvider{
		db: db,
	}, nil
}

func (p *SqlProvider) Close() {
	p.db.Close()
}

func (p *SqlProvider) BookRepository() repositories.BookRepository {
	return &repositories.SqlBookRepository{Db: p.db}
}

func (p *SqlProvider) UserRepository() repositories.UserRepository {
	return &repositories.SqlUserRepository{Db: p.db}
}
