package entity

import (
	"time"

	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Fullname  string `gorm:"not null;column:full_name"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null;default:member"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}
