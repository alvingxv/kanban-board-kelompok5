package entity

import (
	"time"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      bool   `gorm:"not null;default:false"`
	UserID      uint
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) TaskToCategoryTaskResponse() dto.CategoryTask {
	return dto.CategoryTask{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		UserID:      t.UserID,
		CategoryID:  t.CategoryID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
