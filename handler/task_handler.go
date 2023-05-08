package handler

import (
	"fmt"
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/helpers"
	"github.com/alvingxv/kanban-board-kelompok5/service"
	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) taskHandler {
	return taskHandler{
		taskService: taskService,
	}
}

func (th *taskHandler) CreateTask(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(*entity.User)

	var taskRequest dto.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&taskRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := th.taskService.CreateTask(taskRequest, userData.ID)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (th *taskHandler) EditTask(ctx *gin.Context) {

	var taskRequest dto.EditTaskRequest

	if err := ctx.ShouldBindJSON(&taskRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	userData := ctx.MustGet("userData").(*entity.User)

	result, err := th.taskService.EditTask(taskRequest, id, userData.ID)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (th *taskHandler) UpdateTaskStatus(ctx *gin.Context) {
	var updateRequest dto.UpdateTaskStatusRequest

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		fmt.Print(err.Error())
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	userData := ctx.MustGet("userData").(*entity.User)

	result, err := th.taskService.UpdateStatusTask(updateRequest, id, userData.ID)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
