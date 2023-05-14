package database

import (
	"fmt"
	"os"

	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
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
