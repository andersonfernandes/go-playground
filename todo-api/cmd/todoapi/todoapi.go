package main

import (
	"database/sql"
	"go-playground/simple-webserver/pkg/restserever"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Task struct {
  Id          int       `db:"id" json:"id"`
  Name        string    `db:"name" json:"name"`
  Description string    `db:"description" json:"description"`
  CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

// TODO: Extract all funcs to the correct files
func rootHandler(w http.ResponseWriter, r *http.Request) {
	rb := restserever.ResponseBody{Message: "OK"}
	rb.WriteJsonResponse(w, http.StatusOK)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "postgres://postgres:@localhost:5439/todo_go?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rb := restserever.ResponseBody{}

	switch r.Method {
	case http.MethodGet:
		rows, _ := db.Query("SELECT * FROM tasks")

		tasks := []Task{}
		for rows.Next() {
			task := Task{}
			rows.Scan(&task.Id, &task.Name, &task.Description, &task.CreatedAt)
			tasks = append(tasks, task)
		}

		rb.Message = "Get All"
    rb.Data = tasks
	case http.MethodPost:
		rb.Message = "Create"
	}

	rb.WriteJsonResponse(w, http.StatusOK)
}

func main() {
	server := restserever.Server{
		Mux:  *http.NewServeMux(),
		Port: "8088",
	}

	server.Mux.HandleFunc("/", rootHandler)
	server.Mux.HandleFunc("/tasks", tasksHandler)

	server.Start()
}
