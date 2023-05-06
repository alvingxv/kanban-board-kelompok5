package task_pg

import (
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
