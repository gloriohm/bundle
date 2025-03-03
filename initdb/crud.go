package initdb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

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

func CreateTask[T tasks.Task](db *sql.DB, task T, keys []string) (int64, error) {
	table := task.TableName()
	query, values, err := constructQuery(table, task, keys)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(query, values)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId(), nil
}

func constructQuery[T tasks.Task](table string, task T, keys []string) (string, []interface{}, error) {
	v := reflect.ValueOf(task)
	t := v.Type()

	// Ensure we have a struct
	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("expected a struct, got %T", task)
	}

	var columns []string
	var placeholders []string
	var values []interface{}

	// Iterate over struct fields
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		columnName := field.Tag.Get("json")

		if columnName == "" {
			continue // Skip fields without a "db" tag
		}

		columns = append(columns, columnName)
		placeholders = append(placeholders, "?") // SQL placeholders
		values = append(values, v.Field(i).Interface())
	}

	// Construct the query dynamically
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, values, nil
}
