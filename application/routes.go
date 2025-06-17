package application

import (
	"net/http"
	"orders-api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandlers := &handler.Order{}


	router.Post("/", orderHandlers.Create)
	router.Get("/", orderHandlers.List)
	router.Get("/{id}", orderHandlers.GetByID)
	router.Put("/{id}", orderHandlers.UpdateByID)
	router.Delete("/{id}", orderHandlers.DeleteByID)
}