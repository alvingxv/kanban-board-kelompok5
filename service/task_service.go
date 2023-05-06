package service

import "github.com/alvingxv/kanban-board-kelompok5/repository/task_repository"

type taskService struct {
	taskRepo task_repository.TaskRepository
}

type TaskService interface {
}

func NewTaskService(taskRepo task_repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}
