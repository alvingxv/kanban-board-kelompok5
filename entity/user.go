package entity

import (
	"strings"
	"time"

	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/helpers"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secret = helpers.GoDotEnvVariable("SECRET")

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Fullname  string `gorm:"not null;column:full_name"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null;default:member"`
	Tasks     []Task
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

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {

	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return errs.NewUnauthorizedError("Invalid Token")
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return errs.NewUnauthorizedError("Invalid Token")
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return errs.NewUnauthorizedError("Invalid Token")
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errs.NewUnauthorizedError("Invalid Token")
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err

}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthorizedError("Invalid Token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errs.NewUnauthorizedError("Invalid Token")
	}

	return token, nil

}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return errs.NewUnauthorizedError("Invalid Token")
	} else {
		u.ID = uint(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return errs.NewUnauthorizedError("Invalid Token")
	} else {
		u.Email = string(email)
	}

	return nil

}
