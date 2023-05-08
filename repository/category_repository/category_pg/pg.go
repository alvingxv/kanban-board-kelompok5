package category_pg

import (
	"errors"

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

func (c *categoryPG) GetAllCategory(userId uint) ([]entity.Category, errs.MessageErr) {
	var categories []entity.Category

	result := c.db.Model(&entity.Category{}).Preload("Tasks", "user_id = ?", userId).Find(&categories).Error

	if result != nil {
		return nil, errs.NewInternalServerError("something Went Wrong")
	}

	return categories, nil
}

func (c *categoryPG) CreateCategory(category *entity.Category) errs.MessageErr {

	err := c.db.Create(&category).Error

	if err != nil {
		return errs.NewInternalServerError("something Went Wrong")
	}

	return nil

}

func (c *categoryPG) UpdateCategory(category *entity.Category) errs.MessageErr {
	result := c.db.Select("id").First(&category, category.ID)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = c.db.Model(&category).Update("type", category.Type)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (c *categoryPG) DeleteCategory(id uint) errs.MessageErr {
	result := c.db.Select("id").First(&entity.Category{}, id)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = c.db.Delete(&entity.Category{}, id)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (c *categoryPG) GetCategoryById(id uint) errs.MessageErr {
	err := c.db.Debug().First(&entity.Category{}, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("Category didn't exist")
		}
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}
