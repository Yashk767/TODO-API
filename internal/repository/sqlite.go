package repository

import (
	"database/sql"
	"errors"
	"os"
	"path"
	"time"

	"github.com/Yashk767/internal/structs"
	_ "modernc.org/sqlite"
)

const databaseDir = "database"
const databaseFile = "todos.db"

type SqliteRepository struct {
	DB *sql.DB
}

func NewSqliteRepository() (*SqliteRepository, error) {
	if err := os.MkdirAll(databaseDir, os.ModePerm); err != nil {
		return nil, err
	}

	dbPath := path.Join(databaseDir, databaseFile)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		due_date DATETIME NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)`)

	if err != nil {
		return nil, err
	}

	return &SqliteRepository{DB: db}, nil
}

func (r *SqliteRepository) Create(todo structs.CreateTodoRequest) (structs.Todo, error) {
	var createdTodo structs.Todo
	createdTodo.Text = todo.Text
	createdTodo.DueDate = todo.DueDate

	now := time.Now().UTC()

	createdTodo.CreatedAt = now
	createdTodo.UpdatedAt = now

	stmt, err := r.DB.Prepare(`
	INSERT INTO todos (text, due_date, completed, created_at, updated_at)
	VALUES(?,?,?,?,?)
	`)
	if err != nil {
		return structs.Todo{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(createdTodo.Text, createdTodo.DueDate, false, createdTodo.CreatedAt, createdTodo.UpdatedAt)
	if err != nil {
		return structs.Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.Todo{}, err
	}
	createdTodo.Id = int64(id)

	return createdTodo, nil
}

func (r *SqliteRepository) Update(id int64, todo structs.UpdateTodoRequest) (structs.Todo, error) {
	updatedTodo, err := r.GetById(id)
	if err != nil {
		return structs.Todo{}, err
	}

	updatedTodo.Id = id
	if todo.Text != nil {
		updatedTodo.Text = *todo.Text
	}
	if todo.DueDate != nil {
		updatedTodo.DueDate = *todo.DueDate
	}
	if todo.Completed != nil {
		updatedTodo.Completed = *todo.Completed
	}

	updatedTodo.UpdatedAt = time.Now().UTC()

	stmt, err := r.DB.Prepare(`UPDATE todos
	SET text = ? , due_date = ? , completed = ? , updated_at = ?
	WHERE id = ?
	`)

	if err != nil {
		return structs.Todo{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(updatedTodo.Text, updatedTodo.DueDate, updatedTodo.Completed, updatedTodo.UpdatedAt, id)
	if err != nil {
		return structs.Todo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return structs.Todo{}, err
	}

	if rowsAffected == 0 {
		return structs.Todo{}, ErrTodoNotFound
	}

	return updatedTodo, nil
}

func (r *SqliteRepository) Delete(id int64) error {
	stmt, err := r.DB.Prepare(`DELETE FROM todos WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

func (r *SqliteRepository) GetById(id int64) (structs.Todo, error) {
	var todo structs.Todo
	stmt, err := r.DB.Prepare(`SELECT id, text, due_date, completed, created_at, updated_at
	FROM todos
	WHERE id = ? LIMIT 1
	`)
	if err != nil {
		return structs.Todo{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&todo.Id, &todo.Text, &todo.DueDate, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return structs.Todo{}, ErrTodoNotFound
		}
		return structs.Todo{}, err
	}

	return todo, nil
}

func (r *SqliteRepository) List(includeCompleted bool) ([]structs.Todo, error) {
	var todos []structs.Todo

	query := `
		SELECT id, text, due_date, completed, created_at, updated_at
		FROM todos
	`

	if !includeCompleted {
		query += ` WHERE completed = false`
	}

	query += ` ORDER BY due_date ASC`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo structs.Todo

		err = rows.Scan(&todo.Id, &todo.Text, &todo.DueDate, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}
