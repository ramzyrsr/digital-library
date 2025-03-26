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

func (r *AnalyticsRepository) GetMonthlyBorrowingTrends() ([]entity.BorrowingTrends, error) {
	query := `SELECT 
		DATE_TRUNC('month', borrowed_date) as month, COUNT(id) as borrow_count
	FROM
		lending
	GROUP BY 
		month
	ORDER BY
		month
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trends []entity.BorrowingTrends
	for rows.Next() {
		var trend entity.BorrowingTrends
		if err := rows.Scan(&trend.Month, &trend.BorrowCount); err != nil {
			return nil, err
		}
		trends = append(trends, trend)
	}
	return trends, nil
}

func (r *AnalyticsRepository) GetBooksByCategory() ([]entity.BookCategoryDistribution, error) {
	query := `SELECT 
		c.name AS category_name, COUNT(b.id) AS book_count
	FROM
		books b
	JOIN 
		categories c ON b.category_id = c.id
	GROUP BY
		c.name
	ORDER BY 
		book_count DESC;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.BookCategoryDistribution
	for rows.Next() {
		var data entity.BookCategoryDistribution
		if err := rows.Scan(&data.CategoryName, &data.BookCount); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}
