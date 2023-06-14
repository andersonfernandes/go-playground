package api

import (
	"go-playground/simple-webserver/pkg/restserver"
	"go-playground/tasks-api/api/handlers"
	"net/http"
)

func StartServer() {
	server := restserver.Server{
		Mux:  *http.NewServeMux(),
		Port: "8088",
	}

	server.Mux.HandleFunc("/", handlers.RootHandler)
	server.Mux.HandleFunc("/tasks", handlers.TasksHandler)

	server.Start()
}
