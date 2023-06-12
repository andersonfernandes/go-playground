package handlers

import (
	"go-playground/simple-webserver/pkg/restserever"
	"go-playground/todo-api/api/types"
	"go-playground/todo-api/internal/dbconnection"
	"net/http"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
  db := dbconnection.Get()
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
  db.Close()
}
