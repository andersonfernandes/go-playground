package router

import (
	"context"
	"encoding/json"
	"go-playground/notes-api/internal/database/queries"
	"go-playground/notes-api/internal/types"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NotesRoutes(r *chi.Mux) {
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", ListNotes)
		r.Post("/", CreateNote)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(NoteCtx)
			r.Get("/", GetNote)
			r.Put("/", UpdateNote)
			r.Delete("/", DeleteNote)
		})
	})
}

func validateRequestParams(n *types.Note) []string {
	e := []string{}

	if n.Title == "" {
		e = append(e, "title is required")
	}

	if n.Content == "" {
		e = append(e, "content is required")
	}

	return e
}

func NoteCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		note, err := queries.GetNote(id)

		if err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, ResponseBody{Success: false, Errors: []string{"Could not find Note with id=" + id}})
			return
		}

		ctx := context.WithValue(r.Context(), "note", note)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	note, ok := ctx.Value("note").(*types.Note)
	if !ok {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, ResponseBody{Success: false})
		return
	}

	render.JSON(w, r, ResponseBody{Success: true, Data: note})
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, ResponseBody{Success: true, Data: queries.GetAllNotes()})
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	note := &types.Note{CreatedAt: now, UpdatedAt: now}

	json.NewDecoder(r.Body).Decode(note)
	re := validateRequestParams(note)

	if len(re) != 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ResponseBody{Success: false, Errors: re})
		return
	}

	err := queries.InsertNote(note)

	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, ResponseBody{Success: false})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, ResponseBody{Success: true, Data: note})
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	note, ok := ctx.Value("note").(*types.Note)
	if !ok {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, ResponseBody{Success: false})
		return
	}

	newNote := &types.Note{}
	json.NewDecoder(r.Body).Decode(newNote)

	note.UpdatedAt = time.Now()
	note.Title = newNote.Title
	note.Content = newNote.Content

	err := queries.UpdateNote(note)

	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, ResponseBody{Success: false})
		return
	}

	render.JSON(w, r, ResponseBody{Success: true, Data: note})
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := queries.DeleteNote(id)

	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, ResponseBody{Success: false})
		return
	}

	render.JSON(w, r, ResponseBody{Success: true})
}
