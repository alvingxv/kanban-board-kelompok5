package category_repository

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) errs.MessageErr
	UpdateCategory(category *entity.Category) errs.MessageErr
	DeleteCategory(id uint) errs.MessageErr
	GetCategoryById(id uint) errs.MessageErr
}
