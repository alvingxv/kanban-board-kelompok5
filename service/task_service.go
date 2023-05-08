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
	EditTask(payload dto.EditTaskRequest, taskId uint, userId uint) (*dto.EditTaskResponse, errs.MessageErr)
	UpdateStatusTask(payload dto.UpdateTaskStatusRequest, taskId uint, userId uint) (*dto.UpdateTaskStatusResponse, errs.MessageErr)
	UpdateTaskCategory(payload dto.UpdateTaskCategoryRequest, taskId uint, userId uint) (*dto.UpdateTaskCategoryResponse, errs.MessageErr)
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

func (ts *taskService) EditTask(payload dto.EditTaskRequest, taskId uint, userId uint) (*dto.EditTaskResponse, errs.MessageErr) {

	task, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("Unauthorized")
	}

	task.Title = payload.Title
	task.Description = payload.Description

	err = ts.taskRepo.EditTask(task)

	if err != nil {
		return nil, err
	}

	response := dto.EditTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		UpdatedAt:   task.UpdatedAt,
	}

	return &response, nil
}

func (ts *taskService) UpdateStatusTask(payload dto.UpdateTaskStatusRequest, taskId uint, userId uint) (*dto.UpdateTaskStatusResponse, errs.MessageErr) {

	task, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("Unauthorized")
	}

	task.Status = payload.Status

	err = ts.taskRepo.EditTask(task)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateTaskStatusResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		UpdatedAt:   task.UpdatedAt,
	}

	return &response, nil
}

func (ts *taskService) UpdateTaskCategory(payload dto.UpdateTaskCategoryRequest, taskId uint, userId uint) (*dto.UpdateTaskCategoryResponse, errs.MessageErr) {
	task, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("Unauthorized")
	}

	err = ts.categoryRepo.GetCategoryById(payload.CategoryId)

	if err != nil {
		return nil, err
	}

	task.CategoryID = payload.CategoryId

	err = ts.taskRepo.EditTask(task)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateTaskCategoryResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		UpdatedAt:   task.UpdatedAt,
	}

	return &response, nil
}
