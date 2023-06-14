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

func NoteCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    note, err := queries.GetNote(id)

    if err != nil {
      http.Error(w, http.StatusText(404), 404)
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
    http.Error(w, http.StatusText(422), 422)
    return
  }

  render.JSON(w, r, note)
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, queries.GetAllNotes())
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
  now := time.Now()
  note := &types.Note{CreatedAt: now, UpdatedAt: now}

  json.NewDecoder(r.Body).Decode(note)
  err := queries.InsertNote(note)

  if err != nil {
    http.Error(w, http.StatusText(422), 422)
    return
  }

  render.Status(r, http.StatusCreated)
	render.JSON(w, r, note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request)  {
  ctx := r.Context()
  note, ok := ctx.Value("note").(*types.Note)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }

  newNote := &types.Note{}
  json.NewDecoder(r.Body).Decode(newNote)
  
  note.UpdatedAt = time.Now()
  note.Title = newNote.Title
  note.Content = newNote.Content

  err := queries.UpdateNote(note)

  if err != nil {
    http.Error(w, http.StatusText(422), 422)
    return
  }

	render.JSON(w, r, note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  err := queries.DeleteNote(id)

  if err != nil {
    http.Error(w, http.StatusText(422), 422)
    return
  }

  render.Status(r, http.StatusNoContent)
  render.JSON(w, r, "")
}
