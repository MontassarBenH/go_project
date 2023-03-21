package main

import "time"

type Task struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Completed   bool      `db:"completed"`
	CreatedAt   time.Time `db:"created_at"`
}
