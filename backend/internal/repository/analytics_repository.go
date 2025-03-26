package repository

import (
	"database/sql"

	"github.com/ramzyrsr/digital-library/internal/entity"
)

type AnalyticsRepository struct {
	DB *sql.DB
}

func (r *AnalyticsRepository) GetMostBorrowedBooks(limit int) ([]entity.BookAnalytics, error) {
	query := `SELECT 
		b.id, b.title, b.author, COUNT(l.id) as borrow_count 
	FROM
		books b
	JOIN
		lending l ON b.id = l.book_id
	GROUP BY 
		b.id, b.title, b.author
	ORDER BY 
		borrow_count DESC
	LIMIT $1
	`
	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.BookAnalytics
	for rows.Next() {
		var book entity.BookAnalytics
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.BorrowCount); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
