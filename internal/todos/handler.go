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
