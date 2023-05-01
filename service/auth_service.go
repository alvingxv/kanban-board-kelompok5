package service

import (
	"github.com/alvingxv/kanban-board-kelompok5/entity"
	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/alvingxv/kanban-board-kelompok5/repository/user_repository"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (as *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bearerToken := ctx.GetHeader("Authorization")

		var user = entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		usr, err := as.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), errs.NewUnauthenticatedError("Invalid Token"))
			return
		}

		ctx.Set("userData", usr)

		ctx.Next()

	}
}
