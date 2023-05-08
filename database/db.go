package database

import (
	"fmt"

	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"golang.org/x/crypto/bcrypt"
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

	db.AutoMigrate(entity.User{}, entity.Category{}, entity.Task{})

	var user entity.User
	db.First(&user, "role = ?", "admin")

	if user.ID == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		user = entity.User{
			Fullname: "admin",
			Email:    "admin@gmail.com",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		err := db.Create(&user).Error

		if err != nil {
			panic("failed create Admin")
		}
	}
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
