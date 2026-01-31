package handler

import (
	"myChat/internal/modules/auth/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	uc *usecase.Usecase
}

func Init(r *gin.Engine, uc *usecase.Usecase) {
	h := &handler{uc: uc}

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
	}
}
