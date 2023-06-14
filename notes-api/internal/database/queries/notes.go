package queries

import (
	"go-playground/notes-api/internal/database"
	"go-playground/notes-api/internal/types"
)

func GetAllNotes() []types.Note {
	db := database.GetConnection()
	rows, _ := db.Query("SELECT * FROM notes")

	notes := []types.Note{}
	for rows.Next() {
		note := types.Note{}
		rows.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		notes = append(notes, note)
	}

	db.Close()
	return notes
}

func GetNote(id string) (*types.Note, error) {
	db := database.GetConnection()

  note := &types.Note{}
	s := "SELECT * FROM notes WHERE id=$1"
	err := db.QueryRow(s, id).Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

	db.Close()

  return note, err
}

func InsertNote(note *types.Note) error {
	db := database.GetConnection()

	s := "INSERT INTO notes (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(s, note.Title, note.Content, note.CreatedAt, note.UpdatedAt).Scan(&note.Id)

	db.Close()

  return err
}

func UpdateNote(note *types.Note) error {
	db := database.GetConnection()

	s := "UPDATE notes SET title=$1, content=$2, updated_at=$3 WHERE id=$4"
	_, err := db.Exec(s, note.Title, note.Content, note.UpdatedAt, note.Id)

	db.Close()

  return err
}

func DeleteNote(id string) error {
	db := database.GetConnection()

	s := "DELETE FROM notes WHERE id=$1"
	_, err := db.Exec(s, id)

	db.Close()

  return err
}

