package tasks

import (
	"database/sql"
	"fmt"
)

func CreateGoal(db *sql.DB, name string) (int64, error) {
	result, err := db.Exec("INSERT INTO goals (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("failes to insert task: %w", err)
	}
	return result.LastInsertId()
}

func GetGoals(db *sql.DB) ([]Goal, error) {
	rows, err := db.Query("SELECT id, name, created_at FROM goals")
	if err != nil {
		return nil, fmt.Errorf("failes to query task: %w", err)
	}
	defer rows.Close()

	var goals []Goal
	for rows.Next() {
		var t Goal
		err := rows.Scan(&t.ID, &t.Name, &t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		goals = append(goals, t)
	}
	return goals, nil
}
