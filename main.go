package main

import (
	"fmt"
	"log"

	"github.com/gloriohm/bundle/initdb"
)

func main() {
	fmt.Println("Starting Bundle...")

	dbPath, err := initdb.Initialize()
	if err != nil {
		log.Fatalf("Error initializing: %v", err)
	}

	fmt.Println("Using database:", dbPath)
}
