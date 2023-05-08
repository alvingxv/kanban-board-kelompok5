package handler

import (
	"github.com/alvingxv/kanban-board-kelompok5/database"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/helpers"
	"github.com/alvingxv/kanban-board-kelompok5/repository/category_repository/category_pg"
	"github.com/alvingxv/kanban-board-kelompok5/repository/task_repository/task_pg"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository/user_pg"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/override/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {

	database.HandleDatabaseConnection()
	db := database.GetDatabaseInstance()

	// User Injection
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// Category Injection
	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	// Task Injection
	taskRepo := task_pg.NewTaskPG(db)
	taskService := service.NewTaskService(taskRepo, categoryRepo)
	taskHandler := NewTaskHandler(taskService)

	// Auth Injecttion
	authService := service.NewAuthService(userRepo)

	port := helpers.GoDotEnvVariable("PORT")
	r := gin.Default()

	docs.SwaggerInfo.Title = "Kanban Board Kelompok 5"
	docs.SwaggerInfo.Description = "Final Project 3 Hactiv8 by Kelompok 5"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)

		userRoute.PUT("/update-account", authService.Authentication(), userHandler.UpdateUser)
		userRoute.DELETE("delete-account", authService.Authentication(), userHandler.DeleteUser)

	}

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.POST("", authService.Authentication(), categoryHandler.CreateCategory)
		categoryRoute.PATCH("/:id", authService.Authentication(), categoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:id", authService.Authentication(), categoryHandler.DeleteCategory)
	}

	taskRoute := r.Group("/tasks")
	{
		taskRoute.POST("", authService.Authentication(), taskHandler.CreateTask)
		taskRoute.PUT("/:id", authService.Authentication(), taskHandler.EditTask)
	}
	r.Run("127.0.0.1:" + port)
}
