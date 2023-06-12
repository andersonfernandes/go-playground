package handlers

import (
	"database/sql"
	"go-playground/simple-webserver/pkg/restserever"
	"go-playground/todo-api/api/types"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "postgres://postgres:@localhost:5439/todo_go?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rb := restserever.ResponseBody{}

	switch r.Method {
	case http.MethodGet:
		rows, _ := db.Query("SELECT * FROM tasks")

		tasks := []types.Task{}
		for rows.Next() {
			task := types.Task{}
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
