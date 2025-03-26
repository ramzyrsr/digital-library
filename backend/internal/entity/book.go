package entity

import "time"

type Book struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	ISBN       string    `json:"isbn"`
	Quantity   uint      `json:"quantity"`
	CategoryID uint      `json:"category_id"`
	CreatedBy  uint      `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}
