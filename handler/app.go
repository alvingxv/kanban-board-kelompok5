package handler

import (
	"github.com/alvingxv/kanban-board-kelompok5/database"
	"github.com/alvingxv/kanban-board-kelompok5/pkg"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository/user_pg"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := pkg.GoDotEnvVariable("PORT")

	database.HandleDatabaseConnection()

	db := database.GetDatabaseInstance()

	r := gin.Default()

	// User Injection
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.Run("127.0.0.1:" + port)
}
