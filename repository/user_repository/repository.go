package user_repository

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user *entity.User) errs.MessageErr
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	UpdateUser(user *entity.User) errs.MessageErr
	DeleteUser(id uint) errs.MessageErr
}
