package domain

import "time"

type Todo struct {
	ID          int
	Title       string
	Description string
	Done        bool
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
