package auth

import (
	"context"
	"myChat/config"
	"myChat/internal/modules/auth/handler"
	"myChat/internal/modules/auth/usecase"

	"github.com/gin-gonic/gin"
)

type Module struct {
}

func New(gin *gin.Engine, ctx context.Context, cfg *config.Config) {
	uc := usecase.New(cfg)

	handler.Init(gin, uc)
}
