package handler

import (
	"os"

	"github.com/alvingxv/kanban-board-kelompok5/database"
	"github.com/alvingxv/kanban-board-kelompok5/docs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/category_repository/category_pg"
	"github.com/alvingxv/kanban-board-kelompok5/repository/task_repository/task_pg"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository/user_pg"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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

	port := os.Getenv("PORT")
	// port := "4000"
	// port := helpers.GoDotEnvVariable("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	docs.SwaggerInfo.Title = "Kanban Board Kelompok 5"
	docs.SwaggerInfo.Description = "Final Project 3 Hactiv8 by Kelompok 5"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Schemes = []string{"https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)

		userRoute.Use(authService.Authentication())

		userRoute.PUT("/update-account", userHandler.UpdateUser)
		userRoute.DELETE("delete-account", userHandler.DeleteUser)

	}

	categoryRoute := r.Group("/categories")
	{

		categoryRoute.Use(authService.Authentication())

		categoryRoute.GET("", categoryHandler.GetCategory)
		categoryRoute.POST("", categoryHandler.CreateCategory)
		categoryRoute.PATCH("/:id", categoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	taskRoute := r.Group("/tasks")
	{
		taskRoute.Use(authService.Authentication())

		taskRoute.GET("", taskHandler.GetTasks)
		taskRoute.POST("", taskHandler.CreateTask)
		taskRoute.PUT("/:id", taskHandler.EditTask)
		taskRoute.PATCH("/update-status/:id", taskHandler.UpdateTaskStatus)
		taskRoute.PATCH("/update-category/:id", taskHandler.UpdateTaskCategory)
		taskRoute.DELETE("/:id", taskHandler.DeleteTask)
	}
	r.Run(":" + port)
}
