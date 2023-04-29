package handler

import (
	"github.com/alvingxv/kanban-board-kelompok5/database"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := "8000"

	database.HandleDatabaseConnection()

	db := database.GetDatabaseInstance()

	_ = db

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "asd")
	})

	r.Run(":" + port)
}
