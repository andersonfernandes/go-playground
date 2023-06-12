package main

import (
	"go-playground/todo-api/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	api.StartServer()
}
