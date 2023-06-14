package api

import (
	"fmt"
	"go-playground/notes-api/api/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	router.RegisterRoutes(r)

	port := "8088"
	fmt.Println("Server started at http://localhost:" + port + ".")
	fmt.Print("Press ctrl + c to stop.\n\n")
	http.ListenAndServe(":"+port, r)
}
