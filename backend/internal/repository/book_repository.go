package repository

import (
	"database/sql"

	"github.com/ramzyrsr/digital-library/internal/entity"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) CreateBook(book *entity.Book) error {
	query := `INSERT INTO books (title, author, isbn, quantity, category_id, created_by)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.Exec(query, book.Title, book.Author, book.ISBN, book.Quantity, book.CategoryID, book.CreatedBy)

	return err
}
