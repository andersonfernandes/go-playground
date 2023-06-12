package handlers

import (
	"encoding/json"
	"go-playground/simple-webserver/pkg/restserever"
	"go-playground/todo-api/api/types"
	"go-playground/todo-api/internal/database/queries"
	"net/http"
	"time"
)

func getAll() []types.Task {
	return queries.GetAllTasks()
}

func create(r *http.Request) (*types.Task, error) {
	task := types.Task{CreatedAt: time.Now()}

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		return nil, err
	}

	queries.InsertTask(&task)

	return &task, nil
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	rb := restserever.ResponseBody{}
	s := http.StatusOK

	switch r.Method {
	case http.MethodGet:
		rb.Message = "All Tasks"
		rb.Data = getAll()
	case http.MethodPost:
		task, err := create(r)

		if err != nil {
			s = http.StatusBadRequest
			rb.Message = "Could not create Task"
			rb.Data = err.Error()
		} else {
			s = http.StatusCreated
			rb.Message = "Task Created"
			rb.Data = task
		}
	}

	rb.WriteJsonResponse(w, s)
}
