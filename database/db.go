package database

import (
	"fmt"

	"github.com/alvingxv/kanban-board-kelompok5/entity"
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

	var err error
	db, err = gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(entity.User{})
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
