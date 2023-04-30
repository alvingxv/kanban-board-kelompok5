package user_repository

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user *entity.User) errs.MessageErr
}
