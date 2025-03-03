package tasks

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func ShowTasks(db *sql.DB) {
	goals, err := GetGoals(db)
	if err != nil {
		log.Println("Error fetching goals:", err)
		return
	}

	if len(goals) == 0 {
		fmt.Println("No goals found.")
		return
	}

	fmt.Println("\n==== Goals ====")
	for _, t := range goals {
		fmt.Printf("[%d] %s (Created: %s)\n", t.ID, t.Name, t.CreatedAt)
	}
}

func AddTask(db *sql.DB, reader *bufio.Reader) {
	var newTask []Goal
	var updatedFields []string
	// receive update date
	fmt.Print("Enter task name: ")
	taskTitle, _ := reader.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)

	if taskTitle == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}

	updatedFields = append(updatedFields, "name")

	taskID, err := CreateTask(db, newTask, updatedFields)
	if err != nil {
		log.Println("Failed to create task:", err)
		return
	}
	fmt.Printf("New task created with ID: %d\n", taskID)
}

func DeleteTask(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter task ID to delete: ")
	taskID, _ := reader.ReadString('\n')
	taskID = strings.TrimSpace(taskID)

	err := tasks.DeleteTask(db, taskID)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	fmt.Printf("Task %s deleted successfully!\n", taskID)
}
