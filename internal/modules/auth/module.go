package auth

import (
	"context"
	"myChat/config"
	"myChat/internal/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

type Module struct {
}

func New(gin *gin.Engine, ctx context.Context, cfg *config.Config) {
	handler.Init(gin)
}
