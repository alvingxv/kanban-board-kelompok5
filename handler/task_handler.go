package handler

import (
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
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
