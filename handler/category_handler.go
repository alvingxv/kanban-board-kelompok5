package handler

import (
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{
		categoryService: categoryService,
	}
}

func (ch *categoryHandler) CreateCategory(ctx *gin.Context) {

	userData := ctx.MustGet("userData").(*entity.User)

	if userData.Role != "admin" {
		errUnauthorized := errs.NewUnauthorizedError("Unauthorized user")

		ctx.JSON(errUnauthorized.Status(), errUnauthorized)
		return
	}

	var categoryRequest dto.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ch.categoryService.CreateCategory(categoryRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}
