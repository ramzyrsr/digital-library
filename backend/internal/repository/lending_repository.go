package repository

import (
	"database/sql"
	"errors"

	"github.com/ramzyrsr/digital-library/internal/entity"
)

type LendingRepository struct {
	DB *sql.DB
}

func (r *LendingRepository) BorrowBook(lending *entity.Lending) error {
	var totalAvailable int
	query := `SELECT CASE
		WHEN 
			available_qty - borrowed_qty < 0 THEN 0
			ELSE available_qty - borrowed_qty
		END AS qty_difference 
	FROM 
		book_status 
	WHERE
		book_id = $1`
	err := r.DB.QueryRow(query, lending.BookID).Scan(&totalAvailable)
	if err != nil {
		return err
	}
	if totalAvailable == 0 {
		return errors.New("Cant borrow book. All already borrowed")
	}

	query = `INSERT INTO lending (book_id, member_id, borrowed_date, due_date, status, created_by)
	          VALUES ($1, $2, NOW(), NOW() + INTERVAL '2 weeks', $3, $4) RETURNING id`
	err = r.DB.QueryRow(query, lending.BookID, lending.MemberID, "borrowed", lending.CreatedBy).Scan(&lending.ID)
	if err != nil {
		return err
	}

	query = `UPDATE book_status SET borrowed_qty = borrowed_qty + 1 WHERE book_id = $1`
	_, err = r.DB.Exec(query, lending.BookID)

	return err
}

func (r *LendingRepository) ReturnBook(lendingID int) error {
	var bookID int
	var returnDate interface{}

	query := `SELECT return_date, book_id FROM lending WHERE id = $1`
	err := r.DB.QueryRow(query, lendingID).Scan(&returnDate, &bookID)
	if err != nil {
		return err
	}

	if returnDate != nil {
		return errors.New("book has already been returned")
	}

	query = `UPDATE lending SET return_date = NOW(), status = $1 WHERE id = $2`
	_, err = r.DB.Exec(query, "returned", lendingID)
	if err != nil {
		return err
	}

	query = `UPDATE book_status SET borrowed_qty = borrowed_qty - 1 WHERE book_id = $1`
	_, err = r.DB.Exec(query, bookID)

	return err
}
