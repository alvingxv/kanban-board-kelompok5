package helpers

import (
	"strconv"

	"github.com/alvingxv/kanban-board-kelompok5/pkg/errs"
	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context, key string) (uint, errs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter id")
	}

	return uint(id), nil
}
