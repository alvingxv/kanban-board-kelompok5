package task_repository

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
)

type TaskRepository interface {
	GetTasks(userId uint) ([]entity.Task, errs.MessageErr)
	CreateTask(task *entity.Task) errs.MessageErr
	GetTaskById(taskId uint) (*entity.Task, errs.MessageErr)
	EditTask(task *entity.Task) errs.MessageErr
	DeleteTask(taskId uint) errs.MessageErr
}
