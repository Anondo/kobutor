package server

import (
	"kobutor/api"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
)

var (
	router = chi.NewRouter()
)

func prepareRouter() {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		renderer.New().JSON(w, http.StatusNotFound, renderer.M{
			"message": "Route not found!",
		})
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		renderer.New().JSON(w, http.StatusNotFound, renderer.M{
			"message": "Method not allowed!",
		})
	})

	registerRoutes()
}

func registerRoutes() {
	router.Route("/v1/", func(r chi.Router) {
		r.Mount("/kobutor", mailHandlers())
	})

}

func mailHandlers() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Use(api.BasicAuth)
		h.Post("/", api.SendMail)
	})

	return h
}
