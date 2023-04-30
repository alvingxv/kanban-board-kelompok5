package handler

import (
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Register(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

func (uh *userHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(loginRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
