package main

import (
	"database/sql"
	"fmt"
)

var print = fmt.Println

func createTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		priority TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		print("Failed to create the table")
		
	}
}

func addTask(db *sql.DB, description, priority string) int {
	insertTaskSQL := `INSERT INTO tasks (description, priority, completed) VALUES (?, ?, ?)`
	result, err := db.Exec(insertTaskSQL, description, priority, false)
	if err != nil {
		print("Failed to add the task")
		return -1
	}
	taskID, _ := result.LastInsertId()
	print("Task was added successfully:", description)
	return int(taskID)
}

func delTask(db *sql.DB, id int) {
	deleteTaskSQL := `DELETE FROM tasks WHERE id = ?`
	_, err := db.Exec(deleteTaskSQL, id)
	if err != nil {
		print("Failed to delete task with id:", id)
		return
	}
	print("Deleted task with id:", id)
}

func getTasks(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT id, description, priority, completed FROM tasks")
	if err != nil {
		print("Failed to fetch tasks:", err)
	}
	return rows
}

func completeTask(db *sql.DB, id string, completed bool) {
	completeTaskSQL := `UPDATE tasks SET completed = ? WHERE id = ?`
	_, err := db.Exec(completeTaskSQL, completed, id)
	if err != nil {
		print("Failed to update task:", err)
	}
	print("Updated task ID", id, "completion status to", completed)
}

func clearTasks(db *sql.DB) {
	clearTasksSQL := `DELETE FROM tasks`
	_, err := db.Exec(clearTasksSQL)
	if err != nil {
		print("Failed to clear tasks:", err)
	}
	print("Cleared all tasks")
}
