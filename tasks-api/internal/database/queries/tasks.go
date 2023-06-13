package queries

import (
	"go-playground/tasks-api/api/types"
	"go-playground/tasks-api/internal/database"
)

func GetAllTasks() []types.Task {
	db := database.GetConnection()
	rows, _ := db.Query("SELECT * FROM tasks")

	tasks := []types.Task{}
	for rows.Next() {
		task := types.Task{}
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt)
		tasks = append(tasks, task)
	}

	db.Close()
	return tasks
}

func InsertTask(task *types.Task) {
	db := database.GetConnection()

	insertSQL := "INSERT INTO tasks (title, description, created_at) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(insertSQL, task.Title, task.Description, task.CreatedAt).Scan(&task.Id)

	db.Close()

	if err != nil {
		panic(err)
	}
}
