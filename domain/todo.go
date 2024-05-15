package domain

import "time"

type Todo struct {
	ID        int
	Title     string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
