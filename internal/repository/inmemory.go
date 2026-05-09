package repository

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/Yashk767/internal/structs"
)

var ErrTodoNotFound = errors.New("todo not found")

type InMemoryRepository struct {
	mu       sync.RWMutex
	todoList map[int64]structs.Todo
	nextId   int64
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		mu:       sync.RWMutex{},
		todoList: make(map[int64]structs.Todo),
		nextId:   1,
	}
}

func (r *InMemoryRepository) Create(todo structs.Todo) (structs.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.Id = r.nextId
	todo.CreatedAt = time.Now().UTC()
	todo.UpdatedAt = time.Now().UTC()

	r.todoList[todo.Id] = todo
	r.nextId++

	return todo, nil
}

func (r *InMemoryRepository) Update(id int64, updateTodo structs.UpdateTodoRequest) (structs.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	exisitingTodo, ok := r.todoList[id]
	if !ok {
		return structs.Todo{}, ErrTodoNotFound
	}

	if updateTodo.Text != nil {
		exisitingTodo.Text = *updateTodo.Text
	}

	if updateTodo.DueDate != nil {
		exisitingTodo.DueDate = *updateTodo.DueDate
	}

	if updateTodo.Completed != nil {
		exisitingTodo.Completed = *updateTodo.Completed
	}

	exisitingTodo.UpdatedAt = time.Now().UTC()
	r.todoList[id] = exisitingTodo

	return exisitingTodo, nil
}

func (r *InMemoryRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.todoList[id]; !ok {
		return ErrTodoNotFound
	}

	delete(r.todoList, id)
	return nil
}

func (r *InMemoryRepository) GetById(id int64) (structs.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo, ok := r.todoList[id]
	if !ok {
		return structs.Todo{}, ErrTodoNotFound
	}

	return todo, nil
}

func (r *InMemoryRepository) List(shouldInclude bool) ([]structs.Todo, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var todos []structs.Todo

	for _, todo := range r.todoList {
		if shouldInclude || !todo.Completed {
			todos = append(todos, todo)
		}
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].DueDate.Before(todos[j].DueDate)
	})

	return todos, nil

}
