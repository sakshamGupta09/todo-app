package todos

import (
	"net/http"
	"todo-app/internal/app"
)

func Setup(mux *http.ServeMux, app *app.App) {
	repo := NewRepository(app.DB)
	service := NewService(repo)
	handler := NewHandler(service, app)

	RegisterRoutes(mux, handler)
}
