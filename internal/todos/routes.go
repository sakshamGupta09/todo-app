package todos

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("GET /todos", handler.GetTodos)
	mux.HandleFunc("GET /todos/{id}", handler.GetTodoById)
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)
}
