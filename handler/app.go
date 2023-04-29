package handler

import (
	"github.com/alvingxv/kanban-board-fp/database"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := "8000"

	database.HandleDatabaseConnection()

	db := database.GetDatabaseInstance()

	_ = db

	r := gin.Default()

	r.Run(":" + port)
}
