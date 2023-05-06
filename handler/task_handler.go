package handler

import "github.com/alvingxv/kanban-board-kelompok5/service"

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) taskHandler {

	return taskHandler{
		taskService: taskService,
	}
}
