package entity

type BookAnalytics struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	BorrowCount uint   `json:"borrow_count"`
}

type BorrowingTrends struct {
	Month       string `json:"month"`
	BorrowCount uint   `json:"borrow_count"`
}

