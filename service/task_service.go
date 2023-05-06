package service

import (
	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/category_repository"
	"github.com/alvingxv/kanban-board-kelompok5/repository/task_repository"
)

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
}

type TaskService interface {
	CreateTask(payload dto.CreateTaskRequest, userId uint) (*dto.CreateTaskResponse, errs.MessageErr)
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (ts *taskService) CreateTask(payload dto.CreateTaskRequest, userId uint) (*dto.CreateTaskResponse, errs.MessageErr) {

	err := ts.categoryRepo.GetCategoryById(payload.CategoryId)

	if err != nil {
		return nil, err
	}

	task := &entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
		CategoryID:  payload.CategoryId,
		UserID:      userId,
	}

	err = ts.taskRepo.CreateTask(task)

	if err != nil {
		return nil, err
	}

	response := dto.CreateTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Status:      task.Status,
		Description: task.Description,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}

	return &response, nil
}
