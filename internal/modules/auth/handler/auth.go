package handler

import (
	"myChat/internal/modules/auth/domain"

	"github.com/gin-gonic/gin"
)

func (obj *handler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	err := obj.uc.RegisterUser(ctx, &domain.LogUpRequest{})
	if err != nil {
		return
	}
}
