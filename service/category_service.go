package service

import (
	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/category_repository"
	"github.com/asaskevich/govalidator"
)

type categoryService struct {
	categoryRepository category_repository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(payload dto.CategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr)
	UpdateCategory(payload dto.CategoryRequest, id uint) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(id uint) errs.MessageErr
}

func NewCategoryService(categoryRepository category_repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *categoryService) CreateCategory(payload dto.CategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr) {

	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	category := &entity.Category{
		Type: payload.Type,
	}

	err := cs.categoryRepository.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	response := dto.CreateCategoryResponse{
		Id:        category.ID,
		Type:      category.Type,
		CreatedAt: category.CreatedAt,
	}

	return &response, nil
}

func (cs *categoryService) UpdateCategory(payload dto.CategoryRequest, id uint) (*dto.UpdateCategoryResponse, errs.MessageErr) {

	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	category := &entity.Category{
		ID:   id,
		Type: payload.Type,
	}

	err := cs.categoryRepository.UpdateCategory(category)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateCategoryResponse{
		Id:        category.ID,
		Type:      category.Type,
		UpdatedAt: category.UpdatedAt,
	}

	return &response, nil
}

func (cs *categoryService) DeleteCategory(id uint) errs.MessageErr {

	err := cs.categoryRepository.DeleteCategory(id)

	if err != nil {
		return err
	}

	return nil

}
