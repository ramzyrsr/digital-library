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
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.DB.QueryRow(query, book.Title, book.Author, book.ISBN, book.Quantity, book.CategoryID, book.CreatedBy).Scan(&book.ID)
	if err != nil {
		return err
	}

	query = `INSERT INTO book_status (book_id, available_qty, borrowed_qty)
			 VALUES ($1, $2, $3)`
	_, err = r.DB.Exec(query, book.ID, book.Quantity, 0)

	return err
}

func (r *BookRepository) GetBooks(limit, offset int) ([]entity.Book, int, int, error) {
	query := `SELECT id, title, author, isbn, quantity, category_id, created_by, created_at 
	          FROM books LIMIT $1 OFFSET $2`
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Quantity, &book.CategoryID, &book.CreatedBy, &book.CreatedAt); err != nil {
			return nil, 0, 0, err
		}
		books = append(books, book)
	}

	var totalData int
	countQuery := `SELECT COUNT(id) FROM books`
	err = r.DB.QueryRow(countQuery).Scan(&totalData)
	if err != nil {
		return nil, 0, 0, err
	}

	return books, len(books), totalData, nil
}

func (r *BookRepository) GetBooksByTitle(limit, offset int, filter string) ([]entity.Book, int, int, error) {
	query := `SELECT id, title, author, isbn, quantity, category_id, created_by, created_at 
	          FROM books WHERE title ILIKE $1 LIMIT $2 OFFSET $3`
	rows, err := r.DB.Query(query, "%"+filter+"%", limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Quantity, &book.CategoryID, &book.CreatedBy, &book.CreatedAt); err != nil {
			return nil, 0, 0, err
		}
		books = append(books, book)
	}

	var totalData int
	countQuery := `SELECT COUNT(id) FROM books`
	err = r.DB.QueryRow(countQuery).Scan(&totalData)
	if err != nil {
		return nil, 0, 0, err
	}

	return books, len(books), totalData, nil
}

func (r *BookRepository) DeleteBook(bookID int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.DB.Exec(query, bookID)

	return err
}
