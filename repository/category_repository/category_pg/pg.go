package category_pg

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/category_repository"
	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG(db *gorm.DB) category_repository.CategoryRepository {
	return &categoryPG{
		db: db,
	}
}

func (c *categoryPG) CreateCategory(category *entity.Category) errs.MessageErr {

	err := c.db.Create(&category).Error

	if err != nil {
		return errs.NewInternalServerError("something Went Wrong")
	}

	return nil

}
