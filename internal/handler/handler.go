package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Yashk767/internal/service"
	"github.com/Yashk767/internal/structs"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodoReq structs.CreateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&createTodoReq)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to decode request body")
		return
	}

	todoReq, err := h.service.Create(createTodoReq)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, todoReq)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodoReq structs.UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&updateTodoReq)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to decode request body")
		return
	}

	id, err := GetIdInt64(r)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to get id from request")
		return
	}

	todoReq, err := h.service.Update(id, updateTodoReq)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, todoReq)
}

func (h *Handler) GetToDo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdInt64(r)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to get id from request")
		return
	}

	todoReq, err := h.service.GetById(id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, todoReq)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetIdInt64(r)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to get id from request")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "todo deleted successfully"})
}

func (h *Handler) GetAllToDos(w http.ResponseWriter, r *http.Request) {
	includeCompleted := r.PathValue("include_completed")

	includeCompletedBool, err := strconv.ParseBool(includeCompleted)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to parse include_completed query")
		return
	}

	todos, err := h.service.List(includeCompletedBool)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, todos)
}

func GetIdInt64(r *http.Request) (int64, error) {
	id := r.PathValue("id")

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return idInt64, nil
}
