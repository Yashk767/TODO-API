package repository

import "github.com/Yashk767/internal/structs"

type Repository interface {
	Create(todo structs.CreateTodoRequest) (structs.Todo, error)
	Update(id int64, todo structs.UpdateTodoRequest) (structs.Todo, error)
	Delete(id int64) error
	GetById(id int64) (structs.Todo, error)
	List(shouldInclude bool) ([]structs.Todo, error)
}
