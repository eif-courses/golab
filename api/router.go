package api

import (
	"github.com/eif-courses/golab/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/v1/users", controllers.GetAllUsers)
	router.Get("/api/v1/users/user/{id}", controllers.GetUserById)
	router.Post("/api/v1/users/user", controllers.CreateUser)
	router.Put("/api/v1/users/user/{id}", controllers.UpdateUser)
	router.Delete("/api/v1/users/user/{id}", controllers.DeleteUser)
	return router
}
