package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := Initialize()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	fmt.Println("Table created successfully!")
}
