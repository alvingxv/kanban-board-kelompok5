package handler

import (
	"github.com/alvingxv/kanban-board-kelompok5/database"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/helpers"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository/user_pg"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := helpers.GoDotEnvVariable("PORT")

	database.HandleDatabaseConnection()

	db := database.GetDatabaseInstance()

	r := gin.Default()

	// User Injection
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// Auth Injecttion
	authService := service.NewAuthService(userRepo)

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)

		userRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)

	}

	r.Run("127.0.0.1:" + port)
}
