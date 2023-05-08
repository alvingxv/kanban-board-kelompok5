package handler

import (
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/helpers"
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

	var categoryRequest dto.CategoryRequest

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

func (ch *categoryHandler) UpdateCategory(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(*entity.User)

	if userData.Role != "admin" {
		errUnauthorized := errs.NewUnauthorizedError("Unauthorized user")

		ctx.JSON(errUnauthorized.Status(), errUnauthorized)
		return
	}

	var categoryRequest dto.CategoryRequest

	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	result, err := ch.categoryService.UpdateCategory(categoryRequest, id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ch *categoryHandler) DeleteCategory(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(*entity.User)

	if userData.Role != "admin" {
		errUnauthorized := errs.NewUnauthorizedError("Unauthorized user")

		ctx.JSON(errUnauthorized.Status(), errUnauthorized)
		return
	}

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	err = ch.categoryService.DeleteCategory(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}

func (ch *categoryHandler) GetCategory(ctx *gin.Context) {

	userData := ctx.MustGet("userData").(*entity.User)

	result, err := ch.categoryService.GetCategory(userData.ID)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
