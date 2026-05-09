package structs

import "time"

type Todo struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool      `json:"completed"`
}

type UpdateTodoRequest struct {
	Text      *string    `json:"text"`
	DueDate   *time.Time `json:"due_date"`
	Completed *bool      `json:"completed"`
}
