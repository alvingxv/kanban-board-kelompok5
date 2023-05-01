package service

import (
	"net/http"

	"github.com/alvingxv/kanban-board-kelompok5/dto"
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository"
	"github.com/asaskevich/govalidator"
)

type UserService interface {
	Register(payload dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(payload dto.UpdateRequest, userdata *entity.User) (*dto.UpdateResponse, errs.MessageErr)
	DeleteUser(id uint) errs.MessageErr
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

	errs = u.userRepo.CreateNewUser(&user)

	if errs != nil {
		return nil, errs
	}

	response := dto.RegisterResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return &response, nil
}

func (u *userService) Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {

	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Token: token,
	}

	return &response, nil
}

func (u *userService) UpdateUser(payload dto.UpdateRequest, userdata *entity.User) (*dto.UpdateResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}
	user := entity.User{
		ID:        userdata.ID,
		Fullname:  payload.Fullname,
		Email:     payload.Email,
		Password:  userdata.Password,
		Role:      userdata.Role,
		CreatedAt: userdata.CreatedAt,
		UpdatedAt: userdata.UpdatedAt,
	}

	err := u.userRepo.UpdateUser(&user)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) DeleteUser(id uint) errs.MessageErr {
	err := u.userRepo.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil

}
