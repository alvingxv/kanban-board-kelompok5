package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "root"
	dbname   = "kanban-hacktiv"
	dialect  = "postgres"
)

var db *gorm.DB

func HandleDatabaseConnection() {
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
