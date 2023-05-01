package user_pg

import (
	"errors"
	"strings"

	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository"
	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}
func (u *userPG) CreateNewUser(user *entity.User) errs.MessageErr {
	result := u.db.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			return errs.NewBadRequest("User Already Exist")
		}
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {

	user := entity.User{}

	err := u.db.Debug().Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User didn't exist")
		}
		return nil, errs.NewInternalServerError("Internal Server Error")
	}

	return &user, nil

}

func (u *userPG) UpdateUser(user *entity.User) errs.MessageErr {
	result := u.db.Save(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint \"users_email_key") {
			return errs.NewBadRequest("User Already Exist")
		}
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}
