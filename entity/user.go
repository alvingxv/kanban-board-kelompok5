package entity

import (
	"os"
	"time"

	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secret string = os.Getenv("SECRET")

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

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
	}
}

func (u *User) signToken(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, _ := token.SignedString([]byte(secret))

	return tokenstring
}
