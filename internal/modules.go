package internal

import (
	"context"
	"myChat/config"
	"myChat/internal/modules/auth"

	"github.com/gin-gonic/gin"
)

func InitModules(r *gin.Engine, ctx context.Context, cfg *config.Config) {
	// InitModules
	auth.New(r, ctx, cfg)

}
