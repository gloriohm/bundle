package initdb

import (
	"database/sql"
	"fmt"

	"github.com/gloriohm/bundle/tasks"
)

func CreateGoal(db *sql.DB, name string) (int64, error) {
	result, err := db.Exec("INSERT INTO goals (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("failes to insert task: %w", err)
	}
	return result.LastInsertId()
}

func CreateQuest(db *sql.DB, quest tasks.Quest) (int64, error) {
	result, err := db.Exec("INSERT INTO goals (name, priority, deadlines) VALUES (?, ?, ?)", quest.Name, quest.Priority, quest.Deadline)
	if err != nil {
		return 0, fmt.Errorf("failes to insert task: %w", err)
	}
	return result.LastInsertId()
}

func CreateTask[T tasks.Task](db *sql.DB, task T, fields map[string]interface{}) error {
	_, err := db.Exec("INSERT INTO")
}
