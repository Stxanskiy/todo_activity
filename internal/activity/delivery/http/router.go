package routes

import (
	"github.com/go-chi/chi/v5"
	"todo_activity/internal/activity/handler/user"

	"todo_activity/pkg/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/users", user.RegisterUser)
	r.Post("/login", user.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/users", user.GetUsers)
	})

	return r
}
