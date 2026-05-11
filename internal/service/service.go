package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Yashk767/internal/repository"
	"github.com/Yashk767/internal/structs"
)

var (
	ErrInvalidTodoText = errors.New("invalid todo text")
	ErrInvalidDueDate  = errors.New("invalid due date")
	ErrDueDateInPast   = errors.New("due date cannot be in the past")
	ErrInvalidId       = errors.New("invalid id")
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(createReq structs.CreateTodoRequest) (structs.Todo, error) {
	createReq.Text = strings.TrimSpace(createReq.Text)

	if createReq.Text == "" {
		return structs.Todo{}, ErrInvalidTodoText
	}

	if createReq.DueDate.IsZero() {
		return structs.Todo{}, ErrInvalidDueDate
	}

	if time.Now().UTC().After(createReq.DueDate) {
		return structs.Todo{}, ErrDueDateInPast
	}

	return s.repository.Create(createReq)
}

func (s *Service) Update(id int64, updateReq structs.UpdateTodoRequest) (structs.Todo, error) {
	if id <= 0 {
		return structs.Todo{}, ErrInvalidId
	}
	if updateReq.Text != nil {
		trimmedText := strings.TrimSpace(*updateReq.Text)

		updateReq.Text = &trimmedText
	}

	if updateReq.DueDate != nil {
		if updateReq.DueDate.IsZero() {
			return structs.Todo{}, ErrInvalidDueDate
		}

		if time.Now().UTC().After(*updateReq.DueDate) {
			return structs.Todo{}, ErrDueDateInPast
		}
	}
	return s.repository.Update(id, updateReq)
}

func (s *Service) Delete(id int64) error {
	if id <= 0 {
		return ErrInvalidId
	}
	return s.repository.Delete(id)
}

func (s *Service) GetById(id int64) (structs.Todo, error) {
	if id <= 0 {
		return structs.Todo{}, ErrInvalidId
	}
	return s.repository.GetById(id)
}

func (s *Service) List(shouldInclude bool) ([]structs.Todo, error) {
	return s.repository.List(shouldInclude)
}
