package todos

import (
	"encoding/json"
	"net/http"
	"todo-app/internal/app"
	"todo-app/internal/utils"
)

type Handler struct {
	service *Service
	app     *app.App
}

func NewHandler(service *Service, app *app.App) *Handler {
	return &Handler{service: service, app: app}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, VALIDATION, http.StatusBadRequest)
		return
	}
	res, err := h.service.CreateTodo(r.Context(), req)

	if err != nil {
		e := utils.CreateAPIError(err)
		utils.WriteJSON(w, e.StatusCode, e)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, res)
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var req GetTodosRequest

	if err := h.app.Decoder.Decode(&req, r.URL.Query()); err != nil {
		http.Error(w, VALIDATION, http.StatusBadRequest)
		return
	}
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	res, error := h.service.GetTodos(r.Context(), req)

	if error != nil {
		e := utils.CreateAPIError(error)
		utils.WriteJSON(w, e.StatusCode, e)
		return
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *Handler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	todoId, err := utils.ToInt(r, "id")

	if err != nil {
		http.Error(w, ID_MISSING, http.StatusBadRequest)
		return
	}

	res, error := h.service.GetTodoDetails(r.Context(), todoId)

	if error != nil {
		e := utils.CreateAPIError(error)
		utils.WriteJSON(w, e.StatusCode, e)
		return
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := utils.ToInt(r, "id")

	if err != nil {
		http.Error(w, ID_MISSING, http.StatusBadRequest)
		return
	}

	error := h.service.DeleteTodo(r.Context(), todoId)

	if error != nil {
		e := utils.CreateAPIError(error)
		utils.WriteJSON(w, e.StatusCode, e)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
