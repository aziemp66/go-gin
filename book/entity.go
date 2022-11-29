package book

import "time"

type Book struct {
	ID          uint
	Title       string
	Description string
	Price       float64
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
