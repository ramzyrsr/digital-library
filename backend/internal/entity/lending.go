package entity

import "time"

type Lending struct {
	ID          uint      `json:"id"`
	BookID      uint      `json:"book_id"`
	MemberID    uint      `json:"member_id"`
	BorroweDate time.Time `json:"borrowed_date"`
	DueDate     time.Time `json:"due_date"`
	ReturnDate  time.Time `json:"return_date"`
	Status      string    `json:"status"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
