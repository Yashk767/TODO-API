package todo

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool      `json:"completed"`
}
