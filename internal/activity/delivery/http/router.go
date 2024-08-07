package routes

import (
	"github.com/go-chi/chi/v5"
	"todo_activity/internal/activity/handler"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users", handler.GetUsers)
	r.Post("/users", handler.CreateUser)

	return r
}
