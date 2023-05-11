package task_pg

import (
	"errors"

	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/task_repository"
	"gorm.io/gorm"
)

type taskPG struct {
	db *gorm.DB
}

func NewTaskPG(db *gorm.DB) task_repository.TaskRepository {
	return &taskPG{
		db: db,
	}
}

func (t *taskPG) GetTasks(userId uint) ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task
	var user entity.User

	err := t.db.Preload("Tasks").Where("id = ?", userId).First(&user).Error

	tasks = user.Tasks

	if err != nil {
		return nil, errs.NewInternalServerError("something Went Wrong")
	}

	return tasks, nil

}

func (t *taskPG) CreateTask(task *entity.Task) errs.MessageErr {
	err := t.db.Debug().Create(&task).Error

	if err != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *taskPG) GetTaskById(taskId uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	err := t.db.Debug().First(&task, taskId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Task didn't exist")
		}
		return nil, errs.NewInternalServerError("Internal Server Error")
	}

	return &task, nil
}

func (t *taskPG) EditTask(task *entity.Task) errs.MessageErr {

	result := t.db.Save(&task)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *taskPG) DeleteTask(taskId uint) errs.MessageErr {
	result := t.db.Delete(&entity.Task{}, taskId)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}
