package repositories

import (
	"database/sql"
	"github.com/viky1sr/go_cache.git/app/models"
)

type SqlBookRepository struct {
	Db *sql.DB
}

func NewSqlBookRepository(db *sql.DB) *SqlBookRepository {
	return &SqlBookRepository{
		Db: db,
	}
}

func (repo *SqlBookRepository) GetAllBooks() ([]*models.Book, error) {
	books := make([]*models.Book, 0)

	rows, err := repo.Db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(models.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *SqlBookRepository) GetBookByID(id int) (*models.Book, error) {
	book := new(models.Book)

	row := repo.Db.QueryRow("SELECT * FROM books WHERE id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (repo *SqlBookRepository) CreateBook(book *models.Book) error {
	_, err := repo.Db.Exec("INSERT INTO books (title, author, year) VALUES ($1, $2, $3)",
		book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SqlBookRepository) UpdateBook(id int, book *models.Book) error {
	_, err := repo.Db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4",
		book.Title, book.Author, book.Year, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SqlBookRepository) DeleteBook(id int) error {
	_, err := repo.Db.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
