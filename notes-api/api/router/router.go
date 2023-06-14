package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, ResponseBody{Success: true})
	})

	NotesRoutes(r)
}
