package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Employee represents the employee model
type Employee struct {
	gorm.Model
	Name string
	Role string
}

// TableName overrides the default table name
func (Employee) TableName() string {
	return "employees"
}

// Initialize creates the table and sets up the database connection
func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gotest sslmode=disable password=ilovemessi")
	if err != nil {
		return nil, err
	}

	// Automigrate the Employee model to create the table
	db.AutoMigrate(&Employee{})

	return db, nil
}
