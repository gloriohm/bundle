package initdb

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "bundle.db"

func Initialize() (string, error) {
	config, err := LoadConfig()
	if err != nil || config.DBPath == "" {
		fmt.Println("No database path found. Running onboarding process...")
		dbPath, err := specifyPath()
		if err != nil {
			return "", err
		}
		return createDatabase(dbPath)
	}

	if _, err := os.Stat(config.DBPath); os.IsNotExist(err) {
		fmt.Println("Database file not found at:", config.DBPath)
		return handleMissingDB(config.DBPath)
	}

	fmt.Println("Database found at:", config.DBPath)
	return config.DBPath, nil
}

func specifyPath() (string, error) {
	// holds the path in a variable
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a path to store the database file: ")
	dbPath, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// ensure path is correctly formatted before saving
	cleanedPath, err := CleanPath(dbPath)
	if err != nil {
		return "", err
	}
	return cleanedPath, nil
}

func createDatabase(dbPath string) (string, error) {
	// Ensure directory exists and make it if it doesn't
	err := os.MkdirAll(dbPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	} else {
		fmt.Printf("Successfully created directory: %s\n", dbPath)
	}

	// append path with database name - bundle.db
	fullPath := dbPath + dbName

	// Create database file
	err = createSQLiteDB(fullPath)
	if err != nil {
		return "", err
	}

	// Save the DB path in config
	err = SaveConfig(fullPath)
	if err != nil {
		return "", err
	}

	fmt.Println("Database created at:", dbPath)
	return dbPath, nil
}

func createSQLiteDB(dbPath string) error {
	fmt.Println("Creating new SQLite database...")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	defer db.Close()

	// Create an initial table (example)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	fmt.Println("Database intialized successfully!")
	return nil
}

func handleMissingDB(dbPath string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nNo database found. What would you like to do?")
		fmt.Printf("1. Create a new database at %v \n", dbPath)
		fmt.Println("2. Specify a new path")
		fmt.Print("Enter your choice (1 or 2): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			return createDatabase(dbPath)
		case "2":
			return searchForDatabase()
		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")
		}
	}
}

func searchForDatabase() (string, error) {
	fmt.Print("Enter a new directory to search for 'bundle.db': ")
	dbPath, err := specifyPath()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("Database found at:", dbPath)
		SaveConfig(dbPath)
		return dbPath, nil
	}

	fmt.Println("No file bundle.db found at:", dbPath)
	return handleMissingDB(dbPath)
}
