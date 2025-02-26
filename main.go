package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gloriohm/bundle/initdb"
	"github.com/gloriohm/bundle/tasks"
)

func main() {
	fmt.Println("Starting Bundle...")

	dbPath, err := initdb.Initialize()
	if err != nil {
		log.Fatalf("Error initializing: %v", err)
	}

	fmt.Println("Using database:", dbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\n==== Task Manager ====")
		fmt.Println("1) Show Tasks")
		fmt.Println("2) Add Task")
		fmt.Println("2) Delete Task")
		fmt.Println("4) Exit")
		fmt.Print("Choose an option: ")

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			tasks.ShowTasks(db)
		case "2":
			tasks.AddTask(db, reader)
		case "3":
			tasks.DeleteTask(db, reader)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
