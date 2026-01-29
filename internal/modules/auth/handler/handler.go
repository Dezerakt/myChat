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

	otpAuth := r.Group("/otp")
	{
		otpAuth.POST("/log-in", h.LogIn)
		otpAuth.POST("/register-user", h.LogUp)
	}
}
