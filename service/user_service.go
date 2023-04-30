package service

import (
	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository"
	"github.com/asaskevich/govalidator"
)

type UserService interface {
	Register(payload dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(payload dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {

	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return nil, errs.NewBadRequest(err.Error())
	}

	user := entity.User{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Password: payload.Password,
	}

	errs := user.HashPassword()

	if errs != nil {
		return nil, errs
	}

	returned, errs := u.userRepo.CreateNewUser(user)

	if errs != nil {
		return nil, errs
	}

	response := dto.RegisterResponse{
		ID:        returned.ID,
		Fullname:  returned.Fullname,
		Email:     returned.Email,
		CreatedAt: returned.CreatedAt,
	}

	return &response, nil
}
