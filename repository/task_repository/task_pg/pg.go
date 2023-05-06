package task_pg

import (
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

func (t *taskPG) CreateTask(task *entity.Task) errs.MessageErr {
	err := t.db.Debug().Create(&task).Error

	if err != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}
