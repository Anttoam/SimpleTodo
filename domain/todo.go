package domain

import "time"

type Todo struct {
	Title     string
	Body      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
